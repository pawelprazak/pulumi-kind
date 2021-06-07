using Pulumi;
using Pulumi.Kind;
using Pulumi.Kind.Inputs;

class MyStack : Stack
{
    public MyStack()
    {
        var cluster = new Cluster("pulumi-cluster", new ClusterArgs
        {
            KindConfig = new ClusterKindConfigArgs
            {
                Kind = "Cluster",
                ApiVersion = "kind.x-k8s.io/v1alpha4",
                Nodes = {
                    new ClusterKindConfigNodeArgs
                    {
                        Role = "control-plane"
                    },
                    new ClusterKindConfigNodeArgs
                    {
                        Role = "worker"
                    }
                }
            }
        });
        this.Name = cluster.Name;
    }

    [Output]
    public Output<string> Name { get; set; }
}
