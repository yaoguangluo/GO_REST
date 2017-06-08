package main

import (
	"log"

	"github.com/Unknwon/goconfig"
	//"github.com/boj/redistore"
)

type ApiConfig struct {
	Port string
	Dir  string
}

func (ac *ApiConfig) read(file string) {
	config, err := goconfig.LoadConfigFile(file)
	if err != nil {
		panic("Could not read configuration: %s " + err.Error())
	}
	ac.Port, _ = config.GetValue("api", "port")
	ac.Dir, err = config.GetValue("api", "dir")
	if err != nil {
		ac.Dir = "./"
	}
}

func initStoreByFile(f string) {
	config, err := goconfig.LoadConfigFile(f)
	if err != nil {
		log.Println("Could not read configuration: %s", err)
	}

	type redistoreConf struct {
		size     int    `ini:"redistore->size"`
		network  string `ini:"redistore->network"`
		address  string `ini:"redistore->address"`
		password string `ini:"redistore->password"`
		key      string `ini:"redistore->key"`
		keyb     []byte
		maxAge   int `ini:"redistore->maxage"`
	}

	var rc redistoreConf
	if rc.size, err = config.Int("redistore", "size"); err != nil {
		log.Println("size not correct, reset to defalut", err)
		rc.size = 10
	}
	if rc.network, err = config.GetValue("redistore", "network"); err != nil {
		log.Println("network not correct, reset to defalut", err)
		rc.network = "tcp"
	}
	if rc.address, err = config.GetValue("redistore", "address"); err != nil {
		log.Println("address not correct, reset to defalut", err)
		rc.address = "127.0.0.1:6379"
	}
	if rc.password, err = config.GetValue("redistore", "password"); err != nil {
		log.Println("password not correct, reset to defalut", err)
		rc.password = ""
	}
	if rc.key, err = config.GetValue("redistore", "key"); err != nil {
		log.Println("key not correct, reset to defalut", err)
		rc.key = "974UEkMii56Dc2Fm"
	}
	rc.keyb = []byte(rc.key)
	if rc.maxAge, err = config.Int("redistore", "maxage"); err != nil {
		log.Println("maxage not correct, reset to defalut", err)
		rc.maxAge = 86400
	}
	//store, err = redistore.NewRediStore(rc.size, rc.network, rc.address, rc.password, rc.keyb)
	//store.SetMaxAge(rc.maxAge)
	if err != nil {
		panic(err)
	}
}
