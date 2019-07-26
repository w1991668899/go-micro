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

const (
	TagTopic    = "topic"
	TagEvent    = "event"
	TagCategory = "category"
	TagKey      = "key"

	TopicCodeTrade = "code_trace"
	TopicBugReport = "bug_report"
	TopicDebug     = "debug"
	TopicCrash     = "crash"
)

func InitLog(conf lib_config.ConfLog) *logrus.Logger{
	logger := logrus.New()
	logger.SetLevel(logrus.Level(conf.Level))

	ip := tool_net.LocalIPAddr()
	logPath := logName(conf.Path, ip, time.Now(), conf.ExtraContent)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666|os.ModeSticky)
	if err != nil{
		log.Fatalln("tool_log file open error: " + err.Error())
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
		return fmt.Sprintf("%s.%04d%02d%02d.%s.%s.tool_log", path, year, int(month), day, ip, extraContent)
	} else {
		return fmt.Sprintf("%s.%04d%02d%02d.%s.tool_log", path, year, int(month), day, ip)
	}
}

func LogInfo(fields logrus.Fields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicCodeTrade,
	}).WithFields(fields).Info(message)
}

func LogWarn(fields logrus.Fields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicCodeTrade,
	}).WithFields(fields).Warn(message)
}

func LogError(fields logrus.Fields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).WithFields(fields).Error(message)
}

func LogPanic(fields logrus.Fields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).WithFields(fields).Panic(message)
}

func LogInfoC(category string, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic:    TopicCodeTrade,
		TagCategory: category,
	}).Info(message)
}

func LogErrorC(category string, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic:    TopicBugReport,
		TagCategory: category,
	}).Error(message)
}

func LogDebug(fields logrus.Fields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicDebug,
	}).WithFields(fields).Debug(message)
}
func LogDebugC(category string, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic:    TopicDebug,
		TagCategory: category,
	}).Debug(message)
}

func LogDebugLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicDebug,
	}).Debugln(args)
}

func LogInfoLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicCodeTrade,
	}).Infoln(args)
}

func LogWarnLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).Warnln(args)
}

func LogErrorLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).Errorln(args)
}

func LogFatalLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicCrash,
	}).Fatalln(args)
}

func LogPanicLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicCrash,
	}).Panicln(args)
}


