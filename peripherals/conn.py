import asyncio
from queue import Queue, Empty

import websockets

messageQ = Queue()


def conn(addr, readfunc, writefunc):
  asyncio.get_event_loop() \
    .run_until_complete(handle('ws://localhost:8080/' + addr, readfunc, writefunc))


async def handle(uri, readfunc, writeQ):
  while True:
    try:
      async with websockets.connect(uri) as ws:
        print("connected, sending ready...")
        await ws.send("ready")

        receiveTask = asyncio.create_task(receive(ws, readfunc))
        writeTask = asyncio.create_task(write(ws, writeQ))

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
  print("run receiver")
  async for message in websocket:
    print("recv: ", message)
    if message == "ping":
      messageQ.put("pong")
    else:
      if func is not None:
        func(message)


async def write(websocket, writeQ):
  print("run writer")
  while True:
    try:
      message = messageQ.get_nowait()
      messageQ.task_done()
      # print("sending ", message)
      await websocket.send(message)
    except Empty as e:
      # print("write nothing")
      await asyncio.sleep(.01)

    try:
      message = writeQ.get_nowait()
      writeQ.task_done()
      # print("sending ", message)
      await websocket.send(message)
    except Empty as e:
      # print("write nothing")
      await asyncio.sleep(.01)
