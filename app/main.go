package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	metricUpdatePeriod := flag.Int64("update", 0, "Update metrics period")
	metricPort := flag.String("port", "0", "Metrics port")
	flag.Parse()

	// Function to check the presence of mandatory application flags
	if *metricUpdatePeriod == 0 || *metricPort == "0" {
		fmt.Println("ERROR! Added run params -update and -port")
		return
	}

	// Cycle to create metrics from application parameters
	for _, pName := range os.Args[5:] {
		go setMetric(*metricUpdatePeriod, pName, regMetric(pName))
	}

	// Running a web server to give back metrics
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+*metricPort, nil)
}

// Function to register the metric
func regMetric(pName string) prometheus.Gauge {
	metricName := pName + "_number_proccess"
	return promauto.NewGauge(prometheus.GaugeOpts{
		Name:        metricName,
		Help:        "Number of processes running",
		ConstLabels: prometheus.Labels{"Proccess_name": pName},
	})

}

// Function to register the metric
func setMetric(updateTime int64, pName string, metric prometheus.Gauge) {
	for {
		proc := "ps cax | grep " + pName + "|grep -v -e 'grep'| grep -v -e '" + os.Args[0] + "'|wc -l"
		result, _ := exec.Command("bash", "-c", proc).Output()
		metricStr := strings.ReplaceAll(strings.ReplaceAll(string(result), " ", ""), "\n", "")
		metricFloat, _ := strconv.ParseFloat(string(metricStr), 64)
		metric.Set(metricFloat)
		time.Sleep(time.Duration(updateTime) * time.Second)
	}
}
