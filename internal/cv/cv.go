package cv

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path"

	"github.com/rs/zerolog/log"
)

// CVHelper 用于处理图像识别相关的操作
type CVHelper struct {
	img    image.Image
	logDir string
}

// NewCVHelper 创建一个 CVHelper 实例
func NewCVHelper(img image.Image, logDir string) *CVHelper {
	if img == nil {
		log.Fatal().Msg("img is nil")
	}
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err = os.Mkdir(logDir, os.ModePerm); err != nil {
			log.Fatal().Err(err).Msg("os.Mkdir")
		}
	}
	return &CVHelper{
		img:    img,
		logDir: logDir,
	}
}

// DrawRect 在给定的图像上绘制一个颜色为 col 的矩形边框，并指定边框厚度
func DrawRect(rgba *image.RGBA, rect image.Rectangle, col color.Color, thickness int) {
	// 上边
	for t := 0; t < thickness; t++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			rgba.Set(x, rect.Min.Y+t, col)
		}
	}
	// 下边
	for t := 0; t < thickness; t++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			rgba.Set(x, rect.Max.Y-1-t, col)
		}
	}
	// 左边
	for t := 0; t < thickness; t++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			rgba.Set(rect.Min.X+t, y, col)
		}
	}
	// 右边
	for t := 0; t < thickness; t++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			rgba.Set(rect.Max.X-1-t, y, col)
		}
	}
}

func (cvh *CVHelper) width() int {
	return cvh.img.Bounds().Max.X
}

func (cvh *CVHelper) height() int {
	return cvh.img.Bounds().Max.Y
}

func (cvh *CVHelper) tranXFrom4KTo1080p(x int) int {
	return x * cvh.width() / 3840
}

func (cvh *CVHelper) tranYFrom4KTo1080p(y int) int {
	return y * cvh.height() / 2160
}

// getRect: 获取指定行列遗器锁定按钮的矩形区域
// row: 第几行
// column: 第几列
func (cvh *CVHelper) getRect(row int, column int) image.Rectangle {
	x := 258 + 250*(column-1)
	x = cvh.tranXFrom4KTo1080p(x)

	y := 551 + (849-551)*(row-1)
	y = cvh.tranYFrom4KTo1080p(y)

	w := 38 + 5
	w = cvh.tranXFrom4KTo1080p(w)
	h := 43 + 5
	h = cvh.tranYFrom4KTo1080p(h)

	return image.Rect(x, y, x+w, y+h)
}

// 获取未锁定的 point 列表，返回 1080p 分辨率下的绝对坐标
func (cvh *CVHelper) GetUnlockedPoints() []image.Point {
	log.Debug().Msg("GetUnlockedPoints")

	// log.Debug().Interface("color mode", img.ColorModel()).Msg("color mode")
	rgba := image.NewRGBA(cvh.img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), cvh.img, image.Point{}, draw.Src)
	points := []image.Point{}

	// 在 4k 分辨率下，200% 缩放，黑色 白色 一般是 460, 200 左右
	// 这里设置 100 作为阈值
	threshold := 100 * cvh.width() / 3840 * cvh.height() / 2160

	for r := 1; r <= 5; r++ {
		for c := 1; c <= 4; c++ {
			rect := cvh.getRect(r, c)

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

			// if blackCount > threshold && whiteCount > threshold {
			if blackCount > threshold {
				DrawRect(rgba, rect, color.RGBA{R: 255, G: 0, B: 0, A: 255}, 5)

				// locked
			} else {
				DrawRect(rgba, rect, color.RGBA{R: 0, G: 0, B: 255, A: 255}, 5)
				// points = append(points, image.Point{X: rect.Min.X, Y: rect.Min.Y})

				// point := image.Point{X: cvh.tranXFrom4KTo1080p(rect.Min.X), Y: cvh.tranYFrom4KTo1080p(rect.Min.Y)}

				x, y := rect.Min.X, rect.Min.Y
				// 放缩到 1080p, 以便 robotgo 使用
				x, y = x*1920/cvh.width(), y*1080/cvh.height()

				point := image.Point{X: x, Y: y}
				points = append(points, point)
				log.Debug().Interface("point", point).Msg("GetUnlockedPoints")
			}
		}
	}

	rbgaToFile(rgba, path.Join(cvh.logDir, "GetUnlockedPoints.png"))

	return points
}

func rbgaToFile(img *image.RGBA, file string) {
	outFile, err := os.Create(file)
	if err != nil {
		log.Fatal().Err(err).Msg("os.Create")
	}
	defer outFile.Close()

	// 编码并保存图片
	if err = png.Encode(outFile, img); err != nil {
		log.Fatal().Err(err).Msg("png.Encode")
	}
}

// 获取遗器槽的坐标, 1080p 分辨率下的绝对坐标
func (cvh *CVHelper) GetRelicPoints() []image.Point {
	points := []image.Point{}
	// 4K: (269,280), (406,288)

	for i := 0; i < 6; i++ {
		x, y := 269+137*i, 280
		// x, y = cvh.tranXFrom4KTo1080p(x), cvh.tranYFrom4KTo1080p(y)
		x, y = x/2, y/2
		points = append(points, image.Point{X: x, Y: y})
	}

	log.Debug().Interface("points", points).Msg("GetRelicPoints")

	return points
}
