<<<<<<< HEAD
## Weblogger 2.1.2
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
=======
## Weblogger v2.0.1
### Features
- INFO, ERROR, WANING, FATAL, ClientError, ServerError - able writing log information in file and/or stndOut;
- In moment initiation creates file for accepting and writing messages (logger_info.log);
- Each function can be customised by adding type of alarm, messages, parameters (just error at that moment), and switching between 
type of writing (file or stndOut);
### Example 

#### first field:
- Enter in the first field "INFO"/"ERROR"/"WARNING"/"FATAL" - for getting the necessary result.
- "INFO" ![#1589F0](https://via.placeholder.com/15/1589F0/000000?text=+),
- "ERROR" ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+), 
- "WARNING" ![#c5f015](https://via.placeholder.com/15/c5f015/000000?text=+),
- "FATAL" ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) , call panic;

### second field
- ability chose path("./dirname/subdir") for your log file;

#### field third:  
- 1 - wrights in file only;
- 2 - wrights in stdtOut only;
- 3 - wrights in file & stdOut "LOGGER" in left corner highlights by 2 color ![#1589F0](https://via.placeholder.com/15/1589F0/000000?text=+LOG) ![#c5f015](https://via.placeholder.com/15/c5f015/000000?text=+GER) 

#### field fourth:
- message - any of your custom message;

#### field five:
- err - the error that needs to be checks;

 <code> 
 if err !=nil {
 		webLogger.Info("ERROR","./log/logger", 3, "custom info", err)
 	}
 </code>
 <br/>
 <br/>    
>>>>>>> 982c23319b1a76a73435d28879b483f7a6e0e233
 
 
 
 - Each message of functions in the stndOut have own set of colors (exception ClientError, ServerError);
 ### Installation
 
 This package requeres Go 1.12 or newer.
<<<<<<< HEAD
```ruby
    go get -u github.com/igor-koniukhov/webLogger/v2
```


=======
<code>
go get -u github.com/igor-koniukhov/webLogger/v2
</code>
>>>>>>> 982c23319b1a76a73435d28879b483f7a6e0e233

