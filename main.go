package main

import (
	"github.com/fatih/color"
	"os/exec"
	"os"
	"os/signal"
	"fmt"
	"strings"
	"runtime"
)

func commandWrapper(inst func()string){
	color.Blue("Installing ffmpeg, enter password below")
	o := inst()
	fmt.Println(o)
}

func DistroExists(distro string, output []byte) bool{
	if strings.Contains(strings.ToLower(strings.Split(string(output),"\n")[0]),strings.ToLower(distro)){
		return true
	}
	return false
}

func InstallffmpegNix(){
	cmd := exec.Command("cat","/etc/os-release")
	output ,_ := cmd.Output()
	
	if DistroExists("solus",output){
		
		commandWrapper(func() string{
			o, _ := exec.Command("sudo","eopkg","install","ffmpeg").Output()
			return string(o)
		})
	}

	if DistroExists("ubuntu",output) || DistroExists("debian",output){
		commandWrapper(func() string{
			o, _ := exec.Command("sudo","apt","install","ffmpeg").Output()
			return string(o)
		})
	}

	if DistroExists("manjaro",output) || DistroExists("arch",output){
		commandWrapper(func() string{
			o, _ := exec.Command("sudo","pacman","-S","ffmpeg").Output()
				return string(o)
		})
	}
}


func checkForffmpeg(){
	cmd := exec.Command("ffmpeg","-h")
	_, err := cmd.Output()
	if err !=nil{
		if runtime.GOOS=="linux"{
		InstallffmpegNix()
	}}
}


func main(){
	var output string

	v := color.New(color.FgMagenta)

	v.Printf("Enter file Name: ")
	fmt.Scanf("%s",&output)


	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func(){
		<-c
		color.Cyan("\nRecording stopped")
		os.Exit(0)

	}()

	color.Green("Recording...")
	command := exec.Command("ffmpeg", "-f" ,"x11grab" ,"-i" ,":0.0" ,output)
	_, err := command.Output()

	if err !=nil{
		panic(err)
	}
	
	
}