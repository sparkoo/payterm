import asyncio
import websockets
import time


async def hello(uri):
  async with websockets.connect(uri) as websocket:
    await websocket.send("ready")

    while True:
      try:
        message = await websocket.recv()
        print(message)
        await websocket.send("ready")
      except KeyboardInterrupt:
        print("quitting ...")
        await websocket.close()
        exit(0)
      except ConnectionError:
        print("no connection, trying again ...")
        time.sleep(1)

asyncio.get_event_loop().run_until_complete(
    hello('ws://localhost:8080/display'))
