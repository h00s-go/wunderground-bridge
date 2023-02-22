FROM golang:1.20-alpine AS build

WORKDIR /src

COPY . ./

RUN go mod download && \
    go build -o /out/wunderground-bridge

FROM alpine

COPY --from=build /out/wunderground-bridge /bin

EXPOSE 8080

CMD [ "/bin/wunderground-bridge" ]