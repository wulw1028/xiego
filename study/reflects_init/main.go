package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func configIni(filename string, v interface{}) (err error) {
	rType := reflect.TypeOf(v)
	rValue := reflect.ValueOf(v)
	// 判断是否为指针
	if rType.Kind() != reflect.Ptr {
		err = errors.New("v is not ptr")
		return
	}
	// 判断是否为struct
	if rType.Elem().Kind() != reflect.Struct {
		err = errors.New("v is not struct")
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	var structName string
	fileSlice := strings.Split(string(b), "\n")
	for idx, line := range fileSlice {
		strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		// 匹配到 [ 开始，作为一个节开始处理，直到遇到下一次 [
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			for i := 0; i < rType.Elem().NumField(); i++ {
				field := rType.Elem().Field(i)
				if field.Tag.Get("ini") == sectionName {
					structName = field.Name
				}
			}

		} else {
			dSlice := strings.Split(line, "=")
			if len(dSlice) == 1 && dSlice[0] == "" {
				continue
			}
			if len(dSlice) == 2 && (dSlice[0] == "" || dSlice[1] == "") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			sValue := rValue.Elem().FieldByName(structName)
			sType := sValue.Type()

			if sValue.Kind() != reflect.Struct {
				err = fmt.Errorf("%s is not struct", structName)
			}
			dKey := dSlice[0]
			dValue := dSlice[1]

			var fieldName string
			for i := 0; i < sType.NumField(); i++ {
				field := sType.Field(i)
				if field.Tag.Get("ini") == dKey {
					fieldName = field.Name
					break
				}
			}

			fieldObj := sValue.FieldByName(fieldName)
			switch fieldObj.Type().Kind() {
			case reflect.String:
				fieldObj.SetString(dValue)
			case reflect.Int:
				var strInt int64
				strInt, err = strconv.ParseInt(dValue, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d syntax error", idx+1)
					return
				}
				fieldObj.SetInt(strInt)
			}
		}
	}

	return
}

func main() {
	var cfg Config

	err := configIni("./config.ini", &cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", cfg)
}
