import random
import socket
import struct
import sys
from json import dump

# Pass 2 args: number of hosts to generate, output file
# set number of hosts to generate in the settings below by setting nodes

hf = {
    "hosts": [],
    "fronts": [],
    "settings": {
        "tsx-size": 128,
        "rate": 100000,
        "nodes": 4,

        "decision-port": 3000,
        "carrier-port": 4000,
        "client-port": 5000,

        "mempool-threshold": 100,
        "forward-mode": 0,
        "log-level": "info",

        "carrier-conn-retry-delay": 1000,
        "carrier-conn-max-retry": 0,
        "node-conn-retry-delay": 1000,
        "node-conn-max-retry": 0,

        "local-base-port": 6000,
        "local-front-port": 9000,
    }
}

for i in range(hf["settings"]["nodes"]):
    ip = "127.0.0.1"
    front = ip + ":" + str(9000+i)
    hf["hosts"].append(ip)
    hf["fronts"].append(front)

with open(sys.argv[1], 'w') as f:
    dump(hf, f, indent=4)

