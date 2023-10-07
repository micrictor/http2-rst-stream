# http2-rst-stream


Sort of like Slowloris, but for HTTP2

This package contains a client that sends stream RSTs to a local server.


It doesn't handle concurrent streams well, so instead it parallelizes by
creating N unique HTTP clients. This probably makes it less effective at
resource exhaustion, but I'm not actually trying to harm stuff, just prove
out what is described in
https://edg.io/lp/blog/resets-leaks-ddos-and-the-tale-of-a-hidden-cve
