package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/golang/glog"
	"sort"
	"strings"
	"time"

	_ "gorm.io/driver/clickhouse"
)

var (
	ip             = flag.String("ip", "", "")
	ips            = flag.String("ips", "", "")
	port           = flag.Int("port", 9000, "")
	user           = flag.String("user", "presto", "")
	password       = flag.String("password", "qazpresto#2021", "")
	concurrency    = flag.Int("concurrency", 1, "")
	singleQuery    = flag.Bool("singleQuery", false, "")
	minorBeginTime = flag.String("minorBeginTime", "", "") // minorTime 范围的，请求比例 1/6
	minorEndTime   = flag.String("minorEndTime", "", "")
	beginTime      = flag.String("beginTime", "2021-11-14 00:00:00", "")
	endTime        = flag.String("endTime", "2021-11-21 00:00:00", "")

	query = flag.String("q", "show databases", "")
	sql1  = flag.String("q1", "", "")
	sql2  = flag.String("q2", "", "")
	sql3  = flag.String("q3", "", "")
	sql4  = flag.String("q4", "", "")
	sql5  = flag.String("q5", "", "")
	sql6  = flag.String("q6", "", "")
)

func main() {
	flag.Parse()

	fmt.Println("ip", *ip)
	fmt.Println("ips", *ips)
	fmt.Println("port", *port)
	fmt.Println("user", *user)
	fmt.Println("password", *password)
	fmt.Println("concurrency", *concurrency)
	fmt.Println("minorBeginTime", *minorBeginTime)
	fmt.Println("minorEndTime", *minorEndTime)
	fmt.Println("beginTime", *beginTime)
	fmt.Println("endTime", *endTime)

	sqls := []string{
		*sql1, *sql2, *sql3, *sql4, *sql5, *sql6,
	}

	minorQueries := make([]string, 0)
	for _, u := range sqls {
		if u == "" {
			continue
		}
		if *minorBeginTime != "" && *minorEndTime != "" {
			replaceQueryTime(&u, minorBeginTime, minorEndTime)
			minorQueries = append(minorQueries, u)
			fmt.Println("minor sql: ", u)
		}
	}

	queries := make([]string, 0)
	for _, u := range sqls {
		if u == "" {
			continue
		}
		replaceQueryTime(&u, beginTime, endTime)
		queries = append(queries, u)
		fmt.Println("sql: ", u)
	}

	fmt.Println("queries len: ", len(queries))

	db, err := NewCHConnection(*port, *ip, *user, *password)
	if err != nil {
		fmt.Println("cluster connect failed ", err)
		return
	}
	fmt.Println("cluster connected")

	if *singleQuery {
		replaceQueryTime(query, beginTime, endTime)
		fmt.Println("execute single query: ", *query)
		ret, err := Query(db, *query)
		if err != nil {
			fmt.Println("cluster connecte failed ", err)
			return
		}
		columns, err := ret.Columns()
		fmt.Println(columns)
		return
	}

	c := make(chan QueryResult)

	start := time.Now()

	var minorQueriesCnt = 0
	if len(minorQueries) > 0 {
		minorQueriesCnt = *concurrency/6 + 1
	}
	fmt.Println("minorQueriesCnt", minorQueriesCnt)
	for i := 0; i < *concurrency; i++ {
		if minorQueriesCnt > 0 {
			fmt.Println("minor query ", i, " started")
			go ConcurrentQuery(db, minorQueries[i%len(minorQueries)], c)
			minorQueriesCnt--
			continue
		}
		fmt.Println("query ", i, " started")
		go ConcurrentQuery(db, queries[i%len(queries)], c)
	}

	res := make([]QueryResult, 0)

	for i := 0; i < *concurrency; i++ {
		x := <-c
		res = append(res, x)
	}
	elapsed := time.Since(start)
	summary := makeSummary(res)

	fmt.Println("执行完成")
	fmt.Println("开始时间: ", *beginTime)
	fmt.Println("结束时间: ", *endTime)
	fmt.Println("总计: ", summary.total, "成功: ", summary.successNum, "失败: ", summary.failedNum)
	fmt.Println("耗时: ", elapsed)
	fmt.Println("min: ", summary.min)
	fmt.Println("max: ", summary.max)
	fmt.Println("avg: ", summary.avgSuccess)
	fmt.Println("tp50: ", summary.tp50Success)
	fmt.Println("tp90: ", summary.tp90Success)
	fmt.Println("detail: ")
	for _, u := range summary.successDetail {
		fmt.Println(u)
	}
}

type Summary struct {
	min           int64
	max           int64
	total         int
	successNum    int
	failedNum     int
	avgSuccess    int64
	tp50Success   int64
	tp90Success   int64
	successDetail []int64
}

func makeSummary(res []QueryResult) Summary {

	succ := make([]int64, 0)
	fail := make([]int64, 0)

	var totalCost int64
	totalCost = 0
	for _, u := range res {
		if u.status == 0 {
			t := u.elapsed.Milliseconds()
			totalCost += t
			succ = append(succ, t)
			continue
		}
		fail = append(fail, u.elapsed.Milliseconds())
	}
	sort.Slice(succ, func(i, j int) bool { return succ[i] < succ[j] })

	if len(succ) == 0 {
		return Summary{failedNum: len(fail)}
	}
	total := len(res)
	successNum := len(succ)
	failedNum := len(fail)
	avgSuccess := totalCost / int64(successNum)
	tp50Success := succ[successNum/2]
	tp90Success := succ[successNum*9/10-1]

	summary := Summary{
		min:           succ[0],
		max:           succ[successNum-1],
		total:         total,
		successNum:    successNum,
		failedNum:     failedNum,
		avgSuccess:    avgSuccess,
		tp50Success:   tp50Success,
		tp90Success:   tp90Success,
		successDetail: succ,
	}
	return summary
}

func replaceQueryTime(query *string, beginTime *string, endTime *string) {
	*query = strings.Replace(*query, "{{begintime}}", *beginTime, -1)
	*query = strings.Replace(*query, "{{endtime}}", *endTime, -1)
}

type QueryResult struct {
	status  int
	elapsed time.Duration
}

func ConcurrentQuery(db *sql.DB, query string, c chan QueryResult) {
	start := time.Now()
	// fmt.Println("executing query: ", query)
	_, err := Query(db, query)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Println(err)
		res := QueryResult{
			status:  1,
			elapsed: elapsed,
		}
		c <- res
		return
	}
	res := QueryResult{
		status:  0,
		elapsed: elapsed,
	}
	c <- res
}

func Query(conn *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func NewCHConnection(port int, ip, user, passwd string) (*sql.DB, error) {
	var (
		conn = &sql.DB{}
		err  error
	)

	if len(user) > 0 && len(passwd) > 0 {
		conn, err = sql.Open("clickhouse", "tcp://"+ip+":"+fmt.Sprint(port)+"?username="+user+"&password="+passwd+"&debug=false&read_timeout=30m")
	} else {
		conn, err = sql.Open("clickhouse", "tcp://"+ip+":"+fmt.Sprint(port)+"?debug=false&read_timeout=30m")
	}
	if err != nil {
		glog.Errorf("Open connection to %s failed, error %s", ip, err)
		return nil, err
	}
	glog.Infof("Open connection to %s success", ip)

	if err := conn.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			glog.Errorf("Exception code: [%d], %s, %s", exception.Code, exception.Message, exception.StackTrace)
			conn.Close()
			return nil, err
		} else {
			glog.Errorf("%s", err)
		}
	} else {
		glog.Infof("Ping %s successfuly", ip)
	}
	return conn, nil
}

// 容器
// Map, Set, Stack, Queue, Deque, List
// Set, SortedSet, TreeSet
// List, list.sort(), Comparator

/**
5935. 适合打劫银行的日子

security = [5,3,3,3,5,6,2], time = 2
-> [2,3]

*/
func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	left := make([]int, n) // 左边有几个大于等于 i 这个且递减的
	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if security[i+1] >= security[i] {
			right[i] = right[i+1] + 1
		}
	}
	var ans []int // 定义不限大小的空数组
	for i := 0; i < n; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return ans
}

/**
5934. 找到和最大的长度为 K 的子序列

输入：nums = [2,1,3,3], k = 2
输出：[3,3]

输入：nums = [-1,-2,3,4], k = 3
输出：[-1,3,4]
*/
func maxSubsequence(nums []int, k int) []int {
	id := make([]int, len(nums))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id,
		func(i, j int) bool {
			return nums[id[i]] > nums[id[j]]
		})
	sort.Ints(id[:k])
	ans := make([]int, k)
	for i, j := range id[:k] {
		ans[i] = nums[j]
	}
	return ans
}
