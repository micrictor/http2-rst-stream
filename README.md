# http2-rst-stream


Sort of like Slowloris, but for HTTP2

This package contains a client that sends stream RSTs to a local server.


It doesn't handle concurrent streams well, so instead it parallelizes individual HTTP2 clients.
