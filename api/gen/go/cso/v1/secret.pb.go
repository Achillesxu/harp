// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: cso/v1/secret.proto

package csov1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// RingLevel enumerates all cso ring level values.
type RingLevel int32

const (
	// Default value when no enumeration is specified.
	RingLevel_RING_LEVEL_INVALID RingLevel = 0
	// Explicitly Unknown object value.
	RingLevel_RING_LEVEL_UNKNOWN RingLevel = 1
	// Defines secret used by secrets.
	RingLevel_RING_LEVEL_META RingLevel = 2
	// Defines infrastructure level secrets.
	RingLevel_RING_LEVEL_INFRASTRUCTURE RingLevel = 3
	// Defines platform level secrets.
	RingLevel_RING_LEVEL_PLATFORM RingLevel = 4
	// Defines product level secrets.
	RingLevel_RING_LEVEL_PRODUCT RingLevel = 5
	// Defines application level secrets.
	RingLevel_RING_LEVEL_APPLICATION RingLevel = 6
	// Defines artifact level secrets.
	RingLevel_RING_LEVEL_ARTIFACT RingLevel = 7
)

// Enum value maps for RingLevel.
var (
	RingLevel_name = map[int32]string{
		0: "RING_LEVEL_INVALID",
		1: "RING_LEVEL_UNKNOWN",
		2: "RING_LEVEL_META",
		3: "RING_LEVEL_INFRASTRUCTURE",
		4: "RING_LEVEL_PLATFORM",
		5: "RING_LEVEL_PRODUCT",
		6: "RING_LEVEL_APPLICATION",
		7: "RING_LEVEL_ARTIFACT",
	}
	RingLevel_value = map[string]int32{
		"RING_LEVEL_INVALID":        0,
		"RING_LEVEL_UNKNOWN":        1,
		"RING_LEVEL_META":           2,
		"RING_LEVEL_INFRASTRUCTURE": 3,
		"RING_LEVEL_PLATFORM":       4,
		"RING_LEVEL_PRODUCT":        5,
		"RING_LEVEL_APPLICATION":    6,
		"RING_LEVEL_ARTIFACT":       7,
	}
)

func (x RingLevel) Enum() *RingLevel {
	p := new(RingLevel)
	*p = x
	return p
}

func (x RingLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RingLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_cso_v1_secret_proto_enumTypes[0].Descriptor()
}

func (RingLevel) Type() protoreflect.EnumType {
	return &file_cso_v1_secret_proto_enumTypes[0]
}

func (x RingLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RingLevel.Descriptor instead.
func (RingLevel) EnumDescriptor() ([]byte, []int) {
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{0}
}

// QualityLevel enumerates entity quality level values.
type QualityLevel int32

const (
	// Default value when no enumeration is specified.
	QualityLevel_QUALITY_LEVEL_INVALID QualityLevel = 0
	// Explicitly Unknown object value.
	QualityLevel_QUALITY_LEVEL_UNKNOWN QualityLevel = 1
	// Production grade
	QualityLevel_QUALITY_LEVEL_PRODUCTION QualityLevel = 2
	// Staging grade
	QualityLevel_QUALITY_LEVEL_STAGING QualityLevel = 3
	// QA Grade
	QualityLevel_QUALITY_LEVEL_QA QualityLevel = 4
	// Dev grade
	QualityLevel_QUALITY_LEVEL_DEV QualityLevel = 5
)

// Enum value maps for QualityLevel.
var (
	QualityLevel_name = map[int32]string{
		0: "QUALITY_LEVEL_INVALID",
		1: "QUALITY_LEVEL_UNKNOWN",
		2: "QUALITY_LEVEL_PRODUCTION",
		3: "QUALITY_LEVEL_STAGING",
		4: "QUALITY_LEVEL_QA",
		5: "QUALITY_LEVEL_DEV",
	}
	QualityLevel_value = map[string]int32{
		"QUALITY_LEVEL_INVALID":    0,
		"QUALITY_LEVEL_UNKNOWN":    1,
		"QUALITY_LEVEL_PRODUCTION": 2,
		"QUALITY_LEVEL_STAGING":    3,
		"QUALITY_LEVEL_QA":         4,
		"QUALITY_LEVEL_DEV":        5,
	}
)

func (x QualityLevel) Enum() *QualityLevel {
	p := new(QualityLevel)
	*p = x
	return p
}

func (x QualityLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QualityLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_cso_v1_secret_proto_enumTypes[1].Descriptor()
}

func (QualityLevel) Type() protoreflect.EnumType {
	return &file_cso_v1_secret_proto_enumTypes[1]
}

func (x QualityLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QualityLevel.Descriptor instead.
func (QualityLevel) EnumDescriptor() ([]byte, []int) {
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{1}
}

// Secret represents secret value and metadata.
type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RingLevel RingLevel `protobuf:"varint,1,opt,name=ring_level,json=ringLevel,proto3,enum=cso.v1.RingLevel" json:"ring_level,omitempty"`
	Value     *Value    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Types that are assignable to Path:
	//	*Secret_Meta
	//	*Secret_Infrastructure
	//	*Secret_Platform
	//	*Secret_Product
	//	*Secret_Application
	//	*Secret_Artifact
	Path isSecret_Path `protobuf_oneof:"path"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{0}
}

func (x *Secret) GetRingLevel() RingLevel {
	if x != nil {
		return x.RingLevel
	}
	return RingLevel_RING_LEVEL_INVALID
}

func (x *Secret) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (m *Secret) GetPath() isSecret_Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (x *Secret) GetMeta() *Meta {
	if x, ok := x.GetPath().(*Secret_Meta); ok {
		return x.Meta
	}
	return nil
}

func (x *Secret) GetInfrastructure() *Infrastructure {
	if x, ok := x.GetPath().(*Secret_Infrastructure); ok {
		return x.Infrastructure
	}
	return nil
}

func (x *Secret) GetPlatform() *Platform {
	if x, ok := x.GetPath().(*Secret_Platform); ok {
		return x.Platform
	}
	return nil
}

func (x *Secret) GetProduct() *Product {
	if x, ok := x.GetPath().(*Secret_Product); ok {
		return x.Product
	}
	return nil
}

func (x *Secret) GetApplication() *Application {
	if x, ok := x.GetPath().(*Secret_Application); ok {
		return x.Application
	}
	return nil
}

func (x *Secret) GetArtifact() *Artifact {
	if x, ok := x.GetPath().(*Secret_Artifact); ok {
		return x.Artifact
	}
	return nil
}

type isSecret_Path interface {
	isSecret_Path()
}

type Secret_Meta struct {
	Meta *Meta `protobuf:"bytes,10,opt,name=meta,proto3,oneof"`
}

type Secret_Infrastructure struct {
	Infrastructure *Infrastructure `protobuf:"bytes,11,opt,name=infrastructure,proto3,oneof"`
}

type Secret_Platform struct {
	Platform *Platform `protobuf:"bytes,12,opt,name=platform,proto3,oneof"`
}

type Secret_Product struct {
	Product *Product `protobuf:"bytes,13,opt,name=product,proto3,oneof"`
}

type Secret_Application struct {
	Application *Application `protobuf:"bytes,14,opt,name=application,proto3,oneof"`
}

type Secret_Artifact struct {
	Artifact *Artifact `protobuf:"bytes,15,opt,name=artifact,proto3,oneof"`
}

func (*Secret_Meta) isSecret_Path() {}

func (*Secret_Infrastructure) isSecret_Path() {}

func (*Secret_Platform) isSecret_Path() {}

func (*Secret_Product) isSecret_Path() {}

func (*Secret_Application) isSecret_Path() {}

func (*Secret_Artifact) isSecret_Path() {}

// Value represents an encoded secret value.
type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Value) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// Meta describes secrets of secrets path components.
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_secret_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{2}
}

func (x *Meta) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Infrastructure describes infrastructure secret path components.
type Infrastructure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cloud provider name
	CloudProvider string `protobuf:"bytes,1,opt,name=cloud_provider,json=cloudProvider,proto3" json:"cloud_provider,omitempty"`
	// Cloud provider account identifier or alias
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// Cloud provider region
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// Service name used
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Key is the free part of the namming specification.
	Key string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Infrastructure) Reset() {
	*x = Infrastructure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_secret_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Infrastructure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Infrastructure) ProtoMessage() {}

func (x *Infrastructure) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_secret_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Infrastructure.ProtoReflect.Descriptor instead.
func (*Infrastructure) Descriptor() ([]byte, []int) {
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{3}
}

func (x *Infrastructure) GetCloudProvider() string {
	if x != nil {
		return x.CloudProvider
	}
	return ""
}

func (x *Infrastructure) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Infrastructure) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Infrastructure) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Infrastructure) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Platform describes platform secret path components.
type Platform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Quality level
	Stage QualityLevel `protobuf:"varint,1,opt,name=stage,proto3,enum=cso.v1.QualityLevel" json:"stage,omitempty"`
	// Paltform name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Platform region
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// Platform service name
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Key is the free part of the namming specification.
	Key string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Platform) Reset() {
	*x = Platform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_secret_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_secret_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{4}
}

func (x *Platform) GetStage() QualityLevel {
	if x != nil {
		return x.Stage
	}
	return QualityLevel_QUALITY_LEVEL_INVALID
}

func (x *Platform) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Platform) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Platform) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Platform) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Product describes product secret path components.
type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Product name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Product version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Product component name
	ComponentName string `protobuf:"bytes,3,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	// Key is the free part of the namming specification.
	Key string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_secret_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_secret_proto_msgTypes[5]
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
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{5}
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Product) GetComponentName() string {
	if x != nil {
		return x.ComponentName
	}
	return ""
}

func (x *Product) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Application describes application secret path components.
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Quality level
	Stage QualityLevel `protobuf:"varint,1,opt,name=stage,proto3,enum=cso.v1.QualityLevel" json:"stage,omitempty"`
	// Platform name
	PlatformName string `protobuf:"bytes,2,opt,name=platform_name,json=platformName,proto3" json:"platform_name,omitempty"`
	// Product name
	ProductName string `protobuf:"bytes,3,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	// Product version
	ProductVersion string `protobuf:"bytes,4,opt,name=product_version,json=productVersion,proto3" json:"product_version,omitempty"`
	// Product component name
	ComponentName string `protobuf:"bytes,5,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	// Key is the free part of the namming specification.
	Key string `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_secret_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_secret_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{6}
}

func (x *Application) GetStage() QualityLevel {
	if x != nil {
		return x.Stage
	}
	return QualityLevel_QUALITY_LEVEL_INVALID
}

func (x *Application) GetPlatformName() string {
	if x != nil {
		return x.PlatformName
	}
	return ""
}

func (x *Application) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Application) GetProductVersion() string {
	if x != nil {
		return x.ProductVersion
	}
	return ""
}

func (x *Application) GetComponentName() string {
	if x != nil {
		return x.ComponentName
	}
	return ""
}

func (x *Application) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Artifact describes artifact secret path components.
type Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Artifact type
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Artifact id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Key is the free part of the namming specification.
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Artifact) Reset() {
	*x = Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_secret_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifact) ProtoMessage() {}

func (x *Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_secret_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifact.ProtoReflect.Descriptor instead.
func (*Artifact) Descriptor() ([]byte, []int) {
	return file_cso_v1_secret_proto_rawDescGZIP(), []int{7}
}

func (x *Artifact) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Artifact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Artifact) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_cso_v1_secret_proto protoreflect.FileDescriptor

var file_cso_v1_secret_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x73, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x93, 0x03,
	0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63,
	0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x09, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x73, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x22, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x73,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x08, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x48,
	0x00, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x22, 0x2f, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0x18, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xa3,
	0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x97, 0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x70,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0xe3, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x40, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x2a, 0xd5, 0x01, 0x0a, 0x09, 0x52, 0x69, 0x6e,
	0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x16,
	0x0a, 0x12, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x52,
	0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x52, 0x41, 0x53,
	0x54, 0x52, 0x55, 0x43, 0x54, 0x55, 0x52, 0x45, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x49,
	0x4e, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x52,
	0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x49, 0x4e, 0x47, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x46, 0x41, 0x43, 0x54, 0x10, 0x07,
	0x2a, 0xaa, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x19, 0x0a, 0x15, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15,
	0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x51, 0x55, 0x41, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x14, 0x0a, 0x10, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x51, 0x41, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x56, 0x10, 0x05, 0x42, 0x7c, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x65, 0x6c, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x65, 0x63, 0x2e, 0x63, 0x73, 0x6f,
	0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x68, 0x61, 0x72, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x73, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x73,
	0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x43, 0x73, 0x6f, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x06, 0x43, 0x73, 0x6f, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cso_v1_secret_proto_rawDescOnce sync.Once
	file_cso_v1_secret_proto_rawDescData = file_cso_v1_secret_proto_rawDesc
)

func file_cso_v1_secret_proto_rawDescGZIP() []byte {
	file_cso_v1_secret_proto_rawDescOnce.Do(func() {
		file_cso_v1_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_cso_v1_secret_proto_rawDescData)
	})
	return file_cso_v1_secret_proto_rawDescData
}

var (
	file_cso_v1_secret_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
	file_cso_v1_secret_proto_msgTypes  = make([]protoimpl.MessageInfo, 8)
	file_cso_v1_secret_proto_goTypes   = []interface{}{
		(RingLevel)(0),         // 0: cso.v1.RingLevel
		(QualityLevel)(0),      // 1: cso.v1.QualityLevel
		(*Secret)(nil),         // 2: cso.v1.Secret
		(*Value)(nil),          // 3: cso.v1.Value
		(*Meta)(nil),           // 4: cso.v1.Meta
		(*Infrastructure)(nil), // 5: cso.v1.Infrastructure
		(*Platform)(nil),       // 6: cso.v1.Platform
		(*Product)(nil),        // 7: cso.v1.Product
		(*Application)(nil),    // 8: cso.v1.Application
		(*Artifact)(nil),       // 9: cso.v1.Artifact
	}
)

var file_cso_v1_secret_proto_depIdxs = []int32{
	0,  // 0: cso.v1.Secret.ring_level:type_name -> cso.v1.RingLevel
	3,  // 1: cso.v1.Secret.value:type_name -> cso.v1.Value
	4,  // 2: cso.v1.Secret.meta:type_name -> cso.v1.Meta
	5,  // 3: cso.v1.Secret.infrastructure:type_name -> cso.v1.Infrastructure
	6,  // 4: cso.v1.Secret.platform:type_name -> cso.v1.Platform
	7,  // 5: cso.v1.Secret.product:type_name -> cso.v1.Product
	8,  // 6: cso.v1.Secret.application:type_name -> cso.v1.Application
	9,  // 7: cso.v1.Secret.artifact:type_name -> cso.v1.Artifact
	1,  // 8: cso.v1.Platform.stage:type_name -> cso.v1.QualityLevel
	1,  // 9: cso.v1.Application.stage:type_name -> cso.v1.QualityLevel
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cso_v1_secret_proto_init() }
func file_cso_v1_secret_proto_init() {
	if File_cso_v1_secret_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cso_v1_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_cso_v1_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_cso_v1_secret_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_cso_v1_secret_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Infrastructure); i {
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
		file_cso_v1_secret_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform); i {
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
		file_cso_v1_secret_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_cso_v1_secret_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_cso_v1_secret_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifact); i {
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
	file_cso_v1_secret_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Secret_Meta)(nil),
		(*Secret_Infrastructure)(nil),
		(*Secret_Platform)(nil),
		(*Secret_Product)(nil),
		(*Secret_Application)(nil),
		(*Secret_Artifact)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cso_v1_secret_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cso_v1_secret_proto_goTypes,
		DependencyIndexes: file_cso_v1_secret_proto_depIdxs,
		EnumInfos:         file_cso_v1_secret_proto_enumTypes,
		MessageInfos:      file_cso_v1_secret_proto_msgTypes,
	}.Build()
	File_cso_v1_secret_proto = out.File
	file_cso_v1_secret_proto_rawDesc = nil
	file_cso_v1_secret_proto_goTypes = nil
	file_cso_v1_secret_proto_depIdxs = nil
}
