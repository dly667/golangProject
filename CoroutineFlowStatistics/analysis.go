package CoroutineFlowStatistics

import (
	"bufio"
	"flag"
	"io"
	"os"
	"time"
	"github.com/sirupsen/logrus"
)

type cmdParams struct {
	logFilePath string
	routineNum int
}
type digData struct {
	time	string
	url		string
	refer	string
	ua 		string
}
type urlData struct {
	data 	digData
	uid 	string
}
type urlNode struct {
	
} 
type storageBlock struct {
	counterType string
	storageModel string
	unode 		urlNode
}

var log = logrus.New()
func init(){
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)

}

func main(){
	// 获取参数
	logFilePath := flag.String("logFilePath", "默认值","log file")
	routineNum := flag.Int("routineNum",5,"routineNum")
	l := flag.String("l", "/tmp/log","this program log ")
	flag.Parse()

	params := cmdParams{*logFilePath, *routineNum}
	// 打日志
	logFd, err := os.OpenFile(*l, os.O_CREATE|os.O_WRONLY, 0644 )
	if err == nil{
		log.Out = logFd
		defer logFd.Close()
	}
	log.Infoln("Exec start.")
	log.Infof("Params: %s", params.logFilePath)
	// 初始化一些channel，用于数据传递
	var logChannel = make(chan string, 3 * params.routineNum)
	var pvChannel = make(chan urlData, params.routineNum)
	var uvChannel = make(chan urlData, params.routineNum)
	var storageChannel = make(chan storageBlock, params.routineNum)
	// 日志消费者
	go readFileLineByLine(params, logChannel )

	// 创建一组日志处理
	for i:=0; i<params.routineNum; i++{
		go logConsumer(logChannel, pvChannel, uvChannel)
	}
	// 创建PV UV 统计器
	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel)
	//此处可统计器扩展

	// 创建存储器
	go dataStorage(storageChannel)

	time.Sleep(100 * time.Second)
}

func readFileLineByLine(params cmdParams, logChannel chan string) error{
	fd, err := os.OpenFile(params.logFilePath)
	if err != nil{
		log.Warningf("ReadFileLineByLine can't open %s",params.logFilePath)
		return err
	}
	defer fd.Close()

	count := 0
	bufferRead := bufio.NewReader(fd)
	for {
		line, err := bufferRead.ReadString('\n')
		logChannel <- line
		count++

		if count%(1000 * params.routineNum) == 0{
			log.Infof("ReadFileLineByLine %d", count)
		}
		if err != nil{
			if err == io.EOF{
				time.Sleep(3 * time.Second)
				log.Infof("ReadFileLineByLine wait, readline %d", count)
			}else {
				log.Warningf("ReadFileLineByLine read log err %d",count )
			}
		}
	}
}

func logConsumer(logChannel chan string, pvChannel, uvChannel chan urlData){

}

func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock ){

}

func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock)  {

}

func dataStorage(storageChannel chan storageBlock)  {

}
