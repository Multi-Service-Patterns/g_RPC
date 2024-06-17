// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.20.3
// source: music_service.proto

package music_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MusicService_GetUser_FullMethodName        = "/music.MusicService/GetUser"
	MusicService_CreateUser_FullMethodName     = "/music.MusicService/CreateUser"
	MusicService_UpdateUser_FullMethodName     = "/music.MusicService/UpdateUser"
	MusicService_DeleteUser_FullMethodName     = "/music.MusicService/DeleteUser"
	MusicService_GetPlaylist_FullMethodName    = "/music.MusicService/GetPlaylist"
	MusicService_CreatePlaylist_FullMethodName = "/music.MusicService/CreatePlaylist"
	MusicService_UpdatePlaylist_FullMethodName = "/music.MusicService/UpdatePlaylist"
	MusicService_DeletePlaylist_FullMethodName = "/music.MusicService/DeletePlaylist"
	MusicService_GetSong_FullMethodName        = "/music.MusicService/GetSong"
	MusicService_CreateSong_FullMethodName     = "/music.MusicService/CreateSong"
	MusicService_UpdateSong_FullMethodName     = "/music.MusicService/UpdateSong"
	MusicService_DeleteSong_FullMethodName     = "/music.MusicService/DeleteSong"
)

// MusicServiceClient is the client API for MusicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MusicServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*Empty, error)
	GetPlaylist(ctx context.Context, in *GetPlaylistRequest, opts ...grpc.CallOption) (*Playlist, error)
	CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error)
	UpdatePlaylist(ctx context.Context, in *UpdatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error)
	DeletePlaylist(ctx context.Context, in *DeletePlaylistRequest, opts ...grpc.CallOption) (*Empty, error)
	GetSong(ctx context.Context, in *GetSongRequest, opts ...grpc.CallOption) (*Song, error)
	CreateSong(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*Song, error)
	UpdateSong(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*Song, error)
	DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*Empty, error)
}

type musicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMusicServiceClient(cc grpc.ClientConnInterface) MusicServiceClient {
	return &musicServiceClient{cc}
}

func (c *musicServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, MusicService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, MusicService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, MusicService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, MusicService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) GetPlaylist(ctx context.Context, in *GetPlaylistRequest, opts ...grpc.CallOption) (*Playlist, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Playlist)
	err := c.cc.Invoke(ctx, MusicService_GetPlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Playlist)
	err := c.cc.Invoke(ctx, MusicService_CreatePlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) UpdatePlaylist(ctx context.Context, in *UpdatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Playlist)
	err := c.cc.Invoke(ctx, MusicService_UpdatePlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) DeletePlaylist(ctx context.Context, in *DeletePlaylistRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, MusicService_DeletePlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) GetSong(ctx context.Context, in *GetSongRequest, opts ...grpc.CallOption) (*Song, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Song)
	err := c.cc.Invoke(ctx, MusicService_GetSong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) CreateSong(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*Song, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Song)
	err := c.cc.Invoke(ctx, MusicService_CreateSong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) UpdateSong(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*Song, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Song)
	err := c.cc.Invoke(ctx, MusicService_UpdateSong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, MusicService_DeleteSong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MusicServiceServer is the server API for MusicService service.
// All implementations must embed UnimplementedMusicServiceServer
// for forward compatibility
type MusicServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*Empty, error)
	GetPlaylist(context.Context, *GetPlaylistRequest) (*Playlist, error)
	CreatePlaylist(context.Context, *CreatePlaylistRequest) (*Playlist, error)
	UpdatePlaylist(context.Context, *UpdatePlaylistRequest) (*Playlist, error)
	DeletePlaylist(context.Context, *DeletePlaylistRequest) (*Empty, error)
	GetSong(context.Context, *GetSongRequest) (*Song, error)
	CreateSong(context.Context, *CreateSongRequest) (*Song, error)
	UpdateSong(context.Context, *UpdateSongRequest) (*Song, error)
	DeleteSong(context.Context, *DeleteSongRequest) (*Empty, error)
	mustEmbedUnimplementedMusicServiceServer()
}

// UnimplementedMusicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMusicServiceServer struct {
}

func (UnimplementedMusicServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedMusicServiceServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMusicServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedMusicServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedMusicServiceServer) GetPlaylist(context.Context, *GetPlaylistRequest) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylist not implemented")
}
func (UnimplementedMusicServiceServer) CreatePlaylist(context.Context, *CreatePlaylistRequest) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlaylist not implemented")
}
func (UnimplementedMusicServiceServer) UpdatePlaylist(context.Context, *UpdatePlaylistRequest) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlaylist not implemented")
}
func (UnimplementedMusicServiceServer) DeletePlaylist(context.Context, *DeletePlaylistRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlaylist not implemented")
}
func (UnimplementedMusicServiceServer) GetSong(context.Context, *GetSongRequest) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSong not implemented")
}
func (UnimplementedMusicServiceServer) CreateSong(context.Context, *CreateSongRequest) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSong not implemented")
}
func (UnimplementedMusicServiceServer) UpdateSong(context.Context, *UpdateSongRequest) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSong not implemented")
}
func (UnimplementedMusicServiceServer) DeleteSong(context.Context, *DeleteSongRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedMusicServiceServer) mustEmbedUnimplementedMusicServiceServer() {}

// UnsafeMusicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MusicServiceServer will
// result in compilation errors.
type UnsafeMusicServiceServer interface {
	mustEmbedUnimplementedMusicServiceServer()
}

func RegisterMusicServiceServer(s grpc.ServiceRegistrar, srv MusicServiceServer) {
	s.RegisterService(&MusicService_ServiceDesc, srv)
}

func _MusicService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_GetPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).GetPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_GetPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).GetPlaylist(ctx, req.(*GetPlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_CreatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).CreatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_CreatePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).CreatePlaylist(ctx, req.(*CreatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_UpdatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).UpdatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_UpdatePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).UpdatePlaylist(ctx, req.(*UpdatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_DeletePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).DeletePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_DeletePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).DeletePlaylist(ctx, req.(*DeletePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_GetSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).GetSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_GetSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).GetSong(ctx, req.(*GetSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_CreateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).CreateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_CreateSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).CreateSong(ctx, req.(*CreateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_UpdateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).UpdateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_UpdateSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).UpdateSong(ctx, req.(*UpdateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_DeleteSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).DeleteSong(ctx, req.(*DeleteSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MusicService_ServiceDesc is the grpc.ServiceDesc for MusicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MusicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "music.MusicService",
	HandlerType: (*MusicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _MusicService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _MusicService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _MusicService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _MusicService_DeleteUser_Handler,
		},
		{
			MethodName: "GetPlaylist",
			Handler:    _MusicService_GetPlaylist_Handler,
		},
		{
			MethodName: "CreatePlaylist",
			Handler:    _MusicService_CreatePlaylist_Handler,
		},
		{
			MethodName: "UpdatePlaylist",
			Handler:    _MusicService_UpdatePlaylist_Handler,
		},
		{
			MethodName: "DeletePlaylist",
			Handler:    _MusicService_DeletePlaylist_Handler,
		},
		{
			MethodName: "GetSong",
			Handler:    _MusicService_GetSong_Handler,
		},
		{
			MethodName: "CreateSong",
			Handler:    _MusicService_CreateSong_Handler,
		},
		{
			MethodName: "UpdateSong",
			Handler:    _MusicService_UpdateSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _MusicService_DeleteSong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "music_service.proto",
}
