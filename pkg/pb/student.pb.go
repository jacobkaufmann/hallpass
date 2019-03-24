// Code generated by protoc-gen-go. DO NOT EDIT.
// source: student.proto

package hallpass

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Student represents a student at a school
type Student struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	School               string   `protobuf:"bytes,2,opt,name=school,proto3" json:"school,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{0}
}

func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Student) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *Student) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Student) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type AddStudentRequest struct {
	School               string   `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStudentRequest) Reset()         { *m = AddStudentRequest{} }
func (m *AddStudentRequest) String() string { return proto.CompactTextString(m) }
func (*AddStudentRequest) ProtoMessage()    {}
func (*AddStudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{1}
}

func (m *AddStudentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStudentRequest.Unmarshal(m, b)
}
func (m *AddStudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStudentRequest.Marshal(b, m, deterministic)
}
func (m *AddStudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStudentRequest.Merge(m, src)
}
func (m *AddStudentRequest) XXX_Size() int {
	return xxx_messageInfo_AddStudentRequest.Size(m)
}
func (m *AddStudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddStudentRequest proto.InternalMessageInfo

func (m *AddStudentRequest) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *AddStudentRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *AddStudentRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type AddStudentResponse struct {
	Added                bool     `protobuf:"varint,1,opt,name=added,proto3" json:"added,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStudentResponse) Reset()         { *m = AddStudentResponse{} }
func (m *AddStudentResponse) String() string { return proto.CompactTextString(m) }
func (*AddStudentResponse) ProtoMessage()    {}
func (*AddStudentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{2}
}

func (m *AddStudentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStudentResponse.Unmarshal(m, b)
}
func (m *AddStudentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStudentResponse.Marshal(b, m, deterministic)
}
func (m *AddStudentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStudentResponse.Merge(m, src)
}
func (m *AddStudentResponse) XXX_Size() int {
	return xxx_messageInfo_AddStudentResponse.Size(m)
}
func (m *AddStudentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStudentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddStudentResponse proto.InternalMessageInfo

func (m *AddStudentResponse) GetAdded() bool {
	if m != nil {
		return m.Added
	}
	return false
}

func (m *AddStudentResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UpdateStudentRequest struct {
	School               string   `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStudentRequest) Reset()         { *m = UpdateStudentRequest{} }
func (m *UpdateStudentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateStudentRequest) ProtoMessage()    {}
func (*UpdateStudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{3}
}

func (m *UpdateStudentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStudentRequest.Unmarshal(m, b)
}
func (m *UpdateStudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStudentRequest.Marshal(b, m, deterministic)
}
func (m *UpdateStudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStudentRequest.Merge(m, src)
}
func (m *UpdateStudentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateStudentRequest.Size(m)
}
func (m *UpdateStudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStudentRequest proto.InternalMessageInfo

func (m *UpdateStudentRequest) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

type UpdateStudentResponse struct {
	Updated              bool     `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStudentResponse) Reset()         { *m = UpdateStudentResponse{} }
func (m *UpdateStudentResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateStudentResponse) ProtoMessage()    {}
func (*UpdateStudentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{4}
}

func (m *UpdateStudentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStudentResponse.Unmarshal(m, b)
}
func (m *UpdateStudentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStudentResponse.Marshal(b, m, deterministic)
}
func (m *UpdateStudentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStudentResponse.Merge(m, src)
}
func (m *UpdateStudentResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateStudentResponse.Size(m)
}
func (m *UpdateStudentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStudentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStudentResponse proto.InternalMessageInfo

func (m *UpdateStudentResponse) GetUpdated() bool {
	if m != nil {
		return m.Updated
	}
	return false
}

func (m *UpdateStudentResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetStudentRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStudentRequest) Reset()         { *m = GetStudentRequest{} }
func (m *GetStudentRequest) String() string { return proto.CompactTextString(m) }
func (*GetStudentRequest) ProtoMessage()    {}
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{5}
}

func (m *GetStudentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStudentRequest.Unmarshal(m, b)
}
func (m *GetStudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStudentRequest.Marshal(b, m, deterministic)
}
func (m *GetStudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentRequest.Merge(m, src)
}
func (m *GetStudentRequest) XXX_Size() int {
	return xxx_messageInfo_GetStudentRequest.Size(m)
}
func (m *GetStudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentRequest proto.InternalMessageInfo

func (m *GetStudentRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListStudentsRequest struct {
	School               string   `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListStudentsRequest) Reset()         { *m = ListStudentsRequest{} }
func (m *ListStudentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListStudentsRequest) ProtoMessage()    {}
func (*ListStudentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{6}
}

func (m *ListStudentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStudentsRequest.Unmarshal(m, b)
}
func (m *ListStudentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStudentsRequest.Marshal(b, m, deterministic)
}
func (m *ListStudentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStudentsRequest.Merge(m, src)
}
func (m *ListStudentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListStudentsRequest.Size(m)
}
func (m *ListStudentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStudentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListStudentsRequest proto.InternalMessageInfo

func (m *ListStudentsRequest) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *ListStudentsRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ListStudentsRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func init() {
	proto.RegisterType((*Student)(nil), "hallpass.Student")
	proto.RegisterType((*AddStudentRequest)(nil), "hallpass.AddStudentRequest")
	proto.RegisterType((*AddStudentResponse)(nil), "hallpass.AddStudentResponse")
	proto.RegisterType((*UpdateStudentRequest)(nil), "hallpass.UpdateStudentRequest")
	proto.RegisterType((*UpdateStudentResponse)(nil), "hallpass.UpdateStudentResponse")
	proto.RegisterType((*GetStudentRequest)(nil), "hallpass.GetStudentRequest")
	proto.RegisterType((*ListStudentsRequest)(nil), "hallpass.ListStudentsRequest")
}

func init() { proto.RegisterFile("student.proto", fileDescriptor_94a1c1b032ad0c00) }

var fileDescriptor_94a1c1b032ad0c00 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0x87, 0x69, 0x79, 0x5f, 0xfe, 0x4c, 0x84, 0x84, 0x11, 0x4d, 0x03, 0xa2, 0x64, 0xbd, 0x70,
	0x6a, 0x8c, 0xde, 0x4d, 0x4c, 0x48, 0x8c, 0xd1, 0x78, 0x28, 0xf1, 0x6c, 0x56, 0x76, 0x84, 0x26,
	0x2d, 0xad, 0xcc, 0xe2, 0x37, 0xf3, 0xfb, 0x19, 0xda, 0xae, 0xad, 0xa5, 0x12, 0x2f, 0x1e, 0x67,
	0x9f, 0xe9, 0x3e, 0xd3, 0xdf, 0xb4, 0xd0, 0x61, 0xbd, 0x51, 0xb4, 0xd2, 0x6e, 0xbc, 0x8e, 0x74,
	0x84, 0xad, 0xa5, 0x0c, 0x82, 0x58, 0x32, 0x8b, 0x10, 0x9a, 0xb3, 0x14, 0x61, 0x17, 0x6c, 0x5f,
	0x39, 0xd6, 0xd8, 0x9a, 0xd4, 0x3d, 0xdb, 0x57, 0x78, 0x0c, 0x0d, 0x9e, 0x2f, 0xa3, 0x28, 0x70,
	0xec, 0xb1, 0x35, 0x69, 0x7b, 0x59, 0x85, 0x23, 0x80, 0x57, 0x7f, 0xcd, 0xfa, 0x79, 0x25, 0x43,
	0x72, 0xea, 0x09, 0x6b, 0x27, 0x27, 0x8f, 0x32, 0x24, 0x1c, 0x42, 0x3b, 0x90, 0x86, 0xfe, 0x4b,
	0x68, 0x6b, 0x7b, 0xb0, 0x85, 0x62, 0x01, 0xbd, 0x1b, 0xa5, 0x32, 0xa3, 0x47, 0x6f, 0x1b, 0x62,
	0x5d, 0x10, 0x59, 0x7b, 0x44, 0xf6, 0x5e, 0x51, 0xbd, 0x24, 0x9a, 0x02, 0x16, 0x45, 0x1c, 0x47,
	0x2b, 0x26, 0xec, 0xc3, 0x7f, 0xa9, 0x14, 0xa5, 0x6f, 0xd9, 0xf2, 0xd2, 0x02, 0x1d, 0x68, 0x86,
	0xc4, 0x2c, 0x17, 0x46, 0x62, 0x4a, 0xe1, 0x42, 0xff, 0x29, 0x56, 0x52, 0xd3, 0xef, 0x26, 0x16,
	0xf7, 0x70, 0x54, 0xea, 0xcf, 0xc4, 0x0e, 0x34, 0x37, 0x09, 0x30, 0x6a, 0x53, 0xee, 0x91, 0x9f,
	0x43, 0xef, 0x96, 0x74, 0xc9, 0x5c, 0x5a, 0x92, 0xf0, 0xe1, 0xf0, 0xc1, 0x67, 0xd3, 0xc5, 0x7f,
	0x18, 0xe9, 0xe5, 0x87, 0x0d, 0xdd, 0xcc, 0x33, 0xa3, 0xf5, 0xbb, 0x3f, 0x27, 0xbc, 0x03, 0xc8,
	0x53, 0xc6, 0xa1, 0x6b, 0x3e, 0x2b, 0x77, 0x67, 0xc9, 0x83, 0x93, 0x6a, 0x98, 0xe6, 0x23, 0x6a,
	0xe8, 0x41, 0xe7, 0x5b, 0x74, 0x78, 0x9a, 0x3f, 0x50, 0xb5, 0x83, 0xc1, 0xd9, 0x8f, 0xfc, 0xeb,
	0xce, 0x6b, 0x80, 0x3c, 0xc1, 0xe2, 0x78, 0x3b, 0xb9, 0x0e, 0x7a, 0x39, 0xcc, 0x88, 0xa8, 0xe1,
	0x14, 0x0e, 0x8a, 0xe1, 0xe2, 0x28, 0x6f, 0xaa, 0x08, 0xbd, 0xf2, 0x8e, 0x0b, 0xeb, 0xa5, 0x91,
	0xfc, 0x73, 0x57, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x82, 0x6d, 0xdd, 0x84, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentServiceClient interface {
	AddStudent(ctx context.Context, in *AddStudentRequest, opts ...grpc.CallOption) (*AddStudentResponse, error)
	UpdateStudent(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*UpdateStudentResponse, error)
	GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*Student, error)
	ListStudents(ctx context.Context, in *ListStudentsRequest, opts ...grpc.CallOption) (StudentService_ListStudentsClient, error)
}

type studentServiceClient struct {
	cc *grpc.ClientConn
}

func NewStudentServiceClient(cc *grpc.ClientConn) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) AddStudent(ctx context.Context, in *AddStudentRequest, opts ...grpc.CallOption) (*AddStudentResponse, error) {
	out := new(AddStudentResponse)
	err := c.cc.Invoke(ctx, "/hallpass.StudentService/AddStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateStudent(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*UpdateStudentResponse, error) {
	out := new(UpdateStudentResponse)
	err := c.cc.Invoke(ctx, "/hallpass.StudentService/UpdateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/hallpass.StudentService/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ListStudents(ctx context.Context, in *ListStudentsRequest, opts ...grpc.CallOption) (StudentService_ListStudentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StudentService_serviceDesc.Streams[0], "/hallpass.StudentService/ListStudents", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentServiceListStudentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StudentService_ListStudentsClient interface {
	Recv() (*Student, error)
	grpc.ClientStream
}

type studentServiceListStudentsClient struct {
	grpc.ClientStream
}

func (x *studentServiceListStudentsClient) Recv() (*Student, error) {
	m := new(Student)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StudentServiceServer is the server API for StudentService service.
type StudentServiceServer interface {
	AddStudent(context.Context, *AddStudentRequest) (*AddStudentResponse, error)
	UpdateStudent(context.Context, *UpdateStudentRequest) (*UpdateStudentResponse, error)
	GetStudent(context.Context, *GetStudentRequest) (*Student, error)
	ListStudents(*ListStudentsRequest, StudentService_ListStudentsServer) error
}

// UnimplementedStudentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (*UnimplementedStudentServiceServer) AddStudent(ctx context.Context, req *AddStudentRequest) (*AddStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudent not implemented")
}
func (*UnimplementedStudentServiceServer) UpdateStudent(ctx context.Context, req *UpdateStudentRequest) (*UpdateStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (*UnimplementedStudentServiceServer) GetStudent(ctx context.Context, req *GetStudentRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (*UnimplementedStudentServiceServer) ListStudents(req *ListStudentsRequest, srv StudentService_ListStudentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStudents not implemented")
}

func RegisterStudentServiceServer(s *grpc.Server, srv StudentServiceServer) {
	s.RegisterService(&_StudentService_serviceDesc, srv)
}

func _StudentService_AddStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).AddStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hallpass.StudentService/AddStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).AddStudent(ctx, req.(*AddStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hallpass.StudentService/UpdateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).UpdateStudent(ctx, req.(*UpdateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hallpass.StudentService/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ListStudents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListStudentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentServiceServer).ListStudents(m, &studentServiceListStudentsServer{stream})
}

type StudentService_ListStudentsServer interface {
	Send(*Student) error
	grpc.ServerStream
}

type studentServiceListStudentsServer struct {
	grpc.ServerStream
}

func (x *studentServiceListStudentsServer) Send(m *Student) error {
	return x.ServerStream.SendMsg(m)
}

var _StudentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hallpass.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStudent",
			Handler:    _StudentService_AddStudent_Handler,
		},
		{
			MethodName: "UpdateStudent",
			Handler:    _StudentService_UpdateStudent_Handler,
		},
		{
			MethodName: "GetStudent",
			Handler:    _StudentService_GetStudent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStudents",
			Handler:       _StudentService_ListStudents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "student.proto",
}