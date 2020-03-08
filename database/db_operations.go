package database

import (
	"activity-tracker/model"
	"database/sql"
	"fmt"
	"log"
	"time"
)



type DataBase struct {
	Db *sql.DB
}
func NewDataBaseConnection()(*DataBase, error){
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/activity?parseTime=true")
	if err != nil{
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Print(err)
	}
	return &DataBase{db},nil
}

func (d *DataBase) AddActivities(activityName string, activityType string, activityTime time.Time){
	//err := d.db.Ping()
	stmtIns, err := d.Db.Prepare("INSERT INTO activity_tracker VALUES(?, ?, ?, ? )")
	if err != nil {
		panic(err)
	}

	// Close the statement when we leave main() / the program terminates
	defer stmtIns.Close()

	// our id field auto increments so we don't need to pass actual value for it.
	_, err = stmtIns.Exec(nil,activityName,activityType,activityTime)
	if err != nil {
		panic(err)
	}
}
func (d *DataBase) GetActivity(ID string) (activity model.Activity){
	//err := d.db.Ping()
	sqlStatement := `SELECT * FROM activity_tracker WHERE ID=?;`
	row := d.Db.QueryRow(sqlStatement, ID)
	var id int
	var name string
	var activityType string
	var created_at string
	switch err := row.Scan(&id, &name, &activityType, &created_at); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		activity = model.FetchActivity(id,name,activityType,created_at)
		return activity
	default:
		panic(err)
	}
	return activity
}






