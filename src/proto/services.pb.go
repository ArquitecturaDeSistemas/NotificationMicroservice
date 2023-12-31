// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/services.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Titulo       string `protobuf:"bytes,1,opt,name=titulo,proto3" json:"titulo,omitempty"`
	UsuarioId    string `protobuf:"bytes,2,opt,name=usuario_id,json=usuarioId,proto3" json:"usuario_id,omitempty"`
	TareaId      string `protobuf:"bytes,3,opt,name=tarea_id,json=tareaId,proto3" json:"tarea_id,omitempty"`
	FechaTermino string `protobuf:"bytes,4,opt,name=fecha_termino,json=fechaTermino,proto3" json:"fecha_termino,omitempty"`
}

func (x *CreateNotificationRequest) Reset() {
	*x = CreateNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotificationRequest) ProtoMessage() {}

func (x *CreateNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotificationRequest.ProtoReflect.Descriptor instead.
func (*CreateNotificationRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNotificationRequest) GetTitulo() string {
	if x != nil {
		return x.Titulo
	}
	return ""
}

func (x *CreateNotificationRequest) GetUsuarioId() string {
	if x != nil {
		return x.UsuarioId
	}
	return ""
}

func (x *CreateNotificationRequest) GetTareaId() string {
	if x != nil {
		return x.TareaId
	}
	return ""
}

func (x *CreateNotificationRequest) GetFechaTermino() string {
	if x != nil {
		return x.FechaTermino
	}
	return ""
}

type CreateNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateNotificationResponse) Reset() {
	*x = CreateNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotificationResponse) ProtoMessage() {}

func (x *CreateNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotificationResponse.ProtoReflect.Descriptor instead.
func (*CreateNotificationResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNotificationResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UpdateNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Titulo       string `protobuf:"bytes,1,opt,name=titulo,proto3" json:"titulo,omitempty"`
	UsuarioId    string `protobuf:"bytes,2,opt,name=usuario_id,json=usuarioId,proto3" json:"usuario_id,omitempty"`
	TareaId      string `protobuf:"bytes,3,opt,name=tarea_id,json=tareaId,proto3" json:"tarea_id,omitempty"`
	FechaTermino string `protobuf:"bytes,4,opt,name=fecha_termino,json=fechaTermino,proto3" json:"fecha_termino,omitempty"`
}

func (x *UpdateNotificationRequest) Reset() {
	*x = UpdateNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNotificationRequest) ProtoMessage() {}

func (x *UpdateNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNotificationRequest.ProtoReflect.Descriptor instead.
func (*UpdateNotificationRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNotificationRequest) GetTitulo() string {
	if x != nil {
		return x.Titulo
	}
	return ""
}

func (x *UpdateNotificationRequest) GetUsuarioId() string {
	if x != nil {
		return x.UsuarioId
	}
	return ""
}

func (x *UpdateNotificationRequest) GetTareaId() string {
	if x != nil {
		return x.TareaId
	}
	return ""
}

func (x *UpdateNotificationRequest) GetFechaTermino() string {
	if x != nil {
		return x.FechaTermino
	}
	return ""
}

type UpdateNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateNotificationResponse) Reset() {
	*x = UpdateNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNotificationResponse) ProtoMessage() {}

func (x *UpdateNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNotificationResponse.ProtoReflect.Descriptor instead.
func (*UpdateNotificationResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNotificationResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type DeleteNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsuarioId string `protobuf:"bytes,1,opt,name=usuario_id,json=usuarioId,proto3" json:"usuario_id,omitempty"`
	TareaId   string `protobuf:"bytes,2,opt,name=tarea_id,json=tareaId,proto3" json:"tarea_id,omitempty"`
}

func (x *DeleteNotificationRequest) Reset() {
	*x = DeleteNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNotificationRequest) ProtoMessage() {}

func (x *DeleteNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNotificationRequest.ProtoReflect.Descriptor instead.
func (*DeleteNotificationRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteNotificationRequest) GetUsuarioId() string {
	if x != nil {
		return x.UsuarioId
	}
	return ""
}

func (x *DeleteNotificationRequest) GetTareaId() string {
	if x != nil {
		return x.TareaId
	}
	return ""
}

type DeleteNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteNotificationResponse) Reset() {
	*x = DeleteNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNotificationResponse) ProtoMessage() {}

func (x *DeleteNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNotificationResponse.ProtoReflect.Descriptor instead.
func (*DeleteNotificationResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNotificationResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type ListActiveNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsuarioId string `protobuf:"bytes,1,opt,name=usuario_id,json=usuarioId,proto3" json:"usuario_id,omitempty"`
	TareaId   string `protobuf:"bytes,2,opt,name=tarea_id,json=tareaId,proto3" json:"tarea_id,omitempty"`
}

func (x *ListActiveNotificationRequest) Reset() {
	*x = ListActiveNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActiveNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveNotificationRequest) ProtoMessage() {}

func (x *ListActiveNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActiveNotificationRequest.ProtoReflect.Descriptor instead.
func (*ListActiveNotificationRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{6}
}

func (x *ListActiveNotificationRequest) GetUsuarioId() string {
	if x != nil {
		return x.UsuarioId
	}
	return ""
}

func (x *ListActiveNotificationRequest) GetTareaId() string {
	if x != nil {
		return x.TareaId
	}
	return ""
}

type ListActiveNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *ListActiveNotificationResponse) Reset() {
	*x = ListActiveNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActiveNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveNotificationResponse) ProtoMessage() {}

func (x *ListActiveNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActiveNotificationResponse.ProtoReflect.Descriptor instead.
func (*ListActiveNotificationResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{7}
}

func (x *ListActiveNotificationResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Titulo       string `protobuf:"bytes,1,opt,name=titulo,proto3" json:"titulo,omitempty"`
	FechaTermino string `protobuf:"bytes,2,opt,name=fecha_termino,json=fechaTermino,proto3" json:"fecha_termino,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{8}
}

func (x *Notification) GetTitulo() string {
	if x != nil {
		return x.Titulo
	}
	return ""
}

func (x *Notification) GetFechaTermino() string {
	if x != nil {
		return x.FechaTermino
	}
	return ""
}

var File_proto_services_proto protoreflect.FileDescriptor

var file_proto_services_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x69, 0x74, 0x75, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x74,
	0x75, 0x6c, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x65, 0x63, 0x68, 0x61, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65, 0x63, 0x68, 0x61, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x6f, 0x22, 0x34, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x74, 0x75, 0x6c, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x74, 0x75, 0x6c, 0x6f, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x63, 0x68,
	0x61, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x66, 0x65, 0x63, 0x68, 0x61, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x6f, 0x22, 0x34, 0x0a,
	0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x55, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x1a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x59, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x1e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4b, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x74, 0x75,
	0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x74, 0x75, 0x6c, 0x6f,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x63, 0x68, 0x61, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65, 0x63, 0x68, 0x61, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x6f, 0x32, 0x8d, 0x03, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x71, 0x75, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x61,
	0x44, 0x65, 0x53, 0x69, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_services_proto_rawDescOnce sync.Once
	file_proto_services_proto_rawDescData = file_proto_services_proto_rawDesc
)

func file_proto_services_proto_rawDescGZIP() []byte {
	file_proto_services_proto_rawDescOnce.Do(func() {
		file_proto_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_services_proto_rawDescData)
	})
	return file_proto_services_proto_rawDescData
}

var file_proto_services_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_services_proto_goTypes = []interface{}{
	(*CreateNotificationRequest)(nil),      // 0: proto.CreateNotificationRequest
	(*CreateNotificationResponse)(nil),     // 1: proto.CreateNotificationResponse
	(*UpdateNotificationRequest)(nil),      // 2: proto.UpdateNotificationRequest
	(*UpdateNotificationResponse)(nil),     // 3: proto.UpdateNotificationResponse
	(*DeleteNotificationRequest)(nil),      // 4: proto.DeleteNotificationRequest
	(*DeleteNotificationResponse)(nil),     // 5: proto.DeleteNotificationResponse
	(*ListActiveNotificationRequest)(nil),  // 6: proto.ListActiveNotificationRequest
	(*ListActiveNotificationResponse)(nil), // 7: proto.ListActiveNotificationResponse
	(*Notification)(nil),                   // 8: proto.Notification
}
var file_proto_services_proto_depIdxs = []int32{
	8, // 0: proto.ListActiveNotificationResponse.notifications:type_name -> proto.Notification
	0, // 1: proto.NotificationService.CreateNotification:input_type -> proto.CreateNotificationRequest
	2, // 2: proto.NotificationService.UpdateNotification:input_type -> proto.UpdateNotificationRequest
	4, // 3: proto.NotificationService.DeleteNotification:input_type -> proto.DeleteNotificationRequest
	6, // 4: proto.NotificationService.ListActiveNotification:input_type -> proto.ListActiveNotificationRequest
	1, // 5: proto.NotificationService.CreateNotification:output_type -> proto.CreateNotificationResponse
	3, // 6: proto.NotificationService.UpdateNotification:output_type -> proto.UpdateNotificationResponse
	5, // 7: proto.NotificationService.DeleteNotification:output_type -> proto.DeleteNotificationResponse
	7, // 8: proto.NotificationService.ListActiveNotification:output_type -> proto.ListActiveNotificationResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_services_proto_init() }
func file_proto_services_proto_init() {
	if File_proto_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNotificationRequest); i {
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
		file_proto_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNotificationResponse); i {
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
		file_proto_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNotificationRequest); i {
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
		file_proto_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNotificationResponse); i {
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
		file_proto_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNotificationRequest); i {
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
		file_proto_services_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNotificationResponse); i {
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
		file_proto_services_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActiveNotificationRequest); i {
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
		file_proto_services_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActiveNotificationResponse); i {
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
		file_proto_services_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_proto_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_services_proto_goTypes,
		DependencyIndexes: file_proto_services_proto_depIdxs,
		MessageInfos:      file_proto_services_proto_msgTypes,
	}.Build()
	File_proto_services_proto = out.File
	file_proto_services_proto_rawDesc = nil
	file_proto_services_proto_goTypes = nil
	file_proto_services_proto_depIdxs = nil
}
