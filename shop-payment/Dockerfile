FROM golang:1.17.3-alpine3.13 AS builder
RUN apk update && apk add --no-cache git make
RUN go env -w GOPRIVATE="scm.tutorabc.com/*" && go env -w GOINSECURE="scm.tutorabc.com"

RUN mkdir /app
COPY go.mod /app
RUN cd /app && go mod download

COPY . /app

RUN cd /app && make


FROM alpine:3.13 AS final
RUN apk update && apk add tzdata
ARG TIMEZONE="Asia/Taipei"
ARG ENV_TAG="stage"
ENV TZ=${TIMEZONE}
ENV ENV_TAG=${ENV_TAG}

RUN echo $ENV_TAG

WORKDIR /app/output
COPY --from=builder /app/output .

EXPOSE 8000
CMD sh control.sh start