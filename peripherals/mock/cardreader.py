import importlib.util
spec = importlib.util.spec_from_file_location("conn", "../lib/conn.py")
conn = importlib.util.module_from_spec(spec)
spec.loader.exec_module(conn)


def write():
  return input("cr: ")


# print(os.environ)
conn.writeConn("cardreader", write)
