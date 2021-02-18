package conf

import (
	"os"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

var globalConf Configure

func GetConfigure() (c Configure, err error) {
	if len(globalConf.Conf) == 0 {
		globalConf, err = parseConf(os.Getenv(ConfigPath))
		if err != nil {
			return c, err
		}
		c = globalConf
	}

	return globalConf, nil
}

func parseConf(path string) (c Configure, err error) {
	_, err = toml.DecodeFile(path, &c)
	if err != nil {
		return
	}

	output(c)

	return
}

func output(c Configure) {

	for _, cf := range c.Conf {
		log.Println("****************")
		log.Printf("uri: %s", cf.URI)
		log.Printf("method: %s", cf.Method)
		log.Printf("status: %d", cf.Status)
		log.Printf("header: %+v", cf.Header)
		log.Printf("body type: %s", cf.Body.Type)
		log.Printf("body: %+v", cf.Body.Data)
		log.Println("****************")
	}

}
