package utils

import (
	"fmt"
	"io/ioutil"
)

func GetViewPaths() (name []string) {
	dir, err := ioutil.ReadDir("views/home")
	if err != nil {
		fmt.Println(err)
		return name
	}

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			name = append(name, fi.Name())

		}
	}

	return name
}
