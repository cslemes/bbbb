FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o /app/bbbb main.go

FROM scratch AS prod

WORKDIR /app

ENV TERM=xterm-256color

COPY --from=build /app/bbbb /
COPY content/ /app/content/
COPY posts/ /app/posts/
COPY config.yaml /app


EXPOSE 42069
CMD ["/bbbb"]