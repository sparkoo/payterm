from websocket import create_connection
import time

def connectionLoop(url, func, extra):
    while True:
        try:
            ws = create_connection("ws://localhost:8080/" + url)
            print("connected ...")
            func(ws, extra)
        except KeyboardInterrupt:
            print("quitting ...")
            ws.close()
            exit(0)
        except ConnectionError:
            print("no connection, trying again ...")
            time.sleep(1)


def readLoop(ws, prefix):
    while True:
        ws.send("ready")
        print(prefix,end="")
        result = ws.recv()
        print(result)


def writeLoop(ws, prompt):
    key = input("key: ")
    ws.send(key)
    result = ws.recv()
    print(prompt, result)
