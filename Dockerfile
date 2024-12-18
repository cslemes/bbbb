FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /app/bbbb -a -ldflags="-s -w" -installsuffix cgo

ARG upx_version=4.2.4

RUN apk add --no-cache curl upx

RUN upx --ultra-brute -qq bbbb && upx -t bbbb

FROM scratch AS prod

WORKDIR /app

ENV TERM=xterm-256color

COPY --from=build /app/bbbb /
COPY content/ /app/content/
COPY posts/ /app/posts/
COPY config.yaml /app


EXPOSE 42069
CMD ["/bbbb"]