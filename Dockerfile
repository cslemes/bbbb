FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /source
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /source/cris_term cmd/ssh/*


FROM scratch AS prod

WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /source/cris_term /


COPY --from=build /source/content/ /content/
EXPOSE 42069
CMD ["/source/cris_term"]