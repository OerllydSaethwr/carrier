import random
import socket
import struct
import sys
from json import dump

# Pass 2 args: number of hosts to generate, output file

hf = {"hosts": [], "fronts": []}

for i in range(int(sys.argv[1]):
    ip = socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff)))
    front = ip + ":" + "9000"
    hf["hosts"].append(ip)
    hf["fronts"].append(front)

with open(sys.argv[2], 'w') as f:
    dump(hf, f, indent=4)

