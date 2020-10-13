// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: proto/product/product.proto

package product

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nom         string `protobuf:"bytes,1,opt,name=nom,proto3" json:"nom,omitempty"`
	Prenom      string `protobuf:"bytes,2,opt,name=prenom,proto3" json:"prenom,omitempty"`
	Dateofbirth string `protobuf:"bytes,3,opt,name=dateofbirth,proto3" json:"dateofbirth,omitempty"`
	Abjcardno   string `protobuf:"bytes,4,opt,name=abjcardno,proto3" json:"abjcardno,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetNom() string {
	if x != nil {
		return x.Nom
	}
	return ""
}

func (x *Client) GetPrenom() string {
	if x != nil {
		return x.Prenom
	}
	return ""
}

func (x *Client) GetDateofbirth() string {
	if x != nil {
		return x.Dateofbirth
	}
	return ""
}

func (x *Client) GetAbjcardno() string {
	if x != nil {
		return x.Abjcardno
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Police struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Police string `protobuf:"bytes,1,opt,name=police,proto3" json:"police,omitempty"`
}

func (x *Police) Reset() {
	*x = Police{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Police) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Police) ProtoMessage() {}

func (x *Police) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Police.ProtoReflect.Descriptor instead.
func (*Police) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *Police) GetPolice() string {
	if x != nil {
		return x.Police
	}
	return ""
}

type Etat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Police          string `protobuf:"bytes,1,opt,name=police,proto3" json:"police,omitempty"`
	Datedebuteffet  string `protobuf:"bytes,2,opt,name=datedebuteffet,proto3" json:"datedebuteffet,omitempty"`
	Datefineffet    string `protobuf:"bytes,3,opt,name=datefineffet,proto3" json:"datefineffet,omitempty"`
	Libelleproduit  string `protobuf:"bytes,4,opt,name=libelleproduit,proto3" json:"libelleproduit,omitempty"`
	Modereglement   string `protobuf:"bytes,5,opt,name=modereglement,proto3" json:"modereglement,omitempty"`
	Fractionnement  string `protobuf:"bytes,6,opt,name=fractionnement,proto3" json:"fractionnement,omitempty"`
	Numeropayeur    string `protobuf:"bytes,7,opt,name=numeropayeur,proto3" json:"numeropayeur,omitempty"`
	Nompayeur       string `protobuf:"bytes,8,opt,name=nompayeur,proto3" json:"nompayeur,omitempty"`
	Telephone       string `protobuf:"bytes,9,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Profession      string `protobuf:"bytes,10,opt,name=profession,proto3" json:"profession,omitempty"`
	Adresse         string `protobuf:"bytes,11,opt,name=adresse,proto3" json:"adresse,omitempty"`
	Datenaissance   string `protobuf:"bytes,12,opt,name=datenaissance,proto3" json:"datenaissance,omitempty"`
	Lieunaissance   string `protobuf:"bytes,13,opt,name=lieunaissance,proto3" json:"lieunaissance,omitempty"`
	Nomsouscripteur string `protobuf:"bytes,14,opt,name=nomsouscripteur,proto3" json:"nomsouscripteur,omitempty"`
	Quittance       string `protobuf:"bytes,15,opt,name=quittance,proto3" json:"quittance,omitempty"`
	Prime           string `protobuf:"bytes,16,opt,name=prime,proto3" json:"prime,omitempty"`
	Datecomptable   string `protobuf:"bytes,17,opt,name=datecomptable,proto3" json:"datecomptable,omitempty"`
	Datequittance   string `protobuf:"bytes,18,opt,name=datequittance,proto3" json:"datequittance,omitempty"`
	Etatquittance   string `protobuf:"bytes,19,opt,name=etatquittance,proto3" json:"etatquittance,omitempty"`
	Montantsolde    string `protobuf:"bytes,20,opt,name=montantsolde,proto3" json:"montantsolde,omitempty"`
	Montantimpaye   string `protobuf:"bytes,21,opt,name=montantimpaye,proto3" json:"montantimpaye,omitempty"`
	Nombresolde     string `protobuf:"bytes,22,opt,name=nombresolde,proto3" json:"nombresolde,omitempty"`
	Nombreimpaye    string `protobuf:"bytes,23,opt,name=nombreimpaye,proto3" json:"nombreimpaye,omitempty"`
}

func (x *Etat) Reset() {
	*x = Etat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Etat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Etat) ProtoMessage() {}

func (x *Etat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Etat.ProtoReflect.Descriptor instead.
func (*Etat) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *Etat) GetPolice() string {
	if x != nil {
		return x.Police
	}
	return ""
}

func (x *Etat) GetDatedebuteffet() string {
	if x != nil {
		return x.Datedebuteffet
	}
	return ""
}

func (x *Etat) GetDatefineffet() string {
	if x != nil {
		return x.Datefineffet
	}
	return ""
}

func (x *Etat) GetLibelleproduit() string {
	if x != nil {
		return x.Libelleproduit
	}
	return ""
}

func (x *Etat) GetModereglement() string {
	if x != nil {
		return x.Modereglement
	}
	return ""
}

func (x *Etat) GetFractionnement() string {
	if x != nil {
		return x.Fractionnement
	}
	return ""
}

func (x *Etat) GetNumeropayeur() string {
	if x != nil {
		return x.Numeropayeur
	}
	return ""
}

func (x *Etat) GetNompayeur() string {
	if x != nil {
		return x.Nompayeur
	}
	return ""
}

func (x *Etat) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *Etat) GetProfession() string {
	if x != nil {
		return x.Profession
	}
	return ""
}

func (x *Etat) GetAdresse() string {
	if x != nil {
		return x.Adresse
	}
	return ""
}

func (x *Etat) GetDatenaissance() string {
	if x != nil {
		return x.Datenaissance
	}
	return ""
}

func (x *Etat) GetLieunaissance() string {
	if x != nil {
		return x.Lieunaissance
	}
	return ""
}

func (x *Etat) GetNomsouscripteur() string {
	if x != nil {
		return x.Nomsouscripteur
	}
	return ""
}

func (x *Etat) GetQuittance() string {
	if x != nil {
		return x.Quittance
	}
	return ""
}

func (x *Etat) GetPrime() string {
	if x != nil {
		return x.Prime
	}
	return ""
}

func (x *Etat) GetDatecomptable() string {
	if x != nil {
		return x.Datecomptable
	}
	return ""
}

func (x *Etat) GetDatequittance() string {
	if x != nil {
		return x.Datequittance
	}
	return ""
}

func (x *Etat) GetEtatquittance() string {
	if x != nil {
		return x.Etatquittance
	}
	return ""
}

func (x *Etat) GetMontantsolde() string {
	if x != nil {
		return x.Montantsolde
	}
	return ""
}

func (x *Etat) GetMontantimpaye() string {
	if x != nil {
		return x.Montantimpaye
	}
	return ""
}

func (x *Etat) GetNombresolde() string {
	if x != nil {
		return x.Nombresolde
	}
	return ""
}

func (x *Etat) GetNombreimpaye() string {
	if x != nil {
		return x.Nombreimpaye
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{4}
}

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nom         string `protobuf:"bytes,1,opt,name=nom,proto3" json:"nom,omitempty"`
	Prenom      string `protobuf:"bytes,2,opt,name=prenom,proto3" json:"prenom,omitempty"`
	Dateofbirth string `protobuf:"bytes,3,opt,name=dateofbirth,proto3" json:"dateofbirth,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{5}
}

func (x *ClientRequest) GetNom() string {
	if x != nil {
		return x.Nom
	}
	return ""
}

func (x *ClientRequest) GetPrenom() string {
	if x != nil {
		return x.Prenom
	}
	return ""
}

func (x *ClientRequest) GetDateofbirth() string {
	if x != nil {
		return x.Dateofbirth
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product  *Product   `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Etat     string     `protobuf:"bytes,2,opt,name=etat,proto3" json:"etat,omitempty"`
	Products []*Product `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	Errors   []*Error   `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{6}
}

func (x *Response) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *Response) GetEtat() string {
	if x != nil {
		return x.Etat
	}
	return ""
}

func (x *Response) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Response) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{7}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid  bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{8}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *Token) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_proto_product_product_proto protoreflect.FileDescriptor

var file_proto_product_product_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x72, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e,
	0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61,
	0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x61, 0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x62, 0x6a, 0x63, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x62, 0x6a, 0x63, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x22, 0x2d, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x22, 0xa6, 0x06, 0x0a, 0x04,
	0x45, 0x74, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x65, 0x62, 0x75, 0x74, 0x65, 0x66, 0x66, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x65, 0x64, 0x65, 0x62, 0x75, 0x74, 0x65,
	0x66, 0x66, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x66, 0x66, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x66, 0x66, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x69, 0x62, 0x65,
	0x6c, 0x6c, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6c, 0x69, 0x62, 0x65, 0x6c, 0x6c, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x67, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x67,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x70, 0x61, 0x79, 0x65, 0x75, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x70, 0x61, 0x79, 0x65,
	0x75, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x65, 0x75, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x65, 0x75, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65,
	0x6e, 0x61, 0x69, 0x73, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x61, 0x74, 0x65, 0x6e, 0x61, 0x69, 0x73, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x6c, 0x69, 0x65, 0x75, 0x6e, 0x61, 0x69, 0x73, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x65, 0x75, 0x6e, 0x61, 0x69, 0x73, 0x73,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x6f, 0x6d, 0x73, 0x6f, 0x75, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x65, 0x75, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e,
	0x6f, 0x6d, 0x73, 0x6f, 0x75, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x65, 0x75, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x71, 0x75, 0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x71, 0x75, 0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x63,
	0x6f, 0x6d, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65,
	0x71, 0x75, 0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x61, 0x74, 0x65, 0x71, 0x75, 0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x65, 0x74, 0x61, 0x74, 0x71, 0x75, 0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x74, 0x61, 0x74, 0x71, 0x75, 0x69, 0x74, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x61, 0x6e, 0x74, 0x73,
	0x6f, 0x6c, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x74,
	0x61, 0x6e, 0x74, 0x73, 0x6f, 0x6c, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6d, 0x6f, 0x6e, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x64, 0x65, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x64, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x65,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x69, 0x6d,
	0x70, 0x61, 0x79, 0x65, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x5b, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e,
	0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61,
	0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x61, 0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x22, 0xa0, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22,
	0x3d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5b,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0xa1, 0x02, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x74, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65,
	0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x12, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x1a,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product_product_proto_rawDescOnce sync.Once
	file_proto_product_product_proto_rawDescData = file_proto_product_product_proto_rawDesc
)

func file_proto_product_product_proto_rawDescGZIP() []byte {
	file_proto_product_product_proto_rawDescOnce.Do(func() {
		file_proto_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product_product_proto_rawDescData)
	})
	return file_proto_product_product_proto_rawDescData
}

var file_proto_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_product_product_proto_goTypes = []interface{}{
	(*Client)(nil),        // 0: product.Client
	(*Product)(nil),       // 1: product.Product
	(*Police)(nil),        // 2: product.Police
	(*Etat)(nil),          // 3: product.Etat
	(*Request)(nil),       // 4: product.Request
	(*ClientRequest)(nil), // 5: product.ClientRequest
	(*Response)(nil),      // 6: product.Response
	(*Error)(nil),         // 7: product.Error
	(*Token)(nil),         // 8: product.Token
}
var file_proto_product_product_proto_depIdxs = []int32{
	1, // 0: product.Response.product:type_name -> product.Product
	1, // 1: product.Response.products:type_name -> product.Product
	7, // 2: product.Response.errors:type_name -> product.Error
	7, // 3: product.Token.errors:type_name -> product.Error
	1, // 4: product.ProductService.Get:input_type -> product.Product
	4, // 5: product.ProductService.GetAll:input_type -> product.Request
	2, // 6: product.ProductService.GetCotisations:input_type -> product.Police
	2, // 7: product.ProductService.GetlistePoliceExterne:input_type -> product.Police
	0, // 8: product.ProductService.GetClientProducts:input_type -> product.Client
	6, // 9: product.ProductService.Get:output_type -> product.Response
	6, // 10: product.ProductService.GetAll:output_type -> product.Response
	6, // 11: product.ProductService.GetCotisations:output_type -> product.Response
	6, // 12: product.ProductService.GetlistePoliceExterne:output_type -> product.Response
	6, // 13: product.ProductService.GetClientProducts:output_type -> product.Response
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_product_product_proto_init() }
func file_proto_product_product_proto_init() {
	if File_proto_product_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_product_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_proto_product_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_proto_product_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Police); i {
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
		file_proto_product_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Etat); i {
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
		file_proto_product_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_product_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
		file_proto_product_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_product_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_proto_product_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_proto_product_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_product_proto_goTypes,
		DependencyIndexes: file_proto_product_product_proto_depIdxs,
		MessageInfos:      file_proto_product_product_proto_msgTypes,
	}.Build()
	File_proto_product_product_proto = out.File
	file_proto_product_product_proto_rawDesc = nil
	file_proto_product_product_proto_goTypes = nil
	file_proto_product_product_proto_depIdxs = nil
}
