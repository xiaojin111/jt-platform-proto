// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: jthealth/biz/device/v1/device_api.proto

package devicev1

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

// Api Endpoints for DeviceAPI service

func NewDeviceAPIEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DeviceAPI service

type DeviceAPIService interface {
	// CheckDeviceIsUsable 判断试用设备是否可用.
	CheckDeviceIsUsable(ctx context.Context, in *CheckDeviceIsUsableRequest, opts ...client.CallOption) (*CheckDeviceIsUsableResponse, error)
	// ListTrialDevices 试用版硬件清单.
	ListTrialDevices(ctx context.Context, in *ListTrialDevicesRequest, opts ...client.CallOption) (*ListTrialDevicesResponse, error)
	// CreateTrialDevice 添加试用版设备(正式版->试用版).
	CreateTrialDevice(ctx context.Context, in *CreateTrialDeviceRequest, opts ...client.CallOption) (*CreateTrialDeviceResponse, error)
	// DeleteTrialDevice 删除试用版设备(试用版->正式版).
	DeleteTrialDevice(ctx context.Context, in *DeleteTrialDeviceRequest, opts ...client.CallOption) (*DeleteTrialDeviceResponse, error)
	// UpgradeTrialDevice 升级试用版设备(试用版->正式版).
	UpgradeTrialDevice(ctx context.Context, in *UpgradeTrialDeviceRequest, opts ...client.CallOption) (*UpgradeTrialDeviceResponse, error)
	// GetDeviceByID 通过设备ID获取设备信息.
	GetDeviceByID(ctx context.Context, in *GetDeviceByIDRequest, opts ...client.CallOption) (*GetDeviceByIDResponse, error)
	// GetDeviceBySnOrMac 通过sn或mac获取设备.
	GetDeviceBySnOrMac(ctx context.Context, in *GetDeviceBySnOrMacRequest, opts ...client.CallOption) (*GetDeviceBySnOrMacResponse, error)
	//增加设备
	AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...client.CallOption) (*AddDeviceResponse, error)
	//删除设备
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...client.CallOption) (*DeleteDeviceResponse, error)
	//更新设备
	UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...client.CallOption) (*UpdateDeviceResponse, error)
	//上传设备文件
	UploadExcel(ctx context.Context, in *UploadExcelRequest, opts ...client.CallOption) (*UploadExcelResponse, error)
	//设置设备为有赞的销售员
	SetDistributors(ctx context.Context, in *SetDistributorsReq, opts ...client.CallOption) (*SetDistributorsRsp, error)
}

type deviceAPIService struct {
	c    client.Client
	name string
}

func NewDeviceAPIService(name string, c client.Client) DeviceAPIService {
	return &deviceAPIService{
		c:    c,
		name: name,
	}
}

func (c *deviceAPIService) CheckDeviceIsUsable(ctx context.Context, in *CheckDeviceIsUsableRequest, opts ...client.CallOption) (*CheckDeviceIsUsableResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.CheckDeviceIsUsable", in)
	out := new(CheckDeviceIsUsableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) ListTrialDevices(ctx context.Context, in *ListTrialDevicesRequest, opts ...client.CallOption) (*ListTrialDevicesResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.ListTrialDevices", in)
	out := new(ListTrialDevicesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) CreateTrialDevice(ctx context.Context, in *CreateTrialDeviceRequest, opts ...client.CallOption) (*CreateTrialDeviceResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.CreateTrialDevice", in)
	out := new(CreateTrialDeviceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) DeleteTrialDevice(ctx context.Context, in *DeleteTrialDeviceRequest, opts ...client.CallOption) (*DeleteTrialDeviceResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.DeleteTrialDevice", in)
	out := new(DeleteTrialDeviceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) UpgradeTrialDevice(ctx context.Context, in *UpgradeTrialDeviceRequest, opts ...client.CallOption) (*UpgradeTrialDeviceResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.UpgradeTrialDevice", in)
	out := new(UpgradeTrialDeviceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) GetDeviceByID(ctx context.Context, in *GetDeviceByIDRequest, opts ...client.CallOption) (*GetDeviceByIDResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.GetDeviceByID", in)
	out := new(GetDeviceByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) GetDeviceBySnOrMac(ctx context.Context, in *GetDeviceBySnOrMacRequest, opts ...client.CallOption) (*GetDeviceBySnOrMacResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.GetDeviceBySnOrMac", in)
	out := new(GetDeviceBySnOrMacResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...client.CallOption) (*AddDeviceResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.AddDevice", in)
	out := new(AddDeviceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...client.CallOption) (*DeleteDeviceResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.DeleteDevice", in)
	out := new(DeleteDeviceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...client.CallOption) (*UpdateDeviceResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.UpdateDevice", in)
	out := new(UpdateDeviceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) UploadExcel(ctx context.Context, in *UploadExcelRequest, opts ...client.CallOption) (*UploadExcelResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.UploadExcel", in)
	out := new(UploadExcelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIService) SetDistributors(ctx context.Context, in *SetDistributorsReq, opts ...client.CallOption) (*SetDistributorsRsp, error) {
	req := c.c.NewRequest(c.name, "DeviceAPI.SetDistributors", in)
	out := new(SetDistributorsRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceAPI service

type DeviceAPIHandler interface {
	// CheckDeviceIsUsable 判断试用设备是否可用.
	CheckDeviceIsUsable(context.Context, *CheckDeviceIsUsableRequest, *CheckDeviceIsUsableResponse) error
	// ListTrialDevices 试用版硬件清单.
	ListTrialDevices(context.Context, *ListTrialDevicesRequest, *ListTrialDevicesResponse) error
	// CreateTrialDevice 添加试用版设备(正式版->试用版).
	CreateTrialDevice(context.Context, *CreateTrialDeviceRequest, *CreateTrialDeviceResponse) error
	// DeleteTrialDevice 删除试用版设备(试用版->正式版).
	DeleteTrialDevice(context.Context, *DeleteTrialDeviceRequest, *DeleteTrialDeviceResponse) error
	// UpgradeTrialDevice 升级试用版设备(试用版->正式版).
	UpgradeTrialDevice(context.Context, *UpgradeTrialDeviceRequest, *UpgradeTrialDeviceResponse) error
	// GetDeviceByID 通过设备ID获取设备信息.
	GetDeviceByID(context.Context, *GetDeviceByIDRequest, *GetDeviceByIDResponse) error
	// GetDeviceBySnOrMac 通过sn或mac获取设备.
	GetDeviceBySnOrMac(context.Context, *GetDeviceBySnOrMacRequest, *GetDeviceBySnOrMacResponse) error
	//增加设备
	AddDevice(context.Context, *AddDeviceRequest, *AddDeviceResponse) error
	//删除设备
	DeleteDevice(context.Context, *DeleteDeviceRequest, *DeleteDeviceResponse) error
	//更新设备
	UpdateDevice(context.Context, *UpdateDeviceRequest, *UpdateDeviceResponse) error
	//上传设备文件
	UploadExcel(context.Context, *UploadExcelRequest, *UploadExcelResponse) error
	//设置设备为有赞的销售员
	SetDistributors(context.Context, *SetDistributorsReq, *SetDistributorsRsp) error
}

func RegisterDeviceAPIHandler(s server.Server, hdlr DeviceAPIHandler, opts ...server.HandlerOption) error {
	type deviceAPI interface {
		CheckDeviceIsUsable(ctx context.Context, in *CheckDeviceIsUsableRequest, out *CheckDeviceIsUsableResponse) error
		ListTrialDevices(ctx context.Context, in *ListTrialDevicesRequest, out *ListTrialDevicesResponse) error
		CreateTrialDevice(ctx context.Context, in *CreateTrialDeviceRequest, out *CreateTrialDeviceResponse) error
		DeleteTrialDevice(ctx context.Context, in *DeleteTrialDeviceRequest, out *DeleteTrialDeviceResponse) error
		UpgradeTrialDevice(ctx context.Context, in *UpgradeTrialDeviceRequest, out *UpgradeTrialDeviceResponse) error
		GetDeviceByID(ctx context.Context, in *GetDeviceByIDRequest, out *GetDeviceByIDResponse) error
		GetDeviceBySnOrMac(ctx context.Context, in *GetDeviceBySnOrMacRequest, out *GetDeviceBySnOrMacResponse) error
		AddDevice(ctx context.Context, in *AddDeviceRequest, out *AddDeviceResponse) error
		DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, out *DeleteDeviceResponse) error
		UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, out *UpdateDeviceResponse) error
		UploadExcel(ctx context.Context, in *UploadExcelRequest, out *UploadExcelResponse) error
		SetDistributors(ctx context.Context, in *SetDistributorsReq, out *SetDistributorsRsp) error
	}
	type DeviceAPI struct {
		deviceAPI
	}
	h := &deviceAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&DeviceAPI{h}, opts...))
}

type deviceAPIHandler struct {
	DeviceAPIHandler
}

func (h *deviceAPIHandler) CheckDeviceIsUsable(ctx context.Context, in *CheckDeviceIsUsableRequest, out *CheckDeviceIsUsableResponse) error {
	return h.DeviceAPIHandler.CheckDeviceIsUsable(ctx, in, out)
}

func (h *deviceAPIHandler) ListTrialDevices(ctx context.Context, in *ListTrialDevicesRequest, out *ListTrialDevicesResponse) error {
	return h.DeviceAPIHandler.ListTrialDevices(ctx, in, out)
}

func (h *deviceAPIHandler) CreateTrialDevice(ctx context.Context, in *CreateTrialDeviceRequest, out *CreateTrialDeviceResponse) error {
	return h.DeviceAPIHandler.CreateTrialDevice(ctx, in, out)
}

func (h *deviceAPIHandler) DeleteTrialDevice(ctx context.Context, in *DeleteTrialDeviceRequest, out *DeleteTrialDeviceResponse) error {
	return h.DeviceAPIHandler.DeleteTrialDevice(ctx, in, out)
}

func (h *deviceAPIHandler) UpgradeTrialDevice(ctx context.Context, in *UpgradeTrialDeviceRequest, out *UpgradeTrialDeviceResponse) error {
	return h.DeviceAPIHandler.UpgradeTrialDevice(ctx, in, out)
}

func (h *deviceAPIHandler) GetDeviceByID(ctx context.Context, in *GetDeviceByIDRequest, out *GetDeviceByIDResponse) error {
	return h.DeviceAPIHandler.GetDeviceByID(ctx, in, out)
}

func (h *deviceAPIHandler) GetDeviceBySnOrMac(ctx context.Context, in *GetDeviceBySnOrMacRequest, out *GetDeviceBySnOrMacResponse) error {
	return h.DeviceAPIHandler.GetDeviceBySnOrMac(ctx, in, out)
}

func (h *deviceAPIHandler) AddDevice(ctx context.Context, in *AddDeviceRequest, out *AddDeviceResponse) error {
	return h.DeviceAPIHandler.AddDevice(ctx, in, out)
}

func (h *deviceAPIHandler) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, out *DeleteDeviceResponse) error {
	return h.DeviceAPIHandler.DeleteDevice(ctx, in, out)
}

func (h *deviceAPIHandler) UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, out *UpdateDeviceResponse) error {
	return h.DeviceAPIHandler.UpdateDevice(ctx, in, out)
}

func (h *deviceAPIHandler) UploadExcel(ctx context.Context, in *UploadExcelRequest, out *UploadExcelResponse) error {
	return h.DeviceAPIHandler.UploadExcel(ctx, in, out)
}

func (h *deviceAPIHandler) SetDistributors(ctx context.Context, in *SetDistributorsReq, out *SetDistributorsRsp) error {
	return h.DeviceAPIHandler.SetDistributors(ctx, in, out)
}
