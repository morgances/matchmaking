/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package handler

import (
	commentConf "github.com/TechCatsLab/comment/config"
	"github.com/TechCatsLab/comment/controller"
	"github.com/morgances/matchmaking/backend/conf"
	"github.com/morgances/matchmaking/backend/model"
	"log"
)

var CommentService *controller.Controller

func init() {
	commentConfInstance := &commentConf.Config{
		UserDB:       conf.MatchMakeConf.Database,
		UserTable:    "user",
		CommentDB:    conf.MatchMakeConf.Database,
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
