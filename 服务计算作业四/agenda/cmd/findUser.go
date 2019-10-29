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

	"github.com/spf13/cobra"
	. "github.com/user/agenda/entity"
)

// findUserCmd represents the findUser command
var findUserCmd = &cobra.Command{
	Use:   "findUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("findUser called")
		/* 先检查登陆状态 */
		flag, err := checkLogin()
		if err != nil {
			fmt.Println(err)
			return
		} else if !flag {
			fmt.Println("Find failed, please login first!")
			return
		}
		/* 读用户信息 */
		userInfo, err_ := ReadUserInfo("data/user.txt")
		if err_ != nil {
			fmt.Println(err_)
			return
		}
		/* 输出所有用户信息 */
		fmt.Println("User information is as follows:")
		for _, user := range userInfo {
			fmt.Println("Username:", user.Username, " ", "Email:", user.Email, " ", "Phone:", user.Phone)
		}
	},
}

func init() {
	rootCmd.AddCommand(findUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
