package cv_test

import (
	"bytes"
	"fmt"
	// "os"
	"testing"

	"image"
	// "image/draw"
	_ "image/png"

	_ "embed"

	"github.com/117503445/gorobot-demo/internal/cv"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

//go:embed test-input/1.png
var fileImg1 []byte

func TestImage(t *testing.T) {

	fmt.Println("TestImage")
	fmt.Println(len(fileImg1))

	ast := assert.New(t)

	// img1 := image.
	img, format, err := image.Decode(bytes.NewReader(fileImg1))
	ast.NoError(err)
	log.Debug().Str("format", format).Msg("image.Decode")

	rgba := image.NewRGBA(img.Bounds())
	pixel := rgba.At(1, 1)
	r, g, b, a := pixel.RGBA()
	log.Debug().Int("r", int(r)).Int("g", int(g)).Int("b", int(b)).Int("a", int(a)).Msg("pixel")
	points := cv.GetUnlockedPoints(img)

	log.Debug().Interface("points", points).Msg("GetUnlockedPoints")

	// // 创建输出文件
	// outFile, err := os.Create("output.png")
	// ast.NoError(err)
	// defer outFile.Close()

	// // 编码并保存图片
	// err = png.Encode(outFile, rgba)
	// ast.NoError(err)

}
