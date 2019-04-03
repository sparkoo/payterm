from websocket import create_connection

ws = create_connection("ws://localhost:8080/cardreader")

while True:
    key = input("card: ")
    ws.send(key)
    result = ws.recv()
    print(result)

ws.close()
