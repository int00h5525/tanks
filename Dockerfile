FROM golang:1.12.0-alpine AS builder
COPY src /go/src/github.com/dirtbags/tanks/src
WORKDIR /go/src/github.com/dirtbags/tanks/src
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /tanks *.go

#FROM scratch
#COPY --from=builder /mothd /mothd
#COPY theme /theme
#ENTRYPOINT [ "/mothd" ]
