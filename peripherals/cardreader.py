from iohelpers import connectionLoop
from iohelpers import writeLoop

connectionLoop("cardreader", writeLoop, "card: ")
