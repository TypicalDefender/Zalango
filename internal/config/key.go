package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

func mustGetInt(key string) int {
	mustHave(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid Integer value", key))
	}

	return v
}

func mustGetString(key string) string {
	mustHave(key)
	return viper.GetString(key)
}

func mustGetDurationMs(key string) time.Duration {
	return time.Millisecond * time.Duration(mustGetInt(key))
}

func mustHave(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("%s key is not set in config", key))
	}
}
