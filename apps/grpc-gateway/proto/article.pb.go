// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/article.proto

package article

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Articles struct {
	Articles             []*Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Articles) Reset()         { *m = Articles{} }
func (m *Articles) String() string { return proto.CompactTextString(m) }
func (*Articles) ProtoMessage()    {}
func (*Articles) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{0}
}

func (m *Articles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Articles.Unmarshal(m, b)
}
func (m *Articles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Articles.Marshal(b, m, deterministic)
}
func (m *Articles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Articles.Merge(m, src)
}
func (m *Articles) XXX_Size() int {
	return xxx_messageInfo_Articles.Size(m)
}
func (m *Articles) XXX_DiscardUnknown() {
	xxx_messageInfo_Articles.DiscardUnknown(m)
}

var xxx_messageInfo_Articles proto.InternalMessageInfo

func (m *Articles) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type Article struct {
	ArticleId            string           `protobuf:"bytes,1,opt,name=articleId,proto3" json:"articleId,omitempty"`
	Overview             *ArticleOverview `protobuf:"bytes,2,opt,name=overview,proto3" json:"overview,omitempty"`
	Content              *_struct.Struct  `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{1}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetArticleId() string {
	if m != nil {
		return m.ArticleId
	}
	return ""
}

func (m *Article) GetOverview() *ArticleOverview {
	if m != nil {
		return m.Overview
	}
	return nil
}

func (m *Article) GetContent() *_struct.Struct {
	if m != nil {
		return m.Content
	}
	return nil
}

type ArticleOverview struct {
	ArticleName          string               `protobuf:"bytes,1,opt,name=articleName,proto3" json:"articleName,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Editor               *Editor              `protobuf:"bytes,3,opt,name=editor,proto3" json:"editor,omitempty"`
	LastModified         *timestamp.Timestamp `protobuf:"bytes,4,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	Thumbnail            string               `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ArticleOverview) Reset()         { *m = ArticleOverview{} }
func (m *ArticleOverview) String() string { return proto.CompactTextString(m) }
func (*ArticleOverview) ProtoMessage()    {}
func (*ArticleOverview) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{2}
}

func (m *ArticleOverview) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleOverview.Unmarshal(m, b)
}
func (m *ArticleOverview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleOverview.Marshal(b, m, deterministic)
}
func (m *ArticleOverview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleOverview.Merge(m, src)
}
func (m *ArticleOverview) XXX_Size() int {
	return xxx_messageInfo_ArticleOverview.Size(m)
}
func (m *ArticleOverview) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleOverview.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleOverview proto.InternalMessageInfo

func (m *ArticleOverview) GetArticleName() string {
	if m != nil {
		return m.ArticleName
	}
	return ""
}

func (m *ArticleOverview) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArticleOverview) GetEditor() *Editor {
	if m != nil {
		return m.Editor
	}
	return nil
}

func (m *ArticleOverview) GetLastModified() *timestamp.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

func (m *ArticleOverview) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

type Editor struct {
	EditorId             int32    `protobuf:"varint,1,opt,name=editorId,proto3" json:"editorId,omitempty"`
	EditorName           string   `protobuf:"bytes,2,opt,name=editorName,proto3" json:"editorName,omitempty"`
	EditorIcon           string   `protobuf:"bytes,3,opt,name=editorIcon,proto3" json:"editorIcon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Editor) Reset()         { *m = Editor{} }
func (m *Editor) String() string { return proto.CompactTextString(m) }
func (*Editor) ProtoMessage()    {}
func (*Editor) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{3}
}

func (m *Editor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Editor.Unmarshal(m, b)
}
func (m *Editor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Editor.Marshal(b, m, deterministic)
}
func (m *Editor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Editor.Merge(m, src)
}
func (m *Editor) XXX_Size() int {
	return xxx_messageInfo_Editor.Size(m)
}
func (m *Editor) XXX_DiscardUnknown() {
	xxx_messageInfo_Editor.DiscardUnknown(m)
}

var xxx_messageInfo_Editor proto.InternalMessageInfo

func (m *Editor) GetEditorId() int32 {
	if m != nil {
		return m.EditorId
	}
	return 0
}

func (m *Editor) GetEditorName() string {
	if m != nil {
		return m.EditorName
	}
	return ""
}

func (m *Editor) GetEditorIcon() string {
	if m != nil {
		return m.EditorIcon
	}
	return ""
}

type StoreArticleRequest struct {
	Article              *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreArticleRequest) Reset()         { *m = StoreArticleRequest{} }
func (m *StoreArticleRequest) String() string { return proto.CompactTextString(m) }
func (*StoreArticleRequest) ProtoMessage()    {}
func (*StoreArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{4}
}

func (m *StoreArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreArticleRequest.Unmarshal(m, b)
}
func (m *StoreArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreArticleRequest.Marshal(b, m, deterministic)
}
func (m *StoreArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreArticleRequest.Merge(m, src)
}
func (m *StoreArticleRequest) XXX_Size() int {
	return xxx_messageInfo_StoreArticleRequest.Size(m)
}
func (m *StoreArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreArticleRequest proto.InternalMessageInfo

func (m *StoreArticleRequest) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type StoreArticleResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreArticleResponse) Reset()         { *m = StoreArticleResponse{} }
func (m *StoreArticleResponse) String() string { return proto.CompactTextString(m) }
func (*StoreArticleResponse) ProtoMessage()    {}
func (*StoreArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{5}
}

func (m *StoreArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreArticleResponse.Unmarshal(m, b)
}
func (m *StoreArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreArticleResponse.Marshal(b, m, deterministic)
}
func (m *StoreArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreArticleResponse.Merge(m, src)
}
func (m *StoreArticleResponse) XXX_Size() int {
	return xxx_messageInfo_StoreArticleResponse.Size(m)
}
func (m *StoreArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreArticleResponse proto.InternalMessageInfo

func (m *StoreArticleResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StoreEditorRequest struct {
	Editor               *Editor  `protobuf:"bytes,1,opt,name=editor,proto3" json:"editor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreEditorRequest) Reset()         { *m = StoreEditorRequest{} }
func (m *StoreEditorRequest) String() string { return proto.CompactTextString(m) }
func (*StoreEditorRequest) ProtoMessage()    {}
func (*StoreEditorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{6}
}

func (m *StoreEditorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreEditorRequest.Unmarshal(m, b)
}
func (m *StoreEditorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreEditorRequest.Marshal(b, m, deterministic)
}
func (m *StoreEditorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreEditorRequest.Merge(m, src)
}
func (m *StoreEditorRequest) XXX_Size() int {
	return xxx_messageInfo_StoreEditorRequest.Size(m)
}
func (m *StoreEditorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreEditorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreEditorRequest proto.InternalMessageInfo

func (m *StoreEditorRequest) GetEditor() *Editor {
	if m != nil {
		return m.Editor
	}
	return nil
}

type StoreEditorResponse struct {
	Editor               *Editor  `protobuf:"bytes,1,opt,name=editor,proto3" json:"editor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreEditorResponse) Reset()         { *m = StoreEditorResponse{} }
func (m *StoreEditorResponse) String() string { return proto.CompactTextString(m) }
func (*StoreEditorResponse) ProtoMessage()    {}
func (*StoreEditorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60692e730edf4587, []int{7}
}

func (m *StoreEditorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreEditorResponse.Unmarshal(m, b)
}
func (m *StoreEditorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreEditorResponse.Marshal(b, m, deterministic)
}
func (m *StoreEditorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreEditorResponse.Merge(m, src)
}
func (m *StoreEditorResponse) XXX_Size() int {
	return xxx_messageInfo_StoreEditorResponse.Size(m)
}
func (m *StoreEditorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreEditorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreEditorResponse proto.InternalMessageInfo

func (m *StoreEditorResponse) GetEditor() *Editor {
	if m != nil {
		return m.Editor
	}
	return nil
}

func init() {
	proto.RegisterType((*Articles)(nil), "article.Articles")
	proto.RegisterType((*Article)(nil), "article.Article")
	proto.RegisterType((*ArticleOverview)(nil), "article.ArticleOverview")
	proto.RegisterType((*Editor)(nil), "article.Editor")
	proto.RegisterType((*StoreArticleRequest)(nil), "article.StoreArticleRequest")
	proto.RegisterType((*StoreArticleResponse)(nil), "article.StoreArticleResponse")
	proto.RegisterType((*StoreEditorRequest)(nil), "article.StoreEditorRequest")
	proto.RegisterType((*StoreEditorResponse)(nil), "article.StoreEditorResponse")
}

func init() {
	proto.RegisterFile("proto/article.proto", fileDescriptor_60692e730edf4587)
}

var fileDescriptor_60692e730edf4587 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x95, 0x1b, 0x9a, 0xc7, 0x75, 0xd4, 0xc2, 0x6d, 0x51, 0x2d, 0x37, 0x05, 0x6b, 0x36, 0x44,
	0x08, 0xc5, 0x34, 0xb0, 0x40, 0x95, 0xa8, 0x54, 0x21, 0x90, 0x8a, 0x04, 0x48, 0x0e, 0x3f, 0xe0,
	0xd8, 0xd3, 0x30, 0x28, 0xf1, 0x04, 0xcf, 0x24, 0x2c, 0x10, 0x1b, 0xbe, 0x00, 0x89, 0x4f, 0x63,
	0xcb, 0x92, 0x05, 0x9f, 0x81, 0x3c, 0x0f, 0xc7, 0x4e, 0xd2, 0xaa, 0x3b, 0xfb, 0xdc, 0x33, 0xe7,
	0xdc, 0x27, 0x1c, 0xcc, 0x73, 0x2e, 0x79, 0x18, 0xe7, 0x92, 0x25, 0x53, 0x3a, 0x50, 0x7f, 0xd8,
	0x32, 0xbf, 0xfe, 0xc3, 0x09, 0xe7, 0x93, 0x29, 0x0d, 0x15, 0x3c, 0x5e, 0x5c, 0x85, 0x92, 0xcd,
	0xa8, 0x90, 0xf1, 0x6c, 0xae, 0x99, 0x7e, 0xcf, 0x10, 0xe2, 0x39, 0x0b, 0xe3, 0x2c, 0xe3, 0x32,
	0x96, 0x8c, 0x67, 0x62, 0x2d, 0x5a, 0x3e, 0x17, 0x32, 0x5f, 0x24, 0x52, 0x47, 0xc9, 0x0b, 0x68,
	0x5f, 0x68, 0x1f, 0x81, 0x4f, 0xa0, 0x6d, 0x3c, 0x85, 0xe7, 0x04, 0x8d, 0xbe, 0x3b, 0xbc, 0x3b,
	0xb0, 0x39, 0x19, 0x52, 0x54, 0x32, 0xc8, 0x4f, 0x07, 0x5a, 0x06, 0xc5, 0x1e, 0x74, 0x0c, 0x7e,
	0x99, 0x7a, 0x4e, 0xe0, 0xf4, 0x3b, 0xd1, 0x0a, 0xc0, 0xe7, 0xd0, 0xe6, 0x4b, 0x9a, 0x2f, 0x19,
	0xfd, 0xea, 0xed, 0x04, 0x4e, 0xdf, 0x1d, 0x7a, 0xeb, 0xba, 0x1f, 0x4c, 0x3c, 0x2a, 0x99, 0x78,
	0x0a, 0xad, 0x84, 0x67, 0x92, 0x66, 0xd2, 0x6b, 0xa9, 0x47, 0x47, 0x03, 0x5d, 0xc9, 0xc0, 0x56,
	0x32, 0x18, 0xa9, 0x4a, 0x22, 0xcb, 0x23, 0x7f, 0x1c, 0xd8, 0x5f, 0x13, 0xc4, 0x00, 0x5c, 0xe3,
	0xf5, 0x3e, 0x9e, 0x51, 0x93, 0x5c, 0x15, 0x2a, 0x18, 0x29, 0x15, 0x49, 0xce, 0xe6, 0x45, 0xdb,
	0x54, 0x86, 0x9d, 0xa8, 0x0a, 0xe1, 0x23, 0x68, 0xd2, 0x94, 0x49, 0x9e, 0x7b, 0x0d, 0x95, 0xc9,
	0x7e, 0x99, 0xfe, 0x6b, 0x05, 0x47, 0x26, 0x8c, 0xe7, 0xd0, 0x9d, 0xc6, 0x42, 0xbe, 0xe3, 0x29,
	0xbb, 0x62, 0x34, 0xf5, 0xee, 0x28, 0xba, 0xbf, 0x91, 0xf8, 0x47, 0x3b, 0xc1, 0xa8, 0xc6, 0x2f,
	0xfa, 0x28, 0x3f, 0x2d, 0x66, 0xe3, 0x2c, 0x66, 0x53, 0x6f, 0x57, 0xf7, 0xb1, 0x04, 0x48, 0x0a,
	0x4d, 0xed, 0x87, 0x3e, 0xb4, 0xb5, 0xa3, 0x69, 0xf7, 0x6e, 0x54, 0xfe, 0xe3, 0x03, 0x00, 0xfd,
	0xad, 0xea, 0xd5, 0xd5, 0x54, 0x90, 0x55, 0xfc, 0x32, 0xe1, 0x99, 0x2a, 0xa8, 0x8c, 0x17, 0x08,
	0xb9, 0x80, 0x83, 0x91, 0xe4, 0x39, 0xb5, 0x13, 0xa7, 0x5f, 0x16, 0x54, 0x48, 0x7c, 0x0c, 0x76,
	0x21, 0x95, 0xe3, 0xb6, 0xdd, 0xb0, 0x04, 0xf2, 0x14, 0x0e, 0xeb, 0x12, 0x62, 0xce, 0x33, 0x41,
	0xd1, 0x83, 0xd6, 0x8c, 0x0a, 0x11, 0x4f, 0xec, 0x1c, 0xec, 0x2f, 0x79, 0x09, 0xa8, 0x5e, 0x98,
	0x7e, 0x1a, 0xcf, 0x55, 0xdf, 0x9d, 0x1b, 0xfb, 0x4e, 0xce, 0x4d, 0xce, 0xf6, 0xb9, 0xf1, 0xbb,
	0xed, 0xfb, 0xe1, 0xbf, 0x06, 0xec, 0x99, 0x64, 0x47, 0xc5, 0xde, 0x24, 0x14, 0xdf, 0x42, 0xf7,
	0x0d, 0xcb, 0xd2, 0xf2, 0x38, 0xee, 0xad, 0x97, 0x2b, 0xfc, 0x4d, 0x88, 0x1c, 0xfe, 0xf8, 0xfd,
	0xf7, 0xd7, 0xce, 0x1e, 0x76, 0xc3, 0xe5, 0xa9, 0xbd, 0x67, 0x81, 0x23, 0x70, 0x2b, 0x5a, 0xb8,
	0xd1, 0x39, 0x7f, 0x03, 0x21, 0x81, 0x12, 0xf2, 0xd1, 0xab, 0x0a, 0x85, 0xdf, 0xca, 0xa3, 0xfa,
	0x8e, 0x63, 0xc0, 0x8a, 0xe8, 0x2b, 0x7d, 0x02, 0xb7, 0xd2, 0xee, 0x2b, 0x6d, 0x82, 0xc1, 0x75,
	0xda, 0xa1, 0x39, 0x28, 0xfc, 0x0c, 0xdd, 0xea, 0x20, 0xb1, 0x57, 0x6a, 0x6d, 0x59, 0x11, 0xff,
	0xe4, 0x9a, 0xa8, 0x9e, 0x06, 0x39, 0x51, 0xb6, 0x47, 0xa4, 0xd6, 0x9b, 0x33, 0xbb, 0x34, 0x48,
	0xc1, 0xad, 0xcc, 0x10, 0x8f, 0xeb, 0x62, 0xb5, 0xc5, 0xf0, 0x7b, 0xdb, 0x83, 0xc6, 0xe8, 0x58,
	0x19, 0xdd, 0x27, 0x6e, 0x61, 0xa4, 0x27, 0x2c, 0xce, 0xcc, 0xa8, 0xc7, 0x4d, 0x75, 0x84, 0xcf,
	0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xba, 0xc5, 0x53, 0x74, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleServiceClient interface {
	FindArticles(ctx context.Context, in *Articles, opts ...grpc.CallOption) (*Articles, error)
	FindArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error)
	FindArticleContent(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error)
	StoreArticle(ctx context.Context, in *StoreArticleRequest, opts ...grpc.CallOption) (*StoreArticleResponse, error)
	StoreEditor(ctx context.Context, in *StoreEditorRequest, opts ...grpc.CallOption) (*StoreEditorResponse, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) FindArticles(ctx context.Context, in *Articles, opts ...grpc.CallOption) (*Articles, error) {
	out := new(Articles)
	err := c.cc.Invoke(ctx, "/article.ArticleService/FindArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) FindArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/article.ArticleService/FindArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) FindArticleContent(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/article.ArticleService/FindArticleContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) StoreArticle(ctx context.Context, in *StoreArticleRequest, opts ...grpc.CallOption) (*StoreArticleResponse, error) {
	out := new(StoreArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/StoreArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) StoreEditor(ctx context.Context, in *StoreEditorRequest, opts ...grpc.CallOption) (*StoreEditorResponse, error) {
	out := new(StoreEditorResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/StoreEditor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
type ArticleServiceServer interface {
	FindArticles(context.Context, *Articles) (*Articles, error)
	FindArticle(context.Context, *Article) (*Article, error)
	FindArticleContent(context.Context, *Article) (*Article, error)
	StoreArticle(context.Context, *StoreArticleRequest) (*StoreArticleResponse, error)
	StoreEditor(context.Context, *StoreEditorRequest) (*StoreEditorResponse, error)
}

// UnimplementedArticleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (*UnimplementedArticleServiceServer) FindArticles(ctx context.Context, req *Articles) (*Articles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindArticles not implemented")
}
func (*UnimplementedArticleServiceServer) FindArticle(ctx context.Context, req *Article) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindArticle not implemented")
}
func (*UnimplementedArticleServiceServer) FindArticleContent(ctx context.Context, req *Article) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindArticleContent not implemented")
}
func (*UnimplementedArticleServiceServer) StoreArticle(ctx context.Context, req *StoreArticleRequest) (*StoreArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreArticle not implemented")
}
func (*UnimplementedArticleServiceServer) StoreEditor(ctx context.Context, req *StoreEditorRequest) (*StoreEditorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreEditor not implemented")
}

func RegisterArticleServiceServer(s *grpc.Server, srv ArticleServiceServer) {
	s.RegisterService(&_ArticleService_serviceDesc, srv)
}

func _ArticleService_FindArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Articles)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).FindArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/FindArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).FindArticles(ctx, req.(*Articles))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_FindArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).FindArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/FindArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).FindArticle(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_FindArticleContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).FindArticleContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/FindArticleContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).FindArticleContent(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_StoreArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).StoreArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/StoreArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).StoreArticle(ctx, req.(*StoreArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_StoreEditor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreEditorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).StoreEditor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/StoreEditor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).StoreEditor(ctx, req.(*StoreEditorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArticleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "article.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindArticles",
			Handler:    _ArticleService_FindArticles_Handler,
		},
		{
			MethodName: "FindArticle",
			Handler:    _ArticleService_FindArticle_Handler,
		},
		{
			MethodName: "FindArticleContent",
			Handler:    _ArticleService_FindArticleContent_Handler,
		},
		{
			MethodName: "StoreArticle",
			Handler:    _ArticleService_StoreArticle_Handler,
		},
		{
			MethodName: "StoreEditor",
			Handler:    _ArticleService_StoreEditor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/article.proto",
}
