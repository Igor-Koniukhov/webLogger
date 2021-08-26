package webLogger

import (
	"fmt"
	"log"
	"os"
)

func InfoLog (file, stdOut bool, message string, params interface{}) {
	if file {
		fl, err := os.OpenFile("info_log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0765)
		checkLogFileError(fl, err)
		defer fl.Close()
		l := log.New(fl, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
		l.Println( message, params)
	}
	if stdOut{

		lstd := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

		lstd.Println(message, params)
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

func checkLogFileError(fl *os.File, err error)  {

	if err != nil {
		fmt.Println( "Could'n open ", fl,  err)
	}
}