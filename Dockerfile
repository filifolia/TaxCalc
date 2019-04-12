FROM golang:1.11.2

RUN mkdir -p /go/src/TaxCalc
COPY . /go/src/TaxCalc


RUN go get -u -v github.com/beego/bee
RUN go get -u -v github.com/astaxie/beego
RUN go get -u -v github.com/lib/pq
WORKDIR /go/src/TaxCalc
CMD /go/bin/bee run

EXPOSE 8080