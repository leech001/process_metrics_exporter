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
		go updateMetric(*metricUpdatePeriod, pName)
	}

	// Running a web server to give back metrics
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+*metricPort, nil)
}

// Function to register and update the metric
func updateMetric(updateTime int64, pName string) {
	metric := promauto.NewGauge(prometheus.GaugeOpts{
		Name:        pName + "_number_proccess",
		Help:        "Number of processes running",
		ConstLabels: prometheus.Labels{"proccess_name": pName},
	})
	command := "ps waux | grep " + pName + "|grep -v -e 'grep'| grep -v -e '" + os.Args[0] + "'|wc -l"
	for {
		result, _ := exec.Command("bash", "-c", command).Output()
		valueStr := strings.ReplaceAll(strings.ReplaceAll(string(result), " ", ""), "\n", "")
		valueFloat, _ := strconv.ParseFloat(string(valueStr), 64)
		metric.Set(valueFloat)
		time.Sleep(time.Duration(updateTime) * time.Second)
	}
}
