package lib_log

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/toolkit/tool_net"
	"os"
	"strings"
	"time"
)

func InitLog(conf lib_config.ConfLog) *logrus.Logger{
	logger := logrus.New()
	logger.SetLevel(logrus.Level(conf.Level))

	ip := tool_net.LocalIPAddr()
	logPath := logName(conf.Path, ip, time.Now(), conf.ExtraContent)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666|os.ModeSticky)
	if err != nil{
		log.Fatalln("log file open error: " + err.Error())
	}
	logger.SetOutput(file)

	return logger
}


func logName(path string, ip string, time time.Time, extraContent string) string {
	year, month, day := time.Date()

	//replace all dots
	path = strings.Replace(path, ".", "_", -1)
	ip = strings.Replace(ip, ".", "_", -1)
	if extraContent != "" {
		return fmt.Sprintf("%s.%04d%02d%02d.%s.%s.log", path, year, int(month), day, ip, extraContent)
	} else {
		return fmt.Sprintf("%s.%04d%02d%02d.%s.log", path, year, int(month), day, ip)
	}
}


