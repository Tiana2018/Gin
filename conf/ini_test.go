package conf

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"testing"
)

func TestApp( t *testing.T)  {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}
	 // 典型读取操作，默认分区使用空字符串表示
	 fmt.Println("App PageSize:",cfg.Section("app").Key("PAGE_SIZE").String())
	fmt.Println("RUN_MODE:",cfg.Section("").Key("RUN_MODE").String())

	// 我们可以做一些候选值限制的操作
	fmt.Println("DataBase Type:",cfg.Section("database").Key("TYPE").In("mysql",[]string{"mysql","pql"}))
	// 如果读取的值不在候选列表内，则会回退使用提供的默认值
	fmt.Println("DataBase Type:",cfg.Section("database").Key("TYPE").In("pql",[]string{"db","pql"}))

	// 试一试自动类型转换, 没有就自动添加了
	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))

	// 修改某个值然后进行保存
	cfg.Section("").Key("RUN_MODE").SetValue("")
	cfg.SaveTo("my.ini")

}
