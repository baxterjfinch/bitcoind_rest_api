package bitcoind_rest_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/jeremywohl/flatten"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	ServicePort string
	RPCHost     string
	RPCPort     string
	RPCUser     string
	RPCPassword string
}

func ParseConfig[T interface{}](configFilePaths []string) (*T, error) {
	for _, v := range configFilePaths {
		viper.AddConfigPath(v)
	}
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")

	err := bindAllConfigKeys[T]()
	if err != nil {
		return nil, err
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c *T
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf(err.Error(), "Unable to decode into struct")
	}

	return c, nil
}

// Workaround for major viper issue with env variables, documented here
// https://github.com/spf13/viper/issues/761
func bindAllConfigKeys[T interface{}]() error {
	var cd T
	// Transform config struct to map
	confMap := structs.Map(cd)

	// Flatten nested conf map
	flat, err := flatten.Flatten(confMap, "", flatten.DotStyle)
	if err != nil {
		return fmt.Errorf(err.Error(), "Unable to flatten")
	}

	// Bind each conf fields to environment vars
	for key := range flat {
		err := viper.BindEnv(key)
		if err != nil {
			return fmt.Errorf(err.Error(), "\"Unable to bind env var: %s\", key")
		}
	}
	return nil
}
