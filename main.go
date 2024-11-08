package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"sync"
)


var filePaths = map[string]string{
     
//    "here goes file video file name-1 " : "rtsp://localhost:8554/s1",
//    "here goes file video file name-2 " : "rtsp://localhost:8554/s2",

} 


func F_FFMPEGLIB() { 
    wg := sync.WaitGroup{}
    

	for file, url := range filePaths {
            fmt.Println(url)
            filePath := "/home/dbarun/mediamtx/" + file

            // Split the command and arguments correctly
            cmd := exec.Command("ffmpeg", "-re", "-stream_loop", "-1", "-i", filePath, "-c", "copy", "-f", "rtsp", "-rtsp_transport", "tcp", "-bufsize", "70000k", url)

            // Redirect output to the console
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            
        wg.Add(1)
        go func(){
            err := cmd.Run()
            if err != nil {
                slog.Error("Failed to run the command", "Details", err.Error())
                return
            }
            fmt.Println("Successfully started streaming:", file)    
        }()
	}

    wg.Wait()
    
}

func main(){      
    F_FFMPEGLIB()
}
