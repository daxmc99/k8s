// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/api/settings/v1alpha1/generated.proto

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/api/settings/v1alpha1/generated.proto

	It has these top-level messages:
		PodPreset
		PodPresetList
		PodPresetSpec
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import k8s_io_api_core_v1 "github.com/ericchiang/k8s/apis/core/v1"
import k8s_io_apimachinery_pkg_apis_meta_v1 "github.com/ericchiang/k8s/apis/meta/v1"
import _ "github.com/ericchiang/k8s/runtime"
import _ "github.com/ericchiang/k8s/runtime/schema"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PodPreset is a policy resource that defines additional runtime
// requirements for a Pod.
type PodPreset struct {
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// +optional
	Spec             *PodPresetSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *PodPreset) Reset()                    { *m = PodPreset{} }
func (m *PodPreset) String() string            { return proto.CompactTextString(m) }
func (*PodPreset) ProtoMessage()               {}
func (*PodPreset) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PodPreset) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PodPreset) GetSpec() *PodPresetSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// PodPresetList is a list of PodPreset objects.
type PodPresetList struct {
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Items is a list of schema objects.
	Items            []*PodPreset `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PodPresetList) Reset()                    { *m = PodPresetList{} }
func (m *PodPresetList) String() string            { return proto.CompactTextString(m) }
func (*PodPresetList) ProtoMessage()               {}
func (*PodPresetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *PodPresetList) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PodPresetList) GetItems() []*PodPreset {
	if m != nil {
		return m.Items
	}
	return nil
}

// PodPresetSpec is a description of a pod preset.
type PodPresetSpec struct {
	// Selector is a label query over a set of resources, in this case pods.
	// Required.
	Selector *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// Env defines the collection of EnvVar to inject into containers.
	// +optional
	Env []*k8s_io_api_core_v1.EnvVar `protobuf:"bytes,2,rep,name=env" json:"env,omitempty"`
	// EnvFrom defines the collection of EnvFromSource to inject into containers.
	// +optional
	EnvFrom []*k8s_io_api_core_v1.EnvFromSource `protobuf:"bytes,3,rep,name=envFrom" json:"envFrom,omitempty"`
	// Volumes defines the collection of Volume to inject into the pod.
	// +optional
	Volumes []*k8s_io_api_core_v1.Volume `protobuf:"bytes,4,rep,name=volumes" json:"volumes,omitempty"`
	// VolumeMounts defines the collection of VolumeMount to inject into containers.
	// +optional
	VolumeMounts     []*k8s_io_api_core_v1.VolumeMount `protobuf:"bytes,5,rep,name=volumeMounts" json:"volumeMounts,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *PodPresetSpec) Reset()                    { *m = PodPresetSpec{} }
func (m *PodPresetSpec) String() string            { return proto.CompactTextString(m) }
func (*PodPresetSpec) ProtoMessage()               {}
func (*PodPresetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *PodPresetSpec) GetSelector() *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *PodPresetSpec) GetEnv() []*k8s_io_api_core_v1.EnvVar {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *PodPresetSpec) GetEnvFrom() []*k8s_io_api_core_v1.EnvFromSource {
	if m != nil {
		return m.EnvFrom
	}
	return nil
}

func (m *PodPresetSpec) GetVolumes() []*k8s_io_api_core_v1.Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *PodPresetSpec) GetVolumeMounts() []*k8s_io_api_core_v1.VolumeMount {
	if m != nil {
		return m.VolumeMounts
	}
	return nil
}

func init() {
	proto.RegisterType((*PodPreset)(nil), "k8s.io.api.settings.v1alpha1.PodPreset")
	proto.RegisterType((*PodPresetList)(nil), "k8s.io.api.settings.v1alpha1.PodPresetList")
	proto.RegisterType((*PodPresetSpec)(nil), "k8s.io.api.settings.v1alpha1.PodPresetSpec")
}
func (m *PodPreset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodPreset) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Spec != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
		n2, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PodPresetList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodPresetList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n3, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PodPresetSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodPresetSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Selector != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Selector.Size()))
		n4, err := m.Selector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.Env) > 0 {
		for _, msg := range m.Env {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.EnvFrom) > 0 {
		for _, msg := range m.EnvFrom {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Volumes) > 0 {
		for _, msg := range m.Volumes {
			dAtA[i] = 0x22
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.VolumeMounts) > 0 {
		for _, msg := range m.VolumeMounts {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PodPreset) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PodPresetList) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PodPresetSpec) Size() (n int) {
	var l int
	_ = l
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Env) > 0 {
		for _, e := range m.Env {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.EnvFrom) > 0 {
		for _, e := range m.EnvFrom {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Volumes) > 0 {
		for _, e := range m.Volumes {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.VolumeMounts) > 0 {
		for _, e := range m.VolumeMounts {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PodPreset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodPreset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodPreset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &PodPresetSpec{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodPresetList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodPresetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodPresetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &PodPreset{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodPresetSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodPresetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodPresetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = append(m.Env, &k8s_io_api_core_v1.EnvVar{})
			if err := m.Env[len(m.Env)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvFrom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnvFrom = append(m.EnvFrom, &k8s_io_api_core_v1.EnvFromSource{})
			if err := m.EnvFrom[len(m.EnvFrom)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volumes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Volumes = append(m.Volumes, &k8s_io_api_core_v1.Volume{})
			if err := m.Volumes[len(m.Volumes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeMounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VolumeMounts = append(m.VolumeMounts, &k8s_io_api_core_v1.VolumeMount{})
			if err := m.VolumeMounts[len(m.VolumeMounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/api/settings/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8a, 0x13, 0x31,
	0x1c, 0x87, 0x9d, 0x76, 0x97, 0xad, 0x59, 0xbd, 0xe4, 0x34, 0x0c, 0x32, 0xae, 0xbd, 0xb8, 0xe0,
	0x92, 0xd8, 0x75, 0x11, 0x41, 0x44, 0x50, 0xf4, 0x20, 0xbb, 0xec, 0x92, 0xc2, 0x1e, 0xbc, 0xa5,
	0xe9, 0x9f, 0x69, 0xec, 0x4c, 0x12, 0x92, 0xcc, 0x80, 0x6f, 0x22, 0x3e, 0x90, 0x78, 0xf4, 0x11,
	0xa4, 0xbe, 0x88, 0x64, 0xda, 0x99, 0xd6, 0xd6, 0xd1, 0xde, 0x4a, 0xf8, 0xbe, 0x5f, 0x3f, 0xfe,
	0x0c, 0x3a, 0x9b, 0xbf, 0x70, 0x44, 0x6a, 0xca, 0x8d, 0xa4, 0x0e, 0xbc, 0x97, 0x2a, 0x73, 0xb4,
	0x1a, 0xf1, 0xdc, 0xcc, 0xf8, 0x88, 0x66, 0xa0, 0xc0, 0x72, 0x0f, 0x53, 0x62, 0xac, 0xf6, 0x1a,
	0x3f, 0x58, 0xd2, 0x84, 0x1b, 0x49, 0x1a, 0x9a, 0x34, 0x74, 0x32, 0xdc, 0xd8, 0x12, 0xda, 0x02,
	0xad, 0x76, 0x16, 0x92, 0x8b, 0x35, 0x53, 0x70, 0x31, 0x93, 0x0a, 0xec, 0x67, 0x6a, 0xe6, 0x59,
	0x78, 0x70, 0xb4, 0x00, 0xcf, 0xff, 0x66, 0xd1, 0x2e, 0xcb, 0x96, 0xca, 0xcb, 0x02, 0x76, 0x84,
	0xe7, 0xff, 0x13, 0x9c, 0x98, 0x41, 0xc1, 0xb7, 0xbd, 0xe1, 0xd7, 0x08, 0xdd, 0xbd, 0xd1, 0xd3,
	0x1b, 0x0b, 0x0e, 0x3c, 0xbe, 0x44, 0x83, 0x50, 0x34, 0xe5, 0x9e, 0xc7, 0xd1, 0x49, 0x74, 0x7a,
	0x7c, 0xfe, 0x94, 0xac, 0x2f, 0xd0, 0x0e, 0x13, 0x33, 0xcf, 0xc2, 0x83, 0x23, 0x81, 0x26, 0xd5,
	0x88, 0x5c, 0x4f, 0x3e, 0x81, 0xf0, 0x57, 0xe0, 0x39, 0x6b, 0x17, 0xf0, 0x6b, 0x74, 0xe0, 0x0c,
	0x88, 0xb8, 0x57, 0x2f, 0x3d, 0x21, 0xff, 0xba, 0x25, 0x69, 0x23, 0xc6, 0x06, 0x04, 0xab, 0xc5,
	0x10, 0x77, 0xbf, 0x7d, 0xbf, 0x94, 0xce, 0xe3, 0x0f, 0x3b, 0x81, 0x64, 0xbf, 0xc0, 0x60, 0x6f,
	0xe5, 0xbd, 0x42, 0x87, 0xd2, 0x43, 0xe1, 0xe2, 0xde, 0x49, 0xff, 0xf4, 0xf8, 0xfc, 0xf1, 0x9e,
	0x7d, 0x6c, 0x69, 0x0d, 0xbf, 0xf5, 0x36, 0xe2, 0x42, 0x34, 0xbe, 0x46, 0x03, 0x07, 0x39, 0x08,
	0xaf, 0xed, 0x2a, 0xee, 0xd9, 0x9e, 0x71, 0x7c, 0x02, 0xf9, 0x78, 0xa5, 0xb2, 0x76, 0x04, 0x9f,
	0xa1, 0x3e, 0xa8, 0x6a, 0xd5, 0x97, 0x6c, 0xf6, 0x85, 0xaf, 0x2d, 0x98, 0xef, 0x54, 0x75, 0xcb,
	0x2d, 0x0b, 0x18, 0x7e, 0x89, 0x8e, 0x40, 0x55, 0xef, 0xad, 0x2e, 0xe2, 0x7e, 0x6d, 0x3c, 0xea,
	0x30, 0x02, 0x32, 0xd6, 0xa5, 0x15, 0xc0, 0x1a, 0x03, 0x5f, 0xa0, 0xa3, 0x4a, 0xe7, 0x65, 0x01,
	0x2e, 0x3e, 0xe8, 0xfe, 0xbb, 0xdb, 0x1a, 0x61, 0x0d, 0x8a, 0xdf, 0xa2, 0x7b, 0xcb, 0x9f, 0x57,
	0xba, 0x54, 0xde, 0xc5, 0x87, 0xb5, 0xfa, 0xb0, 0x5b, 0xad, 0x39, 0xf6, 0x87, 0xf4, 0x26, 0xf9,
	0xbe, 0x48, 0xa3, 0x1f, 0x8b, 0x34, 0xfa, 0xb9, 0x48, 0xa3, 0x2f, 0xbf, 0xd2, 0x3b, 0x1f, 0x07,
	0xcd, 0xd5, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x54, 0x5b, 0x55, 0xae, 0x03, 0x00, 0x00,
}
