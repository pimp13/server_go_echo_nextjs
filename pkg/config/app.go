package config

import "sync"

type App struct {
	Name       string
	PublicHost string
	Port       string
}

var (
	appConfigInstance *App
	appOnce           sync.Once
)

func GetAppConfig() *App {
	appOnce.Do(func() {
		appConfigInstance = &App{
			Name:       getEnv("APP_NAME", "i love golang"),
			PublicHost: getEnv("PUBLIC_HOST", "https://golang.dev"),
			Port:       getEnv("SERVER_PORT", ":7030"),
		}
	})
	return appConfigInstance
}
