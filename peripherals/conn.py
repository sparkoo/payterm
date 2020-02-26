import http.client


def readConn(addr, func):
  client = http.client.HTTPConnection("localhost:8080")
  while True:
    client.request("GET", "/" + addr, "ready")
    response = client.getresponse().read().decode()
    func(response)


def writeConn(addr, func):
  while True:
    client = http.client.HTTPConnection("localhost:8080")
    client.request("GET", "/" + addr, func())
