import asyncio
from queue import Queue

import websockets


messageQ = Queue()


def conn(addr, func):
  asyncio.get_event_loop() \
    .run_until_complete(handle('ws://localhost:8080/' + addr, func))


async def handle(uri, func):
  while True:
    try:
      async with websockets.connect(uri) as websocket:
        print("connected, sending ready...")
        await websocket.send("ready")
        await handler(websocket, func)

    except KeyboardInterrupt:
      print("quitting ...")
      await websocket.close()
      exit(0)
    except Exception:
      print("something failed, trying connect again ...")
      await asyncio.sleep(1)


async def producer():
  global messageQ
  message = messageQ.get()
  print("send: ", message)
  messageQ.task_done()
  return message


async def consumer(message, func):
  print("recv: ", message)
  global messageQ
  if message == "ping":
    print(message)
    messageQ.put("pong")
  else:
    func(message)


async def consumer_handler(websocket, func):
  print("1")
  async for message in websocket:
    print("Revc: ", message)
    await consumer(message, func)


async def producer_handler(websocket, func):
  print("2")
  while True:
    message = await producer()
    print("send: ", message)
    await websocket.send(message)


async def handler(websocket, func):
  consumer_task = asyncio.ensure_future(
      consumer_handler(websocket, func))
  producer_task = asyncio.ensure_future(
      producer_handler(websocket, func))
  done, pending = await asyncio.wait(
      [consumer_task, producer_task],
      return_when=asyncio.FIRST_COMPLETED,
  )
  for task in pending:
    task.cancel()
