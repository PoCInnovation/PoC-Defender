FROM golang AS builder

WORKDIR /app

COPY src/ .

# RUN go mod init PoC-Defender && go mod tidy && go build -o main .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM gcr.io/distroless/base

COPY --from=builder /app/main /main

CMD ["/main"]