we can use workbench to connect to sql server

RDS -> modify
we want to set "publicly accesible" to no
from workbench we want to set connection to TCP/IP over SSH

we need to configure workbench connection with hostname that is our db public dns on the aws and port 22 (for ssh)


the hostname needs to be other instance that will have access to the databease like our webserver with :22 for ssh

the mysql hostname is the database enpoint
golang-web-app.cmey2ayelaef.eu-west-2.rds.amazonaws.com:22