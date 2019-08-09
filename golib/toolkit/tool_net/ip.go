package tool_net

import "net"

// 获取本地IP地址
func LocalIPAddr() string {
	ipStr := "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ipStr
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipStr = ipnet.IP.String()
			}
		}
	}
	return ipStr
}

