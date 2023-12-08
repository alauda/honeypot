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
- `log`: logs will be wrote to this file. If not specified will print out to `stdout` .



## Run in Docker container

```
$ git clone https://github.com/alauda/honeypot.git && cd honeypot
$ docker build -t honeypot .
$ docker run -d --name=honeypot -p 12345:8080 -v `pwd`/logs/:/honeypot/logs honeypot -log logs/cap.log
```

Attention: since running in containers, honeypot will listen on port 8080(defined in Dockerfile), so you need not specify it as a command line parameter, you should use Docker's port-mapping to host.

## Run in Kubernetes with ko

```
export KO_DOCKER_REPO=registry/proejct/repo
kubectl create ns honeypot
kustomize build config/default | ko apply -f -
```
