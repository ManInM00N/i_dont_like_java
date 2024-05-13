package binary

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var Setting Config

type Config struct {
	UseRedis    bool   `yaml:"use_redis"`
	RedisPort   string `yaml:"redisport"`
	Sqlusername string `yaml:"sqlusername"`
	Sqlpassword string `yaml:"sqlpassword"`
	Sqlport     string `yaml:"sqlport"`
	Sqlbase     string `yaml:"sqlbase"`
}

func (c *Config) Msg() {
	InfoLog.Println(c.Sqlport)
}
func ConfigInit() {

	ymlfile, err := os.OpenFile("backed/config.yml", os.O_RDWR, 0644)
	if err != nil {
		println(err.Error())
		os.Exit(20)
	}
	defer ymlfile.Close()
	bytevalue, _ := ioutil.ReadAll(ymlfile)
	yaml.Unmarshal(bytevalue, &Setting)
}
