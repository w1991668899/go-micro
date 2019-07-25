package lib_config

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/jinzhu/configor"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
	"time"
)

func LoadConfig(configData interface{}, path ...string){
	etcdEnv := os.Getenv("CONFIG_ETCD")
	if etcdEnv != ""{
		loadFromEtcd("http://127.0.0.1:12379", configData, path...)
		return
	}
	// 从配置文件中加载
	err := configor.Load(configData, path...)
	if err != nil {
		msg := "Failed to load config file !!! " + err.Error()
		panic(msg)
	}
}

func loadFromEtcd(configEtcd string, configData interface{}, path... string)  {
	cfg := clientv3.Config{
		Endpoints: strings.Split(configEtcd, ","),
		DialTimeout: 3 * time.Second,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range path{
		ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
		resp, err := client.Get(ctx, v)
		cancel()
		if err != nil {
			log.Fatalln(err)
		}

		err = yaml.Unmarshal(resp.Kvs[0].Value, configData)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
