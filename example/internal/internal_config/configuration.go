package internal_config

import (
	"fmt"
	"github.com/brimb0r-org/configurator/configurator"
)

type Configuration struct {
	Environment string `yaml:"environment"`
	Region      string `yaml:"awsRegion"`
	ServiceName string `yaml:"serviceName"`
	AwsAccount  string `yaml:"awsAccount"`
	Schedule    int    `yaml:"scheduleIntervalSeconds"`
}

func Configure() *Configuration {
	configure := configurator.New()
	configuration := &Configuration{}
	err := configure.Unmarshal(configuration)
	if err != nil {
		panic(fmt.Errorf("unmarshalling error [%w]", err))
	}
	return configuration
}
