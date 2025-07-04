// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: order.proto

package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Positions     []*OrderPosition       `protobuf:"bytes,3,rep,name=positions,proto3" json:"positions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
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
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetPositions() []*OrderPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

type OrderPosition struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CatalogueItemId string                 `protobuf:"bytes,2,opt,name=catalogue_item_id,json=catalogueItemId,proto3" json:"catalogue_item_id,omitempty"`
	Title           string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Quantity        int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *OrderPosition) Reset() {
	*x = OrderPosition{}
	mi := &file_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPosition) ProtoMessage() {}

func (x *OrderPosition) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPosition.ProtoReflect.Descriptor instead.
func (*OrderPosition) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderPosition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderPosition) GetCatalogueItemId() string {
	if x != nil {
		return x.CatalogueItemId
	}
	return ""
}

func (x *OrderPosition) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OrderPosition) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Positions     []*OrderPositionInput  `protobuf:"bytes,2,rep,name=positions,proto3" json:"positions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
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
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetPositions() []*OrderPositionInput {
	if x != nil {
		return x.Positions
	}
	return nil
}

type OrderPositionInput struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	CatalogueItemId string                 `protobuf:"bytes,1,opt,name=catalogue_item_id,json=catalogueItemId,proto3" json:"catalogue_item_id,omitempty"`
	Title           string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Quantity        int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *OrderPositionInput) Reset() {
	*x = OrderPositionInput{}
	mi := &file_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderPositionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPositionInput) ProtoMessage() {}

func (x *OrderPositionInput) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPositionInput.ProtoReflect.Descriptor instead.
func (*OrderPositionInput) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderPositionInput) GetCatalogueItemId() string {
	if x != nil {
		return x.CatalogueItemId
	}
	return ""
}

func (x *OrderPositionInput) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OrderPositionInput) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type OrderList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderList) Reset() {
	*x = OrderList{}
	mi := &file_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderList) ProtoMessage() {}

func (x *OrderList) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderList.ProtoReflect.Descriptor instead.
func (*OrderList) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderList) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

const file_order_proto_rawDesc = "" +
	"\n" +
	"\vorder.proto\x12\x05order\"\a\n" +
	"\x05Empty\"d\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x122\n" +
	"\tpositions\x18\x03 \x03(\v2\x14.order.OrderPositionR\tpositions\"}\n" +
	"\rOrderPosition\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12*\n" +
	"\x11catalogue_item_id\x18\x02 \x01(\tR\x0fcatalogueItemId\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x1a\n" +
	"\bquantity\x18\x04 \x01(\x05R\bquantity\"f\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x127\n" +
	"\tpositions\x18\x02 \x03(\v2\x19.order.OrderPositionInputR\tpositions\"r\n" +
	"\x12OrderPositionInput\x12*\n" +
	"\x11catalogue_item_id\x18\x01 \x01(\tR\x0fcatalogueItemId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x1a\n" +
	"\bquantity\x18\x03 \x01(\x05R\bquantity\"1\n" +
	"\tOrderList\x12$\n" +
	"\x06orders\x18\x01 \x03(\v2\f.order.OrderR\x06orders2t\n" +
	"\fOrderService\x126\n" +
	"\vCreateOrder\x12\x19.order.CreateOrderRequest\x1a\f.order.Order\x12,\n" +
	"\n" +
	"ListOrders\x12\f.order.Empty\x1a\x10.order.OrderListB\fZ\n" +
	"./;orderpbb\x06proto3"

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData []byte
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)))
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_order_proto_goTypes = []any{
	(*Empty)(nil),              // 0: order.Empty
	(*Order)(nil),              // 1: order.Order
	(*OrderPosition)(nil),      // 2: order.OrderPosition
	(*CreateOrderRequest)(nil), // 3: order.CreateOrderRequest
	(*OrderPositionInput)(nil), // 4: order.OrderPositionInput
	(*OrderList)(nil),          // 5: order.OrderList
}
var file_order_proto_depIdxs = []int32{
	2, // 0: order.Order.positions:type_name -> order.OrderPosition
	4, // 1: order.CreateOrderRequest.positions:type_name -> order.OrderPositionInput
	1, // 2: order.OrderList.orders:type_name -> order.Order
	3, // 3: order.OrderService.CreateOrder:input_type -> order.CreateOrderRequest
	0, // 4: order.OrderService.ListOrders:input_type -> order.Empty
	1, // 5: order.OrderService.CreateOrder:output_type -> order.Order
	5, // 6: order.OrderService.ListOrders:output_type -> order.OrderList
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
