package main

import (
	"fmt"
	"os"
	"time"

	_ "embed"
	"sync/atomic"

	"github.com/117503445/goutils"
	"github.com/117503445/starrail-relic/internal/cv"
	"github.com/117503445/starrail-relic/internal/lowos"
	"github.com/gen2brain/beeep"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/rs/zerolog/log"
)

var logsDir string

//go:embed version.json
var versionInfo []byte

var isRunning atomic.Bool

func notify(message string) {
	if err := beeep.Notify("starrail-relic", message, ""); err != nil {
		log.Fatal().Err(err).Msg("beeep.Notify")
	}
}

func AltWCallback(e hook.Event) {
	if !isRunning.CompareAndSwap(false, true) {
		notify("任务进行中，无法执行其他任务")
		return
	}

	go func() {
		notify("锁定每个角色前 20 个推荐的遗器")

		c := 0
		for {
			robotgo.MoveClick(1472/2, 584/2)

			img := robotgo.CaptureImg()

			// 轮换遗器槽
			points := cv.GetRelicPoints(img)
			for _, p := range points {
				robotgo.Move(p.X, p.Y)
				robotgo.Click()

				// 等待加载
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
					robotgo.Move(p.X, p.Y)
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
	if !isRunning.CompareAndSwap(false, true) {
		notify("任务进行中，无法执行其他任务")
		return
	}

	// 取消选择所有仪器
	go func() {
		notify("解锁所有遗器")

		robotgo.Move(3624/2, 546/2)
		for {
			robotgo.Click()
		}
	}()
}

func main() {
	goutils.InitZeroLog()

	log.Info().RawJSON("versionInfo", versionInfo).Msg("")

	if !lowos.IsAdmin() {
		log.Error().Msg("请以管理员身份运行")
		fmt.Scanln()
		os.Exit(1)
	}

	// runID as a unique identifier for the current run, example: 20240803.203942
	runID := time.Now().Format("20060102.150405")

	logsDir = fmt.Sprintf("logs/%s", runID)
	if err := os.MkdirAll(logsDir, os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("os.MkdirAll")
	}

	fmt.Println(`使用方法:
	alt + a: 退出程序
	alt + f: 解锁所有遗器。请先打开星穹铁道，进入 背包 - 遗器 页面，筛选 - 状态 - 已锁定，再按下 alt + f 快捷键。当所有遗器都被解锁后，按下 alt + a 退出程序。
	alt + w: 锁定每个角色前 20 个推荐的遗器。请先打开星穹铁道，进入 角色详情 - 第一个角色 - 遗器 页面，再按下 alt + w 快捷键。当所有遗器都被锁定后，按下 alt + a 退出程序。
	`)

	robotgo.MouseSleep = 200
	robotgo.KeySleep = 200

	hook.Register(hook.KeyDown, []string{"alt", "a"}, func(e hook.Event) {
		notify("退出程序")

		// TODO: use channel
		os.Exit(1)
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"alt", "w"}, AltWCallback)
	hook.Register(hook.KeyDown, []string{"alt", "f"}, AltFCallback)

	s := hook.Start()
	<-hook.Process(s)
}
