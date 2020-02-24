import os

from peripherals.conn import conn


def handle(message):
  print(message)


# print(os.environ)
conn("display", handle)
