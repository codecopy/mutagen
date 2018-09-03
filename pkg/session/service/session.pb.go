// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session/service/session.proto

package service // import "github.com/havoc-io/mutagen/pkg/session/service"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import session "github.com/havoc-io/mutagen/pkg/session"
import url "github.com/havoc-io/mutagen/pkg/url"

import (
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

type CreateRequest struct {
	Alpha                *url.URL               `protobuf:"bytes,1,opt,name=alpha,proto3" json:"alpha,omitempty"`
	Beta                 *url.URL               `protobuf:"bytes,2,opt,name=beta,proto3" json:"beta,omitempty"`
	Configuration        *session.Configuration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Response             string                 `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetAlpha() *url.URL {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func (m *CreateRequest) GetBeta() *url.URL {
	if m != nil {
		return m.Beta
	}
	return nil
}

func (m *CreateRequest) GetConfiguration() *session.Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *CreateRequest) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type CreateResponse struct {
	Session              string   `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Prompt               string   `protobuf:"bytes,3,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateResponse) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type ListRequest struct {
	PreviousStateIndex   uint64   `protobuf:"varint,1,opt,name=previousStateIndex,proto3" json:"previousStateIndex,omitempty"`
	Specifications       []string `protobuf:"bytes,2,rep,name=specifications,proto3" json:"specifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{2}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPreviousStateIndex() uint64 {
	if m != nil {
		return m.PreviousStateIndex
	}
	return 0
}

func (m *ListRequest) GetSpecifications() []string {
	if m != nil {
		return m.Specifications
	}
	return nil
}

type ListResponse struct {
	StateIndex           uint64           `protobuf:"varint,1,opt,name=stateIndex,proto3" json:"stateIndex,omitempty"`
	SessionStates        []*session.State `protobuf:"bytes,2,rep,name=sessionStates,proto3" json:"sessionStates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{3}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetStateIndex() uint64 {
	if m != nil {
		return m.StateIndex
	}
	return 0
}

func (m *ListResponse) GetSessionStates() []*session.State {
	if m != nil {
		return m.SessionStates
	}
	return nil
}

type PauseRequest struct {
	Specifications       []string `protobuf:"bytes,1,rep,name=specifications,proto3" json:"specifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseRequest) Reset()         { *m = PauseRequest{} }
func (m *PauseRequest) String() string { return proto.CompactTextString(m) }
func (*PauseRequest) ProtoMessage()    {}
func (*PauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{4}
}
func (m *PauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseRequest.Unmarshal(m, b)
}
func (m *PauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseRequest.Marshal(b, m, deterministic)
}
func (dst *PauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseRequest.Merge(dst, src)
}
func (m *PauseRequest) XXX_Size() int {
	return xxx_messageInfo_PauseRequest.Size(m)
}
func (m *PauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PauseRequest proto.InternalMessageInfo

func (m *PauseRequest) GetSpecifications() []string {
	if m != nil {
		return m.Specifications
	}
	return nil
}

type PauseResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseResponse) Reset()         { *m = PauseResponse{} }
func (m *PauseResponse) String() string { return proto.CompactTextString(m) }
func (*PauseResponse) ProtoMessage()    {}
func (*PauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{5}
}
func (m *PauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseResponse.Unmarshal(m, b)
}
func (m *PauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseResponse.Marshal(b, m, deterministic)
}
func (dst *PauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseResponse.Merge(dst, src)
}
func (m *PauseResponse) XXX_Size() int {
	return xxx_messageInfo_PauseResponse.Size(m)
}
func (m *PauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PauseResponse proto.InternalMessageInfo

func (m *PauseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ResumeRequest struct {
	Specifications       []string `protobuf:"bytes,1,rep,name=specifications,proto3" json:"specifications,omitempty"`
	Response             string   `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeRequest) Reset()         { *m = ResumeRequest{} }
func (m *ResumeRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeRequest) ProtoMessage()    {}
func (*ResumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{6}
}
func (m *ResumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeRequest.Unmarshal(m, b)
}
func (m *ResumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeRequest.Marshal(b, m, deterministic)
}
func (dst *ResumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeRequest.Merge(dst, src)
}
func (m *ResumeRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeRequest.Size(m)
}
func (m *ResumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeRequest proto.InternalMessageInfo

func (m *ResumeRequest) GetSpecifications() []string {
	if m != nil {
		return m.Specifications
	}
	return nil
}

func (m *ResumeRequest) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type ResumeResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Prompt               string   `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeResponse) Reset()         { *m = ResumeResponse{} }
func (m *ResumeResponse) String() string { return proto.CompactTextString(m) }
func (*ResumeResponse) ProtoMessage()    {}
func (*ResumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{7}
}
func (m *ResumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeResponse.Unmarshal(m, b)
}
func (m *ResumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeResponse.Marshal(b, m, deterministic)
}
func (dst *ResumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeResponse.Merge(dst, src)
}
func (m *ResumeResponse) XXX_Size() int {
	return xxx_messageInfo_ResumeResponse.Size(m)
}
func (m *ResumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeResponse proto.InternalMessageInfo

func (m *ResumeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResumeResponse) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type TerminateRequest struct {
	Specifications       []string `protobuf:"bytes,1,rep,name=specifications,proto3" json:"specifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateRequest) Reset()         { *m = TerminateRequest{} }
func (m *TerminateRequest) String() string { return proto.CompactTextString(m) }
func (*TerminateRequest) ProtoMessage()    {}
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{8}
}
func (m *TerminateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateRequest.Unmarshal(m, b)
}
func (m *TerminateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateRequest.Marshal(b, m, deterministic)
}
func (dst *TerminateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateRequest.Merge(dst, src)
}
func (m *TerminateRequest) XXX_Size() int {
	return xxx_messageInfo_TerminateRequest.Size(m)
}
func (m *TerminateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateRequest proto.InternalMessageInfo

func (m *TerminateRequest) GetSpecifications() []string {
	if m != nil {
		return m.Specifications
	}
	return nil
}

type TerminateResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateResponse) Reset()         { *m = TerminateResponse{} }
func (m *TerminateResponse) String() string { return proto.CompactTextString(m) }
func (*TerminateResponse) ProtoMessage()    {}
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_7aaea59be0ec6c8e, []int{9}
}
func (m *TerminateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateResponse.Unmarshal(m, b)
}
func (m *TerminateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateResponse.Marshal(b, m, deterministic)
}
func (dst *TerminateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateResponse.Merge(dst, src)
}
func (m *TerminateResponse) XXX_Size() int {
	return xxx_messageInfo_TerminateResponse.Size(m)
}
func (m *TerminateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateResponse proto.InternalMessageInfo

func (m *TerminateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "service.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "service.CreateResponse")
	proto.RegisterType((*ListRequest)(nil), "service.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "service.ListResponse")
	proto.RegisterType((*PauseRequest)(nil), "service.PauseRequest")
	proto.RegisterType((*PauseResponse)(nil), "service.PauseResponse")
	proto.RegisterType((*ResumeRequest)(nil), "service.ResumeRequest")
	proto.RegisterType((*ResumeResponse)(nil), "service.ResumeResponse")
	proto.RegisterType((*TerminateRequest)(nil), "service.TerminateRequest")
	proto.RegisterType((*TerminateResponse)(nil), "service.TerminateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionsClient is the client API for Sessions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionsClient interface {
	Create(ctx context.Context, opts ...grpc.CallOption) (Sessions_CreateClient, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Pause(ctx context.Context, opts ...grpc.CallOption) (Sessions_PauseClient, error)
	Resume(ctx context.Context, opts ...grpc.CallOption) (Sessions_ResumeClient, error)
	Terminate(ctx context.Context, opts ...grpc.CallOption) (Sessions_TerminateClient, error)
}

type sessionsClient struct {
	cc *grpc.ClientConn
}

func NewSessionsClient(cc *grpc.ClientConn) SessionsClient {
	return &sessionsClient{cc}
}

func (c *sessionsClient) Create(ctx context.Context, opts ...grpc.CallOption) (Sessions_CreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sessions_serviceDesc.Streams[0], "/service.Sessions/Create", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionsCreateClient{stream}
	return x, nil
}

type Sessions_CreateClient interface {
	Send(*CreateRequest) error
	Recv() (*CreateResponse, error)
	grpc.ClientStream
}

type sessionsCreateClient struct {
	grpc.ClientStream
}

func (x *sessionsCreateClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionsCreateClient) Recv() (*CreateResponse, error) {
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/service.Sessions/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionsClient) Pause(ctx context.Context, opts ...grpc.CallOption) (Sessions_PauseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sessions_serviceDesc.Streams[1], "/service.Sessions/Pause", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionsPauseClient{stream}
	return x, nil
}

type Sessions_PauseClient interface {
	Send(*PauseRequest) error
	Recv() (*PauseResponse, error)
	grpc.ClientStream
}

type sessionsPauseClient struct {
	grpc.ClientStream
}

func (x *sessionsPauseClient) Send(m *PauseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionsPauseClient) Recv() (*PauseResponse, error) {
	m := new(PauseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionsClient) Resume(ctx context.Context, opts ...grpc.CallOption) (Sessions_ResumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sessions_serviceDesc.Streams[2], "/service.Sessions/Resume", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionsResumeClient{stream}
	return x, nil
}

type Sessions_ResumeClient interface {
	Send(*ResumeRequest) error
	Recv() (*ResumeResponse, error)
	grpc.ClientStream
}

type sessionsResumeClient struct {
	grpc.ClientStream
}

func (x *sessionsResumeClient) Send(m *ResumeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionsResumeClient) Recv() (*ResumeResponse, error) {
	m := new(ResumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionsClient) Terminate(ctx context.Context, opts ...grpc.CallOption) (Sessions_TerminateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sessions_serviceDesc.Streams[3], "/service.Sessions/Terminate", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionsTerminateClient{stream}
	return x, nil
}

type Sessions_TerminateClient interface {
	Send(*TerminateRequest) error
	Recv() (*TerminateResponse, error)
	grpc.ClientStream
}

type sessionsTerminateClient struct {
	grpc.ClientStream
}

func (x *sessionsTerminateClient) Send(m *TerminateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionsTerminateClient) Recv() (*TerminateResponse, error) {
	m := new(TerminateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SessionsServer is the server API for Sessions service.
type SessionsServer interface {
	Create(Sessions_CreateServer) error
	List(context.Context, *ListRequest) (*ListResponse, error)
	Pause(Sessions_PauseServer) error
	Resume(Sessions_ResumeServer) error
	Terminate(Sessions_TerminateServer) error
}

func RegisterSessionsServer(s *grpc.Server, srv SessionsServer) {
	s.RegisterService(&_Sessions_serviceDesc, srv)
}

func _Sessions_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionsServer).Create(&sessionsCreateServer{stream})
}

type Sessions_CreateServer interface {
	Send(*CreateResponse) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type sessionsCreateServer struct {
	grpc.ServerStream
}

func (x *sessionsCreateServer) Send(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionsCreateServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sessions_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Sessions/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sessions_Pause_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionsServer).Pause(&sessionsPauseServer{stream})
}

type Sessions_PauseServer interface {
	Send(*PauseResponse) error
	Recv() (*PauseRequest, error)
	grpc.ServerStream
}

type sessionsPauseServer struct {
	grpc.ServerStream
}

func (x *sessionsPauseServer) Send(m *PauseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionsPauseServer) Recv() (*PauseRequest, error) {
	m := new(PauseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sessions_Resume_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionsServer).Resume(&sessionsResumeServer{stream})
}

type Sessions_ResumeServer interface {
	Send(*ResumeResponse) error
	Recv() (*ResumeRequest, error)
	grpc.ServerStream
}

type sessionsResumeServer struct {
	grpc.ServerStream
}

func (x *sessionsResumeServer) Send(m *ResumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionsResumeServer) Recv() (*ResumeRequest, error) {
	m := new(ResumeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sessions_Terminate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionsServer).Terminate(&sessionsTerminateServer{stream})
}

type Sessions_TerminateServer interface {
	Send(*TerminateResponse) error
	Recv() (*TerminateRequest, error)
	grpc.ServerStream
}

type sessionsTerminateServer struct {
	grpc.ServerStream
}

func (x *sessionsTerminateServer) Send(m *TerminateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionsTerminateServer) Recv() (*TerminateRequest, error) {
	m := new(TerminateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Sessions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Sessions",
	HandlerType: (*SessionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Sessions_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Create",
			Handler:       _Sessions_Create_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Pause",
			Handler:       _Sessions_Pause_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Resume",
			Handler:       _Sessions_Resume_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Terminate",
			Handler:       _Sessions_Terminate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "session/service/session.proto",
}

func init() {
	proto.RegisterFile("session/service/session.proto", fileDescriptor_session_7aaea59be0ec6c8e)
}

var fileDescriptor_session_7aaea59be0ec6c8e = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x69, 0x9a, 0xc6, 0x93, 0x3a, 0x82, 0x05, 0x82, 0x31, 0x50, 0x45, 0x3e, 0xa0, 0x70,
	0xa8, 0x0d, 0xe1, 0xe3, 0x80, 0x2a, 0x55, 0x6a, 0x2f, 0x20, 0xf5, 0x80, 0x36, 0x70, 0x41, 0x5c,
	0x1c, 0x77, 0x9a, 0x58, 0xc4, 0x1f, 0xec, 0xae, 0x23, 0xfe, 0x11, 0x3f, 0x86, 0x3f, 0x85, 0xbc,
	0x5e, 0xbb, 0xbb, 0x6e, 0xa4, 0xd2, 0xe3, 0xbc, 0x99, 0x9d, 0x79, 0xf3, 0xde, 0xd8, 0xf0, 0x82,
	0x23, 0xe7, 0x49, 0x9e, 0x85, 0x1c, 0xd9, 0x36, 0x89, 0x31, 0x54, 0x71, 0x50, 0xb0, 0x5c, 0xe4,
	0xe4, 0x40, 0xc1, 0xde, 0xb3, 0xa6, 0x2e, 0xce, 0xb3, 0xab, 0x64, 0x55, 0xb2, 0x48, 0xb4, 0x55,
	0xde, 0xc3, 0xb6, 0x89, 0x88, 0x04, 0x2a, 0xd0, 0x29, 0xd9, 0x26, 0x2c, 0xd9, 0xa6, 0x0e, 0xfd,
	0x3f, 0x16, 0x38, 0xe7, 0x0c, 0x23, 0x81, 0x14, 0x7f, 0x95, 0xc8, 0x05, 0x39, 0x82, 0xfd, 0x68,
	0x53, 0xac, 0x23, 0xd7, 0x9a, 0x5a, 0xb3, 0xd1, 0x7c, 0x18, 0x54, 0xc5, 0xdf, 0xe8, 0x05, 0xad,
	0x61, 0xf2, 0x1c, 0xfa, 0x4b, 0x14, 0x91, 0xdb, 0xeb, 0xa4, 0x25, 0x4a, 0x4e, 0xc0, 0x31, 0xa8,
	0xb8, 0x7b, 0xb2, 0x6c, 0x12, 0x34, 0x0b, 0x9c, 0xeb, 0x59, 0x6a, 0x16, 0x13, 0x0f, 0x86, 0x0c,
	0x79, 0x91, 0x67, 0x1c, 0xdd, 0xfe, 0xd4, 0x9a, 0xd9, 0xb4, 0x8d, 0xfd, 0x1f, 0x30, 0x6e, 0x88,
	0xd6, 0x08, 0x71, 0xe1, 0x40, 0x75, 0x95, 0x5c, 0x6d, 0xda, 0x84, 0x55, 0x26, 0x45, 0xce, 0xa3,
	0x15, 0x4a, 0x9a, 0x36, 0x6d, 0x42, 0x32, 0x81, 0x41, 0xc1, 0xf2, 0xb4, 0x10, 0x92, 0x98, 0x4d,
	0x55, 0xe4, 0x23, 0x8c, 0x2e, 0x12, 0x2e, 0x1a, 0x11, 0x02, 0x20, 0x05, 0xc3, 0x6d, 0x92, 0x97,
	0x7c, 0x51, 0x89, 0xf7, 0x39, 0xbb, 0xc4, 0xdf, 0x72, 0x4a, 0x9f, 0xee, 0xc8, 0x90, 0x97, 0x30,
	0xe6, 0x05, 0xc6, 0xc9, 0x55, 0x12, 0xcb, 0x4d, 0xb8, 0xdb, 0x9b, 0xee, 0xcd, 0x6c, 0xda, 0x41,
	0xfd, 0x4b, 0x38, 0xac, 0xc7, 0xa8, 0x15, 0x8e, 0x00, 0x78, 0xb7, 0xbf, 0x86, 0x90, 0x77, 0xe0,
	0xa8, 0x9d, 0xe4, 0xb0, 0xba, 0xed, 0x68, 0x3e, 0x6e, 0xe5, 0x94, 0x30, 0x35, 0x8b, 0xfc, 0x0f,
	0x70, 0xf8, 0x25, 0x2a, 0x79, 0x6b, 0xe9, 0x4d, 0x76, 0xd6, 0x4e, 0x76, 0xaf, 0xc0, 0x51, 0xef,
	0xae, 0x15, 0x6e, 0x74, 0xb4, 0x0c, 0x1d, 0xfd, 0x05, 0x38, 0x14, 0x79, 0x99, 0xde, 0x75, 0x86,
	0x61, 0x71, 0xaf, 0x63, 0xf1, 0x19, 0x8c, 0x9b, 0xa6, 0xb7, 0x11, 0xd0, 0x8c, 0xec, 0x19, 0x46,
	0x7e, 0x84, 0xfb, 0x5f, 0x91, 0xa5, 0x49, 0xa6, 0x9d, 0xf4, 0xff, 0xee, 0x7f, 0x0c, 0x0f, 0xb4,
	0xb7, 0xb7, 0x51, 0x98, 0xff, 0xed, 0xc1, 0x70, 0x51, 0x0b, 0xcf, 0xc9, 0x29, 0x0c, 0xea, 0xf3,
	0x24, 0xd5, 0xad, 0xcb, 0xaf, 0x33, 0x30, 0x3e, 0x2c, 0xef, 0xc9, 0x0d, 0x5c, 0xad, 0x7d, 0x6f,
	0x66, 0xbd, 0xb6, 0xc8, 0x7b, 0xe8, 0x57, 0xa7, 0x41, 0x1e, 0xb5, 0x65, 0xda, 0x41, 0x7a, 0x8f,
	0x3b, 0x68, 0xf3, 0x94, 0x9c, 0xc0, 0xbe, 0xf4, 0x8c, 0x5c, 0x57, 0xe8, 0xde, 0x7b, 0x93, 0x2e,
	0x6c, 0x0c, 0x3d, 0x85, 0x41, 0xad, 0xb8, 0xc6, 0xda, 0xf0, 0x55, 0x63, 0x6d, 0x5a, 0xa3, 0x1a,
	0x7c, 0x02, 0xbb, 0x95, 0x8c, 0x3c, 0x6d, 0x6b, 0xbb, 0x16, 0x78, 0xde, 0xae, 0x94, 0xde, 0xe9,
	0xec, 0xcd, 0xf7, 0x70, 0x95, 0x88, 0x75, 0xb9, 0x0c, 0xe2, 0x3c, 0x0d, 0xd7, 0xd1, 0x36, 0x8f,
	0x8f, 0x93, 0x3c, 0x4c, 0x4b, 0x11, 0xad, 0x30, 0x0b, 0x8b, 0x9f, 0xab, 0xb0, 0xf3, 0x53, 0x5c,
	0x0e, 0xe4, 0x3f, 0xec, 0xed, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x07, 0xf2, 0x38, 0x2e,
	0x05, 0x00, 0x00,
}
