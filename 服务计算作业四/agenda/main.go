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
package main

import (
	//"fmt"
	// "io/ioutil"

	"github.com/user/agenda/cmd"
)

// /* 初始化用户数据信息 */
// func init_data(filename, null_data string) {
// 	data := []byte(null_data)
// 	if ioutil.WriteFile(filename, data, 0644) == nil {
// 		fmt.Println(filename, "initialization successful!")
// 	}
// }

func main() {
	// content := ""
	// userdata := "data/user.txt"
	// cur_userdata := "data/cur_user.txt"
	// init_data(userdata, content)
	// init_data(cur_userdata, content)
	cmd.Execute()
}
