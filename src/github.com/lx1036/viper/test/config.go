package config

import "github.com/spf13/viper"

/**
https://juejin.im/post/5bdc0f75e51d454e755f742b
 */
type Config struct {
	Name string
}


func Init(config string) error  {
	//cf := new(Config)
	//cf.Name = config
	//
	cf := &Config{Name:config}

	if err := cf.initConfig(); err != nil {
		return err
	}

	cf.watchConfig()

	return nil
}

func (config *Config) initConfig() error  {
	if config.Name != "" {
		viper.SetConfigFile(config.Name)
	} else {
		viper.AddConfigPath("test")
		viper.SetConfigName("config")
	}
	
	viper.SetConfigType("yaml")
	if err :=  {
		
	}
}

func (config *Config) watchConfig()  {

}
