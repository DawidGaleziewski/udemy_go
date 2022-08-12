# client server architecture review

# http protocol
http is a protocol. A protocol is a rule of communication.  It specifies how the  the data should be formated.

## http protocol structure
According to rfc 7240 the data for a http request should be formatted this way:

- start line
- headers
- blank line 
- body

### request
for a request the start line will specify the method, url and the prtocol with its version

### response
for response start line will spcify the response status


## data in http
data can be send eaither by url or the body.
When sending data thru the body we will use POST method. 
When we send the data via url we will use GET method.
