package clickhouse_pressure

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"math/rand"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func LoadExporterMain() {
	http.Handle("/metrics", promhttp.Handler())

	loadExporter := NewLoadExporter()
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(loadExporter)
	http.Handle("/metricsLoad", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	http.ListenAndServe(":12123", nil)
}

type LoadExporter struct {
}

var (
	myGague = promauto.NewGauge(prometheus.GaugeOpts{
		Name:        "current_cpu_usage",
		Help:        "current cpu usage",
		ConstLabels: map[string]string{},
	})
)

func NewLoadExporter() *LoadExporter {
	exporter := &LoadExporter{}
	return exporter
}

func (e *LoadExporter) Describe(ch chan<- *prometheus.Desc) {
	metricCh := make(chan prometheus.Metric)
	doneCh := make(chan struct{})

	go func() {
		for m := range metricCh {
			ch <- m.Desc()
		}
		close(doneCh)
	}()

	e.Collect(metricCh)
	close(metricCh)
	<-doneCh
}

func (e *LoadExporter) Collect(ch chan<- prometheus.Metric) {
	f := getCurrentCpu()
	myGague.Set(f)
	myGague.Collect(ch)
	fmt.Println(f)
}

// recordMetrics 定时后台获取状态
func recordMetrics() {
	go func() {
		for {
			f := getCurrentCpu()
			fmt.Println(f)
			myGague.Set(f)
			time.Sleep(1000 * time.Millisecond)
		}
	}()
}

func getCurrentCpu() float64 {
	f := execTop()
	if f == nil {
		return -1
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.HasPrefix(text, "%Cpu") {
			continue
		}
		strs := strings.Split(text, ",")
		if len(strs) != 8 {
			continue
		}
		idleStr := strs[3]
		subStrs := strings.Split(idleStr, " ")
		if len(subStrs) != 3 {
			continue
		}
		t, _ := strconv.ParseFloat(subStrs[1], 64)
		t, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", 100-t), 64)
		return t
	}
	return -1
}

func execTop() io.Reader {
	cmd := exec.Command("top", "-bn 1", `-i`, `-c`)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return nil
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("execute command err,", err)
		return nil
	}

	return stdout
}

var (
	nodes       = flag.String("nodes", "9.134.77.2", "")
	collectFreq = flag.Int("collectFreq", 100, "")
)

func getCpuUsageMain() {
	flag.Parse()

	// 并行收集每个节点的cpu使用率
	ips := strings.Split(*nodes, ",")
	n := len(ips)
	freq := float64(*collectFreq)
	c := make(chan float64)
	mp = make(map[string]float64)

	fmt.Println("ips", ips)
	fmt.Println("freq", freq)

	go addLoad()

	for true {
		start := time.Now()
		for _, ip := range ips {
			go getCpuUsage(ip, freq, c)
		}
		res := make([]float64, 0)
		for i := 0; i < n; i++ {
			x := <-c
			res = append(res, x)
		}
		fmt.Println(getAvg(res))
		t := start.Add(time.Duration(freq) * time.Second).Sub(time.Now())
		if t < 0 {
			fmt.Printf("can't get result in %f seconds\n", freq)
			return
		}
		time.Sleep(t)
	}
	fmt.Println("oh shit it stops man")
}

var mp map[string]float64 // ip -> 上次 idle 时间 sum 值

func addLoad() {
	fmt.Println("start add load")
	var t float64
	t = rand.ExpFloat64()
	for true {
		go addLoad()
		time.Sleep(1)
		t *= 14324234
		t /= 214937.2984792
		t += 2419837.234341
		t--
		if t > 34729814523452345237927 {
			t = rand.ExpFloat64()
		}
	}
}

func getCpuUsage(ip string, freq float64, c chan float64) float64 {
	resp, err := http.Get("http://" + ip + ":11123/metrics")
	if err != nil {
		// handle error
		fmt.Println("connect to", ip, "error")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("close body to", ip, "error")
		}
	}(resp.Body)

	var sum float64
	var num float64 // 核心数
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		text := scanner.Text()
		strs := strings.Split(text, " ")
		if len(strs) != 2 {
			continue
		}
		if !strings.HasPrefix(strs[0], "node_cpu_seconds_total") {
			continue
		}
		if !strings.HasSuffix(strs[0], "mode=\"idle\"}") {
			continue
		}
		var floatNum float64
		_, err := fmt.Sscanf(strs[1], "%e", &floatNum)
		if err != nil {
			continue
		}
		sum += floatNum
		num++
	}
	if num == 0 {
		c <- 0
		return 0
	}
	if _, ok := mp[ip]; !ok {
		fmt.Println("start getting cpu usage")
		mp[ip] = sum
		c <- 0
		return 0
	}
	idleAvgTimePerSecond := (sum - mp[ip]) / num / freq
	usage := 1 - idleAvgTimePerSecond
	mp[ip] = sum
	c <- usage
	return usage
}

func getAvg(arr []float64) float64 {
	n := len(arr)
	var sum float64
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	return sum / float64(n)
}
