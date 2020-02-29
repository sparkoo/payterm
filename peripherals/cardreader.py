import RPi.GPIO as GPIO
import time
from mfrc522 import SimpleMFRC522
import importlib.util
spec = importlib.util.spec_from_file_location("conn", "lib/conn.py")
conn = importlib.util.module_from_spec(spec)
spec.loader.exec_module(conn)


def readCard():
  try:
    time.sleep(1)
    reader = SimpleMFRC522()
    cardid, text = reader.read()
    print(cardid)
    print(text)
    return text.strip()
  finally:
    GPIO.cleanup()


conn.writeConn("cardreader", readCard)
