---
title: FAQ
description: Frequently Asked Questions
menu: shortcuts
---

### Is Athens just a proxy? A registry?

_TL;DR We've discovered that "registry" doesn't describe what we're trying to do here. The term "global proxy pool" is probably a better description, but it's still an open question._

A registry is generally run by one entity, is one logical server that provides authentication (and provenance sometimes), and is pretty much the de-facto only source of dependencies. Sometimes it's run by a for-profit company.

That's most definitely not what we in the Athens community are going for, and that would harm our community if we did go down that path.

We think that a federated discovery/auth/provenance system is a great resource for folks building proxies, and although it's young, we think that the Athens proxy is growing into a good quality implementation. But it doesn't have to be the only one.

We're purposefully building this project - and working with the toolchain folks - in a way that everyone who wants to write a proxy can participate. Even if they don't use the federated bits.

So, if you look back to "architecture" above, there are a few discrete "things" involved in this system we're building. The term "proxy" describes what it's trying to do fairly well. But there are other things going on too. The term "global proxy pool" covers everything in the global, federated system.

### Does Athens integrate with the go toolchain?

Athens is currently supported by the Go 1.11 toolchain via the [download protocol](/intro/protocol/).

For the TL;DR of the protocol, it's a REST API that lets the go toolchain (i.e. go get) see lists of versions and fetch source code for a specific version.

Athens is a server that implements the protocol. Both it, the protocol and the toolchain (as you almost certainly know) is open source.

### Is Athens a centralized registry?

We have in mind an architecture that:

- Provides a centralized authentication system, code provenance, and discovery system for modules in VCSs
- Is run by many companies, likely under a foundation
- Allows proxies (i.e. the Athens proxy) to use it, if they want, to serve go modules that live on public VCSs

### Are the packages served by Athens immutable?

_TL;DR Athens does store code in CDNs and has the option to store code in other persistent datastores._

The longer version:

It's virtually impossible to ensure immutable builds when source code comes from Github. We have been annoyed by that problem for a long time. The Go modules download protocol is a great opportunity to solve this issue. The Athens proxy works pretty simply at a high level:

1. go get github.com/my/module@v1 happens
1. Athens looks in its datastore, it's missing
1. Athens downloads github.com/my/module@v1 from Github (it uses go get on the backend too)
1. Athens stores the module in its datastore
1. Athens serves github.com/my/module@v1 from its datastore forever

To repeat, "datastore" means a CDN (we currently have support for Google Cloud Storage and Azure Blob Storage and AWS S3) or another datastore (we have support for MongoDB, disk and some others).

And, sidenote - we don't have many concrete details on this aforementioned foundation/group, but we would like to see them pay for CDN hosting (and other hosting fees). We are
currently coordinating with hosting providers on these questions.

One final note - we use "caching" in lots of our docs, and that's technically wrong because no data is evicted or expires. We'll need to update that terminology.

### Can the proxy authenticate to private repositories?

_tldr: yes, with proper authentication configuration defined on the Athens proxy host._

When the GOPROXY environment variable is set on the client-side, the go 1.11+ cli
does not attempt to request the meta tags, via a request that looks like `https://example.org/pkg/foo?go-get=1`.

Internally Athens uses `go get` under the hood (`go mod download` to be exact)
without the GOPROXY environment variable set so that `go` will in turn request
the meta tags using the standard authentication mechanisms supported by `go`.
Therefore, if `go` before v1.11 worked for you, then go 1.11+ with GOPROXY
should work as well, provided that the Athens proxy host is configured with the
proper authentication.

### Is there support for monitoring and observability for Proxy?

Right now, we have structured logs for proxy. Along with that, we have added tracing to help developers identify critical code paths and debug latency issues. While there is no setup required for logs, tracing requires some installation. We currently support exporting traces with [Jaeger](https://www.jaegertracing.io/), [GCP Stackdriver](https://cloud.google.com/stackdriver/) & [Datadog](https://docs.datadoghq.com/tracing/) (untested). Further support for other exporters is in progress.

To try out tracing with Jaeger, do the following:

- Set the environment to development (otherwise traces will be sampled)
- Run `docker-compose up -d` that is found in the athens source root directory to initialize the services required
- Run the walkthrough tutorial
- Open `http://localhost:16686/search`

Observability is not a hard requirement for the proxy. So, if the infrastructure is not properly set up, they will fail with an information log. For eg. if Jaeger is not running / if the wrong URL to the exporter is provided, proxy will continue to run. However, it will not collect any traces or metrics (when the exporter backend is unavailable).

