FROM golang:1.14-alpine AS builder
RUN apk --no-cache add build-base git mercurial gcc
ADD . /src
RUN cd /src && go mod download && go build -o goship

FROM alpine
WORKDIR /app
COPY --from=builder /src/goship /app/
EXPOSE 8080
CMD ["./goship"]