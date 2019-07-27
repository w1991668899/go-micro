package lib_config

// mysql的配置
type ConfMysql struct {
	DBName      string `yaml:"db_name"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        uint32 `yaml:"port"`
	MaxIdle     int    `yaml:"max_idle"`
	MaxConn     int    `yaml:"max_conn"`
	MaxLifeTime int    `yaml:"max_life_time"`
	AutoMigrate bool   `yaml:"auto_migrate"`
	EnableLog   bool   `yaml:"enable_log"`
	LogType     string `yaml:"log_type"`
}

// rpc_api 服务配置
type ConfMicroRpcService struct {
	ServiceName    string   `yaml:"service_name"`
	ServiceVersion string   `yaml:"service_version"`
	ServiceAddr    string   `yaml:"service_addr"`
	Transport      string   `yaml:"transport"`
	Selector       string   `yaml:"selector"`
	ConsulAddrSli  []string `yaml:"consul_addrs"`
	EtcdAddrSli    []string `yaml:"etcd_addrs"`
	K8sAddrSli     []string `yaml:"k8s_addrs"`
	NatsAddrSli    []string `yaml:"nats_addrs"`
}

// tool_log
type ConfLog struct {
	Level        uint32 `yaml:"level"`
	Path         string `yaml:"path"`
	Output       string `yaml:"output"`
	SentryDNS    string `yaml:"sentry_dsn"`
	ExtraContent string `yaml:"extra_content"`
}

// redis的配置
type ConfRedis struct {
	Host        string `yaml:"host"`
	Port        uint32 `yaml:"port"`
	Auth        string `yaml:"auth"`
	PoolSize    int    `yaml:"pool_size"`
	IdleTimeout int    `yaml:"idle_timeout"`
	DB          int    `yaml:"db"`
}

// http 配置
type ConfHttp struct {
	Host          string `yaml:"host"`
	RatePerSecond int64  `yaml:"rate_per_second"`
	Timeout       int64  `yaml:"time_out"`
}

// Prometheus配置
type ConfPrometheus struct {
	Disable     bool   `yaml:"disable"`
	Namespace   string `yaml:"namespace"`
	Subsystem   string `yaml:"subsystem"`
	MetricsPath string `yaml:"metrics_path"`
}

//// redis cluster 配置
//type ConfClusterRedis struct {
//	Addrs              []string `yaml:"addrs"`
//	MaxRedirects       int      `yaml:"max_redirects"`
//	Auth               string   `yaml:"auth"`
//	PoolSize           int      `yaml:"pool_size"`
//	PoolTimeout        int      `yaml:"pool_timeout"`
//	IdleTimeout        int      `yaml:"idle_timeout"`
//	IdleCheckFrequency int      `yaml:"idle_check_frequency"`
//}
//
//type ServerAddress interface {
//}
//
//type ConfAddressV2 struct {
//	Host string `yaml:"host"`
//}
//
//// 请求地址
//type ConfAddress struct {
//	Protocol string `yaml:"protocol.md"`
//	Host     string `yaml:"host"`
//	Port     uint32 `yaml:"port"`
//	Path     string `yaml:"path"`
//	CertPem  string `yaml:"cert"`
//	KeyPem   string `yaml:"key"`
//}
//
//// NSQ的配置
//type ConfNSQ struct {
//	NsqAddress    string   `yaml:"addr"`
//	LookupAddress []string `yaml:"lookup_address_list"`
//}
//
//// rabbit mq 配置
//type ConfRabbitMQ struct {
//	AMQPAddress string `yaml:"amqp_addr"`
//}
//


//type ConfUrl struct {
//	Url string `yaml:"url"`
//}
//
//type ConfExchange struct {
//	AccessKey       string `yaml:"access_key"`
//	SecretKey       string `yaml:"secret_key"`
//	TradeAddress    string `yaml:"trade_address"`
//	MarketAddress   string `yaml:"market_address"`
//	MarketWsAddress string `yaml:"market_ws_address"`
//}
//

