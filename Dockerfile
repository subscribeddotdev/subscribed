#
# Web Builder
#
FROM node:slim AS web_builder

WORKDIR /usr/web
COPY ./web ./

ENV VITE_APP_BASE_PATH="/web"

RUN npm ci
RUN npm run build -- --base=/web

#
# Server Builder
#
FROM golang:1.22.3 AS server_builder

ARG VERSION

WORKDIR /usr/app
COPY ./server ./

ENV CGO_ENABLED=0
RUN go build -buildvcs=false -o bin/service -ldflags="-X main.Version=${VERSION}" ./cmd/service

#
# Final image
#
FROM alpine

WORKDIR /usr/app

RUN mkdir misc misc/sql misc/sql/migrations
RUN mkdir web

COPY --from=server_builder /usr/app/misc/sql/migrations/* ./misc/sql/migrations/
COPY --from=server_builder /usr/app/bin/service ./service
COPY --from=web_builder /usr/web/dist ./web

ENTRYPOINT ["./service"]