// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/order.proto

package order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type OrderStatus int32

const (
	OrderStatus_ORDER_STATUS_UNSPECIFIED OrderStatus = 0
	OrderStatus_CREATED                  OrderStatus = 1
	OrderStatus_COMPLETED                OrderStatus = 2
	OrderStatus_CANCELLED                OrderStatus = 3
	OrderStatus_INVENTORY_RESERVED       OrderStatus = 4
	OrderStatus_PAYMENT_PENDING          OrderStatus = 5
	OrderStatus_PAYMENT_FAILED           OrderStatus = 6
	OrderStatus_PAYMENT_SUCCESS          OrderStatus = 7
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "ORDER_STATUS_UNSPECIFIED",
		1: "CREATED",
		2: "COMPLETED",
		3: "CANCELLED",
		4: "INVENTORY_RESERVED",
		5: "PAYMENT_PENDING",
		6: "PAYMENT_FAILED",
		7: "PAYMENT_SUCCESS",
	}
	OrderStatus_value = map[string]int32{
		"ORDER_STATUS_UNSPECIFIED": 0,
		"CREATED":                  1,
		"COMPLETED":                2,
		"CANCELLED":                3,
		"INVENTORY_RESERVED":       4,
		"PAYMENT_PENDING":          5,
		"PAYMENT_FAILED":           6,
		"PAYMENT_SUCCESS":          7,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_order_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_proto_order_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

type OrderWorkflowStatus int32

const (
	OrderWorkflowStatus_ORDER_WORKFLOW_STATUS_UNSPECIFIED OrderWorkflowStatus = 0
	OrderWorkflowStatus_ORDER_WORKFLOW_STATUS_PROCESSING  OrderWorkflowStatus = 1
	OrderWorkflowStatus_ORDER_WORKFLOW_STATUS_COMPLETED   OrderWorkflowStatus = 2
	OrderWorkflowStatus_ORDER_WORKFLOW_STATUS_CANCELLED   OrderWorkflowStatus = 3
)

// Enum value maps for OrderWorkflowStatus.
var (
	OrderWorkflowStatus_name = map[int32]string{
		0: "ORDER_WORKFLOW_STATUS_UNSPECIFIED",
		1: "ORDER_WORKFLOW_STATUS_PROCESSING",
		2: "ORDER_WORKFLOW_STATUS_COMPLETED",
		3: "ORDER_WORKFLOW_STATUS_CANCELLED",
	}
	OrderWorkflowStatus_value = map[string]int32{
		"ORDER_WORKFLOW_STATUS_UNSPECIFIED": 0,
		"ORDER_WORKFLOW_STATUS_PROCESSING":  1,
		"ORDER_WORKFLOW_STATUS_COMPLETED":   2,
		"ORDER_WORKFLOW_STATUS_CANCELLED":   3,
	}
)

func (x OrderWorkflowStatus) Enum() *OrderWorkflowStatus {
	p := new(OrderWorkflowStatus)
	*p = x
	return p
}

func (x OrderWorkflowStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderWorkflowStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_order_proto_enumTypes[1].Descriptor()
}

func (OrderWorkflowStatus) Type() protoreflect.EnumType {
	return &file_proto_order_proto_enumTypes[1]
}

func (x OrderWorkflowStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderWorkflowStatus.Descriptor instead.
func (OrderWorkflowStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

type OrderData struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code            string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Items           []*OrderItem           `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	UserId          string                 `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status          OrderStatus            `protobuf:"varint,5,opt,name=status,proto3,enum=order.OrderStatus" json:"status,omitempty"`
	TotalAmount     float64                `protobuf:"fixed64,6,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	PaymentMethod   string                 `protobuf:"bytes,7,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	Provider        string                 `protobuf:"bytes,8,opt,name=provider,proto3" json:"provider,omitempty"`
	ProviderDetails string                 `protobuf:"bytes,9,opt,name=provider_details,json=providerDetails,proto3" json:"provider_details,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *OrderData) Reset() {
	*x = OrderData{}
	mi := &file_proto_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderData) ProtoMessage() {}

func (x *OrderData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderData.ProtoReflect.Descriptor instead.
func (*OrderData) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderData) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OrderData) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderData) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderData) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

func (x *OrderData) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *OrderData) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *OrderData) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OrderData) GetProviderDetails() string {
	if x != nil {
		return x.ProviderDetails
	}
	return ""
}

type OrderItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItemRequest) Reset() {
	*x = OrderItemRequest{}
	mi := &file_proto_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemRequest) ProtoMessage() {}

func (x *OrderItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemRequest.ProtoReflect.Descriptor instead.
func (*OrderItemRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductName   string                 `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductPrice  float64                `protobuf:"fixed64,3,opt,name=product_price,json=productPrice,proto3" json:"product_price,omitempty"`
	Quantity      int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TotalAmount   float64                `protobuf:"fixed64,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_proto_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[2]
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
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItem) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *OrderItem) GetProductPrice() float64 {
	if x != nil {
		return x.ProductPrice
	}
	return 0
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

// Create
type CreateRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Items           []*OrderItemRequest    `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	UserId          string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PaymentMethod   string                 `protobuf:"bytes,3,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	Provider        string                 `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	ProviderDetails string                 `protobuf:"bytes,5,opt,name=provider_details,json=providerDetails,proto3" json:"provider_details,omitempty"`
	Metadata        map[string]string      `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_proto_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetItems() []*OrderItemRequest {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *CreateRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CreateRequest) GetProviderDetails() string {
	if x != nil {
		return x.ProviderDetails
	}
	return ""
}

func (x *CreateRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderCode     string                 `protobuf:"bytes,1,opt,name=order_code,json=orderCode,proto3" json:"order_code,omitempty"`
	WorkflowId    string                 `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	PaymentUrl    string                 `protobuf:"bytes,3,opt,name=payment_url,json=paymentUrl,proto3" json:"payment_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_proto_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{4}
}

func (x *CreateResponse) GetOrderCode() string {
	if x != nil {
		return x.OrderCode
	}
	return ""
}

func (x *CreateResponse) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *CreateResponse) GetPaymentUrl() string {
	if x != nil {
		return x.PaymentUrl
	}
	return ""
}

// FindOne
type FindOneRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Request:
	//
	//	*FindOneRequest_Id
	//	*FindOneRequest_Code
	Request       isFindOneRequest_Request `protobuf_oneof:"request"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindOneRequest) Reset() {
	*x = FindOneRequest{}
	mi := &file_proto_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRequest) ProtoMessage() {}

func (x *FindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRequest.ProtoReflect.Descriptor instead.
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{5}
}

func (x *FindOneRequest) GetRequest() isFindOneRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *FindOneRequest) GetId() string {
	if x != nil {
		if x, ok := x.Request.(*FindOneRequest_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *FindOneRequest) GetCode() string {
	if x != nil {
		if x, ok := x.Request.(*FindOneRequest_Code); ok {
			return x.Code
		}
	}
	return ""
}

type isFindOneRequest_Request interface {
	isFindOneRequest_Request()
}

type FindOneRequest_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type FindOneRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*FindOneRequest_Id) isFindOneRequest_Request() {}

func (*FindOneRequest_Code) isFindOneRequest_Request() {}

type FindOneResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *OrderData             `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindOneResponse) Reset() {
	*x = FindOneResponse{}
	mi := &file_proto_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneResponse) ProtoMessage() {}

func (x *FindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneResponse.ProtoReflect.Descriptor instead.
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{6}
}

func (x *FindOneResponse) GetOrder() *OrderData {
	if x != nil {
		return x.Order
	}
	return nil
}

// FindMany
type FindManyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status        OrderStatus            `protobuf:"varint,2,opt,name=status,proto3,enum=order.OrderStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindManyRequest) Reset() {
	*x = FindManyRequest{}
	mi := &file_proto_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindManyRequest) ProtoMessage() {}

func (x *FindManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindManyRequest.ProtoReflect.Descriptor instead.
func (*FindManyRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{7}
}

func (x *FindManyRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FindManyRequest) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

type FindManyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*OrderData           `protobuf:"bytes,3,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindManyResponse) Reset() {
	*x = FindManyResponse{}
	mi := &file_proto_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindManyResponse) ProtoMessage() {}

func (x *FindManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindManyResponse.ProtoReflect.Descriptor instead.
func (*FindManyResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{8}
}

func (x *FindManyResponse) GetOrders() []*OrderData {
	if x != nil {
		return x.Orders
	}
	return nil
}

// UpdateStatus
type UpdateStatusRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Request:
	//
	//	*UpdateStatusRequest_Id
	//	*UpdateStatusRequest_Code
	Request       isUpdateStatusRequest_Request `protobuf_oneof:"request"`
	Status        OrderStatus                   `protobuf:"varint,3,opt,name=status,proto3,enum=order.OrderStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStatusRequest) Reset() {
	*x = UpdateStatusRequest{}
	mi := &file_proto_order_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusRequest) ProtoMessage() {}

func (x *UpdateStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateStatusRequest) GetRequest() isUpdateStatusRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *UpdateStatusRequest) GetId() string {
	if x != nil {
		if x, ok := x.Request.(*UpdateStatusRequest_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *UpdateStatusRequest) GetCode() string {
	if x != nil {
		if x, ok := x.Request.(*UpdateStatusRequest_Code); ok {
			return x.Code
		}
	}
	return ""
}

func (x *UpdateStatusRequest) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

type isUpdateStatusRequest_Request interface {
	isUpdateStatusRequest_Request()
}

type UpdateStatusRequest_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type UpdateStatusRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*UpdateStatusRequest_Id) isUpdateStatusRequest_Request() {}

func (*UpdateStatusRequest_Code) isUpdateStatusRequest_Request() {}

type OrderWorkflowParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderCode     string                 `protobuf:"bytes,1,opt,name=order_code,json=orderCode,proto3" json:"order_code,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TotalAmount   float64                `protobuf:"fixed64,3,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderWorkflowParams) Reset() {
	*x = OrderWorkflowParams{}
	mi := &file_proto_order_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderWorkflowParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderWorkflowParams) ProtoMessage() {}

func (x *OrderWorkflowParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderWorkflowParams.ProtoReflect.Descriptor instead.
func (*OrderWorkflowParams) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{10}
}

func (x *OrderWorkflowParams) GetOrderCode() string {
	if x != nil {
		return x.OrderCode
	}
	return ""
}

func (x *OrderWorkflowParams) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderWorkflowParams) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

type OrderWorkflowResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderCode     string                 `protobuf:"bytes,1,opt,name=order_code,json=orderCode,proto3" json:"order_code,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMessage  string                 `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderWorkflowResult) Reset() {
	*x = OrderWorkflowResult{}
	mi := &file_proto_order_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderWorkflowResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderWorkflowResult) ProtoMessage() {}

func (x *OrderWorkflowResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderWorkflowResult.ProtoReflect.Descriptor instead.
func (*OrderWorkflowResult) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{11}
}

func (x *OrderWorkflowResult) GetOrderCode() string {
	if x != nil {
		return x.OrderCode
	}
	return ""
}

func (x *OrderWorkflowResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderWorkflowResult) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_proto_order_proto protoreflect.FileDescriptor

const file_proto_order_proto_rawDesc = "" +
	"\n" +
	"\x11proto/order.proto\x12\x05order\x1a\x1bgoogle/protobuf/empty.proto\"\xad\x02\n" +
	"\tOrderData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04code\x18\x02 \x01(\tR\x04code\x12&\n" +
	"\x05items\x18\x03 \x03(\v2\x10.order.OrderItemR\x05items\x12\x17\n" +
	"\auser_id\x18\x04 \x01(\tR\x06userId\x12*\n" +
	"\x06status\x18\x05 \x01(\x0e2\x12.order.OrderStatusR\x06status\x12!\n" +
	"\ftotal_amount\x18\x06 \x01(\x01R\vtotalAmount\x12%\n" +
	"\x0epayment_method\x18\a \x01(\tR\rpaymentMethod\x12\x1a\n" +
	"\bprovider\x18\b \x01(\tR\bprovider\x12)\n" +
	"\x10provider_details\x18\t \x01(\tR\x0fproviderDetails\"M\n" +
	"\x10OrderItemRequest\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\"\xb1\x01\n" +
	"\tOrderItem\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12!\n" +
	"\fproduct_name\x18\x02 \x01(\tR\vproductName\x12#\n" +
	"\rproduct_price\x18\x03 \x01(\x01R\fproductPrice\x12\x1a\n" +
	"\bquantity\x18\x04 \x01(\x05R\bquantity\x12!\n" +
	"\ftotal_amount\x18\x05 \x01(\x01R\vtotalAmount\"\xc2\x02\n" +
	"\rCreateRequest\x12-\n" +
	"\x05items\x18\x01 \x03(\v2\x17.order.OrderItemRequestR\x05items\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12%\n" +
	"\x0epayment_method\x18\x03 \x01(\tR\rpaymentMethod\x12\x1a\n" +
	"\bprovider\x18\x04 \x01(\tR\bprovider\x12)\n" +
	"\x10provider_details\x18\x05 \x01(\tR\x0fproviderDetails\x12>\n" +
	"\bmetadata\x18\x06 \x03(\v2\".order.CreateRequest.MetadataEntryR\bmetadata\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"q\n" +
	"\x0eCreateResponse\x12\x1d\n" +
	"\n" +
	"order_code\x18\x01 \x01(\tR\torderCode\x12\x1f\n" +
	"\vworkflow_id\x18\x02 \x01(\tR\n" +
	"workflowId\x12\x1f\n" +
	"\vpayment_url\x18\x03 \x01(\tR\n" +
	"paymentUrl\"C\n" +
	"\x0eFindOneRequest\x12\x10\n" +
	"\x02id\x18\x01 \x01(\tH\x00R\x02id\x12\x14\n" +
	"\x04code\x18\x02 \x01(\tH\x00R\x04codeB\t\n" +
	"\arequest\"9\n" +
	"\x0fFindOneResponse\x12&\n" +
	"\x05order\x18\x03 \x01(\v2\x10.order.OrderDataR\x05order\"V\n" +
	"\x0fFindManyRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12*\n" +
	"\x06status\x18\x02 \x01(\x0e2\x12.order.OrderStatusR\x06status\"<\n" +
	"\x10FindManyResponse\x12(\n" +
	"\x06orders\x18\x03 \x03(\v2\x10.order.OrderDataR\x06orders\"t\n" +
	"\x13UpdateStatusRequest\x12\x10\n" +
	"\x02id\x18\x01 \x01(\tH\x00R\x02id\x12\x14\n" +
	"\x04code\x18\x02 \x01(\tH\x00R\x04code\x12*\n" +
	"\x06status\x18\x03 \x01(\x0e2\x12.order.OrderStatusR\x06statusB\t\n" +
	"\arequest\"p\n" +
	"\x13OrderWorkflowParams\x12\x1d\n" +
	"\n" +
	"order_code\x18\x01 \x01(\tR\torderCode\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12!\n" +
	"\ftotal_amount\x18\x03 \x01(\x01R\vtotalAmount\"q\n" +
	"\x13OrderWorkflowResult\x12\x1d\n" +
	"\n" +
	"order_code\x18\x01 \x01(\tR\torderCode\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x12#\n" +
	"\rerror_message\x18\x03 \x01(\tR\ferrorMessage*\xac\x01\n" +
	"\vOrderStatus\x12\x1c\n" +
	"\x18ORDER_STATUS_UNSPECIFIED\x10\x00\x12\v\n" +
	"\aCREATED\x10\x01\x12\r\n" +
	"\tCOMPLETED\x10\x02\x12\r\n" +
	"\tCANCELLED\x10\x03\x12\x16\n" +
	"\x12INVENTORY_RESERVED\x10\x04\x12\x13\n" +
	"\x0fPAYMENT_PENDING\x10\x05\x12\x12\n" +
	"\x0ePAYMENT_FAILED\x10\x06\x12\x13\n" +
	"\x0fPAYMENT_SUCCESS\x10\a*\xac\x01\n" +
	"\x13OrderWorkflowStatus\x12%\n" +
	"!ORDER_WORKFLOW_STATUS_UNSPECIFIED\x10\x00\x12$\n" +
	" ORDER_WORKFLOW_STATUS_PROCESSING\x10\x01\x12#\n" +
	"\x1fORDER_WORKFLOW_STATUS_COMPLETED\x10\x02\x12#\n" +
	"\x1fORDER_WORKFLOW_STATUS_CANCELLED\x10\x032\x88\x02\n" +
	"\fOrderService\x127\n" +
	"\x06Create\x12\x14.order.CreateRequest\x1a\x15.order.CreateResponse\"\x00\x12:\n" +
	"\aFindOne\x12\x15.order.FindOneRequest\x1a\x16.order.FindOneResponse\"\x00\x12=\n" +
	"\bFindMany\x12\x16.order.FindManyRequest\x1a\x17.order.FindManyResponse\"\x00\x12D\n" +
	"\fUpdateStatus\x12\x1a.order.UpdateStatusRequest\x1a\x16.google.protobuf.Empty\"\x00BIZGgithub.com/vogiaan1904/e-commerce-grpc-nest-proto/protogen/golang/orderb\x06proto3"

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData []byte
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_order_proto_rawDesc), len(file_proto_order_proto_rawDesc)))
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_order_proto_goTypes = []any{
	(OrderStatus)(0),            // 0: order.OrderStatus
	(OrderWorkflowStatus)(0),    // 1: order.OrderWorkflowStatus
	(*OrderData)(nil),           // 2: order.OrderData
	(*OrderItemRequest)(nil),    // 3: order.OrderItemRequest
	(*OrderItem)(nil),           // 4: order.OrderItem
	(*CreateRequest)(nil),       // 5: order.CreateRequest
	(*CreateResponse)(nil),      // 6: order.CreateResponse
	(*FindOneRequest)(nil),      // 7: order.FindOneRequest
	(*FindOneResponse)(nil),     // 8: order.FindOneResponse
	(*FindManyRequest)(nil),     // 9: order.FindManyRequest
	(*FindManyResponse)(nil),    // 10: order.FindManyResponse
	(*UpdateStatusRequest)(nil), // 11: order.UpdateStatusRequest
	(*OrderWorkflowParams)(nil), // 12: order.OrderWorkflowParams
	(*OrderWorkflowResult)(nil), // 13: order.OrderWorkflowResult
	nil,                         // 14: order.CreateRequest.MetadataEntry
	(*emptypb.Empty)(nil),       // 15: google.protobuf.Empty
}
var file_proto_order_proto_depIdxs = []int32{
	4,  // 0: order.OrderData.items:type_name -> order.OrderItem
	0,  // 1: order.OrderData.status:type_name -> order.OrderStatus
	3,  // 2: order.CreateRequest.items:type_name -> order.OrderItemRequest
	14, // 3: order.CreateRequest.metadata:type_name -> order.CreateRequest.MetadataEntry
	2,  // 4: order.FindOneResponse.order:type_name -> order.OrderData
	0,  // 5: order.FindManyRequest.status:type_name -> order.OrderStatus
	2,  // 6: order.FindManyResponse.orders:type_name -> order.OrderData
	0,  // 7: order.UpdateStatusRequest.status:type_name -> order.OrderStatus
	5,  // 8: order.OrderService.Create:input_type -> order.CreateRequest
	7,  // 9: order.OrderService.FindOne:input_type -> order.FindOneRequest
	9,  // 10: order.OrderService.FindMany:input_type -> order.FindManyRequest
	11, // 11: order.OrderService.UpdateStatus:input_type -> order.UpdateStatusRequest
	6,  // 12: order.OrderService.Create:output_type -> order.CreateResponse
	8,  // 13: order.OrderService.FindOne:output_type -> order.FindOneResponse
	10, // 14: order.OrderService.FindMany:output_type -> order.FindManyResponse
	15, // 15: order.OrderService.UpdateStatus:output_type -> google.protobuf.Empty
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_order_proto_init() }
func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	file_proto_order_proto_msgTypes[5].OneofWrappers = []any{
		(*FindOneRequest_Id)(nil),
		(*FindOneRequest_Code)(nil),
	}
	file_proto_order_proto_msgTypes[9].OneofWrappers = []any{
		(*UpdateStatusRequest_Id)(nil),
		(*UpdateStatusRequest_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_order_proto_rawDesc), len(file_proto_order_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		EnumInfos:         file_proto_order_proto_enumTypes,
		MessageInfos:      file_proto_order_proto_msgTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}
