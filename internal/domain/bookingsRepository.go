package domain

import (
	"fmt"
	"mini-project/internal/database"
)



func CreateBooking(b Booking) (Booking, error) {
	db, err := database.Connect()
	if err != nil {
		return Booking{}, err
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into booking values (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return Booking{}, err
	}

	SqlResult, err := stmt.Exec(b.Id, b.ClientName, b.PhotographerName, b.Package, b.DateTime, b.Location, b.Status)
	if err != nil {
		return Booking{}, err
	}

	fmt.Println(SqlResult)
	return b, err
}



func DeleteBooking(id string) (string, error) {
	db, err := database.Connect()
	if err != nil {
		return "error to connect DB", err
	}
	defer db.Close()

	query := fmt.Sprintf("delete from booking where id = '%s' ", id)
	SqlResult , err := db.Exec(query)
    if err != nil {
        return "error while query", err
    }

	if rowsEffect, _ := SqlResult.RowsAffected(); rowsEffect == 0 {
		return "No Rows Deleted", err
	}
	

	return fmt.Sprintf("Booking with id %s successfully deleted", id), nil
	
}


func UpdateBooking(b Booking) (Booking, error) {
	db, err := database.Connect()
	if err != nil {
		return Booking{}, err
	}
	defer db.Close()

	_, err = db.Exec("update booking set ClientName = ?, PhotographerName = ?, Package = ?, DateTime = ?, Location =?, Status = ? where id = ?", b.ClientName, b.PhotographerName, b.Package, b.DateTime, b.Location, b.Status, b.Id)
	if err != nil {
		return Booking{}, err
	}

	return b, err

}

func ReadBookings(order string, orderBy string, page int, limit int) ([]Booking, error) {
	db, err := database.Connect()
	var args []interface{}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	query := "select * from booking "
	if orderBy != "" {
		query += "order by " + orderBy
	}
	
	if order != "" {
		query += " " + order
	}

	if limit > 0 {
		query += " limit ? offset ?" 
	}
	args = append(args, limit, limit*page)


	rows, err := db.Query(query,args...)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var result []Booking

	for rows.Next() {
		var each = Booking{}
		var err = rows.Scan(&each.Id, &each.ClientName, &each.PhotographerName, &each.Package, &each.DateTime, &each.Location, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}