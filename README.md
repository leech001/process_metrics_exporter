# Prometheus process metrics exporter
[![GitHub](https://img.shields.io/badge/Git-Hub-purple.svg)](https://github.com/leech001/process_metrics_exporter) [![Docker](https://img.shields.io/badge/Docker-hub-2496ed.svg)](https://hub.docker.com/r/leech001/process-metrics-exporter) [![License: WTFPL](https://img.shields.io/badge/license-WTFPL-brightgreen)](https://github.com/leech001/process_metrics_exporter/blob/main/LICENSE)  

## GO application for counting the number of running processes and forming the corresponding metrics for Prometheus (Alpine Linux images)

The main task is to calculate the number of running processes for a given list of processes, use it to form Gauge metrics and display them on the page with metrics

The application is launched with a command in the 'app' folder

```
go run main -update 5 -port 8080 test1 test2 test3
```

where:

- "main" - application name;

- "-update 5" - metrics refresh period on sec;

- "-port 8080" - port where metrics are fed;

- "test1...3" - names of processes.


## Work with Docker (Alpine Linux images)

Download prepared image with application

```
docker pull leech001/process-metrics-exporter
```

or start it at once

```
docker run -d -p 8080:8080 leech001/process-metrics-exporter ./procheck -update 5 -port 8080 test1 test2 test3
```

metrics are available at

```
http://localhost:8080/metrics
```
