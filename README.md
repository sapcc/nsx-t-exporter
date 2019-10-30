# Prometheus NSX-T Management Node Exporter

Exposes metrics from NSX-T Management Node REST API to a Prometheus compatible endpoint.

## Exporter Configuration

NSX-T Expoerter takes input from environment variables as:

### Mandatory Variables
- `NSXV3_LOGIN_HOST` NSX-T Manager Node hostname or IP address.
- `NSXV3_LOGIN_PORT` NSX-T Manager Node port.
- `NSXV3_LOGIN_USER` NSX-T Manager Node login user.
- `NSXV3_LOGIN_PASSWORD` NSX-T Manager Node login password.

### Optional Variables
- `NSXV3_REQUESTS_PER_SECOND` (10) NSX-T Manager Node requestes per second [<100].
- `NSXV3_SUPPRESS_SSL_WORNINGS` (false) NSX-T Manager Node disables ssl host validattion.
- `LOG_LEVEL` (debug) NSX-T Exporter logging level.

## Build

```sh
git clone https://github.com/sapcc/github.com/sapcc/nsx-t-exporter.git
cd nsx-t-exporter
docker build -t <image-name> .
```

## Run (Simple)

Edit docker-compose.yml
```yaml
version: "2"

services:
  nsxv3-exporter:
    tty: true
    stdin_open: true
    expose:
      - 9191
    ports:
      - 9191:9191
    image: <image-name> 
    environment:
      - NSXV3_LOGIN_HOST=<NSX-T Manager Node hostname or IP address>
      - NSXV3_LOGIN_PORT=<NSX-T Manager Node port>
      - NSXV3_LOGIN_USER=<NSX-T Manager Node login user>
      - NSXV3_LOGIN_PASSWORD=<NSX-T Manager Node login password>
      - NSXV3_REQUESTS_PER_SECOND=<NSX-T Manager Node requestes per second [<100]>
      - NSXV3_CONNECTION_POOL_SIZE=<NSX-T Manager Node connection pool size>
      - NSXV3_REQUEST_TIMEOUT_SECONDS=<NSX-T Manager request timeout in seconds>
      - NSXV3_SUPPRESS_SSL_WORNINGS=<NSX-T Manager Node disables ssl host validattion>
      - SCRAP_PORT=<The exporter scrap port>
      - SCRAP_SCHEDULE_SECONDS=<The exporter scrap NSX-T Manager schedule in seconds>
      - LOG_LEVEL=<NSX-T Exporter logging level.>

```

## Metrics

Metrics will be made available on http://<docker_host>:9191/metrics
Metrics export can be seen at METRICS.md