FROM golang:1.19.1-alpine3.16 AS build
WORKDIR /src
COPY ./gatewayService/go.mod ./gatewayService/go.sum ./
RUN go mod download
COPY ./gatewayService ./
RUN go build -o ./gatewayService ./cmd
EXPOSE 8080
EXPOSE 8190


FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
COPY --from=build /src/gatewayService .
COPY ./gatewayService/config/config.yaml .
CMD [ "./gatewayService" ]




