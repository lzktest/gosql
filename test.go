package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
)
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	PType string `gorm:"size:40;uniqueIndex:unique_index"`
	V0    sql.NullString `gorm:"size:40;uniqueIndex:unique_index"`
	V1    sql.NullString `gorm:"size:40;uniqueIndex:unique_index"`
	V2    sql.NullString `gorm:"size:40;uniqueIndex:unique_index"`
	V3    sql.NullString `gorm:"size:40;uniqueIndex:unique_index"`
	V4    sql.NullString `gorm:"size:40;uniqueIndex:unique_index"`
	V5    sql.NullString `gorm:"size:40;uniqueIndex:unique_index"`
}


func main() {
	db, err := sql.Open("postgres","port=30001 user=db password=Aa123456 host=192.168.3.121 dbname=testdb sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	var aaa string
	aaa = "2"
	qrstr := "select * from casbin_rule where v0=$1;"
	rows, err := db.Query(qrstr,aaa)
	if err != nil {
		fmt.Println(err)
	}
	var str CasbinRule
	for rows.Next(){

		rows.Scan(&str.ID,&str.PType,&str.V0,&str.V1,&str.V2,&str.V3,&str.V4,&str.V5)
		fmt.Println("aaa",str)
	}
	str.ID = 12
	fmt.Println(str)
	strr := make(map[string]interface{})
	for ro :=range strr{
		fmt.Println(ro)
	}
	strr["p_type"] = str.PType
	strr["v0"] = str.V0
	strr["v1"] = str.V1
	strr["v2"] = str.V2
	//strr["v3"] = str.V3
	//strr["v4"] = str.V4
	//strr["v5"] = str.V5
	ExecSql(db,strr)
	//var qstrbt strings.Builder
	//qstrbt.WriteString("insert into casbin_rule(")
	//var valuestr string
	//i :=0
	//for k, v := range strr{
	//	qstrbt.WriteString(",")
	//	qstrbt.WriteString(k)
	//	valuestr += "$"+strconv.Itoa(i)
	//	i++
	//}
	//qstrbt.WriteString(") values(")
	//qstr = qstrbt.String() + valuestr+ ");"

	//_,err =db.Exec("insert into casbin_rule values($1,$2,$3,$4,$5,$6,$7,$8)","",str.PType,str.V0,str.V1,str.V2,str.V3,str.V4,str.V5)
	//_,err =db.Exec(qstr,str.PType,str.V0,str.V1,str.V2,str.V3,str.V4,str.V5)
	//if err != nil {
	//	fmt.Println(err)
	//}

}

func ExecSql(db *sql.DB,arg map[string]interface{}){
	var test []interface{}

	var qstrbt strings.Builder
	qstrbt.WriteString("insert into casbin_rule(")
	var valuestr string
	i :=0
	for k, v := range arg{
		if v == nil || v ==""{

		} else {
			if i != 0 {
				qstrbt.WriteString(",")
				valuestr += ","
			}
			qstrbt.WriteString(k)
			i++
			valuestr += "$"+strconv.Itoa(i)
			test = append(test,v)
		}

	}
	qstrbt.WriteString(") values(")
	qstr := qstrbt.String() + valuestr+ ");"
	fmt.Println(qstr)
	fmt.Println(test)
	_,err :=db.Exec(qstr,test...)
	if err != nil {
		fmt.Println(err)
	}
}
