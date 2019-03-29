from websocket import create_connection

ws = create_connection("ws://localhost:8080/display")

while True:
    ws.send("ready")
    result = ws.recv()
    print(result)

ws.close()