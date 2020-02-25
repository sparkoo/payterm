import asyncio
from queue import Queue, Empty

import websockets

messageQ = Queue()


def conn(addr, func):
  asyncio.get_event_loop() \
    .run_until_complete(handle('ws://localhost:8080/' + addr, func))


async def handle(uri, func):
  while True:
    try:
      async with websockets.connect(uri) as ws:
        print("connected, sending ready...")
        await ws.send("ready")

        receiveTask = asyncio.create_task(receive(ws, func))
        writeTask = asyncio.create_task(write(ws))

        done, pending = await asyncio.wait(
            {receiveTask, writeTask},
            return_when=asyncio.FIRST_COMPLETED
        )

        print("done")
        for task in pending:
          print("closing others")
          task.cancel()

    except KeyboardInterrupt:
      print("quitting ...")
      await ws.close()
      exit(0)
    except Exception:
      print("something failed, trying connect again ...")
      await asyncio.sleep(1)


async def receive(websocket, func):
  async for message in websocket:
    if message == "ping":
      messageQ.put("pong")
    else:
      func(message)


async def write(websocket):
  while True:
    try:
      message = messageQ.get_nowait()
      messageQ.task_done()
      print("sending ", message)
      await websocket.send(message)
    except Empty as e:
      await asyncio.sleep(.01)
