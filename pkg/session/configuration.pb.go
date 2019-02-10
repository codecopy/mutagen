// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session/configuration.proto

package session // import "github.com/havoc-io/mutagen/pkg/session"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import filesystem "github.com/havoc-io/mutagen/pkg/filesystem"
import sync "github.com/havoc-io/mutagen/pkg/sync"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Configuration encodes session configuration parameters. It is used for create
// commands to specify configuration options, for loading global configuration
// options, and for storing a merged configuration inside sessions.
type Configuration struct {
	// SynchronizationMode specifies the synchronization mode that should be
	// used in synchronization.
	SynchronizationMode sync.SynchronizationMode `protobuf:"varint,11,opt,name=synchronizationMode,proto3,enum=sync.SynchronizationMode" json:"synchronizationMode,omitempty"`
	// MaximumEntryCount specifies the maximum number of filesystem entries that
	// endpoints will tolerate managing. A zero value indicates no limit.
	MaximumEntryCount uint64 `protobuf:"varint,12,opt,name=maximumEntryCount,proto3" json:"maximumEntryCount,omitempty"`
	// MaximumStagingFileSize is the maximum (individual) file size that
	// endpoints will stage. A zero value indicates no limit.
	MaximumStagingFileSize uint64 `protobuf:"varint,13,opt,name=maximumStagingFileSize,proto3" json:"maximumStagingFileSize,omitempty"`
	// SymlinkMode specifies the symlink mode that should be used in
	// synchronization.
	SymlinkMode sync.SymlinkMode `protobuf:"varint,1,opt,name=symlinkMode,proto3,enum=sync.SymlinkMode" json:"symlinkMode,omitempty"`
	// WatchMode specifies the filesystem watching mode.
	WatchMode filesystem.WatchMode `protobuf:"varint,21,opt,name=watchMode,proto3,enum=filesystem.WatchMode" json:"watchMode,omitempty"`
	// WatchPollingInterval specifies the interval (in seconds) for poll-based
	// file monitoring. A value of 0 specifies that the default interval should
	// be used.
	WatchPollingInterval uint32 `protobuf:"varint,22,opt,name=watchPollingInterval,proto3" json:"watchPollingInterval,omitempty"`
	// DefaultIgnores specifies the ignore patterns brought in from the global
	// configuration.
	DefaultIgnores []string `protobuf:"bytes,31,rep,name=defaultIgnores,proto3" json:"defaultIgnores,omitempty"`
	// Ignores specifies the ignore patterns brought in from the create request.
	Ignores []string `protobuf:"bytes,32,rep,name=ignores,proto3" json:"ignores,omitempty"`
	// IgnoreVCSMode specifies the VCS ignore mode that should be used in
	// synchronization.
	IgnoreVCSMode sync.IgnoreVCSMode `protobuf:"varint,33,opt,name=ignoreVCSMode,proto3,enum=sync.IgnoreVCSMode" json:"ignoreVCSMode,omitempty"`
	// PermissionDefaultFileMode specifies the default permission mode to use
	// for new files in "portable" permission propagation mode, with
	// endpoint-specific specifications taking priority.
	PermissionDefaultFileMode uint32 `protobuf:"varint,65,opt,name=permissionDefaultFileMode,proto3" json:"permissionDefaultFileMode,omitempty"`
	// PermissionDefaultFileModeAlpha specifies the default permission mode to
	// use for new files on alpha in "portable" permission propagation mode,
	// taking priority over PermissionDefaultFileMode on alpha if specified.
	PermissionDefaultFileModeAlpha uint32 `protobuf:"varint,66,opt,name=permissionDefaultFileModeAlpha,proto3" json:"permissionDefaultFileModeAlpha,omitempty"`
	// PermissionDefaultFileModeBeta specifies the default permission mode to
	// use for new files on beta in "portable" permission propagation mode,
	// taking priority over PermissionDefaultFileMode on beta if specified.
	PermissionDefaultFileModeBeta uint32 `protobuf:"varint,67,opt,name=permissionDefaultFileModeBeta,proto3" json:"permissionDefaultFileModeBeta,omitempty"`
	// PermissionDefaultDirectoryMode specifies the default permission mode to
	// use for new directories in "portable" permission propagation mode, with
	// endpoint-specific specifications taking priority.
	PermissionDefaultDirectoryMode uint32 `protobuf:"varint,68,opt,name=permissionDefaultDirectoryMode,proto3" json:"permissionDefaultDirectoryMode,omitempty"`
	// PermissionDefaultDirectoryModeAlpha specifies the default permission mode
	// to use for new directories on alpha in "portable" permission propagation
	// mode, taking priority over PermissionDefaultDirectoryMode on alpha if
	// specified.
	PermissionDefaultDirectoryModeAlpha uint32 `protobuf:"varint,69,opt,name=permissionDefaultDirectoryModeAlpha,proto3" json:"permissionDefaultDirectoryModeAlpha,omitempty"`
	// PermissionDefaultDirectoryModeBeta specifies the default permission mode
	// to use for new directories on beta in "portable" permission propagation
	// mode, taking priority over PermissionDefaultDirectoryMode on beta if
	// specified.
	PermissionDefaultDirectoryModeBeta uint32 `protobuf:"varint,70,opt,name=permissionDefaultDirectoryModeBeta,proto3" json:"permissionDefaultDirectoryModeBeta,omitempty"`
	// PermissionDefaultUser specifies the default user identifier to use when
	// setting ownership of new files and directories in "portable" permission
	// propagation mode, with endpoint-specific specifications taking priority.
	PermissionDefaultUser string `protobuf:"bytes,71,opt,name=permissionDefaultUser,proto3" json:"permissionDefaultUser,omitempty"`
	// PermissionDefaultUserAlpha specifies the default user identifier to use
	// when setting ownership of new files and directories on alpha in
	// "portable" permission propagation mode, taking priority over
	// PermissionDefaultUser on alpha if specified.
	PermissionDefaultUserAlpha string `protobuf:"bytes,72,opt,name=permissionDefaultUserAlpha,proto3" json:"permissionDefaultUserAlpha,omitempty"`
	// PermissionDefaultUserBeta specifies the default user identifier to use
	// when setting ownership of new files and directories on beta in "portable"
	// permission propagation mode, taking priority over PermissionDefaultUser
	// on beta if specified.
	PermissionDefaultUserBeta string `protobuf:"bytes,73,opt,name=permissionDefaultUserBeta,proto3" json:"permissionDefaultUserBeta,omitempty"`
	// PermissionDefaultGroup specifies the default group identifier to use when
	// setting ownership of new files and directories in "portable" permission
	// propagation mode, with endpoint-specific specifications taking priority.
	PermissionDefaultGroup string `protobuf:"bytes,74,opt,name=permissionDefaultGroup,proto3" json:"permissionDefaultGroup,omitempty"`
	// PermissionDefaultGroupAlpha specifies the default group identifier to use
	// when setting ownership of new files and directories on alpha in
	// "portable" permission propagation mode, taking priority over
	// PermissionDefaultGroup on alpha if specified.
	PermissionDefaultGroupAlpha string `protobuf:"bytes,75,opt,name=permissionDefaultGroupAlpha,proto3" json:"permissionDefaultGroupAlpha,omitempty"`
	// PermissionDefaultGroupBeta specifies the default group identifier to use
	// when setting ownership of new files and directories on beta in "portable"
	// permission propagation mode, taking priority over PermissionDefaultGroup
	// on beta if specified.
	PermissionDefaultGroupBeta string   `protobuf:"bytes,76,opt,name=permissionDefaultGroupBeta,proto3" json:"permissionDefaultGroupBeta,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_20f63b98f7ba3a82, []int{0}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (dst *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(dst, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetSynchronizationMode() sync.SynchronizationMode {
	if m != nil {
		return m.SynchronizationMode
	}
	return sync.SynchronizationMode_SynchronizationModeDefault
}

func (m *Configuration) GetMaximumEntryCount() uint64 {
	if m != nil {
		return m.MaximumEntryCount
	}
	return 0
}

func (m *Configuration) GetMaximumStagingFileSize() uint64 {
	if m != nil {
		return m.MaximumStagingFileSize
	}
	return 0
}

func (m *Configuration) GetSymlinkMode() sync.SymlinkMode {
	if m != nil {
		return m.SymlinkMode
	}
	return sync.SymlinkMode_SymlinkDefault
}

func (m *Configuration) GetWatchMode() filesystem.WatchMode {
	if m != nil {
		return m.WatchMode
	}
	return filesystem.WatchMode_WatchDefault
}

func (m *Configuration) GetWatchPollingInterval() uint32 {
	if m != nil {
		return m.WatchPollingInterval
	}
	return 0
}

func (m *Configuration) GetDefaultIgnores() []string {
	if m != nil {
		return m.DefaultIgnores
	}
	return nil
}

func (m *Configuration) GetIgnores() []string {
	if m != nil {
		return m.Ignores
	}
	return nil
}

func (m *Configuration) GetIgnoreVCSMode() sync.IgnoreVCSMode {
	if m != nil {
		return m.IgnoreVCSMode
	}
	return sync.IgnoreVCSMode_IgnoreVCSDefault
}

func (m *Configuration) GetPermissionDefaultFileMode() uint32 {
	if m != nil {
		return m.PermissionDefaultFileMode
	}
	return 0
}

func (m *Configuration) GetPermissionDefaultFileModeAlpha() uint32 {
	if m != nil {
		return m.PermissionDefaultFileModeAlpha
	}
	return 0
}

func (m *Configuration) GetPermissionDefaultFileModeBeta() uint32 {
	if m != nil {
		return m.PermissionDefaultFileModeBeta
	}
	return 0
}

func (m *Configuration) GetPermissionDefaultDirectoryMode() uint32 {
	if m != nil {
		return m.PermissionDefaultDirectoryMode
	}
	return 0
}

func (m *Configuration) GetPermissionDefaultDirectoryModeAlpha() uint32 {
	if m != nil {
		return m.PermissionDefaultDirectoryModeAlpha
	}
	return 0
}

func (m *Configuration) GetPermissionDefaultDirectoryModeBeta() uint32 {
	if m != nil {
		return m.PermissionDefaultDirectoryModeBeta
	}
	return 0
}

func (m *Configuration) GetPermissionDefaultUser() string {
	if m != nil {
		return m.PermissionDefaultUser
	}
	return ""
}

func (m *Configuration) GetPermissionDefaultUserAlpha() string {
	if m != nil {
		return m.PermissionDefaultUserAlpha
	}
	return ""
}

func (m *Configuration) GetPermissionDefaultUserBeta() string {
	if m != nil {
		return m.PermissionDefaultUserBeta
	}
	return ""
}

func (m *Configuration) GetPermissionDefaultGroup() string {
	if m != nil {
		return m.PermissionDefaultGroup
	}
	return ""
}

func (m *Configuration) GetPermissionDefaultGroupAlpha() string {
	if m != nil {
		return m.PermissionDefaultGroupAlpha
	}
	return ""
}

func (m *Configuration) GetPermissionDefaultGroupBeta() string {
	if m != nil {
		return m.PermissionDefaultGroupBeta
	}
	return ""
}

func init() {
	proto.RegisterType((*Configuration)(nil), "session.Configuration")
}

func init() {
	proto.RegisterFile("session/configuration.proto", fileDescriptor_configuration_20f63b98f7ba3a82)
}

var fileDescriptor_configuration_20f63b98f7ba3a82 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xf1, 0x6b, 0xd3, 0x40,
	0x14, 0xc7, 0x29, 0x8a, 0xa3, 0x6f, 0x76, 0xd2, 0x9b, 0x2d, 0xb7, 0x0e, 0x35, 0x4e, 0xd0, 0x0a,
	0x9a, 0xc0, 0x2a, 0x82, 0x20, 0xe2, 0xda, 0xae, 0xb3, 0x4e, 0x65, 0xa4, 0xa8, 0xe0, 0x6f, 0xb7,
	0xf4, 0x9a, 0x1e, 0x4b, 0xee, 0xca, 0xe5, 0x32, 0xcd, 0x7e, 0xf3, 0x3f, 0x97, 0xbc, 0xa4, 0x5b,
	0x67, 0x93, 0x6c, 0xbf, 0x35, 0xef, 0xf3, 0x79, 0x2f, 0xdf, 0x77, 0x4d, 0x02, 0xbb, 0x11, 0x8f,
	0x22, 0xa1, 0xa4, 0xe3, 0x29, 0x39, 0x13, 0x7e, 0xac, 0x99, 0x11, 0x4a, 0xda, 0x0b, 0xad, 0x8c,
	0x22, 0x1b, 0x39, 0xec, 0xb4, 0x67, 0x22, 0xe0, 0x51, 0x12, 0x19, 0x1e, 0x3a, 0xbf, 0x99, 0xf1,
	0xe6, 0x99, 0xd0, 0x69, 0x46, 0x89, 0xf4, 0x1c, 0xe1, 0x4b, 0xa5, 0x79, 0x5e, 0x7a, 0x80, 0xa5,
	0x50, 0x4d, 0x97, 0x05, 0x82, 0x85, 0x28, 0x09, 0x03, 0x21, 0xcf, 0xb2, 0xda, 0xde, 0x5f, 0x80,
	0xc6, 0x60, 0xf5, 0x86, 0xe4, 0x18, 0xb6, 0x53, 0x6f, 0xae, 0x95, 0x14, 0x17, 0x58, 0xfa, 0xaa,
	0xa6, 0x9c, 0x6e, 0x5a, 0xb5, 0xee, 0xd6, 0xfe, 0x8e, 0x9d, 0x32, 0x7b, 0xb2, 0x2e, 0xb8, 0x45,
	0x5d, 0xe4, 0x15, 0x34, 0x43, 0xf6, 0x47, 0x84, 0x71, 0x78, 0x28, 0x8d, 0x4e, 0x06, 0x2a, 0x96,
	0x86, 0xde, 0xb7, 0x6a, 0xdd, 0xbb, 0xee, 0x3a, 0x20, 0x6f, 0xa1, 0x9d, 0x17, 0x27, 0x86, 0xf9,
	0x42, 0xfa, 0x23, 0x11, 0xf0, 0x89, 0xb8, 0xe0, 0xb4, 0x81, 0x2d, 0x25, 0x94, 0xf4, 0x60, 0x33,
	0xdf, 0x0a, 0xa3, 0xd6, 0x30, 0x6a, 0x73, 0x19, 0xf5, 0x12, 0xb8, 0xab, 0x16, 0xe9, 0x41, 0x1d,
	0x0f, 0x10, 0x5b, 0x5a, 0xd8, 0xd2, 0xb2, 0xaf, 0x4e, 0xd7, 0xfe, 0xb9, 0x84, 0xee, 0x95, 0x47,
	0xf6, 0xe1, 0x21, 0x5e, 0x9c, 0xa8, 0x20, 0x10, 0xd2, 0x1f, 0x4b, 0xc3, 0xf5, 0x39, 0x0b, 0x68,
	0xdb, 0xaa, 0x75, 0x1b, 0x6e, 0x21, 0x23, 0xcf, 0x61, 0x6b, 0xca, 0x67, 0x2c, 0x0e, 0xcc, 0x18,
	0xff, 0x9e, 0x88, 0x3e, 0xb1, 0xee, 0x74, 0xeb, 0xee, 0x7f, 0x55, 0x42, 0x61, 0x43, 0xe4, 0x82,
	0x85, 0xc2, 0xf2, 0x92, 0xbc, 0x83, 0x46, 0xf6, 0xf3, 0xc7, 0x60, 0x82, 0x71, 0x9f, 0x62, 0xdc,
	0xed, 0x6c, 0xc3, 0xf1, 0x2a, 0x72, 0xaf, 0x9b, 0xe4, 0x3d, 0xec, 0x2c, 0xb8, 0x0e, 0x05, 0x3e,
	0x3d, 0xc3, 0xec, 0x86, 0xe9, 0xb9, 0xe1, 0x98, 0x03, 0x4c, 0x5d, 0x2e, 0x90, 0x11, 0x3c, 0x2e,
	0x85, 0x07, 0xc1, 0x62, 0xce, 0x68, 0x1f, 0x47, 0xdc, 0x60, 0x91, 0x21, 0x3c, 0x2a, 0x35, 0xfa,
	0xdc, 0x30, 0x3a, 0xc0, 0x31, 0xd5, 0x52, 0x61, 0x9a, 0xa1, 0xd0, 0xdc, 0x33, 0x4a, 0x27, 0xb8,
	0xd0, 0xb0, 0x24, 0xcd, 0x35, 0x8b, 0x9c, 0xc0, 0xb3, 0x6a, 0x23, 0x5b, 0xed, 0x10, 0x87, 0xdd,
	0x46, 0x25, 0xdf, 0x60, 0xaf, 0x5a, 0xc3, 0x25, 0x47, 0x38, 0xf0, 0x16, 0x26, 0x79, 0x03, 0xad,
	0x35, 0xeb, 0x7b, 0xc4, 0x35, 0x3d, 0xb2, 0x6a, 0xdd, 0xba, 0x5b, 0x0c, 0xc9, 0x07, 0xe8, 0x14,
	0x82, 0x6c, 0x9d, 0x4f, 0xd8, 0x5a, 0x61, 0x14, 0x3e, 0x2b, 0x29, 0xc5, 0xf0, 0x63, 0x6c, 0x2f,
	0x17, 0xd2, 0x97, 0x77, 0x0d, 0x1e, 0x69, 0x15, 0x2f, 0xe8, 0x67, 0x6c, 0x2d, 0xa1, 0xe4, 0x23,
	0xec, 0x16, 0x93, 0x2c, 0xf6, 0x31, 0x36, 0x57, 0x29, 0x85, 0x7b, 0x23, 0xc6, 0xe0, 0x5f, 0x4a,
	0xf6, 0xbe, 0x34, 0xfa, 0x2f, 0x7f, 0xbd, 0xf0, 0x85, 0x99, 0xc7, 0xa7, 0xb6, 0xa7, 0x42, 0x67,
	0xce, 0xce, 0x95, 0xf7, 0x5a, 0x28, 0x27, 0x8c, 0x0d, 0xf3, 0xb9, 0x74, 0x16, 0x67, 0xbe, 0x93,
	0x7f, 0x7e, 0x4f, 0xef, 0xe1, 0x57, 0xb3, 0xf7, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x18, 0x54, 0xa5,
	0xb9, 0xad, 0x05, 0x00, 0x00,
}
