package public

import (
	"crypto/sha256"
	"fmt"
	"github.com/gorhill/cronexpr"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"os"
	"strings"
	"time"
)

func GenSaltPassword(salt, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	str2 := fmt.Sprintf("%x", s2.Sum(nil))
	return str2
}

func IntToBool(a int) bool {
	if a == 0 {
		return false
	}
	return true
}

func GetFileSize(fileanme string) int {
	fileInfo, err := os.Stat(fileanme)
	if err != nil {
		return 0
	}
	tmp := int(fileInfo.Size()) / 1024
	return tmp
}

// CreateDir 创建文件夹
func CreateDir(path string) {
	_exist, _err := HasDir(path)
	if _err != nil {
		log.Logger.Errorf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		return
	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Logger.Errorf("创建目录异常 -> %v\n", err)
		} else {
			log.Logger.Infof("创建文件夹%s成功!", path)
		}
	}
}

func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

func Cronexpr(exprstr string) (string, error) {
	expr, err := cronexpr.Parse(exprstr) // 如果表达式解析错误将返回一个错误
	if err != nil {
		log.Logger.Error(err)
		return "unknown", err
	}
	nextTime := expr.Next(time.Now())
	return nextTime.Format("2006年01月02日15:04:01"), nil
}

func StatusConversion(a int) string {
	switch a {
	case 0:
		return "失败"
	case 1:
		return "成功"
	case 2:
		return "未启用"
	}
	return "unknown"
}

func StringToTime(receiveTime string) time.Time {
	timeLayOut := "2006年01月02日15:04"
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(timeLayOut, receiveTime, loc)
	if err != nil {
		log.Logger.Error("转换时间失败", err)
		return time.Time{}
	}
	return t
}

// GetBeforeDates 生成日期
func GetBeforeDates(day int) []string {
	var out []string
	for i := day; i >= 0; i-- {
		t := time.Now().AddDate(0, 0, -i)
		out = append(out, t.Format("2006-01-02"))
	}
	return out
}

// DateTran 将带年的日期转为月份
func DateTran(date string) string {
	list := strings.Split(date, "-")
	return list[len(list)-2] + "-" + list[len(list)-1]
}
