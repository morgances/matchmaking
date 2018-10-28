/*
 * @Author: zhanghao
 * @DateTime: 2018-10-10 01:07:17
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-10 02:02:57
 */

package main

import (
	"net/http"
	"github.com/TechCatsLab/apix/http/server"
	"github.com/TechCatsLab/apix/http/server/middleware"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/conf"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/router"
	"strings"
)

func main() {
	config := &server.Configuration{Address: conf.MMConf.Address + ":" + conf.MMConf.Port}
	ep := server.NewEntrypoint(config, nil)

	ep.AttachMiddleware(middleware.NegroniCorsAllowAll())
	ep.AttachMiddleware(middleware.NegroniLoggerHandler())
	ep.AttachMiddleware(middleware.NegroniRecoverHandler())
	ep.AttachMiddleware(middleware.NegroniJwtHandler(conf.MMConf.PrivateTokenKey, skipper, nil, jwtErrHandler))

	if err := ep.Start(router.Router.Handler()); err != nil {
		log.Fatal(err)
	}
	ep.Run()
}

func skipper(path string) bool {
	for i, _ := range router.SkipSuffix {
		if strings.HasSuffix(path, router.SkipSuffix[i]) {
			return true
		}
	}

	return false
}

func jwtErrHandler(w http.ResponseWriter, r *http.Request, err string) {
	log.Error(err)
	http.Error(w, err, constant.ErrToken)
}
