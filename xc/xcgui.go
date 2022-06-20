package xc

import (
	"errors"
	"os"
	"strconv"
	"syscall"
)

// xcguiPath 是xcgui.dll的路径(不是目录, 是文件名), 默认值为'xcgui.dll'.
//	如果你想要更改它的位置, 可以在 xc.LoadXCGUI() 之前调用 xc.SetXcguiPath() 更改为其他路径.
var xcguiPath = "xcgui.dll"

var (
	// Library.
	xcgui *syscall.LazyDLL

	// Global Functions.
	xC_wtoa      *syscall.LazyProc
	xC_atow      *syscall.LazyProc
	xC_utf8tow   *syscall.LazyProc
	xC_utf8towEx *syscall.LazyProc
	xC_utf8toa   *syscall.LazyProc
	xC_atoutf8   *syscall.LazyProc
	xC_wtoutf8   *syscall.LazyProc
	xC_wtoutf8Ex *syscall.LazyProc
	/*
			xC_itoa                       *syscall.LazyProc
			xC_itow                       *syscall.LazyProc
			xC_ftoa                       *syscall.LazyProc
			xC_ftow                       *syscall.LazyProc
		 	xC_UnicodeToAnsi *syscall.LazyProc
		   	xC_AnsiToUnicode *syscall.LazyProc
			xDebug_OutputDebugStringA					*syscall.LazyProc
	*/
	xC_MessageBox                     *syscall.LazyProc
	xC_SendMessage                    *syscall.LazyProc
	xC_PostMessage                    *syscall.LazyProc
	xC_CallUiThread                   *syscall.LazyProc
	xDebug_OutputDebugStringW         *syscall.LazyProc
	xC_DebugToFileInfo                *syscall.LazyProc
	xDebug_Set_OutputDebugString_UTF8 *syscall.LazyProc
	xDebug_Print                      *syscall.LazyProc
	xC_IsHELE                         *syscall.LazyProc
	xC_IsHWINDOW                      *syscall.LazyProc
	xC_IsShape                        *syscall.LazyProc
	xC_IsHXCGUI                       *syscall.LazyProc
	xC_hWindowFromHWnd                *syscall.LazyProc
	xC_SetActivateTopWindow           *syscall.LazyProc
	xC_SetProperty                    *syscall.LazyProc
	xC_GetProperty                    *syscall.LazyProc
	xC_RegisterWindowClassName        *syscall.LazyProc
	xC_IsSViewExtend                  *syscall.LazyProc
	xC_GetObjectType                  *syscall.LazyProc
	xC_GetObjectByID                  *syscall.LazyProc
	xC_GetObjectByIDName              *syscall.LazyProc
	xC_GetObjectByUID                 *syscall.LazyProc
	xC_GetObjectByUIDName             *syscall.LazyProc
	xC_GetObjectByName                *syscall.LazyProc
	xC_SetPaintFrequency              *syscall.LazyProc
	xC_SetTextRenderingHint           *syscall.LazyProc
	xC_EnableGdiDrawText              *syscall.LazyProc
	xC_RectInRect                     *syscall.LazyProc
	xC_CombineRect                    *syscall.LazyProc
	xC_ShowLayoutFrame                *syscall.LazyProc
	xC_EnableDebugFile                *syscall.LazyProc
	xC_EnableResMonitor               *syscall.LazyProc
	xC_SetLayoutFrameColor            *syscall.LazyProc
	xC_EnableErrorMessageBox          *syscall.LazyProc
	xC_EnableAutoExitApp              *syscall.LazyProc
	xC_GetTextSize                    *syscall.LazyProc
	xC_GetTextShowSize                *syscall.LazyProc
	xC_GetTextShowSizeEx              *syscall.LazyProc
	xC_GetTextShowRect                *syscall.LazyProc
	xC_GetDefaultFont                 *syscall.LazyProc
	xC_SetDefaultFont                 *syscall.LazyProc
	xC_AddFileSearchPath              *syscall.LazyProc
	xC_InitFont                       *syscall.LazyProc
	xC_Malloc                         *syscall.LazyProc
	xC_Free                           *syscall.LazyProc
	xC_Alert                          *syscall.LazyProc
	xC_Sys_ShellExecute               *syscall.LazyProc
	xC_LoadLibrary                    *syscall.LazyProc
	xC_GetProcAddress                 *syscall.LazyProc
	xC_FreeLibrary                    *syscall.LazyProc
	xC_LoadDll                        *syscall.LazyProc
	xInitXCGUI                        *syscall.LazyProc
	xRunXCGUI                         *syscall.LazyProc
	xExitXCGUI                        *syscall.LazyProc
	xC_PostQuitMessage                *syscall.LazyProc
	xC_GetD2dFactory                  *syscall.LazyProc
	xC_GetWicFactory                  *syscall.LazyProc
	xC_GetDWriteFactory               *syscall.LazyProc
	xC_SetD2dTextRenderingMode        *syscall.LazyProc
	xC_IsEnableD2D                    *syscall.LazyProc
	xMsg_Create                       *syscall.LazyProc
	xMsg_CreateEx                     *syscall.LazyProc
	xC_ShowSvgFrame                   *syscall.LazyProc

	// UI Designer.
	xC_LoadLayout       *syscall.LazyProc
	xC_LoadLayoutZip    *syscall.LazyProc
	xC_LoadLayoutZipMem *syscall.LazyProc
	// xC_LoadLayoutFromString     *syscall.LazyProc
	xC_LoadLayoutFromStringUtf8   *syscall.LazyProc
	xC_LoadStyle                  *syscall.LazyProc
	xC_LoadStyleZip               *syscall.LazyProc
	xC_LoadStyleZipMem            *syscall.LazyProc
	xC_LoadResource               *syscall.LazyProc
	xC_LoadResourceZip            *syscall.LazyProc
	xC_LoadResourceZipMem         *syscall.LazyProc
	xC_LoadResourceFromStringUtf8 *syscall.LazyProc
	xC_LoadStyleFromStringW       *syscall.LazyProc

	/* 	xC_LoadResourceFromString     *syscall.LazyProc
	   	xC_LoadStyleFromString        *syscall.LazyProc */

	// Window.
	xWnd_Create                    *syscall.LazyProc
	xWnd_CreateEx                  *syscall.LazyProc
	xWnd_RegEventC                 *syscall.LazyProc
	xWnd_RegEventC1                *syscall.LazyProc
	xWnd_RemoveEventC              *syscall.LazyProc
	xWnd_ShowWindow                *syscall.LazyProc
	xWnd_SetTop                    *syscall.LazyProc
	xWnd_AddChild                  *syscall.LazyProc
	xWnd_InsertChild               *syscall.LazyProc
	xWnd_Redraw                    *syscall.LazyProc
	xWnd_RedrawRect                *syscall.LazyProc
	xWnd_SetFocusEle               *syscall.LazyProc
	xWnd_GetFocusEle               *syscall.LazyProc
	xWnd_GetStayEle                *syscall.LazyProc
	xWnd_DrawWindow                *syscall.LazyProc
	xWnd_Center                    *syscall.LazyProc
	xWnd_CenterEx                  *syscall.LazyProc
	xWnd_SetCursor                 *syscall.LazyProc
	xWnd_GetCursor                 *syscall.LazyProc
	xWnd_GetHWND                   *syscall.LazyProc
	xWnd_EnableDragBorder          *syscall.LazyProc
	xWnd_EnableDragWindow          *syscall.LazyProc
	xWnd_EnableDragCaption         *syscall.LazyProc
	xWnd_EnableDrawBk              *syscall.LazyProc
	xWnd_EnableAutoFocus           *syscall.LazyProc
	xWnd_EnableMaxWindow           *syscall.LazyProc
	xWnd_EnablemLimitWindowSize    *syscall.LazyProc
	xWnd_EnableLayout              *syscall.LazyProc
	xWnd_EnableLayoutOverlayBorder *syscall.LazyProc
	xWnd_ShowLayoutFrame           *syscall.LazyProc
	xWnd_IsEnableLayout            *syscall.LazyProc
	xWnd_IsMaxWindow               *syscall.LazyProc
	xWnd_SetCaptureEle             *syscall.LazyProc
	xWnd_GetCaptureEle             *syscall.LazyProc
	xWnd_GetDrawRect               *syscall.LazyProc
	xWnd_SetCursorSys              *syscall.LazyProc
	xWnd_SetFont                   *syscall.LazyProc
	xWnd_SetTextColor              *syscall.LazyProc
	xWnd_GetTextColor              *syscall.LazyProc
	xWnd_GetTextColorEx            *syscall.LazyProc
	xWnd_SetID                     *syscall.LazyProc
	xWnd_GetID                     *syscall.LazyProc
	xWnd_SetName                   *syscall.LazyProc
	xWnd_GetName                   *syscall.LazyProc
	xWnd_SetBorderSize             *syscall.LazyProc
	xWnd_GetBorderSize             *syscall.LazyProc
	xWnd_SetPadding                *syscall.LazyProc
	xWnd_SetDragBorderSize         *syscall.LazyProc
	xWnd_GetDragBorderSize         *syscall.LazyProc
	xWnd_SetMinimumSize            *syscall.LazyProc
	xWnd_HitChildEle               *syscall.LazyProc
	xWnd_GetChildCount             *syscall.LazyProc
	xWnd_GetChildByIndex           *syscall.LazyProc
	xWnd_GetChildByID              *syscall.LazyProc
	xWnd_GetChild                  *syscall.LazyProc
	xWnd_CloseWindow               *syscall.LazyProc
	xWnd_AdjustLayout              *syscall.LazyProc
	xWnd_AdjustLayoutEx            *syscall.LazyProc
	xWnd_CreateCaret               *syscall.LazyProc
	xWnd_SetCaretPos               *syscall.LazyProc
	xWnd_SetCaretColor             *syscall.LazyProc
	xWnd_ShowCaret                 *syscall.LazyProc
	xWnd_DestroyCaret              *syscall.LazyProc
	xWnd_GetCaretHELE              *syscall.LazyProc
	xWnd_GetClientRect             *syscall.LazyProc
	xWnd_GetBodyRect               *syscall.LazyProc
	xWnd_GetLayoutRect             *syscall.LazyProc
	xWnd_SetPosition               *syscall.LazyProc
	xWnd_GetRect                   *syscall.LazyProc
	xWnd_SetRect                   *syscall.LazyProc
	xWnd_MaxWindow                 *syscall.LazyProc
	xWnd_SetTimer                  *syscall.LazyProc
	xWnd_KillTimer                 *syscall.LazyProc
	xWnd_SetXCTimer                *syscall.LazyProc
	xWnd_KillXCTimer               *syscall.LazyProc
	xWnd_GetBkManager              *syscall.LazyProc
	xWnd_GetBkManagerEx            *syscall.LazyProc
	xWnd_SetBkMagager              *syscall.LazyProc
	xWnd_SetTransparentType        *syscall.LazyProc
	xWnd_SetTransparentAlpha       *syscall.LazyProc
	xWnd_SetTransparentColor       *syscall.LazyProc
	xWnd_SetShadowInfo             *syscall.LazyProc
	xWnd_GetShadowInfo             *syscall.LazyProc
	xWnd_GetTransparentType        *syscall.LazyProc
	xWnd_Attach                    *syscall.LazyProc
	xWnd_EnableDragFiles           *syscall.LazyProc
	xWnd_Show                      *syscall.LazyProc
	xWnd_GetCaretInfo              *syscall.LazyProc
	xWnd_SetIcon                   *syscall.LazyProc
	xWnd_SetTitle                  *syscall.LazyProc
	xWnd_SetTitleColor             *syscall.LazyProc
	xWnd_GetButton                 *syscall.LazyProc
	xWnd_GetIcon                   *syscall.LazyProc
	xWnd_GetTitle                  *syscall.LazyProc
	xWnd_GetTitleColor             *syscall.LazyProc
	xWnd_AddBkBorder               *syscall.LazyProc
	xWnd_AddBkFill                 *syscall.LazyProc
	xWnd_AddBkImage                *syscall.LazyProc
	xWnd_GetBkInfoCount            *syscall.LazyProc
	xWnd_ClearBkInfo               *syscall.LazyProc
	xWnd_SetBkInfo                 *syscall.LazyProc
	xWnd_IsDragCaption             *syscall.LazyProc
	xWnd_IsDragWindow              *syscall.LazyProc
	xWnd_IsDragBorder              *syscall.LazyProc
	xWnd_SetCaptionMargin          *syscall.LazyProc

	// Widget.
	xWidget_IsShow                 *syscall.LazyProc
	xWidget_Show                   *syscall.LazyProc
	xWidget_EnableLayoutControl    *syscall.LazyProc
	xWidget_IsLayoutControl        *syscall.LazyProc
	xWidget_GetParentEle           *syscall.LazyProc
	xWidget_GetParent              *syscall.LazyProc
	xWidget_GetHWND                *syscall.LazyProc
	xWidget_GetHWINDOW             *syscall.LazyProc
	xWidget_LayoutItem_EnableWrap  *syscall.LazyProc
	xWidget_LayoutItem_EnableSwap  *syscall.LazyProc
	xWidget_LayoutItem_EnableFloat *syscall.LazyProc
	xWidget_LayoutItem_SetWidth    *syscall.LazyProc
	xWidget_LayoutItem_SetHeight   *syscall.LazyProc
	xWidget_LayoutItem_GetWidth    *syscall.LazyProc
	xWidget_LayoutItem_GetHeight   *syscall.LazyProc
	xWidget_LayoutItem_SetAlign    *syscall.LazyProc
	xWidget_LayoutItem_SetMargin   *syscall.LazyProc
	xWidget_LayoutItem_GetMargin   *syscall.LazyProc
	xWidget_LayoutItem_SetMinSize  *syscall.LazyProc
	xWidget_LayoutItem_SetPosition *syscall.LazyProc
	xWidget_SetID                  *syscall.LazyProc
	xWidget_GetID                  *syscall.LazyProc
	xWidget_SetUID                 *syscall.LazyProc
	xWidget_GetUID                 *syscall.LazyProc
	xWidget_SetName                *syscall.LazyProc
	xWidget_GetName                *syscall.LazyProc
	// xUI.
	xUI_SetStyle   *syscall.LazyProc
	xUI_GetStyle   *syscall.LazyProc
	xUI_EnableCSS  *syscall.LazyProc
	xUI_SetCssName *syscall.LazyProc
	xUI_GetCssName *syscall.LazyProc
	// Button.
	xBtn_Create            *syscall.LazyProc
	xBtn_IsCheck           *syscall.LazyProc
	xBtn_SetCheck          *syscall.LazyProc
	xBtn_SetState          *syscall.LazyProc
	xBtn_GetState          *syscall.LazyProc
	xBtn_GetStateEx        *syscall.LazyProc
	xBtn_SetTypeEx         *syscall.LazyProc
	xBtn_SetGroupID        *syscall.LazyProc
	xBtn_GetGroupID        *syscall.LazyProc
	xBtn_SetBindEle        *syscall.LazyProc
	xBtn_GetBindEle        *syscall.LazyProc
	xBtn_SetTextAlign      *syscall.LazyProc
	xBtn_GetTextAlign      *syscall.LazyProc
	xBtn_SetIconAlign      *syscall.LazyProc
	xBtn_SetOffset         *syscall.LazyProc
	xBtn_SetOffsetIcon     *syscall.LazyProc
	xBtn_SetIconSpace      *syscall.LazyProc
	xBtn_SetText           *syscall.LazyProc
	xBtn_GetText           *syscall.LazyProc
	xBtn_SetIcon           *syscall.LazyProc
	xBtn_SetIconDisable    *syscall.LazyProc
	xBtn_GetIcon           *syscall.LazyProc
	xBtn_AddAnimationFrame *syscall.LazyProc
	xBtn_EnableAnimation   *syscall.LazyProc

	// Element.
	xEle_Create                     *syscall.LazyProc
	xEle_RegEventC                  *syscall.LazyProc
	xEle_RegEventC1                 *syscall.LazyProc
	xEle_RemoveEventC               *syscall.LazyProc
	xEle_SendEvent                  *syscall.LazyProc
	xEle_PostEvent                  *syscall.LazyProc
	xEle_GetRect                    *syscall.LazyProc
	xEle_GetRectLogic               *syscall.LazyProc
	xEle_GetClientRect              *syscall.LazyProc
	xEle_SetWidth                   *syscall.LazyProc
	xEle_SetHeight                  *syscall.LazyProc
	xEle_GetWidth                   *syscall.LazyProc
	xEle_GetHeight                  *syscall.LazyProc
	xEle_RectWndClientToEleClient   *syscall.LazyProc
	xEle_PointWndClientToEleClient  *syscall.LazyProc
	xEle_RectClientToWndClient      *syscall.LazyProc
	xEle_PointClientToWndClient     *syscall.LazyProc
	xEle_GetWndClientRect           *syscall.LazyProc
	xEle_GetCursor                  *syscall.LazyProc
	xEle_SetCursor                  *syscall.LazyProc
	xEle_AddChild                   *syscall.LazyProc
	xEle_InsertChild                *syscall.LazyProc
	xEle_SetRect                    *syscall.LazyProc
	xEle_SetRectEx                  *syscall.LazyProc
	xEle_SetRectLogic               *syscall.LazyProc
	xEle_SetPosition                *syscall.LazyProc
	xEle_SetPositionLogic           *syscall.LazyProc
	xEle_IsDrawFocus                *syscall.LazyProc
	xEle_IsEnable                   *syscall.LazyProc
	xEle_IsEnableFocus              *syscall.LazyProc
	xEle_IsMouseThrough             *syscall.LazyProc
	xEle_HitChildEle                *syscall.LazyProc
	xEle_IsBkTransparent            *syscall.LazyProc
	xEle_IsEnableEvent_XE_PAINT_END *syscall.LazyProc
	xEle_IsKeyTab                   *syscall.LazyProc
	xEle_IsSwitchFocus              *syscall.LazyProc
	xEle_IsEnable_XE_MOUSEWHEEL     *syscall.LazyProc
	xEle_IsChildEle                 *syscall.LazyProc
	xEle_IsEnableCanvas             *syscall.LazyProc
	xEle_IsFocus                    *syscall.LazyProc
	xEle_IsFocusEx                  *syscall.LazyProc
	xEle_Enable                     *syscall.LazyProc
	xEle_EnableFocus                *syscall.LazyProc
	xEle_EnableDrawFocus            *syscall.LazyProc
	xEle_EnableDrawBorder           *syscall.LazyProc
	xEle_EnableCanvas               *syscall.LazyProc
	xEle_EnableEvent_XE_PAINT_END   *syscall.LazyProc
	xEle_EnableBkTransparent        *syscall.LazyProc
	xEle_EnableMouseThrough         *syscall.LazyProc
	xEle_EnableKeyTab               *syscall.LazyProc
	xEle_EnableSwitchFocus          *syscall.LazyProc
	xEle_EnableEvent_XE_MOUSEWHEEL  *syscall.LazyProc
	xEle_Remove                     *syscall.LazyProc
	xEle_SetZOrder                  *syscall.LazyProc
	xEle_SetZOrderEx                *syscall.LazyProc
	xEle_GetZOrder                  *syscall.LazyProc
	xEle_EnableTopmost              *syscall.LazyProc
	xEle_Redraw                     *syscall.LazyProc
	xEle_RedrawRect                 *syscall.LazyProc
	xEle_GetChildCount              *syscall.LazyProc
	xEle_GetChildByIndex            *syscall.LazyProc
	xEle_GetChildByID               *syscall.LazyProc
	xEle_SetBorderSize              *syscall.LazyProc
	xEle_GetBorderSize              *syscall.LazyProc
	xEle_SetPadding                 *syscall.LazyProc
	xEle_GetPadding                 *syscall.LazyProc
	xEle_SetDragBorder              *syscall.LazyProc
	xEle_SetDragBorderBindEle       *syscall.LazyProc
	xEle_SetMinSize                 *syscall.LazyProc
	xEle_SetMaxSize                 *syscall.LazyProc
	xEle_SetLockScroll              *syscall.LazyProc
	xEle_SetTextColor               *syscall.LazyProc
	xEle_GetTextColor               *syscall.LazyProc
	xEle_GetTextColorEx             *syscall.LazyProc
	xEle_SetFocusBorderColor        *syscall.LazyProc
	xEle_GetFocusBorderColor        *syscall.LazyProc
	xEle_SetFont                    *syscall.LazyProc
	xEle_GetFont                    *syscall.LazyProc
	xEle_GetFontEx                  *syscall.LazyProc
	xEle_SetAlpha                   *syscall.LazyProc
	xEle_Destroy                    *syscall.LazyProc
	xEle_AddBkBorder                *syscall.LazyProc
	xEle_AddBkFill                  *syscall.LazyProc
	xEle_AddBkImage                 *syscall.LazyProc
	xEle_GetBkInfoCount             *syscall.LazyProc
	xEle_ClearBkInfo                *syscall.LazyProc
	xEle_GetBkManager               *syscall.LazyProc
	xEle_GetBkManagerEx             *syscall.LazyProc
	xEle_SetBkManager               *syscall.LazyProc
	xEle_GetStateFlags              *syscall.LazyProc
	xEle_DrawFocus                  *syscall.LazyProc
	xEle_DrawEle                    *syscall.LazyProc
	xEle_SetUserData                *syscall.LazyProc
	xEle_GetUserData                *syscall.LazyProc
	xEle_GetContentSize             *syscall.LazyProc
	xEle_SetCapture                 *syscall.LazyProc
	xEle_EnableTransparentChannel   *syscall.LazyProc
	xEle_SetXCTimer                 *syscall.LazyProc
	xEle_KillXCTimer                *syscall.LazyProc
	xEle_SetToolTip                 *syscall.LazyProc
	xEle_SetToolTipEx               *syscall.LazyProc
	xEle_GetToolTip                 *syscall.LazyProc
	xEle_PopupToolTip               *syscall.LazyProc
	xEle_AdjustLayout               *syscall.LazyProc
	xEle_AdjustLayoutEx             *syscall.LazyProc
	xEle_GetAlpha                   *syscall.LazyProc
	xEle_GetPosition                *syscall.LazyProc
	xEle_SetSize                    *syscall.LazyProc
	xEle_GetSize                    *syscall.LazyProc
	xEle_SetBkInfo                  *syscall.LazyProc

	// FreameWindow.
	xFrameWnd_Create                 *syscall.LazyProc
	xFrameWnd_CreateEx               *syscall.LazyProc
	xFrameWnd_GetLayoutAreaRect      *syscall.LazyProc
	xFrameWnd_SetView                *syscall.LazyProc
	xFrameWnd_SetPaneSplitBarColor   *syscall.LazyProc
	xFrameWnd_SetTabBarHeight        *syscall.LazyProc
	xFrameWnd_SaveLayoutToFile       *syscall.LazyProc
	xFrameWnd_LoadLayoutFile         *syscall.LazyProc
	xFrameWnd_AddPane                *syscall.LazyProc
	xFrameWnd_MergePane              *syscall.LazyProc
	xFrameWnd_Attach                 *syscall.LazyProc
	xFrameWnd_GetDragFloatWndTopFlag *syscall.LazyProc

	// Menu.
	xMenu_Create               *syscall.LazyProc
	xMenu_AddItem              *syscall.LazyProc
	xMenu_AddItemIcon          *syscall.LazyProc
	xMenu_InsertItem           *syscall.LazyProc
	xMenu_InsertItemIcon       *syscall.LazyProc
	xMenu_GetFirstChildItem    *syscall.LazyProc
	xMenu_GetEndChildItem      *syscall.LazyProc
	xMenu_GetPrevSiblingItem   *syscall.LazyProc
	xMenu_GetNextSiblingItem   *syscall.LazyProc
	xMenu_GetParentItem        *syscall.LazyProc
	xMenu_SetAutoDestroy       *syscall.LazyProc
	xMenu_EnableDrawBackground *syscall.LazyProc
	xMenu_EnableDrawItem       *syscall.LazyProc
	xMenu_Popup                *syscall.LazyProc
	xMenu_DestroyMenu          *syscall.LazyProc
	xMenu_CloseMenu            *syscall.LazyProc
	xMenu_SetBkImage           *syscall.LazyProc
	xMenu_SetItemText          *syscall.LazyProc
	xMenu_GetItemText          *syscall.LazyProc
	xMenu_GetItemTextLength    *syscall.LazyProc
	xMenu_SetItemIcon          *syscall.LazyProc
	xMenu_SetItemFlags         *syscall.LazyProc
	xMenu_SetItemHeight        *syscall.LazyProc
	xMenu_GetItemHeight        *syscall.LazyProc
	xMenu_SetBorderColor       *syscall.LazyProc
	xMenu_SetBorderSize        *syscall.LazyProc
	xMenu_GetLeftWidth         *syscall.LazyProc
	xMenu_GetLeftSpaceText     *syscall.LazyProc
	xMenu_GetItemCount         *syscall.LazyProc
	xMenu_SetItemCheck         *syscall.LazyProc
	xMenu_IsItemCheck          *syscall.LazyProc
	xMenu_SetItemWidth         *syscall.LazyProc

	// ModalWindow.
	xModalWnd_Create          *syscall.LazyProc
	xModalWnd_CreateEx        *syscall.LazyProc
	xModalWnd_EnableAutoClose *syscall.LazyProc
	xModalWnd_EnableEscClose  *syscall.LazyProc
	xModalWnd_DoModal         *syscall.LazyProc
	xModalWnd_EndModal        *syscall.LazyProc
	xModalWnd_Attach          *syscall.LazyProc

	// LayoutBox.
	xLayoutBox_EnableHorizon      *syscall.LazyProc
	xLayoutBox_EnableAutoWrap     *syscall.LazyProc
	xLayoutBox_EnableOverflowHide *syscall.LazyProc
	xLayoutBox_SetAlignH          *syscall.LazyProc
	xLayoutBox_SetAlignV          *syscall.LazyProc
	xLayoutBox_SetAlignBaseline   *syscall.LazyProc
	xLayoutBox_SetSpace           *syscall.LazyProc
	xLayoutBox_SetSpaceRow        *syscall.LazyProc

	// ComboBox.
	xComboBox_Create                       *syscall.LazyProc
	xComboBox_SetSelItem                   *syscall.LazyProc
	xComboBox_CreateAdapter                *syscall.LazyProc
	xComboBox_BindAdapter                  *syscall.LazyProc
	xComboBox_GetAdapter                   *syscall.LazyProc
	xComboBox_SetBindName                  *syscall.LazyProc
	xComboBox_GetButtonRect                *syscall.LazyProc
	xComboBox_SetButtonSize                *syscall.LazyProc
	xComboBox_SetDropHeight                *syscall.LazyProc
	xComboBox_GetDropHeight                *syscall.LazyProc
	xComboBox_SetItemTemplateXML           *syscall.LazyProc
	xComboBox_SetItemTemplateXMLFromString *syscall.LazyProc
	xComboBox_EnableDrawButton             *syscall.LazyProc
	xComboBox_EnableEdit                   *syscall.LazyProc
	xComboBox_EnableDropHeightFixed        *syscall.LazyProc
	xComboBox_GetSelItem                   *syscall.LazyProc
	xComboBox_GetState                     *syscall.LazyProc
	xComboBox_AddItemText                  *syscall.LazyProc
	xComboBox_AddItemTextEx                *syscall.LazyProc
	xComboBox_AddItemImage                 *syscall.LazyProc
	xComboBox_AddItemImageEx               *syscall.LazyProc
	xComboBox_InsertItemText               *syscall.LazyProc
	xComboBox_InsertItemTextEx             *syscall.LazyProc
	xComboBox_InsertItemImage              *syscall.LazyProc
	xComboBox_InsertItemImageEx            *syscall.LazyProc
	xComboBox_SetItemText                  *syscall.LazyProc
	xComboBox_SetItemTextEx                *syscall.LazyProc
	xComboBox_SetItemImage                 *syscall.LazyProc
	xComboBox_SetItemImageEx               *syscall.LazyProc
	xComboBox_SetItemInt                   *syscall.LazyProc
	xComboBox_SetItemIntEx                 *syscall.LazyProc
	xComboBox_SetItemFloat                 *syscall.LazyProc
	xComboBox_SetItemFloatEx               *syscall.LazyProc
	xComboBox_GetItemText                  *syscall.LazyProc
	xComboBox_GetItemTextEx                *syscall.LazyProc
	xComboBox_GetItemImage                 *syscall.LazyProc
	xComboBox_GetItemImageEx               *syscall.LazyProc
	xComboBox_GetItemInt                   *syscall.LazyProc
	xComboBox_GetItemIntEx                 *syscall.LazyProc
	xComboBox_GetItemFloat                 *syscall.LazyProc
	xComboBox_GetItemFloatEx               *syscall.LazyProc
	xComboBox_DeleteItem                   *syscall.LazyProc
	xComboBox_DeleteItemEx                 *syscall.LazyProc
	xComboBox_DeleteItemAll                *syscall.LazyProc
	xComboBox_DeleteColumnAll              *syscall.LazyProc
	xComboBox_GetCount                     *syscall.LazyProc
	xComboBox_GetCountColumn               *syscall.LazyProc
	xComboBox_PopupDropList                *syscall.LazyProc
	xComboBox_SetItemTemplate              *syscall.LazyProc

	// Adapter.
	xAd_AddRef            *syscall.LazyProc
	xAd_Release           *syscall.LazyProc
	xAd_GetRefCount       *syscall.LazyProc
	xAd_Destroy           *syscall.LazyProc
	xAd_EnableAutoDestroy *syscall.LazyProc

	// AdapterListView.
	xAdListView_Create                   *syscall.LazyProc
	xAdListView_Group_AddColumn          *syscall.LazyProc
	xAdListView_Group_AddItemText        *syscall.LazyProc
	xAdListView_Group_AddItemTextEx      *syscall.LazyProc
	xAdListView_Group_AddItemImage       *syscall.LazyProc
	xAdListView_Group_AddItemImageEx     *syscall.LazyProc
	xAdListView_Group_SetText            *syscall.LazyProc
	xAdListView_Group_SetTextEx          *syscall.LazyProc
	xAdListView_Group_SetImage           *syscall.LazyProc
	xAdListView_Group_SetImageEx         *syscall.LazyProc
	xAdListView_Item_AddColumn           *syscall.LazyProc
	xAdListView_Group_GetCount           *syscall.LazyProc
	xAdListView_Item_GetCount            *syscall.LazyProc
	xAdListView_Item_AddItemText         *syscall.LazyProc
	xAdListView_Item_AddItemTextEx       *syscall.LazyProc
	xAdListView_Item_AddItemImage        *syscall.LazyProc
	xAdListView_Item_AddItemImageEx      *syscall.LazyProc
	xAdListView_Group_DeleteItem         *syscall.LazyProc
	xAdListView_Group_DeleteAllChildItem *syscall.LazyProc
	xAdListView_Item_DeleteItem          *syscall.LazyProc
	xAdListView_DeleteAll                *syscall.LazyProc
	xAdListView_DeleteAllGroup           *syscall.LazyProc
	xAdListView_DeleteAllItem            *syscall.LazyProc
	xAdListView_DeleteColumnGroup        *syscall.LazyProc
	xAdListView_DeleteColumnItem         *syscall.LazyProc
	xAdListView_Item_GetTextEx           *syscall.LazyProc
	xAdListView_Item_GetImageEx          *syscall.LazyProc
	xAdListView_Item_SetText             *syscall.LazyProc
	xAdListView_Item_SetTextEx           *syscall.LazyProc
	xAdListView_Item_SetImage            *syscall.LazyProc
	xAdListView_Item_SetImageEx          *syscall.LazyProc
	xAdListView_Group_GetText            *syscall.LazyProc
	xAdListView_Group_GetTextEx          *syscall.LazyProc
	xAdListView_Group_GetImage           *syscall.LazyProc
	xAdListView_Group_GetImageEx         *syscall.LazyProc
	xAdListView_Item_GetText             *syscall.LazyProc
	xAdListView_Item_GetImage            *syscall.LazyProc

	// AdapterTable.
	xAdTable_Create            *syscall.LazyProc
	xAdTable_Sort              *syscall.LazyProc
	xAdTable_GetItemDataType   *syscall.LazyProc
	xAdTable_GetItemDataTypeEx *syscall.LazyProc
	xAdTable_AddColumn         *syscall.LazyProc
	xAdTable_SetColumn         *syscall.LazyProc
	xAdTable_AddItemText       *syscall.LazyProc
	xAdTable_AddItemTextEx     *syscall.LazyProc
	xAdTable_AddItemImage      *syscall.LazyProc
	xAdTable_AddItemImageEx    *syscall.LazyProc
	xAdTable_InsertItemText    *syscall.LazyProc
	xAdTable_InsertItemTextEx  *syscall.LazyProc
	xAdTable_InsertItemImage   *syscall.LazyProc
	xAdTable_InsertItemImageEx *syscall.LazyProc
	xAdTable_SetItemText       *syscall.LazyProc
	xAdTable_SetItemTextEx     *syscall.LazyProc
	xAdTable_SetItemInt        *syscall.LazyProc
	xAdTable_SetItemIntEx      *syscall.LazyProc
	xAdTable_SetItemFloat      *syscall.LazyProc
	xAdTable_SetItemFloatEx    *syscall.LazyProc
	xAdTable_SetItemImage      *syscall.LazyProc
	xAdTable_SetItemImageEx    *syscall.LazyProc
	xAdTable_DeleteItem        *syscall.LazyProc
	xAdTable_DeleteItemEx      *syscall.LazyProc
	xAdTable_DeleteItemAll     *syscall.LazyProc
	xAdTable_DeleteColumnAll   *syscall.LazyProc
	xAdTable_GetCount          *syscall.LazyProc
	xAdTable_GetCountColumn    *syscall.LazyProc
	xAdTable_GetItemText       *syscall.LazyProc
	xAdTable_GetItemTextEx     *syscall.LazyProc
	xAdTable_GetItemImage      *syscall.LazyProc
	xAdTable_GetItemImageEx    *syscall.LazyProc
	xAdTable_GetItemInt        *syscall.LazyProc
	xAdTable_GetItemIntEx      *syscall.LazyProc
	xAdTable_GetItemFloat      *syscall.LazyProc
	xAdTable_GetItemFloatEx    *syscall.LazyProc

	// AdapterMap.
	xAdMap_Create       *syscall.LazyProc
	xAdMap_AddItemText  *syscall.LazyProc
	xAdMap_AddItemImage *syscall.LazyProc
	xAdMap_DeleteItem   *syscall.LazyProc
	xAdMap_GetCount     *syscall.LazyProc
	xAdMap_GetItemText  *syscall.LazyProc
	xAdMap_GetItemImage *syscall.LazyProc
	xAdMap_SetItemText  *syscall.LazyProc
	xAdMap_SetItemImage *syscall.LazyProc

	// AdapterTree.
	xAdTree_Create            *syscall.LazyProc
	xAdTree_AddColumn         *syscall.LazyProc
	xAdTree_SetColumn         *syscall.LazyProc
	xAdTree_InsertItemText    *syscall.LazyProc
	xAdTree_InsertItemTextEx  *syscall.LazyProc
	xAdTree_InsertItemImage   *syscall.LazyProc
	xAdTree_InsertItemImageEx *syscall.LazyProc
	xAdTree_GetCount          *syscall.LazyProc
	xAdTree_GetCountColumn    *syscall.LazyProc
	xAdTree_SetItemText       *syscall.LazyProc
	xAdTree_SetItemTextEx     *syscall.LazyProc
	xAdTree_SetItemImage      *syscall.LazyProc
	xAdTree_SetItemImageEx    *syscall.LazyProc
	xAdTree_GetItemText       *syscall.LazyProc
	xAdTree_GetItemTextEx     *syscall.LazyProc
	xAdTree_GetItemImage      *syscall.LazyProc
	xAdTree_GetItemImageEx    *syscall.LazyProc
	xAdTree_DeleteItem        *syscall.LazyProc
	xAdTree_DeleteItemAll     *syscall.LazyProc
	xAdTree_DeleteColumnAll   *syscall.LazyProc

	// Editor.
	xEditor_Create                     *syscall.LazyProc
	xEidtor_EnableAutoMatchSpaseSelect *syscall.LazyProc
	xEditor_IsBreakpoint               *syscall.LazyProc
	xEditor_SetBreakpoint              *syscall.LazyProc
	xEditor_RemoveBreakpoint           *syscall.LazyProc
	xEditor_ClearBreakpoint            *syscall.LazyProc
	xEditor_SetRunRow                  *syscall.LazyProc
	xEditor_GetColor                   *syscall.LazyProc
	xEditor_SetColor                   *syscall.LazyProc
	xEditor_SetStyleKeyword            *syscall.LazyProc
	xEditor_SetStyleFunction           *syscall.LazyProc
	xEditor_SetStyleVar                *syscall.LazyProc
	xEditor_SetStyleDataType           *syscall.LazyProc
	xEditor_SetStyleClass              *syscall.LazyProc
	xEditor_SetStyleMacro              *syscall.LazyProc
	xEditor_SetStyleString             *syscall.LazyProc
	xEditor_SetStyleComment            *syscall.LazyProc
	xEditor_SetStyleNumber             *syscall.LazyProc
	xEditor_GetBreakpointCount         *syscall.LazyProc
	xEditor_GetBreakpoints             *syscall.LazyProc
	xEditor_SetCurRow                  *syscall.LazyProc
	xEditor_GetDepth                   *syscall.LazyProc
	xEditor_ToExpandRow                *syscall.LazyProc
	xEditor_ExpandEx                   *syscall.LazyProc
	xEditor_ExpandAll                  *syscall.LazyProc
	xEditor_Expand                     *syscall.LazyProc
	xEditor_AddKeyword                 *syscall.LazyProc
	xEditor_AddConst                   *syscall.LazyProc
	xEditor_AddFunction                *syscall.LazyProc
	xEditor_AddExcludeDefVarKeyword    *syscall.LazyProc
	xEditor_GetExpandState             *syscall.LazyProc
	xEditor_SetExpandState             *syscall.LazyProc
	xEditor_GetIndentation             *syscall.LazyProc
	xEidtor_IsEmptyRow                 *syscall.LazyProc
	xEditor_SetAutoMatchMode           *syscall.LazyProc

	// Edit.
	xEdit_Create               *syscall.LazyProc
	xEdit_CreateEx             *syscall.LazyProc
	xEdit_EnableAutoWrap       *syscall.LazyProc
	xEdit_EnableReadOnly       *syscall.LazyProc
	xEdit_EnableMultiLine      *syscall.LazyProc
	xEdit_EnablePassword       *syscall.LazyProc
	xEdit_EnableAutoSelAll     *syscall.LazyProc
	xEdit_EnableAutoCancelSel  *syscall.LazyProc
	xEdit_IsReadOnly           *syscall.LazyProc
	xEdit_IsMultiLine          *syscall.LazyProc
	xEdit_IsPassword           *syscall.LazyProc
	xEdit_IsAutoWrap           *syscall.LazyProc
	xEdit_IsEmpty              *syscall.LazyProc
	xEdit_IsInSelect           *syscall.LazyProc
	xEdit_GetRowCount          *syscall.LazyProc
	xEdit_GetData              *syscall.LazyProc
	xEdit_AddData              *syscall.LazyProc
	xEdit_FreeData             *syscall.LazyProc
	xEdit_SetDefaultText       *syscall.LazyProc
	xEdit_SetDefaultTextColor  *syscall.LazyProc
	xEdit_SetPasswordCharacter *syscall.LazyProc
	xEdit_SetTextAlign         *syscall.LazyProc
	xEdit_SetTabSpace          *syscall.LazyProc
	xEdit_SetText              *syscall.LazyProc
	xEdit_SetTextInt           *syscall.LazyProc
	xEdit_GetText              *syscall.LazyProc
	xEdit_GetTextRow           *syscall.LazyProc
	xEdit_GetLength            *syscall.LazyProc
	xEdit_GetLengthRow         *syscall.LazyProc
	xEdit_GetAt                *syscall.LazyProc
	xEdit_InsertText           *syscall.LazyProc
	xEdit_AddTextUser          *syscall.LazyProc
	xEdit_AddText              *syscall.LazyProc
	xEdit_AddTextEx            *syscall.LazyProc
	xEdit_AddObject            *syscall.LazyProc
	xEdit_AddByStyle           *syscall.LazyProc
	xEdit_AddStyle             *syscall.LazyProc
	xEdit_AddStyleEx           *syscall.LazyProc
	xEdit_GetStyleInfo         *syscall.LazyProc
	xEdit_SetCurStyle          *syscall.LazyProc
	xEdit_SetCaretColor        *syscall.LazyProc
	xEdit_SetCaretWidth        *syscall.LazyProc
	xEdit_SetSelectBkColor     *syscall.LazyProc
	xEdit_SetRowHeight         *syscall.LazyProc
	xEdit_SetRowHeightEx       *syscall.LazyProc
	xEdit_SetCurPos            *syscall.LazyProc
	xEdit_GetCurPos            *syscall.LazyProc
	xEdit_GetCurRow            *syscall.LazyProc
	xEdit_GetCurCol            *syscall.LazyProc
	xEdit_GetPoint             *syscall.LazyProc
	xEdit_AutoScroll           *syscall.LazyProc
	xEdit_AutoScrollEx         *syscall.LazyProc
	xEdit_PosToRowCol          *syscall.LazyProc
	xEdit_RowColToPos          *syscall.LazyProc
	xEdit_SelectAll            *syscall.LazyProc
	xEdit_CancelSelect         *syscall.LazyProc
	xEdit_DeleteSelect         *syscall.LazyProc
	xEdit_SetSelect            *syscall.LazyProc
	xEdit_GetSelectText        *syscall.LazyProc
	xEdit_GetSelectRange       *syscall.LazyProc
	xEdit_GetVisibleRowRange   *syscall.LazyProc
	xEdit_Delete               *syscall.LazyProc
	xEdit_DeleteRow            *syscall.LazyProc
	xEdit_ClipboardCut         *syscall.LazyProc
	xEdit_ClipboardCopy        *syscall.LazyProc
	xEdit_ClipboardPaste       *syscall.LazyProc
	xEdit_Undo                 *syscall.LazyProc
	xEdit_Redo                 *syscall.LazyProc
	xEdit_AddChatBegin         *syscall.LazyProc
	xEdit_AddChatEnd           *syscall.LazyProc
	xEdit_SetChatIndentation   *syscall.LazyProc
	xEdit_SetRowSpace          *syscall.LazyProc
	xEdit_SetCurPosEx          *syscall.LazyProc
	xEdit_GetCurPosEx          *syscall.LazyProc
	xEdit_MoveEnd              *syscall.LazyProc
	xEdit_SetBackFont          *syscall.LazyProc
	xEdit_ReleaseStyle         *syscall.LazyProc
	xEdit_ModifyStyle          *syscall.LazyProc
	xEdit_SetSpaceSize         *syscall.LazyProc
	xEdit_SetCharSpaceSize     *syscall.LazyProc
	xEdit_GetSelectTextLength  *syscall.LazyProc
	xEdit_SetSelectTextStyle   *syscall.LazyProc

	// LayoutEle.
	xLayout_Create          *syscall.LazyProc
	xLayout_CreateEx        *syscall.LazyProc
	xLayout_IsEnableLayout  *syscall.LazyProc
	xLayout_EnableLayout    *syscall.LazyProc
	xLayout_ShowLayoutFrame *syscall.LazyProc
	xLayout_GetWidthIn      *syscall.LazyProc
	xLayout_GetHeightIn     *syscall.LazyProc

	// LayoutFrame.
	xLayoutFrame_Create          *syscall.LazyProc
	xLayoutFrame_ShowLayoutFrame *syscall.LazyProc

	// ProgressBar.
	xProgBar_Create         *syscall.LazyProc
	xProgBar_SetRange       *syscall.LazyProc
	xProgBar_GetRange       *syscall.LazyProc
	xProgBar_SetImageLoad   *syscall.LazyProc
	xProgBar_SetPos         *syscall.LazyProc
	xProgBar_GetPos         *syscall.LazyProc
	xProgBar_EnableHorizon  *syscall.LazyProc
	xProgBar_EnableShowText *syscall.LazyProc
	xProgBar_EnableStretch  *syscall.LazyProc

	// TextLink.
	xTextLink_Create                 *syscall.LazyProc
	xTextLink_EnableUnderlineLeave   *syscall.LazyProc
	xTextLink_EnableUnderlineStay    *syscall.LazyProc
	xTextLink_SetTextColorStay       *syscall.LazyProc
	xTextLink_SetUnderlineColorLeave *syscall.LazyProc
	xTextLink_SetUnderlineColorStay  *syscall.LazyProc

	// Shape.
	xShape_RemoveShape      *syscall.LazyProc
	xShape_GetZOrder        *syscall.LazyProc
	xShape_Redraw           *syscall.LazyProc
	xShape_GetWidth         *syscall.LazyProc
	xShape_GetHeight        *syscall.LazyProc
	xShape_SetPosition      *syscall.LazyProc
	xShape_GetRect          *syscall.LazyProc
	xShape_SetRect          *syscall.LazyProc
	xShape_SetRectLogic     *syscall.LazyProc
	xShape_GetRectLogic     *syscall.LazyProc
	xShape_GetWndClientRect *syscall.LazyProc
	xShape_GetContentSize   *syscall.LazyProc
	xShape_ShowLayout       *syscall.LazyProc
	xShape_AdjustLayout     *syscall.LazyProc
	xShape_Destroy          *syscall.LazyProc
	xShape_GetPosition      *syscall.LazyProc
	xShape_SetSize          *syscall.LazyProc
	xShape_GetSize          *syscall.LazyProc
	xShape_SetAlpha         *syscall.LazyProc
	xShape_GetAlpha         *syscall.LazyProc

	// ShapeText.
	xShapeText_Create        *syscall.LazyProc
	xShapeText_SetText       *syscall.LazyProc
	xShapeText_GetText       *syscall.LazyProc
	xShapeText_GetTextLength *syscall.LazyProc
	xShapeText_SetFont       *syscall.LazyProc
	xShapeText_GetFont       *syscall.LazyProc
	xShapeText_SetTextColor  *syscall.LazyProc
	xShapeText_GetTextColor  *syscall.LazyProc
	xShapeText_SetTextAlign  *syscall.LazyProc
	xShapeText_SetOffset     *syscall.LazyProc

	// ShapePicture.
	xShapePic_Create   *syscall.LazyProc
	xShapePic_SetImage *syscall.LazyProc
	xShapePic_GetImage *syscall.LazyProc

	// ShapeGif.
	xShapeGif_Create   *syscall.LazyProc
	xShapeGif_SetImage *syscall.LazyProc
	xShapeGif_GetImage *syscall.LazyProc

	// ShapeRect.
	xShapeRect_Create           *syscall.LazyProc
	xShapeRect_SetBorderColor   *syscall.LazyProc
	xShapeRect_SetFillColor     *syscall.LazyProc
	xShapeRect_SetRoundAngle    *syscall.LazyProc
	xShapeRect_GetRoundAngle    *syscall.LazyProc
	xShapeRect_EnableBorder     *syscall.LazyProc
	xShapeRect_EnableFill       *syscall.LazyProc
	xShapeRect_EnableRoundAngle *syscall.LazyProc

	// ShapeEllipse.
	xShapeEllipse_Create         *syscall.LazyProc
	xShapeEllipse_SetBorderColor *syscall.LazyProc
	xShapeEllipse_SetFillColor   *syscall.LazyProc
	xShapeEllipse_EnableBorder   *syscall.LazyProc
	xShapeEllipse_EnableFill     *syscall.LazyProc

	// ShapeGroupBox.
	xShapeGroupBox_Create           *syscall.LazyProc
	xShapeGroupBox_SetBorderColor   *syscall.LazyProc
	xShapeGroupBox_SetTextColor     *syscall.LazyProc
	xShapeGroupBox_SetFontX         *syscall.LazyProc
	xShapeGroupBox_SetTextOffset    *syscall.LazyProc
	xShapeGroupBox_SetRoundAngle    *syscall.LazyProc
	xShapeGroupBox_SetText          *syscall.LazyProc
	xShapeGroupBox_GetTextOffset    *syscall.LazyProc
	xShapeGroupBox_GetRoundAngle    *syscall.LazyProc
	xShapeGroupBox_EnableRoundAngle *syscall.LazyProc

	// ShapeLine.
	xShapeLine_Create      *syscall.LazyProc
	xShapeLine_SetPosition *syscall.LazyProc
	xShapeLine_SetColor    *syscall.LazyProc

	// Table.
	xTable_Create           *syscall.LazyProc
	xTable_Reset            *syscall.LazyProc
	xTable_ComboRow         *syscall.LazyProc
	xTable_ComboCol         *syscall.LazyProc
	xTable_SetColWidth      *syscall.LazyProc
	xTable_SetRowHeight     *syscall.LazyProc
	xTable_SetBorderColor   *syscall.LazyProc
	xTable_SetTextColor     *syscall.LazyProc
	xTable_SetFont          *syscall.LazyProc
	xTable_SetItemPadding   *syscall.LazyProc
	xTable_SetItemText      *syscall.LazyProc
	xTable_SetItemFont      *syscall.LazyProc
	xTable_SetItemTextAlign *syscall.LazyProc
	xTable_SetItemTextColor *syscall.LazyProc
	xTable_SetItemBkColor   *syscall.LazyProc
	xTable_SetItemLine      *syscall.LazyProc
	xTable_SetItemFlag      *syscall.LazyProc
	xTable_GetItemRect      *syscall.LazyProc

	// BkManager.
	xBkM_Create            *syscall.LazyProc
	xBkM_Destroy           *syscall.LazyProc
	xBkM_SetBkInfo         *syscall.LazyProc
	xBkM_AddInfo           *syscall.LazyProc
	xBkM_AddBorder         *syscall.LazyProc
	xBkM_AddFill           *syscall.LazyProc
	xBkM_AddImage          *syscall.LazyProc
	xBkM_GetCount          *syscall.LazyProc
	xBkM_Clear             *syscall.LazyProc
	xBkM_Draw              *syscall.LazyProc
	xBkM_DrawEx            *syscall.LazyProc
	xBkM_EnableAutoDestroy *syscall.LazyProc
	xBkM_AddRef            *syscall.LazyProc
	xBkM_Release           *syscall.LazyProc
	xBkM_GetRefCount       *syscall.LazyProc
	xBkM_SetInfo           *syscall.LazyProc
	xBkM_GetStateTextColor *syscall.LazyProc
	xBkM_GetObject         *syscall.LazyProc

	// Draw.
	xDraw_Create                  *syscall.LazyProc
	xDraw_Destroy                 *syscall.LazyProc
	xDraw_SetOffset               *syscall.LazyProc
	xDraw_GetOffset               *syscall.LazyProc
	xDraw_GDI_RestoreGDIOBJ       *syscall.LazyProc
	xDraw_GetHDC                  *syscall.LazyProc
	xDraw_SetBrushColor           *syscall.LazyProc
	xDraw_SetTextVertical         *syscall.LazyProc
	xDraw_SetTextAlign            *syscall.LazyProc
	xDraw_SetFont                 *syscall.LazyProc
	xDraw_SetLineWidth            *syscall.LazyProc
	xDraw_SetLineWidthF           *syscall.LazyProc
	xDraw_GDI_SetBkMode           *syscall.LazyProc
	xDraw_SetClipRect             *syscall.LazyProc
	xDraw_ClearClip               *syscall.LazyProc
	xDraw_EnableSmoothingMode     *syscall.LazyProc
	xDraw_EnableWndTransparent    *syscall.LazyProc
	xDraw_GDI_CreateSolidBrush    *syscall.LazyProc
	xDraw_GDI_CreatePen           *syscall.LazyProc
	xDraw_GDI_CreateRectRgn       *syscall.LazyProc
	xDraw_GDI_CreateRoundRectRgn  *syscall.LazyProc
	xDraw_GDI_CreatePolygonRgn    *syscall.LazyProc
	xDraw_GDI_SelectClipRgn       *syscall.LazyProc
	xDraw_FillRect                *syscall.LazyProc
	xDraw_FillRectF               *syscall.LazyProc
	xDraw_FillRectColor           *syscall.LazyProc
	xDraw_FillRectColorF          *syscall.LazyProc
	xDraw_GDI_FillRgn             *syscall.LazyProc
	xDraw_FillEllipse             *syscall.LazyProc
	xDraw_FillEllipseF            *syscall.LazyProc
	xDraw_DrawEllipse             *syscall.LazyProc
	xDraw_FillRoundRect           *syscall.LazyProc
	xDraw_FillRoundRectF          *syscall.LazyProc
	xDraw_DrawRoundRect           *syscall.LazyProc
	xDraw_DrawRoundRectF          *syscall.LazyProc
	xDraw_FillRoundRectEx         *syscall.LazyProc
	xDraw_FillRoundRectExF        *syscall.LazyProc
	xDraw_DrawRoundRectEx         *syscall.LazyProc
	xDraw_DrawRoundRectExF        *syscall.LazyProc
	xDraw_GDI_Rectangle           *syscall.LazyProc
	xDraw_GradientFill2           *syscall.LazyProc
	xDraw_GradientFill2F          *syscall.LazyProc
	xDraw_GradientFill4           *syscall.LazyProc
	xDraw_GradientFill4F          *syscall.LazyProc
	xDraw_GDI_FrameRgn            *syscall.LazyProc
	xDraw_DrawLine                *syscall.LazyProc
	xDraw_DrawLineF               *syscall.LazyProc
	xDraw_DrawCurve               *syscall.LazyProc
	xDraw_DrawCurveF              *syscall.LazyProc
	xDraw_FocusRect               *syscall.LazyProc
	xDraw_FocusRectF              *syscall.LazyProc
	xDraw_GDI_MoveToEx            *syscall.LazyProc
	xDraw_GDI_LineTo              *syscall.LazyProc
	xDraw_GDI_Polyline            *syscall.LazyProc
	xDraw_Dottedline              *syscall.LazyProc
	xDraw_DottedlineF             *syscall.LazyProc
	xDraw_GDI_SetPixel            *syscall.LazyProc
	xDraw_GDI_DrawIconEx          *syscall.LazyProc
	xDraw_GDI_BitBlt              *syscall.LazyProc
	xDraw_GDI_BitBlt2             *syscall.LazyProc
	xDraw_GDI_AlphaBlend          *syscall.LazyProc
	xDraw_DrawPolygon             *syscall.LazyProc
	xDraw_DrawPolygonF            *syscall.LazyProc
	xDraw_FillPolygon             *syscall.LazyProc
	xDraw_FillPolygonF            *syscall.LazyProc
	xDraw_Image                   *syscall.LazyProc
	xDraw_ImageF                  *syscall.LazyProc
	xDraw_ImageEx                 *syscall.LazyProc
	xDraw_ImageExF                *syscall.LazyProc
	xDraw_ImageAdaptive           *syscall.LazyProc
	xDraw_ImageAdaptiveF          *syscall.LazyProc
	xDraw_ImageSuper              *syscall.LazyProc
	xDraw_ImageSuperF             *syscall.LazyProc
	xDraw_ImageSuperEx            *syscall.LazyProc
	xDraw_ImageSuperExF           *syscall.LazyProc
	xDraw_ImageSuperMask          *syscall.LazyProc
	xDraw_ImageMask               *syscall.LazyProc
	xDraw_DrawText                *syscall.LazyProc
	xDraw_DrawTextF               *syscall.LazyProc
	xDraw_DrawTextUnderline       *syscall.LazyProc
	xDraw_DrawTextUnderlineF      *syscall.LazyProc
	xDraw_TextOut                 *syscall.LazyProc
	xDraw_TextOutF                *syscall.LazyProc
	xDraw_TextOutEx               *syscall.LazyProc
	xDraw_TextOutExF              *syscall.LazyProc
	xDraw_TextOutA                *syscall.LazyProc
	xDraw_TextOutAF               *syscall.LazyProc
	xDraw_DrawArc                 *syscall.LazyProc
	xDraw_DrawArcF                *syscall.LazyProc
	xDraw_DrawRect                *syscall.LazyProc
	xDraw_DrawRectF               *syscall.LazyProc
	xDraw_GDI_Ellipse             *syscall.LazyProc
	xDraw_GetD2dRenderTarget      *syscall.LazyProc
	xDraw_ImageTile               *syscall.LazyProc
	xDraw_ImageTileF              *syscall.LazyProc
	xDraw_SetD2dTextRenderingMode *syscall.LazyProc
	xDraw_CreateGDI               *syscall.LazyProc
	xDraw_SetTextRenderingHint    *syscall.LazyProc
	xDraw_DrawSvgSrc              *syscall.LazyProc
	xDraw_DrawSvg                 *syscall.LazyProc
	xDraw_DrawSvgEx               *syscall.LazyProc
	xDraw_DrawSvgSize             *syscall.LazyProc
	xDraw_D2D_Clear               *syscall.LazyProc
	xDraw_ImageMaskRect           *syscall.LazyProc
	xDraw_ImageMaskEllipse        *syscall.LazyProc

	// Ease.
	xEase_Linear  *syscall.LazyProc
	xEase_Quad    *syscall.LazyProc
	xEase_Cubic   *syscall.LazyProc
	xEase_Quart   *syscall.LazyProc
	xEase_Quint   *syscall.LazyProc
	xEase_Sine    *syscall.LazyProc
	xEase_Expo    *syscall.LazyProc
	xEase_Circ    *syscall.LazyProc
	xEase_Elastic *syscall.LazyProc
	xEase_Back    *syscall.LazyProc
	xEase_Bounce  *syscall.LazyProc
	xEase_Ex      *syscall.LazyProc

	// FontX.
	xFont_Create            *syscall.LazyProc
	xFont_CreateEx          *syscall.LazyProc
	xFont_CreateFromHFONT   *syscall.LazyProc
	xFont_CreateFromFont    *syscall.LazyProc
	xFont_CreateFromFile    *syscall.LazyProc
	xFont_CreateLOGFONTW    *syscall.LazyProc
	xFont_EnableAutoDestroy *syscall.LazyProc
	xFont_GetFont           *syscall.LazyProc
	xFont_GetFontInfo       *syscall.LazyProc
	xFont_GetLOGFONTW       *syscall.LazyProc
	xFont_Destroy           *syscall.LazyProc
	xFont_AddRef            *syscall.LazyProc
	xFont_GetRefCount       *syscall.LazyProc
	xFont_Release           *syscall.LazyProc
	xFont_CreateFromMem     *syscall.LazyProc
	xFont_CreateFromRes     *syscall.LazyProc
	xFont_CreateFromZip     *syscall.LazyProc
	xFont_CreateFromZipMem  *syscall.LazyProc

	// Image.
	xImage_LoadSrc             *syscall.LazyProc
	xImage_LoadFile            *syscall.LazyProc
	xImage_LoadFileAdaptive    *syscall.LazyProc
	xImage_LoadFileRect        *syscall.LazyProc
	xImage_LoadResAdaptive     *syscall.LazyProc
	xImage_LoadRes             *syscall.LazyProc
	xImage_LoadZip             *syscall.LazyProc
	xImage_LoadZipAdaptive     *syscall.LazyProc
	xImage_LoadZipRect         *syscall.LazyProc
	xImage_LoadZipMem          *syscall.LazyProc
	xImage_LoadMemory          *syscall.LazyProc
	xImage_LoadMemoryRect      *syscall.LazyProc
	xImage_LoadMemoryAdaptive  *syscall.LazyProc
	xImage_LoadFromImage       *syscall.LazyProc
	xImage_LoadFromExtractIcon *syscall.LazyProc
	xImage_LoadFromHICON       *syscall.LazyProc
	xImage_LoadFromHBITMAP     *syscall.LazyProc
	xImage_IsStretch           *syscall.LazyProc
	xImage_IsAdaptive          *syscall.LazyProc
	xImage_IsTile              *syscall.LazyProc
	xImage_SetDrawType         *syscall.LazyProc
	xImage_SetDrawTypeAdaptive *syscall.LazyProc
	xImage_SetTranColor        *syscall.LazyProc
	xImage_SetTranColorEx      *syscall.LazyProc
	xImage_SetRotateAngle      *syscall.LazyProc
	xImage_SetSplitEqual       *syscall.LazyProc
	xImage_EnableTranColor     *syscall.LazyProc
	xImage_EnableAutoDestroy   *syscall.LazyProc
	xImage_EnableCenter        *syscall.LazyProc
	xImage_IsCenter            *syscall.LazyProc
	xImage_GetDrawType         *syscall.LazyProc
	xImage_GetWidth            *syscall.LazyProc
	xImage_GetHeight           *syscall.LazyProc
	xImage_GetImageSrc         *syscall.LazyProc
	xImage_AddRef              *syscall.LazyProc
	xImage_Release             *syscall.LazyProc
	xImage_GetRefCount         *syscall.LazyProc
	xImage_Destroy             *syscall.LazyProc
	xImage_LoadSvg             *syscall.LazyProc
	xImage_LoadSvgFile         *syscall.LazyProc
	xImage_LoadSvgString       *syscall.LazyProc
	xImage_GetSvg              *syscall.LazyProc
	xImage_LoadSvgStringW      *syscall.LazyProc
	xImage_LoadSvgStringUtf8   *syscall.LazyProc
	xImage_SetScaleSize        *syscall.LazyProc

	// Svg.
	xSvg_LoadFile           *syscall.LazyProc
	xSvg_LoadString         *syscall.LazyProc
	xSvg_LoadZip            *syscall.LazyProc
	xSvg_LoadRes            *syscall.LazyProc
	xSvg_SetSize            *syscall.LazyProc
	xSvg_GetSize            *syscall.LazyProc
	xSvg_GetWidth           *syscall.LazyProc
	xSvg_GetHeight          *syscall.LazyProc
	xSvg_SetPosition        *syscall.LazyProc
	xSvg_GetPosition        *syscall.LazyProc
	xSvg_SetPositionF       *syscall.LazyProc
	xSvg_GetPositionF       *syscall.LazyProc
	xSvg_GetViewBox         *syscall.LazyProc
	xSvg_EnableAutoDestroy  *syscall.LazyProc
	xSvg_AddRef             *syscall.LazyProc
	xSvg_Release            *syscall.LazyProc
	xSvg_GetRefCount        *syscall.LazyProc
	xSvg_Destroy            *syscall.LazyProc
	xSvg_SetAlpha           *syscall.LazyProc
	xSvg_GetAlpha           *syscall.LazyProc
	xSvg_SetUserFillColor   *syscall.LazyProc
	xSvg_SetUserStrokeColor *syscall.LazyProc
	xSvg_GetUserFillColor   *syscall.LazyProc
	xSvg_GetUserStrokeColor *syscall.LazyProc
	xSvg_SetRotateAngle     *syscall.LazyProc
	xSvg_GetRotateAngle     *syscall.LazyProc
	xSvg_SetRotate          *syscall.LazyProc
	xSvg_GetRotate          *syscall.LazyProc
	xSvg_Show               *syscall.LazyProc
	xSvg_LoadStringW        *syscall.LazyProc
	xSvg_LoadStringUtf8     *syscall.LazyProc
	xSvg_LoadZipMem         *syscall.LazyProc

	// ListItemTemplate.
	xTemp_Load               *syscall.LazyProc
	xTemp_LoadZip            *syscall.LazyProc
	xTemp_LoadZipMem         *syscall.LazyProc
	xTemp_LoadEx             *syscall.LazyProc
	xTemp_LoadZipEx          *syscall.LazyProc
	xTemp_LoadZipMemEx       *syscall.LazyProc
	xTemp_LoadFromString     *syscall.LazyProc
	xTemp_LoadFromStringEx   *syscall.LazyProc
	xTemp_GetType            *syscall.LazyProc
	xTemp_Destroy            *syscall.LazyProc
	xTemp_Create             *syscall.LazyProc
	xTemp_AddNodeRoot        *syscall.LazyProc
	xTemp_AddNode            *syscall.LazyProc
	xTemp_CreateNode         *syscall.LazyProc
	xTemp_SetNodeAttribute   *syscall.LazyProc
	xTemp_SetNodeAttributeEx *syscall.LazyProc
	xTemp_List_GetNode       *syscall.LazyProc
	xTemp_GetNode            *syscall.LazyProc
	xTemp_CloneNode          *syscall.LazyProc
	xTemp_Clone              *syscall.LazyProc

	// Resource.
	xRes_EnableDelayLoad     *syscall.LazyProc
	xRes_SetLoadFileCallback *syscall.LazyProc
	xRes_GetIDValue          *syscall.LazyProc
	xRes_GetImage            *syscall.LazyProc
	xRes_GetImageEx          *syscall.LazyProc
	xRes_GetColor            *syscall.LazyProc
	xRes_GetFont             *syscall.LazyProc
	xRes_GetBkM              *syscall.LazyProc

	// ListBox.
	xListBox_Create                       *syscall.LazyProc
	xListBox_EnableFixedRowHeight         *syscall.LazyProc
	xListBox_EnablemTemplateReuse         *syscall.LazyProc
	xListBox_EnableVirtualTable           *syscall.LazyProc
	xListBox_SetVirtualRowCount           *syscall.LazyProc
	xListBox_SetDrawItemBkFlags           *syscall.LazyProc
	xListBox_SetItemData                  *syscall.LazyProc
	xListBox_GetItemData                  *syscall.LazyProc
	xListBox_SetItemInfo                  *syscall.LazyProc
	xListBox_GetItemInfo                  *syscall.LazyProc
	xListBox_SetSelectItem                *syscall.LazyProc
	xListBox_GetSelectItem                *syscall.LazyProc
	xListBox_AddSelectItem                *syscall.LazyProc
	xListBox_CancelSelectItem             *syscall.LazyProc
	xListBox_CancelSelectAll              *syscall.LazyProc
	xListBox_GetSelectAll                 *syscall.LazyProc
	xListBox_GetSelectCount               *syscall.LazyProc
	xListBox_GetItemMouseStay             *syscall.LazyProc
	xListBox_SelectAll                    *syscall.LazyProc
	xListBox_VisibleItem                  *syscall.LazyProc
	xListBox_GetVisibleRowRange           *syscall.LazyProc
	xListBox_SetItemHeightDefault         *syscall.LazyProc
	xListBox_GetItemHeightDefault         *syscall.LazyProc
	xListBox_GetItemIndexFromHXCGUI       *syscall.LazyProc
	xListBox_SetRowSpace                  *syscall.LazyProc
	xListBox_GetRowSpace                  *syscall.LazyProc
	xListBox_HitTest                      *syscall.LazyProc
	xListBox_HitTestOffset                *syscall.LazyProc
	xListBox_SetItemTemplateXML           *syscall.LazyProc
	xListBox_SetItemTemplate              *syscall.LazyProc
	xListBox_SetItemTemplateXMLFromString *syscall.LazyProc
	xListBox_GetTemplateObject            *syscall.LazyProc
	xListBox_EnableMultiSel               *syscall.LazyProc
	xListBox_CreateAdapter                *syscall.LazyProc
	xListBox_BindAdapter                  *syscall.LazyProc
	xListBox_GetAdapter                   *syscall.LazyProc
	xListBox_Sort                         *syscall.LazyProc
	xListBox_RefreshData                  *syscall.LazyProc
	xListBox_RefreshItem                  *syscall.LazyProc
	xListBox_AddItemText                  *syscall.LazyProc
	xListBox_AddItemTextEx                *syscall.LazyProc
	xListBox_AddItemImage                 *syscall.LazyProc
	xListBox_AddItemImageEx               *syscall.LazyProc
	xListBox_InsertItemText               *syscall.LazyProc
	xListBox_InsertItemTextEx             *syscall.LazyProc
	xListBox_InsertItemImage              *syscall.LazyProc
	xListBox_InsertItemImageEx            *syscall.LazyProc
	xListBox_SetItemText                  *syscall.LazyProc
	xListBox_SetItemTextEx                *syscall.LazyProc
	xListBox_SetItemImage                 *syscall.LazyProc
	xListBox_SetItemImageEx               *syscall.LazyProc
	xListBox_SetItemInt                   *syscall.LazyProc
	xListBox_SetItemIntEx                 *syscall.LazyProc
	xListBox_SetItemFloat                 *syscall.LazyProc
	xListBox_SetItemFloatEx               *syscall.LazyProc
	xListBox_GetItemText                  *syscall.LazyProc
	xListBox_GetItemTextEx                *syscall.LazyProc
	xListBox_GetItemImage                 *syscall.LazyProc
	xListBox_GetItemImageEx               *syscall.LazyProc
	xListBox_GetItemInt                   *syscall.LazyProc
	xListBox_GetItemIntEx                 *syscall.LazyProc
	xListBox_GetItemFloat                 *syscall.LazyProc
	xListBox_GetItemFloatEx               *syscall.LazyProc
	xListBox_DeleteItem                   *syscall.LazyProc
	xListBox_DeleteItemEx                 *syscall.LazyProc
	xListBox_DeleteItemAll                *syscall.LazyProc
	xListBox_DeleteColumnAll              *syscall.LazyProc
	xListBox_GetCount_AD                  *syscall.LazyProc
	xListBox_GetCountColumn_AD            *syscall.LazyProc
	xListBox_SetSplitLineColor            *syscall.LazyProc
	xListBox_SetDragRectColor             *syscall.LazyProc

	// List.
	xList_Create                       *syscall.LazyProc
	xList_AddColumn                    *syscall.LazyProc
	xList_InsertColumn                 *syscall.LazyProc
	xList_EnableMultiSel               *syscall.LazyProc
	xList_EnableDragChangeColumnWidth  *syscall.LazyProc
	xList_EnableVScrollBarTop          *syscall.LazyProc
	xList_EnableItemBkFullRow          *syscall.LazyProc
	xList_EnableFixedRowHeight         *syscall.LazyProc
	xList_EnablemTemplateReuse         *syscall.LazyProc
	xList_EnableVirtualTable           *syscall.LazyProc
	xList_SetVirtualRowCount           *syscall.LazyProc
	xList_SetSort                      *syscall.LazyProc
	xList_SetDrawItemBkFlags           *syscall.LazyProc
	xList_SetColumnWidth               *syscall.LazyProc
	xList_SetColumnMinWidth            *syscall.LazyProc
	xList_SetColumnWidthFixed          *syscall.LazyProc
	xList_GetColumnWidth               *syscall.LazyProc
	xList_GetColumnCount               *syscall.LazyProc
	xList_SetItemData                  *syscall.LazyProc
	xList_GetItemData                  *syscall.LazyProc
	xList_SetSelectItem                *syscall.LazyProc
	xList_GetSelectItem                *syscall.LazyProc
	xList_GetSelectItemCount           *syscall.LazyProc
	xList_AddSelectItem                *syscall.LazyProc
	xList_SetSelectAll                 *syscall.LazyProc
	xList_GetSelectAll                 *syscall.LazyProc
	xList_VisibleItem                  *syscall.LazyProc
	xList_CancelSelectItem             *syscall.LazyProc
	xList_CancelSelectAll              *syscall.LazyProc
	xList_GetHeaderHELE                *syscall.LazyProc
	xList_DeleteColumn                 *syscall.LazyProc
	xList_DeleteColumnAll              *syscall.LazyProc
	xList_BindAdapter                  *syscall.LazyProc
	xList_BindAdapterHeader            *syscall.LazyProc
	xList_CreateAdapter                *syscall.LazyProc
	xList_CreateAdapterHeader          *syscall.LazyProc
	xList_GetAdapter                   *syscall.LazyProc
	xList_GetAdapterHeader             *syscall.LazyProc
	xList_SetItemTemplateXML           *syscall.LazyProc
	xList_SetItemTemplateXMLFromString *syscall.LazyProc
	xList_SetItemTemplate              *syscall.LazyProc
	xList_GetTemplateObject            *syscall.LazyProc
	xList_GetItemIndexFromHXCGUI       *syscall.LazyProc
	xList_GetHeaderTemplateObject      *syscall.LazyProc
	xList_GetHeaderItemIndexFromHXCGUI *syscall.LazyProc
	xList_SetHeaderHeight              *syscall.LazyProc
	xList_GetHeaderHeight              *syscall.LazyProc
	xList_GetVisibleRowRange           *syscall.LazyProc
	xList_SetItemHeightDefault         *syscall.LazyProc
	xList_GetItemHeightDefault         *syscall.LazyProc
	xList_SetRowSpace                  *syscall.LazyProc
	xList_GetRowSpace                  *syscall.LazyProc
	xList_SetLockColumnLeft            *syscall.LazyProc
	xList_SetLockColumnRight           *syscall.LazyProc
	xList_SetLockRowBottom             *syscall.LazyProc
	xList_SetLockRowBottomOverlap      *syscall.LazyProc
	xList_HitTest                      *syscall.LazyProc
	xList_HitTestOffset                *syscall.LazyProc
	xList_RefreshData                  *syscall.LazyProc
	xList_RefreshItem                  *syscall.LazyProc
	xList_AddColumnText                *syscall.LazyProc
	xList_AddColumnImage               *syscall.LazyProc
	xList_AddItemText                  *syscall.LazyProc
	xList_AddItemTextEx                *syscall.LazyProc
	xList_AddItemImage                 *syscall.LazyProc
	xList_AddItemImageEx               *syscall.LazyProc
	xList_InsertItemText               *syscall.LazyProc
	xList_InsertItemTextEx             *syscall.LazyProc
	xList_InsertItemImage              *syscall.LazyProc
	xList_InsertItemImageEx            *syscall.LazyProc
	xList_SetItemText                  *syscall.LazyProc
	xList_SetItemTextEx                *syscall.LazyProc
	xList_SetItemImage                 *syscall.LazyProc
	xList_SetItemImageEx               *syscall.LazyProc
	xList_SetItemInt                   *syscall.LazyProc
	xList_SetItemIntEx                 *syscall.LazyProc
	xList_SetItemFloat                 *syscall.LazyProc
	xList_SetItemFloatEx               *syscall.LazyProc
	xList_GetItemText                  *syscall.LazyProc
	xList_GetItemTextEx                *syscall.LazyProc
	xList_GetItemImage                 *syscall.LazyProc
	xList_GetItemImageEx               *syscall.LazyProc
	xList_GetItemInt                   *syscall.LazyProc
	xList_GetItemIntEx                 *syscall.LazyProc
	xList_GetItemFloat                 *syscall.LazyProc
	xList_GetItemFloatEx               *syscall.LazyProc
	xList_DeleteItem                   *syscall.LazyProc
	xList_DeleteItemEx                 *syscall.LazyProc
	xList_DeleteItemAll                *syscall.LazyProc
	xList_DeleteColumnAll_AD           *syscall.LazyProc
	xList_GetCount_AD                  *syscall.LazyProc
	xList_GetCountColumn_AD            *syscall.LazyProc
	xList_SetSplitLineColor            *syscall.LazyProc
	xList_SetItemHeight                *syscall.LazyProc
	xList_GetItemHeight                *syscall.LazyProc
	xList_SetDragRectColor             *syscall.LazyProc

	// ListView.
	xListView_Create                       *syscall.LazyProc
	xListView_CreateAdapter                *syscall.LazyProc
	xListView_BindAdapter                  *syscall.LazyProc
	xListView_GetAdapter                   *syscall.LazyProc
	xListView_SetItemTemplateXML           *syscall.LazyProc
	xListView_SetItemTemplateXMLFromString *syscall.LazyProc
	xListView_SetItemTemplate              *syscall.LazyProc
	xListView_GetTemplateObject            *syscall.LazyProc
	xListView_GetTemplateObjectGroup       *syscall.LazyProc
	xListView_GetItemIDFromHXCGUI          *syscall.LazyProc
	xListView_HitTest                      *syscall.LazyProc
	xListView_HitTestOffset                *syscall.LazyProc
	xListView_EnableMultiSel               *syscall.LazyProc
	xListView_EnablemTemplateReuse         *syscall.LazyProc
	xListView_EnableVirtualTable           *syscall.LazyProc
	xListView_SetVirtualItemCount          *syscall.LazyProc
	xListView_SetDrawItemBkFlags           *syscall.LazyProc
	xListView_SetSelectItem                *syscall.LazyProc
	xListView_GetSelectItem                *syscall.LazyProc
	xListView_AddSelectItem                *syscall.LazyProc
	xListView_VisibleItem                  *syscall.LazyProc
	xListView_GetVisibleItemRange          *syscall.LazyProc
	xListView_GetSelectItemCount           *syscall.LazyProc
	xListView_GetSelectAll                 *syscall.LazyProc
	xListView_SetSelectAll                 *syscall.LazyProc
	xListView_CancelSelectAll              *syscall.LazyProc
	xListView_SetColumnSpace               *syscall.LazyProc
	xListView_SetRowSpace                  *syscall.LazyProc
	xListView_SetItemSize                  *syscall.LazyProc
	xListView_GetItemSize                  *syscall.LazyProc
	xListView_SetGroupHeight               *syscall.LazyProc
	xListView_GetGroupHeight               *syscall.LazyProc
	xListView_SetGroupUserData             *syscall.LazyProc
	xListView_SetItemUserData              *syscall.LazyProc
	xListView_GetGroupUserData             *syscall.LazyProc
	xListView_GetItemUserData              *syscall.LazyProc
	xListView_RefreshData                  *syscall.LazyProc
	xListView_RefreshItem                  *syscall.LazyProc
	xListView_ExpandGroup                  *syscall.LazyProc
	xListView_Group_AddColumn              *syscall.LazyProc
	xListView_Group_AddItemText            *syscall.LazyProc
	xListView_Group_AddItemTextEx          *syscall.LazyProc
	xListView_Group_AddItemImage           *syscall.LazyProc
	xListView_Group_AddItemImageEx         *syscall.LazyProc
	xListView_Group_SetText                *syscall.LazyProc
	xListView_Group_SetTextEx              *syscall.LazyProc
	xListView_Group_SetImage               *syscall.LazyProc
	xListView_Group_SetImageEx             *syscall.LazyProc
	xListView_Group_GetCount               *syscall.LazyProc
	xListView_Item_GetCount                *syscall.LazyProc
	xListView_Item_AddColumn               *syscall.LazyProc
	xListView_Item_AddItemText             *syscall.LazyProc
	xListView_Item_AddItemTextEx           *syscall.LazyProc
	xListView_Item_AddItemImage            *syscall.LazyProc
	xListView_Item_AddItemImageEx          *syscall.LazyProc
	xListView_Item_SetText                 *syscall.LazyProc
	xListView_Item_SetTextEx               *syscall.LazyProc
	xListView_Item_SetImage                *syscall.LazyProc
	xListView_Item_SetImageEx              *syscall.LazyProc
	xListView_Group_DeleteItem             *syscall.LazyProc
	xListView_Group_DeleteAllChildItem     *syscall.LazyProc
	xListView_Item_DeleteItem              *syscall.LazyProc
	xListView_DeleteAll                    *syscall.LazyProc
	xListView_DeleteAllGroup               *syscall.LazyProc
	xListView_DeleteAllItem                *syscall.LazyProc
	xListView_DeleteColumnGroup            *syscall.LazyProc
	xListView_DeleteColumnItem             *syscall.LazyProc
	xListView_Item_GetTextEx               *syscall.LazyProc
	xListView_Item_GetImageEx              *syscall.LazyProc
	xListView_Group_GetText                *syscall.LazyProc
	xListView_Group_GetTextEx              *syscall.LazyProc
	xListView_Group_GetImage               *syscall.LazyProc
	xListView_Group_GetImageEx             *syscall.LazyProc
	xListView_Item_GetText                 *syscall.LazyProc
	xListView_Item_GetImage                *syscall.LazyProc
	xListView_SetDragRectColor             *syscall.LazyProc

	// MenuBar.
	xMenuBar_Create          *syscall.LazyProc
	xMenuBar_AddButton       *syscall.LazyProc
	xMenuBar_SetButtonHeight *syscall.LazyProc
	xMenuBar_GetMenu         *syscall.LazyProc
	xMenuBar_DeleteButton    *syscall.LazyProc
	xMenuBar_EnableAutoWidth *syscall.LazyProc
	xMenuBar_GetButton       *syscall.LazyProc

	// Pane.
	xPane_Create           *syscall.LazyProc
	xPane_SetView          *syscall.LazyProc
	xPane_SetTitle         *syscall.LazyProc
	xPane_GetTitle         *syscall.LazyProc
	xPane_SetCaptionHeight *syscall.LazyProc
	xPane_GetCaptionHeight *syscall.LazyProc
	xPane_IsShowPane       *syscall.LazyProc
	xPane_SetSize          *syscall.LazyProc
	xPane_GetState         *syscall.LazyProc
	xPane_GetViewRect      *syscall.LazyProc
	xPane_HidePane         *syscall.LazyProc
	xPane_ShowPane         *syscall.LazyProc
	xPane_DockPane         *syscall.LazyProc
	xPane_LockPane         *syscall.LazyProc
	xPane_FloatPane        *syscall.LazyProc
	xPane_DrawPane         *syscall.LazyProc
	xPane_SetSelect        *syscall.LazyProc
	xPane_IsGroupActivate  *syscall.LazyProc

	// ScrollBar.
	xSBar_Create             *syscall.LazyProc
	xSBar_SetRange           *syscall.LazyProc
	xSBar_GetRange           *syscall.LazyProc
	xSBar_ShowButton         *syscall.LazyProc
	xSBar_SetSliderLength    *syscall.LazyProc
	xSBar_SetSliderMinLength *syscall.LazyProc
	xSBar_SetSliderPadding   *syscall.LazyProc
	xSBar_EnableHorizon      *syscall.LazyProc
	xSBar_GetSliderMaxLength *syscall.LazyProc
	xSBar_ScrollUp           *syscall.LazyProc
	xSBar_ScrollDown         *syscall.LazyProc
	xSBar_ScrollTop          *syscall.LazyProc
	xSBar_ScrollBottom       *syscall.LazyProc
	xSBar_ScrollPos          *syscall.LazyProc
	xSBar_GetButtonUp        *syscall.LazyProc
	xSBar_GetButtonDown      *syscall.LazyProc
	xSBar_GetButtonSlider    *syscall.LazyProc

	// ScrollView.
	xSView_Create                  *syscall.LazyProc
	xSView_SetTotalSize            *syscall.LazyProc
	xSView_GetTotalSize            *syscall.LazyProc
	xSView_SetLineSize             *syscall.LazyProc
	xSView_GetLineSize             *syscall.LazyProc
	xSView_SetScrollBarSize        *syscall.LazyProc
	xSView_GetViewPosH             *syscall.LazyProc
	xSView_GetViewPosV             *syscall.LazyProc
	xSView_GetViewWidth            *syscall.LazyProc
	xSView_GetViewHeight           *syscall.LazyProc
	xSView_GetViewRect             *syscall.LazyProc
	xSView_GetScrollBarH           *syscall.LazyProc
	xSView_GetScrollBarV           *syscall.LazyProc
	xSView_ScrollPosH              *syscall.LazyProc
	xSView_ScrollPosV              *syscall.LazyProc
	xSView_ScrollPosXH             *syscall.LazyProc
	xSView_ScrollPosYV             *syscall.LazyProc
	xSView_ShowSBarH               *syscall.LazyProc
	xSView_ShowSBarV               *syscall.LazyProc
	xSView_EnableAutoShowScrollBar *syscall.LazyProc
	xSView_ScrollLeftLine          *syscall.LazyProc
	xSView_ScrollRightLine         *syscall.LazyProc
	xSView_ScrollTopLine           *syscall.LazyProc
	xSView_ScrollBottomLine        *syscall.LazyProc
	xSView_ScrollLeft              *syscall.LazyProc
	xSView_ScrollRight             *syscall.LazyProc
	xSView_ScrollTop               *syscall.LazyProc
	xSView_ScrollBottom            *syscall.LazyProc

	// SliderBar.
	xSliderBar_Create          *syscall.LazyProc
	xSliderBar_SetRange        *syscall.LazyProc
	xSliderBar_GetRange        *syscall.LazyProc
	xSliderBar_SetImageLoad    *syscall.LazyProc
	xSliderBar_SetButtonWidth  *syscall.LazyProc
	xSliderBar_SetButtonHeight *syscall.LazyProc
	xSliderBar_SetPos          *syscall.LazyProc
	xSliderBar_GetPos          *syscall.LazyProc
	xSliderBar_GetButton       *syscall.LazyProc
	xSliderBar_EnableHorizon   *syscall.LazyProc

	// TabBar.
	xTabBar_Create            *syscall.LazyProc
	xTabBar_AddLabel          *syscall.LazyProc
	xTabBar_InsertLabel       *syscall.LazyProc
	xTabBar_MoveLabel         *syscall.LazyProc
	xTabBar_DeleteLabel       *syscall.LazyProc
	xTabBar_DeleteLabelAll    *syscall.LazyProc
	xTabBar_GetLabel          *syscall.LazyProc
	xTabBar_GetLabelClose     *syscall.LazyProc
	xTabBar_GetButtonLeft     *syscall.LazyProc
	xTabBar_GetButtonRight    *syscall.LazyProc
	xTabBar_GetButtonDropMenu *syscall.LazyProc
	xTabBar_GetSelect         *syscall.LazyProc
	xTabBar_GetLabelSpacing   *syscall.LazyProc
	xTabBar_GetLabelCount     *syscall.LazyProc
	xTabBar_GetindexByEle     *syscall.LazyProc
	xTabBar_SetLabelSpacing   *syscall.LazyProc
	xTabBar_SetPadding        *syscall.LazyProc
	xTabBar_SetSelect         *syscall.LazyProc
	xTabBar_SetUp             *syscall.LazyProc
	xTabBar_SetDown           *syscall.LazyProc
	xTabBar_EnableTile        *syscall.LazyProc
	xTabBar_EnableDropMenu    *syscall.LazyProc
	xTabBar_EnableClose       *syscall.LazyProc
	xTabBar_SetCloseSize      *syscall.LazyProc
	xTabBar_SetTurnButtonSize *syscall.LazyProc
	xTabBar_SetLabelWidth     *syscall.LazyProc
	xTabBar_ShowLabel         *syscall.LazyProc

	// ToolBar.
	xToolBar_Create           *syscall.LazyProc
	xToolBar_InsertEle        *syscall.LazyProc
	xToolBar_InsertSeparator  *syscall.LazyProc
	xToolBar_EnableButtonMenu *syscall.LazyProc
	xToolBar_GetEle           *syscall.LazyProc
	xToolBar_GetButtonLeft    *syscall.LazyProc
	xToolBar_GetButtonRight   *syscall.LazyProc
	xToolBar_GetButtonMenu    *syscall.LazyProc
	xToolBar_SetSpace         *syscall.LazyProc
	xToolBar_DeleteEle        *syscall.LazyProc
	xToolBar_DeleteAllEle     *syscall.LazyProc

	// Tree.
	xTree_Create                          *syscall.LazyProc
	xTree_EnableDragItem                  *syscall.LazyProc
	xTree_EnableConnectLine               *syscall.LazyProc
	xTree_EnableExpand                    *syscall.LazyProc
	xTree_EnablemTemplateReuse            *syscall.LazyProc
	xTree_SetConnectLineColor             *syscall.LazyProc
	xTree_SetExpandButtonSize             *syscall.LazyProc
	xTree_SetConnectLineLength            *syscall.LazyProc
	xTree_SetDragInsertPositionColor      *syscall.LazyProc
	xTree_SetItemTemplateXML              *syscall.LazyProc
	xTree_SetItemTemplateXMLSel           *syscall.LazyProc
	xTree_SetItemTemplate                 *syscall.LazyProc
	xTree_SetItemTemplateSel              *syscall.LazyProc
	xTree_SetItemTemplateXMLFromString    *syscall.LazyProc
	xTree_SetItemTemplateXMLSelFromString *syscall.LazyProc
	xTree_SetDrawItemBkFlags              *syscall.LazyProc
	xTree_SetItemData                     *syscall.LazyProc
	xTree_GetItemData                     *syscall.LazyProc
	xTree_SetSelectItem                   *syscall.LazyProc
	xTree_GetSelectItem                   *syscall.LazyProc
	xTree_VisibleItem                     *syscall.LazyProc
	xTree_IsExpand                        *syscall.LazyProc
	xTree_ExpandItem                      *syscall.LazyProc
	xTree_ExpandAllChildItem              *syscall.LazyProc
	xTree_HitTest                         *syscall.LazyProc
	xTree_HitTestOffset                   *syscall.LazyProc
	xTree_GetFirstChildItem               *syscall.LazyProc
	xTree_GetEndChildItem                 *syscall.LazyProc
	xTree_GetPrevSiblingItem              *syscall.LazyProc
	xTree_GetNextSiblingItem              *syscall.LazyProc
	xTree_GetParentItem                   *syscall.LazyProc
	xTree_CreateAdapter                   *syscall.LazyProc
	xTree_BindAdapter                     *syscall.LazyProc
	xTree_GetAdapter                      *syscall.LazyProc
	xTree_RefreshData                     *syscall.LazyProc
	xTree_RefreshItem                     *syscall.LazyProc
	xTree_SetIndentation                  *syscall.LazyProc
	xTree_GetIndentation                  *syscall.LazyProc
	xTree_SetItemHeightDefault            *syscall.LazyProc
	xTree_GetItemHeightDefault            *syscall.LazyProc
	xTree_SetItemHeight                   *syscall.LazyProc
	xTree_GetItemHeight                   *syscall.LazyProc
	xTree_SetRowSpace                     *syscall.LazyProc
	xTree_GetRowSpace                     *syscall.LazyProc
	xTree_MoveItem                        *syscall.LazyProc
	xTree_GetTemplateObject               *syscall.LazyProc
	xTree_GetItemIDFromHXCGUI             *syscall.LazyProc
	xTree_InsertItemText                  *syscall.LazyProc
	xTree_InsertItemTextEx                *syscall.LazyProc
	xTree_InsertItemImage                 *syscall.LazyProc
	xTree_InsertItemImageEx               *syscall.LazyProc
	xTree_GetCount                        *syscall.LazyProc
	xTree_GetCountColumn                  *syscall.LazyProc
	xTree_SetItemText                     *syscall.LazyProc
	xTree_SetItemTextEx                   *syscall.LazyProc
	xTree_SetItemImage                    *syscall.LazyProc
	xTree_SetItemImageEx                  *syscall.LazyProc
	xTree_GetItemText                     *syscall.LazyProc
	xTree_GetItemTextEx                   *syscall.LazyProc
	xTree_GetItemImage                    *syscall.LazyProc
	xTree_GetItemImageEx                  *syscall.LazyProc
	xTree_DeleteItem                      *syscall.LazyProc
	xTree_DeleteItemAll                   *syscall.LazyProc
	xTree_DeleteColumnAll                 *syscall.LazyProc
	xTree_SetSplitLineColor               *syscall.LazyProc

	// DateTime.
	xDateTime_Create           *syscall.LazyProc
	xDateTime_SetStyle         *syscall.LazyProc
	xDateTime_GetStyle         *syscall.LazyProc
	xDateTime_EnableSplitSlash *syscall.LazyProc
	xDateTime_GetButton        *syscall.LazyProc
	xDateTime_GetSelBkColor    *syscall.LazyProc
	xDateTime_SetSelBkColor    *syscall.LazyProc
	xDateTime_GetDate          *syscall.LazyProc
	xDateTime_SetDate          *syscall.LazyProc
	xDateTime_GetTime          *syscall.LazyProc
	xDateTime_SetTime          *syscall.LazyProc
	xDateTime_Popup            *syscall.LazyProc

	// MonthCal.
	xMonthCal_Create       *syscall.LazyProc
	xMonthCal_GetButton    *syscall.LazyProc
	xMonthCal_SetToday     *syscall.LazyProc
	xMonthCal_GetToday     *syscall.LazyProc
	xMonthCal_GetSelDate   *syscall.LazyProc
	xMonthCal_SetTextColor *syscall.LazyProc

	// XcObjectBase.
	xObj_GetType     *syscall.LazyProc
	xObj_GetTypeBase *syscall.LazyProc
	xObj_GetTypeEx   *syscall.LazyProc
	xObj_SetTypeEx   *syscall.LazyProc

	// Animation.
	xAnima_Run                       *syscall.LazyProc
	xAnima_Release                   *syscall.LazyProc
	xAnima_ReleaseEx                 *syscall.LazyProc
	xAnima_Create                    *syscall.LazyProc
	xAnima_Move                      *syscall.LazyProc
	xAnima_MoveEx                    *syscall.LazyProc
	xAnima_Rotate                    *syscall.LazyProc
	xAnima_RotateEx                  *syscall.LazyProc
	xAnima_Scale                     *syscall.LazyProc
	xAnima_ScaleSize                 *syscall.LazyProc
	xAnima_Alpha                     *syscall.LazyProc
	xAnima_AlphaEx                   *syscall.LazyProc
	xAnima_Color                     *syscall.LazyProc
	xAnima_ColorEx                   *syscall.LazyProc
	xAnima_LayoutWidth               *syscall.LazyProc
	xAnima_LayoutHeight              *syscall.LazyProc
	xAnima_LayoutSize                *syscall.LazyProc
	xAnima_Delay                     *syscall.LazyProc
	xAnima_Show                      *syscall.LazyProc
	xAnimaGroup_Create               *syscall.LazyProc
	xAnimaGroup_AddItem              *syscall.LazyProc
	xAnimaRotate_SetCenter           *syscall.LazyProc
	xAnimaScale_SetPosition          *syscall.LazyProc
	xAnima_GetObjectUI               *syscall.LazyProc
	xAnimaItem_EnableCompleteRelease *syscall.LazyProc
	xAnima_EnableAutoDestroy         *syscall.LazyProc
	xAnima_DestroyObjectUI           *syscall.LazyProc
	xAnima_SetCallBack               *syscall.LazyProc
	xAnima_SetUserData               *syscall.LazyProc
	xAnima_GetUserData               *syscall.LazyProc
	xAnima_Stop                      *syscall.LazyProc
	xAnima_Start                     *syscall.LazyProc
	xAnima_Pause                     *syscall.LazyProc
	xAnimaItem_SetCallback           *syscall.LazyProc
	xAnimaItem_SetUserData           *syscall.LazyProc
	xAnimaItem_GetUserData           *syscall.LazyProc
	xAnimaItem_EnableAutoDestroy     *syscall.LazyProc
	xAnima_DelayEx                   *syscall.LazyProc
	xAnimaMove_SetFlag               *syscall.LazyProc

	// 通知消息.
	xNotifyMsg_Popup            *syscall.LazyProc
	xNotifyMsg_PopupEx          *syscall.LazyProc
	xNotifyMsg_SetDuration      *syscall.LazyProc
	xNotifyMsg_SetParentMargin  *syscall.LazyProc
	xNotifyMsg_WindowPopup      *syscall.LazyProc
	xNotifyMsg_WindowPopupEx    *syscall.LazyProc
	xNotifyMsg_SetCaptionHeight *syscall.LazyProc
	xNotifyMsg_SetWidth         *syscall.LazyProc
	xNotifyMsg_SetSpace         *syscall.LazyProc
	xNotifyMsg_SetBorderSize    *syscall.LazyProc

	// 背景对象.
	xBkObj_SetMargin         *syscall.LazyProc
	xBkObj_SetAlign          *syscall.LazyProc
	xBkObj_SetImage          *syscall.LazyProc
	xBkObj_SetRotate         *syscall.LazyProc
	xBkObj_SetFillColor      *syscall.LazyProc
	xBkObj_SetBorderWidth    *syscall.LazyProc
	xBkObj_SetBorderColor    *syscall.LazyProc
	xBkObj_SetRectRoundAngle *syscall.LazyProc
	xBkObj_EnableFill        *syscall.LazyProc
	xBkObj_EnableBorder      *syscall.LazyProc
	xBkObj_SetText           *syscall.LazyProc
	xBkObj_SetFont           *syscall.LazyProc
	xBkObj_SetTextAlign      *syscall.LazyProc
	xBkObj_GetMargin         *syscall.LazyProc
	xBkObj_GetAlign          *syscall.LazyProc
	xBkObj_GetImage          *syscall.LazyProc
	xBkObj_GetRotate         *syscall.LazyProc
	xBkObj_GetFillColor      *syscall.LazyProc
	xBkObj_GetBorderColor    *syscall.LazyProc
	xBkObj_GetBorderWidth    *syscall.LazyProc
	xBkObj_GetRectRoundAngle *syscall.LazyProc
	xBkObj_IsFill            *syscall.LazyProc
	xBkObj_IsBorder          *syscall.LazyProc
	xBkObj_GetText           *syscall.LazyProc
	xBkObj_GetFont           *syscall.LazyProc
	xBkObj_GetTextAlign      *syscall.LazyProc
)

// SetXcguiPath 手动设置xcgui.dll的路径. 未设置时, 默认值为'xcgui.dll'.
//	@param XcguiPath dll文件名, 不是目录.
//	@return error 如果出错, 要么你输入的文件不存在, 要么你输入的不是dll文件.
//
func SetXcguiPath(XcguiPath string) error {
	// 判断是否为dll文件
	if len(XcguiPath) < 5 {
		return errors.New("XcguiPath 必须是一个dll文件")
	}

	// 判断文件是否存在
	if err := pathExists(XcguiPath); err != nil {
		return errors.New("XcguiPath 指向的文件不存在")
	}
	xcguiPath = XcguiPath
	return nil
}

// GetXcguiPath 获取设置的xcgui.dll的路径.
//	@return string
//
func GetXcguiPath() string {
	return xcguiPath
}

// pathExists 判断文件或文件夹是否存在.
//	@param path 文件或文件夹.
//	@return error
//
func pathExists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
		return err
	}
	return nil
}

// LoadXCGUI 将从 xcguiPath 加载xcgui.dll. xcguiPath 的默认值是'xcgui.dll'.
//	如果你想要更改xcgui.dll的路径, 那么请在调用本函数之前调用 xc.SetXcguiPath().
//
//	注意: app.New() 函数内部会自动调用 xc.LoadXCGUI().
//	所以一般是不需要手动调用的, 除非你没有使用 app.New() 函数, 而是使用了 xc.XInitXCGUI(), 那么你需要在 xc.XInitXCGUI() 之前调用 xc.LoadXCGUI().
//
func LoadXCGUI() {
	// Library.
	xcgui = syscall.NewLazyDLL(xcguiPath)

	// Global Functions.
	xC_MessageBox = xcgui.NewProc("XC_MessageBox")
	xC_SendMessage = xcgui.NewProc("XC_SendMessage")
	xC_PostMessage = xcgui.NewProc("XC_PostMessage")
	xC_CallUiThread = xcgui.NewProc("XC_CallUiThread")
	xC_wtoa = xcgui.NewProc("XC_wtoa")
	xC_atow = xcgui.NewProc("XC_atow")
	xC_utf8tow = xcgui.NewProc("XC_utf8tow")
	xC_utf8towEx = xcgui.NewProc("XC_utf8towEx")
	xC_utf8toa = xcgui.NewProc("XC_utf8toa")
	xC_atoutf8 = xcgui.NewProc("XC_atoutf8")
	xC_wtoutf8 = xcgui.NewProc("XC_wtoutf8")
	xC_wtoutf8Ex = xcgui.NewProc("XC_wtoutf8Ex")
	/*
			xC_itoa = xcgui.NewProc("XC_itoa")
			xC_itow = xcgui.NewProc("XC_itow")
			xC_ftoa = xcgui.NewProc("XC_ftoa")
			xC_ftow = xcgui.NewProc("XC_ftow")
			xC_UnicodeToAnsi = xcgui.NewProc("XC_UnicodeToAnsi")
		   	xC_AnsiToUnicode = xcgui.NewProc("XC_AnsiToUnicode")
	*/
	xC_DebugToFileInfo = xcgui.NewProc("XC_DebugToFileInfo")
	xDebug_OutputDebugStringW = xcgui.NewProc("XDebug_OutputDebugStringW")
	xDebug_Set_OutputDebugString_UTF8 = xcgui.NewProc("XDebug_Set_OutputDebugString_UTF8")
	xDebug_Print = xcgui.NewProc("XDebug_Print")
	xC_IsHELE = xcgui.NewProc("XC_IsHELE")
	xC_IsHWINDOW = xcgui.NewProc("XC_IsHWINDOW")
	xC_IsShape = xcgui.NewProc("XC_IsShape")
	xC_IsHXCGUI = xcgui.NewProc("XC_IsHXCGUI")
	xC_hWindowFromHWnd = xcgui.NewProc("XC_hWindowFromHWnd")
	xC_SetActivateTopWindow = xcgui.NewProc("XC_SetActivateTopWindow")
	xC_SetProperty = xcgui.NewProc("XC_SetProperty")
	xC_GetProperty = xcgui.NewProc("XC_GetProperty")
	xC_RegisterWindowClassName = xcgui.NewProc("XC_RegisterWindowClassName")
	xC_IsSViewExtend = xcgui.NewProc("XC_IsSViewExtend")
	xC_GetObjectType = xcgui.NewProc("XC_GetObjectType")
	xC_GetObjectByID = xcgui.NewProc("XC_GetObjectByID")
	xC_GetObjectByIDName = xcgui.NewProc("XC_GetObjectByIDName")
	xC_GetObjectByUID = xcgui.NewProc("XC_GetObjectByUID")
	xC_GetObjectByUIDName = xcgui.NewProc("XC_GetObjectByUIDName")
	xC_GetObjectByName = xcgui.NewProc("XC_GetObjectByName")
	xC_SetPaintFrequency = xcgui.NewProc("XC_SetPaintFrequency")
	xC_SetTextRenderingHint = xcgui.NewProc("XC_SetTextRenderingHint")
	xC_EnableGdiDrawText = xcgui.NewProc("XC_EnableGdiDrawText")
	xC_RectInRect = xcgui.NewProc("XC_RectInRect")
	xC_CombineRect = xcgui.NewProc("XC_CombineRect")
	xC_ShowLayoutFrame = xcgui.NewProc("XC_ShowLayoutFrame")
	xC_EnableDebugFile = xcgui.NewProc("XC_EnableDebugFile")
	xC_EnableResMonitor = xcgui.NewProc("XC_EnableResMonitor")
	xC_SetLayoutFrameColor = xcgui.NewProc("XC_SetLayoutFrameColor")
	xC_EnableErrorMessageBox = xcgui.NewProc("XC_EnableErrorMessageBox")
	xC_EnableAutoExitApp = xcgui.NewProc("XC_EnableAutoExitApp")
	xC_GetTextSize = xcgui.NewProc("XC_GetTextSize")
	xC_GetTextShowSize = xcgui.NewProc("XC_GetTextShowSize")
	xC_GetTextShowSizeEx = xcgui.NewProc("XC_GetTextShowSizeEx")
	xC_GetTextShowRect = xcgui.NewProc("XC_GetTextShowRect")
	xC_GetDefaultFont = xcgui.NewProc("XC_GetDefaultFont")
	xC_SetDefaultFont = xcgui.NewProc("XC_SetDefaultFont")
	xC_AddFileSearchPath = xcgui.NewProc("XC_AddFileSearchPath")
	xC_InitFont = xcgui.NewProc("XC_InitFont")
	xC_Malloc = xcgui.NewProc("XC_Malloc")
	xC_Free = xcgui.NewProc("XC_Free")
	xC_Alert = xcgui.NewProc("XC_Alert")
	xC_Sys_ShellExecute = xcgui.NewProc("XC_Sys_ShellExecute")
	xC_LoadLibrary = xcgui.NewProc("XC_LoadLibrary")
	xC_GetProcAddress = xcgui.NewProc("XC_GetProcAddress")
	xC_FreeLibrary = xcgui.NewProc("XC_FreeLibrary")
	xC_LoadDll = xcgui.NewProc("XC_LoadDll")
	xInitXCGUI = xcgui.NewProc("XInitXCGUI")
	xRunXCGUI = xcgui.NewProc("XRunXCGUI")
	xExitXCGUI = xcgui.NewProc("XExitXCGUI")
	xC_PostQuitMessage = xcgui.NewProc("XC_PostQuitMessage")
	xC_GetD2dFactory = xcgui.NewProc("XC_GetD2dFactory")
	xC_GetWicFactory = xcgui.NewProc("XC_GetWicFactory")
	xC_GetDWriteFactory = xcgui.NewProc("XC_GetDWriteFactory")
	xC_SetD2dTextRenderingMode = xcgui.NewProc("XC_SetD2dTextRenderingMode")
	xC_IsEnableD2D = xcgui.NewProc("XC_IsEnableD2D")
	xMsg_Create = xcgui.NewProc("XMsg_Create")
	xMsg_CreateEx = xcgui.NewProc("XMsg_CreateEx")
	xC_ShowSvgFrame = xcgui.NewProc("XC_ShowSvgFrame")

	// UI Designer.
	xC_LoadLayout = xcgui.NewProc("XC_LoadLayout")
	xC_LoadLayoutZip = xcgui.NewProc("XC_LoadLayoutZip")
	xC_LoadLayoutZipMem = xcgui.NewProc("XC_LoadLayoutZipMem")
	// xC_LoadLayoutFromString = xcgui.NewProc("XC_LoadLayoutFromString")
	xC_LoadLayoutFromStringUtf8 = xcgui.NewProc("XC_LoadLayoutFromStringUtf8")
	xC_LoadStyle = xcgui.NewProc("XC_LoadStyle")
	xC_LoadStyleZip = xcgui.NewProc("XC_LoadStyleZip")
	xC_LoadStyleZipMem = xcgui.NewProc("XC_LoadStyleZipMem")
	xC_LoadResource = xcgui.NewProc("XC_LoadResource")
	xC_LoadResourceZip = xcgui.NewProc("XC_LoadResourceZip")
	xC_LoadResourceZipMem = xcgui.NewProc("XC_LoadResourceZipMem")
	xC_LoadResourceFromStringUtf8 = xcgui.NewProc("XC_LoadResourceFromStringUtf8")
	xC_LoadStyleFromStringW = xcgui.NewProc("XC_LoadStyleFromStringW")
	/* 	xC_LoadResourceFromString = xcgui.NewProc("XC_LoadResourceFromString")
	   	xC_LoadStyleFromString = xcgui.NewProc("XC_LoadStyleFromString") */

	// Window.
	xWnd_Create = xcgui.NewProc("XWnd_Create")
	xWnd_CreateEx = xcgui.NewProc("XWnd_CreateEx")
	xWnd_RegEventC = xcgui.NewProc("XWnd_RegEventC")
	xWnd_RegEventC1 = xcgui.NewProc("XWnd_RegEventC1")
	xWnd_RemoveEventC = xcgui.NewProc("XWnd_RemoveEventC")
	xWnd_ShowWindow = xcgui.NewProc("XWnd_ShowWindow")
	xWnd_SetTop = xcgui.NewProc("XWnd_SetTop")
	xWnd_InsertChild = xcgui.NewProc("XWnd_InsertChild")
	xWnd_Redraw = xcgui.NewProc("XWnd_Redraw")
	xWnd_RedrawRect = xcgui.NewProc("XWnd_RedrawRect")
	xWnd_SetFocusEle = xcgui.NewProc("XWnd_SetFocusEle")
	xWnd_GetFocusEle = xcgui.NewProc("XWnd_GetFocusEle")
	xWnd_GetStayEle = xcgui.NewProc("XWnd_GetStayEle")
	xWnd_DrawWindow = xcgui.NewProc("XWnd_DrawWindow")
	xWnd_Center = xcgui.NewProc("XWnd_Center")
	xWnd_CenterEx = xcgui.NewProc("XWnd_CenterEx")
	xWnd_SetCursor = xcgui.NewProc("XWnd_SetCursor")
	xWnd_GetCursor = xcgui.NewProc("XWnd_GetCursor")
	xWnd_GetHWND = xcgui.NewProc("XWnd_GetHWND")
	xWnd_EnableDragBorder = xcgui.NewProc("XWnd_EnableDragBorder")
	xWnd_EnableDragWindow = xcgui.NewProc("XWnd_EnableDragWindow")
	xWnd_EnableDragCaption = xcgui.NewProc("XWnd_EnableDragCaption")
	xWnd_EnableDrawBk = xcgui.NewProc("XWnd_EnableDrawBk")
	xWnd_EnableAutoFocus = xcgui.NewProc("XWnd_EnableAutoFocus")
	xWnd_EnableMaxWindow = xcgui.NewProc("XWnd_EnableMaxWindow")
	xWnd_EnablemLimitWindowSize = xcgui.NewProc("XWnd_EnablemLimitWindowSize")
	xWnd_EnableLayout = xcgui.NewProc("XWnd_EnableLayout")
	xWnd_EnableLayoutOverlayBorder = xcgui.NewProc("XWnd_EnableLayoutOverlayBorder")
	xWnd_ShowLayoutFrame = xcgui.NewProc("XWnd_ShowLayoutFrame")
	xWnd_IsEnableLayout = xcgui.NewProc("XWnd_IsEnableLayout")
	xWnd_IsMaxWindow = xcgui.NewProc("XWnd_IsMaxWindow")
	xWnd_SetCaptureEle = xcgui.NewProc("XWnd_SetCaptureEle")
	xWnd_GetCaptureEle = xcgui.NewProc("XWnd_GetCaptureEle")
	xWnd_GetDrawRect = xcgui.NewProc("XWnd_GetDrawRect")
	xWnd_SetCursorSys = xcgui.NewProc("XWnd_SetCursorSys")
	xWnd_SetFont = xcgui.NewProc("XWnd_SetFont")
	xWnd_SetTextColor = xcgui.NewProc("XWnd_SetTextColor")
	xWnd_GetTextColor = xcgui.NewProc("XWnd_GetTextColor")
	xWnd_GetTextColorEx = xcgui.NewProc("XWnd_GetTextColorEx")
	xWnd_SetID = xcgui.NewProc("XWnd_SetID")
	xWnd_GetID = xcgui.NewProc("XWnd_GetID")
	xWnd_SetName = xcgui.NewProc("XWnd_SetName")
	xWnd_GetName = xcgui.NewProc("XWnd_GetName")
	xWnd_SetBorderSize = xcgui.NewProc("XWnd_SetBorderSize")
	xWnd_GetBorderSize = xcgui.NewProc("XWnd_GetBorderSize")
	xWnd_SetPadding = xcgui.NewProc("XWnd_SetPadding")
	xWnd_SetDragBorderSize = xcgui.NewProc("XWnd_SetDragBorderSize")
	xWnd_GetDragBorderSize = xcgui.NewProc("XWnd_GetDragBorderSize")
	xWnd_SetMinimumSize = xcgui.NewProc("XWnd_SetMinimumSize")
	xWnd_HitChildEle = xcgui.NewProc("XWnd_HitChildEle")
	xWnd_GetChildCount = xcgui.NewProc("XWnd_GetChildCount")
	xWnd_GetChildByIndex = xcgui.NewProc("XWnd_GetChildByIndex")
	xWnd_GetChildByID = xcgui.NewProc("XWnd_GetChildByID")
	xWnd_GetChild = xcgui.NewProc("XWnd_GetChild")
	xWnd_CloseWindow = xcgui.NewProc("XWnd_CloseWindow")
	xWnd_AdjustLayout = xcgui.NewProc("XWnd_AdjustLayout")
	xWnd_AdjustLayoutEx = xcgui.NewProc("XWnd_AdjustLayoutEx")
	xWnd_CreateCaret = xcgui.NewProc("XWnd_CreateCaret")
	xWnd_SetCaretPos = xcgui.NewProc("XWnd_SetCaretPos")
	xWnd_SetCaretColor = xcgui.NewProc("XWnd_SetCaretColor")
	xWnd_ShowCaret = xcgui.NewProc("XWnd_ShowCaret")
	xWnd_DestroyCaret = xcgui.NewProc("XWnd_DestroyCaret")
	xWnd_GetCaretHELE = xcgui.NewProc("XWnd_GetCaretHELE")
	xWnd_GetClientRect = xcgui.NewProc("XWnd_GetClientRect")
	xWnd_GetBodyRect = xcgui.NewProc("XWnd_GetBodyRect")
	xWnd_GetLayoutRect = xcgui.NewProc("XWnd_GetLayoutRect")
	xWnd_SetPosition = xcgui.NewProc("XWnd_SetPosition")
	xWnd_GetRect = xcgui.NewProc("XWnd_GetRect")
	xWnd_SetRect = xcgui.NewProc("XWnd_SetRect")
	xWnd_MaxWindow = xcgui.NewProc("XWnd_MaxWindow")
	xWnd_SetTimer = xcgui.NewProc("XWnd_SetTimer")
	xWnd_KillTimer = xcgui.NewProc("XWnd_KillTimer")
	xWnd_SetXCTimer = xcgui.NewProc("XWnd_SetXCTimer")
	xWnd_KillXCTimer = xcgui.NewProc("XWnd_KillXCTimer")
	xWnd_GetBkManager = xcgui.NewProc("XWnd_GetBkManager")
	xWnd_GetBkManagerEx = xcgui.NewProc("XWnd_GetBkManagerEx")
	xWnd_SetBkMagager = xcgui.NewProc("XWnd_SetBkMagager")
	xWnd_SetTransparentType = xcgui.NewProc("XWnd_SetTransparentType")
	xWnd_SetTransparentAlpha = xcgui.NewProc("XWnd_SetTransparentAlpha")
	xWnd_SetTransparentColor = xcgui.NewProc("XWnd_SetTransparentColor")
	xWnd_SetShadowInfo = xcgui.NewProc("XWnd_SetShadowInfo")
	xWnd_GetShadowInfo = xcgui.NewProc("XWnd_GetShadowInfo")
	xWnd_GetTransparentType = xcgui.NewProc("XWnd_GetTransparentType")
	xWnd_Attach = xcgui.NewProc("XWnd_Attach")
	xWnd_EnableDragFiles = xcgui.NewProc("XWnd_EnableDragFiles")
	xWnd_Show = xcgui.NewProc("XWnd_Show")
	xWnd_GetCaretInfo = xcgui.NewProc("XWnd_GetCaretInfo")
	xWnd_SetIcon = xcgui.NewProc("XWnd_SetIcon")
	xWnd_SetTitle = xcgui.NewProc("XWnd_SetTitle")
	xWnd_SetTitleColor = xcgui.NewProc("XWnd_SetTitleColor")
	xWnd_GetButton = xcgui.NewProc("XWnd_GetButton")
	xWnd_GetIcon = xcgui.NewProc("XWnd_GetIcon")
	xWnd_GetTitle = xcgui.NewProc("XWnd_GetTitle")
	xWnd_GetTitleColor = xcgui.NewProc("XWnd_GetTitleColor")
	xWnd_AddBkBorder = xcgui.NewProc("XWnd_AddBkBorder")
	xWnd_AddBkFill = xcgui.NewProc("XWnd_AddBkFill")
	xWnd_AddBkImage = xcgui.NewProc("XWnd_AddBkImage")
	xWnd_GetBkInfoCount = xcgui.NewProc("XWnd_GetBkInfoCount")
	xWnd_ClearBkInfo = xcgui.NewProc("XWnd_ClearBkInfo")
	xWnd_SetBkInfo = xcgui.NewProc("XWnd_SetBkInfo")
	xWnd_IsDragCaption = xcgui.NewProc("XWnd_IsDragCaption")
	xWnd_IsDragWindow = xcgui.NewProc("XWnd_IsDragWindow")
	xWnd_IsDragBorder = xcgui.NewProc("XWnd_IsDragBorder")
	xWnd_SetCaptionMargin = xcgui.NewProc("XWnd_SetCaptionMargin")

	// Widget.
	xWidget_IsShow = xcgui.NewProc("XWidget_IsShow")
	xWidget_Show = xcgui.NewProc("XWidget_Show")
	xWidget_EnableLayoutControl = xcgui.NewProc("XWidget_EnableLayoutControl")
	xWidget_IsLayoutControl = xcgui.NewProc("XWidget_IsLayoutControl")
	xWidget_GetParentEle = xcgui.NewProc("XWidget_GetParentEle")
	xWidget_GetParent = xcgui.NewProc("XWidget_GetParent")
	xWidget_GetHWND = xcgui.NewProc("XWidget_GetHWND")
	xWidget_GetHWINDOW = xcgui.NewProc("XWidget_GetHWINDOW")
	xWidget_LayoutItem_EnableWrap = xcgui.NewProc("XWidget_LayoutItem_EnableWrap")
	xWidget_LayoutItem_EnableSwap = xcgui.NewProc("XWidget_LayoutItem_EnableSwap")
	xWidget_LayoutItem_EnableFloat = xcgui.NewProc("XWidget_LayoutItem_EnableFloat")
	xWidget_LayoutItem_SetWidth = xcgui.NewProc("XWidget_LayoutItem_SetWidth")
	xWidget_LayoutItem_SetHeight = xcgui.NewProc("XWidget_LayoutItem_SetHeight")
	xWidget_LayoutItem_GetWidth = xcgui.NewProc("XWidget_LayoutItem_GetWidth")
	xWidget_LayoutItem_GetHeight = xcgui.NewProc("XWidget_LayoutItem_GetHeight")
	xWidget_LayoutItem_SetAlign = xcgui.NewProc("XWidget_LayoutItem_SetAlign")
	xWidget_LayoutItem_SetMargin = xcgui.NewProc("XWidget_LayoutItem_SetMargin")
	xWidget_LayoutItem_GetMargin = xcgui.NewProc("XWidget_LayoutItem_GetMargin")
	xWidget_LayoutItem_SetMinSize = xcgui.NewProc("XWidget_LayoutItem_SetMinSize")
	xWidget_LayoutItem_SetPosition = xcgui.NewProc("XWidget_LayoutItem_SetPosition")
	xWidget_SetID = xcgui.NewProc("XWidget_SetID")
	xWidget_GetID = xcgui.NewProc("XWidget_GetID")
	xWidget_SetUID = xcgui.NewProc("XWidget_SetUID")
	xWidget_GetUID = xcgui.NewProc("XWidget_GetUID")
	xWidget_SetName = xcgui.NewProc("XWidget_SetName")
	xWidget_GetName = xcgui.NewProc("XWidget_GetName")

	// xUI.
	xUI_SetStyle = xcgui.NewProc("XUI_SetStyle")
	xUI_GetStyle = xcgui.NewProc("XUI_GetStyle")
	xUI_EnableCSS = xcgui.NewProc("XUI_EnableCSS")
	xUI_SetCssName = xcgui.NewProc("XUI_SetCssName")
	xUI_GetCssName = xcgui.NewProc("XUI_GetCssName")

	// Button.
	xBtn_Create = xcgui.NewProc("XBtn_Create")
	xBtn_IsCheck = xcgui.NewProc("XBtn_IsCheck")
	xBtn_SetCheck = xcgui.NewProc("XBtn_SetCheck")
	xBtn_SetState = xcgui.NewProc("XBtn_SetState")
	xBtn_GetState = xcgui.NewProc("XBtn_GetState")
	xBtn_GetStateEx = xcgui.NewProc("XBtn_GetStateEx")
	xBtn_SetTypeEx = xcgui.NewProc("XBtn_SetTypeEx")
	xBtn_SetGroupID = xcgui.NewProc("XBtn_SetGroupID")
	xBtn_GetGroupID = xcgui.NewProc("XBtn_GetGroupID")
	xBtn_SetBindEle = xcgui.NewProc("XBtn_SetBindEle")
	xBtn_GetBindEle = xcgui.NewProc("XBtn_GetBindEle")
	xBtn_SetTextAlign = xcgui.NewProc("XBtn_SetTextAlign")
	xBtn_GetTextAlign = xcgui.NewProc("XBtn_GetTextAlign")
	xBtn_SetIconAlign = xcgui.NewProc("XBtn_SetIconAlign")
	xBtn_SetOffset = xcgui.NewProc("XBtn_SetOffset")
	xBtn_SetOffsetIcon = xcgui.NewProc("XBtn_SetOffsetIcon")
	xBtn_SetIconSpace = xcgui.NewProc("XBtn_SetIconSpace")
	xBtn_SetText = xcgui.NewProc("XBtn_SetText")
	xBtn_GetText = xcgui.NewProc("XBtn_GetText")
	xBtn_SetIcon = xcgui.NewProc("XBtn_SetIcon")
	xBtn_SetIconDisable = xcgui.NewProc("XBtn_SetIconDisable")
	xBtn_GetIcon = xcgui.NewProc("XBtn_GetIcon")
	xBtn_AddAnimationFrame = xcgui.NewProc("XBtn_AddAnimationFrame")
	xBtn_EnableAnimation = xcgui.NewProc("XBtn_EnableAnimation")

	// Element.
	xEle_Create = xcgui.NewProc("XEle_Create")
	xEle_RegEventC = xcgui.NewProc("XEle_RegEventC")
	xEle_RegEventC1 = xcgui.NewProc("XEle_RegEventC1")
	xEle_RemoveEventC = xcgui.NewProc("XEle_RemoveEventC")
	xEle_SendEvent = xcgui.NewProc("XEle_SendEvent")
	xEle_PostEvent = xcgui.NewProc("XEle_PostEvent")
	xEle_GetRect = xcgui.NewProc("XEle_GetRect")
	xEle_GetRectLogic = xcgui.NewProc("XEle_GetRectLogic")
	xEle_GetClientRect = xcgui.NewProc("XEle_GetClientRect")
	xEle_SetWidth = xcgui.NewProc("XEle_SetWidth")
	xEle_SetHeight = xcgui.NewProc("XEle_SetHeight")
	xEle_GetWidth = xcgui.NewProc("XEle_GetWidth")
	xEle_GetHeight = xcgui.NewProc("XEle_GetHeight")
	xEle_RectWndClientToEleClient = xcgui.NewProc("XEle_RectWndClientToEleClient")
	xEle_PointWndClientToEleClient = xcgui.NewProc("XEle_PointWndClientToEleClient")
	xEle_RectClientToWndClient = xcgui.NewProc("XEle_RectClientToWndClient")
	xEle_PointClientToWndClient = xcgui.NewProc("XEle_PointClientToWndClient")
	xEle_GetWndClientRect = xcgui.NewProc("XEle_GetWndClientRect")
	xEle_GetCursor = xcgui.NewProc("XEle_GetCursor")
	xEle_SetCursor = xcgui.NewProc("XEle_SetCursor")
	xEle_AddChild = xcgui.NewProc("XEle_AddChild")
	xEle_InsertChild = xcgui.NewProc("XEle_InsertChild")
	xEle_SetRect = xcgui.NewProc("XEle_SetRect")
	xEle_SetRectEx = xcgui.NewProc("XEle_SetRectEx")
	xEle_SetRectLogic = xcgui.NewProc("XEle_SetRectLogic")
	xEle_SetPosition = xcgui.NewProc("XEle_SetPosition")
	xEle_SetPositionLogic = xcgui.NewProc("XEle_SetPositionLogic")
	xEle_IsDrawFocus = xcgui.NewProc("XEle_IsDrawFocus")
	xEle_IsEnable = xcgui.NewProc("XEle_IsEnable")
	xEle_IsEnableFocus = xcgui.NewProc("XEle_IsEnableFocus")
	xEle_IsMouseThrough = xcgui.NewProc("XEle_IsMouseThrough")
	xEle_HitChildEle = xcgui.NewProc("XEle_HitChildEle")
	xEle_IsBkTransparent = xcgui.NewProc("XEle_IsBkTransparent")
	xEle_IsEnableEvent_XE_PAINT_END = xcgui.NewProc("XEle_IsEnableEvent_XE_PAINT_END")
	xEle_IsKeyTab = xcgui.NewProc("XEle_IsKeyTab")
	xEle_IsSwitchFocus = xcgui.NewProc("XEle_IsSwitchFocus")
	xEle_IsEnable_XE_MOUSEWHEEL = xcgui.NewProc("XEle_IsEnable_XE_MOUSEWHEEL")
	xEle_IsChildEle = xcgui.NewProc("XEle_IsChildEle")
	xEle_IsEnableCanvas = xcgui.NewProc("XEle_IsEnableCanvas")
	xEle_IsFocus = xcgui.NewProc("XEle_IsFocus")
	xEle_IsFocusEx = xcgui.NewProc("XEle_IsFocusEx")
	xEle_Enable = xcgui.NewProc("XEle_Enable")
	xEle_EnableFocus = xcgui.NewProc("XEle_EnableFocus")
	xEle_EnableDrawFocus = xcgui.NewProc("XEle_EnableDrawFocus")
	xEle_EnableDrawBorder = xcgui.NewProc("XEle_EnableDrawBorder")
	xEle_EnableCanvas = xcgui.NewProc("XEle_EnableCanvas")
	xEle_EnableEvent_XE_PAINT_END = xcgui.NewProc("XEle_EnableEvent_XE_PAINT_END")
	xEle_EnableBkTransparent = xcgui.NewProc("XEle_EnableBkTransparent")
	xEle_EnableMouseThrough = xcgui.NewProc("XEle_EnableMouseThrough")
	xEle_EnableKeyTab = xcgui.NewProc("XEle_EnableKeyTab")
	xEle_EnableSwitchFocus = xcgui.NewProc("XEle_EnableSwitchFocus")
	xEle_EnableEvent_XE_MOUSEWHEEL = xcgui.NewProc("XEle_EnableEvent_XE_MOUSEWHEEL")
	xEle_Remove = xcgui.NewProc("XEle_Remove")
	xEle_SetZOrder = xcgui.NewProc("XEle_SetZOrder")
	xEle_SetZOrderEx = xcgui.NewProc("XEle_SetZOrderEx")
	xEle_GetZOrder = xcgui.NewProc("XEle_GetZOrder")
	xEle_EnableTopmost = xcgui.NewProc("XEle_EnableTopmost")
	xEle_Redraw = xcgui.NewProc("XEle_Redraw")
	xEle_RedrawRect = xcgui.NewProc("XEle_RedrawRect")
	xEle_GetChildCount = xcgui.NewProc("XEle_GetChildCount")
	xEle_GetChildByIndex = xcgui.NewProc("XEle_GetChildByIndex")
	xEle_GetChildByID = xcgui.NewProc("XEle_GetChildByID")
	xEle_SetBorderSize = xcgui.NewProc("XEle_SetBorderSize")
	xEle_GetBorderSize = xcgui.NewProc("XEle_GetBorderSize")
	xEle_SetPadding = xcgui.NewProc("XEle_SetPadding")
	xEle_GetPadding = xcgui.NewProc("XEle_GetPadding")
	xEle_SetDragBorder = xcgui.NewProc("XEle_SetDragBorder")
	xEle_SetDragBorderBindEle = xcgui.NewProc("XEle_SetDragBorderBindEle")
	xEle_SetMinSize = xcgui.NewProc("XEle_SetMinSize")
	xEle_SetMaxSize = xcgui.NewProc("XEle_SetMaxSize")
	xEle_SetLockScroll = xcgui.NewProc("XEle_SetLockScroll")
	xEle_SetTextColor = xcgui.NewProc("XEle_SetTextColor")
	xEle_GetTextColor = xcgui.NewProc("XEle_GetTextColor")
	xEle_GetTextColorEx = xcgui.NewProc("XEle_GetTextColorEx")
	xEle_SetFocusBorderColor = xcgui.NewProc("XEle_SetFocusBorderColor")
	xEle_GetFocusBorderColor = xcgui.NewProc("XEle_GetFocusBorderColor")
	xEle_SetFont = xcgui.NewProc("XEle_SetFont")
	xEle_GetFont = xcgui.NewProc("XEle_GetFont")
	xEle_GetFontEx = xcgui.NewProc("XEle_GetFontEx")
	xEle_SetAlpha = xcgui.NewProc("XEle_SetAlpha")
	xEle_Destroy = xcgui.NewProc("XEle_Destroy")
	xEle_AddBkBorder = xcgui.NewProc("XEle_AddBkBorder")
	xEle_AddBkFill = xcgui.NewProc("XEle_AddBkFill")
	xEle_AddBkImage = xcgui.NewProc("XEle_AddBkImage")
	xEle_GetBkInfoCount = xcgui.NewProc("XEle_GetBkInfoCount")
	xEle_ClearBkInfo = xcgui.NewProc("XEle_ClearBkInfo")
	xEle_GetBkManager = xcgui.NewProc("XEle_GetBkManager")
	xEle_GetBkManagerEx = xcgui.NewProc("XEle_GetBkManagerEx")
	xEle_SetBkManager = xcgui.NewProc("XEle_SetBkManager")
	xEle_GetStateFlags = xcgui.NewProc("XEle_GetStateFlags")
	xEle_DrawFocus = xcgui.NewProc("XEle_DrawFocus")
	xEle_DrawEle = xcgui.NewProc("XEle_DrawEle")
	xEle_SetUserData = xcgui.NewProc("XEle_SetUserData")
	xEle_GetUserData = xcgui.NewProc("XEle_GetUserData")
	xEle_GetContentSize = xcgui.NewProc("XEle_GetContentSize")
	xEle_SetCapture = xcgui.NewProc("XEle_SetCapture")
	xEle_EnableTransparentChannel = xcgui.NewProc("XEle_EnableTransparentChannel")
	xEle_SetXCTimer = xcgui.NewProc("XEle_SetXCTimer")
	xEle_KillXCTimer = xcgui.NewProc("XEle_KillXCTimer")
	xEle_SetToolTip = xcgui.NewProc("XEle_SetToolTip")
	xEle_SetToolTipEx = xcgui.NewProc("XEle_SetToolTipEx")
	xEle_GetToolTip = xcgui.NewProc("XEle_GetToolTip")
	xEle_PopupToolTip = xcgui.NewProc("XEle_PopupToolTip")
	xEle_AdjustLayout = xcgui.NewProc("XEle_AdjustLayout")
	xEle_AdjustLayoutEx = xcgui.NewProc("XEle_AdjustLayoutEx")
	xEle_GetAlpha = xcgui.NewProc("XEle_GetAlpha")
	xEle_GetPosition = xcgui.NewProc("XEle_GetPosition")
	xEle_SetSize = xcgui.NewProc("XEle_SetSize")
	xEle_GetSize = xcgui.NewProc("XEle_GetSize")
	xEle_SetBkInfo = xcgui.NewProc("XEle_SetBkInfo")

	// FreameWindow.
	xFrameWnd_Create = xcgui.NewProc("XFrameWnd_Create")
	xFrameWnd_CreateEx = xcgui.NewProc("XFrameWnd_CreateEx")
	xFrameWnd_GetLayoutAreaRect = xcgui.NewProc("XFrameWnd_GetLayoutAreaRect")
	xFrameWnd_SetView = xcgui.NewProc("XFrameWnd_SetView")
	xFrameWnd_SetPaneSplitBarColor = xcgui.NewProc("XFrameWnd_SetPaneSplitBarColor")
	xFrameWnd_SetTabBarHeight = xcgui.NewProc("XFrameWnd_SetTabBarHeight")
	xFrameWnd_SaveLayoutToFile = xcgui.NewProc("XFrameWnd_SaveLayoutToFile")
	xFrameWnd_LoadLayoutFile = xcgui.NewProc("XFrameWnd_LoadLayoutFile")
	xFrameWnd_AddPane = xcgui.NewProc("XFrameWnd_AddPane")
	xFrameWnd_MergePane = xcgui.NewProc("XFrameWnd_MergePane")
	xFrameWnd_Attach = xcgui.NewProc("XFrameWnd_Attach")
	xFrameWnd_GetDragFloatWndTopFlag = xcgui.NewProc("XFrameWnd_GetDragFloatWndTopFlag")

	// Menu.
	xMenu_Create = xcgui.NewProc("XMenu_Create")
	xMenu_AddItem = xcgui.NewProc("XMenu_AddItem")
	xMenu_AddItemIcon = xcgui.NewProc("XMenu_AddItemIcon")
	xMenu_InsertItem = xcgui.NewProc("XMenu_InsertItem")
	xMenu_InsertItemIcon = xcgui.NewProc("XMenu_InsertItemIcon")
	xMenu_GetFirstChildItem = xcgui.NewProc("XMenu_GetFirstChildItem")
	xMenu_GetEndChildItem = xcgui.NewProc("XMenu_GetEndChildItem")
	xMenu_GetPrevSiblingItem = xcgui.NewProc("XMenu_GetPrevSiblingItem")
	xMenu_GetNextSiblingItem = xcgui.NewProc("XMenu_GetNextSiblingItem")
	xMenu_GetParentItem = xcgui.NewProc("XMenu_GetParentItem")
	xMenu_SetAutoDestroy = xcgui.NewProc("XMenu_SetAutoDestroy")
	xMenu_EnableDrawBackground = xcgui.NewProc("XMenu_EnableDrawBackground")
	xMenu_EnableDrawItem = xcgui.NewProc("XMenu_EnableDrawItem")
	xMenu_Popup = xcgui.NewProc("XMenu_Popup")
	xMenu_DestroyMenu = xcgui.NewProc("XMenu_DestroyMenu")
	xMenu_CloseMenu = xcgui.NewProc("XMenu_CloseMenu")
	xMenu_SetBkImage = xcgui.NewProc("XMenu_SetBkImage")
	xMenu_SetItemText = xcgui.NewProc("XMenu_SetItemText")
	xMenu_GetItemText = xcgui.NewProc("XMenu_GetItemText")
	xMenu_GetItemTextLength = xcgui.NewProc("XMenu_GetItemTextLength")
	xMenu_SetItemIcon = xcgui.NewProc("XMenu_SetItemIcon")
	xMenu_SetItemFlags = xcgui.NewProc("XMenu_SetItemFlags")
	xMenu_SetItemHeight = xcgui.NewProc("XMenu_SetItemHeight")
	xMenu_GetItemHeight = xcgui.NewProc("XMenu_GetItemHeight")
	xMenu_SetBorderColor = xcgui.NewProc("XMenu_SetBorderColor")
	xMenu_SetBorderSize = xcgui.NewProc("XMenu_SetBorderSize")
	xMenu_GetLeftWidth = xcgui.NewProc("XMenu_GetLeftWidth")
	xMenu_GetLeftSpaceText = xcgui.NewProc("XMenu_GetLeftSpaceText")
	xMenu_GetItemCount = xcgui.NewProc("XMenu_GetItemCount")
	xMenu_SetItemCheck = xcgui.NewProc("XMenu_SetItemCheck")
	xMenu_IsItemCheck = xcgui.NewProc("XMenu_IsItemCheck")
	xMenu_SetItemWidth = xcgui.NewProc("XMenu_SetItemWidth")

	// ModalWindow.
	xModalWnd_Create = xcgui.NewProc("XModalWnd_Create")
	xModalWnd_CreateEx = xcgui.NewProc("XModalWnd_CreateEx")
	xModalWnd_EnableAutoClose = xcgui.NewProc("XModalWnd_EnableAutoClose")
	xModalWnd_EnableEscClose = xcgui.NewProc("XModalWnd_EnableEscClose")
	xModalWnd_DoModal = xcgui.NewProc("XModalWnd_DoModal")
	xModalWnd_EndModal = xcgui.NewProc("XModalWnd_EndModal")
	xModalWnd_Attach = xcgui.NewProc("XModalWnd_Attach")

	// LayoutBox.
	xLayoutBox_EnableHorizon = xcgui.NewProc("XLayoutBox_EnableHorizon")
	xLayoutBox_EnableAutoWrap = xcgui.NewProc("XLayoutBox_EnableAutoWrap")
	xLayoutBox_EnableOverflowHide = xcgui.NewProc("XLayoutBox_EnableOverflowHide")
	xLayoutBox_SetAlignH = xcgui.NewProc("XLayoutBox_SetAlignH")
	xLayoutBox_SetAlignV = xcgui.NewProc("XLayoutBox_SetAlignV")
	xLayoutBox_SetAlignBaseline = xcgui.NewProc("XLayoutBox_SetAlignBaseline")
	xLayoutBox_SetSpace = xcgui.NewProc("XLayoutBox_SetSpace")
	xLayoutBox_SetSpaceRow = xcgui.NewProc("XLayoutBox_SetSpaceRow")

	// ComboBox.
	xComboBox_Create = xcgui.NewProc("XComboBox_Create")
	xComboBox_SetSelItem = xcgui.NewProc("XComboBox_SetSelItem")
	xComboBox_CreateAdapter = xcgui.NewProc("XComboBox_CreateAdapter")
	xComboBox_BindAdapter = xcgui.NewProc("XComboBox_BindAdapter")
	xComboBox_GetAdapter = xcgui.NewProc("XComboBox_GetAdapter")
	xComboBox_SetBindName = xcgui.NewProc("XComboBox_SetBindName")
	xComboBox_GetButtonRect = xcgui.NewProc("XComboBox_GetButtonRect")
	xComboBox_SetButtonSize = xcgui.NewProc("XComboBox_SetButtonSize")
	xComboBox_SetDropHeight = xcgui.NewProc("XComboBox_SetDropHeight")
	xComboBox_GetDropHeight = xcgui.NewProc("XComboBox_GetDropHeight")
	xComboBox_SetItemTemplateXML = xcgui.NewProc("XComboBox_SetItemTemplateXML")
	xComboBox_SetItemTemplateXMLFromString = xcgui.NewProc("XComboBox_SetItemTemplateXMLFromString")
	xComboBox_EnableDrawButton = xcgui.NewProc("XComboBox_EnableDrawButton")
	xComboBox_EnableEdit = xcgui.NewProc("XComboBox_EnableEdit")
	xComboBox_EnableDropHeightFixed = xcgui.NewProc("XComboBox_EnableDropHeightFixed")
	xComboBox_GetSelItem = xcgui.NewProc("XComboBox_GetSelItem")
	xComboBox_GetState = xcgui.NewProc("XComboBox_GetState")
	xComboBox_AddItemText = xcgui.NewProc("XComboBox_AddItemText")
	xComboBox_AddItemTextEx = xcgui.NewProc("XComboBox_AddItemTextEx")
	xComboBox_AddItemImage = xcgui.NewProc("XComboBox_AddItemImage")
	xComboBox_AddItemImageEx = xcgui.NewProc("XComboBox_AddItemImageEx")
	xComboBox_InsertItemText = xcgui.NewProc("XComboBox_InsertItemText")
	xComboBox_InsertItemTextEx = xcgui.NewProc("XComboBox_InsertItemTextEx")
	xComboBox_InsertItemImage = xcgui.NewProc("XComboBox_InsertItemImage")
	xComboBox_InsertItemImageEx = xcgui.NewProc("XComboBox_InsertItemImageEx")
	xComboBox_SetItemText = xcgui.NewProc("XComboBox_SetItemText")
	xComboBox_SetItemTextEx = xcgui.NewProc("XComboBox_SetItemTextEx")
	xComboBox_SetItemImage = xcgui.NewProc("XComboBox_SetItemImage")
	xComboBox_SetItemImageEx = xcgui.NewProc("XComboBox_SetItemImageEx")
	xComboBox_SetItemInt = xcgui.NewProc("XComboBox_SetItemInt")
	xComboBox_SetItemIntEx = xcgui.NewProc("XComboBox_SetItemIntEx")
	xComboBox_SetItemFloat = xcgui.NewProc("XComboBox_SetItemFloat")
	xComboBox_SetItemFloatEx = xcgui.NewProc("XComboBox_SetItemFloatEx")
	xComboBox_GetItemText = xcgui.NewProc("XComboBox_GetItemText")
	xComboBox_GetItemTextEx = xcgui.NewProc("XComboBox_GetItemTextEx")
	xComboBox_GetItemImage = xcgui.NewProc("XComboBox_GetItemImage")
	xComboBox_GetItemImageEx = xcgui.NewProc("XComboBox_GetItemImageEx")
	xComboBox_GetItemInt = xcgui.NewProc("XComboBox_GetItemInt")
	xComboBox_GetItemIntEx = xcgui.NewProc("XComboBox_GetItemIntEx")
	xComboBox_GetItemFloat = xcgui.NewProc("XComboBox_GetItemFloat")
	xComboBox_GetItemFloatEx = xcgui.NewProc("XComboBox_GetItemFloatEx")
	xComboBox_DeleteItem = xcgui.NewProc("XComboBox_DeleteItem")
	xComboBox_DeleteItemEx = xcgui.NewProc("XComboBox_DeleteItemEx")
	xComboBox_DeleteItemAll = xcgui.NewProc("XComboBox_DeleteItemAll")
	xComboBox_DeleteColumnAll = xcgui.NewProc("XComboBox_DeleteColumnAll")
	xComboBox_GetCount = xcgui.NewProc("XComboBox_GetCount")
	xComboBox_GetCountColumn = xcgui.NewProc("XComboBox_GetCountColumn")
	xComboBox_PopupDropList = xcgui.NewProc("XComboBox_PopupDropList")
	xComboBox_SetItemTemplate = xcgui.NewProc("XComboBox_SetItemTemplate")

	// Adapter.
	xAd_AddRef = xcgui.NewProc("XAd_AddRef")
	xAd_Release = xcgui.NewProc("XAd_Release")
	xAd_GetRefCount = xcgui.NewProc("XAd_GetRefCount")
	xAd_Destroy = xcgui.NewProc("XAd_Destroy")
	xAd_EnableAutoDestroy = xcgui.NewProc("XAd_EnableAutoDestroy")

	// AdapterListView.
	xAdListView_Create = xcgui.NewProc("XAdListView_Create")
	xAdListView_Group_AddColumn = xcgui.NewProc("XAdListView_Group_AddColumn")
	xAdListView_Group_AddItemText = xcgui.NewProc("XAdListView_Group_AddItemText")
	xAdListView_Group_AddItemTextEx = xcgui.NewProc("XAdListView_Group_AddItemTextEx")
	xAdListView_Group_AddItemImage = xcgui.NewProc("XAdListView_Group_AddItemImage")
	xAdListView_Group_AddItemImageEx = xcgui.NewProc("XAdListView_Group_AddItemImageEx")
	xAdListView_Group_SetText = xcgui.NewProc("XAdListView_Group_SetText")
	xAdListView_Group_SetTextEx = xcgui.NewProc("XAdListView_Group_SetTextEx")
	xAdListView_Group_SetImage = xcgui.NewProc("XAdListView_Group_SetImage")
	xAdListView_Group_SetImageEx = xcgui.NewProc("XAdListView_Group_SetImageEx")
	xAdListView_Item_AddColumn = xcgui.NewProc("XAdListView_Item_AddColumn")
	xAdListView_Group_GetCount = xcgui.NewProc("XAdListView_Group_GetCount")
	xAdListView_Item_GetCount = xcgui.NewProc("XAdListView_Item_GetCount")
	xAdListView_Item_AddItemText = xcgui.NewProc("XAdListView_Item_AddItemText")
	xAdListView_Item_AddItemTextEx = xcgui.NewProc("XAdListView_Item_AddItemTextEx")
	xAdListView_Item_AddItemImage = xcgui.NewProc("XAdListView_Item_AddItemImage")
	xAdListView_Item_AddItemImageEx = xcgui.NewProc("XAdListView_Item_AddItemImageEx")
	xAdListView_Group_DeleteItem = xcgui.NewProc("XAdListView_Group_DeleteItem")
	xAdListView_Group_DeleteAllChildItem = xcgui.NewProc("XAdListView_Group_DeleteAllChildItem")
	xAdListView_Item_DeleteItem = xcgui.NewProc("XAdListView_Item_DeleteItem")
	xAdListView_DeleteAll = xcgui.NewProc("XAdListView_DeleteAll")
	xAdListView_DeleteAllGroup = xcgui.NewProc("XAdListView_DeleteAllGroup")
	xAdListView_DeleteAllItem = xcgui.NewProc("XAdListView_DeleteAllItem")
	xAdListView_DeleteColumnGroup = xcgui.NewProc("XAdListView_DeleteColumnGroup")
	xAdListView_DeleteColumnItem = xcgui.NewProc("XAdListView_DeleteColumnItem")
	xAdListView_Item_GetTextEx = xcgui.NewProc("XAdListView_Item_GetTextEx")
	xAdListView_Item_GetImageEx = xcgui.NewProc("XAdListView_Item_GetImageEx")
	xAdListView_Item_SetText = xcgui.NewProc("XAdListView_Item_SetText")
	xAdListView_Item_SetTextEx = xcgui.NewProc("XAdListView_Item_SetTextEx")
	xAdListView_Item_SetImage = xcgui.NewProc("XAdListView_Item_SetImage")
	xAdListView_Item_SetImageEx = xcgui.NewProc("XAdListView_Item_SetImageEx")
	xAdListView_Group_GetText = xcgui.NewProc("XAdListView_Group_GetText")
	xAdListView_Group_GetTextEx = xcgui.NewProc("XAdListView_Group_GetTextEx")
	xAdListView_Group_GetImage = xcgui.NewProc("XAdListView_Group_GetImage")
	xAdListView_Group_GetImageEx = xcgui.NewProc("XAdListView_Group_GetImageEx")
	xAdListView_Item_GetText = xcgui.NewProc("XAdListView_Item_GetText")
	xAdListView_Item_GetImage = xcgui.NewProc("XAdListView_Item_GetImage")

	// AdapterTable.
	xAdTable_Create = xcgui.NewProc("XAdTable_Create")
	xAdTable_Sort = xcgui.NewProc("XAdTable_Sort")
	xAdTable_GetItemDataType = xcgui.NewProc("XAdTable_GetItemDataType")
	xAdTable_GetItemDataTypeEx = xcgui.NewProc("XAdTable_GetItemDataTypeEx")
	xAdTable_AddColumn = xcgui.NewProc("XAdTable_AddColumn")
	xAdTable_SetColumn = xcgui.NewProc("XAdTable_SetColumn")
	xAdTable_AddItemText = xcgui.NewProc("XAdTable_AddItemText")
	xAdTable_AddItemTextEx = xcgui.NewProc("XAdTable_AddItemTextEx")
	xAdTable_AddItemImage = xcgui.NewProc("XAdTable_AddItemImage")
	xAdTable_AddItemImageEx = xcgui.NewProc("XAdTable_AddItemImageEx")
	xAdTable_InsertItemText = xcgui.NewProc("XAdTable_InsertItemText")
	xAdTable_InsertItemTextEx = xcgui.NewProc("XAdTable_InsertItemTextEx")
	xAdTable_InsertItemImage = xcgui.NewProc("XAdTable_InsertItemImage")
	xAdTable_InsertItemImageEx = xcgui.NewProc("XAdTable_InsertItemImageEx")
	xAdTable_SetItemText = xcgui.NewProc("XAdTable_SetItemText")
	xAdTable_SetItemTextEx = xcgui.NewProc("XAdTable_SetItemTextEx")
	xAdTable_SetItemInt = xcgui.NewProc("XAdTable_SetItemInt")
	xAdTable_SetItemIntEx = xcgui.NewProc("XAdTable_SetItemIntEx")
	xAdTable_SetItemFloat = xcgui.NewProc("XAdTable_SetItemFloat")
	xAdTable_SetItemFloatEx = xcgui.NewProc("XAdTable_SetItemFloatEx")
	xAdTable_SetItemImage = xcgui.NewProc("XAdTable_SetItemImage")
	xAdTable_SetItemImageEx = xcgui.NewProc("XAdTable_SetItemImageEx")
	xAdTable_DeleteItem = xcgui.NewProc("XAdTable_DeleteItem")
	xAdTable_DeleteItemEx = xcgui.NewProc("XAdTable_DeleteItemEx")
	xAdTable_DeleteItemAll = xcgui.NewProc("XAdTable_DeleteItemAll")
	xAdTable_DeleteColumnAll = xcgui.NewProc("XAdTable_DeleteColumnAll")
	xAdTable_GetCount = xcgui.NewProc("XAdTable_GetCount")
	xAdTable_GetCountColumn = xcgui.NewProc("XAdTable_GetCountColumn")
	xAdTable_GetItemText = xcgui.NewProc("XAdTable_GetItemText")
	xAdTable_GetItemTextEx = xcgui.NewProc("XAdTable_GetItemTextEx")
	xAdTable_GetItemImage = xcgui.NewProc("XAdTable_GetItemImage")
	xAdTable_GetItemImageEx = xcgui.NewProc("XAdTable_GetItemImageEx")
	xAdTable_GetItemInt = xcgui.NewProc("XAdTable_GetItemInt")
	xAdTable_GetItemIntEx = xcgui.NewProc("XAdTable_GetItemIntEx")
	xAdTable_GetItemFloat = xcgui.NewProc("XAdTable_GetItemFloat")
	xAdTable_GetItemFloatEx = xcgui.NewProc("XAdTable_GetItemFloatEx")
	// AdapterMap.
	xAdMap_Create = xcgui.NewProc("XAdMap_Create")
	xAdMap_AddItemText = xcgui.NewProc("XAdMap_AddItemText")
	xAdMap_AddItemImage = xcgui.NewProc("XAdMap_AddItemImage")
	xAdMap_DeleteItem = xcgui.NewProc("XAdMap_DeleteItem")
	xAdMap_GetCount = xcgui.NewProc("XAdMap_GetCount")
	xAdMap_GetItemText = xcgui.NewProc("XAdMap_GetItemText")
	xAdMap_GetItemImage = xcgui.NewProc("XAdMap_GetItemImage")
	xAdMap_SetItemText = xcgui.NewProc("XAdMap_SetItemText")
	xAdMap_SetItemImage = xcgui.NewProc("XAdMap_SetItemImage")
	// AdapterTree.
	xAdTree_Create = xcgui.NewProc("XAdTree_Create")
	xAdTree_AddColumn = xcgui.NewProc("XAdTree_AddColumn")
	xAdTree_SetColumn = xcgui.NewProc("XAdTree_SetColumn")
	xAdTree_InsertItemText = xcgui.NewProc("XAdTree_InsertItemText")
	xAdTree_InsertItemTextEx = xcgui.NewProc("XAdTree_InsertItemTextEx")
	xAdTree_InsertItemImage = xcgui.NewProc("XAdTree_InsertItemImage")
	xAdTree_InsertItemImageEx = xcgui.NewProc("XAdTree_InsertItemImageEx")
	xAdTree_GetCount = xcgui.NewProc("XAdTree_GetCount")
	xAdTree_GetCountColumn = xcgui.NewProc("XAdTree_GetCountColumn")
	xAdTree_SetItemText = xcgui.NewProc("XAdTree_SetItemText")
	xAdTree_SetItemTextEx = xcgui.NewProc("XAdTree_SetItemTextEx")
	xAdTree_SetItemImage = xcgui.NewProc("XAdTree_SetItemImage")
	xAdTree_SetItemImageEx = xcgui.NewProc("XAdTree_SetItemImageEx")
	xAdTree_GetItemText = xcgui.NewProc("XAdTree_GetItemText")
	xAdTree_GetItemTextEx = xcgui.NewProc("XAdTree_GetItemTextEx")
	xAdTree_GetItemImage = xcgui.NewProc("XAdTree_GetItemImage")
	xAdTree_GetItemImageEx = xcgui.NewProc("XAdTree_GetItemImageEx")
	xAdTree_DeleteItem = xcgui.NewProc("XAdTree_DeleteItem")
	xAdTree_DeleteItemAll = xcgui.NewProc("XAdTree_DeleteItemAll")
	xAdTree_DeleteColumnAll = xcgui.NewProc("XAdTree_DeleteColumnAll")

	// Editor.
	xEditor_Create = xcgui.NewProc("XEditor_Create")
	xEidtor_EnableAutoMatchSpaseSelect = xcgui.NewProc("XEidtor_EnableAutoMatchSpaseSelect")
	xEditor_IsBreakpoint = xcgui.NewProc("XEditor_IsBreakpoint")
	xEditor_SetBreakpoint = xcgui.NewProc("XEditor_SetBreakpoint")
	xEditor_RemoveBreakpoint = xcgui.NewProc("XEditor_RemoveBreakpoint")
	xEditor_ClearBreakpoint = xcgui.NewProc("XEditor_ClearBreakpoint")
	xEditor_SetRunRow = xcgui.NewProc("XEditor_SetRunRow")
	xEditor_GetColor = xcgui.NewProc("XEditor_GetColor")
	xEditor_SetColor = xcgui.NewProc("XEditor_SetColor")
	xEditor_SetStyleKeyword = xcgui.NewProc("XEditor_SetStyleKeyword")
	xEditor_SetStyleFunction = xcgui.NewProc("XEditor_SetStyleFunction")
	xEditor_SetStyleVar = xcgui.NewProc("XEditor_SetStyleVar")
	xEditor_SetStyleDataType = xcgui.NewProc("XEditor_SetStyleDataType")
	xEditor_SetStyleClass = xcgui.NewProc("XEditor_SetStyleClass")
	xEditor_SetStyleMacro = xcgui.NewProc("XEditor_SetStyleMacro")
	xEditor_SetStyleString = xcgui.NewProc("XEditor_SetStyleString")
	xEditor_SetStyleComment = xcgui.NewProc("XEditor_SetStyleComment")
	xEditor_SetStyleNumber = xcgui.NewProc("XEditor_SetStyleNumber")
	xEditor_GetBreakpointCount = xcgui.NewProc("XEditor_GetBreakpointCount")
	xEditor_GetBreakpoints = xcgui.NewProc("XEditor_GetBreakpoints")
	xEditor_SetCurRow = xcgui.NewProc("XEditor_SetCurRow")
	xEditor_GetDepth = xcgui.NewProc("XEditor_GetDepth")
	xEditor_ToExpandRow = xcgui.NewProc("XEditor_ToExpandRow")
	xEditor_ExpandEx = xcgui.NewProc("XEditor_ExpandEx")
	xEditor_ExpandAll = xcgui.NewProc("XEditor_ExpandAll")
	xEditor_Expand = xcgui.NewProc("XEditor_Expand")
	xEditor_AddKeyword = xcgui.NewProc("XEditor_AddKeyword")
	xEditor_AddConst = xcgui.NewProc("XEditor_AddConst")
	xEditor_AddFunction = xcgui.NewProc("XEditor_AddFunction")
	xEditor_AddExcludeDefVarKeyword = xcgui.NewProc("XEditor_AddExcludeDefVarKeyword")
	xEditor_GetExpandState = xcgui.NewProc("XEditor_GetExpandState")
	xEditor_SetExpandState = xcgui.NewProc("XEditor_SetExpandState")
	xEditor_GetIndentation = xcgui.NewProc("XEditor_GetIndentation")
	xEidtor_IsEmptyRow = xcgui.NewProc("XEidtor_IsEmptyRow")
	xEditor_SetAutoMatchMode = xcgui.NewProc("XEditor_SetAutoMatchMode")

	// XEdit.
	xEdit_Create = xcgui.NewProc("XEdit_Create")
	xEdit_CreateEx = xcgui.NewProc("XEdit_CreateEx")
	xEdit_EnableAutoWrap = xcgui.NewProc("XEdit_EnableAutoWrap")
	xEdit_EnableReadOnly = xcgui.NewProc("XEdit_EnableReadOnly")
	xEdit_EnableMultiLine = xcgui.NewProc("XEdit_EnableMultiLine")
	xEdit_EnablePassword = xcgui.NewProc("XEdit_EnablePassword")
	xEdit_EnableAutoSelAll = xcgui.NewProc("XEdit_EnableAutoSelAll")
	xEdit_EnableAutoCancelSel = xcgui.NewProc("XEdit_EnableAutoCancelSel")
	xEdit_IsReadOnly = xcgui.NewProc("XEdit_IsReadOnly")
	xEdit_IsMultiLine = xcgui.NewProc("XEdit_IsMultiLine")
	xEdit_IsPassword = xcgui.NewProc("XEdit_IsPassword")
	xEdit_IsAutoWrap = xcgui.NewProc("XEdit_IsAutoWrap")
	xEdit_IsEmpty = xcgui.NewProc("XEdit_IsEmpty")
	xEdit_IsInSelect = xcgui.NewProc("XEdit_IsInSelect")
	xEdit_GetRowCount = xcgui.NewProc("XEdit_GetRowCount")
	xEdit_GetData = xcgui.NewProc("XEdit_GetData")
	xEdit_AddData = xcgui.NewProc("XEdit_AddData")
	xEdit_FreeData = xcgui.NewProc("XEdit_FreeData")
	xEdit_SetDefaultText = xcgui.NewProc("XEdit_SetDefaultText")
	xEdit_SetDefaultTextColor = xcgui.NewProc("XEdit_SetDefaultTextColor")
	xEdit_SetPasswordCharacter = xcgui.NewProc("XEdit_SetPasswordCharacter")
	xEdit_SetTextAlign = xcgui.NewProc("XEdit_SetTextAlign")
	xEdit_SetTabSpace = xcgui.NewProc("XEdit_SetTabSpace")
	xEdit_SetText = xcgui.NewProc("XEdit_SetText")
	xEdit_SetTextInt = xcgui.NewProc("XEdit_SetTextInt")
	xEdit_GetText = xcgui.NewProc("XEdit_GetText")
	xEdit_GetTextRow = xcgui.NewProc("XEdit_GetTextRow")
	xEdit_GetLength = xcgui.NewProc("XEdit_GetLength")
	xEdit_GetLengthRow = xcgui.NewProc("XEdit_GetLengthRow")
	xEdit_GetAt = xcgui.NewProc("XEdit_GetAt")
	xEdit_InsertText = xcgui.NewProc("XEdit_InsertText")
	xEdit_AddTextUser = xcgui.NewProc("XEdit_AddTextUser")
	xEdit_AddText = xcgui.NewProc("XEdit_AddText")
	xEdit_AddTextEx = xcgui.NewProc("XEdit_AddTextEx")
	xEdit_AddObject = xcgui.NewProc("XEdit_AddObject")
	xEdit_AddByStyle = xcgui.NewProc("XEdit_AddByStyle")
	xEdit_AddStyle = xcgui.NewProc("XEdit_AddStyle")
	xEdit_AddStyleEx = xcgui.NewProc("XEdit_AddStyleEx")
	xEdit_GetStyleInfo = xcgui.NewProc("XEdit_GetStyleInfo")
	xEdit_SetCurStyle = xcgui.NewProc("XEdit_SetCurStyle")
	xEdit_SetCaretColor = xcgui.NewProc("XEdit_SetCaretColor")
	xEdit_SetCaretWidth = xcgui.NewProc("XEdit_SetCaretWidth")
	xEdit_SetSelectBkColor = xcgui.NewProc("XEdit_SetSelectBkColor")
	xEdit_SetRowHeight = xcgui.NewProc("XEdit_SetRowHeight")
	xEdit_SetRowHeightEx = xcgui.NewProc("XEdit_SetRowHeightEx")
	xEdit_SetCurPos = xcgui.NewProc("XEdit_SetCurPos")
	xEdit_GetCurPos = xcgui.NewProc("XEdit_GetCurPos")
	xEdit_GetCurRow = xcgui.NewProc("XEdit_GetCurRow")
	xEdit_GetCurCol = xcgui.NewProc("XEdit_GetCurCol")
	xEdit_GetPoint = xcgui.NewProc("XEdit_GetPoint")
	xEdit_AutoScroll = xcgui.NewProc("XEdit_AutoScroll")
	xEdit_AutoScrollEx = xcgui.NewProc("XEdit_AutoScrollEx")
	xEdit_PosToRowCol = xcgui.NewProc("XEdit_PosToRowCol")
	xEdit_RowColToPos = xcgui.NewProc("XEdit_RowColToPos")
	xEdit_SelectAll = xcgui.NewProc("XEdit_SelectAll")
	xEdit_CancelSelect = xcgui.NewProc("XEdit_CancelSelect")
	xEdit_DeleteSelect = xcgui.NewProc("XEdit_DeleteSelect")
	xEdit_SetSelect = xcgui.NewProc("XEdit_SetSelect")
	xEdit_GetSelectText = xcgui.NewProc("XEdit_GetSelectText")
	xEdit_GetSelectRange = xcgui.NewProc("XEdit_GetSelectRange")
	xEdit_GetVisibleRowRange = xcgui.NewProc("XEdit_GetVisibleRowRange")
	xEdit_Delete = xcgui.NewProc("XEdit_Delete")
	xEdit_DeleteRow = xcgui.NewProc("XEdit_DeleteRow")
	xEdit_ClipboardCut = xcgui.NewProc("XEdit_ClipboardCut")
	xEdit_ClipboardCopy = xcgui.NewProc("XEdit_ClipboardCopy")
	xEdit_ClipboardPaste = xcgui.NewProc("XEdit_ClipboardPaste")
	xEdit_Undo = xcgui.NewProc("XEdit_Undo")
	xEdit_Redo = xcgui.NewProc("XEdit_Redo")
	xEdit_AddChatBegin = xcgui.NewProc("XEdit_AddChatBegin")
	xEdit_AddChatEnd = xcgui.NewProc("XEdit_AddChatEnd")
	xEdit_SetChatIndentation = xcgui.NewProc("XEdit_SetChatIndentation")
	xEdit_SetRowSpace = xcgui.NewProc("XEdit_SetRowSpace")
	xEdit_SetCurPosEx = xcgui.NewProc("XEdit_SetCurPosEx")
	xEdit_GetCurPosEx = xcgui.NewProc("XEdit_GetCurPosEx")
	xEdit_MoveEnd = xcgui.NewProc("XEdit_MoveEnd")
	xEdit_SetBackFont = xcgui.NewProc("XEdit_SetBackFont")
	xEdit_ReleaseStyle = xcgui.NewProc("XEdit_ReleaseStyle")
	xEdit_ModifyStyle = xcgui.NewProc("XEdit_ModifyStyle")
	xEdit_SetSpaceSize = xcgui.NewProc("XEdit_SetSpaceSize")
	xEdit_SetCharSpaceSize = xcgui.NewProc("XEdit_SetCharSpaceSize")
	xEdit_GetSelectTextLength = xcgui.NewProc("XEdit_GetSelectTextLength")
	xEdit_SetSelectTextStyle = xcgui.NewProc("XEdit_SetSelectTextStyle")

	// LayoutEle.
	xLayout_Create = xcgui.NewProc("XLayout_Create")
	xLayout_CreateEx = xcgui.NewProc("XLayout_CreateEx")
	xLayout_IsEnableLayout = xcgui.NewProc("XLayout_IsEnableLayout")
	xLayout_EnableLayout = xcgui.NewProc("XLayout_EnableLayout")
	xLayout_ShowLayoutFrame = xcgui.NewProc("XLayout_ShowLayoutFrame")
	xLayout_GetWidthIn = xcgui.NewProc("XLayout_GetWidthIn")
	xLayout_GetHeightIn = xcgui.NewProc("XLayout_GetHeightIn")
	// LayoutFrame.
	xLayoutFrame_Create = xcgui.NewProc("XLayoutFrame_Create")
	xLayoutFrame_ShowLayoutFrame = xcgui.NewProc("XLayoutFrame_ShowLayoutFrame")
	// ProgressBar.
	xProgBar_Create = xcgui.NewProc("XProgBar_Create")
	xProgBar_SetRange = xcgui.NewProc("XProgBar_SetRange")
	xProgBar_GetRange = xcgui.NewProc("XProgBar_GetRange")
	xProgBar_SetImageLoad = xcgui.NewProc("XProgBar_SetImageLoad")
	xProgBar_SetPos = xcgui.NewProc("XProgBar_SetPos")
	xProgBar_GetPos = xcgui.NewProc("XProgBar_GetPos")
	xProgBar_EnableHorizon = xcgui.NewProc("XProgBar_EnableHorizon")
	xProgBar_EnableShowText = xcgui.NewProc("XProgBar_EnableShowText")
	xProgBar_EnableStretch = xcgui.NewProc("XProgBar_EnableStretch")

	// TextLink.
	xTextLink_Create = xcgui.NewProc("XTextLink_Create")
	xTextLink_EnableUnderlineLeave = xcgui.NewProc("XTextLink_EnableUnderlineLeave")
	xTextLink_EnableUnderlineStay = xcgui.NewProc("XTextLink_EnableUnderlineStay")
	xTextLink_SetTextColorStay = xcgui.NewProc("XTextLink_SetTextColorStay")
	xTextLink_SetUnderlineColorLeave = xcgui.NewProc("XTextLink_SetUnderlineColorLeave")
	xTextLink_SetUnderlineColorStay = xcgui.NewProc("XTextLink_SetUnderlineColorStay")

	// Shape.
	xShape_RemoveShape = xcgui.NewProc("XShape_RemoveShape")
	xShape_GetZOrder = xcgui.NewProc("XShape_GetZOrder")
	xShape_Redraw = xcgui.NewProc("XShape_Redraw")
	xShape_GetWidth = xcgui.NewProc("XShape_GetWidth")
	xShape_GetHeight = xcgui.NewProc("XShape_GetHeight")
	xShape_SetPosition = xcgui.NewProc("XShape_SetPosition")
	xShape_GetRect = xcgui.NewProc("XShape_GetRect")
	xShape_SetRect = xcgui.NewProc("XShape_SetRect")
	xShape_SetRectLogic = xcgui.NewProc("XShape_SetRectLogic")
	xShape_GetRectLogic = xcgui.NewProc("XShape_GetRectLogic")
	xShape_GetWndClientRect = xcgui.NewProc("XShape_GetWndClientRect")
	xShape_GetContentSize = xcgui.NewProc("XShape_GetContentSize")
	xShape_ShowLayout = xcgui.NewProc("XShape_ShowLayout")
	xShape_AdjustLayout = xcgui.NewProc("XShape_AdjustLayout")
	xShape_Destroy = xcgui.NewProc("XShape_Destroy")
	xShape_GetPosition = xcgui.NewProc("XShape_GetPosition")
	xShape_SetSize = xcgui.NewProc("XShape_SetSize")
	xShape_GetSize = xcgui.NewProc("XShape_GetSize")
	xShape_SetAlpha = xcgui.NewProc("XShape_SetAlpha")
	xShape_GetAlpha = xcgui.NewProc("XShape_GetAlpha")

	// ShapeText.
	xShapeText_Create = xcgui.NewProc("XShapeText_Create")
	xShapeText_SetText = xcgui.NewProc("XShapeText_SetText")
	xShapeText_GetText = xcgui.NewProc("XShapeText_GetText")
	xShapeText_GetTextLength = xcgui.NewProc("XShapeText_GetTextLength")
	xShapeText_SetFont = xcgui.NewProc("XShapeText_SetFont")
	xShapeText_GetFont = xcgui.NewProc("XShapeText_GetFont")
	xShapeText_SetTextColor = xcgui.NewProc("XShapeText_SetTextColor")
	xShapeText_GetTextColor = xcgui.NewProc("XShapeText_GetTextColor")
	xShapeText_SetTextAlign = xcgui.NewProc("XShapeText_SetTextAlign")
	xShapeText_SetOffset = xcgui.NewProc("XShapeText_SetOffset")

	// ShapePicture.
	xShapePic_Create = xcgui.NewProc("XShapePic_Create")
	xShapePic_SetImage = xcgui.NewProc("XShapePic_SetImage")
	xShapePic_GetImage = xcgui.NewProc("XShapePic_GetImage")

	// ShapeGif.
	xShapeGif_Create = xcgui.NewProc("XShapeGif_Create")
	xShapeGif_SetImage = xcgui.NewProc("XShapeGif_SetImage")
	xShapeGif_GetImage = xcgui.NewProc("XShapeGif_GetImage")

	// ShapeRect.
	xShapeRect_Create = xcgui.NewProc("XShapeRect_Create")
	xShapeRect_SetBorderColor = xcgui.NewProc("XShapeRect_SetBorderColor")
	xShapeRect_SetFillColor = xcgui.NewProc("XShapeRect_SetFillColor")
	xShapeRect_SetRoundAngle = xcgui.NewProc("XShapeRect_SetRoundAngle")
	xShapeRect_GetRoundAngle = xcgui.NewProc("XShapeRect_GetRoundAngle")
	xShapeRect_EnableBorder = xcgui.NewProc("XShapeRect_EnableBorder")
	xShapeRect_EnableFill = xcgui.NewProc("XShapeRect_EnableFill")
	xShapeRect_EnableRoundAngle = xcgui.NewProc("XShapeRect_EnableRoundAngle")

	// ShapeEllipse.
	xShapeEllipse_Create = xcgui.NewProc("XShapeEllipse_Create")
	xShapeEllipse_SetBorderColor = xcgui.NewProc("XShapeEllipse_SetBorderColor")
	xShapeEllipse_SetFillColor = xcgui.NewProc("XShapeEllipse_SetFillColor")
	xShapeEllipse_EnableBorder = xcgui.NewProc("XShapeEllipse_EnableBorder")
	xShapeEllipse_EnableFill = xcgui.NewProc("XShapeEllipse_EnableFill")

	// ShapeGroupBox.
	xShapeGroupBox_Create = xcgui.NewProc("XShapeGroupBox_Create")
	xShapeGroupBox_SetBorderColor = xcgui.NewProc("XShapeGroupBox_SetBorderColor")
	xShapeGroupBox_SetTextColor = xcgui.NewProc("XShapeGroupBox_SetTextColor")
	xShapeGroupBox_SetFontX = xcgui.NewProc("XShapeGroupBox_SetFontX")
	xShapeGroupBox_SetTextOffset = xcgui.NewProc("XShapeGroupBox_SetTextOffset")
	xShapeGroupBox_SetRoundAngle = xcgui.NewProc("XShapeGroupBox_SetRoundAngle")
	xShapeGroupBox_SetText = xcgui.NewProc("XShapeGroupBox_SetText")
	xShapeGroupBox_GetTextOffset = xcgui.NewProc("XShapeGroupBox_GetTextOffset")
	xShapeGroupBox_GetRoundAngle = xcgui.NewProc("XShapeGroupBox_GetRoundAngle")
	xShapeGroupBox_EnableRoundAngle = xcgui.NewProc("XShapeGroupBox_EnableRoundAngle")

	// ShapeLine.
	xShapeLine_Create = xcgui.NewProc("XShapeLine_Create")
	xShapeLine_SetPosition = xcgui.NewProc("XShapeLine_SetPosition")
	xShapeLine_SetColor = xcgui.NewProc("XShapeLine_SetColor")

	// Table.
	xTable_Create = xcgui.NewProc("XTable_Create")
	xTable_Reset = xcgui.NewProc("XTable_Reset")
	xTable_ComboRow = xcgui.NewProc("XTable_ComboRow")
	xTable_ComboCol = xcgui.NewProc("XTable_ComboCol")
	xTable_SetColWidth = xcgui.NewProc("XTable_SetColWidth")
	xTable_SetRowHeight = xcgui.NewProc("XTable_SetRowHeight")
	xTable_SetBorderColor = xcgui.NewProc("XTable_SetBorderColor")
	xTable_SetTextColor = xcgui.NewProc("XTable_SetTextColor")
	xTable_SetFont = xcgui.NewProc("XTable_SetFont")
	xTable_SetItemPadding = xcgui.NewProc("XTable_SetItemPadding")
	xTable_SetItemText = xcgui.NewProc("XTable_SetItemText")
	xTable_SetItemFont = xcgui.NewProc("XTable_SetItemFont")
	xTable_SetItemTextAlign = xcgui.NewProc("XTable_SetItemTextAlign")
	xTable_SetItemTextColor = xcgui.NewProc("XTable_SetItemTextColor")
	xTable_SetItemBkColor = xcgui.NewProc("XTable_SetItemBkColor")
	xTable_SetItemLine = xcgui.NewProc("XTable_SetItemLine")
	xTable_SetItemFlag = xcgui.NewProc("XTable_SetItemFlag")
	xTable_GetItemRect = xcgui.NewProc("XTable_GetItemRect")

	// BkManager.
	xBkM_Create = xcgui.NewProc("XBkM_Create")
	xBkM_Destroy = xcgui.NewProc("XBkM_Destroy")
	xBkM_SetBkInfo = xcgui.NewProc("XBkM_SetBkInfo")
	xBkM_AddInfo = xcgui.NewProc("XBkM_AddInfo")
	xBkM_AddBorder = xcgui.NewProc("XBkM_AddBorder")
	xBkM_AddFill = xcgui.NewProc("XBkM_AddFill")
	xBkM_AddImage = xcgui.NewProc("XBkM_AddImage")
	xBkM_GetCount = xcgui.NewProc("XBkM_GetCount")
	xBkM_Clear = xcgui.NewProc("XBkM_Clear")
	xBkM_Draw = xcgui.NewProc("XBkM_Draw")
	xBkM_DrawEx = xcgui.NewProc("XBkM_DrawEx")
	xBkM_EnableAutoDestroy = xcgui.NewProc("XBkM_EnableAutoDestroy")
	xBkM_AddRef = xcgui.NewProc("XBkM_AddRef")
	xBkM_Release = xcgui.NewProc("XBkM_Release")
	xBkM_GetRefCount = xcgui.NewProc("XBkM_GetRefCount")
	xBkM_SetInfo = xcgui.NewProc("XBkM_SetInfo")
	xBkM_GetStateTextColor = xcgui.NewProc("XBkM_GetStateTextColor")
	xBkM_GetObject = xcgui.NewProc("XBkM_GetObject")

	// Draw.
	xDraw_Create = xcgui.NewProc("XDraw_Create")
	xDraw_Destroy = xcgui.NewProc("XDraw_Destroy")
	xDraw_SetOffset = xcgui.NewProc("XDraw_SetOffset")
	xDraw_GetOffset = xcgui.NewProc("XDraw_GetOffset")
	xDraw_GDI_RestoreGDIOBJ = xcgui.NewProc("XDraw_GDI_RestoreGDIOBJ")
	xDraw_GetHDC = xcgui.NewProc("XDraw_GetHDC")
	xDraw_SetBrushColor = xcgui.NewProc("XDraw_SetBrushColor")
	xDraw_SetTextVertical = xcgui.NewProc("XDraw_SetTextVertical")
	xDraw_SetTextAlign = xcgui.NewProc("XDraw_SetTextAlign")
	xDraw_SetFont = xcgui.NewProc("XDraw_SetFont")
	xDraw_SetLineWidth = xcgui.NewProc("XDraw_SetLineWidth")
	xDraw_SetLineWidthF = xcgui.NewProc("XDraw_SetLineWidthF")
	xDraw_GDI_SetBkMode = xcgui.NewProc("XDraw_GDI_SetBkMode")
	xDraw_SetClipRect = xcgui.NewProc("XDraw_SetClipRect")
	xDraw_ClearClip = xcgui.NewProc("XDraw_ClearClip")
	xDraw_EnableSmoothingMode = xcgui.NewProc("XDraw_EnableSmoothingMode")
	xDraw_EnableWndTransparent = xcgui.NewProc("XDraw_EnableWndTransparent")
	xDraw_GDI_CreateSolidBrush = xcgui.NewProc("XDraw_GDI_CreateSolidBrush")
	xDraw_GDI_CreatePen = xcgui.NewProc("XDraw_GDI_CreatePen")
	xDraw_GDI_CreateRectRgn = xcgui.NewProc("XDraw_GDI_CreateRectRgn")
	xDraw_GDI_CreateRoundRectRgn = xcgui.NewProc("XDraw_GDI_CreateRoundRectRgn")
	xDraw_GDI_CreatePolygonRgn = xcgui.NewProc("XDraw_GDI_CreatePolygonRgn")
	xDraw_GDI_SelectClipRgn = xcgui.NewProc("XDraw_GDI_SelectClipRgn")
	xDraw_FillRect = xcgui.NewProc("XDraw_FillRect")
	xDraw_FillRectF = xcgui.NewProc("XDraw_FillRectF")
	xDraw_FillRectColor = xcgui.NewProc("XDraw_FillRectColor")
	xDraw_FillRectColorF = xcgui.NewProc("XDraw_FillRectColorF")
	xDraw_GDI_FillRgn = xcgui.NewProc("XDraw_GDI_FillRgn")
	xDraw_FillEllipse = xcgui.NewProc("XDraw_FillEllipse")
	xDraw_FillEllipseF = xcgui.NewProc("XDraw_FillEllipseF")
	xDraw_DrawEllipse = xcgui.NewProc("XDraw_DrawEllipse")
	xDraw_FillRoundRect = xcgui.NewProc("XDraw_FillRoundRect")
	xDraw_FillRoundRectF = xcgui.NewProc("XDraw_FillRoundRectF")
	xDraw_DrawRoundRect = xcgui.NewProc("XDraw_DrawRoundRect")
	xDraw_DrawRoundRectF = xcgui.NewProc("XDraw_DrawRoundRectF")
	xDraw_FillRoundRectEx = xcgui.NewProc("XDraw_FillRoundRectEx")
	xDraw_FillRoundRectExF = xcgui.NewProc("XDraw_FillRoundRectExF")
	xDraw_DrawRoundRectEx = xcgui.NewProc("XDraw_DrawRoundRectEx")
	xDraw_DrawRoundRectExF = xcgui.NewProc("XDraw_DrawRoundRectExF")
	xDraw_GDI_Rectangle = xcgui.NewProc("XDraw_GDI_Rectangle")
	xDraw_GradientFill2 = xcgui.NewProc("XDraw_GradientFill2")
	xDraw_GradientFill2F = xcgui.NewProc("XDraw_GradientFill2F")
	xDraw_GradientFill4 = xcgui.NewProc("XDraw_GradientFill4")
	xDraw_GradientFill4F = xcgui.NewProc("XDraw_GradientFill4F")
	xDraw_GDI_FrameRgn = xcgui.NewProc("XDraw_GDI_FrameRgn")
	xDraw_DrawLine = xcgui.NewProc("XDraw_DrawLine")
	xDraw_DrawLineF = xcgui.NewProc("XDraw_DrawLineF")
	xDraw_DrawCurve = xcgui.NewProc("XDraw_DrawCurve")
	xDraw_DrawCurveF = xcgui.NewProc("XDraw_DrawCurveF")
	xDraw_FocusRect = xcgui.NewProc("XDraw_FocusRect")
	xDraw_FocusRectF = xcgui.NewProc("XDraw_FocusRectF")
	xDraw_GDI_MoveToEx = xcgui.NewProc("XDraw_GDI_MoveToEx")
	xDraw_GDI_LineTo = xcgui.NewProc("XDraw_GDI_LineTo")
	xDraw_GDI_Polyline = xcgui.NewProc("XDraw_GDI_Polyline")
	xDraw_Dottedline = xcgui.NewProc("XDraw_Dottedline")
	xDraw_DottedlineF = xcgui.NewProc("XDraw_DottedlineF")
	xDraw_GDI_SetPixel = xcgui.NewProc("XDraw_GDI_SetPixel")
	xDraw_GDI_DrawIconEx = xcgui.NewProc("XDraw_GDI_DrawIconEx")
	xDraw_GDI_BitBlt = xcgui.NewProc("XDraw_GDI_BitBlt")
	xDraw_GDI_BitBlt2 = xcgui.NewProc("XDraw_GDI_BitBlt2")
	xDraw_GDI_AlphaBlend = xcgui.NewProc("XDraw_GDI_AlphaBlend")
	xDraw_DrawPolygon = xcgui.NewProc("XDraw_DrawPolygon")
	xDraw_DrawPolygonF = xcgui.NewProc("XDraw_DrawPolygonF")
	xDraw_FillPolygon = xcgui.NewProc("XDraw_FillPolygon")
	xDraw_FillPolygonF = xcgui.NewProc("XDraw_FillPolygonF")
	xDraw_Image = xcgui.NewProc("XDraw_Image")
	xDraw_ImageF = xcgui.NewProc("XDraw_ImageF")
	xDraw_ImageEx = xcgui.NewProc("XDraw_ImageEx")
	xDraw_ImageExF = xcgui.NewProc("XDraw_ImageExF")
	xDraw_ImageAdaptive = xcgui.NewProc("XDraw_ImageAdaptive")
	xDraw_ImageAdaptiveF = xcgui.NewProc("XDraw_ImageAdaptiveF")
	xDraw_ImageSuper = xcgui.NewProc("XDraw_ImageSuper")
	xDraw_ImageSuperF = xcgui.NewProc("XDraw_ImageSuperF")
	xDraw_ImageSuperEx = xcgui.NewProc("XDraw_ImageSuperEx")
	xDraw_ImageSuperExF = xcgui.NewProc("XDraw_ImageSuperExF")
	xDraw_ImageSuperMask = xcgui.NewProc("XDraw_ImageSuperMask")
	xDraw_ImageMask = xcgui.NewProc("XDraw_ImageMask")
	xDraw_DrawText = xcgui.NewProc("XDraw_DrawText")
	xDraw_DrawTextF = xcgui.NewProc("XDraw_DrawTextF")
	xDraw_DrawTextUnderline = xcgui.NewProc("XDraw_DrawTextUnderline")
	xDraw_DrawTextUnderlineF = xcgui.NewProc("XDraw_DrawTextUnderlineF")
	xDraw_TextOut = xcgui.NewProc("XDraw_TextOut")
	xDraw_TextOutF = xcgui.NewProc("XDraw_TextOutF")
	xDraw_TextOutEx = xcgui.NewProc("XDraw_TextOutEx")
	xDraw_TextOutExF = xcgui.NewProc("XDraw_TextOutExF")
	xDraw_TextOutA = xcgui.NewProc("XDraw_TextOutA")
	xDraw_TextOutAF = xcgui.NewProc("XDraw_TextOutAF")
	xDraw_DrawArc = xcgui.NewProc("XDraw_DrawArc")
	xDraw_DrawArcF = xcgui.NewProc("XDraw_DrawArcF")
	xDraw_DrawRect = xcgui.NewProc("XDraw_DrawRect")
	xDraw_DrawRectF = xcgui.NewProc("XDraw_DrawRectF")
	xDraw_GDI_Ellipse = xcgui.NewProc("XDraw_GDI_Ellipse")
	xDraw_GetD2dRenderTarget = xcgui.NewProc("XDraw_GetD2dRenderTarget")
	xDraw_ImageTile = xcgui.NewProc("XDraw_ImageTile")
	xDraw_ImageTileF = xcgui.NewProc("XDraw_ImageTileF")
	xDraw_SetD2dTextRenderingMode = xcgui.NewProc("XDraw_SetD2dTextRenderingMode")
	xDraw_CreateGDI = xcgui.NewProc("XDraw_CreateGDI")
	xDraw_SetTextRenderingHint = xcgui.NewProc("XDraw_SetTextRenderingHint")
	xDraw_DrawSvgSrc = xcgui.NewProc("XDraw_DrawSvgSrc")
	xDraw_DrawSvg = xcgui.NewProc("XDraw_DrawSvg")
	xDraw_DrawSvgEx = xcgui.NewProc("XDraw_DrawSvgEx")
	xDraw_DrawSvgSize = xcgui.NewProc("XDraw_DrawSvgSize")
	xDraw_D2D_Clear = xcgui.NewProc("XDraw_D2D_Clear")
	xDraw_ImageMaskRect = xcgui.NewProc("XDraw_ImageMaskRect")
	xDraw_ImageMaskEllipse = xcgui.NewProc("XDraw_ImageMaskEllipse")

	// Ease.
	xEase_Linear = xcgui.NewProc("XEase_Linear")
	xEase_Quad = xcgui.NewProc("XEase_Quad")
	xEase_Cubic = xcgui.NewProc("XEase_Cubic")
	xEase_Quart = xcgui.NewProc("XEase_Quart")
	xEase_Quint = xcgui.NewProc("XEase_Quint")
	xEase_Sine = xcgui.NewProc("XEase_Sine")
	xEase_Expo = xcgui.NewProc("XEase_Expo")
	xEase_Circ = xcgui.NewProc("XEase_Circ")
	xEase_Elastic = xcgui.NewProc("XEase_Elastic")
	xEase_Back = xcgui.NewProc("XEase_Back")
	xEase_Bounce = xcgui.NewProc("XEase_Bounce")
	xEase_Ex = xcgui.NewProc("XEase_Ex")

	// FontX.
	xFont_Create = xcgui.NewProc("XFont_Create")
	xFont_CreateEx = xcgui.NewProc("XFont_CreateEx")
	xFont_CreateFromHFONT = xcgui.NewProc("XFont_CreateFromHFONT")
	xFont_CreateFromFont = xcgui.NewProc("XFont_CreateFromFont")
	xFont_CreateFromFile = xcgui.NewProc("XFont_CreateFromFile")
	xFont_CreateLOGFONTW = xcgui.NewProc("XFont_CreateLOGFONTW")
	xFont_EnableAutoDestroy = xcgui.NewProc("XFont_EnableAutoDestroy")
	xFont_GetFont = xcgui.NewProc("XFont_GetFont")
	xFont_GetFontInfo = xcgui.NewProc("XFont_GetFontInfo")
	xFont_GetLOGFONTW = xcgui.NewProc("XFont_GetLOGFONTW")
	xFont_Destroy = xcgui.NewProc("XFont_Destroy")
	xFont_AddRef = xcgui.NewProc("XFont_AddRef")
	xFont_GetRefCount = xcgui.NewProc("XFont_GetRefCount")
	xFont_Release = xcgui.NewProc("XFont_Release")
	xFont_CreateFromMem = xcgui.NewProc("XFont_CreateFromMem")
	xFont_CreateFromRes = xcgui.NewProc("XFont_CreateFromRes")
	xFont_CreateFromZip = xcgui.NewProc("XFont_CreateFromZip")
	xFont_CreateFromZipMem = xcgui.NewProc("XFont_CreateFromZipMem")

	// Image.
	xImage_LoadSrc = xcgui.NewProc("XImage_LoadSrc")
	xImage_LoadFile = xcgui.NewProc("XImage_LoadFile")
	xImage_LoadFileAdaptive = xcgui.NewProc("XImage_LoadFileAdaptive")
	xImage_LoadFileRect = xcgui.NewProc("XImage_LoadFileRect")
	xImage_LoadResAdaptive = xcgui.NewProc("XImage_LoadResAdaptive")
	xImage_LoadRes = xcgui.NewProc("XImage_LoadRes")
	xImage_LoadZip = xcgui.NewProc("XImage_LoadZip")
	xImage_LoadZipAdaptive = xcgui.NewProc("XImage_LoadZipAdaptive")
	xImage_LoadZipRect = xcgui.NewProc("XImage_LoadZipRect")
	xImage_LoadZipMem = xcgui.NewProc("XImage_LoadZipMem")
	xImage_LoadMemory = xcgui.NewProc("XImage_LoadMemory")
	xImage_LoadMemoryRect = xcgui.NewProc("XImage_LoadMemoryRect")
	xImage_LoadMemoryAdaptive = xcgui.NewProc("XImage_LoadMemoryAdaptive")
	xImage_LoadFromImage = xcgui.NewProc("XImage_LoadFromImage")
	xImage_LoadFromExtractIcon = xcgui.NewProc("XImage_LoadFromExtractIcon")
	xImage_LoadFromHICON = xcgui.NewProc("XImage_LoadFromHICON")
	xImage_LoadFromHBITMAP = xcgui.NewProc("XImage_LoadFromHBITMAP")
	xImage_IsStretch = xcgui.NewProc("XImage_IsStretch")
	xImage_IsAdaptive = xcgui.NewProc("XImage_IsAdaptive")
	xImage_IsTile = xcgui.NewProc("XImage_IsTile")
	xImage_SetDrawType = xcgui.NewProc("XImage_SetDrawType")
	xImage_SetDrawTypeAdaptive = xcgui.NewProc("XImage_SetDrawTypeAdaptive")
	xImage_SetTranColor = xcgui.NewProc("XImage_SetTranColor")
	xImage_SetTranColorEx = xcgui.NewProc("XImage_SetTranColorEx")
	xImage_SetRotateAngle = xcgui.NewProc("XImage_SetRotateAngle")
	xImage_SetSplitEqual = xcgui.NewProc("XImage_SetSplitEqual")
	xImage_EnableTranColor = xcgui.NewProc("XImage_EnableTranColor")
	xImage_EnableAutoDestroy = xcgui.NewProc("XImage_EnableAutoDestroy")
	xImage_EnableCenter = xcgui.NewProc("XImage_EnableCenter")
	xImage_IsCenter = xcgui.NewProc("XImage_IsCenter")
	xImage_GetDrawType = xcgui.NewProc("XImage_GetDrawType")
	xImage_GetWidth = xcgui.NewProc("XImage_GetWidth")
	xImage_GetHeight = xcgui.NewProc("XImage_GetHeight")
	xImage_GetImageSrc = xcgui.NewProc("XImage_GetImageSrc")
	xImage_AddRef = xcgui.NewProc("XImage_AddRef")
	xImage_Release = xcgui.NewProc("XImage_Release")
	xImage_GetRefCount = xcgui.NewProc("XImage_GetRefCount")
	xImage_Destroy = xcgui.NewProc("XImage_Destroy")
	xImage_LoadSvg = xcgui.NewProc("XImage_LoadSvg")
	xImage_LoadSvgFile = xcgui.NewProc("XImage_LoadSvgFile")
	xImage_LoadSvgString = xcgui.NewProc("XImage_LoadSvgString")
	xImage_GetSvg = xcgui.NewProc("XImage_GetSvg")
	xImage_LoadSvgStringW = xcgui.NewProc("XImage_LoadSvgStringW")
	xImage_LoadSvgStringUtf8 = xcgui.NewProc("XImage_LoadSvgStringUtf8")
	xImage_SetScaleSize = xcgui.NewProc("XImage_SetScaleSize")

	// Svg.
	xSvg_LoadFile = xcgui.NewProc("XSvg_LoadFile")
	xSvg_LoadString = xcgui.NewProc("XSvg_LoadString")
	xSvg_LoadZip = xcgui.NewProc("XSvg_LoadZip")
	xSvg_LoadRes = xcgui.NewProc("XSvg_LoadRes")
	xSvg_SetSize = xcgui.NewProc("XSvg_SetSize")
	xSvg_GetSize = xcgui.NewProc("XSvg_GetSize")
	xSvg_GetWidth = xcgui.NewProc("XSvg_GetWidth")
	xSvg_GetHeight = xcgui.NewProc("XSvg_GetHeight")
	xSvg_SetPosition = xcgui.NewProc("XSvg_SetPosition")
	xSvg_GetPosition = xcgui.NewProc("XSvg_GetPosition")
	xSvg_SetPositionF = xcgui.NewProc("XSvg_SetPositionF")
	xSvg_GetPositionF = xcgui.NewProc("XSvg_GetPositionF")
	xSvg_GetViewBox = xcgui.NewProc("XSvg_GetViewBox")
	xSvg_EnableAutoDestroy = xcgui.NewProc("XSvg_EnableAutoDestroy")
	xSvg_AddRef = xcgui.NewProc("XSvg_AddRef")
	xSvg_Release = xcgui.NewProc("XSvg_Release")
	xSvg_GetRefCount = xcgui.NewProc("XSvg_GetRefCount")
	xSvg_Destroy = xcgui.NewProc("XSvg_Destroy")
	xSvg_SetAlpha = xcgui.NewProc("XSvg_SetAlpha")
	xSvg_GetAlpha = xcgui.NewProc("XSvg_GetAlpha")
	xSvg_SetUserFillColor = xcgui.NewProc("XSvg_SetUserFillColor")
	xSvg_SetUserStrokeColor = xcgui.NewProc("XSvg_SetUserStrokeColor")
	xSvg_GetUserFillColor = xcgui.NewProc("XSvg_GetUserFillColor")
	xSvg_GetUserStrokeColor = xcgui.NewProc("XSvg_GetUserStrokeColor")
	xSvg_SetRotateAngle = xcgui.NewProc("XSvg_SetRotateAngle")
	xSvg_GetRotateAngle = xcgui.NewProc("XSvg_GetRotateAngle")
	xSvg_SetRotate = xcgui.NewProc("XSvg_SetRotate")
	xSvg_GetRotate = xcgui.NewProc("XSvg_GetRotate")
	xSvg_Show = xcgui.NewProc("XSvg_Show")
	xSvg_LoadStringW = xcgui.NewProc("XSvg_LoadStringW")
	xSvg_LoadStringUtf8 = xcgui.NewProc("XSvg_LoadStringUtf8")
	xSvg_LoadZipMem = xcgui.NewProc("XSvg_LoadZipMem")

	// ListItemTemplate.
	xTemp_Load = xcgui.NewProc("XTemp_Load")
	xTemp_LoadZip = xcgui.NewProc("XTemp_LoadZip")
	xTemp_LoadZipMem = xcgui.NewProc("XTemp_LoadZipMem")
	xTemp_LoadEx = xcgui.NewProc("XTemp_LoadEx")
	xTemp_LoadZipEx = xcgui.NewProc("XTemp_LoadZipEx")
	xTemp_LoadZipMemEx = xcgui.NewProc("XTemp_LoadZipMemEx")
	xTemp_LoadFromString = xcgui.NewProc("XTemp_LoadFromString")
	xTemp_LoadFromStringEx = xcgui.NewProc("XTemp_LoadFromStringEx")
	xTemp_GetType = xcgui.NewProc("XTemp_GetType")
	xTemp_Destroy = xcgui.NewProc("XTemp_Destroy")
	xTemp_Create = xcgui.NewProc("XTemp_Create")
	xTemp_AddNodeRoot = xcgui.NewProc("XTemp_AddNodeRoot")
	xTemp_AddNode = xcgui.NewProc("XTemp_AddNode")
	xTemp_CreateNode = xcgui.NewProc("XTemp_CreateNode")
	xTemp_SetNodeAttribute = xcgui.NewProc("XTemp_SetNodeAttribute")
	xTemp_SetNodeAttributeEx = xcgui.NewProc("XTemp_SetNodeAttributeEx")
	xTemp_List_GetNode = xcgui.NewProc("XTemp_List_GetNode")
	xTemp_GetNode = xcgui.NewProc("XTemp_GetNode")
	xTemp_CloneNode = xcgui.NewProc("XTemp_CloneNode")
	xTemp_Clone = xcgui.NewProc("XTemp_Clone")

	// Resource.
	xRes_EnableDelayLoad = xcgui.NewProc("XRes_EnableDelayLoad")
	xRes_SetLoadFileCallback = xcgui.NewProc("XRes_SetLoadFileCallback")
	xRes_GetIDValue = xcgui.NewProc("XRes_GetIDValue")
	xRes_GetImage = xcgui.NewProc("XRes_GetImage")
	xRes_GetImageEx = xcgui.NewProc("XRes_GetImageEx")
	xRes_GetColor = xcgui.NewProc("XRes_GetColor")
	xRes_GetFont = xcgui.NewProc("XRes_GetFont")
	xRes_GetBkM = xcgui.NewProc("XRes_GetBkM")

	// ListBox.
	xListBox_Create = xcgui.NewProc("XListBox_Create")
	xListBox_EnableFixedRowHeight = xcgui.NewProc("XListBox_EnableFixedRowHeight")
	xListBox_EnablemTemplateReuse = xcgui.NewProc("XListBox_EnablemTemplateReuse")
	xListBox_EnableVirtualTable = xcgui.NewProc("XListBox_EnableVirtualTable")
	xListBox_SetVirtualRowCount = xcgui.NewProc("XListBox_SetVirtualRowCount")
	xListBox_SetDrawItemBkFlags = xcgui.NewProc("XListBox_SetDrawItemBkFlags")
	xListBox_SetItemData = xcgui.NewProc("XListBox_SetItemData")
	xListBox_GetItemData = xcgui.NewProc("XListBox_GetItemData")
	xListBox_SetItemInfo = xcgui.NewProc("XListBox_SetItemInfo")
	xListBox_GetItemInfo = xcgui.NewProc("XListBox_GetItemInfo")
	xListBox_SetSelectItem = xcgui.NewProc("XListBox_SetSelectItem")
	xListBox_GetSelectItem = xcgui.NewProc("XListBox_GetSelectItem")
	xListBox_AddSelectItem = xcgui.NewProc("XListBox_AddSelectItem")
	xListBox_CancelSelectItem = xcgui.NewProc("XListBox_CancelSelectItem")
	xListBox_CancelSelectAll = xcgui.NewProc("XListBox_CancelSelectAll")
	xListBox_GetSelectAll = xcgui.NewProc("XListBox_GetSelectAll")
	xListBox_GetSelectCount = xcgui.NewProc("XListBox_GetSelectCount")
	xListBox_GetItemMouseStay = xcgui.NewProc("XListBox_GetItemMouseStay")
	xListBox_SelectAll = xcgui.NewProc("XListBox_SelectAll")
	xListBox_VisibleItem = xcgui.NewProc("XListBox_VisibleItem")
	xListBox_GetVisibleRowRange = xcgui.NewProc("XListBox_GetVisibleRowRange")
	xListBox_SetItemHeightDefault = xcgui.NewProc("XListBox_SetItemHeightDefault")
	xListBox_GetItemHeightDefault = xcgui.NewProc("XListBox_GetItemHeightDefault")
	xListBox_GetItemIndexFromHXCGUI = xcgui.NewProc("XListBox_GetItemIndexFromHXCGUI")
	xListBox_SetRowSpace = xcgui.NewProc("XListBox_SetRowSpace")
	xListBox_GetRowSpace = xcgui.NewProc("XListBox_GetRowSpace")
	xListBox_HitTest = xcgui.NewProc("XListBox_HitTest")
	xListBox_HitTestOffset = xcgui.NewProc("XListBox_HitTestOffset")
	xListBox_SetItemTemplateXML = xcgui.NewProc("XListBox_SetItemTemplateXML")
	xListBox_SetItemTemplate = xcgui.NewProc("XListBox_SetItemTemplate")
	xListBox_SetItemTemplateXMLFromString = xcgui.NewProc("XListBox_SetItemTemplateXMLFromString")
	xListBox_GetTemplateObject = xcgui.NewProc("XListBox_GetTemplateObject")
	xListBox_EnableMultiSel = xcgui.NewProc("XListBox_EnableMultiSel")
	xListBox_CreateAdapter = xcgui.NewProc("XListBox_CreateAdapter")
	xListBox_BindAdapter = xcgui.NewProc("XListBox_BindAdapter")
	xListBox_GetAdapter = xcgui.NewProc("XListBox_GetAdapter")
	xListBox_Sort = xcgui.NewProc("XListBox_Sort")
	xListBox_RefreshData = xcgui.NewProc("XListBox_RefreshData")
	xListBox_RefreshItem = xcgui.NewProc("XListBox_RefreshItem")
	xListBox_AddItemText = xcgui.NewProc("XListBox_AddItemText")
	xListBox_AddItemTextEx = xcgui.NewProc("XListBox_AddItemTextEx")
	xListBox_AddItemImage = xcgui.NewProc("XListBox_AddItemImage")
	xListBox_AddItemImageEx = xcgui.NewProc("XListBox_AddItemImageEx")
	xListBox_InsertItemText = xcgui.NewProc("XListBox_InsertItemText")
	xListBox_InsertItemTextEx = xcgui.NewProc("XListBox_InsertItemTextEx")
	xListBox_InsertItemImage = xcgui.NewProc("XListBox_InsertItemImage")
	xListBox_InsertItemImageEx = xcgui.NewProc("XListBox_InsertItemImageEx")
	xListBox_SetItemText = xcgui.NewProc("XListBox_SetItemText")
	xListBox_SetItemTextEx = xcgui.NewProc("XListBox_SetItemTextEx")
	xListBox_SetItemImage = xcgui.NewProc("XListBox_SetItemImage")
	xListBox_SetItemImageEx = xcgui.NewProc("XListBox_SetItemImageEx")
	xListBox_SetItemInt = xcgui.NewProc("XListBox_SetItemInt")
	xListBox_SetItemIntEx = xcgui.NewProc("XListBox_SetItemIntEx")
	xListBox_SetItemFloat = xcgui.NewProc("XListBox_SetItemFloat")
	xListBox_SetItemFloatEx = xcgui.NewProc("XListBox_SetItemFloatEx")
	xListBox_GetItemText = xcgui.NewProc("XListBox_GetItemText")
	xListBox_GetItemTextEx = xcgui.NewProc("XListBox_GetItemTextEx")
	xListBox_GetItemImage = xcgui.NewProc("XListBox_GetItemImage")
	xListBox_GetItemImageEx = xcgui.NewProc("XListBox_GetItemImageEx")
	xListBox_GetItemInt = xcgui.NewProc("XListBox_GetItemInt")
	xListBox_GetItemIntEx = xcgui.NewProc("XListBox_GetItemIntEx")
	xListBox_GetItemFloat = xcgui.NewProc("XListBox_GetItemFloat")
	xListBox_GetItemFloatEx = xcgui.NewProc("XListBox_GetItemFloatEx")
	xListBox_DeleteItem = xcgui.NewProc("XListBox_DeleteItem")
	xListBox_DeleteItemEx = xcgui.NewProc("XListBox_DeleteItemEx")
	xListBox_DeleteItemAll = xcgui.NewProc("XListBox_DeleteItemAll")
	xListBox_DeleteColumnAll = xcgui.NewProc("XListBox_DeleteColumnAll")
	xListBox_GetCount_AD = xcgui.NewProc("XListBox_GetCount_AD")
	xListBox_GetCountColumn_AD = xcgui.NewProc("XListBox_GetCountColumn_AD")
	xListBox_SetSplitLineColor = xcgui.NewProc("XListBox_SetSplitLineColor")
	xListBox_SetDragRectColor = xcgui.NewProc("XListBox_SetDragRectColor")

	// List.
	xList_Create = xcgui.NewProc("XList_Create")
	xList_AddColumn = xcgui.NewProc("XList_AddColumn")
	xList_InsertColumn = xcgui.NewProc("XList_InsertColumn")
	xList_EnableMultiSel = xcgui.NewProc("XList_EnableMultiSel")
	xList_EnableDragChangeColumnWidth = xcgui.NewProc("XList_EnableDragChangeColumnWidth")
	xList_EnableVScrollBarTop = xcgui.NewProc("XList_EnableVScrollBarTop")
	xList_EnableItemBkFullRow = xcgui.NewProc("XList_EnableItemBkFullRow")
	xList_EnableFixedRowHeight = xcgui.NewProc("XList_EnableFixedRowHeight")
	xList_EnablemTemplateReuse = xcgui.NewProc("XList_EnablemTemplateReuse")
	xList_EnableVirtualTable = xcgui.NewProc("XList_EnableVirtualTable")
	xList_SetVirtualRowCount = xcgui.NewProc("XList_SetVirtualRowCount")
	xList_SetSort = xcgui.NewProc("XList_SetSort")
	xList_SetDrawItemBkFlags = xcgui.NewProc("XList_SetDrawItemBkFlags")
	xList_SetColumnWidth = xcgui.NewProc("XList_SetColumnWidth")
	xList_SetColumnMinWidth = xcgui.NewProc("XList_SetColumnMinWidth")
	xList_SetColumnWidthFixed = xcgui.NewProc("XList_SetColumnWidthFixed")
	xList_GetColumnWidth = xcgui.NewProc("XList_GetColumnWidth")
	xList_GetColumnCount = xcgui.NewProc("XList_GetColumnCount")
	xList_SetItemData = xcgui.NewProc("XList_SetItemData")
	xList_GetItemData = xcgui.NewProc("XList_GetItemData")
	xList_SetSelectItem = xcgui.NewProc("XList_SetSelectItem")
	xList_GetSelectItem = xcgui.NewProc("XList_GetSelectItem")
	xList_GetSelectItemCount = xcgui.NewProc("XList_GetSelectItemCount")
	xList_AddSelectItem = xcgui.NewProc("XList_AddSelectItem")
	xList_SetSelectAll = xcgui.NewProc("XList_SetSelectAll")
	xList_GetSelectAll = xcgui.NewProc("XList_GetSelectAll")
	xList_VisibleItem = xcgui.NewProc("XList_VisibleItem")
	xList_CancelSelectItem = xcgui.NewProc("XList_CancelSelectItem")
	xList_CancelSelectAll = xcgui.NewProc("XList_CancelSelectAll")
	xList_GetHeaderHELE = xcgui.NewProc("XList_GetHeaderHELE")
	xList_DeleteColumn = xcgui.NewProc("XList_DeleteColumn")
	xList_DeleteColumnAll = xcgui.NewProc("XList_DeleteColumnAll")
	xList_BindAdapter = xcgui.NewProc("XList_BindAdapter")
	xList_BindAdapterHeader = xcgui.NewProc("XList_BindAdapterHeader")
	xList_CreateAdapter = xcgui.NewProc("XList_CreateAdapter")
	xList_CreateAdapterHeader = xcgui.NewProc("XList_CreateAdapterHeader")
	xList_GetAdapter = xcgui.NewProc("XList_GetAdapter")
	xList_GetAdapterHeader = xcgui.NewProc("XList_GetAdapterHeader")
	xList_SetItemTemplateXML = xcgui.NewProc("XList_SetItemTemplateXML")
	xList_SetItemTemplateXMLFromString = xcgui.NewProc("XList_SetItemTemplateXMLFromString")
	xList_SetItemTemplate = xcgui.NewProc("XList_SetItemTemplate")
	xList_GetTemplateObject = xcgui.NewProc("XList_GetTemplateObject")
	xList_GetItemIndexFromHXCGUI = xcgui.NewProc("XList_GetItemIndexFromHXCGUI")
	xList_GetHeaderTemplateObject = xcgui.NewProc("XList_GetHeaderTemplateObject")
	xList_GetHeaderItemIndexFromHXCGUI = xcgui.NewProc("XList_GetHeaderItemIndexFromHXCGUI")
	xList_SetHeaderHeight = xcgui.NewProc("XList_SetHeaderHeight")
	xList_GetHeaderHeight = xcgui.NewProc("XList_GetHeaderHeight")
	xList_GetVisibleRowRange = xcgui.NewProc("XList_GetVisibleRowRange")
	xList_SetItemHeightDefault = xcgui.NewProc("XList_SetItemHeightDefault")
	xList_GetItemHeightDefault = xcgui.NewProc("XList_GetItemHeightDefault")
	xList_SetRowSpace = xcgui.NewProc("XList_SetRowSpace")
	xList_GetRowSpace = xcgui.NewProc("XList_GetRowSpace")
	xList_SetLockColumnLeft = xcgui.NewProc("XList_SetLockColumnLeft")
	xList_SetLockColumnRight = xcgui.NewProc("XList_SetLockColumnRight")
	xList_SetLockRowBottom = xcgui.NewProc("XList_SetLockRowBottom")
	xList_SetLockRowBottomOverlap = xcgui.NewProc("XList_SetLockRowBottomOverlap")
	xList_HitTest = xcgui.NewProc("XList_HitTest")
	xList_HitTestOffset = xcgui.NewProc("XList_HitTestOffset")
	xList_RefreshData = xcgui.NewProc("XList_RefreshData")
	xList_RefreshItem = xcgui.NewProc("XList_RefreshItem")
	xList_AddColumnText = xcgui.NewProc("XList_AddColumnText")
	xList_AddColumnImage = xcgui.NewProc("XList_AddColumnImage")
	xList_AddItemText = xcgui.NewProc("XList_AddItemText")
	xList_AddItemTextEx = xcgui.NewProc("XList_AddItemTextEx")
	xList_AddItemImage = xcgui.NewProc("XList_AddItemImage")
	xList_AddItemImageEx = xcgui.NewProc("XList_AddItemImageEx")
	xList_InsertItemText = xcgui.NewProc("XList_InsertItemText")
	xList_InsertItemTextEx = xcgui.NewProc("XList_InsertItemTextEx")
	xList_InsertItemImage = xcgui.NewProc("XList_InsertItemImage")
	xList_InsertItemImageEx = xcgui.NewProc("XList_InsertItemImageEx")
	xList_SetItemText = xcgui.NewProc("XList_SetItemText")
	xList_SetItemTextEx = xcgui.NewProc("XList_SetItemTextEx")
	xList_SetItemImage = xcgui.NewProc("XList_SetItemImage")
	xList_SetItemImageEx = xcgui.NewProc("XList_SetItemImageEx")
	xList_SetItemInt = xcgui.NewProc("XList_SetItemInt")
	xList_SetItemIntEx = xcgui.NewProc("XList_SetItemIntEx")
	xList_SetItemFloat = xcgui.NewProc("XList_SetItemFloat")
	xList_SetItemFloatEx = xcgui.NewProc("XList_SetItemFloatEx")
	xList_GetItemText = xcgui.NewProc("XList_GetItemText")
	xList_GetItemTextEx = xcgui.NewProc("XList_GetItemTextEx")
	xList_GetItemImage = xcgui.NewProc("XList_GetItemImage")
	xList_GetItemImageEx = xcgui.NewProc("XList_GetItemImageEx")
	xList_GetItemInt = xcgui.NewProc("XList_GetItemInt")
	xList_GetItemIntEx = xcgui.NewProc("XList_GetItemIntEx")
	xList_GetItemFloat = xcgui.NewProc("XList_GetItemFloat")
	xList_GetItemFloatEx = xcgui.NewProc("XList_GetItemFloatEx")
	xList_DeleteItem = xcgui.NewProc("XList_DeleteItem")
	xList_DeleteItemEx = xcgui.NewProc("XList_DeleteItemEx")
	xList_DeleteItemAll = xcgui.NewProc("XList_DeleteItemAll")
	xList_DeleteColumnAll_AD = xcgui.NewProc("XList_DeleteColumnAll_AD")
	xList_GetCount_AD = xcgui.NewProc("XList_GetCount_AD")
	xList_GetCountColumn_AD = xcgui.NewProc("XList_GetCountColumn_AD")
	xList_SetSplitLineColor = xcgui.NewProc("XList_SetSplitLineColor")
	xList_SetItemHeight = xcgui.NewProc("XList_SetItemHeight")
	xList_GetItemHeight = xcgui.NewProc("XList_GetItemHeight")
	xList_SetDragRectColor = xcgui.NewProc("XList_SetDragRectColor")

	// ListView.
	xListView_Create = xcgui.NewProc("XListView_Create")
	xListView_CreateAdapter = xcgui.NewProc("XListView_CreateAdapter")
	xListView_BindAdapter = xcgui.NewProc("XListView_BindAdapter")
	xListView_GetAdapter = xcgui.NewProc("XListView_GetAdapter")
	xListView_SetItemTemplateXML = xcgui.NewProc("XListView_SetItemTemplateXML")
	xListView_SetItemTemplateXMLFromString = xcgui.NewProc("XListView_SetItemTemplateXMLFromString")
	xListView_SetItemTemplate = xcgui.NewProc("XListView_SetItemTemplate")
	xListView_GetTemplateObject = xcgui.NewProc("XListView_GetTemplateObject")
	xListView_GetTemplateObjectGroup = xcgui.NewProc("XListView_GetTemplateObjectGroup")
	xListView_GetItemIDFromHXCGUI = xcgui.NewProc("XListView_GetItemIDFromHXCGUI")
	xListView_HitTest = xcgui.NewProc("XListView_HitTest")
	xListView_HitTestOffset = xcgui.NewProc("XListView_HitTestOffset")
	xListView_EnableMultiSel = xcgui.NewProc("XListView_EnableMultiSel")
	xListView_EnablemTemplateReuse = xcgui.NewProc("XListView_EnablemTemplateReuse")
	xListView_EnableVirtualTable = xcgui.NewProc("XListView_EnableVirtualTable")
	xListView_SetVirtualItemCount = xcgui.NewProc("XListView_SetVirtualItemCount")
	xListView_SetDrawItemBkFlags = xcgui.NewProc("XListView_SetDrawItemBkFlags")
	xListView_SetSelectItem = xcgui.NewProc("XListView_SetSelectItem")
	xListView_GetSelectItem = xcgui.NewProc("XListView_GetSelectItem")
	xListView_AddSelectItem = xcgui.NewProc("XListView_AddSelectItem")
	xListView_VisibleItem = xcgui.NewProc("XListView_VisibleItem")
	xListView_GetVisibleItemRange = xcgui.NewProc("XListView_GetVisibleItemRange")
	xListView_GetSelectItemCount = xcgui.NewProc("XListView_GetSelectItemCount")
	xListView_GetSelectAll = xcgui.NewProc("XListView_GetSelectAll")
	xListView_SetSelectAll = xcgui.NewProc("XListView_SetSelectAll")
	xListView_CancelSelectAll = xcgui.NewProc("XListView_CancelSelectAll")
	xListView_SetColumnSpace = xcgui.NewProc("XListView_SetColumnSpace")
	xListView_SetRowSpace = xcgui.NewProc("XListView_SetRowSpace")
	xListView_SetItemSize = xcgui.NewProc("XListView_SetItemSize")
	xListView_GetItemSize = xcgui.NewProc("XListView_GetItemSize")
	xListView_SetGroupHeight = xcgui.NewProc("XListView_SetGroupHeight")
	xListView_GetGroupHeight = xcgui.NewProc("XListView_GetGroupHeight")
	xListView_SetGroupUserData = xcgui.NewProc("XListView_SetGroupUserData")
	xListView_SetItemUserData = xcgui.NewProc("XListView_SetItemUserData")
	xListView_GetGroupUserData = xcgui.NewProc("XListView_GetGroupUserData")
	xListView_GetItemUserData = xcgui.NewProc("XListView_GetItemUserData")
	xListView_RefreshData = xcgui.NewProc("XListView_RefreshData")
	xListView_RefreshItem = xcgui.NewProc("XListView_RefreshItem")
	xListView_ExpandGroup = xcgui.NewProc("XListView_ExpandGroup")
	xListView_Group_AddColumn = xcgui.NewProc("XListView_Group_AddColumn")
	xListView_Group_AddItemText = xcgui.NewProc("XListView_Group_AddItemText")
	xListView_Group_AddItemTextEx = xcgui.NewProc("XListView_Group_AddItemTextEx")
	xListView_Group_AddItemImage = xcgui.NewProc("XListView_Group_AddItemImage")
	xListView_Group_AddItemImageEx = xcgui.NewProc("XListView_Group_AddItemImageEx")
	xListView_Group_SetText = xcgui.NewProc("XListView_Group_SetText")
	xListView_Group_SetTextEx = xcgui.NewProc("XListView_Group_SetTextEx")
	xListView_Group_SetImage = xcgui.NewProc("XListView_Group_SetImage")
	xListView_Group_SetImageEx = xcgui.NewProc("XListView_Group_SetImageEx")
	xListView_Group_GetCount = xcgui.NewProc("XListView_Group_GetCount")
	xListView_Item_GetCount = xcgui.NewProc("XListView_Item_GetCount")
	xListView_Item_AddColumn = xcgui.NewProc("XListView_Item_AddColumn")
	xListView_Item_AddItemText = xcgui.NewProc("XListView_Item_AddItemText")
	xListView_Item_AddItemTextEx = xcgui.NewProc("XListView_Item_AddItemTextEx")
	xListView_Item_AddItemImage = xcgui.NewProc("XListView_Item_AddItemImage")
	xListView_Item_AddItemImageEx = xcgui.NewProc("XListView_Item_AddItemImageEx")
	xListView_Item_SetText = xcgui.NewProc("XListView_Item_SetText")
	xListView_Item_SetTextEx = xcgui.NewProc("XListView_Item_SetTextEx")
	xListView_Item_SetImage = xcgui.NewProc("XListView_Item_SetImage")
	xListView_Item_SetImageEx = xcgui.NewProc("XListView_Item_SetImageEx")
	xListView_Group_DeleteItem = xcgui.NewProc("XListView_Group_DeleteItem")
	xListView_Group_DeleteAllChildItem = xcgui.NewProc("XListView_Group_DeleteAllChildItem")
	xListView_Item_DeleteItem = xcgui.NewProc("XListView_Item_DeleteItem")
	xListView_DeleteAll = xcgui.NewProc("XListView_DeleteAll")
	xListView_DeleteAllGroup = xcgui.NewProc("XListView_DeleteAllGroup")
	xListView_DeleteAllItem = xcgui.NewProc("XListView_DeleteAllItem")
	xListView_DeleteColumnGroup = xcgui.NewProc("XListView_DeleteColumnGroup")
	xListView_DeleteColumnItem = xcgui.NewProc("XListView_DeleteColumnItem")
	xListView_Item_GetTextEx = xcgui.NewProc("XListView_Item_GetTextEx")
	xListView_Item_GetImageEx = xcgui.NewProc("XListView_Item_GetImageEx")
	xListView_Group_GetText = xcgui.NewProc("XListView_Group_GetText")
	xListView_Group_GetTextEx = xcgui.NewProc("XListView_Group_GetTextEx")
	xListView_Group_GetImage = xcgui.NewProc("XListView_Group_GetImage")
	xListView_Group_GetImageEx = xcgui.NewProc("XListView_Group_GetImageEx")
	xListView_Item_GetText = xcgui.NewProc("XListView_Item_GetText")
	xListView_Item_GetImage = xcgui.NewProc("XListView_Item_GetImage")
	xListView_SetDragRectColor = xcgui.NewProc("XListView_SetDragRectColor")

	// MenuBar.
	xMenuBar_Create = xcgui.NewProc("XMenuBar_Create")
	xMenuBar_AddButton = xcgui.NewProc("XMenuBar_AddButton")
	xMenuBar_SetButtonHeight = xcgui.NewProc("XMenuBar_SetButtonHeight")
	xMenuBar_GetMenu = xcgui.NewProc("XMenuBar_GetMenu")
	xMenuBar_DeleteButton = xcgui.NewProc("XMenuBar_DeleteButton")
	xMenuBar_EnableAutoWidth = xcgui.NewProc("XMenuBar_EnableAutoWidth")
	xMenuBar_GetButton = xcgui.NewProc("XMenuBar_GetButton")

	// Pane.
	xPane_Create = xcgui.NewProc("XPane_Create")
	xPane_SetView = xcgui.NewProc("XPane_SetView")
	xPane_SetTitle = xcgui.NewProc("XPane_SetTitle")
	xPane_GetTitle = xcgui.NewProc("XPane_GetTitle")
	xPane_SetCaptionHeight = xcgui.NewProc("XPane_SetCaptionHeight")
	xPane_GetCaptionHeight = xcgui.NewProc("XPane_GetCaptionHeight")
	xPane_IsShowPane = xcgui.NewProc("XPane_IsShowPane")
	xPane_SetSize = xcgui.NewProc("XPane_SetSize")
	xPane_GetState = xcgui.NewProc("XPane_GetState")
	xPane_GetViewRect = xcgui.NewProc("XPane_GetViewRect")
	xPane_HidePane = xcgui.NewProc("XPane_HidePane")
	xPane_ShowPane = xcgui.NewProc("XPane_ShowPane")
	xPane_DockPane = xcgui.NewProc("XPane_DockPane")
	xPane_LockPane = xcgui.NewProc("XPane_LockPane")
	xPane_FloatPane = xcgui.NewProc("XPane_FloatPane")
	xPane_DrawPane = xcgui.NewProc("XPane_DrawPane")
	xPane_SetSelect = xcgui.NewProc("XPane_SetSelect")
	xPane_IsGroupActivate = xcgui.NewProc("XPane_IsGroupActivate")

	// ScrollBar.
	xSBar_Create = xcgui.NewProc("XSBar_Create")
	xSBar_SetRange = xcgui.NewProc("XSBar_SetRange")
	xSBar_GetRange = xcgui.NewProc("XSBar_GetRange")
	xSBar_ShowButton = xcgui.NewProc("XSBar_ShowButton")
	xSBar_SetSliderLength = xcgui.NewProc("XSBar_SetSliderLength")
	xSBar_SetSliderMinLength = xcgui.NewProc("XSBar_SetSliderMinLength")
	xSBar_SetSliderPadding = xcgui.NewProc("XSBar_SetSliderPadding")
	xSBar_EnableHorizon = xcgui.NewProc("XSBar_EnableHorizon")
	xSBar_GetSliderMaxLength = xcgui.NewProc("XSBar_GetSliderMaxLength")
	xSBar_ScrollUp = xcgui.NewProc("XSBar_ScrollUp")
	xSBar_ScrollDown = xcgui.NewProc("XSBar_ScrollDown")
	xSBar_ScrollTop = xcgui.NewProc("XSBar_ScrollTop")
	xSBar_ScrollBottom = xcgui.NewProc("XSBar_ScrollBottom")
	xSBar_ScrollPos = xcgui.NewProc("XSBar_ScrollPos")
	xSBar_GetButtonUp = xcgui.NewProc("XSBar_GetButtonUp")
	xSBar_GetButtonDown = xcgui.NewProc("XSBar_GetButtonDown")
	xSBar_GetButtonSlider = xcgui.NewProc("XSBar_GetButtonSlider")

	// ScrollView.
	xSView_Create = xcgui.NewProc("XSView_Create")
	xSView_SetTotalSize = xcgui.NewProc("XSView_SetTotalSize")
	xSView_GetTotalSize = xcgui.NewProc("XSView_GetTotalSize")
	xSView_SetLineSize = xcgui.NewProc("XSView_SetLineSize")
	xSView_GetLineSize = xcgui.NewProc("XSView_GetLineSize")
	xSView_SetScrollBarSize = xcgui.NewProc("XSView_SetScrollBarSize")
	xSView_GetViewPosH = xcgui.NewProc("XSView_GetViewPosH")
	xSView_GetViewPosV = xcgui.NewProc("XSView_GetViewPosV")
	xSView_GetViewWidth = xcgui.NewProc("XSView_GetViewWidth")
	xSView_GetViewHeight = xcgui.NewProc("XSView_GetViewHeight")
	xSView_GetViewRect = xcgui.NewProc("XSView_GetViewRect")
	xSView_GetScrollBarH = xcgui.NewProc("XSView_GetScrollBarH")
	xSView_GetScrollBarV = xcgui.NewProc("XSView_GetScrollBarV")
	xSView_ScrollPosH = xcgui.NewProc("XSView_ScrollPosH")
	xSView_ScrollPosV = xcgui.NewProc("XSView_ScrollPosV")
	xSView_ScrollPosXH = xcgui.NewProc("XSView_ScrollPosXH")
	xSView_ScrollPosYV = xcgui.NewProc("XSView_ScrollPosYV")
	xSView_ShowSBarH = xcgui.NewProc("XSView_ShowSBarH")
	xSView_ShowSBarV = xcgui.NewProc("XSView_ShowSBarV")
	xSView_EnableAutoShowScrollBar = xcgui.NewProc("XSView_EnableAutoShowScrollBar")
	xSView_ScrollLeftLine = xcgui.NewProc("XSView_ScrollLeftLine")
	xSView_ScrollRightLine = xcgui.NewProc("XSView_ScrollRightLine")
	xSView_ScrollTopLine = xcgui.NewProc("XSView_ScrollTopLine")
	xSView_ScrollBottomLine = xcgui.NewProc("XSView_ScrollBottomLine")
	xSView_ScrollLeft = xcgui.NewProc("XSView_ScrollLeft")
	xSView_ScrollRight = xcgui.NewProc("XSView_ScrollRight")
	xSView_ScrollTop = xcgui.NewProc("XSView_ScrollTop")
	xSView_ScrollBottom = xcgui.NewProc("XSView_ScrollBottom")

	// SliderBar.
	xSliderBar_Create = xcgui.NewProc("XSliderBar_Create")
	xSliderBar_SetRange = xcgui.NewProc("XSliderBar_SetRange")
	xSliderBar_GetRange = xcgui.NewProc("XSliderBar_GetRange")
	xSliderBar_SetImageLoad = xcgui.NewProc("XSliderBar_SetImageLoad")
	xSliderBar_SetButtonWidth = xcgui.NewProc("XSliderBar_SetButtonWidth")
	xSliderBar_SetButtonHeight = xcgui.NewProc("XSliderBar_SetButtonHeight")
	xSliderBar_SetPos = xcgui.NewProc("XSliderBar_SetPos")
	xSliderBar_GetPos = xcgui.NewProc("XSliderBar_GetPos")
	xSliderBar_GetButton = xcgui.NewProc("XSliderBar_GetButton")
	xSliderBar_EnableHorizon = xcgui.NewProc("XSliderBar_EnableHorizon")

	// TabBar.
	xTabBar_Create = xcgui.NewProc("XTabBar_Create")
	xTabBar_AddLabel = xcgui.NewProc("XTabBar_AddLabel")
	xTabBar_InsertLabel = xcgui.NewProc("XTabBar_InsertLabel")
	xTabBar_MoveLabel = xcgui.NewProc("XTabBar_MoveLabel")
	xTabBar_DeleteLabel = xcgui.NewProc("XTabBar_DeleteLabel")
	xTabBar_DeleteLabelAll = xcgui.NewProc("XTabBar_DeleteLabelAll")
	xTabBar_GetLabel = xcgui.NewProc("XTabBar_GetLabel")
	xTabBar_GetLabelClose = xcgui.NewProc("XTabBar_GetLabelClose")
	xTabBar_GetButtonLeft = xcgui.NewProc("XTabBar_GetButtonLeft")
	xTabBar_GetButtonRight = xcgui.NewProc("XTabBar_GetButtonRight")
	xTabBar_GetButtonDropMenu = xcgui.NewProc("XTabBar_GetButtonDropMenu")
	xTabBar_GetSelect = xcgui.NewProc("XTabBar_GetSelect")
	xTabBar_GetLabelSpacing = xcgui.NewProc("XTabBar_GetLabelSpacing")
	xTabBar_GetLabelCount = xcgui.NewProc("XTabBar_GetLabelCount")
	xTabBar_GetindexByEle = xcgui.NewProc("XTabBar_GetindexByEle")
	xTabBar_SetLabelSpacing = xcgui.NewProc("XTabBar_SetLabelSpacing")
	xTabBar_SetPadding = xcgui.NewProc("XTabBar_SetPadding")
	xTabBar_SetSelect = xcgui.NewProc("XTabBar_SetSelect")
	xTabBar_SetUp = xcgui.NewProc("XTabBar_SetUp")
	xTabBar_SetDown = xcgui.NewProc("XTabBar_SetDown")
	xTabBar_EnableTile = xcgui.NewProc("XTabBar_EnableTile")
	xTabBar_EnableDropMenu = xcgui.NewProc("XTabBar_EnableDropMenu")
	xTabBar_EnableClose = xcgui.NewProc("XTabBar_EnableClose")
	xTabBar_SetCloseSize = xcgui.NewProc("XTabBar_SetCloseSize")
	xTabBar_SetTurnButtonSize = xcgui.NewProc("XTabBar_SetTurnButtonSize")
	xTabBar_SetLabelWidth = xcgui.NewProc("XTabBar_SetLabelWidth")
	xTabBar_ShowLabel = xcgui.NewProc("XTabBar_ShowLabel")

	// ToolBar.
	xToolBar_Create = xcgui.NewProc("XToolBar_Create")
	xToolBar_InsertEle = xcgui.NewProc("XToolBar_InsertEle")
	xToolBar_InsertSeparator = xcgui.NewProc("XToolBar_InsertSeparator")
	xToolBar_EnableButtonMenu = xcgui.NewProc("XToolBar_EnableButtonMenu")
	xToolBar_GetEle = xcgui.NewProc("XToolBar_GetEle")
	xToolBar_GetButtonLeft = xcgui.NewProc("XToolBar_GetButtonLeft")
	xToolBar_GetButtonRight = xcgui.NewProc("XToolBar_GetButtonRight")
	xToolBar_GetButtonMenu = xcgui.NewProc("XToolBar_GetButtonMenu")
	xToolBar_SetSpace = xcgui.NewProc("XToolBar_SetSpace")
	xToolBar_DeleteEle = xcgui.NewProc("XToolBar_DeleteEle")
	xToolBar_DeleteAllEle = xcgui.NewProc("XToolBar_DeleteAllEle")

	// Tree.
	xTree_Create = xcgui.NewProc("XTree_Create")
	xTree_EnableDragItem = xcgui.NewProc("XTree_EnableDragItem")
	xTree_EnableConnectLine = xcgui.NewProc("XTree_EnableConnectLine")
	xTree_EnableExpand = xcgui.NewProc("XTree_EnableExpand")
	xTree_EnablemTemplateReuse = xcgui.NewProc("XTree_EnablemTemplateReuse")
	xTree_SetConnectLineColor = xcgui.NewProc("XTree_SetConnectLineColor")
	xTree_SetExpandButtonSize = xcgui.NewProc("XTree_SetExpandButtonSize")
	xTree_SetConnectLineLength = xcgui.NewProc("XTree_SetConnectLineLength")
	xTree_SetDragInsertPositionColor = xcgui.NewProc("XTree_SetDragInsertPositionColor")
	xTree_SetItemTemplateXML = xcgui.NewProc("XTree_SetItemTemplateXML")
	xTree_SetItemTemplateXMLSel = xcgui.NewProc("XTree_SetItemTemplateXMLSel")
	xTree_SetItemTemplate = xcgui.NewProc("XTree_SetItemTemplate")
	xTree_SetItemTemplateSel = xcgui.NewProc("XTree_SetItemTemplateSel")
	xTree_SetItemTemplateXMLFromString = xcgui.NewProc("XTree_SetItemTemplateXMLFromString")
	xTree_SetItemTemplateXMLSelFromString = xcgui.NewProc("XTree_SetItemTemplateXMLSelFromString")
	xTree_SetDrawItemBkFlags = xcgui.NewProc("XTree_SetDrawItemBkFlags")
	xTree_SetItemData = xcgui.NewProc("XTree_SetItemData")
	xTree_GetItemData = xcgui.NewProc("XTree_GetItemData")
	xTree_SetSelectItem = xcgui.NewProc("XTree_SetSelectItem")
	xTree_GetSelectItem = xcgui.NewProc("XTree_GetSelectItem")
	xTree_VisibleItem = xcgui.NewProc("XTree_VisibleItem")
	xTree_IsExpand = xcgui.NewProc("XTree_IsExpand")
	xTree_ExpandItem = xcgui.NewProc("XTree_ExpandItem")
	xTree_ExpandAllChildItem = xcgui.NewProc("XTree_ExpandAllChildItem")
	xTree_HitTest = xcgui.NewProc("XTree_HitTest")
	xTree_HitTestOffset = xcgui.NewProc("XTree_HitTestOffset")
	xTree_GetFirstChildItem = xcgui.NewProc("XTree_GetFirstChildItem")
	xTree_GetEndChildItem = xcgui.NewProc("XTree_GetEndChildItem")
	xTree_GetPrevSiblingItem = xcgui.NewProc("XTree_GetPrevSiblingItem")
	xTree_GetNextSiblingItem = xcgui.NewProc("XTree_GetNextSiblingItem")
	xTree_GetParentItem = xcgui.NewProc("XTree_GetParentItem")
	xTree_CreateAdapter = xcgui.NewProc("XTree_CreateAdapter")
	xTree_BindAdapter = xcgui.NewProc("XTree_BindAdapter")
	xTree_GetAdapter = xcgui.NewProc("XTree_GetAdapter")
	xTree_RefreshData = xcgui.NewProc("XTree_RefreshData")
	xTree_RefreshItem = xcgui.NewProc("XTree_RefreshItem")
	xTree_SetIndentation = xcgui.NewProc("XTree_SetIndentation")
	xTree_GetIndentation = xcgui.NewProc("XTree_GetIndentation")
	xTree_SetItemHeightDefault = xcgui.NewProc("XTree_SetItemHeightDefault")
	xTree_GetItemHeightDefault = xcgui.NewProc("XTree_GetItemHeightDefault")
	xTree_SetItemHeight = xcgui.NewProc("XTree_SetItemHeight")
	xTree_GetItemHeight = xcgui.NewProc("XTree_GetItemHeight")
	xTree_SetRowSpace = xcgui.NewProc("XTree_SetRowSpace")
	xTree_GetRowSpace = xcgui.NewProc("XTree_GetRowSpace")
	xTree_MoveItem = xcgui.NewProc("XTree_MoveItem")
	xTree_GetTemplateObject = xcgui.NewProc("XTree_GetTemplateObject")
	xTree_GetItemIDFromHXCGUI = xcgui.NewProc("XTree_GetItemIDFromHXCGUI")
	xTree_InsertItemText = xcgui.NewProc("XTree_InsertItemText")
	xTree_InsertItemTextEx = xcgui.NewProc("XTree_InsertItemTextEx")
	xTree_InsertItemImage = xcgui.NewProc("XTree_InsertItemImage")
	xTree_InsertItemImageEx = xcgui.NewProc("XTree_InsertItemImageEx")
	xTree_GetCount = xcgui.NewProc("XTree_GetCount")
	xTree_GetCountColumn = xcgui.NewProc("XTree_GetCountColumn")
	xTree_SetItemText = xcgui.NewProc("XTree_SetItemText")
	xTree_SetItemTextEx = xcgui.NewProc("XTree_SetItemTextEx")
	xTree_SetItemImage = xcgui.NewProc("XTree_SetItemImage")
	xTree_SetItemImageEx = xcgui.NewProc("XTree_SetItemImageEx")
	xTree_GetItemText = xcgui.NewProc("XTree_GetItemText")
	xTree_GetItemTextEx = xcgui.NewProc("XTree_GetItemTextEx")
	xTree_GetItemImage = xcgui.NewProc("XTree_GetItemImage")
	xTree_GetItemImageEx = xcgui.NewProc("XTree_GetItemImageEx")
	xTree_DeleteItem = xcgui.NewProc("XTree_DeleteItem")
	xTree_DeleteItemAll = xcgui.NewProc("XTree_DeleteItemAll")
	xTree_DeleteColumnAll = xcgui.NewProc("XTree_DeleteColumnAll")
	xTree_SetSplitLineColor = xcgui.NewProc("XTree_SetSplitLineColor")

	// DateTime.
	xDateTime_Create = xcgui.NewProc("XDateTime_Create")
	xDateTime_SetStyle = xcgui.NewProc("XDateTime_SetStyle")
	xDateTime_GetStyle = xcgui.NewProc("XDateTime_GetStyle")
	xDateTime_EnableSplitSlash = xcgui.NewProc("XDateTime_EnableSplitSlash")
	xDateTime_GetButton = xcgui.NewProc("XDateTime_GetButton")
	xDateTime_GetSelBkColor = xcgui.NewProc("XDateTime_GetSelBkColor")
	xDateTime_SetSelBkColor = xcgui.NewProc("XDateTime_SetSelBkColor")
	xDateTime_GetDate = xcgui.NewProc("XDateTime_GetDate")
	xDateTime_SetDate = xcgui.NewProc("XDateTime_SetDate")
	xDateTime_GetTime = xcgui.NewProc("XDateTime_GetTime")
	xDateTime_SetTime = xcgui.NewProc("XDateTime_SetTime")
	xDateTime_Popup = xcgui.NewProc("XDateTime_Popup")

	// MonthCal.
	xMonthCal_Create = xcgui.NewProc("XMonthCal_Create")
	xMonthCal_GetButton = xcgui.NewProc("XMonthCal_GetButton")
	xMonthCal_SetToday = xcgui.NewProc("XMonthCal_SetToday")
	xMonthCal_GetToday = xcgui.NewProc("XMonthCal_GetToday")
	xMonthCal_GetSelDate = xcgui.NewProc("XMonthCal_GetSelDate")
	xMonthCal_SetTextColor = xcgui.NewProc("XMonthCal_SetTextColor")

	// XcObjectBase.
	xObj_GetType = xcgui.NewProc("XObj_GetType")
	xObj_GetTypeBase = xcgui.NewProc("XObj_GetTypeBase")
	xObj_GetTypeEx = xcgui.NewProc("XObj_GetTypeEx")
	xObj_SetTypeEx = xcgui.NewProc("XObj_SetTypeEx")

	// Animation.
	xAnima_Run = xcgui.NewProc("XAnima_Run")
	xAnima_Release = xcgui.NewProc("XAnima_Release")
	xAnima_ReleaseEx = xcgui.NewProc("XAnima_ReleaseEx")
	xAnima_Create = xcgui.NewProc("XAnima_Create")
	xAnima_Move = xcgui.NewProc("XAnima_Move")
	xAnima_MoveEx = xcgui.NewProc("XAnima_MoveEx")
	xAnima_Rotate = xcgui.NewProc("XAnima_Rotate")
	xAnima_RotateEx = xcgui.NewProc("XAnima_RotateEx")
	xAnima_Scale = xcgui.NewProc("XAnima_Scale")
	xAnima_ScaleSize = xcgui.NewProc("XAnima_ScaleSize")
	xAnima_Alpha = xcgui.NewProc("XAnima_Alpha")
	xAnima_AlphaEx = xcgui.NewProc("XAnima_AlphaEx")
	xAnima_Color = xcgui.NewProc("XAnima_Color")
	xAnima_ColorEx = xcgui.NewProc("XAnima_ColorEx")
	xAnima_LayoutWidth = xcgui.NewProc("XAnima_LayoutWidth")
	xAnima_LayoutHeight = xcgui.NewProc("XAnima_LayoutHeight")
	xAnima_LayoutSize = xcgui.NewProc("XAnima_LayoutSize")
	xAnima_Delay = xcgui.NewProc("XAnima_Delay")
	xAnima_Show = xcgui.NewProc("XAnima_Show")
	xAnimaGroup_Create = xcgui.NewProc("XAnimaGroup_Create")
	xAnimaGroup_AddItem = xcgui.NewProc("XAnimaGroup_AddItem")
	xAnimaRotate_SetCenter = xcgui.NewProc("XAnimaRotate_SetCenter")
	xAnimaScale_SetPosition = xcgui.NewProc("XAnimaScale_SetPosition")
	xAnima_GetObjectUI = xcgui.NewProc("XAnima_GetObjectUI")
	xAnimaItem_EnableCompleteRelease = xcgui.NewProc("XAnimaItem_EnableCompleteRelease")
	xAnima_EnableAutoDestroy = xcgui.NewProc("XAnima_EnableAutoDestroy")
	xAnima_DestroyObjectUI = xcgui.NewProc("XAnima_DestroyObjectUI")
	xAnima_SetCallBack = xcgui.NewProc("XAnima_SetCallBack")
	xAnima_SetUserData = xcgui.NewProc("XAnima_SetUserData")
	xAnima_GetUserData = xcgui.NewProc("XAnima_GetUserData")
	xAnima_Stop = xcgui.NewProc("XAnima_Stop")
	xAnima_Start = xcgui.NewProc("XAnima_Start")
	xAnima_Pause = xcgui.NewProc("XAnima_Pause")
	xAnimaItem_SetCallback = xcgui.NewProc("XAnimaItem_SetCallback")
	xAnimaItem_SetUserData = xcgui.NewProc("XAnimaItem_SetUserData")
	xAnimaItem_GetUserData = xcgui.NewProc("XAnimaItem_GetUserData")
	xAnimaItem_EnableAutoDestroy = xcgui.NewProc("XAnimaItem_EnableAutoDestroy")
	xAnima_DelayEx = xcgui.NewProc("XAnima_DelayEx")
	xAnimaMove_SetFlag = xcgui.NewProc("XAnimaMove_SetFlag")

	// 通知消息.
	xNotifyMsg_Popup = xcgui.NewProc("XNotifyMsg_Popup")
	xNotifyMsg_PopupEx = xcgui.NewProc("XNotifyMsg_PopupEx")
	xNotifyMsg_SetDuration = xcgui.NewProc("XNotifyMsg_SetDuration")
	xNotifyMsg_SetParentMargin = xcgui.NewProc("XNotifyMsg_SetParentMargin")
	xNotifyMsg_WindowPopup = xcgui.NewProc("XNotifyMsg_WindowPopup")
	xNotifyMsg_WindowPopupEx = xcgui.NewProc("XNotifyMsg_WindowPopupEx")
	xNotifyMsg_SetCaptionHeight = xcgui.NewProc("XNotifyMsg_SetCaptionHeight")
	xNotifyMsg_SetWidth = xcgui.NewProc("XNotifyMsg_SetWidth")
	xNotifyMsg_SetSpace = xcgui.NewProc("XNotifyMsg_SetSpace")
	xNotifyMsg_SetBorderSize = xcgui.NewProc("XNotifyMsg_SetBorderSize")

	// 背景对象.
	xBkObj_SetMargin = xcgui.NewProc("XBkObj_SetMargin")
	xBkObj_SetAlign = xcgui.NewProc("XBkObj_SetAlign")
	xBkObj_SetImage = xcgui.NewProc("XBkObj_SetImage")
	xBkObj_SetRotate = xcgui.NewProc("XBkObj_SetRotate")
	xBkObj_SetFillColor = xcgui.NewProc("XBkObj_SetFillColor")
	xBkObj_SetBorderWidth = xcgui.NewProc("XBkObj_SetBorderWidth")
	xBkObj_SetBorderColor = xcgui.NewProc("XBkObj_SetBorderColor")
	xBkObj_SetRectRoundAngle = xcgui.NewProc("XBkObj_SetRectRoundAngle")
	xBkObj_EnableFill = xcgui.NewProc("XBkObj_EnableFill")
	xBkObj_EnableBorder = xcgui.NewProc("XBkObj_EnableBorder")
	xBkObj_SetText = xcgui.NewProc("XBkObj_SetText")
	xBkObj_SetFont = xcgui.NewProc("XBkObj_SetFont")
	xBkObj_SetTextAlign = xcgui.NewProc("XBkObj_SetTextAlign")
	xBkObj_GetMargin = xcgui.NewProc("XBkObj_GetMargin")
	xBkObj_GetAlign = xcgui.NewProc("XBkObj_GetAlign")
	xBkObj_GetImage = xcgui.NewProc("XBkObj_GetImage")
	xBkObj_GetRotate = xcgui.NewProc("XBkObj_GetRotate")
	xBkObj_GetFillColor = xcgui.NewProc("XBkObj_GetFillColor")
	xBkObj_GetBorderColor = xcgui.NewProc("XBkObj_GetBorderColor")
	xBkObj_GetBorderWidth = xcgui.NewProc("XBkObj_GetBorderWidth")
	xBkObj_GetRectRoundAngle = xcgui.NewProc("XBkObj_GetRectRoundAngle")
	xBkObj_IsFill = xcgui.NewProc("XBkObj_IsFill")
	xBkObj_IsBorder = xcgui.NewProc("XBkObj_IsBorder")
	xBkObj_GetText = xcgui.NewProc("XBkObj_GetText")
	xBkObj_GetFont = xcgui.NewProc("XBkObj_GetFont")
	xBkObj_GetTextAlign = xcgui.NewProc("XBkObj_GetTextAlign")

}

// Font_Info_Name 将[32]uint16转换到string.
//	@param arr [32]uint16.
//	@return string
//
func Font_Info_Name(arr [32]uint16) string {
	return syscall.UTF16ToString(arr[0:])
}

// ABGR 根据r, g, b, a组合成十进制ABGR 颜色.
//	@param r 红色分量.
//	@param g 绿色分量.
//	@param b 蓝色分量.
//	@param a 透明度.
//	@return int ABGR 颜色.
//
func ABGR(r, g, b, a byte) int {
	return int((uint32(r) & 255) | (uint32(g)&255)<<8 | (uint32(b)&255)<<16 | (uint32(a)&255)<<24)
}

// ABGR2 根据rgb, a组合成十进制ABGR 颜色.
//	@param rgb RGB颜色.
//	@param a 透明度.
//	@return int ABGR 颜色.
//
func ABGR2(rgb int, a byte) int {
	return int((uint32(rgb) & 16777215) | (uint32(a)&255)<<24)
}

// RGB 根据r, g, b组合成十进制RGB颜色.
//	@param r 红色分量.
//	@param g 绿色分量.
//	@param b 蓝色分量.
//	@return int RGB颜色.
//
func RGB(r, g, b byte) int {
	return int(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

// RGBA 根据r, g, b, a组合成十进制ABGR 颜色. 和 ABGR 函数一模一样, 只是为了符合部分人使用习惯.
//	@param r 红色分量.
//	@param g 绿色分量.
//	@param b 蓝色分量.
//	@param a 透明度.
//	@return int ABGR 颜色.
//
func RGBA(r, g, b, a byte) int {
	return int((uint32(r) & 255) | (uint32(g)&255)<<8 | (uint32(b)&255)<<16 | (uint32(a)&255)<<24)
}

// RGBA2 根据rgb, a组合成十进制ABGR 颜色. 和 ABGR2 函数一模一样, 只是为了符合部分人使用习惯.
//	@param rgb RGB颜色.
//	@param a 透明度.
//	@return int ABGR 颜色.
//
func RGBA2(rgb int, a byte) int {
	return int((uint32(rgb) & 16777215) | (uint32(a)&255)<<24)
}

// HEXToRGB 将十六进制颜色转换到RGB颜色.
//	@param str 十六进制颜色.
//	@return int RGB颜色.
//
func HEXToRGB(str string) int {
	r, _ := strconv.ParseInt(str[:2], 16, 10)
	g, _ := strconv.ParseInt(str[2:4], 16, 18)
	b, _ := strconv.ParseInt(str[4:], 16, 10)
	return RGB(byte(r), byte(g), byte(b))
}

// ClientToScreen 将窗口客户区坐标转换到屏幕坐标.
//	@param hWindow GUI库窗口资源句柄.
//	@param pPoint 坐标.
//
func ClientToScreen(hWindow int, pPoint *POINT) {
	var r RECT
	XWnd_GetRect(hWindow, &r)
	pPoint.X += r.Left
	pPoint.Y += r.Top
}
