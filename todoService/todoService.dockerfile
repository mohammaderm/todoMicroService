FROM golang:1.19.1-alpine3.16 AS build
WORKDIR /src
COPY todoService/go.mod todoService/go.sum ./
RUN go mod download
COPY ./todoService ./
RUN go build -o ./todoService ./cmd
EXPOSE 3306
EXPOSE 50053

FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
COPY --from=build /src/todoService .
COPY ./todoService/config/config.yaml .
CMD [ "./todoService" ]



