module pulumi-go

go 1.16

require (
	github.com/pulumi/pulumi-kubernetes/sdk/v3 v3.3.0
	github.com/pulumi/pulumi/sdk/v3 v3.4.0
)

replace github.com/pawelprazak/pulumi-kind/sdk/kind => /opt/pulumi/kind
