// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scrapper/indeedclone.proto

package scrapper

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

type SearchRequest struct {
	Term                 string   `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4728f0b566fded20, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

type JobObject struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Location             string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Salary               string   `protobuf:"bytes,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Summary              string   `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobObject) Reset()         { *m = JobObject{} }
func (m *JobObject) String() string { return proto.CompactTextString(m) }
func (*JobObject) ProtoMessage()    {}
func (*JobObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_4728f0b566fded20, []int{1}
}

func (m *JobObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobObject.Unmarshal(m, b)
}
func (m *JobObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobObject.Marshal(b, m, deterministic)
}
func (m *JobObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobObject.Merge(m, src)
}
func (m *JobObject) XXX_Size() int {
	return xxx_messageInfo_JobObject.Size(m)
}
func (m *JobObject) XXX_DiscardUnknown() {
	xxx_messageInfo_JobObject.DiscardUnknown(m)
}

var xxx_messageInfo_JobObject proto.InternalMessageInfo

func (m *JobObject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JobObject) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *JobObject) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *JobObject) GetSalary() string {
	if m != nil {
		return m.Salary
	}
	return ""
}

func (m *JobObject) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

type SearchResponse struct {
	Jobs                 []*JobObject `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4728f0b566fded20, []int{2}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetJobs() []*JobObject {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "scrapper.searchRequest")
	proto.RegisterType((*JobObject)(nil), "scrapper.jobObject")
	proto.RegisterType((*SearchResponse)(nil), "scrapper.searchResponse")
}

func init() {
	proto.RegisterFile("scrapper/indeedclone.proto", fileDescriptor_4728f0b566fded20)
}

var fileDescriptor_4728f0b566fded20 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x49, 0x9a, 0x86, 0xf6, 0x10, 0x1d, 0x0e, 0x04, 0x56, 0xa6, 0x2a, 0x0c, 0x74, 0x0a,
	0x52, 0x99, 0x18, 0x78, 0x02, 0x06, 0xa4, 0xf4, 0x09, 0x6c, 0xe7, 0x24, 0x1c, 0x25, 0xb9, 0x60,
	0xbb, 0x48, 0x4c, 0xbc, 0x3a, 0xe2, 0xda, 0x14, 0x21, 0x36, 0xff, 0xfc, 0xf3, 0x9f, 0xef, 0x3e,
	0x28, 0x82, 0xf5, 0x7a, 0x1c, 0xc9, 0x3f, 0xb8, 0xa1, 0x21, 0x6a, 0x6c, 0xc7, 0x03, 0x55, 0xa3,
	0xe7, 0xc8, 0xb8, 0x98, 0x5c, 0x79, 0x07, 0x97, 0x81, 0xb4, 0xb7, 0x6f, 0x35, 0xbd, 0xef, 0x29,
	0x44, 0x44, 0xc8, 0x22, 0xf9, 0x5e, 0x25, 0xeb, 0x64, 0xb3, 0xac, 0x65, 0x5d, 0x7e, 0xc1, 0xb2,
	0x65, 0xf3, 0x6a, 0x5a, 0xb2, 0x11, 0x57, 0x90, 0xba, 0xe6, 0xa8, 0x53, 0xd7, 0xe0, 0x35, 0xcc,
	0xa3, 0x8b, 0x1d, 0xa9, 0x54, 0xb6, 0x0e, 0x80, 0x05, 0x2c, 0x3a, 0xb6, 0x3a, 0x3a, 0x1e, 0xd4,
	0x4c, 0xc4, 0x89, 0xf1, 0x06, 0xf2, 0xa0, 0x3b, 0xed, 0x3f, 0x55, 0x26, 0xe6, 0x48, 0xa8, 0xe0,
	0x3c, 0xec, 0xfb, 0xfe, 0x47, 0xcc, 0x45, 0x4c, 0x58, 0x3e, 0xc1, 0x6a, 0x4a, 0x19, 0x46, 0x1e,
	0x02, 0xe1, 0x3d, 0x64, 0x2d, 0x9b, 0xa0, 0x92, 0xf5, 0x6c, 0x73, 0xb1, 0xbd, 0xaa, 0xa6, 0x81,
	0xaa, 0x53, 0xd0, 0x5a, 0x0e, 0x6c, 0x5f, 0x00, 0x5a, 0x36, 0x3b, 0xf2, 0x1f, 0xce, 0x12, 0x3e,
	0x43, 0xbe, 0x93, 0x87, 0xf0, 0xf6, 0xf7, 0xca, 0x9f, 0x02, 0x0a, 0xf5, 0x5f, 0x1c, 0xfe, 0x2c,
	0xcf, 0x4c, 0x2e, 0xf5, 0x3d, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x72, 0x40, 0xa5, 0x3e, 0x5c,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/scrapper.jobService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scrapper.jobService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scrapper.jobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _JobService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scrapper/indeedclone.proto",
}