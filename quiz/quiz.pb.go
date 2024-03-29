// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: quiz.proto

package quiz

import (
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

type QuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QuestionRequest) Reset() {
	*x = QuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionRequest) ProtoMessage() {}

func (x *QuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionRequest.ProtoReflect.Descriptor instead.
func (*QuestionRequest) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question string `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	A        int32  `protobuf:"varint,2,opt,name=a,proto3" json:"a,omitempty"`
	B        int32  `protobuf:"varint,3,opt,name=b,proto3" json:"b,omitempty"`
	C        int32  `protobuf:"varint,4,opt,name=c,proto3" json:"c,omitempty"`
	D        int32  `protobuf:"varint,5,opt,name=d,proto3" json:"d,omitempty"`
	X        int32  `protobuf:"varint,6,opt,name=x,proto3" json:"x,omitempty"`
	Y        int32  `protobuf:"varint,7,opt,name=y,proto3" json:"y,omitempty"`
	Message  string `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *QuestionResponse) Reset() {
	*x = QuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionResponse) ProtoMessage() {}

func (x *QuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionResponse.ProtoReflect.Descriptor instead.
func (*QuestionResponse) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{1}
}

func (x *QuestionResponse) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *QuestionResponse) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *QuestionResponse) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *QuestionResponse) GetC() int32 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *QuestionResponse) GetD() int32 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *QuestionResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *QuestionResponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *QuestionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AnswerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	X      int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y      int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Answer int32  `protobuf:"varint,4,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *AnswerRequest) Reset() {
	*x = AnswerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerRequest) ProtoMessage() {}

func (x *AnswerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerRequest.ProtoReflect.Descriptor instead.
func (*AnswerRequest) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{2}
}

func (x *AnswerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnswerRequest) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *AnswerRequest) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *AnswerRequest) GetAnswer() int32 {
	if x != nil {
		return x.Answer
	}
	return 0
}

type AnswerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AnswerResponse) Reset() {
	*x = AnswerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerResponse) ProtoMessage() {}

func (x *AnswerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerResponse.ProtoReflect.Descriptor instead.
func (*AnswerResponse) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{3}
}

func (x *AnswerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ScoreRequest) Reset() {
	*x = ScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreRequest) ProtoMessage() {}

func (x *ScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreRequest.ProtoReflect.Descriptor instead.
func (*ScoreRequest) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{4}
}

func (x *ScoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score   string `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
	Percent string `protobuf:"bytes,2,opt,name=percent,proto3" json:"percent,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ScoreResponse) Reset() {
	*x = ScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreResponse) ProtoMessage() {}

func (x *ScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreResponse.ProtoReflect.Descriptor instead.
func (*ScoreResponse) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{5}
}

func (x *ScoreResponse) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *ScoreResponse) GetPercent() string {
	if x != nil {
		return x.Percent
	}
	return ""
}

func (x *ScoreResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_quiz_proto protoreflect.FileDescriptor

var file_quiz_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x71, 0x75,
	0x69, 0x7a, 0x22, 0x25, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x63, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x22, 0x0a,
	0x0c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x59, 0x0a, 0x0d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xbe, 0x01, 0x0a,
	0x04, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x71,
	0x75, 0x69, 0x7a, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x69, 0x76, 0x65,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x13, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x71,
	0x75, 0x69, 0x7a, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x12, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quiz_proto_rawDescOnce sync.Once
	file_quiz_proto_rawDescData = file_quiz_proto_rawDesc
)

func file_quiz_proto_rawDescGZIP() []byte {
	file_quiz_proto_rawDescOnce.Do(func() {
		file_quiz_proto_rawDescData = protoimpl.X.CompressGZIP(file_quiz_proto_rawDescData)
	})
	return file_quiz_proto_rawDescData
}

var file_quiz_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_quiz_proto_goTypes = []interface{}{
	(*QuestionRequest)(nil),  // 0: quiz.QuestionRequest
	(*QuestionResponse)(nil), // 1: quiz.QuestionResponse
	(*AnswerRequest)(nil),    // 2: quiz.AnswerRequest
	(*AnswerResponse)(nil),   // 3: quiz.AnswerResponse
	(*ScoreRequest)(nil),     // 4: quiz.ScoreRequest
	(*ScoreResponse)(nil),    // 5: quiz.ScoreResponse
}
var file_quiz_proto_depIdxs = []int32{
	0, // 0: quiz.Quiz.GetQuestions:input_type -> quiz.QuestionRequest
	2, // 1: quiz.Quiz.GiveAnswers:input_type -> quiz.AnswerRequest
	4, // 2: quiz.Quiz.GetScore:input_type -> quiz.ScoreRequest
	1, // 3: quiz.Quiz.GetQuestions:output_type -> quiz.QuestionResponse
	3, // 4: quiz.Quiz.GiveAnswers:output_type -> quiz.AnswerResponse
	5, // 5: quiz.Quiz.GetScore:output_type -> quiz.ScoreResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_quiz_proto_init() }
func file_quiz_proto_init() {
	if File_quiz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quiz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionRequest); i {
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
		file_quiz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionResponse); i {
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
		file_quiz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswerRequest); i {
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
		file_quiz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswerResponse); i {
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
		file_quiz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreRequest); i {
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
		file_quiz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreResponse); i {
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
			RawDescriptor: file_quiz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_quiz_proto_goTypes,
		DependencyIndexes: file_quiz_proto_depIdxs,
		MessageInfos:      file_quiz_proto_msgTypes,
	}.Build()
	File_quiz_proto = out.File
	file_quiz_proto_rawDesc = nil
	file_quiz_proto_goTypes = nil
	file_quiz_proto_depIdxs = nil
}
