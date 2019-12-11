/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2017/3/27
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

import (
	"time"
	"testing"
	"fmt"
	"github.com/gradyZhang/summon-core/aliens/log"
)

//func test() {
//	SetData("key","data01")
//	fmt.Println(GetData("key"))
//
//	strs := []string {"a","q","w"}
//	LPush("strs",strs)
//	values := LRange("strs")
//	for _,v := range values {
//		fmt.Print(string(v))
//	}
//
//	imap,key := map[string]string{"k1":"job","k2":"toby","k3":"mark"},"m"
//	SetMap(key,imap)
//	fmt.Println(GetMap(key))
//
//	key = "testHset"
//	isSuccess := Hset(key,"fld","this is a hset")
//	if isSuccess {
//		fmt.Printf(Hget(key,"fld"))
//	}
//}

var glo *testing.T

func TestRedisCacheClient_ZAdd(t *testing.T) {
	redisClient := &RedisCacheClient{
		MaxIdle:5,
		MaxActive:20,
		Address:"127.0.0.1:20101",
		IdleTimeout:180 * time.Second,
	}
	redisClient.Start()
	redisClient.ZAdd("haha", 2, "s2")
	redisClient.ZAdd("haha", 4, "s4")
	redisClient.ZAdd("haha", 1, "s1")
	redisClient.ZAdd("haha", 5, "s5")
	//redisClient.ZAdds("haha", string[] {})
	result := redisClient.ZRevRangeWithScore("haha", 0, 6)
	log.Debug("%v", result)

}

func TestRedisCacheClient_GetData(t *testing.T) {
	redisClient := &RedisCacheClient{
		MaxIdle:5,
		MaxActive:20,
		Address:"127.0.0.1:20101",
		IdleTimeout:180 * time.Second,
	}
	redisClient.Start()
	redisClient.SetData("abc", "cba")
	t.Logf("get data %v", redisClient.GetData("abc"))
}

func TestRedisCacheClient_Publish(t *testing.T) {
	glo = t
	redisClient := &RedisCacheClient{
		MaxIdle:5,
		MaxActive:20,
		Address:"127.0.0.1:20101",
		IdleTimeout:180 * time.Second,
	}
	redisClient.Start()
	// The following function calls publish a message using another
	// connection to the Redis server.
	redisClient.Publish("example", "hello")
	redisClient.Publish("example", "world")
	redisClient.Publish("pexample", "foo")
	redisClient.Publish("pexample", "bar")
	time.Sleep(time.Second)

}


func TestRedisCacheClient_Sub(t *testing.T) {
	glo = t
	redisClient := &RedisCacheClient{
		MaxIdle:5,
		MaxActive:20,
		Address:"127.0.0.1:20101",
		IdleTimeout:180 * time.Second,
	}
	redisClient.Start()
	//psc := redis.PubSubConn{Conn :redisClient.pool.Get()}
	redisClient.Subscribe(scallback,"example")
	redisClient.PSubscribe(pscallback, "p*")
	time.Sleep(time.Second * 10)

}


func scallback(channel, data string) {
	fmt.Printf("sub  %v    %v", channel, data)
}

func pscallback(pattern, channel, data string) {
	fmt.Printf("sub %v  %v    %v",pattern, channel, data)
}

