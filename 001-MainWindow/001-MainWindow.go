//go:generate goversioninfo
package main

import (
	"runtime"
	"structs"
	"unsafe"

	"github.com/ebitengine/purego"
	"golang.org/x/sys/windows"
)

const (
	WS_OVERLAPPEDWINDOW = 0x00000000 | 0x00C00000 | 0x00080000 | 0x00040000 | 0x00020000 | 0x00010000
	CW_USEDEFAULT       = ^0x7fffffff
	SW_SHOW             = 5
	WM_DESTROY          = 2
	WM_PAINT            = 15
	WM_SETCURSOR        = 32
)

// HitTest
const (
	HTCLIENT            = 1
)

// Standard Cursor ID
const (
	IDC_ARROW       = 32512
	IDC_IBEAM       = 32513
	IDC_WAIT        = 32514
	IDC_CROSS       = 32515
	IDC_UPARROW     = 32516
	IDC_SIZENWSE    = 32642
	IDC_SIZENESW    = 32643
	IDC_SIZEWE      = 32644
	IDC_SIZENS      = 32645
	IDC_SIZEALL     = 32646
	IDC_NO          = 32648
	IDC_HAND        = 32649
	IDC_APPSTARTING = 32650
	IDC_HELP        = 32651
	IDC_ICON        = 32641
	IDC_SIZE        = 32640
)

const (
	COLOR_3DDKSHADOW              = 21
	COLOR_3DFACE                  = 15
	COLOR_3DHILIGHT               = 20
	COLOR_3DHIGHLIGHT             = 20
	COLOR_3DLIGHT                 = 22
	COLOR_BTNHILIGHT              = 20
	COLOR_3DSHADOW                = 16
	COLOR_ACTIVEBORDER            = 10
	COLOR_ACTIVECAPTION           = 2
	COLOR_APPWORKSPACE            = 12
	COLOR_BACKGROUND              = 1
	COLOR_DESKTOP                 = 1
	COLOR_BTNFACE                 = 15
	COLOR_BTNHIGHLIGHT            = 20
	COLOR_BTNSHADOW               = 16
	COLOR_BTNTEXT                 = 18
	COLOR_CAPTIONTEXT             = 9
	COLOR_GRAYTEXT                = 17
	COLOR_HIGHLIGHT               = 13
	COLOR_HIGHLIGHTTEXT           = 14
	COLOR_INACTIVEBORDER          = 11
	COLOR_INACTIVECAPTION         = 3
	COLOR_INACTIVECAPTIONTEXT     = 19
	COLOR_INFOBK                  = 24
	COLOR_INFOTEXT                = 23
	COLOR_MENU                    = 4
	COLOR_MENUTEXT                = 7
	COLOR_SCROLLBAR               = 0
	COLOR_WINDOW                  = 5
	COLOR_WINDOWFRAME             = 6
	COLOR_WINDOWTEXT              = 8
	COLOR_HOTLIGHT                = 26
	COLOR_GRADIENTACTIVECAPTION   = 27
	COLOR_GRADIENTINACTIVECAPTION = 28
)

type (
	ATOM      uint16
	HANDLE    uintptr
	HINSTANCE HANDLE
	HICON     HANDLE
	HCURSOR   HANDLE
	HBRUSH    HANDLE
	HWND      HANDLE
	HMENU     HANDLE
	HDC       HANDLE
	BOOL      int32
)

type WNDCLASSEX struct {
	_          structs.HostLayout
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   *uint16
	ClassName  *uint16
	IconSm     HICON
}

type RECT struct {
	_      structs.HostLayout
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type POINT struct {
	_ structs.HostLayout
	X int32
	Y int32
}

type MSG struct {
	_       structs.HostLayout
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type PAINTSTRUCT struct {
	_           structs.HostLayout
	Hdc         HDC
	FErase      BOOL
	RcPaint     RECT
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]byte
}

var (
	GetModuleHandle func(modulename *uint16) HINSTANCE
	RegisterClassEx func(w *WNDCLASSEX) ATOM
	CreateWindowEx  func(
		exStyle uint,
		className *uint16,
		windowName *uint16,
		style uint,
		x, y, width, height int,
		parent HWND,
		menu HMENU,
		instance HINSTANCE,
		param unsafe.Pointer,
	) HWND
	AdjustWindowRect func(rect *RECT, style uint, menu bool) bool
	ShowWindow       func(hwnd HWND, cmdshow int) bool
	GetMessage       func(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) int
	TranslateMessage func(msg *MSG) bool
	DispatchMessage  func(msg *MSG) uintptr
	DefWindowProc    func(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr
	PostQuitMessage  func(exitCode int)
	BeginPaint       func(hwnd HWND, lpPaint *PAINTSTRUCT) HDC
	EndPaint         func(hwnd HWND, lpPaint *PAINTSTRUCT) bool
	GetClientRect    func(hwnd HWND, lpRect *RECT) bool
	FillRect         func(hdc HDC, lprc *RECT, hbr HBRUSH) int
	LoadCursor       func(instance HINSTANCE, cursorName *uint16) HCURSOR
	SetCursor        func(cursor HCURSOR)
)

func init() {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll").Handle()
	purego.RegisterLibFunc(&GetModuleHandle, kernel32, "GetModuleHandleW")

	user32 := windows.NewLazySystemDLL("user32.dll").Handle()
	purego.RegisterLibFunc(&RegisterClassEx, user32, "RegisterClassExW")
	purego.RegisterLibFunc(&CreateWindowEx, user32, "CreateWindowExW")
	purego.RegisterLibFunc(&AdjustWindowRect, user32, "AdjustWindowRect")
	purego.RegisterLibFunc(&ShowWindow, user32, "ShowWindow")
	purego.RegisterLibFunc(&GetMessage, user32, "GetMessageW")
	purego.RegisterLibFunc(&TranslateMessage, user32, "TranslateMessage")
	purego.RegisterLibFunc(&DispatchMessage, user32, "DispatchMessageW")
	purego.RegisterLibFunc(&DefWindowProc, user32, "DefWindowProcW")
	purego.RegisterLibFunc(&PostQuitMessage, user32, "PostQuitMessage")
	purego.RegisterLibFunc(&BeginPaint, user32, "BeginPaint")
	purego.RegisterLibFunc(&EndPaint, user32, "EndPaint")
	purego.RegisterLibFunc(&GetClientRect, user32, "GetClientRect")
	purego.RegisterLibFunc(&FillRect, user32, "FillRect")
	purego.RegisterLibFunc(&SetCursor, user32, "SetCursor")
	purego.RegisterLibFunc(&LoadCursor, user32, "LoadCursorW")

	runtime.LockOSThread()
}

func main() {
	className, err := windows.UTF16PtrFromString("Sample Window Class")
	if err != nil {
		panic(err)
	}
	instance := GetModuleHandle(className)

	wc := WNDCLASSEX{
		Size:      uint32(unsafe.Sizeof(WNDCLASSEX{})),
		WndProc:   windows.NewCallback(wndProc),
		Instance:  instance,
		ClassName: className,
	}

	RegisterClassEx(&wc)

	title, err := windows.UTF16PtrFromString("purego 001-MainWindow")
	if err != nil {
		panic(err)
	}

	hwnd := CreateWindowEx(
		0,                   // dwExStyle
		className,           // lpClassName
		title,               // lpWindowName
		WS_OVERLAPPEDWINDOW, // dwStyle
		CW_USEDEFAULT,       // X
		CW_USEDEFAULT,       // Y
		600,                 // nWidth
		400,                 // nHeight
		0,                   // hWndParent
		0,                   // hMenu
		instance,            // instance
		nil,                 // lpParam
	)
	if hwnd == 0 {
		panic(windows.GetLastError())
	}

	ShowWindow(hwnd, SW_SHOW)

	var msg MSG
	for GetMessage(&msg, 0, 0, 0) != 0 {
		TranslateMessage(&msg)
		DispatchMessage(&msg)
	}
}

func LOWORD(x uintptr) uint16 {
	return uint16(x & 0xffff)
}

func HIWORD(x uintptr) uint16 {
	return uint16((x >> 16) & 0xffff)
}

func MAKEINTRESOURCE(resId uint16) *uint16 {
    return (*uint16)(unsafe.Pointer(uintptr(resId)))
}

func wndProc(hwnd HWND, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {

	case WM_PAINT:
		var rect RECT
		GetClientRect(hwnd, &rect)

		var ps PAINTSTRUCT
		hdc := BeginPaint(hwnd, &ps)
		FillRect(hdc, &rect, (HBRUSH)(COLOR_WINDOW+1))
		EndPaint(hwnd, &ps)

	case WM_SETCURSOR:
		hitTest := LOWORD(lparam)

		// クライアント領域上にカーソルがある場合カーソルをセット
		if hitTest == HTCLIENT {
			cursor := LoadCursor(0, MAKEINTRESOURCE(IDC_ARROW)) // 矢印カーソル
			SetCursor(cursor)
			return 1 // 処理済(TRUE)を示す
		}

		// クライアント領域外（タイトルバーや枠線など）は
		// DefWindowProc() に流してシステム既定の挙動とする

	case WM_DESTROY:
		PostQuitMessage(0)
	}

	return DefWindowProc(hwnd, msg, wparam, lparam)
}
