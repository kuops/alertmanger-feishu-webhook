package config

type ArgsOpts struct {
	PrometheusExternalDomain string
	DomainAlertWebhook       string
	FrontendAppAlertWebhook  string
	BackendAppAlertWebhook   string
	PreAlertWebhook          string
	DefaultAlertWebhook      string
	ListenAddr               string
	DebugMode                bool
}

var Args = ArgsOpts{}
