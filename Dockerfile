FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /source
COPY . .

RUN go mod tidy
RUN go build -o /source/cris_term cmd/ssh/*


FROM scratch AS prod

WORKDIR /
COPY --from=build /source/cris_term /
# COPY --from=build /source/content/ /content/
EXPOSE 42069
ENTRYPOINT ["/cris_term"]