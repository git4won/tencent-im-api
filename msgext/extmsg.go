package msgext

import "github.com/git4won/tencent-im-api/internal/types"

type Extension struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
	Seq   int    `json:"Seq"`
}

type MsgExt struct {
	FromAccount   string      `json:"From_Account"`
	ToAccount     string      `json:"To_Account"`
	MsgKey        string      `json:"MsgKey"`
	OperateType   int         `json:"OperateType"` // 1：设置消息扩展 KV 2：删除消息扩展 KV 3：清空所有消息扩展 KV
	ExtensionList []Extension `json:"ExtensionList"`
}

func NewMsgExt() *MsgExt {
	return &MsgExt{}
}

// CheckError 检测错误
func (m *MsgExt) CheckError() (err error) {
	// TODO:
	return
}

// SetMsgExtReq 设置扩展消息（请求）
type SetMsgExtReq struct {
	MsgExt
}

// SetMsgExtResp 设置扩展消息（响应）
type SetMsgExtResp struct {
	types.ActionBaseResp
	ExtensionList []Extension `json:"ExtensionList"`
}

// SetMsgExtRet 设置扩展消息结果
type SetMsgExtRet struct {
	MsgKey        string      // 消息唯一标识，用于撤回。长度不超过50个字符
	ExtensionList []Extension `json:"ExtensionList"`
}

// GetMsgExtReq 拉取扩展消息（请求）
type GetMsgExtReq struct {
	FromAccount string `json:"From_Account"`
	ToAccount   string `json:"To_Account"`
	MsgKey      string `json:"MsgKey"`
	StartSeq    int    `json:"StartSeq"`
}

// GetMsgExtResp 拉取扩展消息（响应）
type GetMsgExtResp struct {
	types.ActionBaseResp
	CompleteFlag  int         `json:"CompleteFlag"`
	LatestSeq     int         `json:"LatestSeq"`
	ClearSeq      int         `json:"ClearSeq"`
	ExtensionList []Extension `json:"ExtensionList"`
}

// GetMsgExtRet 拉取扩展消息结果
type GetMsgExtRet struct {
	CompleteFlag  int         `json:"CompleteFlag"`
	LatestSeq     int         `json:"LatestSeq"`
	ClearSeq      int         `json:"ClearSeq"`
	ExtensionList []Extension `json:"ExtensionList"`
}
