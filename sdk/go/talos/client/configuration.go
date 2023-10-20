// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package client

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Generate client configuration for a Talos cluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/client"
//	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/machine"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			thisSecrets, err := machine.NewSecrets(ctx, "thisSecrets", nil)
//			if err != nil {
//				return err
//			}
//			_ = client.ConfigurationOutput(ctx, client.ConfigurationOutputArgs{
//				ClusterName:         pulumi.String("example-cluster"),
//				ClientConfiguration: thisSecrets.ClientConfiguration,
//				Nodes: pulumi.StringArray{
//					pulumi.String("10.5.0.2"),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func Configuration(ctx *pulumi.Context, args *ConfigurationArgs, opts ...pulumi.InvokeOption) (*ConfigurationResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv ConfigurationResult
	err := ctx.Invoke("talos:client/configuration:Configuration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Configuration.
type ConfigurationArgs struct {
	// The client configuration data
	ClientConfiguration ConfigurationClientConfiguration `pulumi:"clientConfiguration"`
	// The name of the cluster in the generated config
	ClusterName string `pulumi:"clusterName"`
	// endpoints to set in the generated config
	Endpoints []string `pulumi:"endpoints"`
	// nodes to set in the generated config
	Nodes []string `pulumi:"nodes"`
}

// A collection of values returned by Configuration.
type ConfigurationResult struct {
	// The client configuration data
	ClientConfiguration ConfigurationClientConfiguration `pulumi:"clientConfiguration"`
	// The name of the cluster in the generated config
	ClusterName string `pulumi:"clusterName"`
	// endpoints to set in the generated config
	Endpoints []string `pulumi:"endpoints"`
	// The ID of this resource
	Id string `pulumi:"id"`
	// nodes to set in the generated config
	Nodes []string `pulumi:"nodes"`
	// The generated client configuration
	TalosConfig string `pulumi:"talosConfig"`
}

func ConfigurationOutput(ctx *pulumi.Context, args ConfigurationOutputArgs, opts ...pulumi.InvokeOption) ConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ConfigurationResult, error) {
			args := v.(ConfigurationArgs)
			r, err := Configuration(ctx, &args, opts...)
			var s ConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ConfigurationResultOutput)
}

// A collection of arguments for invoking Configuration.
type ConfigurationOutputArgs struct {
	// The client configuration data
	ClientConfiguration ConfigurationClientConfigurationInput `pulumi:"clientConfiguration"`
	// The name of the cluster in the generated config
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// endpoints to set in the generated config
	Endpoints pulumi.StringArrayInput `pulumi:"endpoints"`
	// nodes to set in the generated config
	Nodes pulumi.StringArrayInput `pulumi:"nodes"`
}

func (ConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationArgs)(nil)).Elem()
}

// A collection of values returned by Configuration.
type ConfigurationResultOutput struct{ *pulumi.OutputState }

func (ConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationResult)(nil)).Elem()
}

func (o ConfigurationResultOutput) ToConfigurationResultOutput() ConfigurationResultOutput {
	return o
}

func (o ConfigurationResultOutput) ToConfigurationResultOutputWithContext(ctx context.Context) ConfigurationResultOutput {
	return o
}

// The client configuration data
func (o ConfigurationResultOutput) ClientConfiguration() ConfigurationClientConfigurationOutput {
	return o.ApplyT(func(v ConfigurationResult) ConfigurationClientConfiguration { return v.ClientConfiguration }).(ConfigurationClientConfigurationOutput)
}

// The name of the cluster in the generated config
func (o ConfigurationResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// endpoints to set in the generated config
func (o ConfigurationResultOutput) Endpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConfigurationResult) []string { return v.Endpoints }).(pulumi.StringArrayOutput)
}

// The ID of this resource
func (o ConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// nodes to set in the generated config
func (o ConfigurationResultOutput) Nodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConfigurationResult) []string { return v.Nodes }).(pulumi.StringArrayOutput)
}

// The generated client configuration
func (o ConfigurationResultOutput) TalosConfig() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationResult) string { return v.TalosConfig }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationResultOutput{})
}