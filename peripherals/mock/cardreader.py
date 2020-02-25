import os
import threading
from queue import Queue

import time

from peripherals.conn import conn


messageQ = Queue()

def write():
  while True:
    messageQ.put(input("key: "))


def read(message):
  print(message)


# print(os.environ)


threading.Thread(target=write).start()

conn("cardreader", read, messageQ)
