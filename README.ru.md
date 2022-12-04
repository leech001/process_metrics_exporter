# Экспотер количества запущенных процессов для Prometheus
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/leech001/process_metrics_exporter/Go%20release%20builder?label=BUILD%20AND%20RELEASE&logo=github) ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/leech001/process_metrics_exporter/Publish%20Docker%20image?label=BUILD%20AND%20PUBLISH%20TO%20DOCKERHUB&logo=github) [![GitHub](https://img.shields.io/badge/Git-Hub-purple.svg)](https://github.com/leech001/process_metrics_exporter) [![Docker](https://img.shields.io/badge/Docker-hub-2496ed.svg)](https://hub.docker.com/r/leech001/process-metrics-exporter) [![License: WTFPL](https://img.shields.io/badge/license-WTFPL-brightgreen)](https://github.com/leech001/process_metrics_exporter/blob/main/LICENSE)  

## Программа на GO котороя считает количество запущенных процессов одного приложения и выводит их количество как Gauge метрику для Prometheus
---
### Пример с готовым (бинарным) файлом программы 

Не забудьте установить перед запуском bash в систему

```
procheck-linux-amd64 -update 5 -port 8080 test1 test2 test3
```

где:

- "procheck-linux-amd64" - бинарный файл приложения (береться из релизов для репозитория под вашу платформу);

- "-update 5" - частота обновления метрики в секундах;

- "-port 8080" - порт на котором отдаются метрики (http://localhost:8080/metrics);

- "test1...3" - имена процессов которые мониторим (количство ограниченно только вашей фантазией и памятью машинки).
---
### Запуск из готового Docker образа (Alpine Linux images)

Скачиваем и запускаем одной командой

```
docker run -d -p 8080:8080 leech001/process-metrics-exporter ./procheck -update 5 -port 8080 test1 test2 test3
```

где:

- "-update 5" - частота обновления метрики в секундах;

- "-port 8080" - порт на котором отдаются метрики;

- "test1...3" - имена процессов которые мониторим (количство ограниченно только вашей фантазией и памятью машинки).

метрики доступны по следующему адресу

```
http://localhost:8080/metrics
```
