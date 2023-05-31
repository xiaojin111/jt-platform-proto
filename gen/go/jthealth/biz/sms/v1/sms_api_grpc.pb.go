// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package smsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SmsAPIClient is the client API for SmsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsAPIClient interface {
	// JingTongSendVerificationCode 发送手机验证码.
	JingTongSendVerificationCode(ctx context.Context, in *SendVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error)
	// JingTongSendGetLastestVerificationCode 获取最新的短信.
	JingTongSendGetLastestVerificationCode(ctx context.Context, in *GetLastestVerificationCodeRequest, opts ...grpc.CallOption) (*GetLastestVerificationCodeResponse, error)
}

type smsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsAPIClient(cc grpc.ClientConnInterface) SmsAPIClient {
	return &smsAPIClient{cc}
}

func (c *smsAPIClient) JingTongSendVerificationCode(ctx context.Context, in *SendVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error) {
	out := new(SendVerificationCodeResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.sms.v1.SmsAPI/JingTongSendVerificationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsAPIClient) JingTongSendGetLastestVerificationCode(ctx context.Context, in *GetLastestVerificationCodeRequest, opts ...grpc.CallOption) (*GetLastestVerificationCodeResponse, error) {
	out := new(GetLastestVerificationCodeResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.sms.v1.SmsAPI/JingTongSendGetLastestVerificationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsAPIServer is the server API for SmsAPI service.
// All implementations must embed UnimplementedSmsAPIServer
// for forward compatibility
type SmsAPIServer interface {
	// JingTongSendVerificationCode 发送手机验证码.
	JingTongSendVerificationCode(context.Context, *SendVerificationCodeRequest) (*SendVerificationCodeResponse, error)
	// JingTongSendGetLastestVerificationCode 获取最新的短信.
	JingTongSendGetLastestVerificationCode(context.Context, *GetLastestVerificationCodeRequest) (*GetLastestVerificationCodeResponse, error)
	mustEmbedUnimplementedSmsAPIServer()
}

// UnimplementedSmsAPIServer must be embedded to have forward compatible implementations.
type UnimplementedSmsAPIServer struct {
}

func (UnimplementedSmsAPIServer) JingTongSendVerificationCode(context.Context, *SendVerificationCodeRequest) (*SendVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JingTongSendVerificationCode not implemented")
}
func (UnimplementedSmsAPIServer) JingTongSendGetLastestVerificationCode(context.Context, *GetLastestVerificationCodeRequest) (*GetLastestVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JingTongSendGetLastestVerificationCode not implemented")
}
func (UnimplementedSmsAPIServer) mustEmbedUnimplementedSmsAPIServer() {}

// UnsafeSmsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsAPIServer will
// result in compilation errors.
type UnsafeSmsAPIServer interface {
	mustEmbedUnimplementedSmsAPIServer()
}

func RegisterSmsAPIServer(s *grpc.Server, srv SmsAPIServer) {
	s.RegisterService(&_SmsAPI_serviceDesc, srv)
}

func _SmsAPI_JingTongSendVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsAPIServer).JingTongSendVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.sms.v1.SmsAPI/JingTongSendVerificationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsAPIServer).JingTongSendVerificationCode(ctx, req.(*SendVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsAPI_JingTongSendGetLastestVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastestVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsAPIServer).JingTongSendGetLastestVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.sms.v1.SmsAPI/JingTongSendGetLastestVerificationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsAPIServer).JingTongSendGetLastestVerificationCode(ctx, req.(*GetLastestVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmsAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jthealth.biz.sms.v1.SmsAPI",
	HandlerType: (*SmsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JingTongSendVerificationCode",
			Handler:    _SmsAPI_JingTongSendVerificationCode_Handler,
		},
		{
			MethodName: "JingTongSendGetLastestVerificationCode",
			Handler:    _SmsAPI_JingTongSendGetLastestVerificationCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jthealth/biz/sms/v1/sms_api.proto",
}
