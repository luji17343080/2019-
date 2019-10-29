package entity

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

/* 读取用户信息，即将用户信息的json格式转化为数组 */
func ReadUserInfo(filename string) ([]User, error) {
	var userInfo []User
	msg, err := ioutil.ReadFile(filename) //将文件中的信息转化为字符串
	if err != nil {
		return userInfo, err
	}
	/* 调用json包中的Unmarshal函数将json转化为数组 */
	json_ := string(msg)
	json.Unmarshal([]byte(json_), &userInfo)
	return userInfo, nil
}

func WriteUserInfo(filename string, userInfo []User) error {
	userInfo_json, err := json.Marshal(userInfo) //调用json包中的Marshal函数将数组转化为json
	if err == nil {
		return ioutil.WriteFile(filename, []byte(userInfo_json), os.ModeAppend)
	} else {
		return err
	}
}
