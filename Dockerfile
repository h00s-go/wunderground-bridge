FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS build

ARG TARGETOS

ARG TARGETARCH

ARG TARGETVARIANT

WORKDIR /src

COPY . ./

RUN go mod download && \
    GOOS=$TARGETOS GOARCH=$TARGETARCH GOARM=${TARGETVARIANT#v} \
    go build -o /out/wunderground-bridge

FROM alpine

COPY --from=build /out/wunderground-bridge /bin

EXPOSE 8080

CMD [ "/bin/wunderground-bridge" ]
