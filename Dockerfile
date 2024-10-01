FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /source
COPY . .

RUN go mod download
RUN  go build -o /source/sshblog cmd/ssh/*


FROM scratch AS prod

WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /source/sshblog /


# COPY --from=build /source/content/ /content/
EXPOSE 42069
CMD ["/source/sshblog"]