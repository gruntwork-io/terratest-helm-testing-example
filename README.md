# Helm Chart Testing Example using Terratest

This repository contains the example code from the blog post ["Automated Testing for Kubernetes and Helm Charts using Terratest"](https://blog.gruntwork.io/automated-testing-for-kubernetes-and-helm-charts-using-terratest-a4ddc4e67344).

The directory structure represents a typical helm chart repository containing various charts in the `charts` directory
and the tests for those charts under the `test` directory. The `test` folder contains the corresponding terratest code
for testing each of the charts in the `charts` directory.

## Quickstart

### Prerequisites

To run the tests, you will need a working go install. See [here](https://golang.org/doc/install) for instructions on
installing go on to your platform. Make sure to use a version >=1.13.

### Kubernetes cluster

Some of the tests (specifically the integration tests) require a Kubernetes cluster to run. The easiest way to get a
Kuberenetes cluster is to use [`minikube`](https://kubernetes.io/docs/setup/minikube/), which is a local single node
installation of Kubernetes. Follow the instructions in the link to install and setup `minikube`.

Once `minikube` is installed, you will also need to install the helm client run the tests based on `helm install`.
Follow the [official guide](https://helm.sh/docs/intro/install/) for instructions on installing `helm`.

### Running the tests

To run the tests, first change directory to the `test` folder of this repository, and then use `go test` to run the tests:

```
cd test
go test -v .
```
