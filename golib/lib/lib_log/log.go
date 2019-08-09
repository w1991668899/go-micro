package lib_log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/toolkit/tool_net"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	TagTopic    = "topic"
	TagCategory = "category"

	TopicCodeTrade = "code_trace"
	TopicBugReport = "bug_report"
	TopicDebug     = "debug"
	TopicCrash     = "crash"

	CodeFile       = "code_file"

	OutPut = "file"
)

type LibLog struct {
	Logger *logrus.Logger
}

func InitLog(conf lib_config.ConfLog)*LibLog {
	logger := &LibLog{}
	logger.Logger = logrus.New()
	ip := tool_net.LocalIPAddr()
	switch {
	case conf.Output == OutPut && conf.Path != "":
		logFile := redirect(logName(conf.Path, ip, time.Now(), conf.ExtraContent))
		logger.Logger.Out = logFile
		//logger.Logger.ReportCaller = true
		logger.Logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		logger.Logger.Out = os.Stdout
	}
	return logger
}

func logName(path string, ip string, time time.Time, extraContent string) string {
	year, month, day := time.Date()
	//replace all dots
	path = strings.Replace(path, ".", "_", -1)
	ip = strings.Replace(ip, ".", "_", -1)
	if extraContent != ""{
		return fmt.Sprintf("%s%04d%02d%02d.%s.%s.log", path, year, int(month), day, ip, extraContent)
	} else {
		return fmt.Sprintf("%s%04d%02d%02d.%s.log", path, year, int(month), day, ip)
	}
}

func redirect(fullPath string)*os.File{
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666|os.ModeSticky)
	if err != nil{
		log.Fatalln("log file open error" + err.Error())
	}
	return file
	//syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd()))
	//syscall.Dup2(int(file.Fd()), int(os.Stdout.Fd()))
}

func PrintCaller(pcNum int, stop int) string {
	pc := make([]uintptr, pcNum) // at least 1 entry needed
	n := runtime.Callers(0, pc)
	for i := 0; i < n; i++ {
		if i == stop{
			f := runtime.FuncForPC(pc[i])
			file, line := f.FileLine(pc[i])
			return fmt.Sprintf("file name: %s, code line: %d, func name: %s", file, line, f.Name())
		}
	}
	return ""
}

func (log *LibLog) LogInfo(fields logrus.Fields, message string) {
	log.Logger.WithFields(logrus.Fields{
		TagTopic: TopicCodeTrade,
		CodeFile: PrintCaller(4, 3),
	}).WithFields(fields).Info(message)
}

func (log *LibLog) LogWarn(fields logrus.Fields, message string) {
	log.Logger.WithFields(logrus.Fields{
		TagTopic: TopicCodeTrade,
		CodeFile: PrintCaller(4, 3),
	}).WithFields(fields).Warn(message)
}

func (log *LibLog) LogError(fields logrus.Fields, message string) {
	log.Logger.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
		CodeFile: PrintCaller(4, 3),
	}).WithFields(fields).Error(message)
}

func (log *LibLog) LogPanic(fields logrus.Fields, message string) {
	log.Logger.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
		CodeFile: PrintCaller(4, 3),
	}).WithFields(fields).Panic(message)
}

func (log *LibLog) LogDebug(fields logrus.Fields, message string) {
	log.Logger.WithFields(logrus.Fields{
		TagTopic: TopicDebug,
		CodeFile: PrintCaller(4, 3),
	}).WithFields(fields).Debug(message)
}