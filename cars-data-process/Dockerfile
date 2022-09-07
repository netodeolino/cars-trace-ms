FROM golang:1.19.0 AS build

WORKDIR /app

COPY services/process/process.go ./services/process/
COPY services/reader/reader.go ./services/reader/
COPY services/splitter/splitter.go ./services/splitter/
COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN go build -o /server


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY receiver/car-data-example.csv ./receiver/
COPY --from=build /server /server

USER nonroot:nonroot

ENV RECEIVED_CSV_PATH=receiver/car-data-example.csv

ENTRYPOINT [ "/server" ]