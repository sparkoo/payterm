import http.client
import time


def readConn(addr, func):
  client = http.client.HTTPConnection("localhost:8080")
  while True:
    try:
      client.request("GET", "/" + addr, "ready")
      response = client.getresponse().read().decode()
      func(response)
    except (ConnectionRefusedError, http.client.RemoteDisconnected):
      print("Connection refused, maybe server isn't running.")
      print("Trying again in 5s ...")
      client = http.client.HTTPConnection("localhost:8080")
      time.sleep(5)


def writeConn(addr, func):
  while True:
    client = http.client.HTTPConnection("localhost:8080")
    client.request("GET", "/" + addr, func())
