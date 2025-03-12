FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/webapp

FROM scratch
COPY --from=build /bin/demo /bin/webapp
ENTRYPOINT ["/bin/webapp"]
