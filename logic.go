package imgeditor

import (
	"bytes"
	"errors"
	"image"
	"image/color"
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
	FontColor      color.Color
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

	ctx, _ := createFontContextWithOptions(newImg, param)
	fakeCtx, _ := createFontContextWithOptions(fakeImg, param)

	lastXPos = param.PosX
	lastYPos = param.PosY

	// split text to slice of words
	words := strings.Fields(param.Text)
	for _, word := range words {
		fakePos := freetype.Pt(lastXPos, lastYPos)
		fakeP, _ := fakeCtx.DrawString(word+" ", fakePos)

		newLineBorderX := newImg.Bounds().Dx()
		if param.NewLineBorderX != 0 {
			newLineBorderX = param.NewLineBorderX
		}

		// if drawed string more than new line border X, make a new line
		if fakeP.X.Round() >= newLineBorderX {
			lastXPos = param.PosX
			lastYPos += 38
		}

		newLineBorderY := newImg.Bounds().Dy()
		if param.NewLineBorderY != 0 {
			newLineBorderY = param.NewLineBorderY
		}

		// if drawed string more than new line border Y, do nothing
		if fakeP.Y.Round() >= newLineBorderY {
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

func createFontContextWithOptions(img *image.RGBA, param *Param) (*freetype.Context, error) {
	// Parse the font
	font, err := truetype.Parse(param.FontType)
	if err != nil {
		return nil, err
	}

	// Parse the color
	fontColor := image.NewUniform(param.FontColor)

	ctx := freetype.NewContext()
	ctx.SetDPI(150)
	ctx.SetDst(img)
	ctx.SetFont(font)
	ctx.SetFontSize(float64(param.FontSize))
	ctx.SetSrc(fontColor)
	ctx.SetClip(img.Bounds())

	return ctx, nil
}
