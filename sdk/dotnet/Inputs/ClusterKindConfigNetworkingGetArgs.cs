// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kind.Inputs
{

    public sealed class ClusterKindConfigNetworkingGetArgs : Pulumi.ResourceArgs
    {
        [Input("apiServerAddress")]
        public Input<string>? ApiServerAddress { get; set; }

        [Input("apiServerPort")]
        public Input<int>? ApiServerPort { get; set; }

        [Input("disableDefaultCni")]
        public Input<bool>? DisableDefaultCni { get; set; }

        [Input("ipFamily")]
        public Input<string>? IpFamily { get; set; }

        [Input("kubeProxyMode")]
        public Input<string>? KubeProxyMode { get; set; }

        [Input("podSubnet")]
        public Input<string>? PodSubnet { get; set; }

        [Input("serviceSubnet")]
        public Input<string>? ServiceSubnet { get; set; }

        public ClusterKindConfigNetworkingGetArgs()
        {
        }
    }
}
