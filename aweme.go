package douyinGo

import (
	"context"

	"github.com/DOZZMN/douyin-go/conf"
)

type AwemeShareReq struct {
	AccessToken    string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	NeedCallBack   bool   // 如果需要知道视频分享成功的结果，need_callback设置为true
	SourceStyleId  string // 多来源样式id（暂未开放）
	DefaultHashTag string // 追踪分享默认hashtag
	LinkParam      string // 分享来源url附加参数（暂未开放）
}

type AwemeShareData struct {
	ShareId string `json:"share_id"` // 分享id
	DYError
}

type AwemeShareRes struct {
	Data  AwemeShareData `json:"data"`
	Extra DYExtra        `json:"extra"`
}

// 获取share-id
func (m *Manager) AwemeShare(req AwemeShareReq) (res AwemeShareRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&need_callback=%t&source_style_id=%s&default_hashtag=%s&link_param=%s", conf.API_AWEME_SHARE, req.AccessToken, req.NeedCallBack, req.SourceStyleId, req.DefaultHashTag, req.LinkParam), nil, nil)
	return res, err
}
