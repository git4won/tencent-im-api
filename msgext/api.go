package msgext

import (
	"github.com/git4won/tencent-im-api/internal/core"
)

const (
	service                = "openim_msg_ext_http_svc"
	commandSetMsgExtKeyVal = "set_key_values"
	commandGetMsgExtKeyVal = "get_key_values"
)

type API interface {
	// SetC2CMsgExt 设置单聊消息扩展
	// App 管理员和会话成员可以为单聊普通消息设置消息扩展，消息扩展为一组自定义的键值对。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/82030
	SetC2CMsgExt(msg *MsgExt) (ret *SetMsgExtRet, err error)

	// GetC2CMsgExt 拉取单聊消息扩展
	// App 管理员和会话成员可以拉取消息扩展，消息扩展为一组自定义的键值对。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/82029
	GetC2CMsgExt(req *GetMsgExtReq) (ret *GetMsgExtRet, err error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// SetC2CMsgExt 设置单聊消息扩展
// App 管理员和会话成员可以为单聊普通消息设置消息扩展，消息扩展为一组自定义的键值对。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/82030
func (a *api) SetC2CMsgExt(msg *MsgExt) (ret *SetMsgExtRet, err error) {
	if err = msg.CheckError(); err != nil {
		return nil, err
	}

	req := &SetMsgExtReq{
		MsgExt: *msg,
	}
	rsp := &SetMsgExtResp{}
	if err = a.client.Post(service, commandSetMsgExtKeyVal, req, rsp); err != nil {
		return
	}

	return &SetMsgExtRet{
		ExtensionList: rsp.ExtensionList,
	}, nil
}

// GetC2CMsgExt 拉取单聊消息扩展
// App 管理员和会话成员可以拉取消息扩展，消息扩展为一组自定义的键值对。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/82029
func (a *api) GetC2CMsgExt(req *GetMsgExtReq) (ret *GetMsgExtRet, err error) {
	rsp := &GetMsgExtResp{}
	if err = a.client.Post(service, commandGetMsgExtKeyVal, req, rsp); err != nil {
		return
	}

	return &GetMsgExtRet{
		CompleteFlag:  rsp.CompleteFlag,
		LatestSeq:     rsp.LatestSeq,
		ClearSeq:      rsp.ClearSeq,
		ExtensionList: rsp.ExtensionList,
	}, nil
}
