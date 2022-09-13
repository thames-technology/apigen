// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: acme/shelf/v1/request_response.proto

package shelfv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The standard List request definition.
type ListShelvesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only retrieve shelves after this time.
	AfterTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=after_time,json=afterTime,proto3" json:"after_time,omitempty"`
	// Only retrieve shelves before this time.
	BeforeTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=before_time,json=beforeTime,proto3" json:"before_time,omitempty"`
	// The start index for pagination.
	Start uint64 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	// The maximum number of shelves to return.
	MaxSize uint64 `protobuf:"varint,4,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	// The unique id of the parent example for which to list the shelves.
	ExampleId string `protobuf:"bytes,5,opt,name=example_id,json=exampleId,proto3" json:"example_id,omitempty"`
}

func (x *ListShelvesRequest) Reset() {
	*x = ListShelvesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShelvesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShelvesRequest) ProtoMessage() {}

func (x *ListShelvesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShelvesRequest.ProtoReflect.Descriptor instead.
func (*ListShelvesRequest) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{0}
}

func (x *ListShelvesRequest) GetAfterTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AfterTime
	}
	return nil
}

func (x *ListShelvesRequest) GetBeforeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.BeforeTime
	}
	return nil
}

func (x *ListShelvesRequest) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ListShelvesRequest) GetMaxSize() uint64 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *ListShelvesRequest) GetExampleId() string {
	if x != nil {
		return x.ExampleId
	}
	return ""
}

// The standard List response definition.
type ListShelvesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The retrieved list of shelves.
	Shelves []*Shelf `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
	// True if more shelves are available.
	Next bool `protobuf:"varint,2,opt,name=next,proto3" json:"next,omitempty"`
}

func (x *ListShelvesResponse) Reset() {
	*x = ListShelvesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShelvesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShelvesResponse) ProtoMessage() {}

func (x *ListShelvesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShelvesResponse.ProtoReflect.Descriptor instead.
func (*ListShelvesResponse) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{1}
}

func (x *ListShelvesResponse) GetShelves() []*Shelf {
	if x != nil {
		return x.Shelves
	}
	return nil
}

func (x *ListShelvesResponse) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

// The standard BatchGet request definition.
type BatchGetShelvesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ids of the requested shelves.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *BatchGetShelvesRequest) Reset() {
	*x = BatchGetShelvesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetShelvesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetShelvesRequest) ProtoMessage() {}

func (x *BatchGetShelvesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetShelvesRequest.ProtoReflect.Descriptor instead.
func (*BatchGetShelvesRequest) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{2}
}

func (x *BatchGetShelvesRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// The standard BatchGet response definition.
type BatchGetShelvesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The retrieved shelves.
	Shelves []*Shelf `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
}

func (x *BatchGetShelvesResponse) Reset() {
	*x = BatchGetShelvesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetShelvesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetShelvesResponse) ProtoMessage() {}

func (x *BatchGetShelvesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetShelvesResponse.ProtoReflect.Descriptor instead.
func (*BatchGetShelvesResponse) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{3}
}

func (x *BatchGetShelvesResponse) GetShelves() []*Shelf {
	if x != nil {
		return x.Shelves
	}
	return nil
}

// The standard Get request definition.
type GetShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the requested shelf.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetShelfRequest) Reset() {
	*x = GetShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelfRequest) ProtoMessage() {}

func (x *GetShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelfRequest.ProtoReflect.Descriptor instead.
func (*GetShelfRequest) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{4}
}

func (x *GetShelfRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The standard Get response definition.
type GetShelfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The retrieved shelf.
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *GetShelfResponse) Reset() {
	*x = GetShelfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelfResponse) ProtoMessage() {}

func (x *GetShelfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelfResponse.ProtoReflect.Descriptor instead.
func (*GetShelfResponse) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{5}
}

func (x *GetShelfResponse) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

// The standard Create request definition.
type CreateShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The shelf to create.
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The unique id of the parent example where the shelf is to be created.
	ExampleId string `protobuf:"bytes,4,opt,name=example_id,json=exampleId,proto3" json:"example_id,omitempty"`
}

func (x *CreateShelfRequest) Reset() {
	*x = CreateShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelfRequest) ProtoMessage() {}

func (x *CreateShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShelfRequest.ProtoReflect.Descriptor instead.
func (*CreateShelfRequest) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{6}
}

func (x *CreateShelfRequest) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

func (x *CreateShelfRequest) GetExampleId() string {
	if x != nil {
		return x.ExampleId
	}
	return ""
}

// The standard Create response definition.
type CreateShelfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The created shelf.
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *CreateShelfResponse) Reset() {
	*x = CreateShelfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShelfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelfResponse) ProtoMessage() {}

func (x *CreateShelfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShelfResponse.ProtoReflect.Descriptor instead.
func (*CreateShelfResponse) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{7}
}

func (x *CreateShelfResponse) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

// The standard Update request definition.
type UpdateShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the shelf to be updated.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The shelf which replaces the shelf on the server.
	Shelf *Shelf `protobuf:"bytes,2,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The update mask applied to the shelf. See https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateShelfRequest) Reset() {
	*x = UpdateShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShelfRequest) ProtoMessage() {}

func (x *UpdateShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShelfRequest.ProtoReflect.Descriptor instead.
func (*UpdateShelfRequest) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateShelfRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateShelfRequest) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

func (x *UpdateShelfRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// The standard Update response definition.
type UpdateShelfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The updated shelf.
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *UpdateShelfResponse) Reset() {
	*x = UpdateShelfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateShelfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShelfResponse) ProtoMessage() {}

func (x *UpdateShelfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShelfResponse.ProtoReflect.Descriptor instead.
func (*UpdateShelfResponse) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateShelfResponse) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

// The standard Delete request definition.
type DeleteShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the shelf to be deleted.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteShelfRequest) Reset() {
	*x = DeleteShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelfRequest) ProtoMessage() {}

func (x *DeleteShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShelfRequest.ProtoReflect.Descriptor instead.
func (*DeleteShelfRequest) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteShelfRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The standard Delete response definition.
type DeleteShelfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The deleted shelf.
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *DeleteShelfResponse) Reset() {
	*x = DeleteShelfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_shelf_v1_request_response_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShelfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelfResponse) ProtoMessage() {}

func (x *DeleteShelfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acme_shelf_v1_request_response_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShelfResponse.ProtoReflect.Descriptor instead.
func (*DeleteShelfResponse) Descriptor() ([]byte, []int) {
	return file_acme_shelf_v1_request_response_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteShelfResponse) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

var File_acme_shelf_v1_request_response_proto protoreflect.FileDescriptor

var file_acme_shelf_v1_request_response_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x61, 0x78, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65,
	0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07,
	0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x07, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74,
	0x22, 0x2a, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c,
	0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x49, 0x0a, 0x17,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x68, 0x65, 0x6c, 0x76,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e,
	0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x07,
	0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x5f, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x8d,
	0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x41,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x99, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31,
	0x42, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x18, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x6d,
	0x65, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x58, 0xaa, 0x02, 0x0d, 0x41, 0x63, 0x6d, 0x65, 0x2e,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x41, 0x63, 0x6d, 0x65, 0x5c,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x41, 0x63, 0x6d, 0x65, 0x5c,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x41, 0x63, 0x6d, 0x65, 0x3a, 0x3a, 0x53, 0x68, 0x65,
	0x6c, 0x66, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acme_shelf_v1_request_response_proto_rawDescOnce sync.Once
	file_acme_shelf_v1_request_response_proto_rawDescData = file_acme_shelf_v1_request_response_proto_rawDesc
)

func file_acme_shelf_v1_request_response_proto_rawDescGZIP() []byte {
	file_acme_shelf_v1_request_response_proto_rawDescOnce.Do(func() {
		file_acme_shelf_v1_request_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_acme_shelf_v1_request_response_proto_rawDescData)
	})
	return file_acme_shelf_v1_request_response_proto_rawDescData
}

var file_acme_shelf_v1_request_response_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_acme_shelf_v1_request_response_proto_goTypes = []interface{}{
	(*ListShelvesRequest)(nil),      // 0: acme.shelf.v1.ListShelvesRequest
	(*ListShelvesResponse)(nil),     // 1: acme.shelf.v1.ListShelvesResponse
	(*BatchGetShelvesRequest)(nil),  // 2: acme.shelf.v1.BatchGetShelvesRequest
	(*BatchGetShelvesResponse)(nil), // 3: acme.shelf.v1.BatchGetShelvesResponse
	(*GetShelfRequest)(nil),         // 4: acme.shelf.v1.GetShelfRequest
	(*GetShelfResponse)(nil),        // 5: acme.shelf.v1.GetShelfResponse
	(*CreateShelfRequest)(nil),      // 6: acme.shelf.v1.CreateShelfRequest
	(*CreateShelfResponse)(nil),     // 7: acme.shelf.v1.CreateShelfResponse
	(*UpdateShelfRequest)(nil),      // 8: acme.shelf.v1.UpdateShelfRequest
	(*UpdateShelfResponse)(nil),     // 9: acme.shelf.v1.UpdateShelfResponse
	(*DeleteShelfRequest)(nil),      // 10: acme.shelf.v1.DeleteShelfRequest
	(*DeleteShelfResponse)(nil),     // 11: acme.shelf.v1.DeleteShelfResponse
	(*timestamppb.Timestamp)(nil),   // 12: google.protobuf.Timestamp
	(*Shelf)(nil),                   // 13: acme.shelf.v1.Shelf
	(*fieldmaskpb.FieldMask)(nil),   // 14: google.protobuf.FieldMask
}
var file_acme_shelf_v1_request_response_proto_depIdxs = []int32{
	12, // 0: acme.shelf.v1.ListShelvesRequest.after_time:type_name -> google.protobuf.Timestamp
	12, // 1: acme.shelf.v1.ListShelvesRequest.before_time:type_name -> google.protobuf.Timestamp
	13, // 2: acme.shelf.v1.ListShelvesResponse.shelves:type_name -> acme.shelf.v1.Shelf
	13, // 3: acme.shelf.v1.BatchGetShelvesResponse.shelves:type_name -> acme.shelf.v1.Shelf
	13, // 4: acme.shelf.v1.GetShelfResponse.shelf:type_name -> acme.shelf.v1.Shelf
	13, // 5: acme.shelf.v1.CreateShelfRequest.shelf:type_name -> acme.shelf.v1.Shelf
	13, // 6: acme.shelf.v1.CreateShelfResponse.shelf:type_name -> acme.shelf.v1.Shelf
	13, // 7: acme.shelf.v1.UpdateShelfRequest.shelf:type_name -> acme.shelf.v1.Shelf
	14, // 8: acme.shelf.v1.UpdateShelfRequest.update_mask:type_name -> google.protobuf.FieldMask
	13, // 9: acme.shelf.v1.UpdateShelfResponse.shelf:type_name -> acme.shelf.v1.Shelf
	13, // 10: acme.shelf.v1.DeleteShelfResponse.shelf:type_name -> acme.shelf.v1.Shelf
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_acme_shelf_v1_request_response_proto_init() }
func file_acme_shelf_v1_request_response_proto_init() {
	if File_acme_shelf_v1_request_response_proto != nil {
		return
	}
	file_acme_shelf_v1_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_acme_shelf_v1_request_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShelvesRequest); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShelvesResponse); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetShelvesRequest); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetShelvesResponse); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelfRequest); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelfResponse); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShelfRequest); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShelfResponse); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateShelfRequest); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateShelfResponse); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShelfRequest); i {
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
		file_acme_shelf_v1_request_response_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShelfResponse); i {
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
			RawDescriptor: file_acme_shelf_v1_request_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acme_shelf_v1_request_response_proto_goTypes,
		DependencyIndexes: file_acme_shelf_v1_request_response_proto_depIdxs,
		MessageInfos:      file_acme_shelf_v1_request_response_proto_msgTypes,
	}.Build()
	File_acme_shelf_v1_request_response_proto = out.File
	file_acme_shelf_v1_request_response_proto_rawDesc = nil
	file_acme_shelf_v1_request_response_proto_goTypes = nil
	file_acme_shelf_v1_request_response_proto_depIdxs = nil
}
