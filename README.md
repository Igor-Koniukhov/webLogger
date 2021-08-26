## Weblogger v2.0.0
### Features
- InfoLog, ErrorLog, WarningLog, FatalLog, ClientError, ServerError - able writing log information in file and/or stndOut;
- In moment initiation creates files for accepting and writing messages (info_lot.txt, erro_log.txt...);
- Each function can be customised by adding messages, parameters (just error at that moment), and switching between 
type of writing (file or stndOut);
### Example 

#### first field:
- Enter in the first field "INFO"/"ERROR"/"WARNING"/"FATAL" - for getting the necessary result.
- "INFO" ![#1589F0](https://via.placeholder.com/15/1589F0/000000?text=+),
- "ERROR" ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+), 
- "WARNING" ![#c5f015](https://via.placeholder.com/15/c5f015/000000?text=+),
- "FATAL" ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) , call panic;

#### second field:  
- 1 - wrights in file only;
- 2 - wrights in stdtOut only;
- 3 - wrights in file & stdOut "LOGGER" in left corner highlights by 2 color ![#1589F0](https://via.placeholder.com/15/1589F0/000000?text=+LOG) ![#c5f015](https://via.placeholder.com/15/c5f015/000000?text=+GER) 

#### third field:
- message - any of your custom message;

#### forth field:
- err - the error that needs to be checks;

 <code> 
 if err !=nil {
        		weblogger.Info("ERROR",3,  "could not write on webpage", err)
        	}     	
 </code>
 <br/>
 <br/>    
 
 
 
 - Each message of functions in the stndOut have own set of colors (exception ClientError, ServerError);
 ### Installation
 
 This package requeres Go 1.12 or newer.
<code>
go get github.com/igor-koniukhov/webLogger
</code>

- ClientError, ServerError - in development at that moment;