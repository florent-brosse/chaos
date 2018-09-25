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