// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agentlocalpb/agentlocal.proto

package agentlocalpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	inventorypb "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ServerInfo contains information about the PMM Server.
type ServerInfo struct {
	// PMM Server URL in a form https://HOST:PORT/.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// PMM Server's TLS certificate validation should be skipped if true.
	InsecureTls bool `protobuf:"varint,2,opt,name=insecure_tls,json=insecureTls,proto3" json:"insecure_tls,omitempty"`
	// True if pmm-agent is currently connected to the server.
	Connected bool `protobuf:"varint,3,opt,name=connected,proto3" json:"connected,omitempty"`
	// PMM Server version (if agent is connected).
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Ping time from pmm-agent to pmm-managed (if agent is connected).
	Latency *duration.Duration `protobuf:"bytes,5,opt,name=latency,proto3" json:"latency,omitempty"`
	// Clock drift from PMM Server (if agent is connected).
	ClockDrift           *duration.Duration `protobuf:"bytes,6,opt,name=clock_drift,json=clockDrift,proto3" json:"clock_drift,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dbe6cbaa67ca12c, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ServerInfo) GetInsecureTls() bool {
	if m != nil {
		return m.InsecureTls
	}
	return false
}

func (m *ServerInfo) GetConnected() bool {
	if m != nil {
		return m.Connected
	}
	return false
}

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetLatency() *duration.Duration {
	if m != nil {
		return m.Latency
	}
	return nil
}

func (m *ServerInfo) GetClockDrift() *duration.Duration {
	if m != nil {
		return m.ClockDrift
	}
	return nil
}

// AgentInfo contains information about Agent managed by pmm-agent.
type AgentInfo struct {
	AgentId              string                  `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	AgentType            inventorypb.AgentType   `protobuf:"varint,2,opt,name=agent_type,json=agentType,proto3,enum=inventory.AgentType" json:"agent_type,omitempty"`
	Status               inventorypb.AgentStatus `protobuf:"varint,3,opt,name=status,proto3,enum=inventory.AgentStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AgentInfo) Reset()         { *m = AgentInfo{} }
func (m *AgentInfo) String() string { return proto.CompactTextString(m) }
func (*AgentInfo) ProtoMessage()    {}
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dbe6cbaa67ca12c, []int{1}
}

func (m *AgentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentInfo.Unmarshal(m, b)
}
func (m *AgentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentInfo.Marshal(b, m, deterministic)
}
func (m *AgentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentInfo.Merge(m, src)
}
func (m *AgentInfo) XXX_Size() int {
	return xxx_messageInfo_AgentInfo.Size(m)
}
func (m *AgentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AgentInfo proto.InternalMessageInfo

func (m *AgentInfo) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *AgentInfo) GetAgentType() inventorypb.AgentType {
	if m != nil {
		return m.AgentType
	}
	return inventorypb.AgentType_AGENT_TYPE_INVALID
}

func (m *AgentInfo) GetStatus() inventorypb.AgentStatus {
	if m != nil {
		return m.Status
	}
	return inventorypb.AgentStatus_AGENT_STATUS_INVALID
}

type StatusRequest struct {
	// Returns network info (latency and clock_drift) if true.
	GetNetworkInfo       bool     `protobuf:"varint,1,opt,name=get_network_info,json=getNetworkInfo,proto3" json:"get_network_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dbe6cbaa67ca12c, []int{2}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

func (m *StatusRequest) GetGetNetworkInfo() bool {
	if m != nil {
		return m.GetNetworkInfo
	}
	return false
}

type StatusResponse struct {
	AgentId      string       `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	RunsOnNodeId string       `protobuf:"bytes,2,opt,name=runs_on_node_id,json=runsOnNodeId,proto3" json:"runs_on_node_id,omitempty"`
	ServerInfo   *ServerInfo  `protobuf:"bytes,3,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`
	AgentsInfo   []*AgentInfo `protobuf:"bytes,4,rep,name=agents_info,json=agentsInfo,proto3" json:"agents_info,omitempty"`
	// Config file path if pmm-agent was started with one.
	ConfigFilepath string `protobuf:"bytes,5,opt,name=config_filepath,json=configFilepath,proto3" json:"config_filepath,omitempty"`
	// PMM Agent version.
	AgentVersion         string   `protobuf:"bytes,6,opt,name=agent_version,json=agentVersion,proto3" json:"agent_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dbe6cbaa67ca12c, []int{3}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *StatusResponse) GetRunsOnNodeId() string {
	if m != nil {
		return m.RunsOnNodeId
	}
	return ""
}

func (m *StatusResponse) GetServerInfo() *ServerInfo {
	if m != nil {
		return m.ServerInfo
	}
	return nil
}

func (m *StatusResponse) GetAgentsInfo() []*AgentInfo {
	if m != nil {
		return m.AgentsInfo
	}
	return nil
}

func (m *StatusResponse) GetConfigFilepath() string {
	if m != nil {
		return m.ConfigFilepath
	}
	return ""
}

func (m *StatusResponse) GetAgentVersion() string {
	if m != nil {
		return m.AgentVersion
	}
	return ""
}

type ReloadRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReloadRequest) Reset()         { *m = ReloadRequest{} }
func (m *ReloadRequest) String() string { return proto.CompactTextString(m) }
func (*ReloadRequest) ProtoMessage()    {}
func (*ReloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dbe6cbaa67ca12c, []int{4}
}

func (m *ReloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadRequest.Unmarshal(m, b)
}
func (m *ReloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadRequest.Marshal(b, m, deterministic)
}
func (m *ReloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadRequest.Merge(m, src)
}
func (m *ReloadRequest) XXX_Size() int {
	return xxx_messageInfo_ReloadRequest.Size(m)
}
func (m *ReloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadRequest proto.InternalMessageInfo

// ReloadRequest may not be received by the client due to pmm-agent restart.
type ReloadResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReloadResponse) Reset()         { *m = ReloadResponse{} }
func (m *ReloadResponse) String() string { return proto.CompactTextString(m) }
func (*ReloadResponse) ProtoMessage()    {}
func (*ReloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dbe6cbaa67ca12c, []int{5}
}

func (m *ReloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadResponse.Unmarshal(m, b)
}
func (m *ReloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadResponse.Marshal(b, m, deterministic)
}
func (m *ReloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadResponse.Merge(m, src)
}
func (m *ReloadResponse) XXX_Size() int {
	return xxx_messageInfo_ReloadResponse.Size(m)
}
func (m *ReloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ServerInfo)(nil), "agentlocal.ServerInfo")
	proto.RegisterType((*AgentInfo)(nil), "agentlocal.AgentInfo")
	proto.RegisterType((*StatusRequest)(nil), "agentlocal.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "agentlocal.StatusResponse")
	proto.RegisterType((*ReloadRequest)(nil), "agentlocal.ReloadRequest")
	proto.RegisterType((*ReloadResponse)(nil), "agentlocal.ReloadResponse")
}

func init() {
	proto.RegisterFile("agentlocalpb/agentlocal.proto", fileDescriptor_7dbe6cbaa67ca12c)
}

var fileDescriptor_7dbe6cbaa67ca12c = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5f, 0x8b, 0xd3, 0x4e,
	0x14, 0x25, 0xed, 0xfe, 0xfa, 0xe7, 0xa6, 0x4d, 0xcb, 0xf0, 0x73, 0x49, 0xcb, 0x2a, 0x35, 0x22,
	0x96, 0x7d, 0x48, 0xa0, 0x05, 0xc5, 0xbe, 0x29, 0x8b, 0x50, 0x90, 0x15, 0xb2, 0x8b, 0xa0, 0x2f,
	0x21, 0x4d, 0x6e, 0x62, 0xd8, 0x30, 0x13, 0x67, 0x26, 0x95, 0xbe, 0xfa, 0xea, 0x93, 0xf8, 0x8d,
	0xfc, 0x0a, 0x7e, 0x05, 0x9f, 0xfc, 0x14, 0x92, 0x99, 0x64, 0xdb, 0xaa, 0xe8, 0x5b, 0xee, 0xb9,
	0xe7, 0xde, 0xb9, 0xe7, 0xf4, 0x50, 0xb8, 0x1b, 0xa6, 0x48, 0x65, 0xce, 0xa2, 0x30, 0x2f, 0x36,
	0xde, 0xbe, 0x70, 0x0b, 0xce, 0x24, 0x23, 0xb0, 0x47, 0xa6, 0x67, 0x29, 0x63, 0x69, 0x8e, 0x5e,
	0x58, 0x64, 0x5e, 0x48, 0x29, 0x93, 0xa1, 0xcc, 0x18, 0x15, 0x9a, 0x39, 0xbd, 0x57, 0x77, 0x55,
	0xb5, 0x29, 0x13, 0x2f, 0x2e, 0xb9, 0x22, 0xd4, 0x7d, 0x3b, 0xa3, 0x5b, 0xa4, 0x92, 0xf1, 0x5d,
	0xf3, 0x4e, 0x3d, 0xe9, 0xfc, 0x30, 0x00, 0xae, 0x90, 0x6f, 0x91, 0xaf, 0x69, 0xc2, 0xc8, 0x18,
	0xda, 0x25, 0xcf, 0x6d, 0x63, 0x66, 0xcc, 0xfb, 0x7e, 0xf5, 0x49, 0xee, 0xc3, 0x20, 0xa3, 0x02,
	0xa3, 0x92, 0x63, 0x20, 0x73, 0x61, 0xb7, 0x66, 0xc6, 0xbc, 0xe7, 0x9b, 0x0d, 0x76, 0x9d, 0x0b,
	0x72, 0x06, 0xfd, 0x88, 0x51, 0x8a, 0x91, 0xc4, 0xd8, 0x6e, 0xab, 0xfe, 0x1e, 0x20, 0x36, 0x74,
	0xb7, 0xc8, 0x45, 0xc6, 0xa8, 0x7d, 0xa2, 0xd6, 0x36, 0x25, 0x59, 0x42, 0x37, 0x0f, 0x25, 0xd2,
	0x68, 0x67, 0xff, 0x37, 0x33, 0xe6, 0xe6, 0x62, 0xe2, 0x6a, 0x1d, 0x6e, 0xa3, 0xc3, 0xbd, 0xa8,
	0x75, 0xf8, 0x0d, 0x93, 0xac, 0xc0, 0x8c, 0x72, 0x16, 0xdd, 0x04, 0x31, 0xcf, 0x12, 0x69, 0x77,
	0xfe, 0x35, 0x08, 0x8a, 0x7d, 0x51, 0x91, 0x9d, 0x4f, 0x06, 0xf4, 0x9f, 0x55, 0xea, 0x95, 0xd6,
	0x09, 0xf4, 0x94, 0x15, 0x41, 0x16, 0xd7, 0x82, 0xbb, 0xaa, 0x5e, 0xc7, 0x64, 0x09, 0xda, 0xfb,
	0x40, 0xee, 0x0a, 0x54, 0x92, 0xad, 0xc5, 0xff, 0xee, 0xad, 0x89, 0xae, 0x5a, 0x72, 0xbd, 0x2b,
	0xd0, 0xef, 0x87, 0xcd, 0x27, 0x71, 0xa1, 0x23, 0x64, 0x28, 0x4b, 0xa1, 0x3c, 0xb0, 0x16, 0xa7,
	0xbf, 0x0e, 0x5c, 0xa9, 0xae, 0x5f, 0xb3, 0x9c, 0xa7, 0x30, 0xac, 0x11, 0x7c, 0x5f, 0xa2, 0x90,
	0x64, 0x0e, 0xe3, 0x14, 0x65, 0x40, 0x51, 0x7e, 0x60, 0xfc, 0x26, 0xc8, 0x68, 0xc2, 0xd4, 0x61,
	0x3d, 0xdf, 0x4a, 0x51, 0x5e, 0x6a, 0xb8, 0x3a, 0xdd, 0xf9, 0xdc, 0x02, 0xab, 0x99, 0x15, 0x05,
	0xa3, 0x02, 0xff, 0xa6, 0xe6, 0x21, 0x8c, 0x78, 0x49, 0x45, 0xc0, 0x68, 0x40, 0x59, 0x8c, 0x15,
	0xa3, 0xa5, 0x18, 0x83, 0x0a, 0x7e, 0x45, 0x2f, 0x59, 0x8c, 0xeb, 0x98, 0x3c, 0x01, 0x53, 0xa8,
	0x24, 0xe8, 0x97, 0xdb, 0xca, 0xd9, 0x53, 0xf7, 0x20, 0x96, 0xfb, 0xa0, 0xf8, 0x20, 0xf6, 0xa1,
	0x79, 0x0c, 0xa6, 0xce, 0x94, 0x1e, 0x3c, 0x99, 0xb5, 0xe7, 0xe6, 0xe2, 0xce, 0xe1, 0xe0, 0xad,
	0xe9, 0xbe, 0xf6, 0x55, 0xa8, 0xb9, 0x47, 0x30, 0x8a, 0x18, 0x4d, 0xb2, 0x34, 0x48, 0xb2, 0x1c,
	0x8b, 0x50, 0xbe, 0x53, 0x39, 0xe8, 0xfb, 0x96, 0x86, 0x5f, 0xd4, 0x28, 0x79, 0x00, 0x43, 0xad,
	0xad, 0x09, 0x52, 0x47, 0x9f, 0xaf, 0xc0, 0xd7, 0x1a, 0x73, 0x46, 0x30, 0xf4, 0x31, 0x67, 0x61,
	0x5c, 0xdb, 0xe9, 0x8c, 0xc1, 0x6a, 0x00, 0xed, 0xd1, 0xe2, 0xab, 0x01, 0xa0, 0x4e, 0x79, 0x59,
	0x5d, 0x45, 0xde, 0x40, 0x47, 0x9b, 0x48, 0x26, 0x47, 0x2a, 0x0f, 0x7f, 0x94, 0xe9, 0xf4, 0x4f,
	0x2d, 0xbd, 0xcf, 0xb1, 0x3f, 0x7e, 0xfb, 0xfe, 0xa5, 0x45, 0x9c, 0xa1, 0xa7, 0xda, 0x9e, 0x6e,
	0xaf, 0x8c, 0xf3, 0x6a, 0xb5, 0x7e, 0xfb, 0x78, 0xf5, 0xd1, 0x81, 0xc7, 0xab, 0x8f, 0x4f, 0xfd,
	0x6d, 0xb5, 0x6e, 0xaf, 0x8c, 0xf3, 0xe7, 0xd6, 0xdb, 0xc1, 0xe1, 0xdf, 0xc6, 0xa6, 0xa3, 0x32,
	0xbf, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0x87, 0x37, 0x25, 0x80, 0x4d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentLocalClient is the client API for AgentLocal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentLocalClient interface {
	// Status returns current pmm-agent status.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// Reload reloads pmm-agent configuration.
	Reload(ctx context.Context, in *ReloadRequest, opts ...grpc.CallOption) (*ReloadResponse, error)
}

type agentLocalClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentLocalClient(cc grpc.ClientConnInterface) AgentLocalClient {
	return &agentLocalClient{cc}
}

func (c *agentLocalClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agentlocal.AgentLocal/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentLocalClient) Reload(ctx context.Context, in *ReloadRequest, opts ...grpc.CallOption) (*ReloadResponse, error) {
	out := new(ReloadResponse)
	err := c.cc.Invoke(ctx, "/agentlocal.AgentLocal/Reload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentLocalServer is the server API for AgentLocal service.
type AgentLocalServer interface {
	// Status returns current pmm-agent status.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// Reload reloads pmm-agent configuration.
	Reload(context.Context, *ReloadRequest) (*ReloadResponse, error)
}

// UnimplementedAgentLocalServer can be embedded to have forward compatible implementations.
type UnimplementedAgentLocalServer struct {
}

func (*UnimplementedAgentLocalServer) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedAgentLocalServer) Reload(ctx context.Context, req *ReloadRequest) (*ReloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reload not implemented")
}

func RegisterAgentLocalServer(s *grpc.Server, srv AgentLocalServer) {
	s.RegisterService(&_AgentLocal_serviceDesc, srv)
}

func _AgentLocal_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentLocalServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentlocal.AgentLocal/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentLocalServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentLocal_Reload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentLocalServer).Reload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentlocal.AgentLocal/Reload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentLocalServer).Reload(ctx, req.(*ReloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentLocal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agentlocal.AgentLocal",
	HandlerType: (*AgentLocalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AgentLocal_Status_Handler,
		},
		{
			MethodName: "Reload",
			Handler:    _AgentLocal_Reload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agentlocalpb/agentlocal.proto",
}
