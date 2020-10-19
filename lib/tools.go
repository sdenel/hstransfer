package lib

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

func StreamToBytes(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func BytesToStream(inBytes []byte) *bytes.Reader {
	return bytes.NewReader(inBytes)
}

func SaveAsFile(inBytes []byte, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, BytesToStream(inBytes))
	return err
}

func ArgsContainFlag(searchterm string) bool {
	s := strings.Join(os.Args[1:], ",") // TODO something is fucked with the input
	return strings.Contains(s, searchterm)
}

func ExecCommand(command string) {
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
