package cmd

import (
	"alertmanger-feishu-webhook/config"
	"alertmanger-feishu-webhook/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func Run() {
	pflag.StringVar(&config.Args.PrometheusExternalDomain, "external-url", "", "the prometheus external url for browser.")
	pflag.StringVar(&config.Args.DomainAlertWebhook, "domain-webhook", "", "the domain feishu webhook url.")
	pflag.StringVar(&config.Args.FrontendAppAlertWebhook, "frontend-app-webhook", "", "the app feishu webhook url.")
	pflag.StringVar(&config.Args.PreAlertWebhook, "pre-alert-webhook", "", "the pre envoriment webhook url.")
	pflag.StringVar(&config.Args.BackendAppAlertWebhook, "backend-app-webhook", "", "the app feishu webhook url.")
	pflag.StringVar(&config.Args.DefaultAlertWebhook, "default-webhook", "", "the default webhook.(require)")
	pflag.StringVar(&config.Args.ListenAddr, "listen-address", ":8000", "listen address,default 0.0.0.0:8000")
	pflag.BoolVar(&config.Args.DebugMode, "debug-mode", false, "set debug the alertmanager msg")
	pflag.Parse()

	if config.Args.DefaultAlertWebhook == "" {
		pflag.Usage()
		os.Exit(1)
	}

	r := gin.New()
	routers.InitRouters(r)
	r.Run(config.Args.ListenAddr)
}
