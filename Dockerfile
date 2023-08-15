# Build stage
FROM golang:1.20 AS build-env
WORKDIR /opt
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
# Final stage
FROM distroless/base as final
WORKDIR /opt
COPY --from=build-env /opt/app .
ENTRYPOINT ["./app"]
