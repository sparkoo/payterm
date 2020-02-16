from peripherals.conn import conn


def handle(message):
  print(message)


conn("display", handle)
