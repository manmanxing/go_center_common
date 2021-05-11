// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderservice.proto

package orderservice

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PingRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_orderservice_2cc384490b6f5679, []int{0}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PingReply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_orderservice_2cc384490b6f5679, []int{1}
}
func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (dst *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(dst, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type CreateOrderRequest struct {
	UserId               int64              `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AppId                string             `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	MchId                string             `protobuf:"bytes,3,opt,name=mch_id,json=mchId,proto3" json:"mch_id,omitempty"`
	SubAppId             string             `protobuf:"bytes,4,opt,name=sub_app_id,json=subAppId,proto3" json:"sub_app_id,omitempty"`
	SubMchId             string             `protobuf:"bytes,5,opt,name=sub_mch_id,json=subMchId,proto3" json:"sub_mch_id,omitempty"`
	BusType              string             `protobuf:"bytes,6,opt,name=bus_type,json=busType,proto3" json:"bus_type,omitempty"`
	SourceType           string             `protobuf:"bytes,7,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	ChannelType          string             `protobuf:"bytes,8,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	ReceiptSymbol        string             `protobuf:"bytes,9,opt,name=receipt_symbol,json=receiptSymbol,proto3" json:"receipt_symbol,omitempty"`
	BusTradeNo           string             `protobuf:"bytes,10,opt,name=bus_trade_no,json=busTradeNo,proto3" json:"bus_trade_no,omitempty"`
	BusProducts          []*BusProductsList `protobuf:"bytes,11,rep,name=bus_products,json=busProducts,proto3" json:"bus_products,omitempty"`
	PayType              string             `protobuf:"bytes,12,opt,name=payType,proto3" json:"payType,omitempty"`
	TransactionMode      int64              `protobuf:"varint,13,opt,name=transaction_mode,json=transactionMode,proto3" json:"transaction_mode,omitempty"`
	Currency             string             `protobuf:"bytes,14,opt,name=currency,proto3" json:"currency,omitempty"`
	TradeAmount          int64              `protobuf:"varint,15,opt,name=trade_amount,json=tradeAmount,proto3" json:"trade_amount,omitempty"`
	PayAmount            int64              `protobuf:"varint,16,opt,name=pay_amount,json=payAmount,proto3" json:"pay_amount,omitempty"`
	Freight              int64              `protobuf:"varint,17,opt,name=freight,proto3" json:"freight,omitempty"`
	MerchantCoupon       int64              `protobuf:"varint,18,opt,name=merchant_coupon,json=merchantCoupon,proto3" json:"merchant_coupon,omitempty"`
	PlatformCoupon       int64              `protobuf:"varint,19,opt,name=platform_coupon,json=platformCoupon,proto3" json:"platform_coupon,omitempty"`
	TradeDescription     string             `protobuf:"bytes,20,opt,name=trade_description,json=tradeDescription,proto3" json:"trade_description,omitempty"`
	ExpireTime           string             `protobuf:"bytes,21,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	Remark               string             `protobuf:"bytes,22,opt,name=remark,proto3" json:"remark,omitempty"`
	Sign                 string             `protobuf:"bytes,23,opt,name=sign,proto3" json:"sign,omitempty"`
	Timestamp            int64              `protobuf:"varint,24,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_orderservice_2cc384490b6f5679, []int{2}
}
func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (dst *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(dst, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateOrderRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CreateOrderRequest) GetMchId() string {
	if m != nil {
		return m.MchId
	}
	return ""
}

func (m *CreateOrderRequest) GetSubAppId() string {
	if m != nil {
		return m.SubAppId
	}
	return ""
}

func (m *CreateOrderRequest) GetSubMchId() string {
	if m != nil {
		return m.SubMchId
	}
	return ""
}

func (m *CreateOrderRequest) GetBusType() string {
	if m != nil {
		return m.BusType
	}
	return ""
}

func (m *CreateOrderRequest) GetSourceType() string {
	if m != nil {
		return m.SourceType
	}
	return ""
}

func (m *CreateOrderRequest) GetChannelType() string {
	if m != nil {
		return m.ChannelType
	}
	return ""
}

func (m *CreateOrderRequest) GetReceiptSymbol() string {
	if m != nil {
		return m.ReceiptSymbol
	}
	return ""
}

func (m *CreateOrderRequest) GetBusTradeNo() string {
	if m != nil {
		return m.BusTradeNo
	}
	return ""
}

func (m *CreateOrderRequest) GetBusProducts() []*BusProductsList {
	if m != nil {
		return m.BusProducts
	}
	return nil
}

func (m *CreateOrderRequest) GetPayType() string {
	if m != nil {
		return m.PayType
	}
	return ""
}

func (m *CreateOrderRequest) GetTransactionMode() int64 {
	if m != nil {
		return m.TransactionMode
	}
	return 0
}

func (m *CreateOrderRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *CreateOrderRequest) GetTradeAmount() int64 {
	if m != nil {
		return m.TradeAmount
	}
	return 0
}

func (m *CreateOrderRequest) GetPayAmount() int64 {
	if m != nil {
		return m.PayAmount
	}
	return 0
}

func (m *CreateOrderRequest) GetFreight() int64 {
	if m != nil {
		return m.Freight
	}
	return 0
}

func (m *CreateOrderRequest) GetMerchantCoupon() int64 {
	if m != nil {
		return m.MerchantCoupon
	}
	return 0
}

func (m *CreateOrderRequest) GetPlatformCoupon() int64 {
	if m != nil {
		return m.PlatformCoupon
	}
	return 0
}

func (m *CreateOrderRequest) GetTradeDescription() string {
	if m != nil {
		return m.TradeDescription
	}
	return ""
}

func (m *CreateOrderRequest) GetExpireTime() string {
	if m != nil {
		return m.ExpireTime
	}
	return ""
}

func (m *CreateOrderRequest) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *CreateOrderRequest) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *CreateOrderRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type BusProductsList struct {
	BusType              string   `protobuf:"bytes,1,opt,name=bus_type,json=busType,proto3" json:"bus_type,omitempty"`
	ProductType          string   `protobuf:"bytes,2,opt,name=product_type,json=productType,proto3" json:"product_type,omitempty"`
	CommoditySkuId       string   `protobuf:"bytes,3,opt,name=commodity_sku_id,json=commoditySkuId,proto3" json:"commodity_sku_id,omitempty"`
	BusOrderNo           string   `protobuf:"bytes,4,opt,name=bus_order_no,json=busOrderNo,proto3" json:"bus_order_no,omitempty"`
	GoodsAmount          int64    `protobuf:"varint,5,opt,name=goods_amount,json=goodsAmount,proto3" json:"goods_amount,omitempty"`
	Amount               int64    `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	MerchantCoupon       int64    `protobuf:"varint,7,opt,name=merchant_coupon,json=merchantCoupon,proto3" json:"merchant_coupon,omitempty"`
	PlatformCoupon       int64    `protobuf:"varint,8,opt,name=platform_coupon,json=platformCoupon,proto3" json:"platform_coupon,omitempty"`
	Freight              int64    `protobuf:"varint,9,opt,name=freight,proto3" json:"freight,omitempty"`
	SupplierName         int64    `protobuf:"varint,10,opt,name=supplier_name,json=supplierName,proto3" json:"supplier_name,omitempty"`
	PurchaseAmount       int64    `protobuf:"varint,11,opt,name=purchase_amount,json=purchaseAmount,proto3" json:"purchase_amount,omitempty"`
	Remark               string   `protobuf:"bytes,12,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BusProductsList) Reset()         { *m = BusProductsList{} }
func (m *BusProductsList) String() string { return proto.CompactTextString(m) }
func (*BusProductsList) ProtoMessage()    {}
func (*BusProductsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_orderservice_2cc384490b6f5679, []int{3}
}
func (m *BusProductsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BusProductsList.Unmarshal(m, b)
}
func (m *BusProductsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BusProductsList.Marshal(b, m, deterministic)
}
func (dst *BusProductsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BusProductsList.Merge(dst, src)
}
func (m *BusProductsList) XXX_Size() int {
	return xxx_messageInfo_BusProductsList.Size(m)
}
func (m *BusProductsList) XXX_DiscardUnknown() {
	xxx_messageInfo_BusProductsList.DiscardUnknown(m)
}

var xxx_messageInfo_BusProductsList proto.InternalMessageInfo

func (m *BusProductsList) GetBusType() string {
	if m != nil {
		return m.BusType
	}
	return ""
}

func (m *BusProductsList) GetProductType() string {
	if m != nil {
		return m.ProductType
	}
	return ""
}

func (m *BusProductsList) GetCommoditySkuId() string {
	if m != nil {
		return m.CommoditySkuId
	}
	return ""
}

func (m *BusProductsList) GetBusOrderNo() string {
	if m != nil {
		return m.BusOrderNo
	}
	return ""
}

func (m *BusProductsList) GetGoodsAmount() int64 {
	if m != nil {
		return m.GoodsAmount
	}
	return 0
}

func (m *BusProductsList) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *BusProductsList) GetMerchantCoupon() int64 {
	if m != nil {
		return m.MerchantCoupon
	}
	return 0
}

func (m *BusProductsList) GetPlatformCoupon() int64 {
	if m != nil {
		return m.PlatformCoupon
	}
	return 0
}

func (m *BusProductsList) GetFreight() int64 {
	if m != nil {
		return m.Freight
	}
	return 0
}

func (m *BusProductsList) GetSupplierName() int64 {
	if m != nil {
		return m.SupplierName
	}
	return 0
}

func (m *BusProductsList) GetPurchaseAmount() int64 {
	if m != nil {
		return m.PurchaseAmount
	}
	return 0
}

func (m *BusProductsList) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type CreateOrderReply struct {
	TradeNo              string         `protobuf:"bytes,1,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`
	BusTradeNo           string         `protobuf:"bytes,2,opt,name=bus_trade_no,json=busTradeNo,proto3" json:"bus_trade_no,omitempty"`
	PayNo                string         `protobuf:"bytes,3,opt,name=pay_no,json=payNo,proto3" json:"pay_no,omitempty"`
	OrderNos             []*OrderNoList `protobuf:"bytes,4,rep,name=orderNos,proto3" json:"orderNos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateOrderReply) Reset()         { *m = CreateOrderReply{} }
func (m *CreateOrderReply) String() string { return proto.CompactTextString(m) }
func (*CreateOrderReply) ProtoMessage()    {}
func (*CreateOrderReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_orderservice_2cc384490b6f5679, []int{4}
}
func (m *CreateOrderReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderReply.Unmarshal(m, b)
}
func (m *CreateOrderReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderReply.Marshal(b, m, deterministic)
}
func (dst *CreateOrderReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderReply.Merge(dst, src)
}
func (m *CreateOrderReply) XXX_Size() int {
	return xxx_messageInfo_CreateOrderReply.Size(m)
}
func (m *CreateOrderReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderReply proto.InternalMessageInfo

func (m *CreateOrderReply) GetTradeNo() string {
	if m != nil {
		return m.TradeNo
	}
	return ""
}

func (m *CreateOrderReply) GetBusTradeNo() string {
	if m != nil {
		return m.BusTradeNo
	}
	return ""
}

func (m *CreateOrderReply) GetPayNo() string {
	if m != nil {
		return m.PayNo
	}
	return ""
}

func (m *CreateOrderReply) GetOrderNos() []*OrderNoList {
	if m != nil {
		return m.OrderNos
	}
	return nil
}

type OrderNoList struct {
	OrderNo              string   `protobuf:"bytes,1,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	BusOrderNo           string   `protobuf:"bytes,2,opt,name=bus_order_no,json=busOrderNo,proto3" json:"bus_order_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderNoList) Reset()         { *m = OrderNoList{} }
func (m *OrderNoList) String() string { return proto.CompactTextString(m) }
func (*OrderNoList) ProtoMessage()    {}
func (*OrderNoList) Descriptor() ([]byte, []int) {
	return fileDescriptor_orderservice_2cc384490b6f5679, []int{5}
}
func (m *OrderNoList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNoList.Unmarshal(m, b)
}
func (m *OrderNoList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNoList.Marshal(b, m, deterministic)
}
func (dst *OrderNoList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNoList.Merge(dst, src)
}
func (m *OrderNoList) XXX_Size() int {
	return xxx_messageInfo_OrderNoList.Size(m)
}
func (m *OrderNoList) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNoList.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNoList proto.InternalMessageInfo

func (m *OrderNoList) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *OrderNoList) GetBusOrderNo() string {
	if m != nil {
		return m.BusOrderNo
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "orderservice.PingRequest")
	proto.RegisterType((*PingReply)(nil), "orderservice.PingReply")
	proto.RegisterType((*CreateOrderRequest)(nil), "orderservice.CreateOrderRequest")
	proto.RegisterType((*BusProductsList)(nil), "orderservice.busProductsList")
	proto.RegisterType((*CreateOrderReply)(nil), "orderservice.CreateOrderReply")
	proto.RegisterType((*OrderNoList)(nil), "orderservice.orderNoList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	// ping 测试连通性
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	// 创建订单
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderReply, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/orderservice.OrderService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderReply, error) {
	out := new(CreateOrderReply)
	err := c.cc.Invoke(ctx, "/orderservice.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	// ping 测试连通性
	Ping(context.Context, *PingRequest) (*PingReply, error)
	// 创建订单
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderReply, error)
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderservice.OrderService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderservice.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderservice.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _OrderService_Ping_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orderservice.proto",
}

func init() { proto.RegisterFile("orderservice.proto", fileDescriptor_orderservice_2cc384490b6f5679) }

var fileDescriptor_orderservice_2cc384490b6f5679 = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcf, 0x6f, 0xdb, 0x36,
	0x18, 0x9d, 0x63, 0x47, 0xb6, 0x3f, 0x29, 0x89, 0xcb, 0x2d, 0x09, 0x13, 0x34, 0x6b, 0xea, 0x61,
	0x58, 0x86, 0x01, 0x3d, 0x74, 0xd8, 0x6d, 0x87, 0x15, 0xdd, 0x25, 0xc3, 0x9a, 0x75, 0x4e, 0xef,
	0x02, 0x2d, 0xb1, 0xb6, 0x10, 0x53, 0xe4, 0xf8, 0x63, 0x98, 0xfe, 0x98, 0x61, 0xd7, 0xfe, 0x99,
	0x03, 0x3f, 0x52, 0x8a, 0xec, 0x18, 0xc3, 0x6e, 0xe2, 0x7b, 0x8f, 0xe4, 0x47, 0xbe, 0xef, 0x51,
	0x40, 0xa4, 0x2e, 0xb9, 0x36, 0x5c, 0xff, 0x59, 0x15, 0xfc, 0x95, 0xd2, 0xd2, 0x4a, 0x92, 0xf5,
	0xb1, 0xf9, 0x0b, 0x48, 0xdf, 0x57, 0xf5, 0x6a, 0xc1, 0xff, 0x70, 0xdc, 0x58, 0x32, 0x83, 0xa1,
	0x30, 0x2b, 0x3a, 0xb8, 0x1e, 0xdc, 0x4c, 0x17, 0xfe, 0x73, 0x7e, 0x05, 0xd3, 0x20, 0x50, 0x9b,
	0x66, 0x0f, 0xfd, 0x29, 0x01, 0xf2, 0x56, 0x73, 0x66, 0xf9, 0x6f, 0x7e, 0xd9, 0x76, 0x9d, 0x73,
	0x18, 0x3b, 0xc3, 0x75, 0x5e, 0x95, 0x28, 0x1e, 0x2e, 0x12, 0x3f, 0xbc, 0x2d, 0xc9, 0x29, 0x24,
	0x4c, 0x29, 0x8f, 0x1f, 0xe0, 0x22, 0x87, 0x4c, 0xa9, 0x00, 0x8b, 0x62, 0xed, 0xe1, 0x61, 0x80,
	0x45, 0xb1, 0xbe, 0x2d, 0xc9, 0x73, 0x00, 0xe3, 0x96, 0x79, 0x9c, 0x31, 0x42, 0x6a, 0x62, 0xdc,
	0xf2, 0x0d, 0x4e, 0x8a, 0x6c, 0x9c, 0x78, 0xd8, 0xb1, 0xef, 0x70, 0xee, 0x05, 0x4c, 0x96, 0xce,
	0xe4, 0xb6, 0x51, 0x9c, 0x26, 0xc8, 0x8d, 0x97, 0xce, 0x7c, 0x68, 0x14, 0x27, 0x2f, 0x20, 0x35,
	0xd2, 0xe9, 0x82, 0x07, 0x76, 0x8c, 0x2c, 0x04, 0x08, 0x05, 0x2f, 0x21, 0x2b, 0xd6, 0xac, 0xae,
	0xf9, 0x26, 0x28, 0x26, 0xa8, 0x48, 0x23, 0x86, 0x92, 0xaf, 0xe1, 0x58, 0xf3, 0x82, 0x57, 0xca,
	0xe6, 0xa6, 0x11, 0x4b, 0xb9, 0xa1, 0x53, 0x14, 0x1d, 0x45, 0xf4, 0x1e, 0x41, 0x72, 0x0d, 0x19,
	0x56, 0xa1, 0x59, 0xc9, 0xf3, 0x5a, 0x52, 0x08, 0x7b, 0xf9, 0x4a, 0x3c, 0x74, 0x27, 0xc9, 0x4f,
	0x41, 0xa1, 0xb4, 0x2c, 0x5d, 0x61, 0x0d, 0x4d, 0xaf, 0x87, 0x37, 0xe9, 0xeb, 0xab, 0x57, 0x5b,
	0xd6, 0x2d, 0x9d, 0x79, 0x1f, 0x05, 0xbf, 0x56, 0xc6, 0x2e, 0xd2, 0x1e, 0x40, 0x28, 0x8c, 0x15,
	0x6b, 0x7c, 0x55, 0x34, 0x0b, 0x07, 0x8d, 0x43, 0xf2, 0x2d, 0xcc, 0xac, 0x66, 0xb5, 0x61, 0x85,
	0xad, 0x64, 0x9d, 0x0b, 0x59, 0x72, 0x7a, 0x84, 0x7e, 0x9c, 0xf4, 0xf0, 0x77, 0xb2, 0xe4, 0xe4,
	0x12, 0x26, 0x85, 0xd3, 0x9a, 0xd7, 0x45, 0x43, 0x8f, 0xc3, 0x55, 0xb6, 0x63, 0x7f, 0x1d, 0xe1,
	0x00, 0x4c, 0x48, 0x57, 0x5b, 0x7a, 0x82, 0x4b, 0xa4, 0x88, 0xbd, 0x41, 0x88, 0x5c, 0x01, 0x28,
	0xd6, 0xb4, 0x82, 0x19, 0x0a, 0xa6, 0x8a, 0x35, 0x91, 0xa6, 0x30, 0xfe, 0xa8, 0x79, 0xb5, 0x5a,
	0x5b, 0xfa, 0x0c, 0xb9, 0x76, 0x48, 0xbe, 0x81, 0x13, 0xc1, 0xb5, 0xbf, 0x59, 0x9b, 0x17, 0xd2,
	0x29, 0x59, 0x53, 0x82, 0x8a, 0xe3, 0x16, 0x7e, 0x8b, 0xa8, 0x17, 0xaa, 0x0d, 0xb3, 0x1f, 0xa5,
	0x16, 0xad, 0xf0, 0xf3, 0x20, 0x6c, 0xe1, 0x28, 0xfc, 0x0e, 0x9e, 0x85, 0x6a, 0x4b, 0x6e, 0x0a,
	0x5d, 0x29, 0x7f, 0x44, 0xfa, 0x05, 0x1e, 0x69, 0x86, 0xc4, 0xcf, 0x8f, 0xb8, 0x6f, 0x05, 0xfe,
	0x97, 0xaa, 0x34, 0xcf, 0x6d, 0x25, 0x38, 0x3d, 0x0d, 0xf6, 0x04, 0xe8, 0x43, 0x25, 0x38, 0x39,
	0x83, 0x44, 0x73, 0xc1, 0xf4, 0x03, 0x3d, 0x43, 0x2e, 0x8e, 0x08, 0x81, 0x91, 0xa9, 0x56, 0x35,
	0x3d, 0x47, 0x14, 0xbf, 0xc9, 0x73, 0x98, 0xfa, 0x55, 0x8c, 0x65, 0x42, 0x51, 0x1a, 0xee, 0xa0,
	0x03, 0xe6, 0x9f, 0x86, 0x70, 0xb2, 0xe3, 0xe3, 0x56, 0x93, 0x0e, 0xb6, 0x9b, 0xf4, 0x25, 0x64,
	0xb1, 0x27, 0x02, 0x1d, 0xf2, 0x92, 0x46, 0x0c, 0x25, 0x37, 0x30, 0x2b, 0xa4, 0x10, 0xb2, 0xac,
	0x6c, 0x93, 0x9b, 0x07, 0xf7, 0x98, 0x9f, 0xe3, 0x0e, 0xbf, 0x7f, 0x70, 0xb7, 0x65, 0xdb, 0x86,
	0xd8, 0x53, 0xbe, 0x0d, 0x47, 0x5d, 0x1b, 0x62, 0x6c, 0xef, 0xa4, 0xdf, 0x6e, 0x25, 0x65, 0x69,
	0x5a, 0x0b, 0x0f, 0x83, 0xc7, 0x88, 0x45, 0x13, 0xcf, 0x20, 0x89, 0x64, 0x12, 0x32, 0x1d, 0x46,
	0xfb, 0x2c, 0x1c, 0xff, 0x5f, 0x0b, 0x27, 0x7b, 0x2d, 0xec, 0xb5, 0xcb, 0x74, 0xbb, 0x5d, 0xbe,
	0x82, 0x23, 0xe3, 0x94, 0xda, 0x54, 0xfe, 0x1c, 0x4c, 0x70, 0x0c, 0xd4, 0x70, 0x91, 0xb5, 0xe0,
	0x1d, 0x13, 0x1c, 0xf7, 0x71, 0x7e, 0x67, 0xd3, 0xb5, 0x6c, 0x1a, 0xf7, 0x89, 0xf0, 0xe3, 0x89,
	0xa2, 0xb9, 0x59, 0xdf, 0xdc, 0xf9, 0xdf, 0x03, 0x98, 0x6d, 0xbd, 0x6a, 0xfe, 0xf1, 0xbb, 0x80,
	0x49, 0x17, 0xe3, 0xe8, 0x95, 0x8d, 0x19, 0xde, 0x4d, 0xf9, 0xc1, 0x93, 0x94, 0x9f, 0x42, 0xe2,
	0xf3, 0x51, 0xcb, 0xf6, 0x81, 0x53, 0xac, 0xb9, 0x93, 0xe4, 0x07, 0x98, 0xc8, 0x60, 0x80, 0xa1,
	0x23, 0x0c, 0xfe, 0xc5, 0x76, 0xf0, 0x23, 0x8b, 0xa1, 0xef, 0xa4, 0xf3, 0x5f, 0x20, 0xed, 0x11,
	0xbe, 0xb2, 0xce, 0xd9, 0x58, 0x59, 0xa4, 0x9f, 0x18, 0x7f, 0xb0, 0x6b, 0xfc, 0xeb, 0x7f, 0x06,
	0x90, 0xe1, 0xf7, 0x7d, 0xd8, 0x92, 0xfc, 0x08, 0x23, 0xff, 0xe2, 0x93, 0x9d, 0x4a, 0x7a, 0xbf,
	0x89, 0xcb, 0xf3, 0x7d, 0x94, 0xda, 0x34, 0xf3, 0xcf, 0xc8, 0xef, 0x90, 0xf6, 0x6e, 0x8e, 0x5c,
	0x6f, 0x2b, 0x9f, 0xfe, 0x2a, 0x2e, 0xbf, 0xfc, 0x0f, 0x05, 0x2e, 0xb9, 0x4c, 0xf0, 0xc7, 0xf5,
	0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x94, 0x55, 0x96, 0xce, 0x06, 0x00, 0x00,
}