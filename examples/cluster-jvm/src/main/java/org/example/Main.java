package org.example;

import io.pulumi.deployment.Deployment;

public class Main {
    public static void main(String[] args) {
        var code = Deployment.runAsyncStack(MyStack.class).join();
        System.out.println("Example ended with: " + code);
        System.exit(code);
    }
}
