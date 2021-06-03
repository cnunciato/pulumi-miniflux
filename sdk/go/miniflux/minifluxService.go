// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package miniflux

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MinifluxService struct {
	pulumi.ResourceState

	// The URL of the Miniflux service.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
}

// NewMinifluxService registers a new resource with the given unique name, arguments, and options.
func NewMinifluxService(ctx *pulumi.Context,
	name string, args *MinifluxServiceArgs, opts ...pulumi.ResourceOption) (*MinifluxService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminPassword == nil {
		return nil, errors.New("invalid value for required argument 'AdminPassword'")
	}
	if args.DbPassword == nil {
		return nil, errors.New("invalid value for required argument 'DbPassword'")
	}
	if args.AdminUsername == nil {
		args.AdminUsername = pulumi.StringPtr("admin")
	}
	if args.DbName == nil {
		args.DbName = pulumi.StringPtr("miniflux")
	}
	if args.DbUsername == nil {
		args.DbUsername = pulumi.StringPtr("miniflux")
	}
	var resource MinifluxService
	err := ctx.RegisterRemoteComponentResource("miniflux:index:MinifluxService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type minifluxServiceArgs struct {
	// The Miniflux administrator's password.
	AdminPassword string `pulumi:"adminPassword"`
	// The username to use for the Miniflux service administrator account.
	AdminUsername *string `pulumi:"adminUsername"`
	// The name of the PostgreSQL database to be used by Miniflux.
	DbName *string `pulumi:"dbName"`
	// The PostgreSQL user's password.
	DbPassword string `pulumi:"dbPassword"`
	// The username of the PostgreSQL account to be used by the Miniflux service.
	DbUsername *string `pulumi:"dbUsername"`
}

// The set of arguments for constructing a MinifluxService resource.
type MinifluxServiceArgs struct {
	// The Miniflux administrator's password.
	AdminPassword pulumi.StringInput
	// The username to use for the Miniflux service administrator account.
	AdminUsername pulumi.StringPtrInput
	// The name of the PostgreSQL database to be used by Miniflux.
	DbName pulumi.StringPtrInput
	// The PostgreSQL user's password.
	DbPassword pulumi.StringInput
	// The username of the PostgreSQL account to be used by the Miniflux service.
	DbUsername pulumi.StringPtrInput
}

func (MinifluxServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*minifluxServiceArgs)(nil)).Elem()
}

type MinifluxServiceInput interface {
	pulumi.Input

	ToMinifluxServiceOutput() MinifluxServiceOutput
	ToMinifluxServiceOutputWithContext(ctx context.Context) MinifluxServiceOutput
}

func (*MinifluxService) ElementType() reflect.Type {
	return reflect.TypeOf((*MinifluxService)(nil))
}

func (i *MinifluxService) ToMinifluxServiceOutput() MinifluxServiceOutput {
	return i.ToMinifluxServiceOutputWithContext(context.Background())
}

func (i *MinifluxService) ToMinifluxServiceOutputWithContext(ctx context.Context) MinifluxServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MinifluxServiceOutput)
}

func (i *MinifluxService) ToMinifluxServicePtrOutput() MinifluxServicePtrOutput {
	return i.ToMinifluxServicePtrOutputWithContext(context.Background())
}

func (i *MinifluxService) ToMinifluxServicePtrOutputWithContext(ctx context.Context) MinifluxServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MinifluxServicePtrOutput)
}

type MinifluxServicePtrInput interface {
	pulumi.Input

	ToMinifluxServicePtrOutput() MinifluxServicePtrOutput
	ToMinifluxServicePtrOutputWithContext(ctx context.Context) MinifluxServicePtrOutput
}

type minifluxServicePtrType MinifluxServiceArgs

func (*minifluxServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MinifluxService)(nil))
}

func (i *minifluxServicePtrType) ToMinifluxServicePtrOutput() MinifluxServicePtrOutput {
	return i.ToMinifluxServicePtrOutputWithContext(context.Background())
}

func (i *minifluxServicePtrType) ToMinifluxServicePtrOutputWithContext(ctx context.Context) MinifluxServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MinifluxServicePtrOutput)
}

// MinifluxServiceArrayInput is an input type that accepts MinifluxServiceArray and MinifluxServiceArrayOutput values.
// You can construct a concrete instance of `MinifluxServiceArrayInput` via:
//
//          MinifluxServiceArray{ MinifluxServiceArgs{...} }
type MinifluxServiceArrayInput interface {
	pulumi.Input

	ToMinifluxServiceArrayOutput() MinifluxServiceArrayOutput
	ToMinifluxServiceArrayOutputWithContext(context.Context) MinifluxServiceArrayOutput
}

type MinifluxServiceArray []MinifluxServiceInput

func (MinifluxServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MinifluxService)(nil))
}

func (i MinifluxServiceArray) ToMinifluxServiceArrayOutput() MinifluxServiceArrayOutput {
	return i.ToMinifluxServiceArrayOutputWithContext(context.Background())
}

func (i MinifluxServiceArray) ToMinifluxServiceArrayOutputWithContext(ctx context.Context) MinifluxServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MinifluxServiceArrayOutput)
}

// MinifluxServiceMapInput is an input type that accepts MinifluxServiceMap and MinifluxServiceMapOutput values.
// You can construct a concrete instance of `MinifluxServiceMapInput` via:
//
//          MinifluxServiceMap{ "key": MinifluxServiceArgs{...} }
type MinifluxServiceMapInput interface {
	pulumi.Input

	ToMinifluxServiceMapOutput() MinifluxServiceMapOutput
	ToMinifluxServiceMapOutputWithContext(context.Context) MinifluxServiceMapOutput
}

type MinifluxServiceMap map[string]MinifluxServiceInput

func (MinifluxServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MinifluxService)(nil))
}

func (i MinifluxServiceMap) ToMinifluxServiceMapOutput() MinifluxServiceMapOutput {
	return i.ToMinifluxServiceMapOutputWithContext(context.Background())
}

func (i MinifluxServiceMap) ToMinifluxServiceMapOutputWithContext(ctx context.Context) MinifluxServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MinifluxServiceMapOutput)
}

type MinifluxServiceOutput struct {
	*pulumi.OutputState
}

func (MinifluxServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MinifluxService)(nil))
}

func (o MinifluxServiceOutput) ToMinifluxServiceOutput() MinifluxServiceOutput {
	return o
}

func (o MinifluxServiceOutput) ToMinifluxServiceOutputWithContext(ctx context.Context) MinifluxServiceOutput {
	return o
}

func (o MinifluxServiceOutput) ToMinifluxServicePtrOutput() MinifluxServicePtrOutput {
	return o.ToMinifluxServicePtrOutputWithContext(context.Background())
}

func (o MinifluxServiceOutput) ToMinifluxServicePtrOutputWithContext(ctx context.Context) MinifluxServicePtrOutput {
	return o.ApplyT(func(v MinifluxService) *MinifluxService {
		return &v
	}).(MinifluxServicePtrOutput)
}

type MinifluxServicePtrOutput struct {
	*pulumi.OutputState
}

func (MinifluxServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MinifluxService)(nil))
}

func (o MinifluxServicePtrOutput) ToMinifluxServicePtrOutput() MinifluxServicePtrOutput {
	return o
}

func (o MinifluxServicePtrOutput) ToMinifluxServicePtrOutputWithContext(ctx context.Context) MinifluxServicePtrOutput {
	return o
}

type MinifluxServiceArrayOutput struct{ *pulumi.OutputState }

func (MinifluxServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MinifluxService)(nil))
}

func (o MinifluxServiceArrayOutput) ToMinifluxServiceArrayOutput() MinifluxServiceArrayOutput {
	return o
}

func (o MinifluxServiceArrayOutput) ToMinifluxServiceArrayOutputWithContext(ctx context.Context) MinifluxServiceArrayOutput {
	return o
}

func (o MinifluxServiceArrayOutput) Index(i pulumi.IntInput) MinifluxServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MinifluxService {
		return vs[0].([]MinifluxService)[vs[1].(int)]
	}).(MinifluxServiceOutput)
}

type MinifluxServiceMapOutput struct{ *pulumi.OutputState }

func (MinifluxServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MinifluxService)(nil))
}

func (o MinifluxServiceMapOutput) ToMinifluxServiceMapOutput() MinifluxServiceMapOutput {
	return o
}

func (o MinifluxServiceMapOutput) ToMinifluxServiceMapOutputWithContext(ctx context.Context) MinifluxServiceMapOutput {
	return o
}

func (o MinifluxServiceMapOutput) MapIndex(k pulumi.StringInput) MinifluxServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MinifluxService {
		return vs[0].(map[string]MinifluxService)[vs[1].(string)]
	}).(MinifluxServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(MinifluxServiceOutput{})
	pulumi.RegisterOutputType(MinifluxServicePtrOutput{})
	pulumi.RegisterOutputType(MinifluxServiceArrayOutput{})
	pulumi.RegisterOutputType(MinifluxServiceMapOutput{})
}
