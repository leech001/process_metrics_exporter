# Prometheus process metrics exporter

## Контейнер для подсчета количества запущенных процессов

docker run -d -p 8080:8080 process_metrics_exporter_app ./main 1 8080 proc