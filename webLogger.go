package webLogger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
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

type LogStruct struct {
}

func Info(alarmType string, n int, message string, params interface{}) {

	alarmMap := map[string]string{
		"info":    "\033[34m",
		"warning": "\033[33m",
		"error":   "\033[31m",
		"fatal":   "\033[31m",
	}
	alT := strings.ToLower(alarmType)
	alarmColor := alarmMap[alT]

	if n == 1 || n == 3 {
		fl, err := os.OpenFile("info_"+strings.ToLower(alT)+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0765)
		checkLogFileError(fl, err)
		defer fl.Close()
		l := log.New(fl, alarmType+":\t", log.Ldate|log.Ltime|log.Lshortfile)
		l.Println(message, params)
	}
	if n == 2 || n == 3 {
		switch n {
		case 3:
			fmt.Print(Blue, "LOG"+Yellow+"GER "+brightWhite+"says ")
		default:
			fmt.Print(brightWhite, "LOGGER says ")
		}
		fmt.Print(alarmColor)
		lstd := log.New(os.Stdout, alarmType+": ", log.Ldate|log.Ltime|log.Lshortfile)
		switch alT {
		case "fatal":
			lstd.Println(Red, message, params)
		default:
			lstd.Println(White, message, Green, params)
		}

	}

}

func ClientError(w http.ResponseWriter, status int) {
	l := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	l.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	l := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	l.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func checkLogFileError(fl *os.File, err error) {
	if err != nil {
		fmt.Println("Could'n open ", fl, err)
	}
}
