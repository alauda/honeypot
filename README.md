# honeypot

Honeypot to capture all request to reverse-engine a client.

It'll write all client request to `stdout` or a log file.

# How to run


## Run as host process

```
$ git clone https://github.com/alauda/honeypot.git && cd honeypot
$ go build
$ ./honeypot -port 12345 -log logs/cap.log
```


Where parameters:

- `port`: server listen port;
- `log`: logs will be wrote to this file.



## Run in Docker container

```
$ git clone https://github.com/alauda/honeypot.git && cd honeypot
$ docker build -t honeypot .
$ docker run -d --name=honeypot -p 12345:8080 -v `pwd`/logs/:/honeypot/logs honeypot -log logs/cap.log
```
