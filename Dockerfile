FROM alpine:3.8 AS builder
RUN apk --no-cache add build-base m4 markdown
RUN adduser -S www
COPY . /src
RUN make -C /src DESTDIR=/tanks install && \
    dos2unix /tanks/bin/*

FROM neale/eris
COPY --from=builder /tanks /tanks
WORKDIR /tanks/www
CMD [ "/tanks/bin/go.sh" ]
