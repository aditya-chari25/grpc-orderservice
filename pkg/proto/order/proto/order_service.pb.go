// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/order_service.proto

package order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Order item containing product and quantity
type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     int32                  `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UnitPrice     float32                `protobuf:"fixed32,3,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	TotalPrice    float32                `protobuf:"fixed32,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_proto_order_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetUnitPrice() float32 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *OrderItem) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

// Request to create a new order
type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_proto_order_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateOrderRequest) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// Response after creating an order
type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int32                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Total         float32                `protobuf:"fixed32,2,opt,name=total,proto3" json:"total,omitempty"`
	Success       bool                   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_proto_order_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *CreateOrderResponse) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CreateOrderResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Order details
type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Total         float32                `protobuf:"fixed32,4,opt,name=total,proto3" json:"total,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_proto_order_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Request to get order details
type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int32                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_proto_order_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

// Response with order details
type GetOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	mi := &file_proto_order_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

// Request to list orders for a user
type ListOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	mi := &file_proto_order_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListOrdersRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListOrdersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListOrdersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Response with list of orders
type ListOrdersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	mi := &file_proto_order_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *ListOrdersResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

// Request to update order status
type UpdateOrderStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int32                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"` // pending, confirmed, shipped, delivered, cancelled
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusRequest) Reset() {
	*x = UpdateOrderStatusRequest{}
	mi := &file_proto_order_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusRequest) ProtoMessage() {}

func (x *UpdateOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateOrderStatusRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *UpdateOrderStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Response for order status update
type UpdateOrderStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusResponse) Reset() {
	*x = UpdateOrderStatusResponse{}
	mi := &file_proto_order_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusResponse) ProtoMessage() {}

func (x *UpdateOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateOrderStatusResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateOrderStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_order_service_proto protoreflect.FileDescriptor

const file_proto_order_service_proto_rawDesc = "" +
	"\n" +
	"\x19proto/order_service.proto\x12\x0fecommerce.order\x1a\x1fgoogle/protobuf/timestamp.proto\"\x86\x01\n" +
	"\tOrderItem\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\x05R\tproductId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\x12\x1d\n" +
	"\n" +
	"unit_price\x18\x03 \x01(\x02R\tunitPrice\x12\x1f\n" +
	"\vtotal_price\x18\x04 \x01(\x02R\n" +
	"totalPrice\"_\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x05R\x06userId\x120\n" +
	"\x05items\x18\x02 \x03(\v2\x1a.ecommerce.order.OrderItemR\x05items\"\x92\x01\n" +
	"\x13CreateOrderResponse\x12\x19\n" +
	"\border_id\x18\x01 \x01(\x05R\aorderId\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x02R\x05total\x12\x18\n" +
	"\asuccess\x18\x03 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x04 \x01(\tR\amessage\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\"\x86\x02\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x05R\x06userId\x120\n" +
	"\x05items\x18\x03 \x03(\v2\x1a.ecommerce.order.OrderItemR\x05items\x12\x14\n" +
	"\x05total\x18\x04 \x01(\x02R\x05total\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\",\n" +
	"\x0fGetOrderRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\x05R\aorderId\"@\n" +
	"\x10GetOrderResponse\x12,\n" +
	"\x05order\x18\x01 \x01(\v2\x16.ecommerce.order.OrderR\x05order\"Z\n" +
	"\x11ListOrdersRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x05R\x06userId\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x03 \x01(\x05R\x06offset\"e\n" +
	"\x12ListOrdersResponse\x12.\n" +
	"\x06orders\x18\x01 \x03(\v2\x16.ecommerce.order.OrderR\x06orders\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"M\n" +
	"\x18UpdateOrderStatusRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\x05R\aorderId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\"O\n" +
	"\x19UpdateOrderStatusResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\xfc\x02\n" +
	"\fOrderService\x12X\n" +
	"\vCreateOrder\x12#.ecommerce.order.CreateOrderRequest\x1a$.ecommerce.order.CreateOrderResponse\x12O\n" +
	"\bGetOrder\x12 .ecommerce.order.GetOrderRequest\x1a!.ecommerce.order.GetOrderResponse\x12U\n" +
	"\n" +
	"ListOrders\x12\".ecommerce.order.ListOrdersRequest\x1a#.ecommerce.order.ListOrdersResponse\x12j\n" +
	"\x11UpdateOrderStatus\x12).ecommerce.order.UpdateOrderStatusRequest\x1a*.ecommerce.order.UpdateOrderStatusResponseB%Z#github.com/your-project/proto/orderb\x06proto3"

var (
	file_proto_order_service_proto_rawDescOnce sync.Once
	file_proto_order_service_proto_rawDescData []byte
)

func file_proto_order_service_proto_rawDescGZIP() []byte {
	file_proto_order_service_proto_rawDescOnce.Do(func() {
		file_proto_order_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_order_service_proto_rawDesc), len(file_proto_order_service_proto_rawDesc)))
	})
	return file_proto_order_service_proto_rawDescData
}

var file_proto_order_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_order_service_proto_goTypes = []any{
	(*OrderItem)(nil),                 // 0: ecommerce.order.OrderItem
	(*CreateOrderRequest)(nil),        // 1: ecommerce.order.CreateOrderRequest
	(*CreateOrderResponse)(nil),       // 2: ecommerce.order.CreateOrderResponse
	(*Order)(nil),                     // 3: ecommerce.order.Order
	(*GetOrderRequest)(nil),           // 4: ecommerce.order.GetOrderRequest
	(*GetOrderResponse)(nil),          // 5: ecommerce.order.GetOrderResponse
	(*ListOrdersRequest)(nil),         // 6: ecommerce.order.ListOrdersRequest
	(*ListOrdersResponse)(nil),        // 7: ecommerce.order.ListOrdersResponse
	(*UpdateOrderStatusRequest)(nil),  // 8: ecommerce.order.UpdateOrderStatusRequest
	(*UpdateOrderStatusResponse)(nil), // 9: ecommerce.order.UpdateOrderStatusResponse
	(*timestamppb.Timestamp)(nil),     // 10: google.protobuf.Timestamp
}
var file_proto_order_service_proto_depIdxs = []int32{
	0,  // 0: ecommerce.order.CreateOrderRequest.items:type_name -> ecommerce.order.OrderItem
	0,  // 1: ecommerce.order.Order.items:type_name -> ecommerce.order.OrderItem
	10, // 2: ecommerce.order.Order.created_at:type_name -> google.protobuf.Timestamp
	10, // 3: ecommerce.order.Order.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 4: ecommerce.order.GetOrderResponse.order:type_name -> ecommerce.order.Order
	3,  // 5: ecommerce.order.ListOrdersResponse.orders:type_name -> ecommerce.order.Order
	1,  // 6: ecommerce.order.OrderService.CreateOrder:input_type -> ecommerce.order.CreateOrderRequest
	4,  // 7: ecommerce.order.OrderService.GetOrder:input_type -> ecommerce.order.GetOrderRequest
	6,  // 8: ecommerce.order.OrderService.ListOrders:input_type -> ecommerce.order.ListOrdersRequest
	8,  // 9: ecommerce.order.OrderService.UpdateOrderStatus:input_type -> ecommerce.order.UpdateOrderStatusRequest
	2,  // 10: ecommerce.order.OrderService.CreateOrder:output_type -> ecommerce.order.CreateOrderResponse
	5,  // 11: ecommerce.order.OrderService.GetOrder:output_type -> ecommerce.order.GetOrderResponse
	7,  // 12: ecommerce.order.OrderService.ListOrders:output_type -> ecommerce.order.ListOrdersResponse
	9,  // 13: ecommerce.order.OrderService.UpdateOrderStatus:output_type -> ecommerce.order.UpdateOrderStatusResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_order_service_proto_init() }
func file_proto_order_service_proto_init() {
	if File_proto_order_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_order_service_proto_rawDesc), len(file_proto_order_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_service_proto_goTypes,
		DependencyIndexes: file_proto_order_service_proto_depIdxs,
		MessageInfos:      file_proto_order_service_proto_msgTypes,
	}.Build()
	File_proto_order_service_proto = out.File
	file_proto_order_service_proto_goTypes = nil
	file_proto_order_service_proto_depIdxs = nil
}
