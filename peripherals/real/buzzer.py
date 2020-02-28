from peripherals.conn import readConn
import RPi.GPIO as GPIO
import time

buzzer_pin = 4

GPIO.setmode(GPIO.BCM)
GPIO.setup(buzzer_pin, GPIO.IN)
GPIO.setup(buzzer_pin, GPIO.OUT)


def bzz(message):
  print("bzz: " + message)
  buzz = str(message).split(":")

  frequency = int(buzz[0])
  length = int(buzz[1])
  period = 1.0 / frequency
  delayValue = period / 2

  try:
    while True:
      GPIO.output(buzzer_pin, True)
      time.sleep(delayValue)
      GPIO.output(buzzer_pin, False)
      time.sleep(delayValue)
  except KeyboardInterrupt:
    print(" bye")
  finally:
    GPIO.cleanup()


readConn("display", bzz)
