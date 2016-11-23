# HTTP Server Tests

Testing several HTTP servers in Python, Golang and NodeJS.

## Run benchmark
To run benchmark:

```
$ chmod +x http_benchmarking.sh
$ ./http_benchmarking.sh > results.log
```

## Results
Benchmark results:

```
Starting benchmark...

Simple HTTP Servers with a JSON response:
 {Name:HelloWorld,Types:[hello,world]}

Versions:
	Operative System - 	 [GNU/Linux]
	Kernel - 	 [4.4.0-47-generic #68-Ubuntu SMP Wed Oct 26 19:39:52 UTC 2016]
	Hardware-platform - 	 [x86_64]
	Python - 	 [Python 3.5.2]
	Golang - 	 [go version go1.6.2 linux/amd64]
	NodeJS - 	 [v6.9.1]

Start process: python3 python/sanic_simplehttpserver.py
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.60ms    4.10ms 216.89ms   75.01%
    Req/Sec    40.46k     3.22k   43.91k    86.00%
  201531 requests in 5.02s, 31.90MB read
Requests/sec:  40124.52
Transfer/sec:      6.35MB

Finish process: python3 python/sanic_simplehttpserver.py
------

Start process: ./golang/simple_http_server_json
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.72ms  840.89us  17.37ms   80.58%
    Req/Sec   138.24k     4.36k  147.06k    80.00%
  690362 requests in 5.05s, 102.05MB read
Requests/sec: 136688.02
Transfer/sec:     20.21MB

Finish process: ./golang/simple_http_server_json
------

Start process: node nodejs/simple_httpserver_json.js
Listening at http://localhost:8090
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    13.72ms    5.18ms 244.65ms   98.82%
    Req/Sec    27.61k     3.16k   30.43k    92.00%
  137239 requests in 5.02s, 19.24MB read
Requests/sec:  27321.32
Transfer/sec:      3.83MB

Finish process: node nodejs/simple_httpserver_json.js
------

Start process: python3 python/aiohttp_simple_http_server.py
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    98.48ms  210.66ms   2.00s    94.93%
    Req/Sec     3.44k    93.52     3.70k    78.00%
  17109 requests in 5.03s, 3.39MB read
  Socket errors: connect 0, read 0, write 0, timeout 67
Requests/sec:   3400.27
Transfer/sec:    690.71KB

Finish process: python3 python/aiohttp_simple_http_server.py
------

Start process: python3 python/aiohttp_uvloop_simple_http_server.py
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    88.92ms   26.35ms 237.68ms   83.19%
    Req/Sec     4.46k     1.54k    7.34k    73.47%
  21948 requests in 5.03s, 4.35MB read
Requests/sec:   4361.77
Transfer/sec:      0.87MB

Finish process: python3 python/aiohttp_uvloop_simple_http_server.py
------

Start process: python3 python/flask_simple_http_server.py
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    69.50ms  103.22ms   1.77s    94.72%
    Req/Sec     2.43k   590.71     2.81k    91.11%
  11053 requests in 5.04s, 1.63MB read
  Socket errors: connect 0, read 0, write 0, timeout 1
Requests/sec:   2192.74
Transfer/sec:    331.91KB

Finish process: python3 python/flask_simple_http_server.py
------

```
