package main

import (
	"fmt"
	// "image"
	"os"
	"time"

	"github.com/117503445/gorobot-demo/internal/cv"
	"github.com/117503445/goutils"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/rs/zerolog/log"
)

var logsDir string

func AltWCallback(e hook.Event) {
	go func() {
		c := 0
		for {

			robotgo.MoveClick(1472/2, 584/2)

			// 轮换遗器槽
			points := cv.GetRelicPoints()
			for _, p := range points {
				robotgo.Move(p.X/2, p.Y/2)
				robotgo.Click()

				time.Sleep(1 * time.Second)

				img := robotgo.CaptureImg()

				imgFile := fmt.Sprintf("%s/%s.png", logsDir, time.Now().Format("20060102.150405"))
				if err := robotgo.Save(img, imgFile); err != nil {
					log.Fatal().Err(err).Msg("robotgo.Save")
				}

				log.Debug().Str("imgFile", imgFile).Msg("CaptureImg")

				// 选中当前未选中的遗器
				points := cv.GetUnlockedPoints(img)
				log.Debug().Interface("points", points).Msg("GetUnlockedPoints")
				for _, p := range points {
					robotgo.Move(p.X/2, p.Y/2)
					robotgo.Click()
					robotgo.Move(3726/2, 394/2)
					robotgo.Click()
				}
			}

			robotgo.KeyTap("esc")
			c += 1
			robotgo.MoveClick(129/2, 2008/2)
			if c%4 == 0 {
				robotgo.KeyTap("a")
				robotgo.KeyTap("a")
				robotgo.KeyTap("a")
				robotgo.KeyTap("s")
			} else {
				robotgo.KeyTap("d")
			}
			robotgo.MoveClick(1472/2, 584/2)
		}
	}()
}

func AltFCallback(e hook.Event) {
	// 取消选择所有仪器
	go func() {
		robotgo.Move(3624/2, 546/2)
		for {
			// time.Sleep(3 * time.Second)

			robotgo.Click()
			// robotgo.KeyTap("d")
			// time.Sleep(3 * time.Second)
		}
	}()
}

func main() {
	goutils.InitZeroLog()
	// runID as a unique identifier for the current run, filename
	// 20240803.203942
	runID := time.Now().Format("20060102.150405")

	logsDir = fmt.Sprintf("logs/%s", runID)
	if err := os.MkdirAll(logsDir, os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("os.MkdirAll")
	}

	robotgo.MouseSleep = 200
	robotgo.KeySleep = 200

	hook.Register(hook.KeyDown, []string{"alt", "a"}, func(e hook.Event) {
		os.Exit(1)
		hook.End()
		// exit(0)
	})

	// c := 0

	fmt.Println("--- Please press alt+w---")
	hook.Register(hook.KeyDown, []string{"alt", "w"}, AltWCallback)
	hook.Register(hook.KeyDown, []string{"alt", "f"}, AltFCallback)

	s := hook.Start()
	<-hook.Process(s)

	// fmt.Println("Press enter to exit")
	// // python input() to prevent the program from exiting immediately
	// fmt.Scanln()
}
