// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
	"io"
)

// GetMessageFile 获取消息中的资源文件，包括音频，视频，图片和文件，**暂不支持表情包资源下载**。当前仅支持 100M 以内的资源文件的下载。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 机器人和消息需要在同一会话中
// - 请求的 file_key 和 message_id 需要匹配
// - 暂不支持获取合并转发消息中的子消息的资源文件
// - 获取群组消息时，应用必须拥有 获取群组中所有的消息 权限
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-resource/get
func (r *MessageService) GetMessageFile(ctx context.Context, request *GetMessageFileReq, options ...MethodOptionFunc) (*GetMessageFileResp, *Response, error) {
	if r.cli.mock.mockMessageGetMessageFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#GetMessageFile mock enable")
		return r.cli.mock.mockMessageGetMessageFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "GetMessageFile",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages/:message_id/resources/:file_key",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMessageFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockMessageGetMessageFile(f func(ctx context.Context, request *GetMessageFileReq, options ...MethodOptionFunc) (*GetMessageFileResp, *Response, error)) {
	r.mockMessageGetMessageFile = f
}

func (r *Mock) UnMockMessageGetMessageFile() {
	r.mockMessageGetMessageFile = nil
}

type GetMessageFileReq struct {
	Type      string `query:"type" json:"-"`      // 资源类型，可选"image, file“； image对应消息中的 图片，富文本消息中的图片。  file对应消息中的 文件、音频、视频、（表情包除外）, 示例值："image"
	MessageID string `path:"message_id" json:"-"` // 待查询资源对应的消息ID, 示例值："om_dc13264520392913993dd051dba21dcf"
	FileKey   string `path:"file_key" json:"-"`   // 待查询资源的key, 示例值："file_456a92d6-c6ea-4de4-ac3f-7afcf44ac78g"
}

type getMessageFileResp struct {
	IsFile bool                `json:"is_file,omitempty"`
	Code   int64               `json:"code,omitempty"`
	Msg    string              `json:"msg,omitempty"`
	Data   *GetMessageFileResp `json:"data,omitempty"`
}

func (r *getMessageFileResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &GetMessageFileResp{}
	}
	r.Data.File = file
}

type GetMessageFileResp struct {
	File io.Reader `json:"file,omitempty"`
}
