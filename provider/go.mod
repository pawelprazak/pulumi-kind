module github.com/pawelprazak/pulumi-kind/provider

go 1.16

replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0

replace k8s.io/client-go => k8s.io/client-go v0.18.5 // same as in terraform-provider-kind v0.0.7

require (
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/kyma-incubator/terraform-provider-kind v0.0.9
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.5.2-0.20210623115523-414367963f50
)

replace (
	github.com/pulumi/pulumi-terraform-bridge/v3 => /Users/pprazak/repos/vl/pulumi-terraform-bridge
	github.com/pulumi/pulumi/pkg/v3 => /Users/pprazak/repos/vl/pulumi/pkg
	github.com/pulumi/pulumi/sdk/v3 => /Users/pprazak/repos/vl/pulumi/sdk
)
