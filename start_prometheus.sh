# Start Prometheus using Docker
docker run --rm -d \
  -p 9090:9090 \
  -v "$(pwd)/default-prometheus.yml:/etc/prometheus/prometheus.yml" \
  --name prometheus \
  prom/prometheus \
  --config.file=/etc/prometheus/prometheus.yml --enable-feature=remote-write-receiver
