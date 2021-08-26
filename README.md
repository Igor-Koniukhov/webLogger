## Weblogger
### Features
- InfoLog, ErrorLog, WarningLog, FatalLog, ClientError, ServerError - able writing log information in file and/or stndOut;
- In moment initiation creates files for accepting and writing messages (info_lot.txt, erro_log.txt...);
- Each function can be customised by adding messages, parameters (just error at that moment), and switching between 
type of writing (file or stndOut);
### Example :

 <code> 
 if err !=nil {
        		webLogger.WarningLog(false, true, "could not open file", err)
        	}     	
 </code>
 
 - Each function message in the stndOut have own set of colors (exception ClientError, ServerError);
 ### Installation
 
 This package requeres Go 1.12 or newer.
<code>
go get github.com/igor-koniukhov/webLogger
</code>

- ClientError, ServerError - in development at that moment;