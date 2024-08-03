package image_test

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

	"github.com/stretchr/testify/assert"
)

//go:embed test-input/1.png
var fileImg1 []byte

// drawRect 在给定的图像上绘制一个颜色为 col 的矩形边框，并指定边框厚度
func drawRect(img *image.RGBA, rect image.Rectangle, col color.Color, thickness int) {
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
	rect := image.Rect(50, 50, 200, 200) // 这里定义了矩形的起始点和结束点

	// 绘制矩形边框
	drawRect(rgba, rect, red, 5)

	// 创建输出文件
	outFile, err := os.Create("output.png")
	ast.NoError(err)
	defer outFile.Close()

	// 编码并保存图片
	err = png.Encode(outFile, rgba)
	ast.NoError(err)

}
