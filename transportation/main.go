package main

import (
	"flag"
)

var (
	CfgNatsUri      string
	CfgMongoDsn     string
	CfgMysqlDsn     string
	CfgRedisAddress string
)

func init() {
	registerFlags()
}

func main() {
	flag.Parse()
}

func registerFlags() {
	flag.StringVar(&CfgNatsUri, "nats", "nats://nats.any:4222", "NATS URI")
	flag.StringVar(&CfgMongoDsn, "mongo", "mongodb://mongo.any/", "Mongo DB DSN")
	flag.StringVar(&CfgMysqlDsn, "mysql", "race:phi0lambda@tcp(mysql.race)/?parseTime=true", "Mysql DB DSN")
	flag.StringVar(&CfgRedisAddress, "redis", "redis.race:6379", "Redis Address")
}
