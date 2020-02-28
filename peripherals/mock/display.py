import importlib.util
spec = importlib.util.spec_from_file_location("conn", "../lib/conn.py")
conn = importlib.util.module_from_spec(spec)
spec.loader.exec_module(conn)


def read(message):
  print("rrr: ", message)


# print(os.environ)
conn.readConn("display", read)
