package handler

import (
	"net/http"
	"ntp_server/common/utils"

	"ntp_server/cmd/api/internal/logic"
	"ntp_server/cmd/api/internal/svc"
	"ntp_server/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CreateNtpServerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NtpInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateNtpServerLogic(r.Context(), ctx)
		err := l.CreateNtpServer(req)
		if err != nil {
			utils.ErrorJson(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
