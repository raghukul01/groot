# Groot
A Distributed log analyser written in Golang.

We have two entities - groot and worker. Groot is our main query server which runs on the admin side. 
A worker runs on every node. All the queries and their results are communicated via network calls.

To get the system running we have to set the config file according to their roles. Mainly we have config for groot and worker. 

Dependencies : 
1. Linux environment on both groot and worker side.
2. Go installed on both sides.

### Running Groot

1. Modify config file according to requirements :
    ```{
	"HOSTNAME": "0.0.0.0",
    "PORT": "8080",
	"NODE_COUNT": 1,
	"NUM_PROCESS": [4],
	"NODES": {
		"node0": "http://localhost/8005"
	}
    }```


