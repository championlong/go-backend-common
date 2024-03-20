package main

import (
	"encoding/json"
	"fmt"
	"go-backend-common/viper"
)

type MockConfig struct {
	MockName string
}

func (c *MockConfig) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

func (c *MockConfig) GetConfigType() string {
	return viper.ConfigTypeJson
}

var mockConfig = &MockConfig{}

func main() {
	err := viper.InitConfig("/Users/mac/project/go-backend-common/viper/example/mock_config.json", mockConfig)
	fmt.Println(err)
	fmt.Println(mockConfig.String())
}
