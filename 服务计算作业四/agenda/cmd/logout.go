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

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")
		/* 检查当前是否处于登陆状态 */
		flag, err := checkLogin()
		if err != nil { //检查文件是否读取成功
			fmt.Println(err)
			return
		} else if !flag { //不是登陆状态返回错误信息
			fmt.Println("Logout failed, please login first!")
			return
		}
		/* 否则修改当前用户信息并输出成功提示 */
		userLogout()
		fmt.Println("Logout successfully!")
	},
}

/* 写当前用户信息文件 */
func userLogout() error {
	return ioutil.WriteFile("data/cur_user.txt", []byte("No user is currently logged in!"), os.ModeAppend)
}

func init() {
	rootCmd.AddCommand(logoutCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
