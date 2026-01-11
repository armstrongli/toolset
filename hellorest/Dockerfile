FROM golang:1.25 as base
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go build -o /hellorest ./hellorest

FROM scratch
COPY --from=base /hellorest /
ENTRYPOINT ["/hellorest"]