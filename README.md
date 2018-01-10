# Kubernetes Example Application

## Description

Example application that illustrates how to build an application that is Cloud Native and will correctly run within Kubernetes.
Created by following [this tutorial](https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/). All credit goes to [Elena](https://github.com/rumyantseva) for making a great tutorial

The application itself is quite trivial in that it has a basic root endpoint that just shows some build information. The interesting parts of the application are that:
- it closely follows the [12-factor app](https://12factor.net) methodology.
- has metrics instrumented with Prometheus.
- employs a health endpoint that Kubernetes can use to evaluate the health of a replica of the app.
- employs a readiness endpoint that Kubernetes can use to evaluate if a replica of the app is ready to accept traffic.
- embeds build and versioning information at build time through the Makefile.

## Prerequisites

### Golang

It is expected that you have Golang installed and your `$GOPATH` configured correctly. Please follow the [following guide](https://golang.org/doc/install) to set this up.

### Kubernetes

As this is a Kubernetes example application, it is expected that you have `kubectl` installed to your path and its context is set to a live Kubernetes cluster.

## Usage

There is a Makefile for the app located within the `app/` directory the commands described below will help build, test and deploy this application.

- To run `make run`
- To build `make container`
- To deploy app to Kubernetes `make deploy`

#### Docker for Mac Kubernetes Support

I've been testing the single node Kubernetes cluster that comes with Docker for Mac 17.12.0-ce so you can use the `docker-compose.yml` file to deploy the stack to Kubernetes using the `docker stack deploy --compose-file=docker-compose.yml example-app`

## TO DO

- [x] Add Kubernetes files
- [ ] Expand background for Documentation
- [ ] Add examples to Documentation

## License

MIT
