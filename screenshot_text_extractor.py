import tempfile
import subprocess
from google.cloud import vision
import io
import os

def capture_screenshot():
    try:
        with tempfile.NamedTemporaryFile(suffix=".png", delete=False) as temp:
            temp_image_path = temp.name
            result = subprocess.run(["screencapture", "-i", temp_image_path])
        return temp_image_path
    except subprocess.SubprocessError as e:
        print(f"Error capturing screenshot (SubprocessError): {e}")
        return None
    except Exception as e:
        print(f"Error capturing screenshot (General Error): {e}")
        return None

def detect_text(image_path):
    if image_path is None:
        return ""

    try:
        client = vision.ImageAnnotatorClient()

        with io.open(image_path, 'rb') as image_file:
            content = image_file.read()

        image = vision.Image(content=content)
        response = client.text_detection(image=image)

        if response.error.message:
            raise Exception(f"{response.error.message}\nFor more info on error messages, check: https://cloud.google.com/apis/design/errors")

        texts = response.text_annotations
        return texts[0].description
    except Exception as e:
        print(f"Error detecting text: {e}")
        return ""

def copy_to_clipboard(text):
    if not text:
        return False

    try:
        process = subprocess.Popen(['pbcopy'], stdin=subprocess.PIPE)
        process.communicate(text.encode('utf-8'))
        returncode = process.returncode
        return returncode == 0
    except subprocess.SubprocessError as e:
        print(f"Error copying text to clipboard: {e}")
        return False

def main():
    image_path = capture_screenshot()
    text = detect_text(image_path)
    success = copy_to_clipboard(text)

    if success:
        print("Text copied to clipboard.")
    else:
        print("Error copying text to clipboard.")
        
    if image_path:
        try:
            os.remove(image_path)
        except OSError as e:
            print(f"Error removing temporary image file: {e}")

if __name__ == "__main__":
    main()
