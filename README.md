# promwrite

Prometheus Remote Write Go client with minimal dependencies. Supports Prometheus, Cortex, VictoriaMetrics etc.

### Update

- Add stale markers
  // Send Stale Markers
  // Stale markers MUST be signalled by the special NaN value 0x7ff0000000000002. This value MUST NOT be used otherwise.
  // ref; https://prometheus.io/docs/specs/remote_write_spec/
  // https://github.com/prometheus/prometheus/blob/main/model/value/value.go

### Install

```
go get github.com/castai/promwrite
```

### Example Usage

```go
client := promwrite.NewClient("http://prometheus:8428/api/v1/write")
resp, err := client.Write(context.Background(), &promwrite.WriteRequest{
	TimeSeries: []promwrite.TimeSeries{
		{
			Labels: []promwrite.Label{
				{
					Name:  "__name__",
					Value: "my_metric_name",
				},
			},
			Sample: promwrite.Sample{
				Time:  time.Now(),
				Value: 123,
			},
		},
	},
})
```

```bash
# start local prom
./start_prometheus.sh

# push test data
go run .
```
