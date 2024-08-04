package cv

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/rs/zerolog/log"
)

// DrawRect 在给定的图像上绘制一个颜色为 col 的矩形边框，并指定边框厚度
func DrawRect(img *image.RGBA, rect image.Rectangle, col color.Color, thickness int) {
	// 上边
	for t := 0; t < thickness; t++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			img.Set(x, rect.Min.Y+t, col)
		}
	}
	// 下边
	for t := 0; t < thickness; t++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			img.Set(x, rect.Max.Y-1-t, col)
		}
	}
	// 左边
	for t := 0; t < thickness; t++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			img.Set(rect.Min.X+t, y, col)
		}
	}
	// 右边
	for t := 0; t < thickness; t++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			img.Set(rect.Max.X-1-t, y, col)
		}
	}
}

func tranXFrom4KTo1080p(img image.Image, x int) int {
	return x * img.Bounds().Max.X / 3840
}

func tranYFrom4KTo1080p(img image.Image, y int) int {
	return y * img.Bounds().Max.Y / 2160
}

// getRect: 获取指定行列遗器锁定按钮的矩形区域，1080p 分辨率下的绝对坐标
// row: 第几行
// column: 第几列
func getRect(img image.Image, row int, column int) image.Rectangle {
	x := 258 + 250*(column-1)
	log.Debug().Int("x", x).Msg("getRect")
	x = tranXFrom4KTo1080p(img, x)
	log.Debug().Int("x", x).Msg("getRect after tran")

	y := 551 + (849-551)*(row-1)
	y = tranYFrom4KTo1080p(img, y)

	w := 38
	w = tranXFrom4KTo1080p(img, w)
	h := 43
	h = tranYFrom4KTo1080p(img, h)

	return image.Rect(x, y, x+w, y+h)
}

// 获取未锁定的 point 列表，返回 1080p 分辨率下的绝对坐标
func GetUnlockedPoints(img image.Image) []image.Point {
	if img == nil {
		log.Fatal().Msg("img is nil")
	}
	log.Debug().Int("width", img.Bounds().Max.X).Int("height", img.Bounds().Max.Y).Msg("GetUnlockedPoints")

	// log.Debug().Interface("color mode", img.ColorModel()).Msg("color mode")
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
	points := []image.Point{}

	// 在 4k 分辨率下，200% 缩放，黑色 白色 一般是 600 左右
	// 这里设置 300 作为阈值
	threshold := 300 * img.Bounds().Max.X / 3840 * img.Bounds().Max.Y / 2160

	for r := 1; r <= 5; r++ {
		for c := 1; c <= 4; c++ {
			rect := getRect(img, r, c)

			blackCount := 0
			whiteCount := 0
			for x := rect.Min.X; x < rect.Max.X; x++ {
				for y := rect.Min.Y; y < rect.Max.Y; y++ {
					pixel := rgba.At(x, y)
					r, g, b, a := pixel.RGBA()
					if a == 0 {
						log.Fatal().Int("x", x).Int("y", y).Msg("a == 0")
					}

					// log.Debug().Int("r", int(r)).Int("g", int(g)).Int("b", int(b)).Int("a", int(a)).Msg("pixel")

					if r*255/a <= 20 && g*255/a <= 20 && b*255/a <= 20 {
						blackCount++
					}
					if r*255/a >= 235 && g*255/a >= 235 && b*255/a >= 235 {
						whiteCount++
					}

				}
			}
			log.Debug().Int("blackCount", blackCount).Int("whiteCount", whiteCount).Int("threshold", threshold).Msg("count")

			if blackCount > threshold && whiteCount > threshold {
				// locked
			} else {
				// points = append(points, image.Point{X: rect.Min.X, Y: rect.Min.Y})
				point := image.Point{X: tranXFrom4KTo1080p(img, rect.Min.X), Y: tranYFrom4KTo1080p(img, rect.Min.Y)}
				points = append(points, point)
				log.Debug().Interface("point", point).Msg("GetUnlockedPoints")
			}
		}
	}

	return points
}

// 获取遗器槽的坐标, 1080p 分辨率下的绝对坐标
func GetRelicPoints(img image.Image) []image.Point {
	points := []image.Point{}
	// (269,280), (406,288)

	for i := 0; i < 6; i++ {
		x, y := 269+137*i, 280
		x, y = tranXFrom4KTo1080p(img, x), tranYFrom4KTo1080p(img, y)
		points = append(points, image.Point{X: x, Y: y})
	}

	return points
}
