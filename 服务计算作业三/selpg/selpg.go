package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

type selpg_args struct {
	start_page  int    //开始页码
	end_page    int    //结束页码
	in_filename string //输入的文件名
	page_len    int    //每一页的行数
	page_type   bool   //分页的类型
	print_dest  string //打印的位置
}

func parseArgs(args *selpg_args) {
	pflag.IntVarP(&args.start_page, "start_page", "s", 0, "开始页")
	pflag.IntVarP(&args.end_page, "end_page", "e", 0, "结束页")
	pflag.IntVarP(&args.page_len, "page_len", "l", 10, "每页行数")
	pflag.BoolVarP(&args.page_type, "page_type", "f", false, "是否用换页符换页")
	pflag.StringVarP(&args.print_dest, "print_dest", "d", "", "打印位置")
	pflag.Parse()
}

var command_line string

func checkArgs(psa *selpg_args) {
	/* 检查参数的长度 */
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", command_line)
		os.Exit(1)
	}
	/* 检查第一个参数 -s，以及start_page */
	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "%s: 1st arg error\n", command_line)
		os.Exit(2)
	}
	if psa.start_page < 1 {
		fmt.Fprintf(os.Stderr, "%s: invalid start_page %s\n", command_line, psa.start_page)
		os.Exit(3)
	}

	/* 检查end_page */
	if psa.end_page < 1 || psa.end_page < psa.start_page {
		fmt.Fprintf(os.Stderr, "%s: invalid end_page %s\n", command_line, psa.end_page)
		os.Exit(5)
	}

	/* 检查传入文件的行数 */
	if psa.page_len != 10 {
		if psa.page_len < 1 {
			fmt.Fprintf(os.Stderr, "%s: invalid page_len %s\n", command_line, psa.page_len)
			os.Exit(6)
		}
	}

	if pflag.NArg() > 0 {
		psa.in_filename = pflag.Arg(0)
		/* 检查文件是否存在 */
		file, err := os.Open(psa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n", command_line, psa.in_filename)
			os.Exit(7)
		}
		/* 检查文件是否可读 */
		file, err = os.OpenFile(psa.in_filename, os.O_RDONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				fmt.Fprintf(os.Stderr, "%s: input file \"%s\" exists but cannot be read\n", command_line, psa.in_filename)
				os.Exit(8)
			}
		}
		file.Close()
	}
}

func processInput(psa *selpg_args) {
	fin := os.Stdin
	fout := os.Stdout
	var (
		page_ctr int
		line_ctr int
		err      error
		err1     error
		err2     error
		line     string
		cmd      *exec.Cmd
		stdin    io.WriteCloser
	)

	if psa.in_filename != "" {
		fin, err1 = os.Open(psa.in_filename)
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n", command_line, psa.in_filename)
			os.Exit(11)
		}
	}

	if psa.print_dest != "" {
		cmd = exec.Command("cat", "-n")
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		stdin = nil
	}

	rd := bufio.NewReader(fin)
	if psa.page_type == false {
		line_ctr = 0
		page_ctr = 1
		for true {
			line, err2 = rd.ReadString('\n')
			if err2 != nil {
				break
			}

			line_ctr++
			if line_ctr > psa.page_len {
				page_ctr++
				line_ctr = 1
			}

			if page_ctr >= psa.start_page && page_ctr <= psa.end_page {
				fmt.Fprintf(fout, "%s", line)
			}
		}
	} else {
		page_ctr = 1
		for true {
			c, err3 := rd.ReadByte()
			if err3 != nil {
				break
			}
			if c == '\f' {
				page_ctr++
			}
			if page_ctr >= psa.start_page && page_ctr <= psa.end_page {
				fmt.Fprintf(fout, "%c", c)
			}
		}
		fmt.Print("\n")
	}

	if page_ctr < psa.start_page {
		fmt.Fprintf(os.Stderr, "%s: start_page (%d) greater than total pages (%d), no output written\n", command_line, psa.start_page, page_ctr)
	} else if page_ctr < psa.end_page {
		fmt.Fprintf(os.Stderr, "%s: end_page (%d) greater than total pages (%d), less output than expected\n", command_line, psa.end_page, page_ctr)
	}

	if psa.print_dest != "" {
		stdin.Close()
		cmd.Stdout = fout
		cmd.Run()
	}

	fmt.Fprintf(os.Stderr, "\n---------------\nProcess end\n")
	fin.Close()
	fout.Close()
}

func main() {
	sa := selpg_args{0, 0, "", 10, false, ""} //创建一个结构体并初始化，其中分页方式为-l line number
	command_line = os.Args[0]                 //获得可执行文件的name
	parseArgs(&sa)                            //初始化结构体数据
	checkArgs(&sa)                            //检查参数的合法性
	processInput(&sa)                         //通道传递参数，对输入文件进行处理
}
