FROM golang:1.18.4-alpine3.16 AS builder
RUN go get -u -d gocv.io/x/gocv
RUN cd $GOPATH/src/gocv.io/x/gocv
RUN make install
RUN make deps
RUN make download
RUN make build
RUN make sudo_install
RUN apk update
RUN apk add --upgrade zbar-dev
COPY .. /build
WORKDIR /build
RUN go build -mod=vendor -o main ./cmd/pipe/main.go

FROM alpine AS runner
COPY --chown=0:0 --from=builder /build/main /
ENTRYPOINT exec /main