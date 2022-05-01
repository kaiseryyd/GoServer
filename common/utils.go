package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func TimestampToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func GetUnix() int64 {
	fmt.Println(time.Now().Unix())
	return time.Now().Unix()
}

func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}
