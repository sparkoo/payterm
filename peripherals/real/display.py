import board
import digitalio
import adafruit_character_lcd.character_lcd as characterlcd
import asyncio
import websockets
import time

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
  lcd_rs = digitalio.DigitalInOut(board.D22)
  lcd_en = digitalio.DigitalInOut(board.D17)
  lcd_d4 = digitalio.DigitalInOut(board.D25)
  lcd_d5 = digitalio.DigitalInOut(board.D24)
  lcd_d6 = digitalio.DigitalInOut(board.D23)
  lcd_d7 = digitalio.DigitalInOut(board.D18)

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


async def ws():
  writeToDisplay("connecting ...")
  async with websockets.connect(url) as websocket:
    writeToDisplay("connected...\nsending ready...")
    await websocket.send("ready")

    while True:
      try:
        message = await websocket.recv()
        if message == "ping":
          print("received ping, sending pong")
          await websocket.send("ready")
        else:
          writeToDisplay(message)
          print(message)
          await websocket.send("ready")
      except KeyboardInterrupt:
        print("quitting ...")
        await websocket.close()
        exit(0)


initDisplay()

while True:
  try:
    asyncio.get_event_loop().run_until_complete(ws())
  except Exception as e:
    print(e)
    asyncio.sleep(1)
