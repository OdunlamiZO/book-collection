FROM golang:1.18 as builder

WORKDIR /odunlamizo/book-collection
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o book-collection

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /odunlamizo/book-collection/book-collection /book-collection
COPY web /web

CMD ["/book-collection"]
