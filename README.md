# Wunderground Bridge

[![Go Report Card](https://goreportcard.com/badge/github.com/h00s-go/wunderground-bridge)](https://goreportcard.com/report/h00s-go/wunderground-bridge) [![license](https://img.shields.io/github/license/mashape/apistatus.svg)]()

Wunderground Bridge is a web API (bridge) between Renkforce WH2600 Weather Station (also known as Ambient Weather 1400-IP) and Wunderground. Renkforce WH2600 is quality and affordable weather station. The only problem that data from this station can be uploaded only to Wunderground servers. This bridge allows you to gather data from weather station (for additional processing and storing) and optionally can pass data to Wunderground so you can keep that feature.

It has support for:
- API for receiving data from weather station and optionally passing it to Wunderground
- API for getting current weather data gathered from weather station
- Publishing received data to MQTT broker
- Websocket API for getting current weather data (in development)
- Watchdog for restarting weather station when it fails to receive data for configured times

## Installation

Easiest way for running this bridge is using Docker. You can find Docker image on Docker Hub. You can also build it yourself using Dockerfile in this repository.

### Minimal setup

Receive data from weather station and pass it to Wunderground.

```yaml
version: '3.5'
services:
  wunderground-bridge:
    image: h00s/wunderground-bridge:latest
    container_name: wunderground-bridge
    restart: always
    ports:
      - "8080:8080"
    environment:
      - STATION_ID=IABCDEF
      - STATION_PASSWORD=xyz12345
```

### Minimal setup with watchdog

Receive data from weather station and pass it to Wunderground with watchdog. Watchdog will restart weather station when it fails to receive data for configured times.

**NOTE**: Weather station should be connected to the same network as this bridge (url to weather station is needed).

```yaml
version: '3.5'
services:
  wunderground-bridge:
    image: h00s/wunderground-bridge:latest
    container_name: wunderground-bridge
    restart: always
    ports:
      - "8080:8080"
    environment:
      - STATION_ID=IABCDEF
      - STATION_PASSWORD=xzy12345
      - STATION_URL=http://192.168.0.99
      - STATION_WATCHDOG_ENABLED=true
```

### Full setup

Receive data from weather station, publish it to MQTT and pass it to Wunderground. Watchdog will restart weather station when it fails to receive data for configured times.

**NOTE**: Weather station should be connected to the same network as this bridge (url to weather station is needed).

```yaml
version: '3.5'
services:
  wunderground-bridge:
    image: h00s/wunderground-bridge:latest
    container_name: wunderground-bridge
    restart: always
    ports:
      - "8080:8080"
    environment:
      - STATION_ID=IABCDEF
      - STATION_PASSWORD=xzy12345
      - STATION_URL=http://192.168.0.99
      - STATION_WATCHDOG_ENABLED=true
      - MQTT_ENABLED=true
      - MQTT_BROKER=1.2.3.4:1883
      - MQTT_USERNAME=user
      - MQTT_PASSWORD=password
      - MQTT_CLIENT_ID=wunderground-bridge
      - MQTT_UPDATE_TOPIC=weather-station/update
```