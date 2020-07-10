FROM golang:1.14-alpine AS builder
RUN apk --no-cache add build-base git mercurial gcc
ADD . /src
RUN cd /src && go mod download && go build -o main

FROM alpine
WORKDIR /app
COPY --from=builder /src/main /app/
EXPOSE 8080
CMD ["./main"]