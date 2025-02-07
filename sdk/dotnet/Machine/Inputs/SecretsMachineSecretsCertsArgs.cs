// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Talos.Machine.Inputs
{

    public sealed class SecretsMachineSecretsCertsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate and key pair
        /// </summary>
        [Input("etcd")]
        public Input<Inputs.SecretsMachineSecretsCertsEtcdArgs>? Etcd { get; set; }

        /// <summary>
        /// The certificate and key pair
        /// </summary>
        [Input("k8s")]
        public Input<Inputs.SecretsMachineSecretsCertsK8sArgs>? K8s { get; set; }

        /// <summary>
        /// The certificate and key pair
        /// </summary>
        [Input("k8sAggregator")]
        public Input<Inputs.SecretsMachineSecretsCertsK8sAggregatorArgs>? K8sAggregator { get; set; }

        /// <summary>
        /// The service account secrets
        /// </summary>
        [Input("k8sServiceaccount")]
        public Input<Inputs.SecretsMachineSecretsCertsK8sServiceaccountArgs>? K8sServiceaccount { get; set; }

        /// <summary>
        /// The certificate and key pair
        /// </summary>
        [Input("os")]
        public Input<Inputs.SecretsMachineSecretsCertsOsArgs>? Os { get; set; }

        public SecretsMachineSecretsCertsArgs()
        {
        }
        public static new SecretsMachineSecretsCertsArgs Empty => new SecretsMachineSecretsCertsArgs();
    }
}
