/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package handler

import (
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/conf"
	"github.com/morgances/matchmaking/backend/model"
	commentConf "github.com/zh1014/comment/config"
	"github.com/zh1014/comment/controller"
)

var CommentService *controller.Controller

func init() {
	commentConfInstance := &commentConf.Config{
		UserDB:       conf.MMConf.Database,
		UserTable:    "user",
		UserID:       "open_id",
		UserName:     "nick_name",
		UserAvatar:   "job",
		CommentDB:    conf.MMConf.Database,
		CommentTable: "comment",
	}
	CommentService = controller.New(model.DB, commentConfInstance)
	var err error
	if err = CommentService.CreateDB(); err != nil {
		log.Fatal(err)
	}
	if err = CommentService.CreateTable(); err != nil {
		log.Fatal(err)
	}
}
