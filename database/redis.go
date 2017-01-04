package database

import (
	"fmt"
	"gopkg.in/redis.v4"
	"log"
	"time"
)

var Redisdb *redis.Client

func InitRedisDb() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	Redisdb = client
}

// get active seller setbit from redis
func GetActiveSellerByte(keyActiveSeller string) string {
	log.SetFlags(-1)
	log.SetPrefix("redis query = ")
	keyActiveSeller = "active_seller_daily:" + keyActiveSeller
	log.Println(keyActiveSeller)

	result, _ := Redisdb.Get(keyActiveSeller).Result()
	//	for _, c := range result {
	//		binString = fmt.Sprintf("%s%b", binString, c)
	//	}
	//

	return result
}

//checking the input id in redis exist or not
func IsIdExist(someId int64) bool {
	//format time now to yyy-mm-dd
	t := time.Now().Local()
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_daily:" + formatTime

	result, err := Redisdb.GetBit(keyActiveSeller, someId).Result()

	if err != nil {
		log.Println("Error IsIdExist = ", err.Error())
		return false
	}

	if result == 1 {
		return true
	} else {
		return false
	}
}

func InsertActiveSellerDaily(userId int64) {
	//format time now to yyy-mm-dd
	t := time.Now().Local()
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_daily:" + formatTime

	_, err := Redisdb.SetBit(keyActiveSeller, userId, 1).Result()
	if err != nil {
		log.Println("Error Insert = ", err.Error())
	}

	//use expire time 5 seconds
	seconds := 3600
	Redisdb.Expire(keyActiveSeller, time.Duration(seconds)*time.Second)

}

func InsertActiveSellerWeekly(userId int64) {
	t := time.Now().Local()
	//format time now to yyy-mm-dd
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_weekly:" + formatTime

	Redisdb.SetBit(keyActiveSeller, userId, 1)

}

func InsertActiveSellerMonthly(userId int64) {
	t := time.Now().Local()
	//format time now to yyy-mm-dd
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_monthly:" + formatTime

	Redisdb.SetBit(keyActiveSeller, userId, 1)

}