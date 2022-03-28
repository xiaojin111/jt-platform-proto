// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package devicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceAPIClient is the client API for DeviceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceAPIClient interface {
	// CheckDeviceIsUsable 判断试用设备是否可用.
	CheckDeviceIsUsable(ctx context.Context, in *CheckDeviceIsUsableRequest, opts ...grpc.CallOption) (*CheckDeviceIsUsableResponse, error)
	// ListTrialDevices 试用版硬件清单.
	ListTrialDevices(ctx context.Context, in *ListTrialDevicesRequest, opts ...grpc.CallOption) (*ListTrialDevicesResponse, error)
	// CreateTrialDevice 添加试用版设备(正式版->试用版).
	CreateTrialDevice(ctx context.Context, in *CreateTrialDeviceRequest, opts ...grpc.CallOption) (*CreateTrialDeviceResponse, error)
	// DeleteTrialDevice 删除试用版设备(试用版->正式版).
	DeleteTrialDevice(ctx context.Context, in *DeleteTrialDeviceRequest, opts ...grpc.CallOption) (*DeleteTrialDeviceResponse, error)
	// UpgradeTrialDevice 升级试用版设备(试用版->正式版).
	UpgradeTrialDevice(ctx context.Context, in *UpgradeTrialDeviceRequest, opts ...grpc.CallOption) (*UpgradeTrialDeviceResponse, error)
	// GetDeviceByID 通过设备ID获取设备信息.
	GetDeviceByID(ctx context.Context, in *GetDeviceByIDRequest, opts ...grpc.CallOption) (*GetDeviceByIDResponse, error)
	// GetDeviceBySnOrMac 通过sn或mac获取设备.
	GetDeviceBySnOrMac(ctx context.Context, in *GetDeviceBySnOrMacRequest, opts ...grpc.CallOption) (*GetDeviceBySnOrMacResponse, error)
	//增加设备
	AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error)
	//删除设备
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error)
	//更新设备
	UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error)
	//上传设备文件
	UploadExcel(ctx context.Context, in *UploadExcelRequest, opts ...grpc.CallOption) (*UploadExcelResponse, error)
	//设置设备为有赞的销售员
	SetDistributors(ctx context.Context, in *SetDistributorsReq, opts ...grpc.CallOption) (*SetDistributorsRsp, error)
}

type deviceAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceAPIClient(cc grpc.ClientConnInterface) DeviceAPIClient {
	return &deviceAPIClient{cc}
}

func (c *deviceAPIClient) CheckDeviceIsUsable(ctx context.Context, in *CheckDeviceIsUsableRequest, opts ...grpc.CallOption) (*CheckDeviceIsUsableResponse, error) {
	out := new(CheckDeviceIsUsableResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/CheckDeviceIsUsable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) ListTrialDevices(ctx context.Context, in *ListTrialDevicesRequest, opts ...grpc.CallOption) (*ListTrialDevicesResponse, error) {
	out := new(ListTrialDevicesResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/ListTrialDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) CreateTrialDevice(ctx context.Context, in *CreateTrialDeviceRequest, opts ...grpc.CallOption) (*CreateTrialDeviceResponse, error) {
	out := new(CreateTrialDeviceResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/CreateTrialDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) DeleteTrialDevice(ctx context.Context, in *DeleteTrialDeviceRequest, opts ...grpc.CallOption) (*DeleteTrialDeviceResponse, error) {
	out := new(DeleteTrialDeviceResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/DeleteTrialDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) UpgradeTrialDevice(ctx context.Context, in *UpgradeTrialDeviceRequest, opts ...grpc.CallOption) (*UpgradeTrialDeviceResponse, error) {
	out := new(UpgradeTrialDeviceResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/UpgradeTrialDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) GetDeviceByID(ctx context.Context, in *GetDeviceByIDRequest, opts ...grpc.CallOption) (*GetDeviceByIDResponse, error) {
	out := new(GetDeviceByIDResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/GetDeviceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) GetDeviceBySnOrMac(ctx context.Context, in *GetDeviceBySnOrMacRequest, opts ...grpc.CallOption) (*GetDeviceBySnOrMacResponse, error) {
	out := new(GetDeviceBySnOrMacResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/GetDeviceBySnOrMac", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error) {
	out := new(AddDeviceResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/AddDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error) {
	out := new(DeleteDeviceResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error) {
	out := new(UpdateDeviceResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/UpdateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) UploadExcel(ctx context.Context, in *UploadExcelRequest, opts ...grpc.CallOption) (*UploadExcelResponse, error) {
	out := new(UploadExcelResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/UploadExcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAPIClient) SetDistributors(ctx context.Context, in *SetDistributorsReq, opts ...grpc.CallOption) (*SetDistributorsRsp, error) {
	out := new(SetDistributorsRsp)
	err := c.cc.Invoke(ctx, "/jthealth.biz.device.v1.DeviceAPI/SetDistributors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceAPIServer is the server API for DeviceAPI service.
// All implementations must embed UnimplementedDeviceAPIServer
// for forward compatibility
type DeviceAPIServer interface {
	// CheckDeviceIsUsable 判断试用设备是否可用.
	CheckDeviceIsUsable(context.Context, *CheckDeviceIsUsableRequest) (*CheckDeviceIsUsableResponse, error)
	// ListTrialDevices 试用版硬件清单.
	ListTrialDevices(context.Context, *ListTrialDevicesRequest) (*ListTrialDevicesResponse, error)
	// CreateTrialDevice 添加试用版设备(正式版->试用版).
	CreateTrialDevice(context.Context, *CreateTrialDeviceRequest) (*CreateTrialDeviceResponse, error)
	// DeleteTrialDevice 删除试用版设备(试用版->正式版).
	DeleteTrialDevice(context.Context, *DeleteTrialDeviceRequest) (*DeleteTrialDeviceResponse, error)
	// UpgradeTrialDevice 升级试用版设备(试用版->正式版).
	UpgradeTrialDevice(context.Context, *UpgradeTrialDeviceRequest) (*UpgradeTrialDeviceResponse, error)
	// GetDeviceByID 通过设备ID获取设备信息.
	GetDeviceByID(context.Context, *GetDeviceByIDRequest) (*GetDeviceByIDResponse, error)
	// GetDeviceBySnOrMac 通过sn或mac获取设备.
	GetDeviceBySnOrMac(context.Context, *GetDeviceBySnOrMacRequest) (*GetDeviceBySnOrMacResponse, error)
	//增加设备
	AddDevice(context.Context, *AddDeviceRequest) (*AddDeviceResponse, error)
	//删除设备
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error)
	//更新设备
	UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error)
	//上传设备文件
	UploadExcel(context.Context, *UploadExcelRequest) (*UploadExcelResponse, error)
	//设置设备为有赞的销售员
	SetDistributors(context.Context, *SetDistributorsReq) (*SetDistributorsRsp, error)
	mustEmbedUnimplementedDeviceAPIServer()
}

// UnimplementedDeviceAPIServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceAPIServer struct {
}

func (UnimplementedDeviceAPIServer) CheckDeviceIsUsable(context.Context, *CheckDeviceIsUsableRequest) (*CheckDeviceIsUsableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDeviceIsUsable not implemented")
}
func (UnimplementedDeviceAPIServer) ListTrialDevices(context.Context, *ListTrialDevicesRequest) (*ListTrialDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrialDevices not implemented")
}
func (UnimplementedDeviceAPIServer) CreateTrialDevice(context.Context, *CreateTrialDeviceRequest) (*CreateTrialDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrialDevice not implemented")
}
func (UnimplementedDeviceAPIServer) DeleteTrialDevice(context.Context, *DeleteTrialDeviceRequest) (*DeleteTrialDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrialDevice not implemented")
}
func (UnimplementedDeviceAPIServer) UpgradeTrialDevice(context.Context, *UpgradeTrialDeviceRequest) (*UpgradeTrialDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeTrialDevice not implemented")
}
func (UnimplementedDeviceAPIServer) GetDeviceByID(context.Context, *GetDeviceByIDRequest) (*GetDeviceByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceByID not implemented")
}
func (UnimplementedDeviceAPIServer) GetDeviceBySnOrMac(context.Context, *GetDeviceBySnOrMacRequest) (*GetDeviceBySnOrMacResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceBySnOrMac not implemented")
}
func (UnimplementedDeviceAPIServer) AddDevice(context.Context, *AddDeviceRequest) (*AddDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevice not implemented")
}
func (UnimplementedDeviceAPIServer) DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (UnimplementedDeviceAPIServer) UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevice not implemented")
}
func (UnimplementedDeviceAPIServer) UploadExcel(context.Context, *UploadExcelRequest) (*UploadExcelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadExcel not implemented")
}
func (UnimplementedDeviceAPIServer) SetDistributors(context.Context, *SetDistributorsReq) (*SetDistributorsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDistributors not implemented")
}
func (UnimplementedDeviceAPIServer) mustEmbedUnimplementedDeviceAPIServer() {}

// UnsafeDeviceAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceAPIServer will
// result in compilation errors.
type UnsafeDeviceAPIServer interface {
	mustEmbedUnimplementedDeviceAPIServer()
}

func RegisterDeviceAPIServer(s *grpc.Server, srv DeviceAPIServer) {
	s.RegisterService(&_DeviceAPI_serviceDesc, srv)
}

func _DeviceAPI_CheckDeviceIsUsable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDeviceIsUsableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).CheckDeviceIsUsable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/CheckDeviceIsUsable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).CheckDeviceIsUsable(ctx, req.(*CheckDeviceIsUsableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_ListTrialDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTrialDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).ListTrialDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/ListTrialDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).ListTrialDevices(ctx, req.(*ListTrialDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_CreateTrialDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrialDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).CreateTrialDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/CreateTrialDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).CreateTrialDevice(ctx, req.(*CreateTrialDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_DeleteTrialDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTrialDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).DeleteTrialDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/DeleteTrialDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).DeleteTrialDevice(ctx, req.(*DeleteTrialDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_UpgradeTrialDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeTrialDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).UpgradeTrialDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/UpgradeTrialDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).UpgradeTrialDevice(ctx, req.(*UpgradeTrialDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_GetDeviceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).GetDeviceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/GetDeviceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).GetDeviceByID(ctx, req.(*GetDeviceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_GetDeviceBySnOrMac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceBySnOrMacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).GetDeviceBySnOrMac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/GetDeviceBySnOrMac",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).GetDeviceBySnOrMac(ctx, req.(*GetDeviceBySnOrMacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/AddDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).AddDevice(ctx, req.(*AddDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).DeleteDevice(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_UpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).UpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/UpdateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).UpdateDevice(ctx, req.(*UpdateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_UploadExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadExcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).UploadExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/UploadExcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).UploadExcel(ctx, req.(*UploadExcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAPI_SetDistributors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDistributorsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAPIServer).SetDistributors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.device.v1.DeviceAPI/SetDistributors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAPIServer).SetDistributors(ctx, req.(*SetDistributorsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jthealth.biz.device.v1.DeviceAPI",
	HandlerType: (*DeviceAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckDeviceIsUsable",
			Handler:    _DeviceAPI_CheckDeviceIsUsable_Handler,
		},
		{
			MethodName: "ListTrialDevices",
			Handler:    _DeviceAPI_ListTrialDevices_Handler,
		},
		{
			MethodName: "CreateTrialDevice",
			Handler:    _DeviceAPI_CreateTrialDevice_Handler,
		},
		{
			MethodName: "DeleteTrialDevice",
			Handler:    _DeviceAPI_DeleteTrialDevice_Handler,
		},
		{
			MethodName: "UpgradeTrialDevice",
			Handler:    _DeviceAPI_UpgradeTrialDevice_Handler,
		},
		{
			MethodName: "GetDeviceByID",
			Handler:    _DeviceAPI_GetDeviceByID_Handler,
		},
		{
			MethodName: "GetDeviceBySnOrMac",
			Handler:    _DeviceAPI_GetDeviceBySnOrMac_Handler,
		},
		{
			MethodName: "AddDevice",
			Handler:    _DeviceAPI_AddDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _DeviceAPI_DeleteDevice_Handler,
		},
		{
			MethodName: "UpdateDevice",
			Handler:    _DeviceAPI_UpdateDevice_Handler,
		},
		{
			MethodName: "UploadExcel",
			Handler:    _DeviceAPI_UploadExcel_Handler,
		},
		{
			MethodName: "SetDistributors",
			Handler:    _DeviceAPI_SetDistributors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jthealth/biz/device/v1/device_api.proto",
}
