/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/user/agenda/entity"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		/* 读用户信息 */
		userInfo, userReadingerr := ReadUserInfo("data/user.txt")
		if userReadingerr != nil {
			fmt.Println(userReadingerr)
			return
		}
		/* 获取控制台参数信息 */
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		/* 检查是否当前有用户登陆 */
		fl, err1 := checkLogin()
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		if fl {
			fmt.Println("A user has logged in!")
			return
		}

		/* 检查用户名和密码是否正确 */
		if username == "" || password == "" { //用户名和密码不能为空
			fmt.Println("Username and password cannot be empty!")
			return
		}
		flag := false
		for _, user := range userInfo { //用户名和密码必须匹配
			if user.Username == username && user.Password == password {
				userLogin(username)
				fmt.Println("Login successfully!")
				flag = true
				break
			} else if user.Username == username && user.Password != password {
				fmt.Println("Wrong password!")
				return
			}
		}
		if flag != true {
			fmt.Println("The username does not exist!")
		}
		return

	},
}

/* 将当前用户信息写入cur_user.txt中 */
func userLogin(username string) error {
	return ioutil.WriteFile("data/cur_user.txt", []byte(username), os.ModeAppend)
}

/* 检查当前是否有用户登陆 */
func checkLogin() (bool, error) {
	msg, err := ioutil.ReadFile("data/cur_user.txt")
	if err != nil {
		return false, err
	}
	msg_ := string(msg)
	if msg_ == "No user is currently logged in!" {
		return false, nil
	} else if msg_ != "" {
		return true, nil
	}
	return false, nil
}

var (
	username *string
	password *string
)

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.
	username = loginCmd.Flags().StringP("username", "u", "", "your username")
	password = loginCmd.Flags().StringP("password", "p", "", "your password")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
