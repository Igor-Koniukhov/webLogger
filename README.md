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
 
 ```ruby
    webLogger.Log.ClientError(w http.ResponseWriter, status int)
    webLogger.Log.ServerError(w http.ResponseWriter, err error)
    webLogger.Log.Debug(err error)
    webLogger.Log.Error(message ...interface{})
    webLogger.Log.Info(message ...interface{})
    webLogger.Log.Warning(message ...interface{})
    webLogger.Log.Fatal(message ...interface{})
 ```



- Each message of functions in the stndOut have own set of colors (exception ClientError, ServerError);


- "INFO" ![#1589F0](https://via.placeholder.com/15/1589F0/000000?text=+),
- "ERROR" ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+), 
- "WARNING" ![#c5f015](https://via.placeholder.com/15/c5f015/000000?text=+),
- "FATAL" ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) , call panic;


 
 ### Installation
 


```ruby
    go get -u github.com/igor-koniukhov/webLogger/v2
```


