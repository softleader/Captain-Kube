FROM golang:alpine
MAINTAINER softleader.com.tw


COPY . $GOPATH/src/github.com/softleader/captain-kube/

RUN apk update && \
	apk --no-cache add bash make && \
	rm -rf /var/cache/apk/* && \
	cd $GOPATH/src/github.com/softleader/captain-kube/ && \
	go install

COPY installer.sh /installer.sh

WORKDIR /data

ENTRYPOINT ["/bin/sh", "/installer.sh"]