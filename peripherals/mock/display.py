from peripherals.conn import conn


def read(message):
  print(message)


def write():
  key = input("eh: ")
  return key


conn("display", read, write)
