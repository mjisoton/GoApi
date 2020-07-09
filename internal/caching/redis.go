package caching

//Dependencies
import "errors"

//Third-Party
import "github.com/mediocregopher/radix/v3"

//Global
var redis *radix.Pool

//Try and create a new pool of connections to Redis NoSQL
func Connect(sock string, min_conn int) error {
	var err error

	//Connect and disables the implicit pipelining to decrease the latency
	redis, err = radix.NewPool("unix", sock, min_conn, radix.PoolPipelineWindow(0, 0))
	if err != nil {
		return err
	}

	if availConn := redis.NumAvailConns(); availConn == 0 {
		return errors.New("The Redis database has 0 connection slots available to use.")
	}

	return nil
}

//Redis GET
func Get(key string) string {
	var r string
	redis.Do(radix.Cmd(&r, "GET", key))
	return r
}

//Redis SET
func Set(key string, val string) error {
	redis.Do(radix.Cmd(nil, "SET", key, val))
	return nil
}

//Redis SETEX
func SetEx(key string, val string, exp string) error {
	redis.Do(radix.Cmd(nil, "SETEX", key, exp, val))
	return nil
}

//Redis TTL
func TTL(key string) int {
	var r int
	redis.Do(radix.Cmd(&r, "TTL", key))
	return r
}

//Redis HSET
func Hset(key string, index string, val string) int {
	redis.Do(radix.Cmd(nil, "HSET", key, index, val))
	return 0
}


//Redis HGET
func Hget(key string, index string) string {
	var r string
	redis.Do(radix.Cmd(&r, "HGET", key, index))
	return r
}
