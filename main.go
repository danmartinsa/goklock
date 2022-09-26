package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/danmartinsa/klock/lib"
	// "github.com/rivo/tview"
)

func main (){
	// TODO implement in go channel
	// // app := tview.NewApplication()
	// box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	// hour := tview.PrintSimple()
	//
	// if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	// 	panic(err)
	// }
	tick := true
	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
 		fmt.Println(lib.PrintTime(time.Now(), tick))
 		if tick {
 			tick = false
 		} else {
 			tick = true
 		}
 		time.Sleep(time.Second)
	}
}
