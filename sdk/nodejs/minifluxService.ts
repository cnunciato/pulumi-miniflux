// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class MinifluxService extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'miniflux:index:MinifluxService';

    /**
     * Returns true if the given object is an instance of MinifluxService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MinifluxService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MinifluxService.__pulumiType;
    }

    /**
     * The URL of the Miniflux service.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;

    /**
     * Create a MinifluxService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MinifluxServiceArgs, opts?: pulumi.ComponentResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.adminPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminPassword'");
            }
            if ((!args || args.dbPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbPassword'");
            }
            inputs["adminPassword"] = args ? args.adminPassword : undefined;
            inputs["adminUsername"] = (args ? args.adminUsername : undefined) ?? "admin";
            inputs["dbName"] = (args ? args.dbName : undefined) ?? "miniflux";
            inputs["dbPassword"] = args ? args.dbPassword : undefined;
            inputs["dbUsername"] = (args ? args.dbUsername : undefined) ?? "miniflux";
            inputs["endpoint"] = undefined /*out*/;
        } else {
            inputs["endpoint"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MinifluxService.__pulumiType, name, inputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a MinifluxService resource.
 */
export interface MinifluxServiceArgs {
    /**
     * The Miniflux administrator's password.
     */
    adminPassword: pulumi.Input<string>;
    /**
     * The username to use for the Miniflux service administrator account.
     */
    adminUsername?: pulumi.Input<string>;
    /**
     * The name of the PostgreSQL database to be used by Miniflux.
     */
    dbName?: pulumi.Input<string>;
    /**
     * The PostgreSQL user's password.
     */
    dbPassword: pulumi.Input<string>;
    /**
     * The username of the PostgreSQL account to be used by the Miniflux service.
     */
    dbUsername?: pulumi.Input<string>;
}
