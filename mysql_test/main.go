package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBTask struct {
	Dbconnection   *sql.DB
	UserInfomation UserTableStruct
}

type UserTableStruct struct {
	Id   sql.NullInt64
	Name sql.NullString
	Age  sql.NullInt64
	Sex  sql.NullString
	Addr sql.NullString
	Tel  sql.NullString
}

func main() {
	dbTask := &DBTask{}
	var err error
	dbTask.Dbconnection, err = sql.Open("mysql", "root:123456@tcp(192.168.2.225:3306)/mysql?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	defer dbTask.Dbconnection.Close()
	dbTask.CreateDatabase()
	dbTask.InsertData( 1024, "Alice", 23, "woman", "shenzhen", "12312312312")
	dbTask.QueryData( 20, 30)
	dbTask.UpdateData( 40, "man", 1024)
	dbTask.QueryData(20,30)
}

func (d *DBTask) CreateDatabase() {
	_, err := d.Dbconnection.Exec("CREATE DATABASE IF NOT EXISTS my_db;")
	if err != nil {
		fmt.Println("failed to create databases",err.Error())
		return
	}
	_, err = d.Dbconnection.Exec("USE my_db;")
	if err != nil {
		fmt.Println("select database failed")
		return
	}
	_, err = d.Dbconnection.Exec("CREATE TABLE IF NOT EXISTS tb_user(id int(10) primary key,name varchar(20),age int(10),sex varchar(5),addr varchar(64),tel varchar(11));")
	if err != nil {
		fmt.Println("create table  failed:", err.Error())
		return
	}
}

func (d *DBTask) InsertData( id int, name string, age int, sex string, addr string, tel string) {
	stmt, _ := d.Dbconnection.Prepare(`INSERT INTO tb_user(id,name,age,sex,addr,tel) values(?,?,?,?,?,?)`)
	defer stmt.Close()
	ret, err := stmt.Exec(id, name, age, sex, addr, tel)
	if err != nil {
		fmt.Println("insert data error", err.Error())
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId", LastInsertId)
	}
	if RowAffect, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowAffect", RowAffect)
	}
}

func (d *DBTask) QueryData(lower int, higher int) {
	stmt, _ := d.Dbconnection.Prepare(`SELECT * FROM tb_user where age>= ? AND age < ?`)
	defer stmt.Close()
	d.UserInfomation = UserTableStruct{}
	rows, err := stmt.Query(lower, higher)
	defer rows.Close()
	if err != nil {
		fmt.Println("insert data error", err.Error())
		return
	}
	for rows.Next() {
		rows.Scan(&d.UserInfomation.Id, &d.UserInfomation.Name, &d.UserInfomation.Age,&d.UserInfomation.Sex, &d.UserInfomation.Addr, &d.UserInfomation.Tel)
		if !d.UserInfomation.Addr.Valid {
			d.UserInfomation.Name.String = ""
		}
		if !d.UserInfomation.Age.Valid {
			d.UserInfomation.Age.Int64 = 0
		}
		fmt.Println("get data ,id:", d.UserInfomation.Id, "name:", d.UserInfomation.Name.String, "age:", "sex:",d.UserInfomation.Sex,d.UserInfomation.Age, "addr:", d.UserInfomation.Addr.String, "tel:", d.UserInfomation.Tel.String)
	}
}

func (d *DBTask) UpdateData( age int, sex string, id int) {
	stmt, err := d.Dbconnection.Prepare(`UPDATE tb_user SET age=?,sex=? WHERE id=?`)
	defer stmt.Close()
	result, err := stmt.Exec( age, sex, id)
	if err != nil {
		fmt.Println("update error", err.Error())
		return
	}
	num, err := result.RowsAffected()
	fmt.Println(num)
}

