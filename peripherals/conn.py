import asyncio

import websockets


def conn(addr, readFunc, writeFunc):
  asyncio.get_event_loop() \
    .run_until_complete(handle('ws://localhost:8080/' + addr, readFunc, writeFunc))


async def handle(uri, readFunc, writeFunc):
  while True:
    try:
      async with websockets.connect(uri) as websocket:
        print("connected, sending ready...")
        await websocket.send("ready")

        await asyncio.create_task(read(websocket, readFunc))
        await asyncio.create_task(write(websocket, writeFunc))

        while True:
          await asyncio.sleep(10)
        # while True:
        #   try:
        #     await asyncio.wait({read(websocket, readFunc), write(websocket, writeFunc)}, return_when='FIRST_COMPLETED')
        #   except KeyboardInterrupt:
        #     print("quitting ...")
        #     await websocket.close()
        #     exit(0)
    except Exception:
      print("something failed, trying connect again ...")
      await asyncio.sleep(1)


async def read(websocket, func):
  while True:
    message = await websocket.recv()
    if message == "ping":
      # print(message)
      await websocket.send("pong")
    else:
      print("recv: ", message)
      func(message)


async def write(websocket, func):
  while True:
    await websocket.send(func())
