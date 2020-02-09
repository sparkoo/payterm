from iohelpers import connectionLoop
from iohelpers import readLoop

connectionLoop("display", readLoop, "display: ")
