// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basket_service.proto

package basket_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type CreateBasketRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBasketRequest) Reset()         { *m = CreateBasketRequest{} }
func (m *CreateBasketRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBasketRequest) ProtoMessage()    {}
func (*CreateBasketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64f990a99596a990, []int{0}
}

func (m *CreateBasketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBasketRequest.Unmarshal(m, b)
}
func (m *CreateBasketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBasketRequest.Marshal(b, m, deterministic)
}
func (m *CreateBasketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBasketRequest.Merge(m, src)
}
func (m *CreateBasketRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBasketRequest.Size(m)
}
func (m *CreateBasketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBasketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBasketRequest proto.InternalMessageInfo

type CreateBasketResponse struct {
	BasketId             int64    `protobuf:"varint,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBasketResponse) Reset()         { *m = CreateBasketResponse{} }
func (m *CreateBasketResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBasketResponse) ProtoMessage()    {}
func (*CreateBasketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64f990a99596a990, []int{1}
}

func (m *CreateBasketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBasketResponse.Unmarshal(m, b)
}
func (m *CreateBasketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBasketResponse.Marshal(b, m, deterministic)
}
func (m *CreateBasketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBasketResponse.Merge(m, src)
}
func (m *CreateBasketResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBasketResponse.Size(m)
}
func (m *CreateBasketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBasketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBasketResponse proto.InternalMessageInfo

func (m *CreateBasketResponse) GetBasketId() int64 {
	if m != nil {
		return m.BasketId
	}
	return 0
}

type AddProductRequest struct {
	BasketId             int64    `protobuf:"varint,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	ProductId            string   `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProductRequest) Reset()         { *m = AddProductRequest{} }
func (m *AddProductRequest) String() string { return proto.CompactTextString(m) }
func (*AddProductRequest) ProtoMessage()    {}
func (*AddProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64f990a99596a990, []int{2}
}

func (m *AddProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProductRequest.Unmarshal(m, b)
}
func (m *AddProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProductRequest.Marshal(b, m, deterministic)
}
func (m *AddProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProductRequest.Merge(m, src)
}
func (m *AddProductRequest) XXX_Size() int {
	return xxx_messageInfo_AddProductRequest.Size(m)
}
func (m *AddProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddProductRequest proto.InternalMessageInfo

func (m *AddProductRequest) GetBasketId() int64 {
	if m != nil {
		return m.BasketId
	}
	return 0
}

func (m *AddProductRequest) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

type AddProductResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProductResponse) Reset()         { *m = AddProductResponse{} }
func (m *AddProductResponse) String() string { return proto.CompactTextString(m) }
func (*AddProductResponse) ProtoMessage()    {}
func (*AddProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64f990a99596a990, []int{3}
}

func (m *AddProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProductResponse.Unmarshal(m, b)
}
func (m *AddProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProductResponse.Marshal(b, m, deterministic)
}
func (m *AddProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProductResponse.Merge(m, src)
}
func (m *AddProductResponse) XXX_Size() int {
	return xxx_messageInfo_AddProductResponse.Size(m)
}
func (m *AddProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddProductResponse proto.InternalMessageInfo

type GetTotalAmountRequest struct {
	BasketId             int64    `protobuf:"varint,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTotalAmountRequest) Reset()         { *m = GetTotalAmountRequest{} }
func (m *GetTotalAmountRequest) String() string { return proto.CompactTextString(m) }
func (*GetTotalAmountRequest) ProtoMessage()    {}
func (*GetTotalAmountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64f990a99596a990, []int{4}
}

func (m *GetTotalAmountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTotalAmountRequest.Unmarshal(m, b)
}
func (m *GetTotalAmountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTotalAmountRequest.Marshal(b, m, deterministic)
}
func (m *GetTotalAmountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTotalAmountRequest.Merge(m, src)
}
func (m *GetTotalAmountRequest) XXX_Size() int {
	return xxx_messageInfo_GetTotalAmountRequest.Size(m)
}
func (m *GetTotalAmountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTotalAmountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTotalAmountRequest proto.InternalMessageInfo

func (m *GetTotalAmountRequest) GetBasketId() int64 {
	if m != nil {
		return m.BasketId
	}
	return 0
}

type GetTotalAmountResponse struct {
	Price                float32  `protobuf:"fixed32,1,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTotalAmountResponse) Reset()         { *m = GetTotalAmountResponse{} }
func (m *GetTotalAmountResponse) String() string { return proto.CompactTextString(m) }
func (*GetTotalAmountResponse) ProtoMessage()    {}
func (*GetTotalAmountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64f990a99596a990, []int{5}
}

func (m *GetTotalAmountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTotalAmountResponse.Unmarshal(m, b)
}
func (m *GetTotalAmountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTotalAmountResponse.Marshal(b, m, deterministic)
}
func (m *GetTotalAmountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTotalAmountResponse.Merge(m, src)
}
func (m *GetTotalAmountResponse) XXX_Size() int {
	return xxx_messageInfo_GetTotalAmountResponse.Size(m)
}
func (m *GetTotalAmountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTotalAmountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTotalAmountResponse proto.InternalMessageInfo

func (m *GetTotalAmountResponse) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type RemoveBasketRequest struct {
	BasketId             int64    `protobuf:"varint,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveBasketRequest) Reset()         { *m = RemoveBasketRequest{} }
func (m *RemoveBasketRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveBasketRequest) ProtoMessage()    {}
func (*RemoveBasketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64f990a99596a990, []int{6}
}

func (m *RemoveBasketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveBasketRequest.Unmarshal(m, b)
}
func (m *RemoveBasketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveBasketRequest.Marshal(b, m, deterministic)
}
func (m *RemoveBasketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveBasketRequest.Merge(m, src)
}
func (m *RemoveBasketRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveBasketRequest.Size(m)
}
func (m *RemoveBasketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveBasketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveBasketRequest proto.InternalMessageInfo

func (m *RemoveBasketRequest) GetBasketId() int64 {
	if m != nil {
		return m.BasketId
	}
	return 0
}

type RemoveBasketResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveBasketResponse) Reset()         { *m = RemoveBasketResponse{} }
func (m *RemoveBasketResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveBasketResponse) ProtoMessage()    {}
func (*RemoveBasketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64f990a99596a990, []int{7}
}

func (m *RemoveBasketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveBasketResponse.Unmarshal(m, b)
}
func (m *RemoveBasketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveBasketResponse.Marshal(b, m, deterministic)
}
func (m *RemoveBasketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveBasketResponse.Merge(m, src)
}
func (m *RemoveBasketResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveBasketResponse.Size(m)
}
func (m *RemoveBasketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveBasketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveBasketResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateBasketRequest)(nil), "prueba_cabify.CreateBasketRequest")
	proto.RegisterType((*CreateBasketResponse)(nil), "prueba_cabify.CreateBasketResponse")
	proto.RegisterType((*AddProductRequest)(nil), "prueba_cabify.AddProductRequest")
	proto.RegisterType((*AddProductResponse)(nil), "prueba_cabify.AddProductResponse")
	proto.RegisterType((*GetTotalAmountRequest)(nil), "prueba_cabify.GetTotalAmountRequest")
	proto.RegisterType((*GetTotalAmountResponse)(nil), "prueba_cabify.GetTotalAmountResponse")
	proto.RegisterType((*RemoveBasketRequest)(nil), "prueba_cabify.RemoveBasketRequest")
	proto.RegisterType((*RemoveBasketResponse)(nil), "prueba_cabify.RemoveBasketResponse")
}

func init() { proto.RegisterFile("basket_service.proto", fileDescriptor_64f990a99596a990) }

var fileDescriptor_64f990a99596a990 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0x5b, 0x8c, 0x46, 0x26, 0x6d, 0xa3, 0x5b, 0xda, 0x18, 0x8c, 0x09, 0xae, 0x9a, 0xf4,
	0xc4, 0xa1, 0xf5, 0x0f, 0xb4, 0x1e, 0x4c, 0x4f, 0x1a, 0xea, 0xa5, 0x5e, 0x08, 0xb0, 0x63, 0x42,
	0xb4, 0x5d, 0x5c, 0x96, 0x26, 0xfe, 0x76, 0x2f, 0x46, 0x16, 0x2b, 0x0b, 0x04, 0x3d, 0x32, 0xbc,
	0x37, 0x0f, 0xbe, 0x97, 0x01, 0x2b, 0x0c, 0xd2, 0x57, 0x94, 0x7e, 0x8a, 0x62, 0x17, 0x47, 0xe8,
	0x26, 0x82, 0x4b, 0x4e, 0xfa, 0x89, 0xc8, 0x30, 0x0c, 0xfc, 0x28, 0x08, 0xe3, 0x97, 0x0f, 0x3a,
	0x82, 0xe1, 0x9d, 0xc0, 0x40, 0xe2, 0x22, 0x17, 0x7b, 0xf8, 0x9e, 0x61, 0x2a, 0xe9, 0x0c, 0x2c,
	0x7d, 0x9c, 0x26, 0x7c, 0x9b, 0x22, 0x39, 0x07, 0xb3, 0xd8, 0x1a, 0xb3, 0xb3, 0xae, 0xd3, 0x9d,
	0x1c, 0x78, 0xc7, 0x6a, 0xb0, 0x64, 0xf4, 0x01, 0x4e, 0xe7, 0x8c, 0x3d, 0x0a, 0xce, 0xb2, 0xe8,
	0x67, 0x53, 0xab, 0x83, 0x5c, 0x00, 0x24, 0x4a, 0xfe, 0xfd, 0xd6, 0x70, 0xba, 0x13, 0xd3, 0x33,
	0x8b, 0xc9, 0x92, 0x51, 0x0b, 0x48, 0x79, 0xa1, 0xfa, 0x06, 0x7a, 0x0b, 0xa3, 0x7b, 0x94, 0x4f,
	0x5c, 0x06, 0x6f, 0xf3, 0x0d, 0xcf, 0xb6, 0xff, 0x8a, 0xa2, 0x2e, 0x8c, 0xab, 0xae, 0xe2, 0x9f,
	0x2c, 0x38, 0x4c, 0x44, 0x1c, 0x61, 0x6e, 0x31, 0x3c, 0xf5, 0x40, 0xa7, 0x30, 0xf4, 0x70, 0xc3,
	0x77, 0x3a, 0x98, 0xf6, 0x8c, 0x31, 0x58, 0xba, 0x47, 0x25, 0x4c, 0x3f, 0x0d, 0xe8, 0xab, 0xd1,
	0x4a, 0x75, 0x41, 0xd6, 0xd0, 0x2b, 0xf3, 0x25, 0xd4, 0xd5, 0x6a, 0x71, 0x1b, 0x3a, 0xb1, 0xaf,
	0x5a, 0x35, 0x05, 0x9c, 0x0e, 0x59, 0x01, 0xfc, 0x42, 0x23, 0x4e, 0xc5, 0x54, 0x2b, 0xc8, 0xbe,
	0x6c, 0x51, 0xec, 0x97, 0xfa, 0x30, 0xd0, 0xe9, 0x91, 0xeb, 0x8a, 0xad, 0xb1, 0x12, 0xfb, 0xe6,
	0x0f, 0xd5, 0x3e, 0x60, 0x0d, 0xbd, 0x32, 0xba, 0x1a, 0x90, 0x86, 0x2e, 0x6a, 0x40, 0x9a, 0xd8,
	0xd3, 0xce, 0xe2, 0xe4, 0x79, 0xa0, 0x5f, 0x42, 0x78, 0x94, 0x9f, 0xc2, 0xec, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0xf7, 0xc2, 0x97, 0x22, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BasketServiceClient is the client API for BasketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasketServiceClient interface {
	CreateBasket(ctx context.Context, in *CreateBasketRequest, opts ...grpc.CallOption) (*CreateBasketResponse, error)
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error)
	GetTotalAmount(ctx context.Context, in *GetTotalAmountRequest, opts ...grpc.CallOption) (*GetTotalAmountResponse, error)
	RemoveBasket(ctx context.Context, in *RemoveBasketRequest, opts ...grpc.CallOption) (*RemoveBasketResponse, error)
}

type basketServiceClient struct {
	cc *grpc.ClientConn
}

func NewBasketServiceClient(cc *grpc.ClientConn) BasketServiceClient {
	return &basketServiceClient{cc}
}

func (c *basketServiceClient) CreateBasket(ctx context.Context, in *CreateBasketRequest, opts ...grpc.CallOption) (*CreateBasketResponse, error) {
	out := new(CreateBasketResponse)
	err := c.cc.Invoke(ctx, "/prueba_cabify.BasketService/CreateBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error) {
	out := new(AddProductResponse)
	err := c.cc.Invoke(ctx, "/prueba_cabify.BasketService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) GetTotalAmount(ctx context.Context, in *GetTotalAmountRequest, opts ...grpc.CallOption) (*GetTotalAmountResponse, error) {
	out := new(GetTotalAmountResponse)
	err := c.cc.Invoke(ctx, "/prueba_cabify.BasketService/GetTotalAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) RemoveBasket(ctx context.Context, in *RemoveBasketRequest, opts ...grpc.CallOption) (*RemoveBasketResponse, error) {
	out := new(RemoveBasketResponse)
	err := c.cc.Invoke(ctx, "/prueba_cabify.BasketService/RemoveBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasketServiceServer is the server API for BasketService service.
type BasketServiceServer interface {
	CreateBasket(context.Context, *CreateBasketRequest) (*CreateBasketResponse, error)
	AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error)
	GetTotalAmount(context.Context, *GetTotalAmountRequest) (*GetTotalAmountResponse, error)
	RemoveBasket(context.Context, *RemoveBasketRequest) (*RemoveBasketResponse, error)
}

// UnimplementedBasketServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBasketServiceServer struct {
}

func (*UnimplementedBasketServiceServer) CreateBasket(ctx context.Context, req *CreateBasketRequest) (*CreateBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBasket not implemented")
}
func (*UnimplementedBasketServiceServer) AddProduct(ctx context.Context, req *AddProductRequest) (*AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (*UnimplementedBasketServiceServer) GetTotalAmount(ctx context.Context, req *GetTotalAmountRequest) (*GetTotalAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalAmount not implemented")
}
func (*UnimplementedBasketServiceServer) RemoveBasket(ctx context.Context, req *RemoveBasketRequest) (*RemoveBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBasket not implemented")
}

func RegisterBasketServiceServer(s *grpc.Server, srv BasketServiceServer) {
	s.RegisterService(&_BasketService_serviceDesc, srv)
}

func _BasketService_CreateBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).CreateBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prueba_cabify.BasketService/CreateBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).CreateBasket(ctx, req.(*CreateBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prueba_cabify.BasketService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_GetTotalAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).GetTotalAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prueba_cabify.BasketService/GetTotalAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).GetTotalAmount(ctx, req.(*GetTotalAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_RemoveBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).RemoveBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prueba_cabify.BasketService/RemoveBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).RemoveBasket(ctx, req.(*RemoveBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BasketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prueba_cabify.BasketService",
	HandlerType: (*BasketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBasket",
			Handler:    _BasketService_CreateBasket_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _BasketService_AddProduct_Handler,
		},
		{
			MethodName: "GetTotalAmount",
			Handler:    _BasketService_GetTotalAmount_Handler,
		},
		{
			MethodName: "RemoveBasket",
			Handler:    _BasketService_RemoveBasket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basket_service.proto",
}
