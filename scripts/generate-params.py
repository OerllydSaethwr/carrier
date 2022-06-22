import random
import socket
import struct
import sys
from json import dump

# Pass 5 args: output file, number of hosts, tsx-size, rate and init-threshold

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

        "init-threshold": 100,
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

hf["settings"]["nodes"] = int(sys.argv[2])
hf["settings"]["tsx-size"] = int(sys.argv[3])
hf["settings"]["rate"] = int(sys.argv[4])
hf["settings"]["init-threshold"] = int(sys.argv[5])

for i in range(hf["settings"]["nodes"]):
    ip = socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff)))
    front = ip + ":" + "9000"
    hf["hosts"].append(ip)
    hf["fronts"].append(front)

with open(sys.argv[1], 'w') as f:
    dump(hf, f, indent=4)

