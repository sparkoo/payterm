from iohelpers import connectionLoop
from iohelpers import writeLoop

connectionLoop("keyboard", writeLoop, "key: ")
