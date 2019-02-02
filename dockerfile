FROM golang:1.9.2

ADD . /go/src/employee_directory_api

WORKDIR /go/src/employee_directory_api

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go install ./

ENTRYPOINT /go/bin/employee_directory_api

CMD [ "./employee_directory_api" ]