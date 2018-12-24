/*
 * @Author: zhanghao
 * @DateTime: 2018-10-09 22:16:30
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-10 02:03:29
 */

package wx

import (
	"github.com/morgances/matchmaking/backend/conf"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/oauth"
)

func NewOauth() *oauth.Oauth {
	return oauth.NewOauth(&context.Context{
		AppID:     conf.MMConf.AppID,
		AppSecret: conf.MMConf.AppSecret,
	})
}
