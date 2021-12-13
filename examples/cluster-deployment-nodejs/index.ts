// @ts-ignore because it's provided by yarn link @pawelprazak/pulumi-kind
import * as kind from "@pawelprazak/pulumi-kind";
import * as k8s from "@pulumi/kubernetes";

let cluster = new kind.Cluster("pulumi-cluster-deployment-nodejs", {
    kindConfig: {
        kind: "Cluster",
        apiVersion: "kind.x-k8s.io/v1alpha4",
        nodes: [
            {
                role: "control-plane"
            },{
                role: "worker"
            }
        ]
    },
    waitForReady: true
});

const appLabels = { app: "nginx" };
const deployment = new k8s.apps.v1.Deployment("nginx",{
    spec: {
        selector: { matchLabels: appLabels },
        replicas: 1,
        template: {
            metadata: { labels: appLabels },
            spec: { containers: [{ name: "nginx", image: "nginx" }] }
        }
    }
},
{
    dependsOn: [cluster]
});

// noinspection JSUnusedGlobalSymbols
export let clusterName = cluster.name;