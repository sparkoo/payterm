from iohelpers import connectionLoop
from iohelpers import readLoop

connectionLoop("testWrite", readLoop, "card: ")
