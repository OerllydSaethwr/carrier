import random
import socket
import struct
import sys
from json import dump

# Pass 2 args: number of hosts to generate, output file

hf = {"hosts": [], "fronts": []}

for i in range(int(sys.argv[1])):
    ip = "127.0.0.1"
    front = ip + ":" + str(9000+i)
    hf["hosts"].append(ip)
    hf["fronts"].append(front)

with open(sys.argv[2], 'w') as f:
    dump(hf, f, indent=4)

