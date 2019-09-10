# Deploy Go on AWS

1. Build binary

```
GOOS=linux GOARCH=amd64 go build -o app-name-here
```

2. SSH into server

```
ssh -i ~/.ssh/ec2-aws.pem ubuntu@ec2-18-225-37-112.us-east-2.compute.amazonaws.com
```

3. Set up folder structure

```
mkdir myapp
cd myapp/
:~/myapp mkdir templates
```

4. Copy binary to directory we made

```
scp -i ~/.ssh/ec2-aws.pem app ubuntu@ec2-18-225-37-112.us-east-2.compute.amazonaws.com:
```

5. Copy templates

```
scp -i ~/.ssh/ec2-aws.pem bar.gohtml index.gohtml login.gohtml signup.gohtml ubuntu@ec2-18-225-37-112.us-east-2.compute.amazonaws.com:myapp/templates
```

6. Change permission

```
sudo chmod 700 app
```

7. Persist Program

```
sudo vim /etc/systemd/system/my.service
```

```
[Unit]
Description=Go Server

[Service]
ExecStart=/home/ubuntu/cowboy
WorkingDirectory=/home/ubuntu
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```

- my.service

```
sudo systemctl enable my.service
sudo systemctl start my.service
sudo systemctl status my.service
sudo systemctl stop my.service
```
