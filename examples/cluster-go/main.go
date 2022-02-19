package main

import (
	kind "github.com/pawelprazak/pulumi-kind/sdk/kind"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cluster, _ := kind.NewCluster(ctx, "pulumi-cluster", &kind.ClusterArgs{
			KindConfig: kind.ClusterKindConfigArgs{
				ApiVersion: pulumi.String("kind.x-k8s.io/v1alpha4"),
				Kind:       pulumi.String("Cluster"),
				Nodes: kind.ClusterKindConfigNodeArray{
					kind.ClusterKindConfigNodeArgs{
						Role: pulumi.String("worker"),
					},
					kind.ClusterKindConfigNodeArgs{
						Role: pulumi.String("control-plane"),
					},
				},
			},
		})

		ctx.Export("name", cluster.Name)

		return nil
	})
}
