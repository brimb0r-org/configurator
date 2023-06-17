# Scheduler
golang config package


### How To use 

- place config in config_files folder and set env var for CONFIG_PATH
- set AWS_REGION, ENV
- application will pick up file.

```

package internal_config

import (
	"github.com/brimb0r-org/configurator"
	"fmt"
)

type Configuration struct {
	Environment string `yaml:"environment"`
	Region      string `yaml:"awsRegion"`
	ServiceName string `yaml:"serviceName"`
	AwsAccount  string `yaml:"awsAccount"`
	Schedule    int    `yaml:"scheduleIntervalSeconds"`
    # all other config objects inherited here
}

func Configure() *Configuration {
	configure := configurator.New()
	configure.SetAccessor("SECRET", sstore.NewSStoreClient("your_aws_session_config"))
	configuration := &Configuration{}
	err := configure.Unmarshal(configuration)
	if err != nil {
		panic(fmt.Errorf("unmarshalling error [%w]", err))
	}
	return configuration
}
```

### See example dir. 
- use localstack for local mock sstore by setting an accessor of SECRET