from peripherals.conn import writeConn


def write():
  return input("cr: ")


# print(os.environ)
writeConn("cardreader", write)
