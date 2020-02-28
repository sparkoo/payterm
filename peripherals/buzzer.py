import RPi.GPIO as GPIO
import time
import importlib.util
spec = importlib.util.spec_from_file_location("conn", "lib/conn.py")
conn = importlib.util.module_from_spec(spec)
spec.loader.exec_module(conn)


buzzer_pin = 1

GPIO.setmode(GPIO.BCM)
GPIO.setup(buzzer_pin, GPIO.IN)
GPIO.setup(buzzer_pin, GPIO.OUT)


def bzz(message):
  print("bzz: " + message)
  buzz = str(message).split(":")

  length = int(buzz[1])
  frequency = int(buzz[0])
  period = 1.0 / frequency
  delayValue = period / 2

  try:
    now = time.time_ns()
    future = now + (length * 1000)
    while time.time_ns() < future:
      GPIO.output(buzzer_pin, True)
      time.sleep(delayValue)
      GPIO.output(buzzer_pin, False)
      time.sleep(delayValue)
  except KeyboardInterrupt:
    print(" bye")
  finally:
    GPIO.cleanup()


conn.readConn("buzzer", bzz)
