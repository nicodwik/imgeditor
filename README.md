
# Image Editor: Add Text, Customize, and Enhance

Easily add text to your images with this powerful and flexible Go package.

## Features

### function `HexToRGBA(hex string)`
Convert hex to RGBA struct

### function `GenerateText(param Param)`
Generate text overlay in your image source

### struct `Param` 
- `FontSize (int)` => set font size
- `FontType ([]byte)` => set font type, you can embed custom TTF file using `go embed`
- `FontColor (color.Color)` => set font color, you can use `color.Black` or define custom color with `HexToRGBA()` function
- `PosX (int)` => set starting pixel point in horizontal scale
- `PosY (int)` => same with `PosX` but in vertical way
- `NewLineBorderX (int)` => set the limit of pixel `PosX` to make new break line
- `NewLineBorderY (int)` => set the limit of pixel `PosY` to skip print text vertically
- `Text (string)` => set text you want to be printed on image

### function `WriteToFile(format string)`
Write edited image to file bytes

## Installation

```bash
  go get github.com/nicodwik/imgeditor
```
    
## Usage/Examples

```
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/color"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/nicodwik/imgeditor"
)

//go:embed image.png
var imageFile []byte

//go:embed circular-std-black.ttf
var circularBlack []byte

//go:embed circular-std-bold.ttf
var circularBold []byte

func main() {

	ie := imgeditor.New(imageFile)

	// generate custom color, if needed
	colorBlue := imgeditor.HexToRGBA("#0000FF")
	fmt.Println(colorBlue)

	// generate song singer
	_, lastYPos, _ := ie.GenerateText(
		&imgeditor.Param{
			FontSize:      24,
			FontType:      circularBlack,
			FontColor:	color.Black,
			PosX:          64,
			PosY:          200,
			LineHeight: 38,
			NewLineBorderX: 516,
			NewLineBorderY: 430,
			Text:          "JMK48",
		})

	// generate song title
	ie.GenerateText(
		&imgeditor.Param{
			FontSize:      17,
			FontType:      circularBold,
			FontColor:	color.Black,
			PosX:          64,
			PosY:          lastYPos + 50,
			LineHeight: 38,
			NewLineBorderX: 516,
			NewLineBorderY: 430,
			Text:          "Fortune Cookie Yang Mencinta Fortune Cookie Yang Mencinta Fortune Cookie Yang Mencinta",
		})

	// write edited image to file bytes
	fileBytes, err := ie.WriteToFile("png")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("new-image.png", fileBytes, 0644)

	fmt.Println("Edit image success!")
}

```

## Output
<img width="701" alt="image" src="https://github.com/user-attachments/assets/fc9ec372-ba87-43a2-9b34-ec39cbc08357">

## Feedback

If you have any feedback, please reach out to us at nicodwika@gmail.com

