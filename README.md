# Prometheus process metrics exporter

## Контейнер для подсчета количества запущенных процессов и формирования соотвествующих метрик для Prometheus


docker run -d -p 8080:8080 process_metrics_exporter_app ./main -update 5 -port 8080 test1 test2 test3