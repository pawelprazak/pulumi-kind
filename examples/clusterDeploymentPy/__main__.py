import pulumi
from pulumi_kubernetes.apps.v1 import Deployment

import pulumi_kind as kind

cluster = kind.Cluster(
    "pulumi-cluster",
    kind_config={
        "kind": "Cluster",
        "apiVersion": "kind.x-k8s.io/v1alpha4",
        "nodes": [
            {
                "role": "control-plane"
            },{
                "role": "worker"
            }
        ]
    },
    wait_for_ready=True
)

app_labels = { "app": "nginx" }

deployment = Deployment(
    "nginx",
    spec={
        "selector": { "match_labels": app_labels },
        "replicas": 1,
        "template": {
            "metadata": { "labels": app_labels },
            "spec": { "containers": [{ "name": "nginx", "image": "nginx" }] }
        }
    },
    opts=pulumi.ResourceOptions(
        depends_on=[cluster]
    )
)
    

pulumi.export("clusterName", cluster.name)