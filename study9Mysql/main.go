package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //init()
)

func main() {
	initDB()
	fmt.Println("===============")
	//insert()
	fmt.Println("===============")
	queryOne()
	//fmt.Println("===============")
	//querySome()
	//fmt.Println("===============")
	//update()
	//fmt.Println("===============")
	//delete()
	//fmt.Println("================")
	//transaction()
}

var db *sql.DB

func initDB() (err error) {
	//数据库信息
	//dsn := "root:Li20031202@tcp(127.0.0.1:3306)/data1"
	dsn := "root:123456@tcp(127.0.0.1:3306)/data1"
	//连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		//fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	//fmt.Println("连接成功")
	return nil
}

// 插入数据
func insert() {
	sqlStr := "insert into tb_user(username,address)values (?,?)"
	//row, err := db.Exec(sqlStr, "项羽", "广西")
	//预处理
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("insert err=", err)
	} else {
		//i, _ := row.RowsAffected()
		//ii, _ := row.LastInsertId()
		//fmt.Printf("i:%v,ii:%v\n", i, ii)
	}
	result, err := stmt.Exec("吕布", "宁夏")
	if err != nil {
		fmt.Println("insert err=", err)
	} else {
		i, _ := result.LastInsertId()
		fmt.Printf("i:%v\n", i)
	}
	defer stmt.Close()
}

type User struct {
	id       int
	username string
	address  string
}

// 查询单行
func queryOne() {
	//sqlStr := "select * from tb_user where id=?"
	//str := "xxx' or 1=1 #"
	//sqlStr := fmt.Sprintf("select * from tb_user where id ='%s'", str)
	sqlStr := "select * from tb_user where id ='xxx' or 1=1 #"
	var user User
	err := db.QueryRow(sqlStr).Scan(&user.id, &user.username, &user.address)
	if err != nil {
		fmt.Println("query err=", err)
	} else {
		fmt.Printf("%v\n", user)
	}
}

// 查询多行
func querySome() {
	sqlStr := "select * from tb_user"
	var user User
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("query err=", err)
	} else {
		for rows.Next() {
			rows.Scan(&user.id, &user.username, &user.address)
			fmt.Printf("%v\n", user)
		}
	}
}

// 更新数据
func update() {
	sqlStr := "update tb_user set username=?,address=? where id=?"
	row, err := db.Exec(sqlStr, "曹操", "广州", 115)
	if err != nil {
		fmt.Println("update err=", err)
	} else {
		i, _ := row.RowsAffected()
		//ii, _ := row.LastInsertId()
		fmt.Printf("i:%v\n", i)
	}
}

// 删除数据
func delete() {
	sqlStr := "delete from tb_user where id=?"
	row, err := db.Exec(sqlStr, 119)
	if err != nil {
		fmt.Println("delete err=", err)
	} else {
		i, _ := row.RowsAffected()
		//ii, _ := row.LastInsertId()
		fmt.Printf("i:%v\n", i)
	}
}

// 事务
func transaction() {
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	sqlStr1 := "update  table_account set money=money+2000 where name=?"
	sqlStr2 := "update  table_account set money=money-2000 where name=?"
	stmt1, err1 := tx.Prepare(sqlStr1)
	if err1 != nil {
		fmt.Println("err1=", err1)
		return
	}
	_, err1 = stmt1.Exec("刘备")
	if err1 != nil {
		fmt.Println("sqlStr1执行出错，已回滚")
		tx.Rollback() //回滚
		return
	}
	stmt2, err2 := tx.Prepare(sqlStr2)
	if err2 != nil {
		fmt.Println("err1=", err1)
		return
	}
	_, err2 = stmt2.Exec("关羽")
	if err2 != nil {
		fmt.Println("sqlStr2执行出错，已回滚")
		tx.Rollback() //回滚
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("事务提交错误，已回滚")
		tx.Rollback()
		return
	}
	fmt.Println("事务提交成功")
}
