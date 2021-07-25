//Q:我们在数据库操作的时候，比如dao层中当遇到一个sql.ErrNoRows的时候，是否应该Wrap这个error，抛给上层。为什么，应该怎么做请写出代码
//A:应该Wrap这个error，抛给上层。因为查询结果为空这个错误在这一层不能做降级处理，不能给打日志或给默认，所以应该让上层判断，进行错误处理
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func mySqlQurey() error {
	// TODO creat db
	var name string
	err := db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			//TODO Wrap errors
			return errors.Wrap(err, "there were no rows")
		} else {
			//TODO 降级处理
			log.Fatal(err)
		}
	}
	//TODO next
	return nil
}

func Call() errors {
	return errors.WithMessage(mySqlQurey(), "find failed")
}

func main() {
	err := Call()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}

	//TODO next things
}
