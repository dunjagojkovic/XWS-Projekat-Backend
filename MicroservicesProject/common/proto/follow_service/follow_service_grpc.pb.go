// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: follow_service.proto

package follow

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FollowServiceClient is the client API for FollowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowServiceClient interface {
	AcceptFollow(ctx context.Context, in *AcceptFollowRequest, opts ...grpc.CallOption) (*AcceptFollowResponse, error)
	Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
	Unfollow(ctx context.Context, in *UnfollowRequest, opts ...grpc.CallOption) (*UnfollowResponse, error)
	FollowRequestRemove(ctx context.Context, in *FollowRequestRemoveRequest, opts ...grpc.CallOption) (*FollowRequestRemoveResponse, error)
	Follows(ctx context.Context, in *FollowsRequest, opts ...grpc.CallOption) (*FollowsResponse, error)
	Followers(ctx context.Context, in *FollowersRequest, opts ...grpc.CallOption) (*FollowersResponse, error)
	FollowRequests(ctx context.Context, in *FollowRequestsRequest, opts ...grpc.CallOption) (*FollowRequestsResponse, error)
	FollowerRequests(ctx context.Context, in *FollowerRequestsRequest, opts ...grpc.CallOption) (*FollowerRequestsResponse, error)
	Relationships(ctx context.Context, in *RelationshipsRequest, opts ...grpc.CallOption) (*RelationshipsResponse, error)
	GetRecommended(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ListId, error)
}

type followServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowServiceClient(cc grpc.ClientConnInterface) FollowServiceClient {
	return &followServiceClient{cc}
}

func (c *followServiceClient) AcceptFollow(ctx context.Context, in *AcceptFollowRequest, opts ...grpc.CallOption) (*AcceptFollowResponse, error) {
	out := new(AcceptFollowResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/AcceptFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	out := new(FollowResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) Unfollow(ctx context.Context, in *UnfollowRequest, opts ...grpc.CallOption) (*UnfollowResponse, error) {
	out := new(UnfollowResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/Unfollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) FollowRequestRemove(ctx context.Context, in *FollowRequestRemoveRequest, opts ...grpc.CallOption) (*FollowRequestRemoveResponse, error) {
	out := new(FollowRequestRemoveResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/FollowRequestRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) Follows(ctx context.Context, in *FollowsRequest, opts ...grpc.CallOption) (*FollowsResponse, error) {
	out := new(FollowsResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/Follows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) Followers(ctx context.Context, in *FollowersRequest, opts ...grpc.CallOption) (*FollowersResponse, error) {
	out := new(FollowersResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/Followers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) FollowRequests(ctx context.Context, in *FollowRequestsRequest, opts ...grpc.CallOption) (*FollowRequestsResponse, error) {
	out := new(FollowRequestsResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/FollowRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) FollowerRequests(ctx context.Context, in *FollowerRequestsRequest, opts ...grpc.CallOption) (*FollowerRequestsResponse, error) {
	out := new(FollowerRequestsResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/FollowerRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) Relationships(ctx context.Context, in *RelationshipsRequest, opts ...grpc.CallOption) (*RelationshipsResponse, error) {
	out := new(RelationshipsResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/Relationships", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) GetRecommended(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ListId, error) {
	out := new(ListId)
	err := c.cc.Invoke(ctx, "/follow.FollowService/GetRecommended", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServiceServer is the server API for FollowService service.
// All implementations must embed UnimplementedFollowServiceServer
// for forward compatibility
type FollowServiceServer interface {
	AcceptFollow(context.Context, *AcceptFollowRequest) (*AcceptFollowResponse, error)
	Follow(context.Context, *FollowRequest) (*FollowResponse, error)
	Unfollow(context.Context, *UnfollowRequest) (*UnfollowResponse, error)
	FollowRequestRemove(context.Context, *FollowRequestRemoveRequest) (*FollowRequestRemoveResponse, error)
	Follows(context.Context, *FollowsRequest) (*FollowsResponse, error)
	Followers(context.Context, *FollowersRequest) (*FollowersResponse, error)
	FollowRequests(context.Context, *FollowRequestsRequest) (*FollowRequestsResponse, error)
	FollowerRequests(context.Context, *FollowerRequestsRequest) (*FollowerRequestsResponse, error)
	Relationships(context.Context, *RelationshipsRequest) (*RelationshipsResponse, error)
	GetRecommended(context.Context, *Id) (*ListId, error)
	mustEmbedUnimplementedFollowServiceServer()
}

// UnimplementedFollowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowServiceServer struct {
}

func (UnimplementedFollowServiceServer) AcceptFollow(context.Context, *AcceptFollowRequest) (*AcceptFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFollow not implemented")
}
func (UnimplementedFollowServiceServer) Follow(context.Context, *FollowRequest) (*FollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedFollowServiceServer) Unfollow(context.Context, *UnfollowRequest) (*UnfollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfollow not implemented")
}
func (UnimplementedFollowServiceServer) FollowRequestRemove(context.Context, *FollowRequestRemoveRequest) (*FollowRequestRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowRequestRemove not implemented")
}
func (UnimplementedFollowServiceServer) Follows(context.Context, *FollowsRequest) (*FollowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follows not implemented")
}
func (UnimplementedFollowServiceServer) Followers(context.Context, *FollowersRequest) (*FollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Followers not implemented")
}
func (UnimplementedFollowServiceServer) FollowRequests(context.Context, *FollowRequestsRequest) (*FollowRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowRequests not implemented")
}
func (UnimplementedFollowServiceServer) FollowerRequests(context.Context, *FollowerRequestsRequest) (*FollowerRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowerRequests not implemented")
}
func (UnimplementedFollowServiceServer) Relationships(context.Context, *RelationshipsRequest) (*RelationshipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Relationships not implemented")
}
func (UnimplementedFollowServiceServer) GetRecommended(context.Context, *Id) (*ListId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommended not implemented")
}
func (UnimplementedFollowServiceServer) mustEmbedUnimplementedFollowServiceServer() {}

// UnsafeFollowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowServiceServer will
// result in compilation errors.
type UnsafeFollowServiceServer interface {
	mustEmbedUnimplementedFollowServiceServer()
}

func RegisterFollowServiceServer(s grpc.ServiceRegistrar, srv FollowServiceServer) {
	s.RegisterService(&FollowService_ServiceDesc, srv)
}

func _FollowService_AcceptFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).AcceptFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/AcceptFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).AcceptFollow(ctx, req.(*AcceptFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).Follow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/Unfollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).Unfollow(ctx, req.(*UnfollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_FollowRequestRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequestRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).FollowRequestRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/FollowRequestRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).FollowRequestRemove(ctx, req.(*FollowRequestRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_Follows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).Follows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/Follows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).Follows(ctx, req.(*FollowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_Followers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).Followers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/Followers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).Followers(ctx, req.(*FollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_FollowRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).FollowRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/FollowRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).FollowRequests(ctx, req.(*FollowRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_FollowerRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).FollowerRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/FollowerRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).FollowerRequests(ctx, req.(*FollowerRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_Relationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationshipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).Relationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/Relationships",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).Relationships(ctx, req.(*RelationshipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_GetRecommended_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).GetRecommended(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/GetRecommended",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).GetRecommended(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowService_ServiceDesc is the grpc.ServiceDesc for FollowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "follow.FollowService",
	HandlerType: (*FollowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptFollow",
			Handler:    _FollowService_AcceptFollow_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _FollowService_Follow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _FollowService_Unfollow_Handler,
		},
		{
			MethodName: "FollowRequestRemove",
			Handler:    _FollowService_FollowRequestRemove_Handler,
		},
		{
			MethodName: "Follows",
			Handler:    _FollowService_Follows_Handler,
		},
		{
			MethodName: "Followers",
			Handler:    _FollowService_Followers_Handler,
		},
		{
			MethodName: "FollowRequests",
			Handler:    _FollowService_FollowRequests_Handler,
		},
		{
			MethodName: "FollowerRequests",
			Handler:    _FollowService_FollowerRequests_Handler,
		},
		{
			MethodName: "Relationships",
			Handler:    _FollowService_Relationships_Handler,
		},
		{
			MethodName: "GetRecommended",
			Handler:    _FollowService_GetRecommended_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "follow_service.proto",
}
