// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: job_service.proto

package job

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type OwnerJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerKey string `protobuf:"bytes,1,opt,name=ownerKey,proto3" json:"ownerKey,omitempty"`
}

func (x *OwnerJob) Reset() {
	*x = OwnerJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnerJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnerJob) ProtoMessage() {}

func (x *OwnerJob) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnerJob.ProtoReflect.Descriptor instead.
func (*OwnerJob) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{0}
}

func (x *OwnerJob) GetOwnerKey() string {
	if x != nil {
		return x.OwnerKey
	}
	return ""
}

type OwnerJobOffersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *OwnerJob `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *OwnerJobOffersRequest) Reset() {
	*x = OwnerJobOffersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnerJobOffersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnerJobOffersRequest) ProtoMessage() {}

func (x *OwnerJobOffersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnerJobOffersRequest.ProtoReflect.Descriptor instead.
func (*OwnerJobOffersRequest) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{1}
}

func (x *OwnerJobOffersRequest) GetKey() *OwnerJob {
	if x != nil {
		return x.Key
	}
	return nil
}

type AddKeyJobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AddKeyJobOffer) Reset() {
	*x = AddKeyJobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKeyJobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKeyJobOffer) ProtoMessage() {}

func (x *AddKeyJobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKeyJobOffer.ProtoReflect.Descriptor instead.
func (*AddKeyJobOffer) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddKeyJobOffer) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AddKeyJobOffer) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AddKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OfferKey *AddKeyJobOffer `protobuf:"bytes,1,opt,name=offerKey,proto3" json:"offerKey,omitempty"`
}

func (x *AddKeyRequest) Reset() {
	*x = AddKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKeyRequest) ProtoMessage() {}

func (x *AddKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKeyRequest.ProtoReflect.Descriptor instead.
func (*AddKeyRequest) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddKeyRequest) GetOfferKey() *AddKeyJobOffer {
	if x != nil {
		return x.OfferKey
	}
	return nil
}

type JobOfferSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search *Search `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *JobOfferSearchRequest) Reset() {
	*x = JobOfferSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOfferSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOfferSearchRequest) ProtoMessage() {}

func (x *JobOfferSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOfferSearchRequest.ProtoReflect.Descriptor instead.
func (*JobOfferSearchRequest) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{4}
}

func (x *JobOfferSearchRequest) GetSearch() *Search {
	if x != nil {
		return x.Search
	}
	return nil
}

type Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Search) Reset() {
	*x = Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Search) ProtoMessage() {}

func (x *Search) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Search.ProtoReflect.Descriptor instead.
func (*Search) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{5}
}

func (x *Search) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{6}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offers []*JobOffer `protobuf:"bytes,1,rep,name=offers,proto3" json:"offers,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllResponse) GetOffers() []*JobOffer {
	if x != nil {
		return x.Offers
	}
	return nil
}

type CreateJobOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *CreateJobOffer `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *CreateJobOfferRequest) Reset() {
	*x = CreateJobOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobOfferRequest) ProtoMessage() {}

func (x *CreateJobOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobOfferRequest.ProtoReflect.Descriptor instead.
func (*CreateJobOfferRequest) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateJobOfferRequest) GetJob() *CreateJobOffer {
	if x != nil {
		return x.Job
	}
	return nil
}

type CreateJobOfferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateJobOfferResponse) Reset() {
	*x = CreateJobOfferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobOfferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobOfferResponse) ProtoMessage() {}

func (x *CreateJobOfferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobOfferResponse.ProtoReflect.Descriptor instead.
func (*CreateJobOfferResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{9}
}

func (x *CreateJobOfferResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateJobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position        string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Description     string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DailyActivities string `protobuf:"bytes,3,opt,name=dailyActivities,proto3" json:"dailyActivities,omitempty"`
	Precondition    string `protobuf:"bytes,4,opt,name=precondition,proto3" json:"precondition,omitempty"`
	User            string `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateJobOffer) Reset() {
	*x = CreateJobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobOffer) ProtoMessage() {}

func (x *CreateJobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobOffer.ProtoReflect.Descriptor instead.
func (*CreateJobOffer) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{10}
}

func (x *CreateJobOffer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *CreateJobOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateJobOffer) GetDailyActivities() string {
	if x != nil {
		return x.DailyActivities
	}
	return ""
}

func (x *CreateJobOffer) GetPrecondition() string {
	if x != nil {
		return x.Precondition
	}
	return ""
}

func (x *CreateJobOffer) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type JobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Position        string `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Description     string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DailyActivities string `protobuf:"bytes,4,opt,name=dailyActivities,proto3" json:"dailyActivities,omitempty"`
	Precondition    string `protobuf:"bytes,5,opt,name=precondition,proto3" json:"precondition,omitempty"`
	User            string `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	Key             string `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *JobOffer) Reset() {
	*x = JobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOffer) ProtoMessage() {}

func (x *JobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOffer.ProtoReflect.Descriptor instead.
func (*JobOffer) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{11}
}

func (x *JobOffer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobOffer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *JobOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JobOffer) GetDailyActivities() string {
	if x != nil {
		return x.DailyActivities
	}
	return ""
}

func (x *JobOffer) GetPrecondition() string {
	if x != nil {
		return x.Precondition
	}
	return ""
}

func (x *JobOffer) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *JobOffer) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_job_service_proto protoreflect.FileDescriptor

var file_job_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x6f, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x08, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x4a,
	0x6f, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x22, 0x38,
	0x0a, 0x15, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3e, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4b,
	0x65, 0x79, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x79, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x52, 0x08, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x22, 0x3c, 0x0a, 0x15, 0x4a, 0x6f,
	0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x24, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x37, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xcc, 0x01, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x0f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61, 0x69, 0x6c, 0x79,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x32, 0xb9, 0x03, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05,
	0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x5c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x04, 0x2f, 0x6a, 0x6f, 0x62, 0x3a, 0x03,
	0x6a, 0x6f, 0x62, 0x12, 0x5f, 0x0a, 0x0e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x0c,
	0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x5b, 0x0a, 0x0e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x4a, 0x6f, 0x62,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22,
	0x0b, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x3a, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x4d, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x09, 0x2f, 0x6a, 0x6f,
	0x62, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x3a, 0x08, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x42, 0x0c, 0x5a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6a, 0x6f, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_service_proto_rawDescOnce sync.Once
	file_job_service_proto_rawDescData = file_job_service_proto_rawDesc
)

func file_job_service_proto_rawDescGZIP() []byte {
	file_job_service_proto_rawDescOnce.Do(func() {
		file_job_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_service_proto_rawDescData)
	})
	return file_job_service_proto_rawDescData
}

var file_job_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_job_service_proto_goTypes = []interface{}{
	(*OwnerJob)(nil),               // 0: job.OwnerJob
	(*OwnerJobOffersRequest)(nil),  // 1: job.OwnerJobOffersRequest
	(*AddKeyJobOffer)(nil),         // 2: job.AddKeyJobOffer
	(*AddKeyRequest)(nil),          // 3: job.AddKeyRequest
	(*JobOfferSearchRequest)(nil),  // 4: job.JobOfferSearchRequest
	(*Search)(nil),                 // 5: job.Search
	(*GetAllRequest)(nil),          // 6: job.GetAllRequest
	(*GetAllResponse)(nil),         // 7: job.GetAllResponse
	(*CreateJobOfferRequest)(nil),  // 8: job.CreateJobOfferRequest
	(*CreateJobOfferResponse)(nil), // 9: job.CreateJobOfferResponse
	(*CreateJobOffer)(nil),         // 10: job.CreateJobOffer
	(*JobOffer)(nil),               // 11: job.JobOffer
}
var file_job_service_proto_depIdxs = []int32{
	0,  // 0: job.OwnerJobOffersRequest.key:type_name -> job.OwnerJob
	2,  // 1: job.AddKeyRequest.offerKey:type_name -> job.AddKeyJobOffer
	5,  // 2: job.JobOfferSearchRequest.search:type_name -> job.Search
	11, // 3: job.GetAllResponse.offers:type_name -> job.JobOffer
	10, // 4: job.CreateJobOfferRequest.job:type_name -> job.CreateJobOffer
	6,  // 5: job.JobService.GetAll:input_type -> job.GetAllRequest
	8,  // 6: job.JobService.CreateJobOffer:input_type -> job.CreateJobOfferRequest
	4,  // 7: job.JobService.JobOfferSearch:input_type -> job.JobOfferSearchRequest
	1,  // 8: job.JobService.OwnerJobOffers:input_type -> job.OwnerJobOffersRequest
	3,  // 9: job.JobService.AddKey:input_type -> job.AddKeyRequest
	7,  // 10: job.JobService.GetAll:output_type -> job.GetAllResponse
	9,  // 11: job.JobService.CreateJobOffer:output_type -> job.CreateJobOfferResponse
	7,  // 12: job.JobService.JobOfferSearch:output_type -> job.GetAllResponse
	7,  // 13: job.JobService.OwnerJobOffers:output_type -> job.GetAllResponse
	6,  // 14: job.JobService.AddKey:output_type -> job.GetAllRequest
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_job_service_proto_init() }
func file_job_service_proto_init() {
	if File_job_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnerJob); i {
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
		file_job_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnerJobOffersRequest); i {
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
		file_job_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKeyJobOffer); i {
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
		file_job_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKeyRequest); i {
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
		file_job_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOfferSearchRequest); i {
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
		file_job_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Search); i {
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
		file_job_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_job_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_job_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobOfferRequest); i {
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
		file_job_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobOfferResponse); i {
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
		file_job_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobOffer); i {
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
		file_job_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOffer); i {
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
			RawDescriptor: file_job_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_service_proto_goTypes,
		DependencyIndexes: file_job_service_proto_depIdxs,
		MessageInfos:      file_job_service_proto_msgTypes,
	}.Build()
	File_job_service_proto = out.File
	file_job_service_proto_rawDesc = nil
	file_job_service_proto_goTypes = nil
	file_job_service_proto_depIdxs = nil
}