package main

import (
	"fmt"
	"github.com/kardianos/osext"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(folderPath)

	switch folderPath {
	case "1":

	default:
		switch {
		case strings.HasSuffix(folderPath, "cirus\\") == true:
			cmd := exec.Command("mkdir", "test")
			cmd.Start()
			cmd1 := exec.Command("xcopy", "lol.exe", "test\\")
			cmd1.Start()
			fi := "run.bat"
			f, err := os.Create(fi)
			if err != nil {
				fmt.Println(err)
			}
			n, err := io.WriteString(f, "del lol.exe \n")
			n, err = io.WriteString(f, "cd test \n")
			n, err = io.WriteString(f, "start lol.exe \n")
			n, err = io.WriteString(f, "start /b \"\" cmd /c del \"%~f0\"&exit /b\n")
			if err != nil {
				fmt.Println(n, err)
			}
			f.Close()
			cmd3 := exec.Command("run.bat")
			cmd3.Start()
			return

		case strings.HasSuffix(folderPath, "test\\") == true:

			time.Sleep(time.Second * 5)
			return

		}

	}
	time.Sleep(time.Second * 10)
}
