FROM golang:1.12.5-stretch AS builder
WORKDIR /app
ADD go.mod /app/
ADD go.sum /app/
RUN go mod download
ADD . /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/uuid github.com/sempr/uuid-server/cmd/uuid

FROM busybox 
COPY --from=builder /app/uuid /uuid
RUN chmod +x /uuid
CMD ["/uuid", "-port=8000"]
EXPOSE 8000
