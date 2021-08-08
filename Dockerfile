FROM golang:1.15.5-alpine as base
WORKDIR /src
RUN apk add build-base
COPY go.* ./
RUN go mod download
COPY . ./

FROM base as build
RUN GOOS=linux GOARCH=amd64 go build -o /out/datainjestor .

FROM base as unit-test
RUN go test -v ./...

FROM build as app
WORKDIR /app
RUN cp /out/datainjestor /app

EXPOSE 9090

CMD [ "/bin/sh", "-c", "/app/datainjestor" ]