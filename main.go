// 宣告程式屬於哪個 package
package main

// 引入套件
import (
    "fmt"
		"log"
		"database/sql"
	_ "github.com/go-sql-driver/mysql"	
)

// 程式執行入口
func main() {
    // 使用 fmt 套件印出字串 word，使用 := 簡化原本變數宣告 var word string = "Hello, World!"
    word := "Hello, World!"
    fmt.Println(word)
			// dsn := 

			// 使用 sql.Open 创建数据库对象。
			// 注意，这不会立即打开一个数据库连接。该函数只是验证DSN格式是否有效。
			db, err := sql.Open("mysql", "root:0000@tcp(172.25.0.2:3306)/go-project")
		
			// 检查错误
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()
		
			// 现在尝试与数据库建立连接
			err = db.Ping()
			if err != nil {
				log.Fatal(err)
			}
		
			fmt.Println("Successfully connected to the database!")
		}