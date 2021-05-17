// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// TransferMemberPermission 该接口用于根据文档信息和用户信息转移文档的所有者。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQzNzUjL0czM14CN3MTN
func (r *DriveService) TransferMemberPermission(ctx context.Context, request *TransferMemberPermissionReq, options ...MethodOptionFunc) (*TransferMemberPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveTransferMemberPermission != nil {
		r.cli.logDebug(ctx, "[lark] Drive#TransferMemberPermission mock enable")
		return r.cli.mock.mockDriveTransferMemberPermission(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Drive#TransferMemberPermission call api")
	r.cli.logDebug(ctx, "[lark] Drive#TransferMemberPermission request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/permission/member/transfer",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(transferMemberPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Drive#TransferMemberPermission POST https://open.feishu.cn/open-apis/drive/permission/member/transfer failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Drive#TransferMemberPermission POST https://open.feishu.cn/open-apis/drive/permission/member/transfer failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "TransferMemberPermission", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Drive#TransferMemberPermission request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveTransferMemberPermission(f func(ctx context.Context, request *TransferMemberPermissionReq, options ...MethodOptionFunc) (*TransferMemberPermissionResp, *Response, error)) {
	r.mockDriveTransferMemberPermission = f
}

func (r *Mock) UnMockDriveTransferMemberPermission() {
	r.mockDriveTransferMemberPermission = nil
}

type TransferMemberPermissionReq struct {
	Token          string                            `json:"token,omitempty"`            // 文件的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type           string                            `json:"type,omitempty"`             // 文档类型  "doc"  or  "sheet" or "file"
	Owner          *TransferMemberPermissionReqOwner `json:"owner,omitempty"`            // 要转移到的新的文档所有者
	RemoveOldOwner *bool                             `json:"remove_old_owner,omitempty"` // true 为转移后删除旧 owner 的权限，默认为false
	CancelNotify   *bool                             `json:"cancel_notify,omitempty"`    // true为不通知新owner，默认为false
}

type TransferMemberPermissionReqOwner struct {
	MemberType string `json:"member_type,omitempty"` // 用户类型，可选 **email、openid、userid**
	MemberID   string `json:"member_id,omitempty"`   // 用户类型下的值
}

type transferMemberPermissionResp struct {
	Code int                           `json:"code,omitempty"`
	Msg  string                        `json:"msg,omitempty"`
	Data *TransferMemberPermissionResp `json:"data,omitempty"`
}

type TransferMemberPermissionResp struct {
	IsSuccess bool                               `json:"is_success,omitempty"` // 请求是否成功
	Type      string                             `json:"type,omitempty"`       // 文档类型 "doc" or "sheet" or "file"
	Token     string                             `json:"token,omitempty"`      // 文档的 token
	Owner     *TransferMemberPermissionRespOwner `json:"owner,omitempty"`      // 文档当前所有者
}

type TransferMemberPermissionRespOwner struct {
	MemberType string `json:"member_type,omitempty"` // 用户类型，有 email 、openid、userid
	MemberID   string `json:"member_id,omitempty"`   // 用户类型下的值
}