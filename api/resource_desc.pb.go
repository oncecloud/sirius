// Code generated by protoc-gen-go.
// source: resource_desc.proto
// DO NOT EDIT!

package sirius

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResourceDescriptor_ResourceState int32

const (
	ResourceDescriptor_RESOURCE_UNKNOWN ResourceDescriptor_ResourceState = 0
	ResourceDescriptor_RESOURCE_IDLE    ResourceDescriptor_ResourceState = 1
	ResourceDescriptor_RESOURCE_BUSY    ResourceDescriptor_ResourceState = 2
	ResourceDescriptor_RESOURCE_LOST    ResourceDescriptor_ResourceState = 3
)

var ResourceDescriptor_ResourceState_name = map[int32]string{
	0: "RESOURCE_UNKNOWN",
	1: "RESOURCE_IDLE",
	2: "RESOURCE_BUSY",
	3: "RESOURCE_LOST",
}
var ResourceDescriptor_ResourceState_value = map[string]int32{
	"RESOURCE_UNKNOWN": 0,
	"RESOURCE_IDLE":    1,
	"RESOURCE_BUSY":    2,
	"RESOURCE_LOST":    3,
}

func (x ResourceDescriptor_ResourceState) Enum() *ResourceDescriptor_ResourceState {
	p := new(ResourceDescriptor_ResourceState)
	*p = x
	return p
}
func (x ResourceDescriptor_ResourceState) String() string {
	return proto.EnumName(ResourceDescriptor_ResourceState_name, int32(x))
}
func (x *ResourceDescriptor_ResourceState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ResourceDescriptor_ResourceState_value, data, "ResourceDescriptor_ResourceState")
	if err != nil {
		return err
	}
	*x = ResourceDescriptor_ResourceState(value)
	return nil
}
func (ResourceDescriptor_ResourceState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 0}
}

type ResourceDescriptor_ResourceType int32

const (
	ResourceDescriptor_RESOURCE_PU          ResourceDescriptor_ResourceType = 0
	ResourceDescriptor_RESOURCE_CORE        ResourceDescriptor_ResourceType = 1
	ResourceDescriptor_RESOURCE_CACHE       ResourceDescriptor_ResourceType = 2
	ResourceDescriptor_RESOURCE_NIC         ResourceDescriptor_ResourceType = 3
	ResourceDescriptor_RESOURCE_DISK        ResourceDescriptor_ResourceType = 4
	ResourceDescriptor_RESOURCE_SSD         ResourceDescriptor_ResourceType = 5
	ResourceDescriptor_RESOURCE_MACHINE     ResourceDescriptor_ResourceType = 6
	ResourceDescriptor_RESOURCE_LOGICAL     ResourceDescriptor_ResourceType = 7
	ResourceDescriptor_RESOURCE_NUMA_NODE   ResourceDescriptor_ResourceType = 8
	ResourceDescriptor_RESOURCE_SOCKET      ResourceDescriptor_ResourceType = 9
	ResourceDescriptor_RESOURCE_COORDINATOR ResourceDescriptor_ResourceType = 10
)

var ResourceDescriptor_ResourceType_name = map[int32]string{
	0:  "RESOURCE_PU",
	1:  "RESOURCE_CORE",
	2:  "RESOURCE_CACHE",
	3:  "RESOURCE_NIC",
	4:  "RESOURCE_DISK",
	5:  "RESOURCE_SSD",
	6:  "RESOURCE_MACHINE",
	7:  "RESOURCE_LOGICAL",
	8:  "RESOURCE_NUMA_NODE",
	9:  "RESOURCE_SOCKET",
	10: "RESOURCE_COORDINATOR",
}
var ResourceDescriptor_ResourceType_value = map[string]int32{
	"RESOURCE_PU":          0,
	"RESOURCE_CORE":        1,
	"RESOURCE_CACHE":       2,
	"RESOURCE_NIC":         3,
	"RESOURCE_DISK":        4,
	"RESOURCE_SSD":         5,
	"RESOURCE_MACHINE":     6,
	"RESOURCE_LOGICAL":     7,
	"RESOURCE_NUMA_NODE":   8,
	"RESOURCE_SOCKET":      9,
	"RESOURCE_COORDINATOR": 10,
}

func (x ResourceDescriptor_ResourceType) Enum() *ResourceDescriptor_ResourceType {
	p := new(ResourceDescriptor_ResourceType)
	*p = x
	return p
}
func (x ResourceDescriptor_ResourceType) String() string {
	return proto.EnumName(ResourceDescriptor_ResourceType_name, int32(x))
}
func (x *ResourceDescriptor_ResourceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ResourceDescriptor_ResourceType_value, data, "ResourceDescriptor_ResourceType")
	if err != nil {
		return err
	}
	*x = ResourceDescriptor_ResourceType(value)
	return nil
}
func (ResourceDescriptor_ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 1}
}

type ResourceDescriptor struct {
	Uuid                *string                           `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	FriendlyName        *string                           `protobuf:"bytes,2,opt,name=friendly_name" json:"friendly_name,omitempty"`
	DescriptiveName     *string                           `protobuf:"bytes,3,opt,name=descriptive_name" json:"descriptive_name,omitempty"`
	State               *ResourceDescriptor_ResourceState `protobuf:"varint,4,opt,name=state,enum=sirius.ResourceDescriptor_ResourceState,def=0" json:"state,omitempty"`
	TaskCapacity        *uint64                           `protobuf:"varint,5,opt,name=task_capacity,def=0" json:"task_capacity,omitempty"`
	LastHeartbeat       *uint64                           `protobuf:"varint,6,opt,name=last_heartbeat" json:"last_heartbeat,omitempty"`
	Type                *ResourceDescriptor_ResourceType  `protobuf:"varint,7,opt,name=type,enum=sirius.ResourceDescriptor_ResourceType" json:"type,omitempty"`
	Schedulable         *bool                             `protobuf:"varint,8,opt,name=schedulable,def=0" json:"schedulable,omitempty"`
	CurrentRunningTasks []uint64                          `protobuf:"varint,9,rep,name=current_running_tasks" json:"current_running_tasks,omitempty"`
	// Stores the number of running tasks on the resources that
	// are below this node,
	NumRunningTasksBelow *uint64 `protobuf:"varint,10,opt,name=num_running_tasks_below,def=0" json:"num_running_tasks_below,omitempty"`
	NumSlotsBelow        *uint64 `protobuf:"varint,11,opt,name=num_slots_below,def=0" json:"num_slots_below,omitempty"`
	// Resource capacity and load tracking
	AvailableResources          *ResourceVector `protobuf:"bytes,12,opt,name=available_resources" json:"available_resources,omitempty"`
	ReservedResources           *ResourceVector `protobuf:"bytes,13,opt,name=reserved_resources" json:"reserved_resources,omitempty"`
	MinAvailableResourcesBelow  *ResourceVector `protobuf:"bytes,14,opt,name=min_available_resources_below" json:"min_available_resources_below,omitempty"`
	MaxAvailableResourcesBelow  *ResourceVector `protobuf:"bytes,15,opt,name=max_available_resources_below" json:"max_available_resources_below,omitempty"`
	MinUnreservedResourcesBelow *ResourceVector `protobuf:"bytes,16,opt,name=min_unreserved_resources_below" json:"min_unreserved_resources_below,omitempty"`
	MaxUnreservedResourcesBelow *ResourceVector `protobuf:"bytes,17,opt,name=max_unreserved_resources_below" json:"max_unreserved_resources_below,omitempty"`
	ResourceCapacity            *ResourceVector `protobuf:"bytes,18,opt,name=resource_capacity" json:"resource_capacity,omitempty"`
	XXX_unrecognized            []byte          `json:"-"`
}

func (m *ResourceDescriptor) Reset()                    { *m = ResourceDescriptor{} }
func (m *ResourceDescriptor) String() string            { return proto.CompactTextString(m) }
func (*ResourceDescriptor) ProtoMessage()               {}
func (*ResourceDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

const Default_ResourceDescriptor_State ResourceDescriptor_ResourceState = ResourceDescriptor_RESOURCE_UNKNOWN
const Default_ResourceDescriptor_TaskCapacity uint64 = 0
const Default_ResourceDescriptor_Schedulable bool = false
const Default_ResourceDescriptor_NumRunningTasksBelow uint64 = 0
const Default_ResourceDescriptor_NumSlotsBelow uint64 = 0

func (m *ResourceDescriptor) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *ResourceDescriptor) GetFriendlyName() string {
	if m != nil && m.FriendlyName != nil {
		return *m.FriendlyName
	}
	return ""
}

func (m *ResourceDescriptor) GetDescriptiveName() string {
	if m != nil && m.DescriptiveName != nil {
		return *m.DescriptiveName
	}
	return ""
}

func (m *ResourceDescriptor) GetState() ResourceDescriptor_ResourceState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return Default_ResourceDescriptor_State
}

func (m *ResourceDescriptor) GetTaskCapacity() uint64 {
	if m != nil && m.TaskCapacity != nil {
		return *m.TaskCapacity
	}
	return Default_ResourceDescriptor_TaskCapacity
}

func (m *ResourceDescriptor) GetLastHeartbeat() uint64 {
	if m != nil && m.LastHeartbeat != nil {
		return *m.LastHeartbeat
	}
	return 0
}

func (m *ResourceDescriptor) GetType() ResourceDescriptor_ResourceType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ResourceDescriptor_RESOURCE_PU
}

func (m *ResourceDescriptor) GetSchedulable() bool {
	if m != nil && m.Schedulable != nil {
		return *m.Schedulable
	}
	return Default_ResourceDescriptor_Schedulable
}

func (m *ResourceDescriptor) GetCurrentRunningTasks() []uint64 {
	if m != nil {
		return m.CurrentRunningTasks
	}
	return nil
}

func (m *ResourceDescriptor) GetNumRunningTasksBelow() uint64 {
	if m != nil && m.NumRunningTasksBelow != nil {
		return *m.NumRunningTasksBelow
	}
	return Default_ResourceDescriptor_NumRunningTasksBelow
}

func (m *ResourceDescriptor) GetNumSlotsBelow() uint64 {
	if m != nil && m.NumSlotsBelow != nil {
		return *m.NumSlotsBelow
	}
	return Default_ResourceDescriptor_NumSlotsBelow
}

func (m *ResourceDescriptor) GetAvailableResources() *ResourceVector {
	if m != nil {
		return m.AvailableResources
	}
	return nil
}

func (m *ResourceDescriptor) GetReservedResources() *ResourceVector {
	if m != nil {
		return m.ReservedResources
	}
	return nil
}

func (m *ResourceDescriptor) GetMinAvailableResourcesBelow() *ResourceVector {
	if m != nil {
		return m.MinAvailableResourcesBelow
	}
	return nil
}

func (m *ResourceDescriptor) GetMaxAvailableResourcesBelow() *ResourceVector {
	if m != nil {
		return m.MaxAvailableResourcesBelow
	}
	return nil
}

func (m *ResourceDescriptor) GetMinUnreservedResourcesBelow() *ResourceVector {
	if m != nil {
		return m.MinUnreservedResourcesBelow
	}
	return nil
}

func (m *ResourceDescriptor) GetMaxUnreservedResourcesBelow() *ResourceVector {
	if m != nil {
		return m.MaxUnreservedResourcesBelow
	}
	return nil
}

func (m *ResourceDescriptor) GetResourceCapacity() *ResourceVector {
	if m != nil {
		return m.ResourceCapacity
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceDescriptor)(nil), "sirius.ResourceDescriptor")
	proto.RegisterEnum("sirius.ResourceDescriptor_ResourceState", ResourceDescriptor_ResourceState_name, ResourceDescriptor_ResourceState_value)
	proto.RegisterEnum("sirius.ResourceDescriptor_ResourceType", ResourceDescriptor_ResourceType_name, ResourceDescriptor_ResourceType_value)
}

func init() { proto.RegisterFile("resource_desc.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0x5f, 0x6f, 0xda, 0x3c,
	0x14, 0xc6, 0x5f, 0x0a, 0xb4, 0xf4, 0x50, 0xc0, 0x75, 0xff, 0xbc, 0x51, 0xa5, 0x4e, 0x88, 0x9b,
	0xf5, 0x0a, 0x6d, 0x4c, 0xbb, 0x41, 0xda, 0x24, 0x96, 0x44, 0x5b, 0x04, 0x4d, 0xa6, 0x04, 0x36,
	0xed, 0x66, 0x91, 0x49, 0xdc, 0x35, 0x5a, 0x48, 0x90, 0xed, 0xb0, 0xf1, 0xa9, 0xb7, 0x8f, 0x30,
	0x07, 0x08, 0x34, 0x65, 0x4a, 0x77, 0x99, 0xe7, 0x39, 0xbf, 0xf3, 0x1c, 0xcb, 0xc7, 0x81, 0x33,
	0x46, 0x79, 0x9c, 0x30, 0x8f, 0xba, 0x3e, 0xe5, 0x5e, 0x77, 0xce, 0x62, 0x11, 0xe3, 0x43, 0x1e,
	0xb0, 0x20, 0xe1, 0x57, 0x17, 0x5b, 0x73, 0x41, 0x3d, 0x11, 0xb3, 0xb5, 0xdd, 0xf9, 0x55, 0x03,
	0x6c, 0x6f, 0x1c, 0x4d, 0x52, 0x2c, 0x98, 0x4b, 0x13, 0x9f, 0x40, 0x25, 0x49, 0x02, 0x5f, 0x29,
	0xb5, 0x0f, 0x6e, 0x8e, 0xf1, 0x05, 0x34, 0xee, 0x58, 0x40, 0x23, 0x3f, 0x5c, 0xba, 0x11, 0x99,
	0x51, 0xe5, 0xa0, 0x5d, 0x92, 0xb2, 0x02, 0xc8, 0xdf, 0x20, 0xc1, 0x82, 0xae, 0x9d, 0xf2, 0xca,
	0x31, 0xa0, 0xca, 0x05, 0x11, 0x54, 0xa9, 0xc8, 0xcf, 0x66, 0xef, 0xa6, 0xbb, 0x1e, 0xa2, 0xbb,
	0x9f, 0xb4, 0x95, 0x9c, 0xb4, 0xbe, 0x8f, 0x6c, 0xdd, 0xb1, 0x26, 0xb6, 0xaa, 0xbb, 0x13, 0x73,
	0x68, 0x5a, 0x9f, 0x4d, 0x19, 0xd2, 0x10, 0x84, 0x7f, 0x77, 0x3d, 0x32, 0x27, 0x5e, 0x20, 0x96,
	0x4a, 0x55, 0xb6, 0xac, 0xf4, 0x4b, 0x2f, 0xf0, 0x25, 0x34, 0x43, 0xc2, 0x85, 0x7b, 0x4f, 0x09,
	0x13, 0x53, 0x4a, 0x84, 0x72, 0x98, 0x5a, 0xf8, 0x35, 0x54, 0xc4, 0x72, 0x4e, 0x95, 0xa3, 0x55,
	0xf6, 0xf3, 0x7f, 0xc8, 0x1e, 0xcb, 0x72, 0x7c, 0x05, 0x75, 0xee, 0xdd, 0x53, 0x3f, 0x09, 0xc9,
	0x34, 0xa4, 0x4a, 0x4d, 0xd2, 0xb5, 0x7e, 0xf5, 0x8e, 0x84, 0x9c, 0xe2, 0x6b, 0xb8, 0xf0, 0x12,
	0xc6, 0x68, 0x24, 0x5c, 0x96, 0x44, 0x51, 0x10, 0x7d, 0x73, 0xd3, 0xa1, 0xb8, 0x72, 0xdc, 0x2e,
	0xcb, 0xc4, 0x0e, 0xfc, 0x1f, 0x25, 0xb3, 0xbc, 0xe5, 0x4e, 0x69, 0x18, 0xff, 0x50, 0x20, 0x9b,
	0xf6, 0x0a, 0x5a, 0x69, 0x0d, 0x0f, 0x63, 0x91, 0x79, 0xf5, 0xcc, 0x7b, 0x05, 0x67, 0x64, 0x41,
	0x82, 0x55, 0xb0, 0x9b, 0xdd, 0x13, 0x57, 0x4e, 0xa4, 0x5f, 0xef, 0x5d, 0x3e, 0x3e, 0xc0, 0xa7,
	0xd5, 0xfd, 0xe1, 0x1e, 0x60, 0x59, 0x4a, 0xd9, 0x82, 0xfa, 0x0f, 0x98, 0x46, 0x21, 0xf3, 0x06,
	0xae, 0x67, 0x41, 0xe4, 0xfe, 0x25, 0x6c, 0x33, 0x52, 0xf3, 0x49, 0x9c, 0xfc, 0x2c, 0xc0, 0x5b,
	0x85, 0xf8, 0x5b, 0x78, 0x96, 0xa6, 0x27, 0xd1, 0xfe, 0xdc, 0x1b, 0x1e, 0x3d, 0xc9, 0xcb, 0xf8,
	0x02, 0xfe, 0xb4, 0x90, 0x7f, 0x09, 0xa7, 0xdb, 0x47, 0xb0, 0x5d, 0x27, 0x5c, 0x84, 0x74, 0xbe,
	0x42, 0x23, 0xb7, 0xa0, 0xf8, 0x1c, 0xf6, 0x56, 0x14, 0xfd, 0x87, 0x4f, 0x65, 0x59, 0xa6, 0x1a,
	0xda, 0x48, 0x47, 0xa5, 0x9c, 0xf4, 0x6e, 0xe2, 0x7c, 0x41, 0x07, 0x39, 0x69, 0x64, 0x39, 0x63,
	0x54, 0xee, 0xfc, 0x2e, 0xc1, 0x49, 0x6e, 0x0b, 0x5b, 0x50, 0xdf, 0xd6, 0x7c, 0x9c, 0x3c, 0x6a,
	0xad, 0x5a, 0x76, 0xda, 0x1a, 0x43, 0x73, 0x27, 0x0d, 0xd4, 0x0f, 0xba, 0xec, 0x8d, 0x64, 0x9f,
	0x4c, 0x33, 0x0d, 0x15, 0x95, 0x73, 0xa0, 0x66, 0x38, 0x43, 0x54, 0xc9, 0x15, 0x39, 0x8e, 0x86,
	0xaa, 0xb9, 0xe3, 0xdc, 0xca, 0x56, 0x86, 0xa9, 0xa3, 0xc3, 0x9c, 0x3a, 0xb2, 0xde, 0x1b, 0xea,
	0x60, 0x84, 0x8e, 0xe4, 0x7b, 0xc3, 0xbb, 0x88, 0xc9, 0xed, 0xc0, 0x35, 0x2d, 0x4d, 0x47, 0x35,
	0x7c, 0x06, 0xad, 0x5d, 0x57, 0x4b, 0x1d, 0xea, 0x63, 0x94, 0xfe, 0x1b, 0xce, 0x1f, 0x8c, 0x6d,
	0xd9, 0x9a, 0x61, 0x0e, 0xc6, 0x96, 0x8d, 0xe0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x6c,
	0xa5, 0xe0, 0xa6, 0x04, 0x00, 0x00,
}
