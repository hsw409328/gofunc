package go_plugin

import (
	"github.com/go-redis/redis"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var (
	end         chan bool
	restartChan = make(chan bool)
	pid         int
	clientRedis *redis.Client
)

func init() {
	clientRedis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	go func() {
		for {
			_, err := clientRedis.Get("restart_symbol").Result()
			if err == nil {
				restartChan <- true
			}
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			select {
			case <-restartChan:
				restart()
				break
			}
		}
	}()
	<-end
}

func restart() {
	//获取当前启动的路径及应用
	filePath, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Println(err)
		return
	}
	var argsParams []string
	for i := 1; i < len(os.Args); i++ {
		// 使用=分割的参数
		tmpArgs := strings.Split(os.Args[i], "=")
		for _, tmpVal := range tmpArgs {
			argsParams = append(argsParams, tmpVal)
		}
	}

	cmd := exec.Command(filePath, argsParams...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Println("启动失败：", err)
		return
	}
	os.Exit(0)
}
