import time
import importlib.util
spec = importlib.util.spec_from_file_location("conn", "../lib/conn.py")
conn = importlib.util.module_from_spec(spec)
spec.loader.exec_module(conn)


def bzz(message):
  print("bzz: " + message)
  buzz = str(message).split(":")

  frequency = int(buzz[0])
  length = int(buzz[1])
  period = 1.0 / frequency
  delayValue = period / 2

  print("bzz %dHz for %dms" % (frequency, length))
  print("delay: %s" % delayValue)

  now = time.time_ns()
  future = now + (length * 1000)
  print("b", end='')
  while time.time_ns() < future:
    print("e", end='')
  print("p")


# print(os.environ)
conn.readConn("buzzer", bzz)
