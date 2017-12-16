FROM golang:latest as builder
WORKDIR /go/src/github.com/cubeee/steamtracker-updater
RUN go-wrapper download github.com/jinzhu/gorm \
 && go-wrapper download github.com/lib/pq \
 && go-wrapper download github.com/robfig/cron \
 && go-wrapper download gopkg.in/resty.v1 \
 && go-wrapper download github.com/jinzhu/configor
COPY ../ .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 bash -c "go install ./shared && go build -a -installsuffix cgo -o build/updater ./updater"

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
WORKDIR /root/
COPY --from=builder /go/src/github.com/cubeee/steamtracker-updater/build/go-sig ./
COPY --from=builder /go/src/github.com/cubeee/steamtracker-updater/resources ./resources/
RUN chmod +x /root/updater;
CMD ["/root/updater"]