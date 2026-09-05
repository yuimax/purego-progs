//go:generate goversioninfo

// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 The Ebitengine Authors
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
	WM_DESTROY          = 2
	WM_PAINT            = 15
	WM_SETCURSOR        = 32
	FALSE               = 0
	TRUE                = 1
)

// ShowWindow constants
const (
	SW_HIDE            = 0
	SW_NORMAL          = 1
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
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

// System colors
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

// SetBkMode の mode
const (
	TRANSPARENT = 1
	OPAQUE      = 2
)

// DrawText の format
const (
	DT_TOP                  = 0x00000000
	DT_LEFT                 = 0x00000000
	DT_CENTER               = 0x00000001
	DT_RIGHT                = 0x00000002
	DT_VCENTER              = 0x00000004
	DT_BOTTOM               = 0x00000008
	DT_WORDBREAK            = 0x00000010
	DT_SINGLELINE           = 0x00000020
	DT_EXPANDTABS           = 0x00000040
	DT_TABSTOP              = 0x00000080
	DT_NOCLIP               = 0x00000100
	DT_EXTERNALLEADING      = 0x00000200
	DT_CALCRECT             = 0x00000400
	DT_NOPREFIX             = 0x00000800
	DT_INTERNAL             = 0x00001000
	DT_EDITCONTROL          = 0x00002000
	DT_PATH_ELLIPSIS        = 0x00004000
	DT_END_ELLIPSIS         = 0x00008000
	DT_MODIFYSTRING         = 0x00010000
	DT_RTLREADING           = 0x00020000
	DT_WORD_ELLIPSIS        = 0x00040000
	DT_NOFULLWIDTHCHARBREAK = 0x00080000
	DT_HIDEPREFIX           = 0x00100000
	DT_PREFIXONLY           = 0x00200000
)

// WM_NCHITTEST constants
const (
	HTBORDER      = 18
	HTBOTTOM      = 15
	HTBOTTOMLEFT  = 16
	HTBOTTOMRIGHT = 17
	HTCAPTION     = 2
	HTCLIENT      = 1
	HTCLOSE       = 20
	HTERROR       = -2
	HTGROWBOX     = 4
	HTHELP        = 21
	HTHSCROLL     = 6
	HTLEFT        = 10
	HTMENU        = 5
	HTMAXBUTTON   = 9
	HTMINBUTTON   = 8
	HTNOWHERE     = 0
	HTREDUCE      = 8
	HTRIGHT       = 11
	HTSIZE        = 4
	HTSYSMENU     = 3
	HTTOP         = 12
	HTTOPLEFT     = 13
	HTTOPRIGHT    = 14
	HTTRANSPARENT = -1
	HTVSCROLL     = 7
	HTZOOM        = 9
)

// Font weight
const (
	FW_DONTCARE   = 0
	FW_THIN       = 100
	FW_EXTRALIGHT = 200
	FW_ULTRALIGHT = 200
	FW_LIGHT      = 300
	FW_NORMAL     = 400
	FW_REGULAR    = 400
	FW_MEDIUM     = 500
	FW_SEMIBOLD   = 600
	FW_DEMIBOLD   = 600
	FW_BOLD       = 700
	FW_EXTRABOLD  = 800
	FW_ULTRABOLD  = 800
	FW_HEAVY      = 900
	FW_BLACK      = 900
)

// Font constants
const (
	OUT_DEFAULT_PRECIS        = 0
	OUT_STRING_PRECIS         = 1
	OUT_CHARACTER_PRECIS      = 2
	OUT_STROKE_PRECIS         = 3
	OUT_TT_PRECIS             = 4
	OUT_DEVICE_PRECIS         = 5
	OUT_RASTER_PRECIS         = 6
	OUT_TT_ONLY_PRECIS        = 7
	OUT_OUTLINE_PRECIS        = 8
	OUT_SCREEN_OUTLINE_PRECIS = 9
	OUT_PS_ONLY_PRECIS        = 10

	CLIP_DEFAULT_PRECIS   = 0
	CLIP_CHARACTER_PRECIS = 1
	CLIP_STROKE_PRECIS    = 2
	CLIP_MASK             = 0xf
	CLIP_LH_ANGLES        = (1 << 4)
	CLIP_TT_ALWAYS        = (2 << 4)
	CLIP_DFA_DISABLE      = (4 << 4)
	CLIP_EMBEDDED         = (8 << 4)

	DEFAULT_QUALITY           = 0
	DRAFT_QUALITY             = 1
	PROOF_QUALITY             = 2
	NONANTIALIASED_QUALITY    = 3
	ANTIALIASED_QUALITY       = 4
	CLEARTYPE_QUALITY         = 5
	CLEARTYPE_NATURAL_QUALITY = 6

	DEFAULT_PITCH  = 0
	FIXED_PITCH    = 1
	VARIABLE_PITCH = 2
	MONO_FONT      = 8

	ANSI_CHARSET        = 0
	DEFAULT_CHARSET     = 1
	SYMBOL_CHARSET      = 2
	SHIFTJIS_CHARSET    = 128
	HANGEUL_CHARSET     = 129
	HANGUL_CHARSET      = 129
	GB2312_CHARSET      = 134
	CHINESEBIG5_CHARSET = 136
	OEM_CHARSET         = 255
	JOHAB_CHARSET       = 130
	HEBREW_CHARSET      = 177
	ARABIC_CHARSET      = 178
	GREEK_CHARSET       = 161
	TURKISH_CHARSET     = 162
	VIETNAMESE_CHARSET  = 163
	THAI_CHARSET        = 222
	EASTEUROPE_CHARSET  = 238
	RUSSIAN_CHARSET     = 204
	MAC_CHARSET         = 77
	BALTIC_CHARSET      = 186

	FS_LATIN1      = 0x00000001
	FS_LATIN2      = 0x00000002
	FS_CYRILLIC    = 0x00000004
	FS_GREEK       = 0x00000008
	FS_TURKISH     = 0x00000010
	FS_HEBREW      = 0x00000020
	FS_ARABIC      = 0x00000040
	FS_BALTIC      = 0x00000080
	FS_VIETNAMESE  = 0x00000100
	FS_THAI        = 0x00010000
	FS_JISJAPAN    = 0x00020000
	FS_CHINESESIMP = 0x00040000
	FS_WANSUNG     = 0x00080000
	FS_CHINESETRAD = 0x00100000
	FS_JOHAB       = 0x00200000
	FS_SYMBOL      = 0x80000000

	FF_DONTCARE   = 0x00
	FF_ROMAN      = 0x10
	FF_SWISS      = 0x20
	FF_MODERN     = 0x30
	FF_SCRIPT     = 0x40
	FF_DECORATIVE = 0x50
)

type (
	ATOM       uint16
	HANDLE     uintptr
	HINSTANCE  HANDLE
	HICON      HANDLE
	HCURSOR    HANDLE
	HBRUSH     HANDLE
	HWND       HANDLE
	HMENU      HANDLE
	HDC        HANDLE
	HFONT      HANDLE
	HGDIOBJECT HANDLE
	BOOL       int32
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
	CreateWindowEx  func(exStyle uint, className, windowName *uint16,
		style uint, x, y, width, height int, parent HWND, menu HMENU,
		instanceance HINSTANCE, param unsafe.Pointer) HWND
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
	DrawText         func(hdc HDC, text *uint16, c int32, rect *RECT, format uint32) bool
	LoadCursor       func(instanceance HINSTANCE, cursorName *uint16) HCURSOR
	SetCursor        func(cursor HCURSOR)

	Rectangle func(hdc HDC, left, right, top, bottom int32) bool
	Ellipse   func(hdc HDC, left, right, top, bottom int32) bool
	SetBkMode func(hdc HDC, mode uint32) int
	TextOut   func(hdc HDC, x, y int32, text *uint16, c int32) bool

	CreateFont func(
		height, width, escapement, orientation, weight int32,
		italic, underline, strikeout uint32,
		charset, outPrecisioin, clipPrecision, quality, pitchandfamily uint32,
		facename *uint16,
	) HFONT
	SelectObject func(hdc HDC, handle HGDIOBJECT) HGDIOBJECT
	DeleteObject func(handle HGDIOBJECT) BOOL
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
	purego.RegisterLibFunc(&DrawText, user32, "DrawTextW")
	purego.RegisterLibFunc(&SetCursor, user32, "SetCursor")
	purego.RegisterLibFunc(&LoadCursor, user32, "LoadCursorW")

	gdi32 := windows.NewLazySystemDLL("gdi32.dll").Handle()
	purego.RegisterLibFunc(&Rectangle, gdi32, "Rectangle")
	purego.RegisterLibFunc(&Ellipse, gdi32, "Ellipse")
	purego.RegisterLibFunc(&SetBkMode, gdi32, "SetBkMode")
	purego.RegisterLibFunc(&TextOut, gdi32, "TextOutW")

	purego.RegisterLibFunc(&CreateFont, gdi32, "CreateFontW")
	purego.RegisterLibFunc(&SelectObject, gdi32, "SelectObject")
	purego.RegisterLibFunc(&DeleteObject, gdi32, "DeleteObject")

	runtime.LockOSThread()
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

	title, err := windows.UTF16PtrFromString("purego 002-Graphics")
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

func wndProc(hwnd HWND, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {

	case WM_PAINT:
		ps := PAINTSTRUCT{}
		hdc := BeginPaint(hwnd, &ps)

		// 1. 背景の塗りつぶし
		clientRect := RECT{}
		GetClientRect(hwnd, &clientRect)
		FillRect(hdc, &clientRect, (HBRUSH)(COLOR_WINDOW+1))

		// 2. 四角形の描画
		Rectangle(hdc, 50, 50, 200, 150)

		// 3. 楕円の描画
		Ellipse(hdc, 100, 100, 250, 200)

		// 4. テキストの描画

		// フォントを作成（例: Noto Sans JP、サイズ20px）
		fontName, err := windows.UTF16PtrFromString("Noto Sans JP")
		if err != nil {
			panic(err)
		}
		font := CreateFont(
			32,                        // フォントの高さ (px)
			0,                         // 幅（0で自動調整）
			0,                         // エスケープメント角度
			0,                         // ベースライン角度
			FW_BOLD,                   // 太さ (FW_NORMAL, FW_BOLD など)
			FALSE,                     // 斜体 (TRUE / FALSE)
			FALSE,                     // 下線
			FALSE,                     // 打消線
			SHIFTJIS_CHARSET,          // 文字セット（日本語環境）
			OUT_DEFAULT_PRECIS,        // 出力精度
			CLIP_DEFAULT_PRECIS,       // クリッピング精度
			DEFAULT_QUALITY,           // 描画品質
			DEFAULT_PITCH|FF_DONTCARE, // ピッチとファミリー
			fontName,                  // フォント名
		)

		//  作成したフォントをデバイスコンテキスト(hdc)に適用し、元のフォントを保存
		oldFont := HFONT(SelectObject(hdc, HGDIOBJECT(font)))

		// 背景透過を設定（テキスト周囲の白背景化を防ぐ）
		SetBkMode(hdc, TRANSPARENT)

		// 文字列を *uint16 に変換し、末尾に 0x0000 を入れる
		text := "Hello world\nこんにちは世界\n"
		textU16, err := windows.UTF16PtrFromString(text)
		if err != nil {
			panic(err)
		}

		// テキストを出力する
		DrawText(hdc,
			textU16,
			-1,	// 文字列長を自動計算
			&RECT{Left: 70, Top: 100}, // DT_NOCLIP の場合、Right と Bottom は不用
			DT_LEFT|DT_TOP|DT_NOCLIP,
		)

		// 元のフォントに戻し、不要になったフォントを破棄する
		SelectObject(hdc, HGDIOBJECT(oldFont))
		DeleteObject(HGDIOBJECT(font))

		// 描画終了処理
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
