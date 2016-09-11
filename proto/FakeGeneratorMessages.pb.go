// Code generated by protoc-gen-go.
// source: FakeGeneratorMessages.proto
// DO NOT EDIT!

package fakegenerator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Address struct {
	Continent     string  `protobuf:"bytes,1,opt,name=Continent,json=continent" json:"Continent,omitempty"`
	Country       string  `protobuf:"bytes,2,opt,name=Country,json=country" json:"Country,omitempty"`
	City          string  `protobuf:"bytes,3,opt,name=City,json=city" json:"City,omitempty"`
	State         string  `protobuf:"bytes,4,opt,name=State,json=state" json:"State,omitempty"`
	StreetAddress string  `protobuf:"bytes,5,opt,name=StreetAddress,json=streetAddress" json:"StreetAddress,omitempty"`
	Zip           string  `protobuf:"bytes,6,opt,name=Zip,json=zip" json:"Zip,omitempty"`
	Latitude      float32 `protobuf:"fixed32,7,opt,name=Latitude,json=latitude" json:"Latitude,omitempty"`
	Longitude     float32 `protobuf:"fixed32,8,opt,name=Longitude,json=longitude" json:"Longitude,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Date struct {
	Year  int64 `protobuf:"varint,1,opt,name=Year,json=year" json:"Year,omitempty"`
	Month int64 `protobuf:"varint,2,opt,name=Month,json=month" json:"Month,omitempty"`
	Day   int64 `protobuf:"varint,3,opt,name=Day,json=day" json:"Day,omitempty"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type User struct {
	FullName    string   `protobuf:"bytes,1,opt,name=FullName,json=fullName" json:"FullName,omitempty"`
	Username    string   `protobuf:"bytes,2,opt,name=Username,json=username" json:"Username,omitempty"`
	Password    string   `protobuf:"bytes,3,opt,name=Password,json=password" json:"Password,omitempty"`
	Email       string   `protobuf:"bytes,4,opt,name=Email,json=email" json:"Email,omitempty"`
	PhoneNumber string   `protobuf:"bytes,5,opt,name=PhoneNumber,json=phoneNumber" json:"PhoneNumber,omitempty"`
	Gender      string   `protobuf:"bytes,6,opt,name=Gender,json=gender" json:"Gender,omitempty"`
	Language    string   `protobuf:"bytes,7,opt,name=Language,json=language" json:"Language,omitempty"`
	Birthday    *Date    `protobuf:"bytes,8,opt,name=Birthday,json=birthday" json:"Birthday,omitempty"`
	Address     *Address `protobuf:"bytes,9,opt,name=Address,json=address" json:"Address,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *User) GetBirthday() *Date {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *User) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type Email struct {
	Title string `protobuf:"bytes,1,opt,name=Title,json=title" json:"Title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=Body,json=body" json:"Body,omitempty"`
	From  string `protobuf:"bytes,3,opt,name=From,json=from" json:"From,omitempty"`
	To    string `protobuf:"bytes,4,opt,name=To,json=to" json:"To,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*Address)(nil), "fakegenerator.Address")
	proto.RegisterType((*Date)(nil), "fakegenerator.Date")
	proto.RegisterType((*User)(nil), "fakegenerator.User")
	proto.RegisterType((*Email)(nil), "fakegenerator.Email")
	proto.RegisterType((*EmptyMessage)(nil), "fakegenerator.EmptyMessage")
}

func init() { proto.RegisterFile("FakeGeneratorMessages.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0xcb, 0x6e, 0xdb, 0x3c,
	0x10, 0x85, 0x61, 0x5d, 0x2c, 0x89, 0xfe, 0x1d, 0xfc, 0x60, 0x8b, 0x40, 0x68, 0xbb, 0x30, 0x84,
	0x2e, 0xb2, 0x72, 0x8b, 0xf6, 0x09, 0x6a, 0x27, 0xce, 0x26, 0x09, 0x02, 0x25, 0x59, 0xb4, 0x3b,
	0xda, 0x1a, 0xcb, 0x44, 0x24, 0x52, 0x20, 0x47, 0x28, 0xd4, 0x37, 0xea, 0x43, 0xf5, 0x5d, 0x8a,
	0xa1, 0xa8, 0xde, 0x76, 0x3a, 0xe7, 0x0c, 0xc0, 0x99, 0x4f, 0x87, 0xbd, 0xde, 0x89, 0x67, 0xb8,
	0x06, 0x05, 0x46, 0xa0, 0x36, 0xb7, 0x60, 0xad, 0xa8, 0xc1, 0xae, 0x3b, 0xa3, 0x51, 0xf3, 0xe5,
	0x51, 0x3c, 0x43, 0x3d, 0x85, 0xc5, 0x8f, 0x19, 0x4b, 0x3e, 0x55, 0x95, 0x01, 0x6b, 0xf9, 0x1b,
	0x96, 0x6d, 0xb5, 0x42, 0xa9, 0x40, 0x61, 0x3e, 0x5b, 0xcd, 0x2e, 0xb2, 0x32, 0x3b, 0x4c, 0x06,
	0xcf, 0x59, 0xb2, 0xd5, 0xbd, 0x42, 0x33, 0xe4, 0x81, 0xcb, 0x92, 0xc3, 0x28, 0x39, 0x67, 0xd1,
	0x56, 0xe2, 0x90, 0x87, 0xce, 0x8e, 0x0e, 0x12, 0x07, 0xfe, 0x92, 0xc5, 0x0f, 0x28, 0x10, 0xf2,
	0xc8, 0x99, 0xb1, 0x25, 0xc1, 0xdf, 0xb2, 0xe5, 0x03, 0x1a, 0x00, 0xf4, 0x4f, 0xe6, 0xb1, 0x4b,
	0x97, 0xf6, 0x4f, 0x93, 0xff, 0xcf, 0xc2, 0x2f, 0xb2, 0xcb, 0xe7, 0x2e, 0x0b, 0xbf, 0xc9, 0x8e,
	0xbf, 0x62, 0xe9, 0x8d, 0x40, 0x89, 0x7d, 0x05, 0x79, 0xb2, 0x9a, 0x5d, 0x04, 0x65, 0xda, 0x78,
	0x4d, 0x5b, 0xdf, 0x68, 0x55, 0x8f, 0x61, 0xea, 0xc2, 0xac, 0x99, 0x8c, 0x62, 0xc3, 0xa2, 0x4b,
	0x7a, 0x99, 0xb3, 0xe8, 0x33, 0x08, 0xe3, 0xce, 0x0a, 0xcb, 0x68, 0x00, 0x61, 0x68, 0xc7, 0x5b,
	0xad, 0xf0, 0xe4, 0xee, 0x09, 0xcb, 0xb8, 0x25, 0x41, 0xaf, 0x5f, 0x8a, 0xf1, 0x98, 0xb0, 0x0c,
	0x2b, 0x31, 0x14, 0xdf, 0x03, 0x16, 0x3d, 0x59, 0x30, 0xb4, 0xc6, 0xae, 0x6f, 0x9a, 0x3b, 0xd1,
	0x82, 0xe7, 0x93, 0x1e, 0xbd, 0xa6, 0x8c, 0x66, 0x14, 0x65, 0x23, 0x9f, 0xb4, 0xf7, 0x9a, 0xb2,
	0x7b, 0x61, 0xed, 0x57, 0x6d, 0x2a, 0x0f, 0x29, 0xed, 0xbc, 0xa6, 0x25, 0xae, 0x5a, 0x21, 0x9b,
	0x09, 0x14, 0x90, 0xe0, 0x2b, 0xb6, 0xb8, 0x3f, 0x69, 0x05, 0x77, 0x7d, 0xbb, 0x07, 0xe3, 0x31,
	0x2d, 0xba, 0xdf, 0x16, 0x3f, 0x67, 0xf3, 0x6b, 0x50, 0x15, 0x18, 0xcf, 0x69, 0x5e, 0x3b, 0x35,
	0xa2, 0x52, 0x75, 0x2f, 0xea, 0x11, 0x55, 0x46, 0xa8, 0x46, 0xcd, 0xdf, 0xb1, 0x74, 0x23, 0x0d,
	0x9e, 0x2a, 0x31, 0x38, 0x52, 0x8b, 0x0f, 0x2f, 0xd6, 0x7f, 0xd5, 0x61, 0x4d, 0xac, 0xca, 0x74,
	0xef, 0x87, 0xf8, 0xfb, 0x5f, 0xe5, 0xc8, 0x33, 0x37, 0x7f, 0xfe, 0xcf, 0xbc, 0x4f, 0xcb, 0x44,
	0x8c, 0x1f, 0xc5, 0x93, 0x3f, 0x87, 0xee, 0x7a, 0x94, 0xd8, 0x4c, 0xa0, 0x62, 0x24, 0x41, 0xbf,
	0x61, 0xa3, 0xab, 0xa9, 0x41, 0xd1, 0x5e, 0x57, 0xae, 0x3e, 0x3b, 0xa3, 0xdb, 0xa9, 0x3e, 0x47,
	0xa3, 0x5b, 0x7e, 0xc6, 0x82, 0x47, 0xed, 0x91, 0x04, 0xa8, 0x8b, 0x33, 0xf6, 0xdf, 0x55, 0xdb,
	0xe1, 0xe0, 0xcb, 0xbc, 0x9f, 0xbb, 0x32, 0x7f, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0x87, 0xd0,
	0x2a, 0xa3, 0xeb, 0x02, 0x00, 0x00,
}