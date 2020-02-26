import socket

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(("localhost", 8080))
s.send(bytes("ready"))
recv = s.recv(1024)
print(recv)
s.close()
