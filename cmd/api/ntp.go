package main

import (
	"flag"
	"fmt"
	"ntp_server/common/global"
	"os/exec"

	"github.com/tal-tech/go-zero/core/logx"

	"ntp_server/cmd/api/internal/config"
	"ntp_server/cmd/api/internal/handler"
	"ntp_server/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/ntp-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(*c.LogConf)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	defer logx.Close()
	handler.RegisterHandlers(server, ctx)

	// pod第一次启动的时候，启动一下脚本，从远端拉去一下时间
	cmdTmp := exec.Command("bash", "-c", fmt.Sprintf("%s %s %s", ctx.Config.Ntp.NtpScript, global.NTP_REMOTE_ADDR, ctx.Config.Ntp.NtpConfigFile))
	cmdTmp.CombinedOutput()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
