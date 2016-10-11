# binsearch

A simple utility for finding a binary pattern (and offset to it) in a file.
Because it is written in Go, it is multi-platform.

## Installation

* Install a Go toolkit
* Set a GOPATH: export GOPATH=$HOME/go
* Get sources: go get github.com/k3a/binsearch
* Enjoy the utility: $HOME/go/bin/binsearch

## Usage

```bash
./binsearch -string StringToFind /file/to/search
./binsearch -hex 1a2b3c /file/to/search
```


