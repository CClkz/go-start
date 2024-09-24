package main

import "fmt"

func insertData() {
	// r, err := Db.Exec("insert into person(user_id, username, sex, email)values(?, ?, ?, ?)", "1", "stu001", "man", "stu01@qq.com")
	r, err := Db.Exec("insert into person( username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func getData() {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)

}

func updateData(username string) {
	if len(username) == 0 {
		username = "name00"
	}
	res, err := Db.Exec("update person set username=? where user_id=?", username, 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)
}

func deleteData() {
	res, err := Db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}

	fmt.Println("delete succ: ", row)
}
