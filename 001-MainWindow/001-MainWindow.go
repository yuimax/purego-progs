//go:generate goversioninfo

// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 The Ebitengine Authors
package main

import (
	"runtime"
	"structs"
	"syscall"
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

// SetBkMode の mode 引数
const (
	TRANSPARENT = 1
	OPAQUE      = 2
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
	_ structs.HostLayout
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
	_ structs.HostLayout
	Left, Top, Right, Bottom int32
}


type POINT struct {
	_ structs.HostLayout
	X, Y int32
}

type MSG struct {
	_ structs.HostLayout
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type PAINTSTRUCT struct {
	_ structs.HostLayout
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
	CreateWindowEx  func(exStyle uint, className, windowName *uint16,
		style uint, x, y, width, height int, parent HWND, menu HMENU,
		instance HINSTANCE, param unsafe.Pointer) HWND
	AdjustWindowRect func(rect *RECT, style uint, menu bool) bool
	ShowWindow       func(hwnd HWND, cmdshow int) bool
	GetMessage       func(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) int
	TranslateMessage func(msg *MSG) bool
	DispatchMessage  func(msg *MSG) uintptr
	DefWindowProc    func(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr
	PostQuitMessage  func(exitCode int)

	BeginPaint    func(hwnd HWND, lpPaint *PAINTSTRUCT) HDC
	EndPaint      func(hwnd HWND, lpPaint *PAINTSTRUCT) bool
	GetClientRect func(hwnd HWND, lpRect *RECT) bool
	FillRect      func(hdc HDC, lprc *RECT, hbr HBRUSH) int

	Rectangle     func(hdc HDC, left, right, top, bottom int32) bool
	Ellipse       func(hdc HDC, left, right, top, bottom int32) bool
	SetBkMode     func(hdc HDC, mode uint32) int
	TextOut       func(hdc HDC, x, y int32, lpString *uint16, c int32) bool
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

	gdi32 := windows.NewLazySystemDLL("gdi32.dll").Handle()
	purego.RegisterLibFunc(&Rectangle, gdi32, "Rectangle")
	purego.RegisterLibFunc(&Ellipse, gdi32, "Ellipse")
	purego.RegisterLibFunc(&SetBkMode, gdi32, "SetBkMode")
	purego.RegisterLibFunc(&TextOut, gdi32, "TextOutW")

	runtime.LockOSThread()
}

func main() {
	className, err := syscall.UTF16PtrFromString("Sample Window Class")
	if err != nil {
		panic(err)
	}
	inst := GetModuleHandle(className)

	wc := WNDCLASSEX{
		Size:      uint32(unsafe.Sizeof(WNDCLASSEX{})),
		WndProc:   syscall.NewCallback(wndProc),
		Instance:  inst,
		ClassName: className,
	}

	RegisterClassEx(&wc)

	wr := RECT{
		Left:   0,
		Top:    0,
		Right:  320,
		Bottom: 240,
	}
	title, err := syscall.UTF16PtrFromString("My Title")
	if err != nil {
		panic(err)
	}
	AdjustWindowRect(&wr, WS_OVERLAPPEDWINDOW, false)
	hwnd := CreateWindowEx(
		0, className,
		title,
		WS_OVERLAPPEDWINDOW,
		CW_USEDEFAULT, CW_USEDEFAULT, int(wr.Right-wr.Left), int(wr.Bottom-wr.Top),
		0, 0, inst, nil,
	)
	if hwnd == 0 {
		panic(syscall.GetLastError())
	}

	ShowWindow(hwnd, SW_SHOW)

	var msg MSG
	for GetMessage(&msg, 0, 0, 0) != 0 {
		TranslateMessage(&msg)
		DispatchMessage(&msg)
	}
}

func wndProc(hwnd HWND, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case WM_PAINT:
		ps := PAINTSTRUCT{}
		hdc := BeginPaint(hwnd, &ps)

		// 1. 背景の塗りつぶし
		clientRect := RECT{}
		GetClientRect(hwnd, &clientRect)
		FillRect(hdc, &clientRect, (HBRUSH)(COLOR_WINDOW + 1))

		// 2. 四角形の描画
		Rectangle(hdc, 50, 50, 200, 150)

		// 3. 楕円の描画
		Ellipse(hdc, 100, 100, 250, 200)

		// 4. テキストの描画

		// 背景透過を設定（テキスト周囲の白背景化を防ぐ）
		SetBkMode(hdc, TRANSPARENT)

		// 文字列をUTF16のスライスに変換する
		// UTF16FromString() は末尾に自動的に 0x0000 を入れるので注意
		text := "こんにちは世界"
		u16s, _ := windows.UTF16FromString(text)

		// TextOutで出力する
		TextOut(hdc, 60, 110,
			(*uint16)(unsafe.Pointer(&u16s[0])),
			(int32)(len(u16s) - 1),
		)

		// 描画終了処理
		EndPaint(hwnd, &ps)

	case WM_DESTROY:
		PostQuitMessage(0)

	}
	return DefWindowProc(hwnd, msg, wparam, lparam)
}
