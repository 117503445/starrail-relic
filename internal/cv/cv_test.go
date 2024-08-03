package cv_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"image"
	"image/color"
	"image/draw"
	"image/png"

	_ "embed"

	"github.com/117503445/gorobot-demo/internal/cv"
	"github.com/stretchr/testify/assert"
)

//go:embed test-input/1.png
var fileImg1 []byte

func TestImage(t *testing.T) {

	fmt.Println("TestImage")
	fmt.Println(len(fileImg1))

	ast := assert.New(t)

	// img1 := image.
	img, _, err := image.Decode(bytes.NewReader(fileImg1))
	ast.NoError(err)

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	// 定义矩形的颜色和位置
	red := color.RGBA{255, 0, 0, 255}
	// rect := image.Rect(258, 551, 296, 594) // 这里定义了矩形的起始点和结束点
	for r := 1; r <= 5; r++ {
		for c := 1; c <= 4; c++ {
			rect := cv.DrawR(r, c)
			cv.DrawRect(rgba, rect, red, 5)

			blackCount := 0
			whiteCount := 0
			for x := rect.Min.X; x < rect.Max.X; x++ {
				for y := rect.Min.Y; y < rect.Max.Y; y++ {
					pixel := rgba.At(x, y)
					r, g, b, a := pixel.RGBA()
					// fmt.Printf("r: %d, g: %d, b: %d, a: %d\n", r, g, b, a)

					if r*255/a <= 20 && g*255/a <= 20 && b*255/a <= 20 {
						blackCount++
					}
					if r*255/a >= 235 && g*255/a >= 235 && b*255/a >= 235 {
						whiteCount++
					}
				}
			}
			fmt.Printf("row: %d, column: %d, blackCount: %d, whiteCount: %d\n", r, c, blackCount, whiteCount)
		}
	}

	// 创建输出文件
	outFile, err := os.Create("output.png")
	ast.NoError(err)
	defer outFile.Close()

	// 编码并保存图片
	err = png.Encode(outFile, rgba)
	ast.NoError(err)

}
