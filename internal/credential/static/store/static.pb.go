// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: controller/storage/credential/static/store/v1/static.proto

// Package store provides protobufs for storing types in the static
// credential package.

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
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

type CredentialStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// name is optional. If set, it must be unique within scope_id.
	// @inject_tag: `gorm:"default:null"`
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" gorm:"default:null"`
	// description is optional.
	// @inject_tag: `gorm:"default:null"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" gorm:"default:null"`
	// The scope_id of the owning scope.
	// It must be set.
	// @inject_tag: `gorm:"not_null"`
	ScopeId string `protobuf:"bytes,6,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" gorm:"not_null"`
	// version allows optimistic locking of the resource.
	// @inject_tag: `gorm:"default:null"`
	Version uint32 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
}

func (x *CredentialStore) Reset() {
	*x = CredentialStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_credential_static_store_v1_static_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialStore) ProtoMessage() {}

func (x *CredentialStore) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_credential_static_store_v1_static_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialStore.ProtoReflect.Descriptor instead.
func (*CredentialStore) Descriptor() ([]byte, []int) {
	return file_controller_storage_credential_static_store_v1_static_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialStore) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *CredentialStore) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CredentialStore) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *CredentialStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CredentialStore) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CredentialStore) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *CredentialStore) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type UsernamePasswordCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// name is optional. If set, it must be unique within scope_id.
	// @inject_tag: `gorm:"default:null"`
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" gorm:"default:null"`
	// description is optional.
	// @inject_tag: `gorm:"default:null"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" gorm:"default:null"`
	// store_id of the owning static credential store.
	// It must be set.
	// @inject_tag: `gorm:"not_null"`
	StoreId string `protobuf:"bytes,6,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty" gorm:"not_null"`
	// version allows optimistic locking of the resource.
	// @inject_tag: `gorm:"default:null"`
	Version uint32 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
	// username is the username associated with the credential.
	// It must be set.
	// @inject_tag: `gorm:"not_null"`
	Username string `protobuf:"bytes,8,opt,name=username,proto3" json:"username,omitempty" gorm:"not_null"`
	// password is the plain-text of the password associated with the credential. We are
	// not storing this plain-text password in the database.
	// @inject_tag: `gorm:"-" wrapping:"pt,password_data"`
	Password []byte `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty" gorm:"-" wrapping:"pt,password_data"`
	// ct_password is the ciphertext of the password. It
	// is stored in the database.
	// @inject_tag: `gorm:"column:password_encrypted;not_null" wrapping:"ct,password_data"`
	CtPassword []byte `protobuf:"bytes,10,opt,name=ct_password,json=ctPassword,proto3" json:"ct_password,omitempty" gorm:"column:password_encrypted;not_null" wrapping:"ct,password_data"`
	// password_hmac is a sha256-hmac of the unencrypted password.  It is recalculated
	// everytime the password is updated.
	// @inject_tag: `gorm:"not_null"`
	PasswordHmac []byte `protobuf:"bytes,11,opt,name=password_hmac,json=passwordHmac,proto3" json:"password_hmac,omitempty" gorm:"not_null"`
	// The key_id of the kms database key used for encrypting this entry.
	// It must be set.
	// @inject_tag: `gorm:"not_null"`
	KeyId string `protobuf:"bytes,12,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty" gorm:"not_null"`
}

func (x *UsernamePasswordCredential) Reset() {
	*x = UsernamePasswordCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_credential_static_store_v1_static_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsernamePasswordCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernamePasswordCredential) ProtoMessage() {}

func (x *UsernamePasswordCredential) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_credential_static_store_v1_static_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernamePasswordCredential.ProtoReflect.Descriptor instead.
func (*UsernamePasswordCredential) Descriptor() ([]byte, []int) {
	return file_controller_storage_credential_static_store_v1_static_proto_rawDescGZIP(), []int{1}
}

func (x *UsernamePasswordCredential) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *UsernamePasswordCredential) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *UsernamePasswordCredential) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *UsernamePasswordCredential) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UsernamePasswordCredential) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UsernamePasswordCredential) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *UsernamePasswordCredential) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UsernamePasswordCredential) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UsernamePasswordCredential) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *UsernamePasswordCredential) GetCtPassword() []byte {
	if x != nil {
		return x.CtPassword
	}
	return nil
}

func (x *UsernamePasswordCredential) GetPasswordHmac() []byte {
	if x != nil {
		return x.PasswordHmac
	}
	return nil
}

func (x *UsernamePasswordCredential) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

var File_controller_storage_credential_static_store_v1_static_proto protoreflect.FileDescriptor

var file_controller_storage_credential_static_store_v1_static_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x10, 0xc2, 0xdd, 0x29, 0x0c, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e,
	0xc2, 0xdd, 0x29, 0x1a, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0xfd, 0x04, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xc2, 0xdd, 0x29, 0x0c, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1e, 0xc2, 0xdd, 0x29, 0x1a, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xdd, 0x29, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x23, 0xc2, 0xdd, 0x29, 0x1f, 0x0a, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x51, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x2c, 0xc2,
	0xdd, 0x29, 0x28, 0x0a, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x6d, 0x61,
	0x63, 0x12, 0x18, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x52, 0x0c, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x6d, 0x61, 0x63, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64,
	0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_credential_static_store_v1_static_proto_rawDescOnce sync.Once
	file_controller_storage_credential_static_store_v1_static_proto_rawDescData = file_controller_storage_credential_static_store_v1_static_proto_rawDesc
)

func file_controller_storage_credential_static_store_v1_static_proto_rawDescGZIP() []byte {
	file_controller_storage_credential_static_store_v1_static_proto_rawDescOnce.Do(func() {
		file_controller_storage_credential_static_store_v1_static_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_credential_static_store_v1_static_proto_rawDescData)
	})
	return file_controller_storage_credential_static_store_v1_static_proto_rawDescData
}

var file_controller_storage_credential_static_store_v1_static_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_storage_credential_static_store_v1_static_proto_goTypes = []interface{}{
	(*CredentialStore)(nil),            // 0: controller.storage.credential.static.store.v1.CredentialStore
	(*UsernamePasswordCredential)(nil), // 1: controller.storage.credential.static.store.v1.UsernamePasswordCredential
	(*timestamp.Timestamp)(nil),        // 2: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_credential_static_store_v1_static_proto_depIdxs = []int32{
	2, // 0: controller.storage.credential.static.store.v1.CredentialStore.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 1: controller.storage.credential.static.store.v1.CredentialStore.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 2: controller.storage.credential.static.store.v1.UsernamePasswordCredential.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 3: controller.storage.credential.static.store.v1.UsernamePasswordCredential.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_controller_storage_credential_static_store_v1_static_proto_init() }
func file_controller_storage_credential_static_store_v1_static_proto_init() {
	if File_controller_storage_credential_static_store_v1_static_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_credential_static_store_v1_static_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialStore); i {
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
		file_controller_storage_credential_static_store_v1_static_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsernamePasswordCredential); i {
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
			RawDescriptor: file_controller_storage_credential_static_store_v1_static_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_credential_static_store_v1_static_proto_goTypes,
		DependencyIndexes: file_controller_storage_credential_static_store_v1_static_proto_depIdxs,
		MessageInfos:      file_controller_storage_credential_static_store_v1_static_proto_msgTypes,
	}.Build()
	File_controller_storage_credential_static_store_v1_static_proto = out.File
	file_controller_storage_credential_static_store_v1_static_proto_rawDesc = nil
	file_controller_storage_credential_static_store_v1_static_proto_goTypes = nil
	file_controller_storage_credential_static_store_v1_static_proto_depIdxs = nil
}
