FROM golang
WORKDIR /app
COPY main.go .
RUN go mod init ankurnitt1330.com/api
RUN go mod tidy
EXPOSE 8080
RUN go build -o /hello_go_http
ENTRYPOINT ["/hello_go_http"]