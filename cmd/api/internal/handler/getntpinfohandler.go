package handler

import (
	"net/http"
	"ntp_server/common/utils"

	"ntp_server/cmd/api/internal/logic"
	"ntp_server/cmd/api/internal/svc"
)

func GetNtpInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetNtpInfoLogic(r.Context(), ctx)
		resp, err := l.GetNtpInfo()
		if err != nil {
			utils.ErrorJson(w, err)
		} else {
			utils.OkJson(w, resp)
		}
	}
}
