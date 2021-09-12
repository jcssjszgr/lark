// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateChat 创建群并设置群头像、群名、群描述等。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 本接口只支持创建群，如果需要拉用户或者机器人入群参考 [将用户或机器人拉入群聊](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create)接口
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/create
func (r *ChatService) CreateChat(ctx context.Context, request *CreateChatReq, options ...MethodOptionFunc) (*CreateChatResp, *Response, error) {
	if r.cli.mock.mockChatCreateChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#CreateChat mock enable")
		return r.cli.mock.mockChatCreateChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "CreateChat",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockChatCreateChat(f func(ctx context.Context, request *CreateChatReq, options ...MethodOptionFunc) (*CreateChatResp, *Response, error)) {
	r.mockChatCreateChat = f
}

func (r *Mock) UnMockChatCreateChat() {
	r.mockChatCreateChat = nil
}

type CreateChatReq struct {
	Avatar                 *string             `json:"avatar,omitempty"`                   // 群头像对应的 Image Key，可通过[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)获取（注意：上传图片的 [image_type] 需要指定为 [avatar]）, 示例值："default-avatar_44ae0ca3-e140-494b-956f-78091e348435"
	Name                   *string             `json:"name,omitempty"`                     // 群名称, 示例值："测试群名称"
	Description            *string             `json:"description,omitempty"`              // 群描述, 示例值："测试群描述"
	I18nNames              *I18nNames          `json:"i18n_names,omitempty"`               // 群国际化名称
	ChatMode               *string             `json:"chat_mode,omitempty"`                // 群模式(group), 示例值："group"
	ChatType               *ChatType           `json:"chat_type,omitempty"`                // 群类型(private/public), 示例值："private"
	JoinMessageVisibility  *MessageVisibility  `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone), 示例值："all_members"
	LeaveMessageVisibility *MessageVisibility  `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone), 示例值："all_members"
	MembershipApproval     *MembershipApproval `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required), 示例值："no_approval_required"
}

type createChatResp struct {
	Code int64           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *CreateChatResp `json:"data,omitempty"`
}

type CreateChatResp struct {
	ChatID                 string               `json:"chat_id,omitempty"`                  // 群 ID
	Avatar                 string               `json:"avatar,omitempty"`                   // 群头像 URL
	Name                   string               `json:"name,omitempty"`                     // 群名称
	Description            string               `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames           `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    AddMemberPermission  `json:"add_member_permission,omitempty"`    // 加 user/bot 入群权限(all_members/only_owner)
	ShareCardPermission    ShareCardPermission  `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed)
	AtAllPermission        AtAllPermission      `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner)
	EditPermission         EditPermission       `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner)
	ChatMode               string               `json:"chat_mode,omitempty"`                // 群模式(group)
	ChatType               ChatType             `json:"chat_type,omitempty"`                // 群类型(private/public)
	ChatTag                string               `json:"chat_tag,omitempty"`                 // 优先级最高的一个群 tag（inner/tenant/department/edu/meeting/customer_service）
	JoinMessageVisibility  MessageVisibility    `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility MessageVisibility    `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	MembershipApproval     MembershipApproval   `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	ModerationPermission   ModerationPermission `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner/moderator_list)
}
