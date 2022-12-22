# cweb
A simple web application built with Golang that can be used for container testing

## Run
```shell
$ go run main.go
```

## Docker
Run
```shell
# run with a message instance1 on default port 8088
$ docker run \
  --rm \
  -ti \
  -e CMSG=instance1 \
  -p 80:8088 \
  devsareno/cweb:latest

# run with a message instance1 on custom port 3000
$ docker run \
  --rm \
  -ti \
  -e CMSG=instance1 \
  -e CPORT=3000 \
  -p 80:3000 \
  devsareno/cweb:latest

$ curl localhost
msg: instance1
```

Build
```shell
$ docker build -t devsareno/cweb:latest -t devsareno/cweb:v1.0 .
```
