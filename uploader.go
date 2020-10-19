package main

import (
	"bufio"
	"hstransfer/lib"
	"os"
	"os/signal"
	"syscall"
)
import . "fmt"

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Println("\r- Ctrl+C pressed in Terminal. Cleaning...")
		lib.ExecCommand("./sh/clean.sh")
		os.Exit(0)
	}()
}
func performSync(generatedEncryptionKey string) {
	println("syncing...")
	directory := "/directory_to_upload/"
	directoryStat, err := os.Stat(directory)
	if err != nil {
		panic(err)
	}
	if !directoryStat.IsDir() {
		panic(Sprintf("%s is not a directory!", directory))
	}
	dirAsBytes, _ := lib.DirectoryToBytes(directory)
	dirAsBytesEncrypted := lib.Aes256Encrypt(generatedEncryptionKey, dirAsBytes)
	lib.SaveAsFile(dirAsBytesEncrypted, "/data_ready_to_upload/content.bin")
	lib.ExecCommand("./sh/rsync.sh")
	println("done.")
}

func uploader(runOnce bool) {
	SetupCloseHandler()
	sessionId := os.Getenv("HSTRANSFER_SESSION_ID")
	Printf("    sessionId=%s\n", sessionId)

	generatedEncryptionKey := lib.GenerateHex256keyAsStr()

	Println("    Building downloader...")
	lib.ExecCommand("./sh/build_downloader.sh")

	performSync(generatedEncryptionKey)
	Printf(
		"encryptionKey %s\nGo download then execute:\n%s\n%s\n",
		generatedEncryptionKey,
		os.Getenv("HSTRANSFER_HTTP_SESSION_PATH")+"/downloader-windows.exe",
		os.Getenv("HSTRANSFER_HTTP_SESSION_PATH")+"/downloader-linux")

	if runOnce {
		performSync(generatedEncryptionKey)
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			performSync(generatedEncryptionKey)
			reader.ReadString('\n')
		}
	}
}
