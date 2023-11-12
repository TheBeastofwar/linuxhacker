package maintance

import (
	"fmt"
	"os"
)

func File_Change(existingfile string, changingfile string) {
	fileInfo, err := os.Stat(existingfile)
	if err != nil {
		fmt.Println("Unable to obtain file information", err)
		os.Exit(1)
	}
	atime := fileInfo.ModTime()
	mtime := fileInfo.ModTime()

	err = os.Chtimes(changingfile, atime, mtime)
	if err != nil {
		fmt.Println("Unable to change file information", err)
		os.Exit(1)
	}
	fmt.Println("success")
}
