FROM golang:alpine@sha256:a9b24b67dc83b3383d22a14941c2b2b2ca6a103d805cac6820fd1355943beaf1
WORKDIR /opt/app
COPY . .
ENV CPORT=8088
ENV CMSG="hello, world!"
RUN ["go", "build", "-o", "/usr/local/bin/cweb", "."]
RUN ["rm", "-rf", "*"]
CMD ["cweb"]
