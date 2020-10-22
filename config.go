package config

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// Load config from config.ini
func Load(filePath string, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer value %s", reflect.TypeOf(v))
	}

	kvMap, err := loadFile(filePath)
	if err != nil {
		return err
	}

	keys := reflect.TypeOf(v).Elem()
	values := reflect.ValueOf(v).Elem()
	for i := 0; i < keys.NumField(); i++ {
		key := keys.Field(i).Name
		if len(keys.Field(i).Tag.Get("key")) != 0 {
			key = keys.Field(i).Tag.Get("key")
		}

		if len(kvMap[key]) != 0 {
			setVal(kvMap[key], values.Field(i))
		}
	}
	return nil
}

func loadFile(filePath string) (map[string]string, error) {
	config := make(map[string]string)

	lines, err := fileHandler(filePath)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		if len(line) > 0 && line[0:1] != "#" {
			val := strings.Split(line, "=")
			if len(val) == 2 {
				config[strings.TrimSpace(val[0])] = strings.TrimSpace(val[1])
			}
		}
	}

	return config, nil
}

func fileHandler(filePath string) ([]string, error) {
	var lines []string
	binary, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	str := string(binary)
	str = strings.Replace(str, "\r", "", -1)
	lines = strings.Split(str, "\n")
	return lines, nil
}

func setVal(varStr string, reflectVal reflect.Value) error {
	switch reflectVal.Kind() {
	case reflect.String:
		reflectVal.SetString(varStr)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num, err := strconv.ParseInt(varStr, 10, 64)
		if err != nil {
			return err
		}
		reflectVal.SetInt(num)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		num, err := strconv.ParseUint(varStr, 10, 64)
		if err != nil {
			return err
		}
		reflectVal.SetUint(num)

	case reflect.Float32, reflect.Float64:
		num, err := strconv.ParseFloat(varStr, 64)
		if err != nil {
			return err
		}
		reflectVal.SetFloat(num)

	case reflect.Complex64, reflect.Complex128:
		num, err := strconv.ParseFloat(varStr, 64)
		if err != nil {
			return err
		}
		reflectVal.SetComplex(complex(num, 0))

	case reflect.Bool:
		boolean, err := strconv.ParseBool(varStr)
		if err != nil {
			return err
		}
		reflectVal.SetBool(boolean)
	}

	return nil
}
