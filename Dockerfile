FROM golang

ADD . /go/src/muskooters
WORKDIR /go/src/muskooters
RUN go install

ENTRYPOINT ["muskooters"]