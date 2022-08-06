package cmd

import (
	"fmt"
	"nc_beanshell_scan/module/nc"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nc_beanshell_scan",
	Short: "nc_beanshell_scan对nc潜在的rce风险进行扫描",
	Long:  "从url或者文件读取需要扫描的用友nc链接,进行用友nc_beanshell扫描",
	Run: func(cmd *cobra.Command, args []string) {
		if urla != "" {
			s := nc.NewScan([]string{urla}, thread, output)
			s.StartScan()
		} else if localfile != "" {
			urls := removeRepeatedElement(nc.LocalFile(localfile))
			s := nc.NewScan(urls, thread, output)
			s.StartScan()
		} else {
			fmt.Println("nc_beanshell_scan,use -h see usage")
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

var (
	urla      string
	localfile string
	thread    int
	output    []string
)

func init() {
	rootCmd.Flags().StringVarP(&urla, "url", "u", "", "识别单条url")
	rootCmd.Flags().StringVarP(&localfile, "local", "l", "", "从本地文件读取资产，进行指纹识别，支持无协议，列如：192.168.1.1:9090 | http://192.168.1.1:9090")
	rootCmd.Flags().IntVarP(&thread, "thread", "t", 20, "指纹识别线程大小。")
}

func removeRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
