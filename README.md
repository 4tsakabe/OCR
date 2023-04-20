# Screenshot Text Extractor (for macOS Monterey)


macOS Montereyでは、OCR機能にて日本語がサポート対象外であるため、このPythonスクリプトが役立ちます。Google Cloud Vision APIを使用して、スクリーンショット内のテキストを検出し、抽出したテキストをクリップボードにコピーできます。画像からのテキストや画面上のアクセス困難なテキスト領域から素早くテキストをコピーするのに便利なツールです。

※Google Cloud Vision APIの登録が必要です。(月1000回まで無料)

※VenturaであればOCR機能に日本語が追加されています。[TRex](https://github.com/amebalabs/TRex)がおすすめです。
※TRex使ってみたけど日本語が中国の文字が入っていたりしそうなので、VisionAPIの方が綺麗な結果っぽいのでVenturaでも使える

<img src="https://user-images.githubusercontent.com/61626658/226303592-0a1fdafe-2eaf-4547-a365-68d40ddb1fb6.gif" width="300">

This Python script allows you to capture a screenshot, detect text within the screenshot using Google Cloud Vision API, and copy the extracted text to your clipboard. It is a convenient tool to quickly copy text from images or inaccessible text regions on your screen.

## Requirements
- Python 3.6 or later
- Google Cloud Vision API credentials
- macOS

## Dependencies
Before you run the script, make sure to install the following dependencies:

```
pip install google-cloud-vision
```

## Setup
1. Create a new Google Cloud project if you haven't already.
2. Enable the Google Cloud Vision API for your project.
3. Create a service account key for the Google Cloud Vision API. Download the JSON key file and set the environment variable `GOOGLE_APPLICATION_CREDENTIALS` to the path of the key file.

```
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/credentials.json"
```

## How to Use

1. Run the script:
```
python screenshot_text_extractor.py
```
2. When prompted, select the area of the screen you want to capture using the crosshair cursor.
3. The script will detect any text within the screenshot and copy it to your clipboard.
4. If the text was successfully copied to the clipboard, you'll see the message "Text copied to clipboard." Otherwise, an error message will be displayed.

## Limitations
- This script is designed for macOS, and may not work on other operating systems due to the use of the `screencapture` command and `pbcopy`.
- The text recognition quality depends on the Google Cloud Vision API.

## Troubleshooting
If you encounter any issues or error messages, make sure you have set up your Google Cloud project, enabled the Vision API, and set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable correctly. If problems persist, refer to the Google Cloud Vision API documentation for more information.
