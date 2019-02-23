# Helm Chart Testing Example using Terratest

This repository contains the example code from the blog post ["Testing Helm Charts using Terratest"](TODO).

The directory structure represents a typical helm chart repository containing various charts in the `charts` directory
and the tests for those charts under the `test` directory. The `test` folder contains the corresponding terratest code
for testing each of the charts in the `charts` directory.

## Quickstart

### Prerequisites

To run the tests, you will need:

1. A working go install. See [here](https://golang.org/doc/install) for instructions on installing go on to your
   platform.
1. [`dep`](https://golang.github.io/dep/) for installing the package dependencies. See
   [here](https://golang.github.io/dep/docs/installation.html) for instructions on installing dep.
1. This repo checked out in your `GOPATH` in the format go expects. The easiest way to do this is to use `go env` to
   lookup the `GOPATH` and derive the directory from there. The following commands should work on most Unix go
   installations:

    ```
    mkdir -p "$(go env GOPATH)/src/github.com/gruntwork-io"
    git clone https://github.com/gruntwork-io/terratest-helm-testing-example.git "$(go env GOPATH)/src/github.com/gruntwork-io/terratest-helm-testing-example"
    ```

   <!-- TODO: figure out Windows environments -->

Once you have the repository checked out and the necessary tools in place, you are ready to run the tests.

### Kubernetes cluster

Some of the tests (specifically the integration tests) require a Kubernetes cluster to run. The easiest way to get a
Kuberenetes cluster is to use [`minikube`](https://kubernetes.io/docs/setup/minikube/), which is a local single node
installation of Kubernetes. Follow the instructions in the link to install and setup `minikube`.

Once `minikube` is installed, you will also need to deploy Tiller to run the tests based on `helm install`. To deploy
Tiller, you will first need to install the helm client. Follow the [official
guide](https://helm.sh/docs/using_helm/#installing-helm) for instructions on installing `helm`.

Once `helm` is installed, you can setup Tiller on `minikube` using `helm init`:

```
helm init --wait
```

To verify Tiller deployed correctly, use `helm version`. If the deploy was successful, you should see both the client
and server versions:

```
$ helm version
Client: &version.Version{SemVer:"v2.11.0", GitCommit:"2e55dbe1fdb5fdb96b75ff144a339489417b146b", GitTreeState:"clean"}
Server: &version.Version{SemVer:"v2.11.0", GitCommit:"2e55dbe1fdb5fdb96b75ff144a339489417b146b", GitTreeState:"clean"}
```

### Running the tests

To run the tests, first change directory to the `test` folder of this repository:

```
cd "$(go env GOPATH)/src/github.com/gruntwork-io/terratest-helm-testing-example/test"
```

You then need to pull in the dependencies using `dep`:

```
dep ensure
```

This will clone all the referenced go packages into a local folder called `vendor` so that go knows where to find them.

Finally, run the tests:

```
go test -v .
```
