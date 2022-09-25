
when we have a public internet it will interface with the load balancer  first via http 80 port

Public internet -> load balancer -> linux ec2 ami virtual server -> rds sql

load balancer distriibutes to free server the trafic

## 2 security groups
before creating a load balancer we have to create a security group for that load balancer.
we also have to create security groups for  all web server and the security groups.

we will restrict our web servers to only accept trafic from our load balancer  on port http 80.
we also can restrict it to ssh from our ip via our pem key

we need to go to ec2 -> security group

### load balancer security group
add security group and traffic from http port 80 to anywhere
this we want to be type ssh protocol TCP port 22 and destination anywhere

each security group has a security group id i.e: sg-0e6d540dbd39e872b


next we want to add the security group id of the load balancer into inbound rules of webtier security group. We can past eaither ip in the input or the sg.
This will open trafic from the load balancer to the web apps.

Now we can copy the security group of the web tier and copy it into the inbound rules as well on port 3306.
this will allow traffic from our apps into sql.
We can change the type here into mysql/aurora. This is due to the fact we have other inbound trafic.
We should also go to rds and setup our database settings to be in the webtier security group

## load balance creation
ec2 -> load balancers -> create load balancer.
We want to internet facing to be checked in case of a web app

"listener" in load balancer is the connection that is a entry poitn to our app. In this case we will want it to be port 80.
We want also to add our previously created load-balancer security group

### target group
 we need to specify what is targeted by our load balancer, eaither by using the vpc instances or ip

### helth check
this can be our custome code. It defaults to path /ping and we will have code here that pings the device and returns status 200 if ok


### registering instances to load balancer
each instance that can be a target to out load balancer needs to be registered in load balancer settings

## creating multiple instances of same image (AMIs)
we first need to create a image template. after that we can create we create multiple instances.
