from websocket import create_connection

ws = create_connection("ws://localhost:8080/keyboard")

while True:
    key = input("key: ")
    ws.send(key)
    result = ws.recv()
    print(result)

ws.close()
