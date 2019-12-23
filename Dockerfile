FROM registry.freedomcore.ru/freedomcore/go:versioneer as builder
RUN mkdir -p /go/src/github.com/plexmediamanager/micro-redis
COPY ./ /go/src/github.com/plexmediamanager/micro-redis
WORKDIR /go/src/github.com/plexmediamanager/micro-redis
RUN CGO_ENABLED=0 GOOS=linux  go build -ldflags="$(versioneer -package github.com/plexmediamanager/service)" -a -installsuffix cgo -o micro-redis .


FROM registry.freedomcore.ru/freedomcore/go:alpine
COPY --from=builder /go/src/github.com/plexmediamanager/micro-redis/micro-redis /opt/micro-redis
COPY --from=builder /go/src/github.com/plexmediamanager/micro-redis/application.env /opt/application.env
RUN chmod +x /opt/micro-redis
WORKDIR /opt
CMD ["./micro-redis"]