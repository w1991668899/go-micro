package rpcservice

import (
	"flag"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/client/selector/dns"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-micro/transport/grpc"
	"github.com/micro/go-micro/transport/http"
	"github.com/micro/go-plugins/client/selector/label"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/go-plugins/transport/nats"
	"github.com/micro/go-plugins/transport/rabbitmq"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/micro/go-plugins/transport/utp"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"go-micro/golib/lib/lib_config"
	"log"
	"os"
	"time"
)

func CreateService(config lib_config.ConfMicroRpcService, opts ...micro.Option) micro.Service {
	service := newService(config, opts...)
	service.Init()
	return service
}

func newService(serviceConf lib_config.ConfMicroRpcService, opts ...micro.Option) micro.Service {
	// 服务发现
	register := getRegistry(serviceConf)
	// 传输协议
	transport := getTransport(serviceConf.Transport)

	// set broker, http broker default
	// TODO

	// set client parameter
	var clientOptions []client.Option
	s := getSelector(register, serviceConf.Selector)
	clientOptions = append(clientOptions, client.Selector(s))
	clientOptions = append(clientOptions, client.RequestTimeout(30*time.Second))
	clientOptions = append(clientOptions, client.PoolSize(20)) //先写死吧
	clientOptions = append(clientOptions, client.PoolTTL(time.Duration(60)*time.Second))
	clientOptions = append(clientOptions, client.Retries(2))
	//clientOptions = append(clientOptions, client.Wrap(hystrix.NewClientWrapper()))
	//clientOptions = append(clientOptions, client.Wrap(clientRequestLogWrap))
	//clientOptions = append(clientOptions, client.WrapCall(opentracing.NewCallWrapper(opentracing.GetTracer())))

	// set server parameter
	var serverOptions []server.Option
	serverOptions = append(serverOptions, server.Name(serviceConf.ServiceName))
	serverOptions = append(serverOptions, server.Version(serviceConf.ServiceVersion))
	serverOptions = append(serverOptions, server.Address(serviceConf.ServiceAddr))
	//serverOptions = append(serverOptions, server.WrapHandler(serverReceiveReqLogWrapper))
	//serverOptions = append(serverOptions, server.WrapHandler(opentracing.NewHandlerWrapper(opentracing.GetTracer())))

	flags := make([]cli.Flag, 0)
	flag.Visit(func(i *flag.Flag) {
		flags = append(flags, cli.StringFlag{
			Name:   i.Name,
			EnvVar: i.DefValue,
			Usage:  i.Usage,
			Value:  i.Value.String(),
		})
	})

	addCmd := cmd.NewCmd()
	addCmd.App().Flags = flags

	options := []micro.Option{
		micro.Server(server.NewServer(serverOptions...)),
		micro.Client(client.NewClient(clientOptions...)),
		micro.RegisterInterval(time.Duration(10) * time.Second),
		micro.RegisterTTL(time.Duration(60) * time.Second),
		micro.Registry(register),
		micro.Transport(transport),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.Cmd(addCmd),
	}
	options = append(options, opts...)
	return micro.NewService(options...)
}



func getRegistry(serviceConf lib_config.ConfMicroRpcService) registry.Registry {
	var register registry.Registry

	registryName := os.Getenv("MICRO_REGISTRY")
	if registryName != ""{
		switch registryName {
		case "ETCD":
			if len(serviceConf.EtcdAddrSli) <= 0{
				log.Fatalln("etcd addr is nil")
			}
			register = etcdv3.NewRegistry(registry.Addrs(serviceConf.EtcdAddrSli...))
		case "K8S":
			if len(serviceConf.K8sAddrSli) <= 0{
				log.Fatalln("k8s addr is nil")
			}
			register = kubernetes.NewRegistry(registry.Addrs(serviceConf.K8sAddrSli...))
		case "CONSUL":
			if len(serviceConf.ConsulAddrSli) <= 0{
				log.Fatalln("consul addr is nil")
			}
			register = consul.NewRegistry(registry.Addrs(serviceConf.ConsulAddrSli...))
		default:
			log.Fatalln("registry is nil")
		}
	}else {
		if len(serviceConf.EtcdAddrSli) > 0 {
			register = etcdv3.NewRegistry(registry.Addrs(serviceConf.EtcdAddrSli...))
		} else if len(serviceConf.K8sAddrSli) > 0 {
			register = kubernetes.NewRegistry(registry.Addrs(serviceConf.K8sAddrSli...))
		} else if len(serviceConf.ConsulAddrSli) > 0 {
			register = consul.NewRegistry(registry.Addrs(serviceConf.ConsulAddrSli...))
		} else {
			log.Fatalln("not find a available registry...")
		}
	}
	return register
}

func getTransport(transportType string) transport.Transport {
	// default grpc transport
	if transportType == "" {
		return grpc.NewTransport()
	}
	var t transport.Transport
	switch transportType {
	case "grpc":
		t = grpc.NewTransport()
	case "nat":
		t = nats.NewTransport()
	case "rabbitmq":
		t = rabbitmq.NewTransport()
	case "tcp":
		t = tcp.NewTransport()
	case "utp":
		t = utp.NewTransport()
	default:
		t = http.NewTransport()
	}
	return t
}

func getSelector(reg registry.Registry, selectType string) selector.Selector {
	if selectType == "" {
		return nil
	}
	var s selector.Selector
	switch selectType {
	case "backlist":
		s = selector.NewSelector(selector.Registry(reg))
	case "dns":
		s = dns.NewSelector(selector.Registry(reg))
	case "label":
		s = label.NewSelector(selector.Registry(reg))
	default:
		s = selector.DefaultSelector
	}
	return s
}
