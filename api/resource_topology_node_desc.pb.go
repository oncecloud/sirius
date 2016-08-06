// Code generated by protoc-gen-go.
// source: resource_topology_node_desc.proto
// DO NOT EDIT!

package sirius

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResourceTopologyNodeDescriptor struct {
	ResourceDesc     *ResourceDescriptor               `protobuf:"bytes,1,req,name=resource_desc" json:"resource_desc,omitempty"`
	Children         []*ResourceTopologyNodeDescriptor `protobuf:"bytes,2,rep,name=children" json:"children,omitempty"`
	ParentId         *string                           `protobuf:"bytes,3,opt,name=parent_id" json:"parent_id,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *ResourceTopologyNodeDescriptor) Reset()                    { *m = ResourceTopologyNodeDescriptor{} }
func (m *ResourceTopologyNodeDescriptor) String() string            { return proto.CompactTextString(m) }
func (*ResourceTopologyNodeDescriptor) ProtoMessage()               {}
func (*ResourceTopologyNodeDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *ResourceTopologyNodeDescriptor) GetResourceDesc() *ResourceDescriptor {
	if m != nil {
		return m.ResourceDesc
	}
	return nil
}

func (m *ResourceTopologyNodeDescriptor) GetChildren() []*ResourceTopologyNodeDescriptor {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *ResourceTopologyNodeDescriptor) GetParentId() string {
	if m != nil && m.ParentId != nil {
		return *m.ParentId
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceTopologyNodeDescriptor)(nil), "sirius.ResourceTopologyNodeDescriptor")
}

func init() { proto.RegisterFile("resource_topology_node_desc.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x8d, 0x2f, 0xc9, 0x2f, 0xc8, 0xcf, 0xc9, 0x4f, 0xaf, 0x8c, 0xcf, 0xcb,
	0x4f, 0x49, 0x8d, 0x4f, 0x49, 0x2d, 0x4e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b,
	0xce, 0x2c, 0xca, 0x2c, 0x2d, 0x96, 0x12, 0x86, 0x2b, 0x45, 0x48, 0x2a, 0x2d, 0x60, 0xe4, 0x92,
	0x0b, 0x82, 0x8a, 0x87, 0x40, 0x4d, 0xf0, 0x03, 0x1a, 0xe0, 0x02, 0x54, 0x52, 0x94, 0x59, 0x50,
	0x92, 0x5f, 0x24, 0x64, 0xc8, 0xc5, 0x8b, 0xa2, 0x53, 0x82, 0x51, 0x81, 0x49, 0x83, 0xdb, 0x48,
	0x4a, 0x0f, 0x62, 0xae, 0x1e, 0x4c, 0x3b, 0x92, 0x16, 0x0b, 0x2e, 0x8e, 0xe4, 0x8c, 0xcc, 0x9c,
	0x94, 0xa2, 0xd4, 0x3c, 0x09, 0x26, 0x05, 0x66, 0xa0, 0x6a, 0x35, 0x74, 0xd5, 0x38, 0x2c, 0x13,
	0xe4, 0xe2, 0x2c, 0x48, 0x04, 0x6a, 0x2b, 0x89, 0xcf, 0x4c, 0x91, 0x60, 0x56, 0x60, 0xd4, 0xe0,
	0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x77, 0xc9, 0x2a, 0xe3, 0x00, 0x00, 0x00,
}
