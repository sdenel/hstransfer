package main

import (
	"bufio"
	. "fmt"
	"hstransfer/lib"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

var httpSessionPath = ""

func downloaderInner(encryptionKey string) {
	os.Mkdir("hstransfer", 0755)
	// Empty the dir: ugly way to ensure a --delete synchronisation
	dir, _ := ioutil.ReadDir("hstransfer")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"tmp", d.Name()}...))
	}

	url := httpSessionPath + "/content.bin"
	Printf("Downloading %s...\n", url)
	bytes := DownloadBytes(url)
	lib.BytesToDirectory(lib.Aes256Decrypt(encryptionKey, bytes), "hstransfer")
	Printf("    Done. Press enter to do again.\n")
}

func downloader(runOnce bool) {
	encryptionKey := os.Getenv("HSTRANSFER_ENCRYPTION_KEY")
	if len(encryptionKey) == 0 {
		println("Please provide encryptionKey:")
		Scanln(&encryptionKey)
	}

	Printf(
		"    httpSessionPath=%s encryptionKey=%s\n",
		httpSessionPath,
		encryptionKey,
	)

	if runOnce {
		downloaderInner(encryptionKey)
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			downloaderInner(encryptionKey)
			reader.ReadString('\n')
		}
	}

}

func DownloadBytes(url string) []byte {

	// Get data
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// Regularly update the user agent to improve legitimacy: https://www.whatismybrowser.com/guides/the-latest-user-agent/chrome
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	return lib.StreamToBytes(resp.Body)
}
