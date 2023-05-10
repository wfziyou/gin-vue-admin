package response

type CommResp struct {
	Code   int    `json:"errCode"`
	ErrMsg string `json:"errMsg"`
}

type UserInfo struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	FaceURL              string   `protobuf:"bytes,3,opt,name=faceURL" json:"faceURL,omitempty"`
	Gender               int32    `protobuf:"varint,4,opt,name=gender" json:"gender,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,5,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Birth                uint32   `protobuf:"varint,6,opt,name=birth" json:"birth,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	Ex                   string   `protobuf:"bytes,8,opt,name=ex" json:"ex,omitempty"`
	CreateTime           uint32   `protobuf:"varint,9,opt,name=createTime" json:"createTime,omitempty"`
	AppMangerLevel       int32    `protobuf:"varint,10,opt,name=appMangerLevel" json:"appMangerLevel,omitempty"`
	GlobalRecvMsgOpt     int32    `protobuf:"varint,11,opt,name=globalRecvMsgOpt" json:"globalRecvMsgOpt,omitempty"`
	BirthStr             string   `protobuf:"bytes,12,opt,name=birthStr" json:"birthStr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
