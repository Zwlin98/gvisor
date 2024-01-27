// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: pkg/sentry/seccheck/points/common.proto

package points_go_proto

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

type MessageType int32

const (
	MessageType_MESSAGE_UNKNOWN                   MessageType = 0
	MessageType_MESSAGE_CONTAINER_START           MessageType = 1
	MessageType_MESSAGE_SENTRY_CLONE              MessageType = 2
	MessageType_MESSAGE_SENTRY_EXEC               MessageType = 3
	MessageType_MESSAGE_SENTRY_EXIT_NOTIFY_PARENT MessageType = 4
	MessageType_MESSAGE_SENTRY_TASK_EXIT          MessageType = 5
	MessageType_MESSAGE_SYSCALL_RAW               MessageType = 6
	MessageType_MESSAGE_SYSCALL_OPEN              MessageType = 7
	MessageType_MESSAGE_SYSCALL_CLOSE             MessageType = 8
	MessageType_MESSAGE_SYSCALL_READ              MessageType = 9
	MessageType_MESSAGE_SYSCALL_CONNECT           MessageType = 10
	MessageType_MESSAGE_SYSCALL_EXECVE            MessageType = 11
	MessageType_MESSAGE_SYSCALL_SOCKET            MessageType = 12
	MessageType_MESSAGE_SYSCALL_CHDIR             MessageType = 13
	MessageType_MESSAGE_SYSCALL_SETID             MessageType = 14
	MessageType_MESSAGE_SYSCALL_SETRESID          MessageType = 15
	MessageType_MESSAGE_SYSCALL_PRLIMIT64         MessageType = 16
	MessageType_MESSAGE_SYSCALL_PIPE              MessageType = 17
	MessageType_MESSAGE_SYSCALL_FCNTL             MessageType = 18
	MessageType_MESSAGE_SYSCALL_DUP               MessageType = 19
	MessageType_MESSAGE_SYSCALL_SIGNALFD          MessageType = 20
	MessageType_MESSAGE_SYSCALL_CHROOT            MessageType = 21
	MessageType_MESSAGE_SYSCALL_EVENTFD           MessageType = 22
	MessageType_MESSAGE_SYSCALL_CLONE             MessageType = 23
	MessageType_MESSAGE_SYSCALL_BIND              MessageType = 24
	MessageType_MESSAGE_SYSCALL_ACCEPT            MessageType = 25
	MessageType_MESSAGE_SYSCALL_TIMERFD_CREATE    MessageType = 26
	MessageType_MESSAGE_SYSCALL_TIMERFD_SETTIME   MessageType = 27
	MessageType_MESSAGE_SYSCALL_TIMERFD_GETTIME   MessageType = 28
	MessageType_MESSAGE_SYSCALL_FORK              MessageType = 29
	MessageType_MESSAGE_SYSCALL_INOTIFY_INIT      MessageType = 30
	MessageType_MESSAGE_SYSCALL_INOTIFY_ADD_WATCH MessageType = 31
	MessageType_MESSAGE_SYSCALL_INOTIFY_RM_WATCH  MessageType = 32
	MessageType_MESSAGE_SYSCALL_SOCKETPAIR        MessageType = 33
	MessageType_MESSAGE_SYSCALL_WRITE             MessageType = 34
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0:  "MESSAGE_UNKNOWN",
		1:  "MESSAGE_CONTAINER_START",
		2:  "MESSAGE_SENTRY_CLONE",
		3:  "MESSAGE_SENTRY_EXEC",
		4:  "MESSAGE_SENTRY_EXIT_NOTIFY_PARENT",
		5:  "MESSAGE_SENTRY_TASK_EXIT",
		6:  "MESSAGE_SYSCALL_RAW",
		7:  "MESSAGE_SYSCALL_OPEN",
		8:  "MESSAGE_SYSCALL_CLOSE",
		9:  "MESSAGE_SYSCALL_READ",
		10: "MESSAGE_SYSCALL_CONNECT",
		11: "MESSAGE_SYSCALL_EXECVE",
		12: "MESSAGE_SYSCALL_SOCKET",
		13: "MESSAGE_SYSCALL_CHDIR",
		14: "MESSAGE_SYSCALL_SETID",
		15: "MESSAGE_SYSCALL_SETRESID",
		16: "MESSAGE_SYSCALL_PRLIMIT64",
		17: "MESSAGE_SYSCALL_PIPE",
		18: "MESSAGE_SYSCALL_FCNTL",
		19: "MESSAGE_SYSCALL_DUP",
		20: "MESSAGE_SYSCALL_SIGNALFD",
		21: "MESSAGE_SYSCALL_CHROOT",
		22: "MESSAGE_SYSCALL_EVENTFD",
		23: "MESSAGE_SYSCALL_CLONE",
		24: "MESSAGE_SYSCALL_BIND",
		25: "MESSAGE_SYSCALL_ACCEPT",
		26: "MESSAGE_SYSCALL_TIMERFD_CREATE",
		27: "MESSAGE_SYSCALL_TIMERFD_SETTIME",
		28: "MESSAGE_SYSCALL_TIMERFD_GETTIME",
		29: "MESSAGE_SYSCALL_FORK",
		30: "MESSAGE_SYSCALL_INOTIFY_INIT",
		31: "MESSAGE_SYSCALL_INOTIFY_ADD_WATCH",
		32: "MESSAGE_SYSCALL_INOTIFY_RM_WATCH",
		33: "MESSAGE_SYSCALL_SOCKETPAIR",
		34: "MESSAGE_SYSCALL_WRITE",
	}
	MessageType_value = map[string]int32{
		"MESSAGE_UNKNOWN":                   0,
		"MESSAGE_CONTAINER_START":           1,
		"MESSAGE_SENTRY_CLONE":              2,
		"MESSAGE_SENTRY_EXEC":               3,
		"MESSAGE_SENTRY_EXIT_NOTIFY_PARENT": 4,
		"MESSAGE_SENTRY_TASK_EXIT":          5,
		"MESSAGE_SYSCALL_RAW":               6,
		"MESSAGE_SYSCALL_OPEN":              7,
		"MESSAGE_SYSCALL_CLOSE":             8,
		"MESSAGE_SYSCALL_READ":              9,
		"MESSAGE_SYSCALL_CONNECT":           10,
		"MESSAGE_SYSCALL_EXECVE":            11,
		"MESSAGE_SYSCALL_SOCKET":            12,
		"MESSAGE_SYSCALL_CHDIR":             13,
		"MESSAGE_SYSCALL_SETID":             14,
		"MESSAGE_SYSCALL_SETRESID":          15,
		"MESSAGE_SYSCALL_PRLIMIT64":         16,
		"MESSAGE_SYSCALL_PIPE":              17,
		"MESSAGE_SYSCALL_FCNTL":             18,
		"MESSAGE_SYSCALL_DUP":               19,
		"MESSAGE_SYSCALL_SIGNALFD":          20,
		"MESSAGE_SYSCALL_CHROOT":            21,
		"MESSAGE_SYSCALL_EVENTFD":           22,
		"MESSAGE_SYSCALL_CLONE":             23,
		"MESSAGE_SYSCALL_BIND":              24,
		"MESSAGE_SYSCALL_ACCEPT":            25,
		"MESSAGE_SYSCALL_TIMERFD_CREATE":    26,
		"MESSAGE_SYSCALL_TIMERFD_SETTIME":   27,
		"MESSAGE_SYSCALL_TIMERFD_GETTIME":   28,
		"MESSAGE_SYSCALL_FORK":              29,
		"MESSAGE_SYSCALL_INOTIFY_INIT":      30,
		"MESSAGE_SYSCALL_INOTIFY_ADD_WATCH": 31,
		"MESSAGE_SYSCALL_INOTIFY_RM_WATCH":  32,
		"MESSAGE_SYSCALL_SOCKETPAIR":        33,
		"MESSAGE_SYSCALL_WRITE":             34,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_sentry_seccheck_points_common_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_pkg_sentry_seccheck_points_common_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_sentry_seccheck_points_common_proto_rawDescGZIP(), []int{0}
}

type Handshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Handshake) Reset() {
	*x = Handshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_seccheck_points_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake) ProtoMessage() {}

func (x *Handshake) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_seccheck_points_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake.ProtoReflect.Descriptor instead.
func (*Handshake) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_seccheck_points_common_proto_rawDescGZIP(), []int{0}
}

func (x *Handshake) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RealUid      uint32 `protobuf:"varint,1,opt,name=real_uid,json=realUid,proto3" json:"real_uid,omitempty"`
	EffectiveUid uint32 `protobuf:"varint,2,opt,name=effective_uid,json=effectiveUid,proto3" json:"effective_uid,omitempty"`
	SavedUid     uint32 `protobuf:"varint,3,opt,name=saved_uid,json=savedUid,proto3" json:"saved_uid,omitempty"`
	RealGid      uint32 `protobuf:"varint,4,opt,name=real_gid,json=realGid,proto3" json:"real_gid,omitempty"`
	EffectiveGid uint32 `protobuf:"varint,5,opt,name=effective_gid,json=effectiveGid,proto3" json:"effective_gid,omitempty"`
	SavedGid     uint32 `protobuf:"varint,6,opt,name=saved_gid,json=savedGid,proto3" json:"saved_gid,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_seccheck_points_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_seccheck_points_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_seccheck_points_common_proto_rawDescGZIP(), []int{1}
}

func (x *Credentials) GetRealUid() uint32 {
	if x != nil {
		return x.RealUid
	}
	return 0
}

func (x *Credentials) GetEffectiveUid() uint32 {
	if x != nil {
		return x.EffectiveUid
	}
	return 0
}

func (x *Credentials) GetSavedUid() uint32 {
	if x != nil {
		return x.SavedUid
	}
	return 0
}

func (x *Credentials) GetRealGid() uint32 {
	if x != nil {
		return x.RealGid
	}
	return 0
}

func (x *Credentials) GetEffectiveGid() uint32 {
	if x != nil {
		return x.EffectiveGid
	}
	return 0
}

func (x *Credentials) GetSavedGid() uint32 {
	if x != nil {
		return x.SavedGid
	}
	return 0
}

type ContextData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeNs                 int64        `protobuf:"varint,1,opt,name=time_ns,json=timeNs,proto3" json:"time_ns,omitempty"`
	ThreadId               int32        `protobuf:"varint,2,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	ThreadStartTimeNs      int64        `protobuf:"varint,3,opt,name=thread_start_time_ns,json=threadStartTimeNs,proto3" json:"thread_start_time_ns,omitempty"`
	ThreadGroupId          int32        `protobuf:"varint,4,opt,name=thread_group_id,json=threadGroupId,proto3" json:"thread_group_id,omitempty"`
	ThreadGroupStartTimeNs int64        `protobuf:"varint,5,opt,name=thread_group_start_time_ns,json=threadGroupStartTimeNs,proto3" json:"thread_group_start_time_ns,omitempty"`
	ContainerId            string       `protobuf:"bytes,6,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Credentials            *Credentials `protobuf:"bytes,7,opt,name=credentials,proto3" json:"credentials,omitempty"`
	Cwd                    string       `protobuf:"bytes,8,opt,name=cwd,proto3" json:"cwd,omitempty"`
	ProcessName            string       `protobuf:"bytes,9,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
}

func (x *ContextData) Reset() {
	*x = ContextData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_seccheck_points_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextData) ProtoMessage() {}

func (x *ContextData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_seccheck_points_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextData.ProtoReflect.Descriptor instead.
func (*ContextData) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_seccheck_points_common_proto_rawDescGZIP(), []int{2}
}

func (x *ContextData) GetTimeNs() int64 {
	if x != nil {
		return x.TimeNs
	}
	return 0
}

func (x *ContextData) GetThreadId() int32 {
	if x != nil {
		return x.ThreadId
	}
	return 0
}

func (x *ContextData) GetThreadStartTimeNs() int64 {
	if x != nil {
		return x.ThreadStartTimeNs
	}
	return 0
}

func (x *ContextData) GetThreadGroupId() int32 {
	if x != nil {
		return x.ThreadGroupId
	}
	return 0
}

func (x *ContextData) GetThreadGroupStartTimeNs() int64 {
	if x != nil {
		return x.ThreadGroupStartTimeNs
	}
	return 0
}

func (x *ContextData) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *ContextData) GetCredentials() *Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *ContextData) GetCwd() string {
	if x != nil {
		return x.Cwd
	}
	return ""
}

func (x *ContextData) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

var File_pkg_sentry_seccheck_points_common_proto protoreflect.FileDescriptor

var file_pkg_sentry_seccheck_points_common_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x73, 0x65, 0x63,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64,
	0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0xc7, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x55, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x73, 0x61, 0x76, 0x65, 0x64, 0x55, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x65, 0x61, 0x6c, 0x5f, 0x67, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x61, 0x6c, 0x47, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x47, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x73, 0x61, 0x76, 0x65, 0x64, 0x47, 0x69, 0x64, 0x22, 0xee, 0x02, 0x0a, 0x0b, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65,
	0x4e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12,
	0x2f, 0x0a, 0x14, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x1a, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x4e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x77, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x63, 0x77, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x8f, 0x08, 0x0a, 0x0b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41,
	0x49, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x43,
	0x4c, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x10, 0x03, 0x12,
	0x25, 0x0a, 0x21, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x52,
	0x59, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x50, 0x41,
	0x52, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x45, 0x58,
	0x49, 0x54, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x41, 0x57, 0x10, 0x06, 0x12, 0x18, 0x0a,
	0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59,
	0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x45, 0x58, 0x45,
	0x43, 0x56, 0x45, 0x10, 0x0b, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x4f, 0x43, 0x4b, 0x45, 0x54, 0x10,
	0x0c, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53,
	0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x48, 0x44, 0x49, 0x52, 0x10, 0x0d, 0x12, 0x19, 0x0a, 0x15,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x53, 0x45, 0x54, 0x49, 0x44, 0x10, 0x0e, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x45, 0x54, 0x52, 0x45,
	0x53, 0x49, 0x44, 0x10, 0x0f, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x50, 0x52, 0x4c, 0x49, 0x4d, 0x49, 0x54,
	0x36, 0x34, 0x10, 0x10, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x50, 0x49, 0x50, 0x45, 0x10, 0x11, 0x12, 0x19,
	0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x46, 0x43, 0x4e, 0x54, 0x4c, 0x10, 0x12, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x44, 0x55, 0x50,
	0x10, 0x13, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59,
	0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x4c, 0x46, 0x44, 0x10, 0x14,
	0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43,
	0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x48, 0x52, 0x4f, 0x4f, 0x54, 0x10, 0x15, 0x12, 0x1b, 0x0a, 0x17,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x46, 0x44, 0x10, 0x16, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x4c, 0x4f,
	0x4e, 0x45, 0x10, 0x17, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x10, 0x18, 0x12, 0x1a,
	0x0a, 0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x19, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x52, 0x46, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x1a, 0x12, 0x23,
	0x0a, 0x1f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x52, 0x46, 0x44, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4d,
	0x45, 0x10, 0x1b, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53,
	0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x52, 0x46, 0x44, 0x5f, 0x47,
	0x45, 0x54, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x1c, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x46, 0x4f, 0x52, 0x4b,
	0x10, 0x1d, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59,
	0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x49, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x49, 0x4e,
	0x49, 0x54, 0x10, 0x1e, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x49, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f,
	0x41, 0x44, 0x44, 0x5f, 0x57, 0x41, 0x54, 0x43, 0x48, 0x10, 0x1f, 0x12, 0x24, 0x0a, 0x20, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x49,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x52, 0x4d, 0x5f, 0x57, 0x41, 0x54, 0x43, 0x48, 0x10,
	0x20, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53,
	0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x4f, 0x43, 0x4b, 0x45, 0x54, 0x50, 0x41, 0x49, 0x52, 0x10,
	0x21, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x59, 0x53,
	0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x22, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_sentry_seccheck_points_common_proto_rawDescOnce sync.Once
	file_pkg_sentry_seccheck_points_common_proto_rawDescData = file_pkg_sentry_seccheck_points_common_proto_rawDesc
)

func file_pkg_sentry_seccheck_points_common_proto_rawDescGZIP() []byte {
	file_pkg_sentry_seccheck_points_common_proto_rawDescOnce.Do(func() {
		file_pkg_sentry_seccheck_points_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sentry_seccheck_points_common_proto_rawDescData)
	})
	return file_pkg_sentry_seccheck_points_common_proto_rawDescData
}

var file_pkg_sentry_seccheck_points_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_sentry_seccheck_points_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_sentry_seccheck_points_common_proto_goTypes = []interface{}{
	(MessageType)(0),    // 0: gvisor.common.MessageType
	(*Handshake)(nil),   // 1: gvisor.common.Handshake
	(*Credentials)(nil), // 2: gvisor.common.Credentials
	(*ContextData)(nil), // 3: gvisor.common.ContextData
}
var file_pkg_sentry_seccheck_points_common_proto_depIdxs = []int32{
	2, // 0: gvisor.common.ContextData.credentials:type_name -> gvisor.common.Credentials
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_sentry_seccheck_points_common_proto_init() }
func file_pkg_sentry_seccheck_points_common_proto_init() {
	if File_pkg_sentry_seccheck_points_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sentry_seccheck_points_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake); i {
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
		file_pkg_sentry_seccheck_points_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
		file_pkg_sentry_seccheck_points_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextData); i {
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
			RawDescriptor: file_pkg_sentry_seccheck_points_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_sentry_seccheck_points_common_proto_goTypes,
		DependencyIndexes: file_pkg_sentry_seccheck_points_common_proto_depIdxs,
		EnumInfos:         file_pkg_sentry_seccheck_points_common_proto_enumTypes,
		MessageInfos:      file_pkg_sentry_seccheck_points_common_proto_msgTypes,
	}.Build()
	File_pkg_sentry_seccheck_points_common_proto = out.File
	file_pkg_sentry_seccheck_points_common_proto_rawDesc = nil
	file_pkg_sentry_seccheck_points_common_proto_goTypes = nil
	file_pkg_sentry_seccheck_points_common_proto_depIdxs = nil
}
