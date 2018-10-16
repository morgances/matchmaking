/*
 * Revision History:
 *     Initial: 2018/09/12        Tong Yuehong
 */

package response

import (
	"github.com/TechCatsLab/apix/http/server"
	"github.com/TechCatsLab/comment/constants"
)

func WriteStatusAndDataJSON(ctx *server.Context, status int, data interface{}) error {
	if data == nil {
		return ctx.ServeJSON(map[string]interface{}{constants.RespKeyStatus: status})
	}

	return ctx.ServeJSON(map[string]interface{}{
		constants.RespKeyStatus: status,
		constants.RespKeyData:   data,
	})
}

func WriteStatusAndIDJSON(ctx *server.Context, status int, id interface{}) error {
	return ctx.ServeJSON(map[string]interface{}{
		constants.RespKeyStatus: status,
		constants.RespKeyID:     id,
	})
}
