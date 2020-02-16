import asyncio
import websockets


def conn(addr, func):
  asyncio.get_event_loop() \
    .run_until_complete(handle('ws://localhost:8080/' + addr, func))


async def handle(uri, func):
  async with websockets.connect(uri) as websocket:
    print("connected, sending ready...")
    await websocket.send("ready")

    while True:
      try:
        message = await websocket.recv()
        if message == "ping":
          print(message)
          await websocket.send("pong")
        else:
          func(message)
      except KeyboardInterrupt:
        print("quitting ...")
        await websocket.close()
        exit(0)
      except ConnectionError:
        print("no connection, trying again ...")
        await asyncio.sleep(1)
