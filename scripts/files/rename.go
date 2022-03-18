package files

import (
	"fmt"
	"io/ioutil"
)

func RenameByName() {
	items, _ := ioutil.ReadDir("../../demo-tokens/zunka")
	for _, item := range items {
		if item.IsDir() {
			subitems, _ := ioutil.ReadDir(item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {
					// handle file there
					fmt.Println(item.Name() + "/" + subitem.Name())
				}
			}
		} else {
			// handle file there
			fmt.Println(item.Name())
		}
	}
}
