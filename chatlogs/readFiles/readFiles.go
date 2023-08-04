package readFiles

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/Fabricio2210/saveData"
)

func ReadFiles(subject string) {
	// Read the directory contents
	files, err := ioutil.ReadDir( fmt.Sprintf("./%s", subject))
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the directory contents and list files and subdirectories
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Directory: %s\n", file.Name())
		} else {
			savedata.Savedata(file.Name(), subject)
		}
	}
}