package 缓存

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"testing"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(pong)
}

func TestRedis(t *testing.T) {
	// set
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// get
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

// 管道
func TestPipeliner(t *testing.T) {
	var a *redis.StringCmd
	var aa *redis.StringCmd
	cmds, err := client.Pipelined(func(pipe redis.Pipeliner) error {
		a = pipe.Get("a")
		aa = pipe.Get("aa")
		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("len:", len(cmds))
	fmt.Println(a.Val())
	fmt.Println(aa.Val())
}

func TestRedis2(t *testing.T) {
	// SET key value EX 10 NX
	set, err := client.SetNX("key", "value", 10*time.Second).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(set)

}

// list,堆栈
func TestRedis3(t *testing.T) {
	//	插入数据
	//c, err := client.LPush("list", "11","22","33").Result()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(c)
	//
	//strs,err := client.LRange("list", 0, 2).Result()
	//panicErr(err)
	//fmt.Println(strs)

	// LPop 获取第一个元素，RPop获取最后一个元素
	strs, err := client.LPop("list").Result()
	panicErr(err)
	fmt.Println(strs)

	//strs,err := client.Sort("list", &redis.Sort{Order:"value"}).Result()
	//panicErr(err)
	//fmt.Println(strs)
}

// hash，哈希
func TestRedis4(t *testing.T) {
	//m := map[string]interface{}{}
	//m["name"] = "xiaoming"
	//_,err := client.HMSet("hash", m).Result()
	//panicErr(err)

	// 获取单个值
	//mm,err := client.HMGet("hash","key").Result()
	//panicErr(err)
	//fmt.Println(mm)

	// 获取对象
	mm, err := client.HGetAll("hash").Result()
	panicErr(err)
	fmt.Println(mm)
	fmt.Println(mm["key"])
	fmt.Println(mm["name"])
}

// set，集合
func TestRedis5(t *testing.T) {
	// 添加
	//_,err := client.SAdd("jihe","mm").Result()
	//panicErr(err)
	//_,err = client.SAdd("jihe","oo","cc","dd","lwem").Result()
	//panicErr(err)

	// 获取
	//strs,err :=client.SMembers("jihe").Result()
	//panicErr(err)
	//fmt.Println(strs)

	// 有序集合
	//err = client.ZAdd("youxu",
	//	&redis.Z{
	//		Score:0,
	//		Member:"str",
	//	},&redis.Z{
	//		Score:0,
	//		Member:"xiaoming",
	//	},&redis.Z{
	//		Score:0,
	//		Member:"qiangge",
	//	},&redis.Z{
	//		Score:0,
	//		Member:"bao",
	//	}).Err()
	//panicErr(err)

	// 获取
	strs, err := client.SMembers("youxu").Result()
	panicErr(err)
	fmt.Println(strs)

	client.ZRangeByScore("youxu", &redis.ZRangeBy{})

}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}

}
