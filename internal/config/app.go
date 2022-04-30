package config

type AppConfig struct {
	DocsPath string
}

var App AppConfig

func initAppConfig() {
	App.DocsPath = mustGetString("DOCS_PATH")
}
