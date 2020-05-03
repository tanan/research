// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/article.proto

package article

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ArticleRequest struct {
	ArticleId            string   `protobuf:"bytes,1,opt,name=articleId,proto3" json:"articleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleRequest) Reset()         { *m = ArticleRequest{} }
func (m *ArticleRequest) String() string { return proto.CompactTextString(m) }
func (*ArticleRequest) ProtoMessage()    {}
func (*ArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{0}
}

func (m *ArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleRequest.Unmarshal(m, b)
}
func (m *ArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleRequest.Marshal(b, m, deterministic)
}
func (m *ArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleRequest.Merge(m, src)
}
func (m *ArticleRequest) XXX_Size() int {
	return xxx_messageInfo_ArticleRequest.Size(m)
}
func (m *ArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleRequest proto.InternalMessageInfo

func (m *ArticleRequest) GetArticleId() string {
	if m != nil {
		return m.ArticleId
	}
	return ""
}

type ArticleResponse struct {
	ArticleId            string               `protobuf:"bytes,1,opt,name=articleId,proto3" json:"articleId,omitempty"`
	ArticleName          string               `protobuf:"bytes,2,opt,name=articleName,proto3" json:"articleName,omitempty"`
	Editor               string               `protobuf:"bytes,3,opt,name=editor,proto3" json:"editor,omitempty"`
	LastModified         *timestamp.Timestamp `protobuf:"bytes,4,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	Thumbnail            string               `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Description          string               `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ArticleResponse) Reset()         { *m = ArticleResponse{} }
func (m *ArticleResponse) String() string { return proto.CompactTextString(m) }
func (*ArticleResponse) ProtoMessage()    {}
func (*ArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{1}
}

func (m *ArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleResponse.Unmarshal(m, b)
}
func (m *ArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleResponse.Marshal(b, m, deterministic)
}
func (m *ArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleResponse.Merge(m, src)
}
func (m *ArticleResponse) XXX_Size() int {
	return xxx_messageInfo_ArticleResponse.Size(m)
}
func (m *ArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleResponse proto.InternalMessageInfo

func (m *ArticleResponse) GetArticleId() string {
	if m != nil {
		return m.ArticleId
	}
	return ""
}

func (m *ArticleResponse) GetArticleName() string {
	if m != nil {
		return m.ArticleName
	}
	return ""
}

func (m *ArticleResponse) GetEditor() string {
	if m != nil {
		return m.Editor
	}
	return ""
}

func (m *ArticleResponse) GetLastModified() *timestamp.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

func (m *ArticleResponse) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *ArticleResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ArticleRequest)(nil), "article.ArticleRequest")
	proto.RegisterType((*ArticleResponse)(nil), "article.ArticleResponse")
}

func init() {
	proto.RegisterFile("proto/article.proto", fileDescriptor_60692e730edf4587)
}

var fileDescriptor_60692e730edf4587 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0xa4, 0xab, 0x76, 0xd9, 0x54, 0x14, 0x22, 0x68, 0x28, 0x0b, 0x96, 0x9e, 0xf6, 0x94, 0xe0,
	0x7a, 0x17, 0xbc, 0x08, 0x1e, 0xf4, 0x50, 0xfc, 0x81, 0x74, 0x93, 0x5d, 0x03, 0x6d, 0x52, 0x9b,
	0x57, 0x2f, 0xe2, 0xc5, 0x5f, 0xf0, 0xd3, 0xfc, 0x05, 0x2f, 0xfe, 0x85, 0xb4, 0x49, 0xea, 0x2e,
	0xc2, 0x1e, 0x67, 0xe6, 0xbd, 0xc7, 0xcc, 0x3c, 0x74, 0xd6, 0xb4, 0x06, 0x0c, 0xe3, 0x2d, 0xa8,
	0x55, 0x25, 0xe9, 0x80, 0xf0, 0xd4, 0xc3, 0xf4, 0x72, 0x63, 0xcc, 0xa6, 0x92, 0x6c, 0xa0, 0xcb,
	0x6e, 0xcd, 0x40, 0xd5, 0xd2, 0x02, 0xaf, 0x1b, 0x37, 0x99, 0xce, 0xfd, 0x00, 0x6f, 0x14, 0xe3,
	0x5a, 0x1b, 0xe0, 0xa0, 0x8c, 0xb6, 0x4e, 0xcd, 0x29, 0x3a, 0xb9, 0x75, 0x97, 0x0a, 0xf9, 0xd2,
	0x49, 0x0b, 0x78, 0x8e, 0x66, 0xfe, 0xf6, 0xbd, 0x20, 0x51, 0x16, 0x2d, 0x66, 0xc5, 0x1f, 0x91,
	0xff, 0x44, 0xe8, 0x74, 0x5c, 0xb0, 0x8d, 0xd1, 0x56, 0xee, 0xdf, 0xc0, 0x19, 0x4a, 0x3c, 0x78,
	0xe4, 0xb5, 0x24, 0x93, 0x41, 0xdf, 0xa6, 0xf0, 0x39, 0x8a, 0xa5, 0x50, 0x60, 0x5a, 0x72, 0x30,
	0x88, 0x1e, 0xe1, 0x1b, 0x74, 0x5c, 0x71, 0x0b, 0x0f, 0x46, 0xa8, 0xb5, 0x92, 0x82, 0x1c, 0x66,
	0xd1, 0x22, 0x59, 0xa6, 0xd4, 0x05, 0xa2, 0x21, 0x31, 0x7d, 0x0a, 0x89, 0x8b, 0x9d, 0xf9, 0xde,
	0x17, 0x3c, 0x77, 0x75, 0xa9, 0xb9, 0xaa, 0xc8, 0x91, 0xf3, 0x35, 0x12, 0xbd, 0x2f, 0x21, 0xed,
	0xaa, 0x55, 0x4d, 0xdf, 0x07, 0x89, 0x9d, 0xaf, 0x2d, 0x6a, 0x59, 0xa3, 0xa9, 0x8f, 0x8a, 0x4b,
	0x94, 0xdc, 0x29, 0x2d, 0x02, 0xbc, 0xa0, 0xe1, 0x1b, 0xbb, 0xe5, 0xa5, 0xe4, 0xbf, 0xe0, 0x4a,
	0xca, 0xb3, 0x8f, 0xaf, 0xef, 0xcf, 0x49, 0x8a, 0x09, 0x7b, 0xbd, 0x0a, 0xbf, 0xb4, 0xec, 0x6d,
	0xec, 0xe9, 0xbd, 0x8c, 0x87, 0x40, 0xd7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xc8, 0x7a,
	0xac, 0xf0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleClient interface {
	FindArticle(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*ArticleResponse, error)
}

type articleClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleClient(cc grpc.ClientConnInterface) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) FindArticle(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*ArticleResponse, error) {
	out := new(ArticleResponse)
	err := c.cc.Invoke(ctx, "/article.Article/FindArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
type ArticleServer interface {
	FindArticle(context.Context, *ArticleRequest) (*ArticleResponse, error)
}

// UnimplementedArticleServer can be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (*UnimplementedArticleServer) FindArticle(ctx context.Context, req *ArticleRequest) (*ArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindArticle not implemented")
}

func RegisterArticleServer(s *grpc.Server, srv ArticleServer) {
	s.RegisterService(&_Article_serviceDesc, srv)
}

func _Article_FindArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).FindArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.Article/FindArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).FindArticle(ctx, req.(*ArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Article_serviceDesc = grpc.ServiceDesc{
	ServiceName: "article.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindArticle",
			Handler:    _Article_FindArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/article.proto",
}