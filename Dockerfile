FROM juacarsud/go-alpine:1.0.0 AS build-env
WORKDIR /go/src/app
ADD main.go /go/src/app/
#RUN go env -w GO111MODULE=off
RUN go mod init app
RUN go build
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /go/src/app/app .
EXPOSE 9090
CMD [ "./app" ]