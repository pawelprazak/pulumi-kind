using Pulumi;
using Pulumi.Kubernetes.Apps.V1;
using Pulumi.Kubernetes.Types.Inputs.Core.V1;
using Pulumi.Kubernetes.Types.Inputs.Apps.V1;
using Pulumi.Kubernetes.Types.Inputs.Meta.V1;
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

        var appLabels = new InputMap<string>
        {
            { "app", "nginx" }
        };

        var deployment = new Pulumi.Kubernetes.Apps.V1.Deployment("nginx", new DeploymentArgs
        {
            Spec = new DeploymentSpecArgs
            {
                Selector = new LabelSelectorArgs
                {
                    MatchLabels = appLabels
                },
                Replicas = 1,
                Template = new PodTemplateSpecArgs
                {
                    Metadata = new ObjectMetaArgs
                    {
                        Labels = appLabels
                    },
                    Spec = new PodSpecArgs
                    {
                        Containers =
                        {
                            new ContainerArgs
                            {
                                Name = "nginx",
                                Image = "nginx",
                                Ports =
                                {
                                    new ContainerPortArgs
                                    {
                                        ContainerPortValue = 80
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        new CustomResourceOptions
        {
            DependsOn = { cluster }
        });
    }

    [Output]
    public Output<string> Name { get; set; }
}