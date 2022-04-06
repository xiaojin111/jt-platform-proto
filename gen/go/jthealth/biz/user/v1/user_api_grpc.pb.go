// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package userv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserAPIClient is the client API for UserAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAPIClient interface {
	// SignInByPhoneCode 手机号验证码登录.
	SignInByPhoneCode(ctx context.Context, in *SignInByPhoneCodeRequest, opts ...grpc.CallOption) (*SignInByPhoneCodeResponse, error)
	//SignInByPassWord 账号密码登录
	SignInByPassWord(ctx context.Context, in *SignInByPassWordRequest, opts ...grpc.CallOption) (*SignInByPassWordResponse, error)
	//GetUserInfosById 获取用户信息
	GetUserInfosById(ctx context.Context, in *GetUserInfosByIdRequest, opts ...grpc.CallOption) (*GetUserInfosByIdResponse, error)
	//UpdateUserInfos  更改用户信息
	UpdateUserInfos(ctx context.Context, in *UpdateUserInfosRequest, opts ...grpc.CallOption) (*UpdateUserInfosResponse, error)
	//CreatedUserProfile 创建用户档案
	CreatedUserProfile(ctx context.Context, in *CreatedUserProfileRequest, opts ...grpc.CallOption) (*CreatedUserProfileResponse, error)
	//GetUserInfosList 获取用户信息列表
	GetUserInfosList(ctx context.Context, in *GetUserInfosListRequest, opts ...grpc.CallOption) (*GetUserInfosListResponse, error)
	//----------------------应用-------------------------------
	//应用申请
	ApplyApplication(ctx context.Context, in *ApplyApplicationRequest, opts ...grpc.CallOption) (*ApplyApplicationResponse, error)
	//获取应用详情
	GetApplicationInfo(ctx context.Context, in *GetApplicationInfoRequest, opts ...grpc.CallOption) (*GetApplicationInfoResponse, error)
	//重置app_secret
	RestAppSecretById(ctx context.Context, in *RestAppSecretByIdRequest, opts ...grpc.CallOption) (*RestAppSecretByIdResponse, error)
	//应用功能申请
	ApplyFeatureAudit(ctx context.Context, in *ApplyFeatureAuditRequest, opts ...grpc.CallOption) (*ApplyFeatureAuditResponse, error)
	//获取审核列表
	GetFeatureAuditList(ctx context.Context, in *GetFeatureAuditListRequest, opts ...grpc.CallOption) (*GetFeatureAuditListResponse, error)
	//应用功能审核
	UpdateFeatureAudit(ctx context.Context, in *UpdateFeatureAuditRequest, opts ...grpc.CallOption) (*UpdateFeatureAuditResponse, error)
}

type userAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAPIClient(cc grpc.ClientConnInterface) UserAPIClient {
	return &userAPIClient{cc}
}

func (c *userAPIClient) SignInByPhoneCode(ctx context.Context, in *SignInByPhoneCodeRequest, opts ...grpc.CallOption) (*SignInByPhoneCodeResponse, error) {
	out := new(SignInByPhoneCodeResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/SignInByPhoneCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) SignInByPassWord(ctx context.Context, in *SignInByPassWordRequest, opts ...grpc.CallOption) (*SignInByPassWordResponse, error) {
	out := new(SignInByPassWordResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/SignInByPassWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) GetUserInfosById(ctx context.Context, in *GetUserInfosByIdRequest, opts ...grpc.CallOption) (*GetUserInfosByIdResponse, error) {
	out := new(GetUserInfosByIdResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/GetUserInfosById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) UpdateUserInfos(ctx context.Context, in *UpdateUserInfosRequest, opts ...grpc.CallOption) (*UpdateUserInfosResponse, error) {
	out := new(UpdateUserInfosResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/UpdateUserInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) CreatedUserProfile(ctx context.Context, in *CreatedUserProfileRequest, opts ...grpc.CallOption) (*CreatedUserProfileResponse, error) {
	out := new(CreatedUserProfileResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/CreatedUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) GetUserInfosList(ctx context.Context, in *GetUserInfosListRequest, opts ...grpc.CallOption) (*GetUserInfosListResponse, error) {
	out := new(GetUserInfosListResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/GetUserInfosList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) ApplyApplication(ctx context.Context, in *ApplyApplicationRequest, opts ...grpc.CallOption) (*ApplyApplicationResponse, error) {
	out := new(ApplyApplicationResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/ApplyApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) GetApplicationInfo(ctx context.Context, in *GetApplicationInfoRequest, opts ...grpc.CallOption) (*GetApplicationInfoResponse, error) {
	out := new(GetApplicationInfoResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/GetApplicationInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) RestAppSecretById(ctx context.Context, in *RestAppSecretByIdRequest, opts ...grpc.CallOption) (*RestAppSecretByIdResponse, error) {
	out := new(RestAppSecretByIdResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/RestAppSecretById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) ApplyFeatureAudit(ctx context.Context, in *ApplyFeatureAuditRequest, opts ...grpc.CallOption) (*ApplyFeatureAuditResponse, error) {
	out := new(ApplyFeatureAuditResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/ApplyFeatureAudit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) GetFeatureAuditList(ctx context.Context, in *GetFeatureAuditListRequest, opts ...grpc.CallOption) (*GetFeatureAuditListResponse, error) {
	out := new(GetFeatureAuditListResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/GetFeatureAuditList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) UpdateFeatureAudit(ctx context.Context, in *UpdateFeatureAuditRequest, opts ...grpc.CallOption) (*UpdateFeatureAuditResponse, error) {
	out := new(UpdateFeatureAuditResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/UpdateFeatureAudit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAPIServer is the server API for UserAPI service.
// All implementations must embed UnimplementedUserAPIServer
// for forward compatibility
type UserAPIServer interface {
	// SignInByPhoneCode 手机号验证码登录.
	SignInByPhoneCode(context.Context, *SignInByPhoneCodeRequest) (*SignInByPhoneCodeResponse, error)
	//SignInByPassWord 账号密码登录
	SignInByPassWord(context.Context, *SignInByPassWordRequest) (*SignInByPassWordResponse, error)
	//GetUserInfosById 获取用户信息
	GetUserInfosById(context.Context, *GetUserInfosByIdRequest) (*GetUserInfosByIdResponse, error)
	//UpdateUserInfos  更改用户信息
	UpdateUserInfos(context.Context, *UpdateUserInfosRequest) (*UpdateUserInfosResponse, error)
	//CreatedUserProfile 创建用户档案
	CreatedUserProfile(context.Context, *CreatedUserProfileRequest) (*CreatedUserProfileResponse, error)
	//GetUserInfosList 获取用户信息列表
	GetUserInfosList(context.Context, *GetUserInfosListRequest) (*GetUserInfosListResponse, error)
	//----------------------应用-------------------------------
	//应用申请
	ApplyApplication(context.Context, *ApplyApplicationRequest) (*ApplyApplicationResponse, error)
	//获取应用详情
	GetApplicationInfo(context.Context, *GetApplicationInfoRequest) (*GetApplicationInfoResponse, error)
	//重置app_secret
	RestAppSecretById(context.Context, *RestAppSecretByIdRequest) (*RestAppSecretByIdResponse, error)
	//应用功能申请
	ApplyFeatureAudit(context.Context, *ApplyFeatureAuditRequest) (*ApplyFeatureAuditResponse, error)
	//获取审核列表
	GetFeatureAuditList(context.Context, *GetFeatureAuditListRequest) (*GetFeatureAuditListResponse, error)
	//应用功能审核
	UpdateFeatureAudit(context.Context, *UpdateFeatureAuditRequest) (*UpdateFeatureAuditResponse, error)
	mustEmbedUnimplementedUserAPIServer()
}

// UnimplementedUserAPIServer must be embedded to have forward compatible implementations.
type UnimplementedUserAPIServer struct {
}

func (UnimplementedUserAPIServer) SignInByPhoneCode(context.Context, *SignInByPhoneCodeRequest) (*SignInByPhoneCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInByPhoneCode not implemented")
}
func (UnimplementedUserAPIServer) SignInByPassWord(context.Context, *SignInByPassWordRequest) (*SignInByPassWordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInByPassWord not implemented")
}
func (UnimplementedUserAPIServer) GetUserInfosById(context.Context, *GetUserInfosByIdRequest) (*GetUserInfosByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfosById not implemented")
}
func (UnimplementedUserAPIServer) UpdateUserInfos(context.Context, *UpdateUserInfosRequest) (*UpdateUserInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfos not implemented")
}
func (UnimplementedUserAPIServer) CreatedUserProfile(context.Context, *CreatedUserProfileRequest) (*CreatedUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedUserProfile not implemented")
}
func (UnimplementedUserAPIServer) GetUserInfosList(context.Context, *GetUserInfosListRequest) (*GetUserInfosListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfosList not implemented")
}
func (UnimplementedUserAPIServer) ApplyApplication(context.Context, *ApplyApplicationRequest) (*ApplyApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyApplication not implemented")
}
func (UnimplementedUserAPIServer) GetApplicationInfo(context.Context, *GetApplicationInfoRequest) (*GetApplicationInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationInfo not implemented")
}
func (UnimplementedUserAPIServer) RestAppSecretById(context.Context, *RestAppSecretByIdRequest) (*RestAppSecretByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestAppSecretById not implemented")
}
func (UnimplementedUserAPIServer) ApplyFeatureAudit(context.Context, *ApplyFeatureAuditRequest) (*ApplyFeatureAuditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyFeatureAudit not implemented")
}
func (UnimplementedUserAPIServer) GetFeatureAuditList(context.Context, *GetFeatureAuditListRequest) (*GetFeatureAuditListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeatureAuditList not implemented")
}
func (UnimplementedUserAPIServer) UpdateFeatureAudit(context.Context, *UpdateFeatureAuditRequest) (*UpdateFeatureAuditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeatureAudit not implemented")
}
func (UnimplementedUserAPIServer) mustEmbedUnimplementedUserAPIServer() {}

// UnsafeUserAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAPIServer will
// result in compilation errors.
type UnsafeUserAPIServer interface {
	mustEmbedUnimplementedUserAPIServer()
}

func RegisterUserAPIServer(s *grpc.Server, srv UserAPIServer) {
	s.RegisterService(&_UserAPI_serviceDesc, srv)
}

func _UserAPI_SignInByPhoneCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInByPhoneCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).SignInByPhoneCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/SignInByPhoneCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).SignInByPhoneCode(ctx, req.(*SignInByPhoneCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_SignInByPassWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInByPassWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).SignInByPassWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/SignInByPassWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).SignInByPassWord(ctx, req.(*SignInByPassWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_GetUserInfosById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfosByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetUserInfosById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/GetUserInfosById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetUserInfosById(ctx, req.(*GetUserInfosByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_UpdateUserInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).UpdateUserInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/UpdateUserInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).UpdateUserInfos(ctx, req.(*UpdateUserInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_CreatedUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).CreatedUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/CreatedUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).CreatedUserProfile(ctx, req.(*CreatedUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_GetUserInfosList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfosListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetUserInfosList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/GetUserInfosList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetUserInfosList(ctx, req.(*GetUserInfosListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_ApplyApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).ApplyApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/ApplyApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).ApplyApplication(ctx, req.(*ApplyApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_GetApplicationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetApplicationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/GetApplicationInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetApplicationInfo(ctx, req.(*GetApplicationInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_RestAppSecretById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestAppSecretByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).RestAppSecretById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/RestAppSecretById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).RestAppSecretById(ctx, req.(*RestAppSecretByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_ApplyFeatureAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyFeatureAuditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).ApplyFeatureAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/ApplyFeatureAudit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).ApplyFeatureAudit(ctx, req.(*ApplyFeatureAuditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_GetFeatureAuditList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeatureAuditListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetFeatureAuditList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/GetFeatureAuditList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetFeatureAuditList(ctx, req.(*GetFeatureAuditListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_UpdateFeatureAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFeatureAuditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).UpdateFeatureAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/UpdateFeatureAudit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).UpdateFeatureAudit(ctx, req.(*UpdateFeatureAuditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jthealth.biz.user.v1.UserAPI",
	HandlerType: (*UserAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignInByPhoneCode",
			Handler:    _UserAPI_SignInByPhoneCode_Handler,
		},
		{
			MethodName: "SignInByPassWord",
			Handler:    _UserAPI_SignInByPassWord_Handler,
		},
		{
			MethodName: "GetUserInfosById",
			Handler:    _UserAPI_GetUserInfosById_Handler,
		},
		{
			MethodName: "UpdateUserInfos",
			Handler:    _UserAPI_UpdateUserInfos_Handler,
		},
		{
			MethodName: "CreatedUserProfile",
			Handler:    _UserAPI_CreatedUserProfile_Handler,
		},
		{
			MethodName: "GetUserInfosList",
			Handler:    _UserAPI_GetUserInfosList_Handler,
		},
		{
			MethodName: "ApplyApplication",
			Handler:    _UserAPI_ApplyApplication_Handler,
		},
		{
			MethodName: "GetApplicationInfo",
			Handler:    _UserAPI_GetApplicationInfo_Handler,
		},
		{
			MethodName: "RestAppSecretById",
			Handler:    _UserAPI_RestAppSecretById_Handler,
		},
		{
			MethodName: "ApplyFeatureAudit",
			Handler:    _UserAPI_ApplyFeatureAudit_Handler,
		},
		{
			MethodName: "GetFeatureAuditList",
			Handler:    _UserAPI_GetFeatureAuditList_Handler,
		},
		{
			MethodName: "UpdateFeatureAudit",
			Handler:    _UserAPI_UpdateFeatureAudit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jthealth/biz/user/v1/user_api.proto",
}
