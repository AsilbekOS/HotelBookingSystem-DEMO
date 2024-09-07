// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.11.2
// source: proto/booking.proto

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

type CreateBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID       int64   `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	HotelID      int64   `protobuf:"varint,2,opt,name=hotelID,proto3" json:"hotelID,omitempty"`
	RoomID       int64   `protobuf:"varint,3,opt,name=roomID,proto3" json:"roomID,omitempty"`
	RoomType     string  `protobuf:"bytes,4,opt,name=roomType,proto3" json:"roomType,omitempty"`
	CheckInDate  string  `protobuf:"bytes,5,opt,name=checkInDate,proto3" json:"checkInDate,omitempty"`
	CheckOutDate string  `protobuf:"bytes,6,opt,name=checkOutDate,proto3" json:"checkOutDate,omitempty"`
	TotalAmount  float64 `protobuf:"fixed64,7,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
}

func (x *CreateBookingRequest) Reset() {
	*x = CreateBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingRequest) ProtoMessage() {}

func (x *CreateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingRequest.ProtoReflect.Descriptor instead.
func (*CreateBookingRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBookingRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateBookingRequest) GetHotelID() int64 {
	if x != nil {
		return x.HotelID
	}
	return 0
}

func (x *CreateBookingRequest) GetRoomID() int64 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

func (x *CreateBookingRequest) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *CreateBookingRequest) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *CreateBookingRequest) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

func (x *CreateBookingRequest) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

type UpdateBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingID    int64   `protobuf:"varint,1,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
	CheckInDate  string  `protobuf:"bytes,2,opt,name=checkInDate,proto3" json:"checkInDate,omitempty"`
	CheckOutDate string  `protobuf:"bytes,3,opt,name=checkOutDate,proto3" json:"checkOutDate,omitempty"`
	TotalAmount  float64 `protobuf:"fixed64,4,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	Status       string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateBookingRequest) Reset() {
	*x = UpdateBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookingRequest) ProtoMessage() {}

func (x *UpdateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookingRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookingRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBookingRequest) GetBookingID() int64 {
	if x != nil {
		return x.BookingID
	}
	return 0
}

func (x *UpdateBookingRequest) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *UpdateBookingRequest) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

func (x *UpdateBookingRequest) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *UpdateBookingRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingID int64 `protobuf:"varint,1,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
}

func (x *GetBookingRequest) Reset() {
	*x = GetBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingRequest) ProtoMessage() {}

func (x *GetBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingRequest.ProtoReflect.Descriptor instead.
func (*GetBookingRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookingRequest) GetBookingID() int64 {
	if x != nil {
		return x.BookingID
	}
	return 0
}

type CancelBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingID int64 `protobuf:"varint,1,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
	RoomID    int64 `protobuf:"varint,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *CancelBookingRequest) Reset() {
	*x = CancelBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBookingRequest) ProtoMessage() {}

func (x *CancelBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBookingRequest.ProtoReflect.Descriptor instead.
func (*CancelBookingRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{3}
}

func (x *CancelBookingRequest) GetBookingID() int64 {
	if x != nil {
		return x.BookingID
	}
	return 0
}

func (x *CancelBookingRequest) GetRoomID() int64 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

type ListUserBookingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *ListUserBookingsRequest) Reset() {
	*x = ListUserBookingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserBookingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserBookingsRequest) ProtoMessage() {}

func (x *ListUserBookingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserBookingsRequest.ProtoReflect.Descriptor instead.
func (*ListUserBookingsRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserBookingsRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingID    int64   `protobuf:"varint,1,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
	UserID       int64   `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	HotelID      int64   `protobuf:"varint,3,opt,name=hotelID,proto3" json:"hotelID,omitempty"`
	RoomID       int64   `protobuf:"varint,4,opt,name=roomID,proto3" json:"roomID,omitempty"`
	RoomType     string  `protobuf:"bytes,5,opt,name=roomType,proto3" json:"roomType,omitempty"`
	CheckInDate  string  `protobuf:"bytes,6,opt,name=checkInDate,proto3" json:"checkInDate,omitempty"`
	CheckOutDate string  `protobuf:"bytes,7,opt,name=checkOutDate,proto3" json:"checkOutDate,omitempty"`
	TotalAmount  float64 `protobuf:"fixed64,8,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	Status       string  `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{5}
}

func (x *BookingResponse) GetBookingID() int64 {
	if x != nil {
		return x.BookingID
	}
	return 0
}

func (x *BookingResponse) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *BookingResponse) GetHotelID() int64 {
	if x != nil {
		return x.HotelID
	}
	return 0
}

func (x *BookingResponse) GetRoomID() int64 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

func (x *BookingResponse) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *BookingResponse) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *BookingResponse) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

func (x *BookingResponse) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *BookingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CancelBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	BookingID int64  `protobuf:"varint,2,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
}

func (x *CancelBookingResponse) Reset() {
	*x = CancelBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBookingResponse) ProtoMessage() {}

func (x *CancelBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBookingResponse.ProtoReflect.Descriptor instead.
func (*CancelBookingResponse) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{6}
}

func (x *CancelBookingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CancelBookingResponse) GetBookingID() int64 {
	if x != nil {
		return x.BookingID
	}
	return 0
}

type ListUserBookingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*BookingResponse `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
}

func (x *ListUserBookingsResponse) Reset() {
	*x = ListUserBookingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserBookingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserBookingsResponse) ProtoMessage() {}

func (x *ListUserBookingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserBookingsResponse.ProtoReflect.Descriptor instead.
func (*ListUserBookingsResponse) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{7}
}

func (x *ListUserBookingsResponse) GetBookings() []*BookingResponse {
	if x != nil {
		return x.Bookings
	}
	return nil
}

var File_proto_booking_proto protoreflect.FileDescriptor

var file_proto_booking_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb4, 0x01, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x31, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x22, 0x4c, 0x0a, 0x14, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x22, 0x31, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x95, 0x02, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x4f, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44,
	0x22, 0x48, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x32, 0xc1, 0x02, 0x0a, 0x0e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x15,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_booking_proto_rawDescOnce sync.Once
	file_proto_booking_proto_rawDescData = file_proto_booking_proto_rawDesc
)

func file_proto_booking_proto_rawDescGZIP() []byte {
	file_proto_booking_proto_rawDescOnce.Do(func() {
		file_proto_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_booking_proto_rawDescData)
	})
	return file_proto_booking_proto_rawDescData
}

var file_proto_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_booking_proto_goTypes = []any{
	(*CreateBookingRequest)(nil),     // 0: CreateBookingRequest
	(*UpdateBookingRequest)(nil),     // 1: UpdateBookingRequest
	(*GetBookingRequest)(nil),        // 2: GetBookingRequest
	(*CancelBookingRequest)(nil),     // 3: CancelBookingRequest
	(*ListUserBookingsRequest)(nil),  // 4: ListUserBookingsRequest
	(*BookingResponse)(nil),          // 5: BookingResponse
	(*CancelBookingResponse)(nil),    // 6: CancelBookingResponse
	(*ListUserBookingsResponse)(nil), // 7: ListUserBookingsResponse
}
var file_proto_booking_proto_depIdxs = []int32{
	5, // 0: ListUserBookingsResponse.bookings:type_name -> BookingResponse
	0, // 1: BookingService.CreateBooking:input_type -> CreateBookingRequest
	2, // 2: BookingService.GetBooking:input_type -> GetBookingRequest
	1, // 3: BookingService.UpdateBooking:input_type -> UpdateBookingRequest
	3, // 4: BookingService.CancelBooking:input_type -> CancelBookingRequest
	4, // 5: BookingService.ListUserBookings:input_type -> ListUserBookingsRequest
	5, // 6: BookingService.CreateBooking:output_type -> BookingResponse
	5, // 7: BookingService.GetBooking:output_type -> BookingResponse
	5, // 8: BookingService.UpdateBooking:output_type -> BookingResponse
	6, // 9: BookingService.CancelBooking:output_type -> CancelBookingResponse
	7, // 10: BookingService.ListUserBookings:output_type -> ListUserBookingsResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_booking_proto_init() }
func file_proto_booking_proto_init() {
	if File_proto_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_booking_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateBookingRequest); i {
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
		file_proto_booking_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBookingRequest); i {
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
		file_proto_booking_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetBookingRequest); i {
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
		file_proto_booking_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CancelBookingRequest); i {
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
		file_proto_booking_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserBookingsRequest); i {
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
		file_proto_booking_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BookingResponse); i {
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
		file_proto_booking_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CancelBookingResponse); i {
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
		file_proto_booking_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserBookingsResponse); i {
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
			RawDescriptor: file_proto_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_booking_proto_goTypes,
		DependencyIndexes: file_proto_booking_proto_depIdxs,
		MessageInfos:      file_proto_booking_proto_msgTypes,
	}.Build()
	File_proto_booking_proto = out.File
	file_proto_booking_proto_rawDesc = nil
	file_proto_booking_proto_goTypes = nil
	file_proto_booking_proto_depIdxs = nil
}
