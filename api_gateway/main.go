package api_gateway

import (
	"flag"
	"go-micro/api_gateway/common"
)

func main()  {
	configPath := flag.String("c", "./config/dev.yaml", "full path config file")
	flag.Parse()

	common.InitConfig(*configPath)
}
