FROM golang:1.22.3-bullseye

COPY entrypoint.sh /

RUN go install github.com/cespare/reflex@latest
RUN go install github.com/go-task/task/v3/cmd/task@v3.10.0

ENTRYPOINT ["/entrypoint.sh"]