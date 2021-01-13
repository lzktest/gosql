package gosql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pg"
	"strconv"
	"strings"
)

type Gosql struct {
	err error
	db *sql.DB
	result interface{}
}

func (g Gosql)ExecSqlline(db *sql.DB,arg map[string]interface{}) error {
	var test []interface{}
	var qstrbt strings.Builder
	var valuestr string
	qstrbt.WriteString("insert into casbin_rule(")
	i :=0
	for k, v := range arg{
		if i != 0 {
			qstrbt.WriteString(",")
			valuestr += ","
		}
		qstrbt.WriteString(k)
		i++
		valuestr += "$"+strconv.Itoa(i)
		test = append(test,v)
	}
	qstrbt.WriteString(") values(")
	qstr := qstrbt.String() + valuestr+ ");"
	_,err :=db.Exec(qstr,test...)
	if err != nil {
		return err
	}
	return nil
}