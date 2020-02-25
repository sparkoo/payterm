import os

from peripherals.conn import conn


def write():
  key = input("key: ")
  return key


# print(os.environ)
conn("keyboard", None, write)
