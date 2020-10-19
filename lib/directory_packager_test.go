package lib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestDirectoryPackager(t *testing.T) {
	// Given
	tmpDir, _ := ioutil.TempDir(os.TempDir(), "totoprogramname-")
	fmt.Println(tmpDir)
	defer os.Remove(tmpDir)

	inputDir := path.Join(tmpDir, "input")
	os.Mkdir(inputDir, 0755)
	ioutil.WriteFile(path.Join(inputDir, "hello.txt"), []byte("Hello world!"), 0644)

	outputDir := path.Join(tmpDir, "output")
	os.Mkdir(outputDir, 0755)

	// When
	bytes, err := DirectoryToBytes(inputDir)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	// Then
	BytesToDirectory(bytes, outputDir)
	_, err = os.Stat(path.Join(outputDir, "hello.txt"))
	assert.Nil(t, err)
	// TODO we might check the content too.
}
