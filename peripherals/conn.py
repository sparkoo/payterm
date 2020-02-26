import asyncio
import websockets


async def read(addr, func):
  uri = "ws://localhost:8080/" + addr
  async with websockets.connect(uri) as websocket:
    await websocket.send("ready")
    while True:
      message = await websocket.recv()
      func(message)


async def write(addr, func):
  uri = "ws://localhost:8080/" + addr
  async with websockets.connect(uri) as websocket:
    await websocket.send("ready")
    while True:
      message = func()
      print("sending: ", message)
      await websocket.send(message)


def readConn(addr, func):
  asyncio.get_event_loop().run_until_complete(read(addr, func))


def writeConn(addr, func):
  asyncio.get_event_loop().run_until_complete(write(addr, func))
