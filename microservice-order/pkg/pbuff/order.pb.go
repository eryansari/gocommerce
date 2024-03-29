// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: order.proto

package pbuff

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string            `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NoInv       string            `protobuf:"bytes,3,opt,name=no_inv,json=noInv,proto3" json:"no_inv,omitempty"`
	Status      bool              `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	OrderDetail []*OrderDetailDto `protobuf:"bytes,5,rep,name=orderDetail,proto3" json:"orderDetail,omitempty"`
}

func (x *OrderDto) Reset() {
	*x = OrderDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDto) ProtoMessage() {}

func (x *OrderDto) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDto.ProtoReflect.Descriptor instead.
func (*OrderDto) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderDto) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderDto) GetNoInv() string {
	if x != nil {
		return x.NoInv
	}
	return ""
}

func (x *OrderDto) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *OrderDto) GetOrderDetail() []*OrderDetailDto {
	if x != nil {
		return x.OrderDetail
	}
	return nil
}

type OrderDetailDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Qty       int64  `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *OrderDetailDto) Reset() {
	*x = OrderDetailDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetailDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetailDto) ProtoMessage() {}

func (x *OrderDetailDto) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetailDto.ProtoReflect.Descriptor instead.
func (*OrderDetailDto) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderDetailDto) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderDetailDto) GetQty() int64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type OrderDtoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*OrderDto `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *OrderDtoList) Reset() {
	*x = OrderDtoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDtoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDtoList) ProtoMessage() {}

func (x *OrderDtoList) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDtoList.ProtoReflect.Descriptor instead.
func (*OrderDtoList) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderDtoList) GetList() []*OrderDto {
	if x != nil {
		return x.List
	}
	return nil
}

type OrderReqDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order []*OrderDetailDto `protobuf:"bytes,1,rep,name=order,proto3" json:"order,omitempty"`
}

func (x *OrderReqDto) Reset() {
	*x = OrderReqDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReqDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReqDto) ProtoMessage() {}

func (x *OrderReqDto) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReqDto.ProtoReflect.Descriptor instead.
func (*OrderReqDto) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderReqDto) GetOrder() []*OrderDetailDto {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x62, 0x75, 0x66, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x74,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x6f,
	0x5f, 0x69, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x49, 0x6e,
	0x76, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x62, 0x75, 0x66, 0x66, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x44, 0x74, 0x6f, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x22, 0x41, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x44, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x71, 0x74, 0x79, 0x22, 0x33, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x74,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x66, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x74, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x44, 0x74, 0x6f, 0x12, 0x2b, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x66,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x44, 0x74, 0x6f, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0xc6, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x41, 0x70, 0x69, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e,
	0x70, 0x62, 0x75, 0x66, 0x66, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x44, 0x74,
	0x6f, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x66, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x74, 0x6f, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x66, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x66,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42,
	0x1e, 0x5a, 0x1c, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x75, 0x66, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_order_proto_goTypes = []interface{}{
	(*OrderDto)(nil),               // 0: pbuff.OrderDto
	(*OrderDetailDto)(nil),         // 1: pbuff.OrderDetailDto
	(*OrderDtoList)(nil),           // 2: pbuff.OrderDtoList
	(*OrderReqDto)(nil),            // 3: pbuff.OrderReqDto
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
}
var file_order_proto_depIdxs = []int32{
	1, // 0: pbuff.OrderDto.orderDetail:type_name -> pbuff.OrderDetailDto
	0, // 1: pbuff.OrderDtoList.list:type_name -> pbuff.OrderDto
	1, // 2: pbuff.OrderReqDto.order:type_name -> pbuff.OrderDetailDto
	3, // 3: pbuff.OrderApi.Create:input_type -> pbuff.OrderReqDto
	4, // 4: pbuff.OrderApi.GetOrderByID:input_type -> google.protobuf.StringValue
	4, // 5: pbuff.OrderApi.ListOrderByUserID:input_type -> google.protobuf.StringValue
	0, // 6: pbuff.OrderApi.Create:output_type -> pbuff.OrderDto
	0, // 7: pbuff.OrderApi.GetOrderByID:output_type -> pbuff.OrderDto
	2, // 8: pbuff.OrderApi.ListOrderByUserID:output_type -> pbuff.OrderDtoList
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetailDto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDtoList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReqDto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderApiClient is the client API for OrderApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderApiClient interface {
	Create(ctx context.Context, in *OrderReqDto, opts ...grpc.CallOption) (*OrderDto, error)
	GetOrderByID(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*OrderDto, error)
	ListOrderByUserID(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*OrderDtoList, error)
}

type orderApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderApiClient(cc grpc.ClientConnInterface) OrderApiClient {
	return &orderApiClient{cc}
}

func (c *orderApiClient) Create(ctx context.Context, in *OrderReqDto, opts ...grpc.CallOption) (*OrderDto, error) {
	out := new(OrderDto)
	err := c.cc.Invoke(ctx, "/pbuff.OrderApi/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderApiClient) GetOrderByID(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*OrderDto, error) {
	out := new(OrderDto)
	err := c.cc.Invoke(ctx, "/pbuff.OrderApi/GetOrderByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderApiClient) ListOrderByUserID(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*OrderDtoList, error) {
	out := new(OrderDtoList)
	err := c.cc.Invoke(ctx, "/pbuff.OrderApi/ListOrderByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderApiServer is the server API for OrderApi service.
type OrderApiServer interface {
	Create(context.Context, *OrderReqDto) (*OrderDto, error)
	GetOrderByID(context.Context, *wrapperspb.StringValue) (*OrderDto, error)
	ListOrderByUserID(context.Context, *wrapperspb.StringValue) (*OrderDtoList, error)
}

// UnimplementedOrderApiServer can be embedded to have forward compatible implementations.
type UnimplementedOrderApiServer struct {
}

func (*UnimplementedOrderApiServer) Create(context.Context, *OrderReqDto) (*OrderDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedOrderApiServer) GetOrderByID(context.Context, *wrapperspb.StringValue) (*OrderDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByID not implemented")
}
func (*UnimplementedOrderApiServer) ListOrderByUserID(context.Context, *wrapperspb.StringValue) (*OrderDtoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrderByUserID not implemented")
}

func RegisterOrderApiServer(s *grpc.Server, srv OrderApiServer) {
	s.RegisterService(&_OrderApi_serviceDesc, srv)
}

func _OrderApi_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReqDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderApiServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbuff.OrderApi/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderApiServer).Create(ctx, req.(*OrderReqDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderApi_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderApiServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbuff.OrderApi/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderApiServer).GetOrderByID(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderApi_ListOrderByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderApiServer).ListOrderByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbuff.OrderApi/ListOrderByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderApiServer).ListOrderByUserID(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbuff.OrderApi",
	HandlerType: (*OrderApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderApi_Create_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _OrderApi_GetOrderByID_Handler,
		},
		{
			MethodName: "ListOrderByUserID",
			Handler:    _OrderApi_ListOrderByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
