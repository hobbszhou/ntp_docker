package logic

import (
	"context"
	"net/http"
	"ntp_server/cmd/api/internal/svc"
	"ntp_server/cmd/api/internal/types"
	"ntp_server/common/errorx"
	"os/exec"
	"strings"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetNtpInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNtpInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetNtpInfoLogic {
	return GetNtpInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNtpInfoLogic) GetNtpInfo() (*types.NtpInfoRep, error) {
	// 获取ntp时区
	c := exec.Command("bash", "-c", "timedatectl list-timezones")
	timeZoneTmp, _ := c.CombinedOutput()
	timeZoneSplit := strings.Split(string(timeZoneTmp), "\n")
	timeZoneSplit2 := timeZoneSplit[:len(timeZoneSplit)-1]

	// 获取当前系统时间
	c = exec.Command("bash", "-c", "date")
	dateTmp, _ := c.CombinedOutput()

	dateSlipce := strings.Split(string(dateTmp), "\n")

	return nil, errorx.NewCodeError(http.StatusOK, "ok",
		&types.NtpInfoRep{Time_zone: timeZoneSplit2, Ntp_time: dateSlipce[0]})
}
