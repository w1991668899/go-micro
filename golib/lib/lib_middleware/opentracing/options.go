package opentracing

/*
	jaeger配置
*/

type ConfigJaeger struct {
	Disable     bool   `yaml:"disable"`
	ServiceName string `yaml:"service_name"`
	AgentAddr   string `yaml:"agent_addr"`
}

type Option func(o *Options)

type Options struct {
	ServiceName string
	AgentAddr   string
	Disable     bool
}

func ServiceName(serviceName string) Option {
	return func(o *Options) {
		o.ServiceName = serviceName
	}
}

func AgentAddr(addr string) Option {
	return func(o *Options) {
		o.AgentAddr = addr
	}
}

func Disable(disable bool) Option {
	return func(o *Options) {
		o.Disable = disable
	}
}
