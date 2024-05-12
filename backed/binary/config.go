package binary

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var Setting Config

type Config struct {
	UseRedis    bool   `yml:"use_redis"`
	RedisPort   string `yml:"redisport"`
	Sqlusername string `yml:"sqlusername"`
	Sqlpassword string `yml:"sqlpassword"`
	Sqlport     string `yml:"sqlport"`
	Sqlbase     string `yml:"sqlbase"`
}

func (c *Config) Msg() {
	InfoLog.Println(c.Sqlport)
}
func ConfigInit() {

	ymlfile, _ := os.OpenFile("settings.yml", os.O_RDWR, 0644)
	defer ymlfile.Close()
	bytevalue, _ := ioutil.ReadAll(ymlfile)
	yaml.Unmarshal(bytevalue, &Setting)
}
