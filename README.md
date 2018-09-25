chaos
===

Chaos is a tool to create chaos in your server.


## Overview

Chaos is an agent tool to create chaos in your local machine. It's manageg by another project chaos-manager.

## Installation

Make sure you have a working Go environment.  Go version 1.2+ is supported.  [See
the install instructions for Go](http://golang.org/doc/install.html).

To install cli, simply run:
```
$ go get github.com/florent-brosse/chaos
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

### Supported platforms

Chaos is tested against multiple versions of Go on Linux

```
chaos --ram --ramusage 80%
```
```
chaos --cpu --cpuusage 80%
```
```
chaos --file --fileusage 1% --filepath /tmp/BIGFILE

or run the server
chaos
```

```
curl -X POST \
  http://localhost:7070/scenarios/1 \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{"name":"add cpu","description":"add cpu","tasks":[{"id":"10","type":4,"start":"2018-09-25T17:15:00.757540298+02:00","duration":60000000000,"tags":["toto","DC1"],"param":{"usage":"30%"},"launched":false,"done":false}],"id":"1","done":false}'
```