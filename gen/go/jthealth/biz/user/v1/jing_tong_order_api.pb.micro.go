// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: jthealth/biz/user/v1/jing_tong_order_api.proto

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

// Api Endpoints for JingTongOrderAPI service

func NewJingTongOrderAPIEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for JingTongOrderAPI service

type JingTongOrderAPIService interface {
	//获取用户剩余测量次数
	JingTongMeasurementCount(ctx context.Context, in *JingTongMeasurementCountRequest, opts ...client.CallOption) (*JingTongMeasurementCountResponse, error)
	//获取用户openId
	JingTongGetWxMiniProOpenIdByCode(ctx context.Context, in *JingTongGetWxMiniProOpenIdByCodeRequest, opts ...client.CallOption) (*JingTongGetWxMiniProOpenIdByCodeResponse, error)
	//创建订单
	JingTongCreateOrder(ctx context.Context, in *JingTongCreateOrderRequest, opts ...client.CallOption) (*JingTongCreateOrderResponse, error)
	//微信支付通知
	JingTongNotify(ctx context.Context, in *JingTongNotifyRequest, opts ...client.CallOption) (*JingTongNotifyResponse, error)
	//获取订单列表
	JingTongGetOrderList(ctx context.Context, in *JingTongGetOrderListRequest, opts ...client.CallOption) (*JingTongGetOrderListResponse, error)
	//首页订单统计
	JingTongOrderStatistical(ctx context.Context, in *JingTongOrderStatisticalRequest, opts ...client.CallOption) (*JingTongOrderStatisticalResponse, error)
	//后台获取订单列表
	JingTongBGGetOrderList(ctx context.Context, in *JingTongBGGetOrderListRequest, opts ...client.CallOption) (*JingTongBGGetOrderListResponse, error)
}

type jingTongOrderAPIService struct {
	c    client.Client
	name string
}

func NewJingTongOrderAPIService(name string, c client.Client) JingTongOrderAPIService {
	return &jingTongOrderAPIService{
		c:    c,
		name: name,
	}
}

func (c *jingTongOrderAPIService) JingTongMeasurementCount(ctx context.Context, in *JingTongMeasurementCountRequest, opts ...client.CallOption) (*JingTongMeasurementCountResponse, error) {
	req := c.c.NewRequest(c.name, "JingTongOrderAPI.JingTongMeasurementCount", in)
	out := new(JingTongMeasurementCountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jingTongOrderAPIService) JingTongGetWxMiniProOpenIdByCode(ctx context.Context, in *JingTongGetWxMiniProOpenIdByCodeRequest, opts ...client.CallOption) (*JingTongGetWxMiniProOpenIdByCodeResponse, error) {
	req := c.c.NewRequest(c.name, "JingTongOrderAPI.JingTongGetWxMiniProOpenIdByCode", in)
	out := new(JingTongGetWxMiniProOpenIdByCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jingTongOrderAPIService) JingTongCreateOrder(ctx context.Context, in *JingTongCreateOrderRequest, opts ...client.CallOption) (*JingTongCreateOrderResponse, error) {
	req := c.c.NewRequest(c.name, "JingTongOrderAPI.JingTongCreateOrder", in)
	out := new(JingTongCreateOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jingTongOrderAPIService) JingTongNotify(ctx context.Context, in *JingTongNotifyRequest, opts ...client.CallOption) (*JingTongNotifyResponse, error) {
	req := c.c.NewRequest(c.name, "JingTongOrderAPI.JingTongNotify", in)
	out := new(JingTongNotifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jingTongOrderAPIService) JingTongGetOrderList(ctx context.Context, in *JingTongGetOrderListRequest, opts ...client.CallOption) (*JingTongGetOrderListResponse, error) {
	req := c.c.NewRequest(c.name, "JingTongOrderAPI.JingTongGetOrderList", in)
	out := new(JingTongGetOrderListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jingTongOrderAPIService) JingTongOrderStatistical(ctx context.Context, in *JingTongOrderStatisticalRequest, opts ...client.CallOption) (*JingTongOrderStatisticalResponse, error) {
	req := c.c.NewRequest(c.name, "JingTongOrderAPI.JingTongOrderStatistical", in)
	out := new(JingTongOrderStatisticalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jingTongOrderAPIService) JingTongBGGetOrderList(ctx context.Context, in *JingTongBGGetOrderListRequest, opts ...client.CallOption) (*JingTongBGGetOrderListResponse, error) {
	req := c.c.NewRequest(c.name, "JingTongOrderAPI.JingTongBGGetOrderList", in)
	out := new(JingTongBGGetOrderListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JingTongOrderAPI service

type JingTongOrderAPIHandler interface {
	//获取用户剩余测量次数
	JingTongMeasurementCount(context.Context, *JingTongMeasurementCountRequest, *JingTongMeasurementCountResponse) error
	//获取用户openId
	JingTongGetWxMiniProOpenIdByCode(context.Context, *JingTongGetWxMiniProOpenIdByCodeRequest, *JingTongGetWxMiniProOpenIdByCodeResponse) error
	//创建订单
	JingTongCreateOrder(context.Context, *JingTongCreateOrderRequest, *JingTongCreateOrderResponse) error
	//微信支付通知
	JingTongNotify(context.Context, *JingTongNotifyRequest, *JingTongNotifyResponse) error
	//获取订单列表
	JingTongGetOrderList(context.Context, *JingTongGetOrderListRequest, *JingTongGetOrderListResponse) error
	//首页订单统计
	JingTongOrderStatistical(context.Context, *JingTongOrderStatisticalRequest, *JingTongOrderStatisticalResponse) error
	//后台获取订单列表
	JingTongBGGetOrderList(context.Context, *JingTongBGGetOrderListRequest, *JingTongBGGetOrderListResponse) error
}

func RegisterJingTongOrderAPIHandler(s server.Server, hdlr JingTongOrderAPIHandler, opts ...server.HandlerOption) error {
	type jingTongOrderAPI interface {
		JingTongMeasurementCount(ctx context.Context, in *JingTongMeasurementCountRequest, out *JingTongMeasurementCountResponse) error
		JingTongGetWxMiniProOpenIdByCode(ctx context.Context, in *JingTongGetWxMiniProOpenIdByCodeRequest, out *JingTongGetWxMiniProOpenIdByCodeResponse) error
		JingTongCreateOrder(ctx context.Context, in *JingTongCreateOrderRequest, out *JingTongCreateOrderResponse) error
		JingTongNotify(ctx context.Context, in *JingTongNotifyRequest, out *JingTongNotifyResponse) error
		JingTongGetOrderList(ctx context.Context, in *JingTongGetOrderListRequest, out *JingTongGetOrderListResponse) error
		JingTongOrderStatistical(ctx context.Context, in *JingTongOrderStatisticalRequest, out *JingTongOrderStatisticalResponse) error
		JingTongBGGetOrderList(ctx context.Context, in *JingTongBGGetOrderListRequest, out *JingTongBGGetOrderListResponse) error
	}
	type JingTongOrderAPI struct {
		jingTongOrderAPI
	}
	h := &jingTongOrderAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&JingTongOrderAPI{h}, opts...))
}

type jingTongOrderAPIHandler struct {
	JingTongOrderAPIHandler
}

func (h *jingTongOrderAPIHandler) JingTongMeasurementCount(ctx context.Context, in *JingTongMeasurementCountRequest, out *JingTongMeasurementCountResponse) error {
	return h.JingTongOrderAPIHandler.JingTongMeasurementCount(ctx, in, out)
}

func (h *jingTongOrderAPIHandler) JingTongGetWxMiniProOpenIdByCode(ctx context.Context, in *JingTongGetWxMiniProOpenIdByCodeRequest, out *JingTongGetWxMiniProOpenIdByCodeResponse) error {
	return h.JingTongOrderAPIHandler.JingTongGetWxMiniProOpenIdByCode(ctx, in, out)
}

func (h *jingTongOrderAPIHandler) JingTongCreateOrder(ctx context.Context, in *JingTongCreateOrderRequest, out *JingTongCreateOrderResponse) error {
	return h.JingTongOrderAPIHandler.JingTongCreateOrder(ctx, in, out)
}

func (h *jingTongOrderAPIHandler) JingTongNotify(ctx context.Context, in *JingTongNotifyRequest, out *JingTongNotifyResponse) error {
	return h.JingTongOrderAPIHandler.JingTongNotify(ctx, in, out)
}

func (h *jingTongOrderAPIHandler) JingTongGetOrderList(ctx context.Context, in *JingTongGetOrderListRequest, out *JingTongGetOrderListResponse) error {
	return h.JingTongOrderAPIHandler.JingTongGetOrderList(ctx, in, out)
}

func (h *jingTongOrderAPIHandler) JingTongOrderStatistical(ctx context.Context, in *JingTongOrderStatisticalRequest, out *JingTongOrderStatisticalResponse) error {
	return h.JingTongOrderAPIHandler.JingTongOrderStatistical(ctx, in, out)
}

func (h *jingTongOrderAPIHandler) JingTongBGGetOrderList(ctx context.Context, in *JingTongBGGetOrderListRequest, out *JingTongBGGetOrderListResponse) error {
	return h.JingTongOrderAPIHandler.JingTongBGGetOrderList(ctx, in, out)
}
