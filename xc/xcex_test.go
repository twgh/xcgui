package xc_test

import (
	"log"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
)

func TestCallUTWithAny(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		type Info struct {
			Name string
			Age  int
		}
		data1 := Info{Name: "张三", Age: 18}
		data2 := Info{Name: "李四", Age: 20}

		btn := widget.NewButton(50, 50, 100, 30, "按钮", w.Handle)
		btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			go func() {
				xc.CallUTAny(func(data ...interface{}) int {
					t.Log("不传参数:", data)
					return 0
				})

				xc.CallUTAny(func(data ...interface{}) int {
					t.Log("1个参数:", data)
					return 0
				}, data1)

				xc.AutoAny(func(data ...interface{}) int {
					t.Log("2个参数:", data)
					return 0
				}, data2, "女")

				xc.CallUTAny(func(data ...interface{}) int {
					t.Log("3个参数:", data)
					return 0
				}, data1, "女", "上海")
			}()
			return 0
		})
	})
}

func TestCallUTAny(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		const goroutineCount = 3000

		var (
			wg           sync.WaitGroup
			successCount atomic.Int32
		)

		go func() {
			time.Sleep(time.Second * 1)
			for i := 0; i < goroutineCount; i++ {
				wg.Add(1)
				go func(id int) {
					defer wg.Done()
					xc.CallUTAny(func(data ...interface{}) int {
						if len(data) > 0 {
							if v, ok := data[0].(int); ok && v == id {
								successCount.Add(1)
							}
						}
						return 0
					}, id)
				}(i)
			}

			wg.Wait()
			time.Sleep(500 * time.Millisecond)

			success := successCount.Load()
			log.Printf("成功执行次数: %d / %d", success, goroutineCount)

			if success != goroutineCount {
				log.Printf("期望成功 %d 次, 实际成功 %d 次", goroutineCount, success)
			}
		}()
	})
}

func TestCallUT(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		const goroutineCount = 10

		var (
			wg           sync.WaitGroup
			successCount atomic.Int32
		)

		go func() {
			time.Sleep(time.Second * 1)
			for i := 0; i < goroutineCount; i++ {
				wg.Add(1)

				go func(id int) {
					defer wg.Done()
					xc.XC_CallUT(func() {
						// 在 UI 线程执行
						successCount.Add(1)

						log.Println(id)
						time.Sleep(time.Millisecond * 100 * time.Duration(id))
					})
				}(i)
			}

			wg.Wait()
			time.Sleep(500 * time.Millisecond)

			success := successCount.Load()
			log.Printf("成功执行次数: %d / %d", success, goroutineCount)

			if success != goroutineCount {
				log.Printf("期望成功 %d 次, 实际成功 %d 次", goroutineCount, success)
			}
		}()
	})
}

func TestCallUiThreadEx(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		const goroutineCount = 1000
		const callCountPerGoroutine = 5

		var (
			wg           sync.WaitGroup
			successCount atomic.Int32
			resultMap    sync.Map
		)

		go func() {
			time.Sleep(time.Second * 1)
			// 启动多个 goroutine 并发调用 XC_CallUiThreadEx
			for i := 0; i < goroutineCount; i++ {
				wg.Add(1)
				go func(goroutineID int) {
					defer wg.Done()
					for j := 0; j < callCountPerGoroutine; j++ {
						callID := goroutineID*1000 + j
						expectedData := callID * 10

						xc.XC_CallUiThreadEx(func(data int) int {
							// 验证回调接收到的数据是否正确
							if data == expectedData {
								successCount.Add(1)
								resultMap.Store(callID, true)
							} else {
								log.Printf("数据不匹配: 期望 %d, 实际 %d", expectedData, data)
								resultMap.Store(callID, false)
							}
							return 0
						}, expectedData)
					}
				}(i)
			}

			// 等待所有 goroutine 完成调用
			wg.Wait()

			// 等待 UI 线程执行完所有回调
			time.Sleep(500 * time.Millisecond)

			// 验证结果
			totalCalls := goroutineCount * callCountPerGoroutine
			success := successCount.Load()
			log.Printf("总调用次数: %d, 成功次数: %d", totalCalls, success)

			if success != int32(totalCalls) {
				log.Printf("期望成功 %d 次, 实际成功 %d 次", totalCalls, success)
			}
		}()
	})
}

// TestXC_CallUiThreadEx_Recursive 测试回调中再次调用 XC_CallUiThreadEx (死锁测试)
func TestXC_CallUiThreadEx_Recursive(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		var (
			wg           sync.WaitGroup
			done         = make(chan struct{})
			successCount atomic.Int32
		)

		wg.Add(1)

		go func() {
			time.Sleep(time.Second * 1)
			// 从非 UI 线程调用
			go func() {
				defer wg.Done()
				xc.XC_CallUiThreadEx(func(data int) int {
					t.Log("第一次调用")
					// 在 UI 线程执行，再次调用 XC_CallUiThreadEx
					xc.XC_CallUiThreadEx(func(data2 int) int {
						t.Log("第二次调用")
						successCount.Add(1)
						close(done)
						return 0
					}, 0)
					return 0
				}, 0)
			}()

			// 等待完成或超时
			select {
			case <-done:
				t.Log("递归调用成功，无死锁")
			case <-time.After(2 * time.Second):
				t.Log("超时！可能发生死锁")
				os.Exit(1)
			}

			wg.Wait()

			if successCount.Load() != 1 {
				t.Error("回调未正确执行")
			}
		}()
	})
}

// TestXC_CallUT_Recursive 测试回调中再次调用 XC_CallUT (死锁测试)
func TestXC_CallUT_Recursive(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		var (
			done         = make(chan struct{})
			successCount atomic.Int32
		)

		go func() {
			time.Sleep(time.Second * 1)

			go func() {
				xc.XC_CallUT(func() {
					t.Log("第一次调用")
					// 在 UI 线程执行，再次调用 XC_CallUT
					xc.XC_CallUT(func() {
						t.Log("第二次调用")
						successCount.Add(1)
						close(done)
					})
				})
			}()

			select {
			case <-done:
				t.Log("递归调用成功，无死锁")
			case <-time.After(2 * time.Second):
				t.Log("超时！可能发生死锁")
				os.Exit(1)
			}

			if successCount.Load() != 1 {
				t.Error("回调未正确执行")
			}
		}()
	})
}
