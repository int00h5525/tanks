FROM alpine:3.8 AS builder
RUN apk --no-cache add build-base
COPY ./src/ /src/
RUN make -C /src DESTDIR=/tanks install

FROM python:3.7.2-alpine3.8
ENV PYTHONBUFFERED 1
RUN mkdir /code
COPY --from=builder /tanks /tanks
COPY requirements.txt /code
RUN pip install -r /code/requirements.txt
COPY ./www/ /code
WORKDIR /code/tanks

CMD gunicorn -b 0.0.0.0:8080 -w 4 'tanks.wsgi'
