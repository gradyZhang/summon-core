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
	"aliens/log"
	"fmt"
	"github.com/FZambia/sentinel"
	"github.com/garyburd/redigo/redis"
	"time"
)

//操作类型
const (
	PARAM_WITHSCORES string = "WITHSCORES"

	OP_SET string = "SET"
	OP_GET string = "GET"
	OP_DEL string = "DEL"
	OP_INCR string = "INCR"

	OP_FLUSHALL string = "FLUSHALL"
	OP_FLUSHDB string = "FLUSHDB"

	OP_H_KEYS string = "HKEYS"
	OP_H_SET   string = "HSET"
	OP_H_GET   string = "HGET"
	OP_H_DEL string = "HDEL"
	OP_L_PUSH  string = "LPUSH"
	OP_L_RANGE string = "LRANGE"
	OP_H_GETALL string = "HGETALL"
	OP_SCAN string = "SCAN"
	OP_KEY string = "KEYS"
	OP_H_LEN string = "HLEN"

	OP_Z_ADD      string = "ZADD"
	OP_Z_SCORE    string = "ZSCORE"
	OP_Z_RANGE    string = "ZRANGE"
	OP_Z_REVRANGE string = "ZREVRANGE"
	OP_Z_REM      string = "ZREM"

	OP_Z_REVRANK string = "ZREVRANK"
	OP_Z_REMRANGE_RANK string = "ZREMRANGEBYRANK"

	OP_PUBLISH string = "PUBLISH"

	OP_SELECE string = "SELECT"

	// set集合
	OP_S_ADD = "SADD" // 一个或多个成员元素加入到集合中
	OP_S_MEMBERS = "SMEMBERS"	// 返回集合中的所有的成员
	OP_S_REM = "SREM"	// 移除集合中的一个或多个成员元素
)

type RedisCacheClient struct {
	MaxIdle     int
	MaxActive   int
	Address     string
	IdleTimeout time.Duration //180 * time.Second
	pool        *redis.Pool
	DBIdx	int	// 默认连接的db索引
	Password string
	SentinelAddr []string	// 哨兵节点地址
	MasterName string	// 主节点别名
}

//redis.pool.maxActive=200  #最大连接数：能够同时建立的“最大链接个数”

//redis.pool.maxIdle=20     #最大空闲数：空闲链接数大于maxIdle时，将进行回收

//redis.pool.minIdle=5      #最小空闲数：低于minIdle时，将创建新的链接

//redis.pool.maxWait=3000    #最大等待时间：单位ms

//启动缓存客户端
func (this *RedisCacheClient) Start() {
	this.pool = &redis.Pool{
		MaxIdle:     this.MaxIdle,
		MaxActive:   this.MaxActive,
		IdleTimeout: this.IdleTimeout, //空闲释放时间
		Dial:
		func() (redis.Conn, error) {
			c,err := redis.Dial("tcp",this.Address)
			if err != nil {
				log.Debug("redis conn err 11: %v",err)
				return nil,err
			}
			if len(this.Password) > 0 {
				if _,err := c.Do("AUTH",this.Password);err != nil {
					c.Close()
					log.Debug("redis conn err 2: %v",err)
					return nil,err
				}
			}
			if _,err := c.Do("SELECT",this.DBIdx);err != nil {
				log.Debug("redis conn err 33: %v",err)
				return nil,err
			}
			return c,nil
		},
	}
}

//启动缓存客户端
func (this *RedisCacheClient) StartSentinel() {
	sntnl := &sentinel.Sentinel{
		Addrs:this.SentinelAddr,
		MasterName:this.MasterName,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 50000 * time.Millisecond
			c,err := redis.DialTimeout("tcp",addr,timeout,timeout,timeout)
			if err != nil {
				return nil,err
			}
			return c,nil
		},
	}
	this.pool = &redis.Pool{
		MaxIdle:     this.MaxIdle,
		MaxActive:   this.MaxActive,
		IdleTimeout: this.IdleTimeout, //空闲释放时间
		Dial:
		func() (redis.Conn, error) {
			masterAddr,err := sntnl.MasterAddr()
			if err != nil {
				return nil,err
			}
			c,err := redis.Dial("tcp",masterAddr)
			if err != nil {
				return nil,err
			}
			if len(this.Password) > 0 {
				if _,err := c.Do("AUTH",this.Password);err != nil {
					c.Close()
					return nil,err
				}
			}
			if _,err := c.Do("SELECT",this.DBIdx);err != nil {
				return nil,err
			}
			return c,nil
		},
		TestOnBorrow:CheckRedisRole,
	}
}

func CheckRedisRole(c redis.Conn, t time.Time) error {
	if !sentinel.TestRole(c, "master") {
		return fmt.Errorf("Role check failed")
	} else {
		return nil
	}
}

//关闭缓存客户端
func (this *RedisCacheClient) Close() {
	if this.pool != nil {
		this.pool.Close()
	}
}

// 设置
func (this *RedisCacheClient) GetConnByDB(dbIndex int) redis.Conn {
	conn,err := redis.Dial("tcp",this.Address)
	if err != nil {
		log.Debug("%v",err)
		return nil
	}
	if len(this.Password) > 0 {
		_,err = conn.Do("AUTH",this.Password)
		if err != nil {
			log.Debug("%v",err)
			return nil
		}
	}
	conn.Do(OP_SELECE,dbIndex)
	return conn
}

// 获取连接，指定db
func (this *RedisCacheClient) getConn() redis.Conn {
	conn := this.pool.Get()
	//_,err := conn.Do(OP_SELECE,this.DBIdx)
	//if err != nil {
	//	log.Debug("%v",err)
	//	return nil
	//}
	if conn == nil {
		log.Error("getConn nil ")
	}
	return conn
}

// 清除db数据
func (this *RedisCacheClient) FlushDB() bool {
	conn := this.getConn()
	defer conn.Close()
	conn.Do(OP_FLUSHDB)
	return true
}

//添加数据
func (this *RedisCacheClient) SetData(key string, value interface{}) bool {
	conn := this.getConn()
	defer conn.Close()
	_, err := conn.Do(OP_SET, key, value)
	if (err != nil) {
		log.Debug("%v", err)
		return false
	}
	return true
}

func (this *RedisCacheClient) IncrData(key string) bool {
	conn := this.getConn()
	defer conn.Close()
	_,err := conn.Do(OP_INCR,key)
	if (err != nil) {
		log.Debug("%v", err)
		return false
	}
	return true
}

func (this *RedisCacheClient) GetDataInt32(key string) int {
	conn := this.getConn()
	defer conn.Close()
	value, err := redis.Int(conn.Do(OP_GET, key))
	if err != nil {
		log.Debug("%v", err)
		return 0
	}
	return value
}

//获取数据
func (this *RedisCacheClient) GetData(key string) string {
	conn := this.getConn()
	defer conn.Close()
	value, err := redis.String(conn.Do(OP_GET, key))
	if err != nil {
		log.Debug("%v", err)
		return ""
	}
	return value
}

//添加数据
func (this *RedisCacheClient) DelData(key string) bool {
	conn := this.getConn()
	defer conn.Close()
	_, err := conn.Do(OP_DEL, key)
	if (err != nil) {
		log.Debug("%v", err)
		return false
	}
	return true
}

//清除所有数据
func (this *RedisCacheClient) FlashAll() {
	conn := this.getConn()
	defer conn.Close()
	conn.Do(OP_FLUSHALL)
}

//
func (this *RedisCacheClient) LPush(key string, values []string) {
	conn := this.getConn()
	defer conn.Close()
	for i := 0; i < len(values); i++ {
		conn.Do(OP_L_PUSH, key, values[i])
	}
}

//
func (this *RedisCacheClient) LRange(key string) []string {
	conn := this.getConn()
	defer conn.Close()
	values, err := redis.Strings(conn.Do(OP_L_RANGE, key, 0, -1))
	if err != nil {
		log.Debug("%v", err)
		//return
	}
	return values
}

//// 存map
//func (this *RedisCacheClient)SetMap(key string ,value map[string]string) bool{
//	conn := this.getConn()
//	defer conn.Close()
//	// 转换成json
//	v,_ := json.Marshal(value)
//	// 存redis
//	_,err := conn.Do("SETNX",key, v)
//	if err != nil {
//		log.Error("%v",err)
//		return false
//	}
//	return true
//}
//
//// 取map
//func (this *RedisCacheClient)GetMap(key string) map[string]string {
//	conn := this.getConn()
//	defer conn.Close()
//	var imap map[string]string
//	value,err := redis.Bytes(conn.Do("GET",key))
//	if err != nil {
//		log.Error("%v",err)
//		return nil
//	}
//	// json转map
//	errShal := json.Unmarshal(value,&imap)
//	if errShal != nil {
//		log.Error("%v",errShal)
//		return nil
//	}
//	return imap
//}

func (this *RedisCacheClient) HSet(key interface{}, field string, value interface{}) bool {
	conn := this.getConn()
	defer conn.Close()
	_, err := conn.Do(OP_H_SET, key, field, value)
	if err != nil {
		log.Error("HSet %v", err)
		return false
	}
	return true
}

func (this *RedisCacheClient) HGet(key interface{}, field string) string {
	conn := this.getConn()
	defer conn.Close()
	value, err := redis.String(conn.Do(OP_H_GET, key, field))
	if err != nil {
		//log.Error("HGet %v", err)
		return ""
	}
	return value
}

func (this *RedisCacheClient) HDel(key interface{},field string) bool {
	conn := this.getConn()
	defer conn.Close()
	_,err := conn.Do(OP_H_DEL,key,field)
	if err != nil {
		log.Error("HDel %v",err)
		return false
	}
	return true
}

func (this *RedisCacheClient) HKeys(key interface{}) []string {
	conn := this.getConn()
	defer conn.Close()
	reply,err := conn.Do(OP_H_KEYS,key)
	value,err := redis.Strings(reply,err)
	if err != nil {
		log.Error("HKeys %v",err)
		return []string{}
	}
	return value
}

func (this *RedisCacheClient) Keys(key interface{}) []string {
	conn := this.getConn()
	defer conn.Close()
	reply,err := conn.Do(OP_KEY,key)
	value,err := redis.Strings(reply,err)
	if err != nil {
		log.Error(" Keys %v",err)
		return []string{}
	}
	return value
}

func (this *RedisCacheClient) HGetBool(key interface{}, field string) bool {
	conn := this.getConn()
	defer conn.Close()
	value, err := redis.Bool(conn.Do(OP_H_GET, key, field))
	if err != nil {
		log.Error("%v", err)
		return false
	}
	return value
}

func (this *RedisCacheClient) HGetInt32(key interface{}, field string) int32 {
	conn := this.getConn()
	defer conn.Close()
	value, err := redis.Int(conn.Do(OP_H_GET, key, field))
	if err != nil {
		log.Error("%v", err)
		return 0
	}
	return int32(value)
}

//订阅数据变更
func (this *RedisCacheClient) Subscribe(callback func(channel, value string), channel ... interface{}) {
	//defer conn.Close()
	psc := redis.PubSubConn{Conn: this.pool.Get()}
	psc.Subscribe(channel...)
	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				value, _ := redis.String(v.Data, nil)
				callback(v.Channel, value)
			case error:
				log.Error("error: %v\n", v)
				return
			}
		}
	}()
}

func (this *RedisCacheClient) PSubscribe(callback func(pattern, channel, value string), channel ... interface{}) {
	//defer conn.Close()
	psc := &redis.PubSubConn{Conn: this.pool.Get()}
	psc.PSubscribe(channel...)
	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.PMessage:
				value, _ := redis.String(v.Data, nil)
				callback(v.Pattern, v.Channel, value)
			case error:
				log.Error("error: %v\n", v)
				return
			}
		}
	}()
}

func (this *RedisCacheClient) Send(cmd [][]string) {
	conn := this.getConn()
	defer conn.Close()

	for _,value := range cmd {
		if len(value) < 3 {
			continue
		}
		err := conn.Send(value[0],value[1],value[2],value[3])
		if err != nil {
			log.Error("Send: %v", err)
			continue
		}
	}
	conn.Flush()
	conn.Receive()
}

//发布数据
func (this *RedisCacheClient) Publish(channel, value interface{}) {
	conn := this.getConn()
	defer conn.Close()
	conn.Do(OP_PUBLISH, channel, value)
}
func (this *RedisCacheClient) GetConn() redis.Conn {
	return this.pool.Get()
}
//加锁
func (this *RedisCacheClient) LockSet(key interface{}, field string, value interface{})  {
	conn := this.getConn()
	defer conn.Close()

	//lock, ok, err := TryLock(conn, "xiaoru.cc", "token", int(10))
	//if err != nil {
	//	log.Fatal("Error while attempting lock")
	//}
	//if !ok {
	//	log.Fatal("Lock")
	//}
	//lock.AddTimeout(100)
	//
	//_, err = conn.Do(OP_H_SET, key, field, value)
	//if err != nil {
	//	log.Error("%v", err)
	//	return false
	//}
	//lock.Unlock()
	//
	//_, err := redis.String(conn.Do(OP_SET, key, field, "EX", int(10), "NX"))
	//if err == redis.ErrNil {
	//	// The lock was not successful, it already exists.
	//	return false, nil
	//}
	//if err != nil {
	//	return false, err
	//}
	//return true, nil
}

func (this *RedisCacheClient) Unlock(key interface{}) (err error) {
	conn := this.getConn()
	defer conn.Close()
	_, err = conn.Do(OP_DEL, key)
	return
}

type Lock struct {
	resource string
	token    string
	conn     redis.Conn
	timeout  int
}

func (lock *Lock) tryLock() (ok bool, err error) {
	_, err = redis.String(lock.conn.Do("SET", lock.key(), lock.token, "EX", int(lock.timeout), "NX"))
	if err == redis.ErrNil {
		// The lock was not successful, it already exists.
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (lock *Lock) Unlock() (err error) {
	_, err = lock.conn.Do(OP_DEL, lock.key())
	return
}

func (lock *Lock) key() string {
	return fmt.Sprintf("redislock:%s", lock.resource)
}

func (lock *Lock) AddTimeout(ex_time int64) (ok bool, err error) {
	ttl_time, err := redis.Int64(lock.conn.Do("TTL", lock.key()))
	fmt.Println(ttl_time)
	if err != nil {
		log.Fatal("redis get failed:", err)
	}
	if ttl_time > 0 {
		fmt.Println(11)
		_, err := redis.String(lock.conn.Do("SET", lock.key(), lock.token, "EX", int(ttl_time+ex_time)))
		if err == redis.ErrNil {
			return false, nil
		}
		if err != nil {
			return false, err
		}
	}
	return false, nil
}

func TryLock(conn redis.Conn, resource string, token string, DefaulTimeout int) (lock *Lock, ok bool, err error) {
	return TryLockWithTimeout(conn, resource, token, DefaulTimeout)
}

func TryLockWithTimeout(conn redis.Conn, resource string, token string, timeout int) (lock *Lock, ok bool, err error) {
	lock = &Lock{resource, token, conn, timeout}

	ok, err = lock.tryLock()

	if !ok || err != nil {
		lock = nil
	}

	return
}

func (this *RedisCacheClient) HLen(key interface{}) int {
	conn := this.getConn()
	defer conn.Close()
	value, err := redis.Int(conn.Do(OP_H_LEN,key))
	if err != nil {
		log.Error("HLen %v",err)
		return 0
	}
	return value
}


func (this *RedisCacheClient) HGetAll(key interface{}) []string {
	conn := this.getConn()
	defer conn.Close()
	reply,err := conn.Do(OP_H_GETALL,key)
	value,err := redis.Strings(reply,err)
	if err != nil {
		log.Error("%v",err)
		return []string{}
	}
	return value
}

func (this *RedisCacheClient) Scan(match interface{},count int) []string {
	var result []string
	cursor := "0"	// 游标从0 开始
	for i:=0;i<10;i++ {	// 只尝试迭代10次，没有就算了
		if len(result) >= count {
			break
		}
		i,values := this.scan(cursor,match,count)
		for _,v := range values {
			if len(result) >= count {
				break
			}
			result = append(result,v)
		}
		cursor = i	// 赋下新的游标
	}
	return result
}

// 模糊查询匹配的 key
func (this *RedisCacheClient) scan(cursor string,match interface{},count int) (string,[]string) {
	conn := this.getConn()
	defer conn.Close()

	var result []string
	value,err := redis.Values(conn.Do(OP_SCAN, cursor,"MATCH", match,"COUNT",count))
	if err != nil {
		log.Error("%v",err)
		return cursor,result
	}
	for i,v := range value {
		if i == 0 {	// idx0 是当前游标
			cursor,err = redis.String(v,err)
			if err != nil {
				log.Debug("err  %v",err)
				continue
			}
		} else {	// idx1 是数值,需要获取的结果
			values,err := redis.Strings(v,err)
			if err != nil {
				log.Debug("err  %v",err)
				continue
			}
			result = append(result,values...)
		}
	}
	return cursor,result
}