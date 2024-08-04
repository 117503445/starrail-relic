package cv_test

import (
	"bytes"
	"fmt"
	"testing"

	"image"
	_ "image/png"

	_ "embed"

	"github.com/117503445/starrail-relic/internal/cv"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

//go:embed test-input/1.png
var fileImg1 []byte

func TestImage(t *testing.T) {

	fmt.Println("TestImage")
	fmt.Println(len(fileImg1))

	ast := assert.New(t)

	img, format, err := image.Decode(bytes.NewReader(fileImg1))
	ast.NoError(err)
	log.Debug().Str("format", format).Msg("image.Decode")

	cvh := cv.NewCVHelper(img, "./logs")

	points := cvh.GetUnlockedPoints()

	log.Debug().Interface("points", points).Msg("GetUnlockedPoints")

}
