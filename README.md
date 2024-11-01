
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

	// send image as multipart/form-data
	err = httpClientWithMultipart(fileBytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Send image success!")
}

func httpClientWithMultipart(fileBytes []byte) error {
	buf := &bytes.Buffer{}
	multipart := multipart.NewWriter(buf)

	wri, err := multipart.CreateFormFile("image", "jancuk.jpg")
	if err != nil {
		return err
	}

	read := bytes.NewReader(fileBytes)
	_, err = io.Copy(wri, read)
	if err != nil {
		return err
	}
	multipart.Close()

	req, err := http.NewRequest("POST", "https://99f3-202-6-237-2.ngrok-free.app/api/test", buf)
	req.Header.Add("Content-Type", multipart.FormDataContentType())
	if err != nil {
		return err
	}
	client := http.Client{}
	client.Do(req)

	return nil
}

```

## Output

![Screenshot](https://nicodwik.github.io/totp-generator/screenshot.png)


## Feedback

If you have any feedback, please reach out to us at nicodwika@gmail.com

