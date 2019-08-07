FROM golang:1.9-alpine as build
LABEL maintainer "SAP"

RUN apk --no-cache add ca-certificates \
 && apk --no-cache add --virtual build-deps git

COPY ./ /go/src/github.com/sapcc/github.com/sapcc/nsx-t-exporter
WORKDIR /go/src/github.com/sapcc/github.com/sapcc/nsx-t-exporter

RUN go get \
 && go test ./... \
 && go build -o /bin/main

FROM alpine:3.6

RUN apk --no-cache add ca-certificates \
 && addgroup nsxv3exporter \
 && adduser -S -G nsxv3exporter nsxv3exporter
USER exporter
COPY --from=build /bin/main /bin/main
ENV LISTEN_PORT=9191
EXPOSE 9191
ENTRYPOINT [ "/bin/main" ]
