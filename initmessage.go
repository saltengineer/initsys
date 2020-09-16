package initsys
// Original code by Salt Engineer, 2016 (systatus.c) >> systatus.hpp >> initmessage.go
// Translated to go on 12/28/2019

import (
	"flag"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"runtime"
	"github.com/logrusorgru/aurora" // For terminal colors
	"github.com/nsf/termbox-go" // For terminal height/width
)

// Global Vars
// Store terminal width, height
var w, h int
var statsize = 8 // Size of status block
//var colors = flag.Bool("colors", true, "Enable color; Needs HKCU/Console/VirtualTerminalLevel (dword) set to 1 on Windows")
var colors bool

var au aurora.Aurora

func init() {
	if runtime.GOOS == "windows" {
		Console, err := registry.OpenKey(registry.CURRENT_USER, "Console", registry.ALL_ACCESS)
		if err != nil {
			flag.BoolVar(&colors, "colors", false, "Enable color; Needs HKCU/Console/VirtualTerminalLevel (dword) set to 1 on Windows")

		} else {
			defer Console.Close()
			//var key []byte
			colorKey, _, err := Console.GetIntegerValue("VirtualTerminalLevel")
			if err != nil {
				flag.BoolVar(&colors, "colors", false, "Enable color; Needs HKCU/Console/VirtualTerminalLevel (dword) set to 1 on Windows")
			} else {
				if colorKey == 1 {
					flag.BoolVar(&colors, "colors", true, "Enable color; Needs HKCU/Console/VirtualTerminalLevel (dword) set to 1 on Windows")
				} else {
					flag.BoolVar(&colors, "colors", false, "Enable color; Needs HKCU/Console/VirtualTerminalLevel (dword) set to 1 on Windows")
				}
			}
		}

	} else {
		fmt.Printf("OS = %s\n", runtime.GOOS)
		flag.BoolVar(&colors, "colors", true, "Enable color; Needs HKCU/Console/VirtualTerminalLevel (dword) set to 1 on Windows")
	}
	{	// Other flags
		//envPtr = nil
		//configPtr = nil
		//if envPtr == nil || configPtr == nil { } // This line intentionally left blank
	}
	//flag.Parse()
	au = aurora.NewAurora(colors)
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, h = termbox.Size()
	termbox.Close()
}

func StatusMessage(message string, status string) string {
	pad := len(message) + statsize
	out := message
	fmt.Printf("%s", message)
	padn := 1
	for padn < (w - pad + 1) {
		out += " "
		fmt.Printf(" ")
		padn++
	}
	out += "[ "
	fmt.Printf("[ ")
	switch stat := status; {
	case stat == "PASS" || stat == "pass":
		fmt.Printf("%s", au.Green("PASS"))
		out += "PASS"
	case stat == "DONE" || stat == "done":
		fmt.Printf("%s", au.Green("DONE"))
		out += "DONE"
	case stat == "FAIL" || stat == "fail":
		fmt.Printf("%s", au.Red(au.SlowBlink(au.Bold("FAIL"))))
		out += "FAIL"
	case stat == "WARN" || stat == "warn":
		fmt.Printf("%s", au.Yellow("WARN"))
		out += "WARN"
	}
	out += " ]"
	fmt.Printf(" ]\n")
	return out
}

func GetConsoleW() int {
	return w
}

