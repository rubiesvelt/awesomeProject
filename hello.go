package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/golang/glog"
	"sort"
	"time"

	_ "gorm.io/driver/clickhouse"
)

// "SELECT *\nFROM \n(\n    SELECT\n        '所有用户' AS compare_user,\n        platid,\n\t\tdate_trunc('day', toDateTime(dteventtime)) gtime,\n        count(distinct multiIf(event = 'MoneyFlow', vopenid, NULL)) AS MoneyFlow_cnt,\n        count(distinct multiIf(event = 'ItemFlow', vopenid, NULL)) AS item_count\n    FROM clickhouse_441386_all\n    WHERE (dteventtime >= '{{begintime}}') AND (dteventtime < '{{endtime}}') AND (event IN ('MoneyFlow', 'ItemFlow'))\n    GROUP BY platid,date_trunc('day', toDateTime(dteventtime))\n    UNION ALL\n    SELECT\n        'izoneareaid=2009用户' AS compare_user,\n        platid,\n\t\tdate_trunc('day', toDateTime(dteventtime)) gtime,\n        count(distinct multiIf(event = 'MoneyFlow', vopenid, NULL)) AS MoneyFlow_cnt,\n        count(distinct multiIf(event = 'ItemFlow', vopenid, NULL)) AS item_count\n    FROM clickhouse_441386_all\n    WHERE (dteventtime >= '{{begintime}}') AND (dteventtime < '{{endtime}}') AND (izoneareaid = '2009') AND (event IN ('MoneyFlow', 'ItemFlow'))\n    GROUP BY platid,date_trunc('day', toDateTime(dteventtime))\n)\nORDER BY\n    platid ASC,\n    compare_user ASC"
// "SELECT\n        '所有用户' AS compare_user,\n        date_trunc('day', toDateTime(t1.dteventtime)) AS gtime,\n        t2.ilevel,\n        max(multiIf(t1.event = 'PlayerLogin', t1.exp6, NULL)) AS PlayerLoginmNum,\n        max(multiIf(t1.event = 'PlayerRegister', t1.exp6, NULL)) AS PlayerRegistermNum,\n        count(distinct multiIf(t1.event = 'PlayerLogin', t1.vopenid, NULL)) AS usernum\n    FROM clickhouse_441386_all AS t1    \n    INNER JOIN\n   (\n        select ilevel,vopenid from clickhouse_441372_all AS x where platid=255\n    ) t2 \n    ON t1.vopenid=t2.vopenid\n    WHERE (t1.dteventtime > '{{begintime}}') AND (t1.dteventtime < '{{endtime}}') AND (t1.event IN ('PlayerLogin', 'PlayerRegister')) and t1.vopenid in \n        (\n       select vopenid from clickhouse_441372_all AS x where ilevel>20 and platid=255\n        ) \n        and t1.vopenid in (select vopenid from clickhouse_441372_all AS x where ilevel=54 and platid=255) and toInt64(t1.exp6)>200\n    GROUP BY\n        t2.ilevel,\n        gtime  \n    SETTINGS max_memory_usage = 11000000000, distributed_product_mode = 'local'\n        union all \n        SELECT\n        'cptest' AS compare_user,\n        date_trunc('day', toDateTime(t1.dteventtime)) AS gtime,\n        t2.ilevel,\n        max(multiIf(t1.event = 'PlayerLogin', t1.exp6, NULL)) AS PlayerLoginmNum,\n        max(multiIf(t1.event = 'PlayerRegister', t1.exp6, NULL)) AS PlayerRegistermNum,\n        count(distinct multiIf(t1.event = 'PlayerLogin', t1.vopenid, NULL)) AS usernum\n    FROM clickhouse_441386_all AS t1    \n    INNER JOIN\n   (\n        select ilevel,vopenid from clickhouse_441372_all AS x where platid=255\n    ) t2 \n    ON t1.vopenid=t2.vopenid\n    WHERE (t1.dteventtime > '{{begintime}}') AND (t1.dteventtime < '{{endtime}}') AND (t1.event IN ('PlayerLogin', 'PlayerRegister')) and t1.vopenid in \n        (\n       select vopenid from clickhouse_441372_all AS x where ilevel>20 and platid=255\n        ) \n        and t1.vopenid in (select vopenid from clickhouse_441372_all AS x where ilevel=54 and platid=255) and toInt64(t1.exp6)>100\n    GROUP BY\n        t2.ilevel,\n        gtime  \n    SETTINGS max_memory_usage = 11000000000, distributed_product_mode = 'local'"
var (
	ip          = flag.String("ip", "clickhouse.dengta-test.cdp.db.", "")
	port        = flag.Int("port", 9000, "")
	user        = flag.String("user", "presto", "")
	password    = flag.String("password", "qazpresto#2021", "")
	concurrency = flag.Int("concurrency", 1, "")
	singleQuery = flag.Bool("singleQuery", false, "")

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
	fmt.Println("port", *port)
	fmt.Println("user", *user)
	fmt.Println("password", *password)
	fmt.Println("concurrency", *concurrency)

	sqls := []string{
		*sql1, *sql2, *sql3, *sql4, *sql5, *sql6,
	}

	queries := make([]string, 0)
	for _, u := range sqls {
		if u == "" {
			continue
		}
		queries = append(queries, u)
		fmt.Println("sql: ", u)
	}

	fmt.Println("queries len: ", len(queries))

	db, err := NewCHConnection(*port, *ip, *user, *password)
	if err != nil {
		fmt.Println("cluster connecte failed ", err)
		return
	}
	fmt.Println("cluster connected")

	if *singleQuery {
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

	c := make(chan int)

	start := time.Now()
	for i := 0; i < *concurrency; i++ {
		fmt.Println("query ", i, " started")
		go ConcurrentQuery(db, queries[i%len(queries)], c)
	}
	success := 0
	failed := 0
	for i := 0; i < *concurrency; i++ {
		x := <-c
		if x == 1 {
			failed++
		} else {
			success++
		}
	}
	elapsed := time.Since(start)
	fmt.Println("执行完成")
	fmt.Println("总计", *concurrency, "成功：", success, "失败：", failed)
	fmt.Println("耗时：", elapsed)
}

func ConcurrentQuery(db *sql.DB, query string, c chan int) {
	_, err := Query(db, query)
	if err != nil {
		fmt.Println(err)
		c <- 1
		return
	}
	c <- 0
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
