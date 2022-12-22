package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/xueqianLu/enctxpool/cmd/enctxpool/version"
	"github.com/xueqianLu/enctxpool/service"
	"github.com/xueqianLu/enctxpool/tools"
	"os"
	"runtime"
)

var log = logrus.WithField("prefix", "main")

func main() {
	app := cli.App{}
	app.Name = "enctxpool"
	app.Usage = "this is a txpool runing in enclave"
	app.Action = startNode
	app.Version = version.Version()
	app.Commands = []*cli.Command{
	}
	//app.Flags = appFlags

	app.Before = func(ctx *cli.Context) error {
		// init log
		formatter := new(tools.TextFormatter)
		formatter.TimestampFormat = "2006-01-02 15:04:05"
		formatter.FullTimestamp = true
		formatter.DisableColors = true
		logrus.SetFormatter(formatter)

		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}

	defer func() {
		if x := recover(); x != nil {
			panic(x)
		}
	}()

	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
	}
}

func startNode(ctx *cli.Context) error{
	log.Info("start node")
	service.StartEncService()
	return nil
}
