from iohelpers import connectionLoop
from iohelpers import readLoop

connectionLoop("buzzer", readLoop, "bzzz: ")
