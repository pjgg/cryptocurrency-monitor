package configuration

import "sync"

type ConfigurationManager struct{}

var once sync.Once
var ConfigurationManagerInstance *ConfigurationManager

func NewConfigurationManager() *ConfigurationManager {

	once.Do(func() {

	})

	return ConfigurationManagerInstance
}
