import os

from peripherals.conn import conn


def read(message):
  print(message)


# print(os.environ)
conn("display", read, None)
