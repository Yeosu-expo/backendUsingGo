package kioskPack

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func PostAndStoreJson(w http.ResponseWriter, r *http.Request) {
	menu := new(Menu)
	read, err := io.ReadAll(r.Body)
	CheckErr(err)
	err = json.Unmarshal(read, &menu)
	CheckErr(err)

	db, err := sql.Open("mysql", "root:9250@tcp(127.0.0.1:3306)/kiosk")
	defer db.Close()
	CheckErr(err)

	insertQuery := "INSERT INTO menu (name, category, price) VALUES(?,?,?)"
	price, _ := strconv.Atoi(menu.Price)
	_, err = db.Exec(insertQuery, menu.Name, menu.Category, price)
	CheckErr(err)
}
