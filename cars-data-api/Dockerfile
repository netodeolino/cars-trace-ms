FROM golang:1.19.0 AS build

WORKDIR /app

COPY request/request.go ./request/
COPY routes/routes.go ./routes/
COPY service/service.go ./service/
COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN go build -o /server


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]