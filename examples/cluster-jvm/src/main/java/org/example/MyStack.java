package org.example;

import io.pulumi.Config;
import io.pulumi.Stack;
import io.pulumi.core.Output;
import io.pulumi.core.internal.annotations.OutputExport;
import io.pulumi.kubernetes.core.v1.Namespace;
import io.pulumi.kubernetes.core.v1.inputs.NamespaceArgs;
import io.pulumi.kubernetes.meta.v1.inputs.ObjectMetaArgs;
import io.pulumi.kubernetes.meta.v1.outputs.ObjectMeta;

import java.nio.charset.StandardCharsets;
import java.util.Base64;
import java.util.Map;

import static java.util.stream.Collectors.toMap;

public class MyStack extends Stack {

    @OutputExport(type = String.class)
    public final Output<String> clusterName;

    public MyStack() {
        var cluster = new Cluster("pulumi-cluster-alone-jvm", ClusterArgs.builder()
                .setKindConfig(ClusterKindConfigArgs.builder()
                        .setKind("Cluster")
                        .setApiVersion("kind.x-k8s.io/v1alpha4")
                        .setNodes(List.of(
                                ClusterKindConfigNodeArgs.builder().setRole("control-plane").build(),
                                ClusterKindConfigNodeArgs.builder().setRole("worker").build()
                        ))
                        .build())
                .setWaitForReady(true)
                .build(),
                null
        );

        this.clusterName = cluster.getName();
    }
}
