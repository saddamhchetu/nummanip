package dtbaseConn

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDB() (*sql.DB, error) {
	userName := "ddlocal"
	password := ""
	database := "4thand1_local_july"

	db, err := sql.Open("mysql", userName + ":" + password + "@/"+ database)
	/*defer func() {
		if r:= recover(); r != nil {
			return err
		}
	}()*/

	if err != nil {
		return db, err
	}

	return db, nil
}	
