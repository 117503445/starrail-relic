package cv

import (
	"image"
	"image/color"
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

// row: 第几行
// column: 第几列
func DrawR(row int, column int) image.Rectangle {
	x := 258 + 250*(column-1)
	y := 551 + (849-551)*(row-1)
	return image.Rect(x, y, x+38, y+43)
}

// 获取未锁定的 point 列表
func GetUnlockedPoints(img image.Image) []image.Point {
	rgba := image.NewRGBA(img.Bounds())
	points := []image.Point{}

	for r := 1; r <= 5; r++ {
		for c := 1; c <= 4; c++ {
			rect := DrawR(r, c)
			blackCount := 0
			whiteCount := 0
			for x := rect.Min.X; x < rect.Max.X; x++ {
				for y := rect.Min.Y; y < rect.Max.Y; y++ {

					pixel := rgba.At(x, y)
					r, g, b, a := pixel.RGBA()
					if r*255/a <= 20 && g*255/a <= 20 && b*255/a <= 20 {
						blackCount++
					}
					if r*255/a >= 235 && g*255/a >= 235 && b*255/a >= 235 {
						whiteCount++
					}

				}
			}
			if blackCount > 200 && whiteCount > 100 {
				// locked
			} else {
				points = append(points, image.Point{X: rect.Min.X, Y: rect.Min.Y})
			}
		}
	}

	return points
}
