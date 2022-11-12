FROM golang:1.19.1-alpine3.16 AS build
WORKDIR /src
COPY authService/go.mod authService/go.sum ./
RUN go mod download
COPY ./authService ./
RUN go build -o ./authService ./cmd
EXPOSE 50054
EXPOSE 5432


FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
COPY --from=build /src/authService .
COPY ./authService/config/config.yaml .
CMD [ "./authService" ]




