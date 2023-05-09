package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/atotto/clipboard"
	"github.com/kbinani/screenshot"
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
)

func captureScreenshot() (string, error) {
	screenRect := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(screenRect)
	if err != nil {
		return "", fmt.Errorf("error capturing screenshot: %v", err)
	}

	tempFile, err := ioutil.TempFile("", "screenshot-*.png")
	if err != nil {
		return "", fmt.Errorf("error creating temp file: %v", err)
	}

	err = png.Encode(tempFile, img)
	if err != nil {
		return "", fmt.Errorf("error encoding image: %v", err)
	}

	tempImagePath := tempFile.Name()
	tempFile.Close()

	return tempImagePath, nil
}

func detectText(imagePath string) (string, error) {
	if imagePath == "" {
		return "", nil
	}

	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return "", fmt.Errorf("error creating image annotator client: %v", err)
	}

	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("error reading image file: %v", err)
	}

	image, err := vision.NewImageFromReader(bytes.NewReader(imageData))
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return "", fmt.Errorf("error detecting text: %v", err)
	}

	if len(annotations) == 0 {
		return "", nil
	}

	return annotations[0].Description, nil
}

func copyToClipboard(text string) error {
	if text == "" {
		return nil
	}

	err := clipboard.WriteAll(text)
	if err != nil {
		return fmt.Errorf("error copying text to clipboard: %v", err)
	}

	return nil
}

func main() {
	imagePath, err := captureScreenshot()
	if err != nil {
		fmt.Printf("Error capturing screenshot: %v\n", err)
		return
	}

	text, err := detectText(imagePath)
	if err != nil {
		fmt.Printf("Error detecting text: %v\n", err)
		return
	}

	err = copyToClipboard(text)
	if err != nil {
		fmt.Printf("Error copying text to clipboard: %v\n", err)
	} else {
		fmt.Println("Text copied to clipboard.")
	}

	if imagePath != "" {
		err = os.Remove(imagePath)
		if err != nil {
			fmt.Printf("Error removing temporary image file: %v\n", err)
		}
	}
}
