package kioskPack

import (
	"log"
)

type Menu struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Category string `json:"category"`
}

type MenusData struct {
	Menus []Menu
}

func CheckErr(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
