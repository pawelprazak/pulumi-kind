module github.com/virtuslab/pulumi-kind/provider

go 1.16

replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
replace k8s.io/client-go => k8s.io/client-go v0.18.5 // same as in terraform-provider-kind v0.0.7

require (
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/kyma-incubator/terraform-provider-kind v0.0.7
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
)
