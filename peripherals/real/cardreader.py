from peripherals.conn import writeConn
import RPi.GPIO as GPIO
from mfrc522 import SimpleMFRC522

reader = SimpleMFRC522()

def readCard():
  try:
    cardid, text = reader.read()
    print(cardid)
    print(text)
    return text
  finally:
    GPIO.cleanup()


writeConn("cardreader", readCard)
