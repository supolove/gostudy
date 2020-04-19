package web开发

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "log"
	"testing"
)

/*
orm

sql builder 使用mysql简短语言封装

大型的互联网公司核心线上业务都会在代码中把SQL放在显眼的位置提供给DBA评审
*/

func TestDatabase(t *testing.T) {
	// db 是一个 sql.DB 类型的对象
	// 该对象线程安全，且内部已包含了一个连接池
	// 连接池的选项可以在 sql.DB 的方法中设置，这里为了简单省略了
	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		Fatal(err)
	}
	defer db.Close()

	var (
		id   int
		name string
	)
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		Fatal(err)
	}

	defer rows.Close()

	// 必须要把 rows 里的内容读完，或者显式调用 Close() 方法，
	// 否则在 defer 的 rows.Close() 执行之前，连接永远不会释放
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			Fatal(err)
		}
		Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		Fatal(err)
	}
}
