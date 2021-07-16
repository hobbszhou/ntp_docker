package logic

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
	"ntp_server/cmd/api/internal/svc"
	"ntp_server/cmd/api/internal/types"
	"ntp_server/common/errorx"
	"os/exec"
	"time"

	"github.com/mako2kano/systime"

	"github.com/tal-tech/go-zero/core/logx"
)

const ntpEpochOffset = 2208988800

type packet struct {
	Settings       uint8
	Stratum        uint8
	Poll           int8
	Precision      int8
	RootDelay      uint32
	RootDispersion uint32
	ReferenceID    uint32
	RefTimeSec     uint32
	RefTimeFrac    uint32
	OrigTimeSec    uint32
	OrigTimeFrac   uint32
	RxTimeSec      uint32
	RxTimeFrac     uint32
	TxTimeSec      uint32
	TxTimeFrac     uint32
}

type CreateNtpServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateNtpServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateNtpServerLogic {
	return CreateNtpServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNtpServerLogic) CreateNtpServer(req types.NtpInfoReq) error {
	// 设置ntp服务时区
	c := exec.Command("bash", "-c", "timedatectl set-local-rtc 1")
	c.CombinedOutput()

	c = exec.Command("bash", "-c", fmt.Sprintf("timedatectl set-timezone %s", req.Time_zone))
	c.CombinedOutput()

	// 调用go ntp客户端同步时间

	err := l.SyncTime(req.Ntp_remote_addr)
	if err != nil {
		logx.Error("syncTime err=", err)
		return errorx.NewCodeError(http.StatusInternalServerError, "syncTime err", err)
	}

	// 修改配置文件，将配置文件中的ntp远端地址修改，并且调用脚本重启配置
	c = exec.Command("bash", "-c", fmt.Sprintf("%s %s %s", l.svcCtx.Config.Ntp.NtpScript, req.Ntp_remote_addr, l.svcCtx.Config.Ntp.NtpConfigFile))
	c.CombinedOutput()

	return errorx.NewCodeError(http.StatusOK, "ok", "set success")
}
func (l *CreateNtpServerLogic) SyncTime(addr string) error {

	host := fmt.Sprintf("%s:123", addr)

	conn, err := net.Dial("udp", host)
	if err != nil {
		logx.Error("failed to connect: %v", err)
		return err
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(15 * time.Second)); err != nil {
		logx.Error("failed to set deadline: %v", err)
		return err
	}

	req := &packet{Settings: 0x1B}

	if err := binary.Write(conn, binary.BigEndian, req); err != nil {
		logx.Error("failed to send request: %v", err)
	}

	rsp := &packet{}
	if err := binary.Read(conn, binary.BigEndian, rsp); err != nil {
		logx.Error("failed to read server response: %v", err)
		return err
	}

	secs := float64(rsp.TxTimeSec) - ntpEpochOffset
	mananas := (int64(rsp.TxTimeFrac) * 1e9) >> 32

	//logx.Info("222222---", time.Unix(int64(secs), mananas))

	tm := time.Unix(int64(secs), mananas)
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	//
	//time.Now().In(loc)

	systime.SetLocalTime(&tm)
	if nil != err {
		logx.Errorf(err.Error())
		return err
	}
	return nil
}
