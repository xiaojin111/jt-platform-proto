// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: jthealth/biz/user/v1/user_api.proto

package userv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserAPI service

func NewUserAPIEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserAPI service

type UserAPIService interface {
	// SignInByPhoneCode 手机号验证码登录.
	SignInByPhoneCode(ctx context.Context, in *SignInByPhoneCodeRequest, opts ...client.CallOption) (*SignInByPhoneCodeResponse, error)
	//SignInByPassWord 账号密码登录
	SignInByPassWord(ctx context.Context, in *SignInByPassWordRequest, opts ...client.CallOption) (*SignInByPassWordResponse, error)
	//GetUserInfosById 获取用户信息
	GetUserInfosById(ctx context.Context, in *GetUserInfosByIdRequest, opts ...client.CallOption) (*GetUserInfosByIdResponse, error)
	//UpdateUserInfos  更改用户信息
	UpdateUserInfos(ctx context.Context, in *UpdateUserInfosRequest, opts ...client.CallOption) (*UpdateUserInfosResponse, error)
	//----------------------应用详情-------------------------------
	//应用申请
	ApplyApplication(ctx context.Context, in *ApplyApplicationRequest, opts ...client.CallOption) (*ApplyApplicationResponse, error)
	//获取应用详情
	GetApplicationInfo(ctx context.Context, in *GetApplicationInfoRequest, opts ...client.CallOption) (*GetApplicationInfoResponse, error)
}

type userAPIService struct {
	c    client.Client
	name string
}

func NewUserAPIService(name string, c client.Client) UserAPIService {
	return &userAPIService{
		c:    c,
		name: name,
	}
}

func (c *userAPIService) SignInByPhoneCode(ctx context.Context, in *SignInByPhoneCodeRequest, opts ...client.CallOption) (*SignInByPhoneCodeResponse, error) {
	req := c.c.NewRequest(c.name, "UserAPI.SignInByPhoneCode", in)
	out := new(SignInByPhoneCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIService) SignInByPassWord(ctx context.Context, in *SignInByPassWordRequest, opts ...client.CallOption) (*SignInByPassWordResponse, error) {
	req := c.c.NewRequest(c.name, "UserAPI.SignInByPassWord", in)
	out := new(SignInByPassWordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIService) GetUserInfosById(ctx context.Context, in *GetUserInfosByIdRequest, opts ...client.CallOption) (*GetUserInfosByIdResponse, error) {
	req := c.c.NewRequest(c.name, "UserAPI.GetUserInfosById", in)
	out := new(GetUserInfosByIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIService) UpdateUserInfos(ctx context.Context, in *UpdateUserInfosRequest, opts ...client.CallOption) (*UpdateUserInfosResponse, error) {
	req := c.c.NewRequest(c.name, "UserAPI.UpdateUserInfos", in)
	out := new(UpdateUserInfosResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIService) ApplyApplication(ctx context.Context, in *ApplyApplicationRequest, opts ...client.CallOption) (*ApplyApplicationResponse, error) {
	req := c.c.NewRequest(c.name, "UserAPI.ApplyApplication", in)
	out := new(ApplyApplicationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIService) GetApplicationInfo(ctx context.Context, in *GetApplicationInfoRequest, opts ...client.CallOption) (*GetApplicationInfoResponse, error) {
	req := c.c.NewRequest(c.name, "UserAPI.GetApplicationInfo", in)
	out := new(GetApplicationInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserAPI service

type UserAPIHandler interface {
	// SignInByPhoneCode 手机号验证码登录.
	SignInByPhoneCode(context.Context, *SignInByPhoneCodeRequest, *SignInByPhoneCodeResponse) error
	//SignInByPassWord 账号密码登录
	SignInByPassWord(context.Context, *SignInByPassWordRequest, *SignInByPassWordResponse) error
	//GetUserInfosById 获取用户信息
	GetUserInfosById(context.Context, *GetUserInfosByIdRequest, *GetUserInfosByIdResponse) error
	//UpdateUserInfos  更改用户信息
	UpdateUserInfos(context.Context, *UpdateUserInfosRequest, *UpdateUserInfosResponse) error
	//----------------------应用详情-------------------------------
	//应用申请
	ApplyApplication(context.Context, *ApplyApplicationRequest, *ApplyApplicationResponse) error
	//获取应用详情
	GetApplicationInfo(context.Context, *GetApplicationInfoRequest, *GetApplicationInfoResponse) error
}

func RegisterUserAPIHandler(s server.Server, hdlr UserAPIHandler, opts ...server.HandlerOption) error {
	type userAPI interface {
		SignInByPhoneCode(ctx context.Context, in *SignInByPhoneCodeRequest, out *SignInByPhoneCodeResponse) error
		SignInByPassWord(ctx context.Context, in *SignInByPassWordRequest, out *SignInByPassWordResponse) error
		GetUserInfosById(ctx context.Context, in *GetUserInfosByIdRequest, out *GetUserInfosByIdResponse) error
		UpdateUserInfos(ctx context.Context, in *UpdateUserInfosRequest, out *UpdateUserInfosResponse) error
		ApplyApplication(ctx context.Context, in *ApplyApplicationRequest, out *ApplyApplicationResponse) error
		GetApplicationInfo(ctx context.Context, in *GetApplicationInfoRequest, out *GetApplicationInfoResponse) error
	}
	type UserAPI struct {
		userAPI
	}
	h := &userAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&UserAPI{h}, opts...))
}

type userAPIHandler struct {
	UserAPIHandler
}

func (h *userAPIHandler) SignInByPhoneCode(ctx context.Context, in *SignInByPhoneCodeRequest, out *SignInByPhoneCodeResponse) error {
	return h.UserAPIHandler.SignInByPhoneCode(ctx, in, out)
}

func (h *userAPIHandler) SignInByPassWord(ctx context.Context, in *SignInByPassWordRequest, out *SignInByPassWordResponse) error {
	return h.UserAPIHandler.SignInByPassWord(ctx, in, out)
}

func (h *userAPIHandler) GetUserInfosById(ctx context.Context, in *GetUserInfosByIdRequest, out *GetUserInfosByIdResponse) error {
	return h.UserAPIHandler.GetUserInfosById(ctx, in, out)
}

func (h *userAPIHandler) UpdateUserInfos(ctx context.Context, in *UpdateUserInfosRequest, out *UpdateUserInfosResponse) error {
	return h.UserAPIHandler.UpdateUserInfos(ctx, in, out)
}

func (h *userAPIHandler) ApplyApplication(ctx context.Context, in *ApplyApplicationRequest, out *ApplyApplicationResponse) error {
	return h.UserAPIHandler.ApplyApplication(ctx, in, out)
}

func (h *userAPIHandler) GetApplicationInfo(ctx context.Context, in *GetApplicationInfoRequest, out *GetApplicationInfoResponse) error {
	return h.UserAPIHandler.GetApplicationInfo(ctx, in, out)
}
