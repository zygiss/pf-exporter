package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	up = prometheus.NewDesc(
		"pf_up",
		"Was executing pfctl successful",
		nil, nil,
	)
)

type PfCollector struct{}

func main() {
	c := PfCollector{}
	prometheus.MustRegister(c)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (c PfCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
}

func (c PfCollector) Collect(ch chan<- prometheus.Metric) {
	filename := "/home/zygis/tmp/pf.stats"
	data, err := getInput(filename)
	if err != nil {
		fmt.Println("Err: %v\n", err)
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
		return
	}
	ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 1)
	fileByLine := strings.Split(string(data), "\n")
	fmt.Println(fileByLine)
}

// getStats reads PF stats from a file generated with `pfctl -vs info`
// for dev purposes.  When done, it'll read stats from `pfctl` directly.
func getStats(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	return data, err
}
