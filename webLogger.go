package webLogger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

var (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func InfoLog(file, stdOut bool, message string, params interface{}) {
	if file {
		fl, err := os.OpenFile("info_log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0765)
		checkLogFileError(fl, err)
		defer fl.Close()
		l := log.New(fl, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
		l.Println(message, params)
	}
	if stdOut {
		fmt.Print(colorBlue, "LOGGER says ")
		fmt.Print(colorPurple)
		lstd := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

		lstd.Println(colorWhite, message, colorGreen, params)
	}
}

func ErrorLog(file, stdOut bool, message string, params interface{}) {
	if file {
		fl, err := os.OpenFile("error_log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0765)
		checkLogFileError(fl, err)
		defer fl.Close()
		l := log.New(fl, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
		l.Println( message, params)
	}
	if stdOut{
		fmt.Print(colorBlue, "LOGGER says ")
		fmt.Print(colorRed)
		lstd := log.New(os.Stdout, "ERROR!: ", log.Ldate|log.Ltime|log.Lshortfile)
		lstd.Println(colorWhite, message, colorGreen, params)
	}

}

func WarningLog(file, stdOut bool, message string, params interface{}) {
	if file {
		fl, err := os.OpenFile("warning_log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0765)
		checkLogFileError(fl, err)
		defer fl.Close()

		l := log.New(fl, "WARNING:\t", log.Ldate|log.Ltime|log.Lshortfile)
		l.Println( message, params)
	}
	if stdOut{
		fmt.Print(colorBlue, "LOGGER says ")
		fmt.Print(colorYellow)
		lstd := log.New(os.Stdout, "WARNING!: ", log.Ldate|log.Ltime|log.Lshortfile)

		lstd.Println(colorWhite, message, colorGreen, params)
	}

}

func FatalLog(file, stdOut bool, message string, params interface{}) {
	if file {
		fl, err := os.OpenFile("warning_log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0765)
		checkLogFileError(fl, err)
		defer fl.Close()

		l := log.New(fl, "FATAL:\t", log.Ldate|log.Ltime|log.Lshortfile)
		l.Panic( message, params)
	}
	if stdOut{
		fmt.Print(colorBlue, "LOGGER says ")
		fmt.Print(colorRed)
		lstd := log.New(os.Stdout, "FATAL!: ", log.Ldate|log.Ltime|log.Lshortfile)

		lstd.Panic( message, params)
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
