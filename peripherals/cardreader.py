import RPi.GPIO as GPIO
from mfrc522 import SimpleMFRC522
import importlib.util
spec = importlib.util.spec_from_file_location("conn", "lib/conn.py")
conn = importlib.util.module_from_spec(spec)
spec.loader.exec_module(conn)

reader = SimpleMFRC522()

def readCard():
  try:
    cardid, text = reader.read()
    print(cardid)
    print(text)
    return text
  finally:
    GPIO.cleanup()


conn.writeConn("cardreader", readCard)
