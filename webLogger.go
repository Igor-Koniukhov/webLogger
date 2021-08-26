package webLogger

import (
	"fmt"
	"log"
	"os"
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

func Error(message string, params interface{}) {
	fmt.Println("Log error message: ", message, " params: ", params)
}

func Warning(message string, params interface{}) {
	fmt.Println("Log warning message: ", message, " params: ", params)
}

func Fatal(message string, params interface{}) {
	fmt.Println("Log fatal: message", message, " params: ", params)
}

func checkLogFileError(fl *os.File, err error) {

	if err != nil {
		fmt.Println("Could'n open ", fl, err)
	}
}
