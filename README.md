# Groot
A Distributed log analyser written in Golang.

#### Why GO ? 
The motivation for using GO was to make things concurrent so that the design is easily scalable and would make things efficient if we have multiple cores. Our code is generic enough to handle changes in configuration on groot and worker side. (Live changes are also supported).
Both groot and worker can be started just by giving an environment variable, the setup is minimal.

We have two entities - groot and worker. Groot is our main query server which runs on the admin side. 
A worker runs on every node. All the queries and their results are communicated via network calls.

To get the system running we have to set the config file according to their roles. Mainly we have config for groot and worker. 

Dependencies : 
1. Linux environment on both groot and worker side.
2. Go installed on both sides.

### Running Groot

1. Modify config file according to requirements (`config/groot.json`):
```
	"HOSTNAME": "0.0.0.0", 
    	"PORT": "8080",
	"NODE_COUNT": 1,
	"NUM_PROCESS": [4],
	"NODES": {
		"node0": "http://localhost/8005"
	}
  ```
Where PORT is the port where groot is listening from all the workers.
NUM_PROCESS is an array denoting the number of processes in each node.
NODES is a dictionary where key is the "node$index" (the format for key shouldn't be changed) and value is the address on which worker is listening. 

2. Steps to run : 

```
$ export ENV=groot
$ go run src/cmd/main.go
-> 
```
3. Groot commands : 
```
-> grep DBG
## this will show the merged results of grep DBG queried on all the workers.
-> live 
## this will start our server at HOSTNAME:PORT and show all live logs from all the workers at the prompt.
-> between T1 T2
## this will show all the logs genereated between timestamp T1 and T2 for all the workers.
```

### Running Worker

1. Modify config file according to requirements (`config/worker.json`):
```
	"HOSTNAME": "0.0.0.0",
    	"PORT": "8005",
	"PROCESS_COUNT": 4,
	"LIVE_ENDPOINT": "http://localhost/api/live/",
	"NODE_DIR": "/home/raghukul/go/src/github.com/raghukul01/groot/log_simulator/HackNode1/",
	"PROCESSES": {
		"0": "Process1.log",
		"1": "Process2.log",
		"2": "Process3.log",
		"3": "Process4.log"
	}
  ```
Where PORT is the port where worker is listening from the groot.
PROCESS_COUNT is the number of process on this worker.
LIVE_ENDPOINT is the endpoint of groot which is listening to live logs.
PROCESSES is a map for storing the log file name for all the process on the worker.
(The key for this map shouldn't be changed.)

2. Steps to run : 

```
$ export ENV=worker
$ go run src/cmd/main.go
## this will start our server at HOSTNAME:PORT 
```
