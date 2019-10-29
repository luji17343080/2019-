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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/user/agenda/entity"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		/* Read data/user.txt */
		userInfo, err_ := ReadUserInfo("data/user.txt")
		if err_ != nil {
			fmt.Println(err_)
			return
		}
		/* 获取控制台参数信息 */
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		/* 用户信息检查 */
		pass, err := checkRegister(userInfo, username, password, email, phone)
		if pass {
			userInfo = append(userInfo, User{username, password, email, phone})
			fmt.Println("Register successfully!")
			WriteUserInfo("data/user.txt", userInfo)
		} else {
			fmt.Println("Register failed!", err)
			return
		}
	},
}

func checkRegister(userInfo []User, username string, password string, email string, phone string) (bool, error) {
	for _, user := range userInfo { //用户信息查重
		if user.Username == username {
			return false, errors.New("Sorry, the username already exists, please try again!")
		}
		if user.Email == email {
			return false, errors.New("Sorry, the email already exists, please try again!")
		}
		if user.Phone == phone {
			return false, errors.New("Sorry, the phone already exists, please try again!")
		}
	}

	if len(password) == 0 {
		return false, errors.New("The password cannot be empty!")
	} else if len(password) < 6 { //password长度不小于6
		return false, errors.New("Password at least 6 characters!")
	} else if len(email) == 0 {
		return false, errors.New("The email cannot be empty!")
	} else if len(phone) != 11 { //phone长度必须为11
		return false, errors.New("The phone must have eleven digits!")
	} else if len(phone) == 11 { //phone只能为0-9数字
		for i := 0; i < 11; i++ {
			if phone[i] < '0' || phone[i] > '9' {
				return false, errors.New("The phone can only contain digits!")
			}
		}
	}
	return true, nil
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.
	registerCmd.Flags().StringP("username", "u", "", "username that haven't be registered")
	registerCmd.Flags().StringP("password", "p", "", "your password, must be no shorter than 6 characters")
	registerCmd.Flags().StringP("email", "e", "", "your email address")
	registerCmd.Flags().StringP("phone", "n", "", "your phone number, must be 11 characters")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
