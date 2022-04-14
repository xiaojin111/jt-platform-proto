// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: jthealth/biz/user/v1/applet_user_api.proto

package userv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Api Endpoints for AppletUserAPI service

func NewAppletUserAPIEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AppletUserAPI service

type AppletUserAPIService interface {
	// AppletUserSignInByPhoneCode 手机号验证码登录.
	AppletUserSignInByPhoneCode(ctx context.Context, in *AppletUserSignInByPhoneCodeRequest, opts ...client.CallOption) (*AppletUserSignInByPhoneCodeResponse, error)
	// SignInByWechat 通过微信小程序登录.
	SignInByWechatMiniProgram(ctx context.Context, in *SignInByWechatMiniProgramRequest, opts ...client.CallOption) (*SignInByWechatMiniProgramResponse, error)
	// GetWechatPhone 获取微信手机号.
	GetWechatPhone(ctx context.Context, in *GetWechatPhoneRequest, opts ...client.CallOption) (*GetWechatPhoneResponse, error)
	// BindWechatPhone 微信登录绑定手机号.
	BindWechatPhone(ctx context.Context, in *BindWechatPhoneRequest, opts ...client.CallOption) (*BindWechatPhoneResponse, error)
}

type appletUserAPIService struct {
	c    client.Client
	name string
}

func NewAppletUserAPIService(name string, c client.Client) AppletUserAPIService {
	return &appletUserAPIService{
		c:    c,
		name: name,
	}
}

func (c *appletUserAPIService) AppletUserSignInByPhoneCode(ctx context.Context, in *AppletUserSignInByPhoneCodeRequest, opts ...client.CallOption) (*AppletUserSignInByPhoneCodeResponse, error) {
	req := c.c.NewRequest(c.name, "AppletUserAPI.AppletUserSignInByPhoneCode", in)
	out := new(AppletUserSignInByPhoneCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appletUserAPIService) SignInByWechatMiniProgram(ctx context.Context, in *SignInByWechatMiniProgramRequest, opts ...client.CallOption) (*SignInByWechatMiniProgramResponse, error) {
	req := c.c.NewRequest(c.name, "AppletUserAPI.SignInByWechatMiniProgram", in)
	out := new(SignInByWechatMiniProgramResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appletUserAPIService) GetWechatPhone(ctx context.Context, in *GetWechatPhoneRequest, opts ...client.CallOption) (*GetWechatPhoneResponse, error) {
	req := c.c.NewRequest(c.name, "AppletUserAPI.GetWechatPhone", in)
	out := new(GetWechatPhoneResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appletUserAPIService) BindWechatPhone(ctx context.Context, in *BindWechatPhoneRequest, opts ...client.CallOption) (*BindWechatPhoneResponse, error) {
	req := c.c.NewRequest(c.name, "AppletUserAPI.BindWechatPhone", in)
	out := new(BindWechatPhoneResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppletUserAPI service

type AppletUserAPIHandler interface {
	// AppletUserSignInByPhoneCode 手机号验证码登录.
	AppletUserSignInByPhoneCode(context.Context, *AppletUserSignInByPhoneCodeRequest, *AppletUserSignInByPhoneCodeResponse) error
	// SignInByWechat 通过微信小程序登录.
	SignInByWechatMiniProgram(context.Context, *SignInByWechatMiniProgramRequest, *SignInByWechatMiniProgramResponse) error
	// GetWechatPhone 获取微信手机号.
	GetWechatPhone(context.Context, *GetWechatPhoneRequest, *GetWechatPhoneResponse) error
	// BindWechatPhone 微信登录绑定手机号.
	BindWechatPhone(context.Context, *BindWechatPhoneRequest, *BindWechatPhoneResponse) error
}

func RegisterAppletUserAPIHandler(s server.Server, hdlr AppletUserAPIHandler, opts ...server.HandlerOption) error {
	type appletUserAPI interface {
		AppletUserSignInByPhoneCode(ctx context.Context, in *AppletUserSignInByPhoneCodeRequest, out *AppletUserSignInByPhoneCodeResponse) error
		SignInByWechatMiniProgram(ctx context.Context, in *SignInByWechatMiniProgramRequest, out *SignInByWechatMiniProgramResponse) error
		GetWechatPhone(ctx context.Context, in *GetWechatPhoneRequest, out *GetWechatPhoneResponse) error
		BindWechatPhone(ctx context.Context, in *BindWechatPhoneRequest, out *BindWechatPhoneResponse) error
	}
	type AppletUserAPI struct {
		appletUserAPI
	}
	h := &appletUserAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&AppletUserAPI{h}, opts...))
}

type appletUserAPIHandler struct {
	AppletUserAPIHandler
}

func (h *appletUserAPIHandler) AppletUserSignInByPhoneCode(ctx context.Context, in *AppletUserSignInByPhoneCodeRequest, out *AppletUserSignInByPhoneCodeResponse) error {
	return h.AppletUserAPIHandler.AppletUserSignInByPhoneCode(ctx, in, out)
}

func (h *appletUserAPIHandler) SignInByWechatMiniProgram(ctx context.Context, in *SignInByWechatMiniProgramRequest, out *SignInByWechatMiniProgramResponse) error {
	return h.AppletUserAPIHandler.SignInByWechatMiniProgram(ctx, in, out)
}

func (h *appletUserAPIHandler) GetWechatPhone(ctx context.Context, in *GetWechatPhoneRequest, out *GetWechatPhoneResponse) error {
	return h.AppletUserAPIHandler.GetWechatPhone(ctx, in, out)
}

func (h *appletUserAPIHandler) BindWechatPhone(ctx context.Context, in *BindWechatPhoneRequest, out *BindWechatPhoneResponse) error {
	return h.AppletUserAPIHandler.BindWechatPhone(ctx, in, out)
}