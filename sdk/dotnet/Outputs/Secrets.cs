// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Talos.Outputs
{

    /// <summary>
    /// Talos Secrets type
    /// </summary>
    [OutputType]
    public sealed class Secrets
    {
        public readonly string? AESCBCEncryptionSecret;
        public readonly string? BootstrapToken;

        [OutputConstructor]
        private Secrets(
            string? AESCBCEncryptionSecret,

            string? BootstrapToken)
        {
            this.AESCBCEncryptionSecret = AESCBCEncryptionSecret;
            this.BootstrapToken = BootstrapToken;
        }
    }
}