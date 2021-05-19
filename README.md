# Pulumi Provider for kind

This is **experimental project**, for demonstration purposes, **do not use in production**.

The Pulumi Provider for *kind* enables [Pulumi](https://www.pulumi.com) to provision local [Kubernetes](https://kubernetes.io) 
clusters on base of [Kubernetes IN Docker (kind)](https://github.com/kubernetes-sigs/kind).

## Development
This provider is a Terraform Bridge for [terraform-provider-kind](https://github.com/kyma-incubator/terraform-provider-kind).

### Creating a Pulumi Terraform Bridge Provider

First, clone this repo with the name of the desired provider in place of `kind`:

```
git clone https://github.com/pulumi/pulumi-tf-provider-boilerplate pulumi-kind
```

Second, replace references to `kind` with the name of your provider:

```
make prepare NAME=kind REPOSITORY=github.com/virtuslab/pulumi-kind
```

Next, list the configuration points for the provider in the area of the README.

> Note: If the name of the desired Pulumi provider differs from the name of the Terraform provider, you will need to carefully distinguish between the references - see https://github.com/pulumi/pulumi-azure for an example.

#### Adding dependencies

In order to properly build the sdks, the following tools are expected:
- `pulumictl` (See the project's README for installation instructions: https://github.com/pulumi/pulumictl)

In the root of the repository, run:

- `(cd provider && go get github.com/kyma-incubator/terraform-provider-kind)`  (where `kind` is the name of the provider - note the parenthesis to run this in a subshell)
- `(cd provider && go mod download)`

#### Building the provider

- Edit `provider/resources.go` to map each resource, and specify provider information
- `make build_sdks`

#### Testing locally
- `make clean cleanup cleanup_local`
- `make provider lint_provider`
- `make install_local`
- `make build_nodejs`
- `make install_nodejs_sdk`
- `make test`

## Installing
TODO

This package is available in many languages in the standard packaging formats.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/kind

or `yarn`:

    $ yarn add @pulumi/kind

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_kind

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/virtuslab/pulumi-kind/sdk/go/...

## Configuration
TODO

## Reference
TODO

For detailed reference documentation, please visit [the API docs][1].

[1]: https://www.pulumi.com/docs/reference/pkg/kind/
