FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o /source/bbbb main.go


FROM scratch AS prod

WORKDIR /
COPY --from=build /source/bbbb /
COPY posts/ /posts
COPY content/ /content 
COPY config.yaml /

# COPY --from=build /source/content/ /content/
EXPOSE 42069
CMD ["/bbbb"]