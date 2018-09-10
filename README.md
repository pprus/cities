# cities

HTTP server that accepts JSON payloads of the form:

`{ “count”: 4, “city”: “Atlanta” }`

And exposes the per-city totals in Prometheus format.

# Installation

`go get github.com/pprus/cities`

# Running

`go run cities.go` (or just run the executable)

The server listens on port 8080 and exposes metrics at http://localhost:8080/metrics.
