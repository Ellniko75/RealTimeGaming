package main

import (
	"fmt"
	"syscall"
	"unicode/utf8"

	"github.com/go-vgo/robotgo"
)

var (
	user32         = syscall.NewLazyDLL("user32.dll")
	procKeybdEvent = user32.NewProc("keybd_event")
)

const (
	KEYEVENTF_KEYUP = 0x0002
	VK_BACK         = 0x08
	VK_TAB          = 0x09
	VK_RETURN       = 0x0D
	VK_SHIFT        = 0x10
	VK_CONTROL      = 0x11
	VK_MENU         = 0x12
	VK_PAUSE        = 0x13
	VK_CAPITAL      = 0x14
	VK_ESCAPE       = 0x1B
	VK_SPACE        = 0x20
	VK_LEFT         = 0x25
	VK_UP           = 0x26
	VK_RIGHT        = 0x27
	VK_DOWN         = 0x28
	VK_W            = 0x57
	VK_A            = 0x41
	VK_S            = 0x53
	VK_D            = 0x44
)

func stringToKeyCode(key string) (byte, bool) {
	switch key {
	case "backspace":
		return VK_BACK, true
	case "tab":
		return VK_TAB, true
	case "enter":
		return VK_RETURN, true
	case "shift":
		return VK_SHIFT, true
	case "ctrl":
		return VK_CONTROL, true
	case "alt":
		return VK_MENU, true
	case "pause":
		return VK_PAUSE, true
	case "capslock":
		return VK_CAPITAL, true
	case "esc":
		return VK_ESCAPE, true
	case "8":
		return VK_SPACE, true
	case "left":
		return VK_LEFT, true
	case "up":
		return VK_UP, true
	case "right":
		return VK_RIGHT, true
	case "down":
		return VK_DOWN, true
	case "87":
		return VK_W, true
	case "65":
		return VK_A, true
	case "83":
		return VK_S, true
	case "68":
		return VK_D, true
	default:
		// Handle ASCII characters
		if len(key) == 1 {
			r, _ := utf8.DecodeRuneInString(key)
			return byte(r), true
		}
		return 0, false
	}
}

func pressKey(key byte) {
	procKeybdEvent.Call(uintptr(key), 0, 0, 0)
	procKeybdEvent.Call(uintptr(key), 0, KEYEVENTF_KEYUP, 0)
	//fmt.Print("KEY PRESSED: ", key)
}

func ExecuteKey(key string) {
	keyCode, ok := stringToKeyCode(key)
	if !ok {
		fmt.Printf("invalid key: %s", key)
	}
	pressKey(keyCode)
}

// Recieves the key as a string ("w","a","s", etc) and executes it with robotgo
func keyDownRobotgo(key string) {
	robotgo.KeyDown(key)
}
func keyUpRobotgo(key string) {
	robotgo.KeyUp(key)
}
