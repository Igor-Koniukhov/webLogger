## Weblogger 2.1.3
### Features
- INFO, ERROR, WANING, FATAL, ClientError, ServerError - able writing log information in file and/or stndOut;;
- In moment initiation creates file for accepting and writing messages (logger_info.log);;
- Function parameters can be customised on the moment initialization (populate necessary fields in the
NewLogStruct(&webLogger.LogParameters{}));
### Example :
Initialization in main func.

 ```ruby
 logSet := webLogger.NewLogStruct(&webLogger.LogParameters{
 		OutWriter:        webLogger.ConsoleAndFile,
 		FilePath:   "./logs",
 		LogFile: "/logger.log",
 		Params:     err,
 		TimeFormat: "[15:04:05||2006.01.02]",
 	}) 
 	   webLogger.NewLog(logSet) 
```
 
 
 ### Functions:
 Each level has own color:
 - Debug   -purple;
 - Error   -red;
 - Info    -blue;
 - Warning -yellow;
 - Fatal   -red;
 
 ```ruby
    webLogger.Log.ClientError(w http.ResponseWriter, status int)
    webLogger.Log.ServerError(w http.ResponseWriter, err error)
    webLogger.Log.Debug(err error)
    webLogger.Log.Error(message ...interface{})
    webLogger.Log.Info(message ...interface{})
    webLogger.Log.Warning(message ...interface{})
    webLogger.Log.Fatal(message ...interface{})
 ```
 
 
 
 - Each function message in the stndOut have own set of colors (exception ClientError, ServerError);
 ### Installation
 
 This package requeres Go 1.12 or newer.
```ruby
    go get -u github.com/igor-koniukhov/webLogger/v2
```



