from peripherals.conn import readConn


def read(message):
  print("rrr: ", message)


# print(os.environ)
readConn("buzzer", read)
