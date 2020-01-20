// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wasme/v1/filter_deployment.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// the state of the filter deployment
type WorkloadStatus_State int32

const (
	WorkloadStatus_Succeeded WorkloadStatus_State = 0
	WorkloadStatus_Failed    WorkloadStatus_State = 1
)

var WorkloadStatus_State_name = map[int32]string{
	0: "Succeeded",
	1: "Failed",
}

var WorkloadStatus_State_value = map[string]int32{
	"Succeeded": 0,
	"Failed":    1,
}

func (x WorkloadStatus_State) String() string {
	return proto.EnumName(WorkloadStatus_State_name, int32(x))
}

func (WorkloadStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8758ba57473cad6a, []int{5, 0}
}

// A FilterDeployment tells the Wasme Operator
// to deploy a filter with the provided configuration
// to the target workloads.
// Currently FilterDeployments support Wasm filters on Istio
type FilterDeploymentSpec struct {
	// the spec of the filter to deploy
	Filter *FilterSpec `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// Spec that selects one or more target workloads in the FilterDeployment namespace
	Deployment           *DeploymentSpec `protobuf:"bytes,2,opt,name=deployment,proto3" json:"deployment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FilterDeploymentSpec) Reset()         { *m = FilterDeploymentSpec{} }
func (m *FilterDeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*FilterDeploymentSpec) ProtoMessage()    {}
func (*FilterDeploymentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8758ba57473cad6a, []int{0}
}
func (m *FilterDeploymentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterDeploymentSpec.Unmarshal(m, b)
}
func (m *FilterDeploymentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterDeploymentSpec.Marshal(b, m, deterministic)
}
func (m *FilterDeploymentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterDeploymentSpec.Merge(m, src)
}
func (m *FilterDeploymentSpec) XXX_Size() int {
	return xxx_messageInfo_FilterDeploymentSpec.Size(m)
}
func (m *FilterDeploymentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterDeploymentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FilterDeploymentSpec proto.InternalMessageInfo

func (m *FilterDeploymentSpec) GetFilter() *FilterSpec {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *FilterDeploymentSpec) GetDeployment() *DeploymentSpec {
	if m != nil {
		return m.Deployment
	}
	return nil
}

// the filter to deploy
type FilterSpec struct {
	// unique identifier that will be used
	// to remove the filter as well as for logging.
	// if id is not set, it will be set automatically to be the name.namespace
	// of the FilterDeployment resource
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name of image which houses the compiled wasm filter
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// string of the config sent to the wasm filter
	// Currently has to be json or will crash
	Config string `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	// the root id must match the root id
	// defined inside the filter.
	// if the user does not provide this field,
	// wasme will attempt to pull the image
	// and set it from the filter_conf
	// the first time it must pull the image and inspect it
	// second time it will cache it locally
	// if the user provides
	RootID               string   `protobuf:"bytes,4,opt,name=rootID,proto3" json:"rootID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterSpec) Reset()         { *m = FilterSpec{} }
func (m *FilterSpec) String() string { return proto.CompactTextString(m) }
func (*FilterSpec) ProtoMessage()    {}
func (*FilterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8758ba57473cad6a, []int{1}
}
func (m *FilterSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterSpec.Unmarshal(m, b)
}
func (m *FilterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterSpec.Marshal(b, m, deterministic)
}
func (m *FilterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterSpec.Merge(m, src)
}
func (m *FilterSpec) XXX_Size() int {
	return xxx_messageInfo_FilterSpec.Size(m)
}
func (m *FilterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FilterSpec proto.InternalMessageInfo

func (m *FilterSpec) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FilterSpec) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *FilterSpec) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *FilterSpec) GetRootID() string {
	if m != nil {
		return m.RootID
	}
	return ""
}

// how to deploy the filter
type DeploymentSpec struct {
	// Types that are valid to be assigned to DeploymentType:
	//	*DeploymentSpec_Istio
	DeploymentType       isDeploymentSpec_DeploymentType `protobuf_oneof:"deploymentType"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DeploymentSpec) Reset()         { *m = DeploymentSpec{} }
func (m *DeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*DeploymentSpec) ProtoMessage()    {}
func (*DeploymentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8758ba57473cad6a, []int{2}
}
func (m *DeploymentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentSpec.Unmarshal(m, b)
}
func (m *DeploymentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentSpec.Marshal(b, m, deterministic)
}
func (m *DeploymentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentSpec.Merge(m, src)
}
func (m *DeploymentSpec) XXX_Size() int {
	return xxx_messageInfo_DeploymentSpec.Size(m)
}
func (m *DeploymentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentSpec proto.InternalMessageInfo

type isDeploymentSpec_DeploymentType interface {
	isDeploymentSpec_DeploymentType()
}

type DeploymentSpec_Istio struct {
	Istio *IstioDeploymentSpec `protobuf:"bytes,2,opt,name=istio,proto3,oneof" json:"istio,omitempty"`
}

func (*DeploymentSpec_Istio) isDeploymentSpec_DeploymentType() {}

func (m *DeploymentSpec) GetDeploymentType() isDeploymentSpec_DeploymentType {
	if m != nil {
		return m.DeploymentType
	}
	return nil
}

func (m *DeploymentSpec) GetIstio() *IstioDeploymentSpec {
	if x, ok := m.GetDeploymentType().(*DeploymentSpec_Istio); ok {
		return x.Istio
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DeploymentSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DeploymentSpec_Istio)(nil),
	}
}

// how to deploy to Istio
type IstioDeploymentSpec struct {
	// the kind of workload to deploy the filter to
	// can either be Deployment or DaemonSet
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// the name of the workload to deploy the filter to
	// the workload must live in the same namespace as the FilterDeployment
	// if empty, the filter will be deployed to all workloads in the namespace
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioDeploymentSpec) Reset()         { *m = IstioDeploymentSpec{} }
func (m *IstioDeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*IstioDeploymentSpec) ProtoMessage()    {}
func (*IstioDeploymentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8758ba57473cad6a, []int{3}
}
func (m *IstioDeploymentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioDeploymentSpec.Unmarshal(m, b)
}
func (m *IstioDeploymentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioDeploymentSpec.Marshal(b, m, deterministic)
}
func (m *IstioDeploymentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioDeploymentSpec.Merge(m, src)
}
func (m *IstioDeploymentSpec) XXX_Size() int {
	return xxx_messageInfo_IstioDeploymentSpec.Size(m)
}
func (m *IstioDeploymentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioDeploymentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioDeploymentSpec proto.InternalMessageInfo

func (m *IstioDeploymentSpec) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *IstioDeploymentSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// the current status of the deployment
type FilterDeploymentStatus struct {
	// the observed generation of the FilterDeployment
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observedGeneration,proto3" json:"observedGeneration,omitempty"`
	// for each workload, was the deployment successful?
	Workloads map[string]*WorkloadStatus `protobuf:"bytes,2,rep,name=workloads,proto3" json:"workloads,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// a human-readable string explaining the error, if any
	Reason               string   `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterDeploymentStatus) Reset()         { *m = FilterDeploymentStatus{} }
func (m *FilterDeploymentStatus) String() string { return proto.CompactTextString(m) }
func (*FilterDeploymentStatus) ProtoMessage()    {}
func (*FilterDeploymentStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8758ba57473cad6a, []int{4}
}
func (m *FilterDeploymentStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterDeploymentStatus.Unmarshal(m, b)
}
func (m *FilterDeploymentStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterDeploymentStatus.Marshal(b, m, deterministic)
}
func (m *FilterDeploymentStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterDeploymentStatus.Merge(m, src)
}
func (m *FilterDeploymentStatus) XXX_Size() int {
	return xxx_messageInfo_FilterDeploymentStatus.Size(m)
}
func (m *FilterDeploymentStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterDeploymentStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FilterDeploymentStatus proto.InternalMessageInfo

func (m *FilterDeploymentStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *FilterDeploymentStatus) GetWorkloads() map[string]*WorkloadStatus {
	if m != nil {
		return m.Workloads
	}
	return nil
}

func (m *FilterDeploymentStatus) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type WorkloadStatus struct {
	State WorkloadStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=wasme.io.WorkloadStatus_State" json:"state,omitempty"`
	// a human-readable string explaining the error, if any
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadStatus) Reset()         { *m = WorkloadStatus{} }
func (m *WorkloadStatus) String() string { return proto.CompactTextString(m) }
func (*WorkloadStatus) ProtoMessage()    {}
func (*WorkloadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8758ba57473cad6a, []int{5}
}
func (m *WorkloadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadStatus.Unmarshal(m, b)
}
func (m *WorkloadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadStatus.Marshal(b, m, deterministic)
}
func (m *WorkloadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadStatus.Merge(m, src)
}
func (m *WorkloadStatus) XXX_Size() int {
	return xxx_messageInfo_WorkloadStatus.Size(m)
}
func (m *WorkloadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadStatus proto.InternalMessageInfo

func (m *WorkloadStatus) GetState() WorkloadStatus_State {
	if m != nil {
		return m.State
	}
	return WorkloadStatus_Succeeded
}

func (m *WorkloadStatus) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterEnum("wasme.io.WorkloadStatus_State", WorkloadStatus_State_name, WorkloadStatus_State_value)
	proto.RegisterType((*FilterDeploymentSpec)(nil), "wasme.io.FilterDeploymentSpec")
	proto.RegisterType((*FilterSpec)(nil), "wasme.io.FilterSpec")
	proto.RegisterType((*DeploymentSpec)(nil), "wasme.io.DeploymentSpec")
	proto.RegisterType((*IstioDeploymentSpec)(nil), "wasme.io.IstioDeploymentSpec")
	proto.RegisterType((*FilterDeploymentStatus)(nil), "wasme.io.FilterDeploymentStatus")
	proto.RegisterMapType((map[string]*WorkloadStatus)(nil), "wasme.io.FilterDeploymentStatus.WorkloadsEntry")
	proto.RegisterType((*WorkloadStatus)(nil), "wasme.io.WorkloadStatus")
}

func init() { proto.RegisterFile("wasme/v1/filter_deployment.proto", fileDescriptor_8758ba57473cad6a) }

var fileDescriptor_8758ba57473cad6a = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x74, 0xa9, 0xe8, 0x9b, 0x88, 0x2a, 0x53, 0x4d, 0x11, 0x12, 0xa8, 0xca, 0x69,
	0x07, 0xb0, 0xb5, 0xc2, 0xc4, 0x84, 0xc4, 0x65, 0x1a, 0x83, 0x1d, 0xb8, 0x78, 0x08, 0x04, 0x17,
	0xe4, 0x26, 0x6f, 0xc5, 0x6a, 0x92, 0x17, 0x25, 0x6e, 0xa7, 0x5e, 0x38, 0x70, 0xe1, 0xdf, 0x46,
	0xb1, 0x13, 0xd2, 0x76, 0xe3, 0x14, 0xbf, 0x97, 0xcf, 0xf7, 0xfb, 0x7e, 0x24, 0x86, 0xe9, 0x9d,
	0xaa, 0x73, 0x14, 0xeb, 0x53, 0x71, 0xab, 0x33, 0x83, 0xd5, 0x8f, 0x14, 0xcb, 0x8c, 0x36, 0x39,
	0x16, 0x86, 0x97, 0x15, 0x19, 0x62, 0x8f, 0x2c, 0xc1, 0x35, 0xc5, 0xbf, 0x60, 0x72, 0x65, 0xa1,
	0xcb, 0x7f, 0xcc, 0x4d, 0x89, 0x09, 0x7b, 0x01, 0x43, 0x27, 0x8e, 0xbc, 0xa9, 0x77, 0x72, 0x34,
	0x9b, 0xf0, 0x4e, 0xc2, 0x1d, 0xdf, 0x50, 0xb2, 0x65, 0xd8, 0x39, 0x40, 0x5f, 0x23, 0xf2, 0xad,
	0x22, 0xea, 0x15, 0xbb, 0xde, 0x72, 0x8b, 0x8d, 0xe7, 0x00, 0xbd, 0x1f, 0x0b, 0xc1, 0xd7, 0xa9,
	0xad, 0x38, 0x92, 0xbe, 0x4e, 0xd9, 0x04, 0x02, 0x9d, 0xab, 0x05, 0x5a, 0xcb, 0x91, 0x74, 0x01,
	0x3b, 0x86, 0x61, 0x42, 0xc5, 0xad, 0x5e, 0x44, 0x03, 0x9b, 0x6e, 0xa3, 0x26, 0x5f, 0x11, 0x99,
	0xeb, 0xcb, 0xe8, 0xd0, 0xe5, 0x5d, 0x14, 0x7f, 0x83, 0x70, 0x6f, 0xba, 0x33, 0x08, 0x74, 0x6d,
	0x34, 0xb5, 0xad, 0x3e, 0xeb, 0x5b, 0xbd, 0x6e, 0xd2, 0xbb, 0xf4, 0xc7, 0x03, 0xe9, 0xe8, 0x8b,
	0x31, 0x84, 0x7d, 0xeb, 0x9f, 0x37, 0x25, 0xc6, 0xef, 0xe0, 0xc9, 0x03, 0x0a, 0xc6, 0xe0, 0x70,
	0xa9, 0x8b, 0x6e, 0x12, 0x7b, 0x6e, 0x72, 0x85, 0xca, 0xbb, 0x51, 0xec, 0x39, 0xfe, 0xe3, 0xc3,
	0xf1, 0xbd, 0xf5, 0x1b, 0x65, 0x56, 0x35, 0xe3, 0xc0, 0x68, 0x5e, 0x63, 0xb5, 0xc6, 0xf4, 0x03,
	0x16, 0x58, 0x29, 0xa3, 0xa9, 0xb0, 0x86, 0x03, 0xf9, 0xc0, 0x1b, 0xf6, 0x09, 0x46, 0x77, 0x54,
	0x2d, 0x33, 0x52, 0x69, 0x1d, 0xf9, 0xd3, 0xc1, 0xc9, 0xd1, 0x4c, 0xec, 0x7f, 0xb3, 0xfd, 0x22,
	0xfc, 0x6b, 0xa7, 0x78, 0x5f, 0x98, 0x6a, 0x23, 0x7b, 0x07, 0xbb, 0x4b, 0x54, 0x35, 0x15, 0xdd,
	0x8e, 0x5d, 0xf4, 0xf4, 0x0b, 0x84, 0xbb, 0x22, 0x36, 0x86, 0xc1, 0x12, 0x37, 0xed, 0xa8, 0xcd,
	0x91, 0x71, 0x08, 0xd6, 0x2a, 0x5b, 0xe1, 0xfd, 0x1f, 0xa1, 0x93, 0xba, 0xf2, 0xd2, 0x61, 0x6f,
	0xfd, 0x73, 0x2f, 0xfe, 0xed, 0xf5, 0xc6, 0xed, 0x06, 0x5e, 0x43, 0x50, 0x1b, 0x65, 0xd0, 0x5a,
	0x87, 0xb3, 0xe7, 0xff, 0xb3, 0xe1, 0xcd, 0x03, 0xa5, 0x83, 0xb7, 0x1a, 0xf7, 0xb7, 0x1b, 0x8f,
	0x63, 0x08, 0x2c, 0xc7, 0x1e, 0xc3, 0xe8, 0x66, 0x95, 0x24, 0x88, 0x29, 0xa6, 0xe3, 0x03, 0x06,
	0x30, 0xbc, 0x52, 0x3a, 0xc3, 0x74, 0xec, 0x5d, 0xbc, 0xf9, 0x7e, 0xb6, 0xd0, 0xe6, 0xe7, 0x6a,
	0xce, 0x13, 0xca, 0x45, 0x4d, 0x19, 0xbd, 0xd4, 0x24, 0xdc, 0x6d, 0xa2, 0xb2, 0xd9, 0x34, 0x55,
	0xa2, 0x5c, 0x2e, 0x84, 0x2a, 0xb5, 0xe8, 0xba, 0x11, 0xeb, 0xd3, 0xf9, 0xd0, 0x5e, 0xab, 0x57,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x21, 0x9c, 0xad, 0xa8, 0x7a, 0x03, 0x00, 0x00,
}