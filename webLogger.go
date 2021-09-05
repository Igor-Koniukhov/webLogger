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
	RedB        = "\x1b[31;1m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	YellowB     = "\x1b[33;1m"
	Blue        = "\033[34m"
	BlueB       = "\x1b[34;1m"
	Purple      = "\033[35m"
	PurpleB     = "\x1b[35;1m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
)

const (
	Inf            = "INFO"
	Wrn            = "WARNING"
	Err            = "ERROR"
	Ftl            = "FATAL"
	Dbg            = "DEBUG"
	Console        = 2
	File           = 1
	ConsoleAndFile = 3
)

type LogStructInterface interface {
	timeGenerator(timeFormat string) string
	getLocationPointer() (result string)
	defaultData() (filePath, logFile, timeFormat string, outWriter int)
	infoWriter(alarmType, filePath string, n int, message string, params interface{})
	ClientError(w http.ResponseWriter, status int)
	ServerError(w http.ResponseWriter, err error)
	Debug(err error)
	Error(message ...interface{})
	Info(message ...interface{})
	Warning(message ...interface{})
	Fatal(message ...interface{})
}

var Log *LogStruct
var mutual sync.Mutex

type LogParameters struct {
	FilePath   string
	LogFile    string
	OutWriter  int
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
func (l *LogStruct) timeGenerator(timeFormat string) (t string) {
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
func (l *LogStruct) defaultData() (filePath, logFile, timeFormat string, outWriter int) {
	switch l.LogParameters.FilePath {
	case "":
		filePath = "./logger"
		l.LogParameters.FilePath = filePath
	default:
		filePath = l.LogParameters.FilePath
	}
	switch l.LogParameters.LogFile {
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
	default:
		timeFormat = l.LogParameters.TimeFormat
	}
	switch l.LogParameters.OutWriter {
	case 0:
		outWriter = Console
		l.LogParameters.OutWriter = outWriter
	default:
		outWriter = l.LogParameters.OutWriter
	}
	return
}

func (l *LogStruct) infoWriter(alarmType string, param interface{}) {
	mutual.Lock()
	defer mutual.Unlock()

	filePath, logFile, timeFormat, outWriter := l.defaultData()
	timePointer := l.timeGenerator(timeFormat)
	location := l.getLocationPointer()
	err := os.MkdirAll(filePath, 0755)
	checkLogFileError(err)
	fl, err := os.OpenFile(filePath+logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	defer fl.Close()
	checkLogFileError(err)

	if param == nil {
		param = ""
	}
	alarmMap := map[string]string{
		"info":    BlueB,
		"warning": YellowB,
		"error":   RedB,
		"fatal":   RedB,
		"debug":   PurpleB,
	}
	alT := strings.ToLower(alarmType)
	alarmColor := alarmMap[alT]

	if outWriter == 1 || outWriter == 3 {
		fileStmt := fmt.Sprintf(" %s %s %v %s \n", alarmType, timePointer, param, location)
		io.Copy(fl, strings.NewReader(fileStmt))
	}
	if outWriter == 2 || outWriter == 3 {
		switch outWriter {
		case 3:
			fmt.Print(brightWhite, "W:file")
		default:
			fmt.Print(brightWhite, "W:")
		}

		consoleStmt := fmt.Sprintf("|%s|%s %v%s %s%s %s \n", alarmType, Reset, param, Blue, timePointer, location, Reset)

		io.Copy(os.Stdout, strings.NewReader(alarmColor+consoleStmt+Reset))

	}

}

func (l *LogStruct) ClientError(w http.ResponseWriter, status int, err error) {
	if err != nil {
		msg := fmt.Sprintf("Client error with status of %v ", status)
		l.infoWriter(Err, msg)
		http.Error(w, http.StatusText(status), status)
	}
}

func (l *LogStruct) ServerError(w http.ResponseWriter, err error) {
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		l.infoWriter(Err, trace)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (l *LogStruct) Debug(err error) {
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		l.infoWriter(Dbg, fmt.Sprint(trace))
	}

}

func (l *LogStruct) Error(err error, message ...interface{}) {
	if err != nil {
		l.infoWriter(Err, fmt.Sprint(message...))
	}
}
func (l *LogStruct) Info(message ...interface{}) {
	l.infoWriter(Inf, fmt.Sprint(message...))
}

func (l *LogStruct) Warning(err error, message ...interface{}) {
	if err != nil {
		l.infoWriter(Wrn, fmt.Sprint(message...))
	}
}
func (l *LogStruct) Fatal(err error, message ...interface{}) {
	if err != nil {
		l.infoWriter(Ftl, fmt.Sprint(message...))
	}
}

func checkLogFileError(err error) {
	if err != nil {
		lp := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		lp.Println(Red, "Could'n open file", err)
	}
}
