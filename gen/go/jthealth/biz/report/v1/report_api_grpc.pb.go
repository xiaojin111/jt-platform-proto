// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package reportv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportAPIClient is the client API for ReportAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportAPIClient interface {
	// SubmitPulseTest 提交采样数据.
	SubmitPulseTest(ctx context.Context, in *SubmitPulseTestRequest, opts ...grpc.CallOption) (*SubmitPulseTestResponse, error)
	// GetReport 获取阶梯报告.
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error)
	// GetMealSuggestion 通过体质获得膳食建议.
	GetMealSuggestion(ctx context.Context, in *GetMealSuggestionRequest, opts ...grpc.CallOption) (*GetMealSuggestionResponse, error)
	//创建风险推荐商品
	CreateRiskCommodity(ctx context.Context, in *CreateRiskCommodityRequest, opts ...grpc.CallOption) (*CreateRiskCommodityResponse, error)
	//获取风险列表
	GetRiskList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRiskListResponse, error)
	//获取推荐商品列表
	GetRiskCommodityList(ctx context.Context, in *GetRiskCommodityListRequest, opts ...grpc.CallOption) (*GetRiskCommodityListResponse, error)
	//编辑报告显示内容
	EditReportShow(ctx context.Context, in *EditReportShowRequest, opts ...grpc.CallOption) (*EditReportShowResponse, error)
	//获取报告显示内容
	GetReportShow(ctx context.Context, in *GetReportShowRequest, opts ...grpc.CallOption) (*GetReportShowResponse, error)
	//编辑报告对比显示功能
	EditReportComparedShow(ctx context.Context, in *EditReportComparedShowRequest, opts ...grpc.CallOption) (*EditReportComparedShowResponse, error)
	//获取报告对比显示功能
	GetReportComparedShow(ctx context.Context, in *GetReportComparedShowRequest, opts ...grpc.CallOption) (*GetReportComparedShowResponse, error)
	// 报告对比接口
	GetComparisonReportNew(ctx context.Context, in *GetComparisonReportNewRequest, opts ...grpc.CallOption) (*GetComparisonReportNewResponse, error)
}

type reportAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewReportAPIClient(cc grpc.ClientConnInterface) ReportAPIClient {
	return &reportAPIClient{cc}
}

func (c *reportAPIClient) SubmitPulseTest(ctx context.Context, in *SubmitPulseTestRequest, opts ...grpc.CallOption) (*SubmitPulseTestResponse, error) {
	out := new(SubmitPulseTestResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/SubmitPulseTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error) {
	out := new(GetReportResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) GetMealSuggestion(ctx context.Context, in *GetMealSuggestionRequest, opts ...grpc.CallOption) (*GetMealSuggestionResponse, error) {
	out := new(GetMealSuggestionResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/GetMealSuggestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) CreateRiskCommodity(ctx context.Context, in *CreateRiskCommodityRequest, opts ...grpc.CallOption) (*CreateRiskCommodityResponse, error) {
	out := new(CreateRiskCommodityResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/CreateRiskCommodity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) GetRiskList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRiskListResponse, error) {
	out := new(GetRiskListResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/GetRiskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) GetRiskCommodityList(ctx context.Context, in *GetRiskCommodityListRequest, opts ...grpc.CallOption) (*GetRiskCommodityListResponse, error) {
	out := new(GetRiskCommodityListResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/GetRiskCommodityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) EditReportShow(ctx context.Context, in *EditReportShowRequest, opts ...grpc.CallOption) (*EditReportShowResponse, error) {
	out := new(EditReportShowResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/EditReportShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) GetReportShow(ctx context.Context, in *GetReportShowRequest, opts ...grpc.CallOption) (*GetReportShowResponse, error) {
	out := new(GetReportShowResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/GetReportShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) EditReportComparedShow(ctx context.Context, in *EditReportComparedShowRequest, opts ...grpc.CallOption) (*EditReportComparedShowResponse, error) {
	out := new(EditReportComparedShowResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/EditReportComparedShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) GetReportComparedShow(ctx context.Context, in *GetReportComparedShowRequest, opts ...grpc.CallOption) (*GetReportComparedShowResponse, error) {
	out := new(GetReportComparedShowResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/GetReportComparedShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportAPIClient) GetComparisonReportNew(ctx context.Context, in *GetComparisonReportNewRequest, opts ...grpc.CallOption) (*GetComparisonReportNewResponse, error) {
	out := new(GetComparisonReportNewResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.report.v1.ReportAPI/GetComparisonReportNew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportAPIServer is the server API for ReportAPI service.
// All implementations must embed UnimplementedReportAPIServer
// for forward compatibility
type ReportAPIServer interface {
	// SubmitPulseTest 提交采样数据.
	SubmitPulseTest(context.Context, *SubmitPulseTestRequest) (*SubmitPulseTestResponse, error)
	// GetReport 获取阶梯报告.
	GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error)
	// GetMealSuggestion 通过体质获得膳食建议.
	GetMealSuggestion(context.Context, *GetMealSuggestionRequest) (*GetMealSuggestionResponse, error)
	//创建风险推荐商品
	CreateRiskCommodity(context.Context, *CreateRiskCommodityRequest) (*CreateRiskCommodityResponse, error)
	//获取风险列表
	GetRiskList(context.Context, *emptypb.Empty) (*GetRiskListResponse, error)
	//获取推荐商品列表
	GetRiskCommodityList(context.Context, *GetRiskCommodityListRequest) (*GetRiskCommodityListResponse, error)
	//编辑报告显示内容
	EditReportShow(context.Context, *EditReportShowRequest) (*EditReportShowResponse, error)
	//获取报告显示内容
	GetReportShow(context.Context, *GetReportShowRequest) (*GetReportShowResponse, error)
	//编辑报告对比显示功能
	EditReportComparedShow(context.Context, *EditReportComparedShowRequest) (*EditReportComparedShowResponse, error)
	//获取报告对比显示功能
	GetReportComparedShow(context.Context, *GetReportComparedShowRequest) (*GetReportComparedShowResponse, error)
	// 报告对比接口
	GetComparisonReportNew(context.Context, *GetComparisonReportNewRequest) (*GetComparisonReportNewResponse, error)
	mustEmbedUnimplementedReportAPIServer()
}

// UnimplementedReportAPIServer must be embedded to have forward compatible implementations.
type UnimplementedReportAPIServer struct {
}

func (UnimplementedReportAPIServer) SubmitPulseTest(context.Context, *SubmitPulseTestRequest) (*SubmitPulseTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPulseTest not implemented")
}
func (UnimplementedReportAPIServer) GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedReportAPIServer) GetMealSuggestion(context.Context, *GetMealSuggestionRequest) (*GetMealSuggestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMealSuggestion not implemented")
}
func (UnimplementedReportAPIServer) CreateRiskCommodity(context.Context, *CreateRiskCommodityRequest) (*CreateRiskCommodityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRiskCommodity not implemented")
}
func (UnimplementedReportAPIServer) GetRiskList(context.Context, *emptypb.Empty) (*GetRiskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRiskList not implemented")
}
func (UnimplementedReportAPIServer) GetRiskCommodityList(context.Context, *GetRiskCommodityListRequest) (*GetRiskCommodityListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRiskCommodityList not implemented")
}
func (UnimplementedReportAPIServer) EditReportShow(context.Context, *EditReportShowRequest) (*EditReportShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditReportShow not implemented")
}
func (UnimplementedReportAPIServer) GetReportShow(context.Context, *GetReportShowRequest) (*GetReportShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReportShow not implemented")
}
func (UnimplementedReportAPIServer) EditReportComparedShow(context.Context, *EditReportComparedShowRequest) (*EditReportComparedShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditReportComparedShow not implemented")
}
func (UnimplementedReportAPIServer) GetReportComparedShow(context.Context, *GetReportComparedShowRequest) (*GetReportComparedShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReportComparedShow not implemented")
}
func (UnimplementedReportAPIServer) GetComparisonReportNew(context.Context, *GetComparisonReportNewRequest) (*GetComparisonReportNewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComparisonReportNew not implemented")
}
func (UnimplementedReportAPIServer) mustEmbedUnimplementedReportAPIServer() {}

// UnsafeReportAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportAPIServer will
// result in compilation errors.
type UnsafeReportAPIServer interface {
	mustEmbedUnimplementedReportAPIServer()
}

func RegisterReportAPIServer(s *grpc.Server, srv ReportAPIServer) {
	s.RegisterService(&_ReportAPI_serviceDesc, srv)
}

func _ReportAPI_SubmitPulseTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitPulseTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).SubmitPulseTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/SubmitPulseTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).SubmitPulseTest(ctx, req.(*SubmitPulseTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_GetMealSuggestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMealSuggestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).GetMealSuggestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/GetMealSuggestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).GetMealSuggestion(ctx, req.(*GetMealSuggestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_CreateRiskCommodity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRiskCommodityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).CreateRiskCommodity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/CreateRiskCommodity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).CreateRiskCommodity(ctx, req.(*CreateRiskCommodityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_GetRiskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).GetRiskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/GetRiskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).GetRiskList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_GetRiskCommodityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRiskCommodityListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).GetRiskCommodityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/GetRiskCommodityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).GetRiskCommodityList(ctx, req.(*GetRiskCommodityListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_EditReportShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditReportShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).EditReportShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/EditReportShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).EditReportShow(ctx, req.(*EditReportShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_GetReportShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).GetReportShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/GetReportShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).GetReportShow(ctx, req.(*GetReportShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_EditReportComparedShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditReportComparedShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).EditReportComparedShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/EditReportComparedShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).EditReportComparedShow(ctx, req.(*EditReportComparedShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_GetReportComparedShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportComparedShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).GetReportComparedShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/GetReportComparedShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).GetReportComparedShow(ctx, req.(*GetReportComparedShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportAPI_GetComparisonReportNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComparisonReportNewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportAPIServer).GetComparisonReportNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.report.v1.ReportAPI/GetComparisonReportNew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportAPIServer).GetComparisonReportNew(ctx, req.(*GetComparisonReportNewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jthealth.biz.report.v1.ReportAPI",
	HandlerType: (*ReportAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitPulseTest",
			Handler:    _ReportAPI_SubmitPulseTest_Handler,
		},
		{
			MethodName: "GetReport",
			Handler:    _ReportAPI_GetReport_Handler,
		},
		{
			MethodName: "GetMealSuggestion",
			Handler:    _ReportAPI_GetMealSuggestion_Handler,
		},
		{
			MethodName: "CreateRiskCommodity",
			Handler:    _ReportAPI_CreateRiskCommodity_Handler,
		},
		{
			MethodName: "GetRiskList",
			Handler:    _ReportAPI_GetRiskList_Handler,
		},
		{
			MethodName: "GetRiskCommodityList",
			Handler:    _ReportAPI_GetRiskCommodityList_Handler,
		},
		{
			MethodName: "EditReportShow",
			Handler:    _ReportAPI_EditReportShow_Handler,
		},
		{
			MethodName: "GetReportShow",
			Handler:    _ReportAPI_GetReportShow_Handler,
		},
		{
			MethodName: "EditReportComparedShow",
			Handler:    _ReportAPI_EditReportComparedShow_Handler,
		},
		{
			MethodName: "GetReportComparedShow",
			Handler:    _ReportAPI_GetReportComparedShow_Handler,
		},
		{
			MethodName: "GetComparisonReportNew",
			Handler:    _ReportAPI_GetComparisonReportNew_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jthealth/biz/report/v1/report_api.proto",
}
