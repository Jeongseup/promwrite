package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"time"

	"github.com/castai/promwrite/promwrite"
	promvalue "github.com/prometheus/prometheus/model/value"
)

const promMetricNameLabel = "__name__"

var metricsName = "my_sequential_metric"

var metricsLabels = []promwrite.Label{
	{
		Name:  promMetricNameLabel,
		Value: metricsName,
	},
	{
		Name:  "label1",
		Value: "this",
	},
	{
		Name:  "label2",
		Value: "is",
	},
	{
		Name:  "label3",
		Value: "test",
	},
	{
		Name:  "label4",
		Value: "labels",
	},
	// it'll be injected a rand hex string
}

func main() {
	// Get the current time
	startTime := time.Now()

	// Print the result
	fmt.Println("Start Time:", startTime)

	// Create client
	client := promwrite.NewClient("http://localhost:9090/api/v1/write")

	// Loop from oneHourAgo to currentTime, incrementing by one second each iteration
	for t, push_value := startTime, 123; !t.After(startTime.Add(10 * time.Second)); t, push_value = t.Add(1*time.Second), push_value+1 {
		newMetricsLabels := append(metricsLabels, promwrite.Label{
			Name:  "rand_label",
			Value: generateRandomHexString(),
		})
		// fmt.Println(newMetricsLabels)
		resp, err := client.Write(context.Background(), &promwrite.WriteRequest{
			TimeSeries: []promwrite.TimeSeries{
				{
					Labels: newMetricsLabels,
					Sample: promwrite.Sample{
						Time:  t,
						Value: float64(push_value),
					},
				},
			},
		})

		if err != nil {
			fmt.Println("Error writing to Prometheus:", err)
		} else {
			fmt.Println("Successfully wrote to Prometheus:", resp)
		}

		// Sleep for a second to avoid overwhelming the server
		time.Sleep(1 * time.Second)

		// Send Stale Markers
		// Stale markers MUST be signalled by the special NaN value 0x7ff0000000000002. This value MUST NOT be used otherwise.
		// ref; https://prometheus.io/docs/specs/remote_write_spec/
		// https://github.com/prometheus/prometheus/blob/main/model/value/value.go
		if t.After(startTime.Add(9 * time.Second)) {
			fmt.Println("become over 10 second, so it'll send stale marker for prometheus")

			// Send Stale Marker after 10 seconds
			staleTime := startTime.Add(11 * time.Second)
			staleValue := math.Float64frombits(promvalue.StaleNaN)

			resp, err := client.Write(context.Background(), &promwrite.WriteRequest{
				TimeSeries: []promwrite.TimeSeries{
					{
						Labels: metricsLabels,
						Sample: promwrite.Sample{
							Time:  staleTime,
							Value: staleValue,
						},
					},
				},
			})

			if err != nil {
				fmt.Println("Error writing to Prometheus:", err)
			} else {
				fmt.Println("Successfully wrote to Prometheus:", resp)
			}

		}
	}
}

func generateRandomHexString() string {
	timestamp := time.Now().Unix()
	bytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bytes[i] = byte(timestamp >> (i * 8))
	}
	return hex.EncodeToString(bytes)
}
