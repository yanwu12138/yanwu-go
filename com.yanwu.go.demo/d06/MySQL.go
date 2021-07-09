package main

/**
 * @Author Baofeng Xu
 * @Date 2021/5/14 11:00
 *
 * Description: go语言使用MySQL数据库
 * ******************************************************** SQL *************************************************************
 * #创建user表
 * DROP TABLE IF EXISTS `order`;
 * DROP TABLE IF EXISTS `user`;
 * CREATE TABLE IF NOT EXISTS `user` (`uid` SERIAL PRIMARY KEY, `name` VARCHAR(20) NOT NULL, `password` VARCHAR(20) NOT NULL) ENGINE=`innodb`, CHARACTER SET=utf8;
 * #创建order表
 * CREATE TABLE IF NOT EXISTS `order`(`oid` SERIAL PRIMARY KEY, `uid`  BIGINT(20) UNSIGNED NOT NULL, `date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (`uid`) REFERENCES `user`(`uid`))ENGINE=innodb,CHARACTER SET=utf8;
 * #插入测试数据
 * INSERT INTO  `user`(`name`,`password`) VALUES('nick', 'nick'),('jacky', 'jacky');
 * NSERT INTO `order`(`uid`) VALUES(1),(2);
 **/

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 引入数据库驱动注册及初始化
)

func main() {
	// ----- 插入
	testInsert()
	// ----- 查询
	testSelect()
	// ----- 更改
	testUpdate()
	// ----- 删除
	testDelete()
	// ----- 事务
	testTransaction()
	// ----- 表达式
	testPreparedStatement()
}

func testInsert() {
	fmt.Println("-------------- insert --------------")
	/**********************************************************
	 * root					* 数据库用户名
	 * yanwu12138			* 数据库密码
	 * 192.168.56.150:3306	* 数据库地址与端口
	 * go_db_test			* 数据库库名
	 ***********************************************************/
	db, err := sql.Open("mysql", "root:yanwu12138@tcp(192.168.56.150:3306)/go_db_test?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	// ----- 释放数据库连接
	defer db.Close()

	// 插入一条新数据
	result, err := db.Exec("INSERT INTO `user`(`name`,`password`) VALUES('tom', 'tom')")
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return
	}
	fmt.Println("insert new record", id)
}

func testSelect() {
	fmt.Println("-------------- select --------------")
	/**********************************************************
	 * root					* 数据库用户名
	 * yanwu12138			* 数据库密码
	 * 192.168.56.150:3306	* 数据库地址与端口
	 * go_db_test			* 数据库库名
	 ***********************************************************/
	db, err := sql.Open("mysql", "root:yanwu12138@tcp(192.168.56.150:3306)/go_db_test?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	// ----- 释放数据库连接
	defer db.Close()

	// ===== 获取USERS表中的前十行记录
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var uid int
		var name, password string
		rows.Scan(&uid, &name, &password)
		fmt.Println("uid:", uid, "name:", name, "password:", password)
	}
}

func testUpdate() {
	fmt.Println("-------------- update --------------")
	/**********************************************************
	 * root					* 数据库用户名
	 * yanwu12138			* 数据库密码
	 * 192.168.56.150:3306	* 数据库地址与端口
	 * go_db_test			* 数据库库名
	 ***********************************************************/
	db, err := sql.Open("mysql", "root:yanwu12138@tcp(192.168.56.150:3306)/go_db_test?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	// ----- 释放数据库连接
	defer db.Close()

	// 更新一条数据
	result, err := db.Exec("UPDATE `user` SET `password`=? WHERE `name`=?", "tom_new_password", "tom")
	if err != nil {
		fmt.Println("update data failed:", err.Error())
		return
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return
	}
	fmt.Println("update recors number", num)
	testSelect()
}

func testDelete() {
	fmt.Println("-------------- delete --------------")
	/**********************************************************
	 * root					* 数据库用户名
	 * yanwu12138			* 数据库密码
	 * 192.168.56.150:3306	* 数据库地址与端口
	 * go_db_test			* 数据库库名
	 ***********************************************************/
	db, err := sql.Open("mysql", "root:yanwu12138@tcp(192.168.56.150:3306)/go_db_test?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	// ----- 释放数据库连接
	defer db.Close()

	// 删除数据
	result, err := db.Exec("DELETE FROM `user` WHERE `name`=?", "tom")
	if err != nil {
		fmt.Println("delete data failed:", err.Error())
		return
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return
	}
	fmt.Println("delete record number", num)
	testSelect()
}

func testTransaction() {
	fmt.Println("-------------- transaction --------------")
	/**********************************************************
	 * root					* 数据库用户名
	 * yanwu12138			* 数据库密码
	 * 192.168.56.150:3306	* 数据库地址与端口
	 * go_db_test			* 数据库库名
	 ***********************************************************/
	db, err := sql.Open("mysql", "root:yanwu12138@tcp(192.168.56.150:3306)/go_db_test?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	// ----- 释放数据库连接
	defer db.Close()

	// ===== 开启事务
	tx, err := db.Begin()
	// ----- 删除操作1
	result, err := tx.Exec("DELETE FROM `order` WHERE uid=? ", 2)
	if err != nil {
		fmt.Println("delete data failed:", err.Error())
		return
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return
	}
	fmt.Println("delete record number", num)
	// ----- 删除操作2
	result, err = tx.Exec("DELETE FROM `user` WHERE uid=? ", 2)
	if err != nil {
		fmt.Println("delete data failed:", err.Error())
		return
	}
	num, err = result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return
	}
	fmt.Println("delete record number", num)
	// ===== 根据条件回滚或者提交
	// tx.Rollback()
	tx.Commit()
}

func testPreparedStatement() {
	fmt.Println("-------------- prepared statement --------------")
	/**********************************************************
	 * root					* 数据库用户名
	 * yanwu12138			* 数据库密码
	 * 192.168.56.150:3306	* 数据库地址与端口
	 * go_db_test			* 数据库库名
	 ***********************************************************/
	db, err := sql.Open("mysql", "root:yanwu12138@tcp(192.168.56.150:3306)/go_db_test?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	// ----- 释放数据库连接
	defer db.Close()

	// 预备表达式
	stmt, err := db.Prepare("DELETE FROM `order` WHERE `oid`=?")
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return
	}
	result, err := stmt.Exec(1)
	if err != nil {
		fmt.Println("delete data failed:", err.Error())
		return
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return
	}
	fmt.Println("delete record number", num)
}
