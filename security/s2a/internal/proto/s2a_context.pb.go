// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s2a_context.proto

package s2a_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type S2AContext struct {
	// The application protocol negotiated for this connection, e.g., 'grpc'.
	ApplicationProtocol string `protobuf:"bytes,1,opt,name=application_protocol,json=applicationProtocol,proto3" json:"application_protocol,omitempty"`
	// The TLS version number that the S2A's handshaker module used to set up the
	// session.
	TlsVersion TLSVersion `protobuf:"varint,2,opt,name=tls_version,json=tlsVersion,proto3,enum=s2a.proto.TLSVersion" json:"tls_version,omitempty"`
	// The TLS ciphersuite negotiated by the S2A's handshaker module.
	Ciphersuite Ciphersuite `protobuf:"varint,3,opt,name=ciphersuite,proto3,enum=s2a.proto.Ciphersuite" json:"ciphersuite,omitempty"`
	// The authenticated identity of the peer.
	PeerIdentity *Identity `protobuf:"bytes,4,opt,name=peer_identity,json=peerIdentity,proto3" json:"peer_identity,omitempty"`
	// The local identity used during session setup. This could be:
	// - The local identity that the client specifies in ClientSessionStartReq.
	// - One of the local identities that the server specifies in
	//   ServerSessionStartReq.
	// - If neither client or server specifies local identities, the S2A picks the
	//   default one. In this case, this field will contain that identity.
	LocalIdentity *Identity `protobuf:"bytes,5,opt,name=local_identity,json=localIdentity,proto3" json:"local_identity,omitempty"`
	// The SHA256 hash of the peer certificate used in the handshake.
	PeerCertFingerprint []byte `protobuf:"bytes,6,opt,name=peer_cert_fingerprint,json=peerCertFingerprint,proto3" json:"peer_cert_fingerprint,omitempty"`
	// The SHA256 hash of the local certificate used in the handshake.
	LocalCertFingerprint []byte   `protobuf:"bytes,7,opt,name=local_cert_fingerprint,json=localCertFingerprint,proto3" json:"local_cert_fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2AContext) Reset()         { *m = S2AContext{} }
func (m *S2AContext) String() string { return proto.CompactTextString(m) }
func (*S2AContext) ProtoMessage()    {}
func (*S2AContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d9bb22991f97e4a, []int{0}
}

func (m *S2AContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2AContext.Unmarshal(m, b)
}
func (m *S2AContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2AContext.Marshal(b, m, deterministic)
}
func (m *S2AContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2AContext.Merge(m, src)
}
func (m *S2AContext) XXX_Size() int {
	return xxx_messageInfo_S2AContext.Size(m)
}
func (m *S2AContext) XXX_DiscardUnknown() {
	xxx_messageInfo_S2AContext.DiscardUnknown(m)
}

var xxx_messageInfo_S2AContext proto.InternalMessageInfo

func (m *S2AContext) GetApplicationProtocol() string {
	if m != nil {
		return m.ApplicationProtocol
	}
	return ""
}

func (m *S2AContext) GetTlsVersion() TLSVersion {
	if m != nil {
		return m.TlsVersion
	}
	return TLSVersion_TLS1_2
}

func (m *S2AContext) GetCiphersuite() Ciphersuite {
	if m != nil {
		return m.Ciphersuite
	}
	return Ciphersuite_AES_128_GCM_SHA256
}

func (m *S2AContext) GetPeerIdentity() *Identity {
	if m != nil {
		return m.PeerIdentity
	}
	return nil
}

func (m *S2AContext) GetLocalIdentity() *Identity {
	if m != nil {
		return m.LocalIdentity
	}
	return nil
}

func (m *S2AContext) GetPeerCertFingerprint() []byte {
	if m != nil {
		return m.PeerCertFingerprint
	}
	return nil
}

func (m *S2AContext) GetLocalCertFingerprint() []byte {
	if m != nil {
		return m.LocalCertFingerprint
	}
	return nil
}

func init() {
	proto.RegisterType((*S2AContext)(nil), "s2a.proto.S2AContext")
}

func init() {
	proto.RegisterFile("s2a_context.proto", fileDescriptor_0d9bb22991f97e4a)
}

var fileDescriptor_0d9bb22991f97e4a = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0x87, 0xa9, 0xd3, 0xc9, 0x5e, 0xbb, 0x81, 0xe9, 0x36, 0x8a, 0xa7, 0xe2, 0xa9, 0xa7, 0x82,
	0x51, 0x64, 0x78, 0x93, 0x82, 0x20, 0x78, 0x90, 0x4e, 0xbc, 0x96, 0x1a, 0xa3, 0x06, 0xb2, 0x24,
	0x24, 0x4f, 0xd1, 0x3f, 0xc3, 0xff, 0x58, 0x96, 0xd4, 0xb5, 0x28, 0xde, 0xf2, 0xf8, 0xbe, 0x8f,
	0x5f, 0xe0, 0xc8, 0xd1, 0xb6, 0x61, 0x5a, 0x21, 0xff, 0xc0, 0xd2, 0x58, 0x8d, 0x9a, 0x4c, 0x1c,
	0x6d, 0xc3, 0xf3, 0x38, 0x61, 0x7a, 0xb3, 0xd1, 0x2a, 0x5c, 0x27, 0x5f, 0x23, 0x80, 0x35, 0xbd,
	0xaa, 0x82, 0x4d, 0x4e, 0x61, 0xde, 0x1a, 0x23, 0x05, 0x6b, 0x51, 0x68, 0xd5, 0x78, 0x87, 0x69,
	0x99, 0x45, 0x79, 0x54, 0x4c, 0xea, 0x74, 0xc0, 0xee, 0x3a, 0x44, 0x2e, 0x20, 0x46, 0xe9, 0x9a,
	0x77, 0x6e, 0x9d, 0xd0, 0x2a, 0xdb, 0xcb, 0xa3, 0x62, 0x46, 0x17, 0xe5, 0x6e, 0xb0, 0xbc, 0xbf,
	0x5d, 0x3f, 0x04, 0x58, 0x03, 0x4a, 0xd7, 0xbd, 0xc9, 0x0a, 0x62, 0x26, 0xcc, 0x2b, 0xb7, 0xee,
	0x4d, 0x20, 0xcf, 0x46, 0xbe, 0x5b, 0x0e, 0xba, 0xaa, 0xa7, 0xf5, 0x50, 0x25, 0x2b, 0x98, 0x1a,
	0xce, 0x6d, 0x23, 0x9e, 0xb8, 0x42, 0x81, 0x9f, 0xd9, 0x7e, 0x1e, 0x15, 0x31, 0x4d, 0x07, 0xed,
	0x4d, 0x87, 0xea, 0x64, 0x6b, 0xfe, 0x5c, 0xe4, 0x12, 0x66, 0x52, 0xb3, 0x56, 0xf6, 0xe9, 0xc1,
	0xff, 0xe9, 0xd4, 0xab, 0xbb, 0x96, 0xc2, 0xc2, 0xaf, 0x32, 0x6e, 0xb1, 0x79, 0x16, 0xea, 0x85,
	0x5b, 0x63, 0x85, 0xc2, 0x6c, 0x9c, 0x47, 0x45, 0x52, 0xa7, 0x5b, 0x58, 0x71, 0x8b, 0xd7, 0x3d,
	0x22, 0xe7, 0xb0, 0x0c, 0x7b, 0x7f, 0xa2, 0x43, 0x1f, 0xcd, 0x3d, 0xfd, 0x55, 0x3d, 0x8e, 0xfd,
	0x47, 0xce, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xf5, 0xdb, 0x49, 0xc8, 0x01, 0x00, 0x00,
}
