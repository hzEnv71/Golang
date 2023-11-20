package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("connect redis failed ,err:%v\n", err)
	}
	fmt.Println("连接成功")
	redis1()
	//redis2()
}

var (
	redisdb *redis.Client
	//ctx     = context.Background()
)

func initClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, //适用默认数据库
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// key
func redis1() {
	ok, err := redisdb.Set("java", 999, time.Duration(time.Second*5)).Result()
	if err != nil {
		fmt.Printf("set failed err=%v\n", err)
	}
	fmt.Println(ok)
	value, err := redisdb.Get("java").Result()
	if err != nil {
		fmt.Printf("get failed err=%v\n", err)
	}
	fmt.Println(value)
	values, err := redisdb.Keys("*").Result()
	if err != nil {
		fmt.Printf("keys failed err=%v\n", err)
	}
	fmt.Println(values)
	//自定义命令
	res, err := redisdb.Do("set", "golang", "999").Result()
	if err != nil {
		fmt.Printf("Do failed err=%v\n", err)
	}
	fmt.Println(res)
	//通配符
	//ctx := context.Background()
	iter := redisdb.Scan(0, "*", 10).Iterator()
	for i := 0; iter.Next(); {
		fmt.Println(i, iter.Val())
		i++
	}
	//pipeline
	pipe := redisdb.Pipeline()

	pipe.Incr("pipeline_counter") //将 key 中储存的数字值增一

	pipe.Expire("golang", time.Second*5) //设置过期时间

	cmder, err := pipe.Exec()

	if err != nil {
		fmt.Println("pipeline err=", err)
	}
	for i, ret := range cmder {
		fmt.Println(i, "    ", ret.Name(), "     ", ret.Args())

	}
	//事务
	txpipe := redisdb.TxPipeline()
	incr := txpipe.Incr("tx_pipeline_counter")
	flag := txpipe.Expire("tx_pipeline_counter", time.Hour)

	_, err = txpipe.Exec()
	if err != nil {
		fmt.Println("tx:", err)
	}

	fmt.Println(incr.Val(), flag.Val())

	//watch
	// 监视watch_count的值，并在值不变的前提下将其值+1
	key := "watch_count"
	err = redisdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)

	value, err = redisdb.Get("watch_count").Result()
	if err != nil {
		fmt.Printf("get failed err=%v\n", err)
	}
	fmt.Println(value)

}

// zset(sorted set：有序集合)
func redis2() {
	zsetKey := "language_rand"
	language := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 97.0, Member: "Python"},
		redis.Z{Score: 99.0, Member: "Rust"},
	}
	//ZADD
	num, err := redisdb.ZAdd(zsetKey, language...).Result()
	if err != nil {
		fmt.Printf("zadd failed,err%v\n", err)
	}
	fmt.Printf("zadd  %d success\n", num)
	//Golang分数+10
	newScore, err := redisdb.ZIncrBy(zsetKey, 10, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed,err%v\n", err)
	}
	fmt.Printf("zincrby  %f success\n", newScore)
	// 取分数最高的3个
	ret, err := redisdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

}
