package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	defer db.Close()

	res, err := db.Exec("CREATE TABLE IF NOT EXISTS `userinfo` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT,`username` VARCHAR(64) NULL,`departname` VARCHAR(64) NULL,`created` DATE NULL)")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err = stmt.Exec("astaxie", "研发部门", time.Now().Format("2006-01-02 15:04:05"))
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("LastInsertId: ", id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("RowsAffected: ", affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("RowsAffected: ", affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
