package lo
// 日志模块
// ...还没想好怎么写

import (
	"os"
	"log"
)

var Instance *Los = newLo()

type Los struct{
	lo *log.Logger
}
func newLo() *Los{
	// 初始化日志
	log_file, ferr := os.Create("nya.log")
	if ferr != nil {
		log.Fatalln("open file error !")
	}
	// defer log_file.Close()
	return &Los{
		lo: log.New(log_file,"[info]",log.LstdFlags),
	}
}

func (this *Los) Info(pre string, str interface{}) {
	this.lo.SetPrefix("["+pre+"]")
	this.lo.Println(str)
}



