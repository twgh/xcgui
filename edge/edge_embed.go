package edge

import (
	"embed"
	"errors"
	"io/fs"
	"strings"
	"sync"
)

// EmbedFSMapping 是虚拟主机名和嵌入文件系统之间的映射的全局实例.
var EmbedFSMapping = newEmbedFSMapping()

// SetVirtualHostNameToEmbedFSMapping 设置虚拟主机名和嵌入文件系统之间的映射，以便通过该主机名对网站可用.
//   - 设置后这个映射关系是存在于全局的, 所有的 WebView 都可以访问, 这是为了复用嵌入文件.
//   - WebView 要想访问必须先 EnableVirtualHostNameToEmbedFSMapping(true).
//   - 如果想要删除映射关系, 请使用 edge.DeleteVirtualHostNameToEmbedFSMapping(hostName).
//
// hostName: 要映射的主机名。
//
// embedFS: 要映射的文件系统。
//
// dir: 上面嵌入的目录名。例如: "assets". 不填则自动获取.
func SetVirtualHostNameToEmbedFSMapping(hostName string, embedFS embed.FS, dir ...string) error {
	return EmbedFSMapping.Set(hostName, embedFS, dir...)
}

// DeleteVirtualHostNameToEmbedFSMapping 删除指定虚拟主机名和嵌入文件系统之间的映射.
func DeleteVirtualHostNameToEmbedFSMapping(hostName string) {
	EmbedFSMapping.Delete(hostName)
}

// DeleteAllVirtualHostNameToEmbedFSMapping 删除所有虚拟主机名和嵌入文件系统之间的映射.
func DeleteAllVirtualHostNameToEmbedFSMapping() {
	EmbedFSMapping.DeleteAll()
}

// JoinUrlHeader 给虚拟主机名拼接 URL 头.
//   - isHttps 不填时返回: "https://" + hostName, 填 false 时返回: "http://" + hostName.
//   - 默认用 https:// 是因为它的页面属于安全上下文(secure context), 地理位置, 摄像头, 麦克风, 剪贴板
//     等 API 才可用; 而 http:// 的页面不是安全上下文, 页面调用这些 API 会被浏览器直接拒绝.
//   - 注意: https:// 的页面里不能引用 http:// 的绝对地址(会被当成混合内容拦截), 页面内请使用相对地址.
//
// hostName: 虚拟主机名.
//
// isHttps: 是否使用 https, 不填默认为 true.
func JoinUrlHeader(hostName string, isHttps ...bool) string {
	if len(isHttps) > 0 && !isHttps[0] {
		return "http://" + hostName
	}
	return "https://" + hostName
}

// Edge 虚拟主机名和嵌入文件系统之间的映射.
type embedFsMapping struct {
	m    map[string]fs.FS
	lock sync.RWMutex
}

func newEmbedFSMapping() *embedFsMapping {
	return &embedFsMapping{
		m: make(map[string]fs.FS),
	}
}

// Set 设置虚拟主机名和嵌入文件系统之间的映射, 如果已存在会覆盖.
//
// hostName: 要映射的主机名。
//
// embedFS: 要映射的文件系统。
//
// dir: 上面嵌入的目录名。例如: "assets". 不填则自动获取.
func (f *embedFsMapping) Set(hostName string, embedFS embed.FS, dir ...string) error {
	var dirName string
	if len(dir) > 0 { // 使用指定的目录名
		dirName = dir[0]
	} else {
		// 自动检测目录名
		entries, err := embedFS.ReadDir(".")
		if err != nil {
			return errors.New("error reported when automatically obtaining directory name: " + err.Error())
		}
		if len(entries) != 1 || !entries[0].IsDir() {
			return errors.New("automatic retrieval of directory name failed")
		}
		dirName = entries[0].Name()
	}

	subFS, err := fs.Sub(embedFS, dirName)
	if err != nil {
		return errors.New("failed to create sub file system: " + err.Error())
	}

	f.lock.Lock()
	defer f.lock.Unlock()
	f.m[hostName] = subFS
	return nil
}

// Get 获取指定虚拟主机名映射的嵌入文件系统.
func (f *embedFsMapping) Get(hostName string) fs.FS {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.m[hostName]
}

// Delete 删除指定虚拟主机名和嵌入文件系统之间的映射.
func (f *embedFsMapping) Delete(hostName string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	delete(f.m, hostName)
}

// DeleteAll 删除所有虚拟主机名和嵌入文件系统之间的映射.
func (f *embedFsMapping) DeleteAll() {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.m = make(map[string]fs.FS)
}

// Has 检测指定虚拟主机名是否存在.
func (f *embedFsMapping) Has(hostName string) bool {
	f.lock.RLock()
	defer f.lock.RUnlock()
	_, ok := f.m[hostName]
	return ok
}

// Match 匹配指定 URI, 如果匹配到了, 返回 uri 删除了虚拟主机名后的字符串(也就是嵌入文件系统中的文件路径)和嵌入文件系统.
//   - http:// 和 https:// 都可以匹配到.
//   - 如果没有匹配到, 则返回空字符串和 nil.
func (f *embedFsMapping) Match(uri string) (string, fs.FS) {
	for k, v := range f.m {
		for _, scheme := range []string{"http://", "https://"} {
			urlPrefix := scheme + k + "/"
			if strings.HasPrefix(uri, urlPrefix) {
				return strings.TrimPrefix(uri, urlPrefix), v
			}
		}
	}
	return "", nil
}
