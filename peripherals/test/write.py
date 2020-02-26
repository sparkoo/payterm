from peripherals.conn import writeConn


def write():
  return input("yes: ")


# print(os.environ)
writeConn("keyboard", write)
