import board
import digitalio
import adafruit_character_lcd.character_lcd as characterlcd
import importlib.util
spec = importlib.util.spec_from_file_location("conn", "lib/conn.py")
conn = importlib.util.module_from_spec(spec)
spec.loader.exec_module(conn)

lcd = None
url = "ws://localhost:8080/display"


def initDisplay():
  global lcd

  # doc:https://circuitpython.readthedocs.io/projects/charlcd/en/latest/api.html

  # Modify this if you have a different sized character LCD
  lcd_columns = 16
  lcd_rows = 2

  # compatible with all versions of RPI as of Jan. 2019
  # v1 - v3B+
  lcd_rs = digitalio.DigitalInOut(board.D26)
  lcd_en = digitalio.DigitalInOut(board.D19)
  lcd_d4 = digitalio.DigitalInOut(board.D13)
  lcd_d5 = digitalio.DigitalInOut(board.D21)
  lcd_d6 = digitalio.DigitalInOut(board.D20)
  lcd_d7 = digitalio.DigitalInOut(board.D16)

  # Initialise the lcd class
  lcd = characterlcd.Character_LCD_Mono(lcd_rs, lcd_en, lcd_d4, lcd_d5, lcd_d6,
                                        lcd_d7, lcd_columns, lcd_rows)

  # wipe LCD screen before we start
  lcd.clear()

  lcd.message = "1234567890123456\nDisplay initialized ..."


def writeToDisplay(message):
  global lcd
  lcd.clear()
  lcd.message = message


def handle(message):
  writeToDisplay(message)
  print(message)


initDisplay()
conn.readConn("display", handle)
