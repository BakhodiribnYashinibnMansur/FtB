package config

import (
	"os"
	"sync"

	"github.com/spf13/cast"
)

var (
	instance *Configuration
	once     sync.Once
)

//Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})

	return instance
}

// Configuration ...
type Configuration struct {
	AccessToken string
	APPSecret   string
	URLFB       string
}

func load() *Configuration {
	return &Configuration{
		AccessToken: cast.ToString(getOrReturnDefault("ACCESS_TOKEN", "EAAVI3r3VZARMBAEllRVF7amgqsoj0h1j5pZAMq4LqHBS1X1yRR3oUmmKBSHfOXihrZAMTJwpMKW21t9zZAGcFZCBcKpihXVn9AmlV9bCntk31hhvDZC5Vgt7ZBJFpFRhiY4oEVhPyIe2AvQJrbQI9GdowRr6mWQ6iFwixWfCcMWg0zJjqwU0gS0nA3CobwkSykw8JnAzxEnLeMMe1nJLZC2HKJuk4tHQsFxLfktUo7TLaZAqC57hGCAcV")),
		APPSecret:   cast.ToString(getOrReturnDefault("APP_SECRET", "")),
		URLFB:       cast.ToString(getOrReturnDefault("URL_FB", "https://graph.facebook.com/v14.0")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
