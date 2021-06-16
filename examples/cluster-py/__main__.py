import pulumi
import pulumi_kind as kind

cluster = kind.Cluster(
    "pulumi-cluster",
    kind_config={
        "kind": "Cluster",
        "apiVersion": "kind.x-k8s.io/v1alpha4",
        "nodes": [{"role": "control-plane"}, {"role": "worker"}],
    },
    wait_for_ready=True,
)


pulumi.export("clusterName", cluster.name)
