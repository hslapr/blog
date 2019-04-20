package data

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

// dataDir is the root directory of blog data
var dataDir string

// Download downloads the blog data intp a temporary directory
func Download() {
	dataDir, err := ioutil.TempDir("../tmp/", "blog-data")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dataDir) // clean up

	_, err = git.PlainClone(dataDir, false, &git.CloneOptions{
		URL: "https://github.com/hslapr/blog-data.git",
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Update updates blog data
func Update() {

}
