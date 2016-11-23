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

[jue nov 24 00:30:54 CET 2016]

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
    Latency     8.69ms    3.70ms 216.80ms   76.65%
    Req/Sec    44.68k     2.59k   47.60k    80.00%
  222564 requests in 5.03s, 35.23MB read
Requests/sec:  44246.66
Transfer/sec:      7.00MB

Finish process: python3 python/sanic_simplehttpserver.py
------

Start process: ./golang/simple_http_server_json
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.21ms    1.39ms  18.15ms   86.37%
    Req/Sec   116.68k    10.99k  128.56k    88.00%
  582084 requests in 5.05s, 86.04MB read
Requests/sec: 115252.54
Transfer/sec:     17.04MB

Finish process: ./golang/simple_http_server_json
------

Start process: ./golang/valyala_simple_httpserver
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.44ms  506.86us  10.36ms   88.38%
    Req/Sec   146.21k     8.63k  173.07k    86.00%
  728228 requests in 5.05s, 129.87MB read
Requests/sec: 144069.91
Transfer/sec:     25.69MB

Finish process: ./golang/valyala_simple_httpserver
------

Start process: node nodejs/simple_httpserver_json.js
Listening at http://localhost:8090
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    12.51ms    3.17ms 220.11ms   91.40%
    Req/Sec    30.16k     4.18k   34.27k    88.00%
  149878 requests in 5.03s, 21.01MB read
Requests/sec:  29821.89
Transfer/sec:      4.18MB

Finish process: node nodejs/simple_httpserver_json.js
------

Start process: python3 python/aiohttp_simple_http_server.py
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   100.89ms  218.74ms   2.00s    94.58%
    Req/Sec     3.46k   173.97     3.74k    78.00%
  17205 requests in 5.03s, 3.41MB read
  Socket errors: connect 0, read 0, write 0, timeout 65
Requests/sec:   3422.19
Transfer/sec:    695.13KB

Finish process: python3 python/aiohttp_simple_http_server.py
------

Start process: python3 python/aiohttp_uvloop_simple_http_server.py
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    82.12ms   26.51ms 266.80ms   80.86%
    Req/Sec     4.82k     1.56k    7.20k    61.22%
  23891 requests in 5.03s, 4.74MB read
Requests/sec:   4754.21
Transfer/sec:      0.94MB

Finish process: python3 python/aiohttp_uvloop_simple_http_server.py
------

Start process: python3 python/flask_simple_http_server.py
Running 5s test @ http://127.0.0.1:8090/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    77.40ms  120.16ms   1.65s    92.38%
    Req/Sec     2.54k   651.34     2.88k    90.24%
  10546 requests in 5.02s, 1.56MB read
Requests/sec:   2100.46
Transfer/sec:    317.94KB

Finish process: python3 python/flask_simple_http_server.py
------

```
