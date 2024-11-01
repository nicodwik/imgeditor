
# Image Editor: Add Text, Customize, and Enhance

Easily add text to your images with this powerful and flexible Go package.

## Features

### function `generateText(param Param)`
Generate text overlay in your image source

### struct `Param` 
- `FontSize (int)`
set font size
- `FontType ([]byte)` 
set font type, you can embed custom TTF file using `go embed`
- `PosX (int)`
set starting pixel point in horizontal scale
- `PosY (int)`
same with `PosX` but in vertical way
- `NewLineBorder (int)`
set the limit of pixel `PosX` to make new break line
- `Text (string)`
set text you want to be printed on image

### function `writeToFile(format string)`
Write edited image to file bytes





## Installation

```bash
  go get github.com/yourusername/go-image-text
```
    
## Usage/Examples

```
package main

import (
	"bytes"
	_ "embed"
	"fmt"
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

	// generate song singer
	_, lastYPos, _ := ie.GenerateText(
		&imgeditor.Param{
			FontSize:      24,
			FontType:      circularBlack,
			PosX:          64,
			PosY:          200,
			NewLineBorder: 516,
			Text:          "JMK48",
		})

	// generate song title
	ie.GenerateText(
		&imgeditor.Param{
			FontSize:      17,
			FontType:      circularBold,
			PosX:          64,
			PosY:          lastYPos + 50,
			NewLineBorder: 516,
			Text:          "Fortune Cookie Yang Mencinta Fortune Cookie Yang Mencinta Fortune Cookie Yang Mencinta",
		})

	// write edited image to file bytes
	fileBytes, err := ie.WriteToFile("png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileBytes)

	fmt.Println("Edit image success!")
}

```

## Output
<img width="701" alt="image" src="https://github.com/user-attachments/assets/fc9ec372-ba87-43a2-9b34-ec39cbc08357">

## Feedback

If you have any feedback, please reach out to us at nicodwika@gmail.com

