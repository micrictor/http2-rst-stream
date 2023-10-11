# http2-rst-stream

Sort of like Slowloris, but for HTTP2. A proof of concept for [CVE-2023-44487](https://nvd.nist.gov/vuln/detail/CVE-2023-44487)

This package contains a client that sends HTTP2 stream resets (RST_STREAM) to a local server. It does this by misusing a Go HTTP client error handler that sends a reset if the request Body is longer than the request Content-Length.

To run:
```
go build && ./http2-rst-stream
```

It intentionally does not allow targeting of remote web servers out of the box. I'm not trying to help people break stuff.

Original inspiration: https://edg.io/lp/blog/resets-leaks-ddos-and-the-tale-of-a-hidden-cve
