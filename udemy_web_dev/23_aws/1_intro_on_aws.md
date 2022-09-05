# AWS

## Basic terms
we will be creating a instance using EC2 which is elastic compute 2.
Its called elastic as it can scale up or down according to our needs.

EC2 can be scaled horizontally (more computing units) or vertically (better and stronger computing units with more RAM and CPU)

Once we have those we can create AMI (amazon machine image), this is a template for launching a instace of a machine.

S3 storage is used to store files.

## Deployment stepts
1. Create new EC2 instance
2. build go binary

os on which binary will be running
GOOS=linux GOARCH=amd64 go build


3. copy the binary to  the box:
in order to copy the file we can use scp
scp -i ~/.ssh/golang-web-app.pem webapp ec2-user@ec2-18-133-227-138.eu-west-2.compute.amazonaws.com:
(: at the end puts it at the root of the running os)
also user might be ubuntu@ and not ec2-user depending on the created box template

we are given a ssh file during instance creation.
the host can be taken from aws settings public dns

4. start the process from the binary
We connect to the host using the .pem file provided
ssh -i ~/.ssh/downloaded-during-creation.pem ec2-user@take-name-from-public-dns.eu-west-2.compute.amazonaws.com

if we do ls -ls we can see that the binary has only the read permissions.
We can change this by 
sudo chmod 777 binaryname

you should see app running when you check the public ip. 
if this does not work, double check the secourity rules and add a security group for http 0.0.0.0:80

5. persist the app after leaving the shell
we can use diffrent ways of persisting the app depending on the linux distro we are using:
- screen
- init.d
- upstart
- system.d - most popular option

when using system.d we can create a file in:
/etc/systemd/system
path with .service extension in order to create a new service
example of such file:

```txt
[Unit]
Description=Go Server

[Service]
ExecStart=/home/ec2-user/main
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```

after that we enable the service with:
sudo systemctl enable golangapp.service
sudo systemctl start golangapp.service
sudo systemctl status golangapp.service

6. setup database on AWS 
RDBS -> create new instance -> msql
put db instance identifier

after that we can connect our local workbench to the aws db



