---
title: Try it out!
description: How to see Athens in action 
menu: shortcuts
---

## Try out Athens

To quickly see Athens in action, follow these steps:

First, make sure you have [Go 1.11 installed](https://gophersource.com/setup/) and that GOPATH/bin is on your path.

Next, use Go to install and run the Athens proxy in a background process

```console
$ go get -u github.com/gomods/athens/cmd/proxy
# the source is downloaded to GOPATH/src/github.com/gomods/athens/
$ proxy &
[1] 25243
INFO[0000] Starting application at 127.0.0.1:3000
```

Next, you will need to enable the [Go Modules](https://github.com/golang/go/wiki/Modules)
feature and configure Go to use the proxy!

**Bash**
```bash
export GO111MODULE=on
export GOPROXY=http://127.0.0.1:3000
```

**PowerShell**
```powershell
$env:GO111MODULE = "on"
$env:GOPROXY = "http://127.0.0.1:3000"
``@ma`


Now, when you build and run this example application, **go** will fetch dependencies via Athens!

```console
$ git clone https://github.com/athens-artifacts/walkthrough.git
$ cd walkthrough
$ go run .
go: finding github.com/athens-artifacts/samplelib v1.0.0
handler: GET /github.com/athens-artifacts/samplelib/@v/v1.0.0.info [200]
handler: GET /github.com/athens-artifacts/samplelib/@v/v1.0.0.mod [200]
go: downloading github.com/athens-artifacts/samplelib v1.0.0
handler: GET /github.com/athens-artifacts/samplelib/@v/v1.0.0.zip [200]
The 🦁 says rawr!
```

The output from `go run .` includes attempts to find the **github.com/athens-artifacts/samplelib** dependency. Since the
proxy was run in the background, you should also see output from Athens indicating that it is handling requests for the dependency.

This should give you an overview of what using Athens is like!

