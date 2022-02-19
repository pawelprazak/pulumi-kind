// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kind

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Cluster
//
// Provides a Kind cluster resource. This can be used to create and delete Kind
// clusters. It does NOT support modification to an existing kind cluster.
type Cluster struct {
	pulumi.CustomResourceState

	// Client certificate for authenticating to cluster.
	ClientCertificate pulumi.StringOutput `pulumi:"clientCertificate"`
	// Client key for authenticating to cluster.
	ClientKey pulumi.StringOutput `pulumi:"clientKey"`
	// Client verifies the server certificate with this CA cert.
	ClusterCaCertificate pulumi.StringOutput `pulumi:"clusterCaCertificate"`
	// Kubernetes APIServer endpoint.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The kindConfig that kind will use.
	KindConfig ClusterKindConfigPtrOutput `pulumi:"kindConfig"`
	// Kubeconfig set after the the cluster is created.
	Kubeconfig pulumi.StringOutput `pulumi:"kubeconfig"`
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath pulumi.StringOutput `pulumi:"kubeconfigPath"`
	// The kind name that is given to the created cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The nodeImage that kind will use (ex: kindest/node:v1.15.3).
	NodeImage pulumi.StringOutput `pulumi:"nodeImage"`
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false.
	WaitForReady pulumi.BoolPtrOutput `pulumi:"waitForReady"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		args = &ClusterArgs{}
	}

	var resource Cluster
	err := ctx.RegisterResource("kind:index/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("kind:index/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Client certificate for authenticating to cluster.
	ClientCertificate *string `pulumi:"clientCertificate"`
	// Client key for authenticating to cluster.
	ClientKey *string `pulumi:"clientKey"`
	// Client verifies the server certificate with this CA cert.
	ClusterCaCertificate *string `pulumi:"clusterCaCertificate"`
	// Kubernetes APIServer endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The kindConfig that kind will use.
	KindConfig *ClusterKindConfig `pulumi:"kindConfig"`
	// Kubeconfig set after the the cluster is created.
	Kubeconfig *string `pulumi:"kubeconfig"`
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath *string `pulumi:"kubeconfigPath"`
	// The kind name that is given to the created cluster.
	Name *string `pulumi:"name"`
	// The nodeImage that kind will use (ex: kindest/node:v1.15.3).
	NodeImage *string `pulumi:"nodeImage"`
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false.
	WaitForReady *bool `pulumi:"waitForReady"`
}

type ClusterState struct {
	// Client certificate for authenticating to cluster.
	ClientCertificate pulumi.StringPtrInput
	// Client key for authenticating to cluster.
	ClientKey pulumi.StringPtrInput
	// Client verifies the server certificate with this CA cert.
	ClusterCaCertificate pulumi.StringPtrInput
	// Kubernetes APIServer endpoint.
	Endpoint pulumi.StringPtrInput
	// The kindConfig that kind will use.
	KindConfig ClusterKindConfigPtrInput
	// Kubeconfig set after the the cluster is created.
	Kubeconfig pulumi.StringPtrInput
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath pulumi.StringPtrInput
	// The kind name that is given to the created cluster.
	Name pulumi.StringPtrInput
	// The nodeImage that kind will use (ex: kindest/node:v1.15.3).
	NodeImage pulumi.StringPtrInput
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false.
	WaitForReady pulumi.BoolPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The kindConfig that kind will use.
	KindConfig *ClusterKindConfig `pulumi:"kindConfig"`
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath *string `pulumi:"kubeconfigPath"`
	// The kind name that is given to the created cluster.
	Name *string `pulumi:"name"`
	// The nodeImage that kind will use (ex: kindest/node:v1.15.3).
	NodeImage *string `pulumi:"nodeImage"`
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false.
	WaitForReady *bool `pulumi:"waitForReady"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The kindConfig that kind will use.
	KindConfig ClusterKindConfigPtrInput
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath pulumi.StringPtrInput
	// The kind name that is given to the created cluster.
	Name pulumi.StringPtrInput
	// The nodeImage that kind will use (ex: kindest/node:v1.15.3).
	NodeImage pulumi.StringPtrInput
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false.
	WaitForReady pulumi.BoolPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

func (i *Cluster) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

type ClusterPtrInput interface {
	pulumi.Input

	ToClusterPtrOutput() ClusterPtrOutput
	ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput
}

type clusterPtrType ClusterArgs

func (*clusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (i *clusterPtrType) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *clusterPtrType) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//          ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Cluster)(nil))
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//          ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Cluster)(nil))
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct {
	*pulumi.OutputState
}

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o.ToClusterPtrOutputWithContext(context.Background())
}

func (o ClusterOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o.ApplyT(func(v Cluster) *Cluster {
		return &v
	}).(ClusterPtrOutput)
}

type ClusterPtrOutput struct {
	*pulumi.OutputState
}

func (ClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (o ClusterPtrOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o
}

func (o ClusterPtrOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Cluster)(nil))
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].([]Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Cluster)(nil))
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].(map[string]Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterPtrOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
