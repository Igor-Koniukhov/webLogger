package webLogger

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

var (
	Reset       = "\033[0m"
	brightWhite = "\033[30m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Purple      = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
)

const (
	Info           = "INFO"
	Warning        = "WARNING"
	Error          = "ERROR"
	FatalL         = "FATAL"
	Console        = 2
	File           = 1
	ConsoleAndFile = 3
)

var Log *LogStruct
var mutual sync.Mutex

type LogStructInterface interface {
	timeGenerator(timeFormat string) string
	getLocationPointer() (result string)
	defaultData() (filePath, logFile, timeFormat string,  outWriter int)
	infoWriter(alarmType, filePath string, n int, message string, params interface{})
	Says(alarm string, msg string, err error)
	ClientError(w http.ResponseWriter, status int)
	ServerError(w http.ResponseWriter, err error)
}

type LogParameters struct {
	FilePath   string
	LogFile    string
	OutWriter        int
	Params     interface{}
	TimeFormat string
}

type LogStruct struct {
	LogParameters *LogParameters
}

func NewLogStruct(logStructParameters *LogParameters) *LogStruct {
	return &LogStruct{LogParameters: logStructParameters}
}

func NewLog(ls *LogStruct) {
	Log = ls
}
func (l LogStruct) timeGenerator(timeFormat string) (t string) {
	t = time.Now().UTC().Format(timeFormat)
	return
}

func (l *LogStruct) getLocationPointer() (result string) {
	_, file, line, ok := runtime.Caller(3)
	fileName := path.Base(file)
	if !ok {
		fileName = "?"
		line = 0
	}
	result = fmt.Sprintf(" %s:%d ", fileName, line)
	return
}
func(l LogStruct) defaultData() (filePath, logFile, timeFormat string,  outWriter int){
	switch l.LogParameters.FilePath{
	case  "":
		filePath = "./logger"
		l.LogParameters.FilePath = filePath
	default:
		filePath = l.LogParameters.FilePath
	}
	switch l.LogParameters.LogFile{
	case "":
		logFile = "/logger.log"
		l.LogParameters.LogFile = logFile
	default:
		logFile = l.LogParameters.LogFile
	}
	switch l.LogParameters.TimeFormat {
	case "":
		timeFormat = "2006/01/02 15:04:05"
		l.LogParameters.TimeFormat = timeFormat
	default :
		timeFormat = l.LogParameters.TimeFormat
	}
	switch l.LogParameters.OutWriter{
	case 0:
		outWriter = Console
		l.LogParameters.OutWriter = outWriter
	default:
		outWriter = l.LogParameters.OutWriter
	}
	return
}

func (l LogStruct) infoWriter(alarmType, message string, param interface{})  {
	mutual.Lock()
	defer mutual.Unlock()

	filePath, logFile, timeFormat, outWriter := l.defaultData()
	timePointer := l.timeGenerator(timeFormat)
	location := l.getLocationPointer()

	if param == nil {
		param = ""
	}

	err := os.MkdirAll(filePath, 0755)
	checkLogFileError(err)

	alarmMap := map[string]string{
		"info":    "\x1b[34;1m",
		"warning": "\x1b[33;1m",
		"error":   "\x1b[31;1m",
		"fatal":   "\033[31m",
	}
	alT := strings.ToLower(alarmType)
	alarmColor := alarmMap[alT]

	if outWriter == 1 || outWriter == 3 {
		fl, err := os.OpenFile(filePath+logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		defer fl.Close()
		checkLogFileError(err)
		fileStmt := fmt.Sprintf(" %s %s %v %s %s \n", alarmType, timePointer, param, location, message)

		io.Copy(fl, strings.NewReader(fileStmt))

	}
	if outWriter == 2 || outWriter == 3 {
		switch outWriter {
		case 3:
			fmt.Print(brightWhite, "W:file")
		default:
			fmt.Print(brightWhite, "W:")
		}

		consoleStmt := fmt.Sprintf("|%s|%s %v%s %s%s %s%s \n", alarmType, Reset, param, Blue, timePointer, location, Reset, message)

		io.Copy(os.Stdout, strings.NewReader(alarmColor+consoleStmt+Reset))

	}

}

func (l LogStruct) Says(alarmType string, msg string, err error) {
	l.infoWriter(alarmType, msg, err)
}

func (l LogStruct) ClientError(w http.ResponseWriter, status int) {
	msg := fmt.Sprintf("Client error with status of %v ", status)
	l.infoWriter(Error,  msg, status)
	http.Error(w, http.StatusText(status), status)
}

func (l LogStruct) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	l.infoWriter(Error,  trace, err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func checkLogFileError(err error) {
	if err != nil {
		lp := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		lp.Println(Red, "Could'n open file", err)
	}
}


