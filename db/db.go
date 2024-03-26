package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Indirect Use
)

// func main() {
// 	OpenDBConnection()
// }

type Album struct {
	glusr_usr_id      int64
	fk_gl_country_iso string
	fk_gl_city_id     int64
	fk_gl_state_id    int64
}

func OpenDBConnection() error {
	var alb Album
	var albums []Album

	d := fmt.Sprintf("host=13.235.198.24 port=5439 user=rd_anurag_93402 " + "password=2B5b9XA8wt dbname=biredshiftdevelopment connect_timeout=2 sslmode=disable")
	db, err := sql.Open("postgres", d)
	if err != nil {
		fmt.Println(err)
	} else {
		rows, _ := db.Query("select glusr_usr_id, fk_gl_country_iso, fk_gl_city_id, fk_gl_state_id from im_dwh.dim_glusr_usr limit 100")
		// columns, _ := rows.Columns()
		// fmt.Println(strings.Join(columns, "\n"))
		for rows.Next() {
			rows.Scan(&alb.glusr_usr_id, &alb.fk_gl_country_iso, &alb.fk_gl_city_id, &alb.fk_gl_state_id)
			albums = append(albums, alb)
		}
		for _, val := range albums {
			fmt.Println(val)
		}
	}
	defer db.Close()
	return err
}
