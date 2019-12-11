/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2017/3/29
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gradyZhang/summon-core/aliens/log"
	"errors"
)

type Rank struct {
	Member 		string
	Score 		int64
}


func Ranks(result interface{}, err error) ([]Rank, error) {
	values, err := redis.Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("redigo: IntMap expects even number of values result")
	}
	ranks := []Rank{}
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].([]byte)
		if !ok {
			return nil, errors.New("redigo: ScanMap key not a bulk string value")
		}
		value, err := redis.Int(values[i+1], nil)
		if err != nil {
			return nil, err
		}
		ranks = append(ranks, Rank{Score:int64(value), Member:string(key)})
	}
	return ranks, nil
}


//ZADD key score member [[score member] [score member] ...]
//将一个或多个 member 元素及其 score 值加入到有序集 key 当中。
func (this *RedisCacheClient)ZAdd(key string, score interface{}, member interface{}) bool{
	conn := this.getConn()
	defer conn.Close()
	_ , err := conn.Do(OP_Z_ADD, key, score, member)
	if  err != nil{
		//log.Error("%v",err)
		return false
	}
	return true
}

////ZADD key score member [[score member] [score member] ...]
////添加多条有序集合
//func (this *RedisCacheClient)ZAdds(key string, scoreMembers ...string) bool{
//	conn := this.pool.Get()
//	defer conn.Close()
//	_ , err := conn.Do(OP_Z_ADD, key, scoreMembers...)
//	if  err != nil{
//		log.Error("%v",err)
//		return false
//	}
//	return true
//}

//ZRANGE key start stop [WITHSCORES]   start:0  stop :-1 显示所有 下标从0开始
//返回有序集 key 中，指定区间内的成员。
//其中成员的位置按 score 值递增(从小到大)来排序。
func (this *RedisCacheClient)ZRange(key string, start int32, stop int32) []string {
	conn := this.getConn()
	defer conn.Close()
	result,err := redis.Strings(conn.Do(OP_Z_RANGE,key,start,stop))
	if  err != nil{
		//log.Error("%v",err)
	}
	return result
}

//ZREVRANGE key start stop [WITHSCORES]
//返回有序集 key 中，指定区间内的成员。
//其中成员的位置按 score 值递减(从大到小)来排列
func (this *RedisCacheClient)ZRevRange(key string, start int32, stop int32) []string {
	conn := this.getConn()
	defer conn.Close()
	result,err := redis.Strings(conn.Do(OP_Z_REVRANGE,key,start,stop))
	if  err != nil{
		//log.Error("%v",err)
	}
	return result
}

func (this *RedisCacheClient)ZRangeWithScore(key string, start int32, stop int32) []Rank {
	conn := this.getConn()
	defer conn.Close()
	result,	err := Ranks(conn.Do(OP_Z_RANGE,key,start,stop, PARAM_WITHSCORES))
	if  err != nil {
		//log.Error("%v",err)
	}
	return result
}

func (this *RedisCacheClient)ZRevRangeWithScore(key string, start int32, stop int32) []Rank {
	conn := this.getConn()
	defer conn.Close()
	result,err := Ranks(conn.Do(OP_Z_REVRANGE,key,start,stop, PARAM_WITHSCORES))
	if  err != nil {
		log.Error("%v",err)
	}
	return result
}

func (this *RedisCacheClient)ZRevRank(key string, member interface{}) int {
	conn := this.getConn()
	defer conn.Close()
	result,err := redis.Int(conn.Do(OP_Z_REVRANK,key, member))
	if  err != nil {
		//log.Error("%v",err)
		return -1
	}
	return result
}

// 移除指定排名的数据
func (this *RedisCacheClient)ZRemRangeByRank(key string,start int32,stop int32) {
	conn := this.getConn()
	defer conn.Close()

	redis.Int(conn.Do(OP_Z_REMRANGE_RANK,key, start,stop))
}

// 移除指定元素排行数据
func (this *RedisCacheClient)ZRem(key string,member ...interface{}) bool {
	conn := this.getConn()
	defer conn.Close()
	_,err := conn.Do(OP_Z_REM,key,member)
	if  err != nil {
		//log.Error("%v",err)
		return false
	}
	return true
}

// 获取指定member的数据
func (this *RedisCacheClient)ZScore(key string,  member interface{}) int64{
	conn := this.getConn()
	defer conn.Close()
	score , err := redis.Int(conn.Do(OP_Z_SCORE, key, member))
	if  err != nil{
		//log.Error("%v",err)
		return 0
	}
	return int64(score)
}

// 添加一个元素到 set集合
func (this *RedisCacheClient) SAdd(key string,member interface{}) bool {
	conn := this.getConn()
	defer conn.Close()
	_,err := conn.Do(OP_S_ADD,key,member)
	if  err != nil {
		//log.Error("%v",err)
		return false
	}
	return true
}

// 移除 set集合的一个元素
func (this *RedisCacheClient) SRem(key string,member interface{}) bool {
	conn := this.getConn()
	defer conn.Close()
	_,err := conn.Do(OP_S_REM,key,member)
	if  err != nil {
		//log.Error("%v",err)
		return false
	}
	return true
}

// 获取set集合所有元素
func (this *RedisCacheClient) SGetAllMember(key string) []string {
	conn := this.getConn()
	defer conn.Close()
	reslut,err := redis.Strings(conn.Do(OP_S_MEMBERS,key))
	if  err != nil {
		//log.Error("%v",err)
		return nil
	}
	return reslut
}