// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BootstrapArgs, BootstrapState } from "./bootstrap";
export type Bootstrap = import("./bootstrap").Bootstrap;
export const Bootstrap: typeof import("./bootstrap").Bootstrap = null as any;
utilities.lazyLoad(exports, ["Bootstrap"], () => require("./bootstrap"));

export { ConfigurationArgs, ConfigurationResult, ConfigurationOutputArgs } from "./configuration";
export const configuration: typeof import("./configuration").configuration = null as any;
export const configurationOutput: typeof import("./configuration").configurationOutput = null as any;
utilities.lazyLoad(exports, ["configuration","configurationOutput"], () => require("./configuration"));

export { ConfigurationApplyArgs, ConfigurationApplyState } from "./configurationApply";
export type ConfigurationApply = import("./configurationApply").ConfigurationApply;
export const ConfigurationApply: typeof import("./configurationApply").ConfigurationApply = null as any;
utilities.lazyLoad(exports, ["ConfigurationApply"], () => require("./configurationApply"));

export { SecretsArgs, SecretsState } from "./secrets";
export type Secrets = import("./secrets").Secrets;
export const Secrets: typeof import("./secrets").Secrets = null as any;
utilities.lazyLoad(exports, ["Secrets"], () => require("./secrets"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "talos:machine/bootstrap:Bootstrap":
                return new Bootstrap(name, <any>undefined, { urn })
            case "talos:machine/configurationApply:ConfigurationApply":
                return new ConfigurationApply(name, <any>undefined, { urn })
            case "talos:machine/secrets:Secrets":
                return new Secrets(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("talos", "machine/bootstrap", _module)
pulumi.runtime.registerResourceModule("talos", "machine/configurationApply", _module)
pulumi.runtime.registerResourceModule("talos", "machine/secrets", _module)