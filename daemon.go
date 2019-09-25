package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/takama/daemon"
)

func main() {

	service, err := daemon.New("golangDaemon", "golangDaemonTest")

	fileLog, err := os.OpenFile("golangDeamon_Log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(fileLog, os.Stdout)
	log.SetOutput(multiWriter)
	defer fileLog.Close()

	if err != nil {
		log.Println(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	status, err1 := service.Status()

	log.Println("Service...")
	ppID := os.Getppid()

	if err1 != nil {
		status, err1 = service.Install()

		log.Println(status)

		if err1 == nil {
			//success install...
			status, err1 = service.Start()

			if err1 != nil {
				log.Println("Error : ", err1)
				log.Println(status)
				os.Exit(-1)
			} else {
				log.Println(status)
				log.Println("Manually Starting...")
				log.Println("Please Check Service..[systemctl status 'servicename']")
				os.Exit(1)
			}
		}
	} else {

		log.Println(status)

		if strings.Contains(status, "is running") {
			log.Println("Already Running..")

			if ppID == 1 {
				go func() {
					backWork()
				}()
			} else {
				log.Println("Manually Execute Program is Aborted")
				os.Exit(1)
			}

		} else {

			log.Println(status)

			if strings.Contains(status, "is stopped") {
				status, err1 = service.Start()
				log.Println(status)
				log.Println("Restart..")
				os.Exit(1)
			} else {
				log.Println(status)
				log.Println("Error : Can't Running..")
				os.Exit(-1)
			}

		}
	}

	//wait.Wait()
	log.Println("channel loop select start..")
	for {
		select {
		case killSignal := <-interrupt:
			log.Println("Got Signal : ", killSignal)
			os.Exit(1)
		}
	}
	log.Println("channel loop select end..")
}

func backWork() {

	for {
		//do it something....
		log.Println("Goroutine....")
		time.Sleep(time.Second * 10)
	}
}
