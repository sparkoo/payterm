from peripherals.conn import conn


def handle(message):
  print("bzz: " + message)


conn("buzzer", handle, None)
