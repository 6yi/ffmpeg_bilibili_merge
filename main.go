package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func Scanf(a *string) {
	//打开系统的输入流
	reader := bufio.NewReader(os.Stdin)
	//读取一整行二进制数据
	data, _, _ := reader.ReadLine()
	//转化为字符串
	*a = string(data)
}

func main() {
	var filename string
	var number string
	fmt.Println("输入需要批量转化的文件夹: ")
	Scanf(&filename)
	fmt.Println("输入文件内部编号(就是缓存视频下分段文件里面还有一个文件夹也是用数字表示的,每个缓存视频的这个编号是一样的)")
	fmt.Scanf("%s",&number)
	f, err := ioutil.ReadDir(filename)
	if err != nil {
		fmt.Println("文件夹不存在")
		os.Exit(0)
	}
	t4:=strconv.Itoa(time.Now().Hour())
	t5:=strconv.Itoa(time.Now().Minute())
	t6:=strconv.Itoa(time.Now().Second())
	newpath :=t4+""+t5+""+t6
	err2 :=os.Mkdir("./"+newpath, os.ModePerm)
	if err2 != nil {
		fmt.Printf("mkdir failed![%v]\n", err2)
	} else {

	}
	for i:=1;i<=len(f);i++{
		pathName:="H:/89238680/"+strconv.Itoa(i)+"/"+number
		cmd :=exec.Command(
			"ffmpeg","-i",
			pathName+"/video.m4s",
			"-i",
			pathName+"/audio.m4s",
			"-codec",
			"copy",
			newpath+"/"+strconv.Itoa(i)+".mp4")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("出错咯")
		}
		fmt.Println(string(out))
	}

}