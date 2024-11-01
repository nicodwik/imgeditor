package imgeditor

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Param struct {
	FontSize       int
	FontType       []byte
	PosX           int
	PosY           int
	NewLineBorderX int
	NewLineBorderY int
	Text           string
}

type Object struct {
	ImageSrc    []byte
	ImageOutput *image.RGBA
}

func New(imageSrc []byte) *Object {
	return &Object{
		ImageSrc: imageSrc,
	}
}

func (ie *Object) GenerateText(param *Param) (lastXPos, lastYPos int, err error) {
	if len(param.Text) == 0 {
		return 0, 0, errors.New("Text can't be empty")
	}

	var imageInput image.Image
	if ie.ImageOutput != nil {
		imageInput = ie.ImageOutput
	} else {
		// Decode the image
		img, _, err := image.Decode(bytes.NewReader(ie.ImageSrc))
		if err != nil {
			return 0, 0, err
		}

		imageInput = img
	}

	// Create a new image with the same dimensions as the original
	fakeImg := image.NewRGBA(imageInput.Bounds())
	newImg := image.NewRGBA(imageInput.Bounds())
	draw.Draw(newImg, newImg.Rect, imageInput, image.Point{0, 0}, draw.Src)

	// Parse the font
	font, err := truetype.Parse(param.FontType)
	if err != nil {
		return 0, 0, err
	}

	ctx := createFontContextWithOptions(newImg, font, param)
	fakeCtx := createFontContextWithOptions(fakeImg, font, param)

	lastXPos = param.PosX
	lastYPos = param.PosY

	// split text to slice of words
	words := strings.Fields(param.Text)
	for _, word := range words {
		fakePos := freetype.Pt(lastXPos, lastYPos)
		fakeP, _ := fakeCtx.DrawString(word+" ", fakePos)

		// if drawed string more than new line border X, make a new line
		if fakeP.X.Round() >= param.NewLineBorderX {
			lastXPos = param.PosX
			lastYPos += 38
		}

		// if drawed string more than new line border Y, do nothing
		if fakeP.Y.Round() >= param.NewLineBorderY {
			continue
		}

		pos := freetype.Pt(lastXPos, lastYPos)
		p, _ := ctx.DrawString(word+" ", pos)

		lastXPos += (p.X.Round() - pos.X.Round())
	}

	ie.ImageOutput = newImg

	return lastXPos, lastYPos, nil
}

func (ie *Object) WriteToFile(format string) ([]byte, error) {

	if ie.ImageOutput == nil {
		return []byte(""), errors.New("Image Output can't be empty")
	}

	// Save the overwritten image
	buff := new(bytes.Buffer)

	// Encode the image to the desired format
	switch format {
	case "png":
		err := png.Encode(buff, ie.ImageOutput)
		if err != nil {
			return []byte(""), err
		}

	case "jpeg":
		err := jpeg.Encode(buff, ie.ImageOutput, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			return []byte(""), err
		}
	default:
		return []byte(""), errors.New("Unknown Format")
	}

	return buff.Bytes(), nil
}

func createFontContextWithOptions(img *image.RGBA, font *truetype.Font, param *Param) *freetype.Context {
	ctx := freetype.NewContext()

	ctx.SetDPI(150)
	ctx.SetDst(img)
	ctx.SetFont(font)
	ctx.SetFontSize(float64(param.FontSize))
	ctx.SetSrc(image.Black)
	ctx.SetClip(img.Bounds())

	return ctx
}
