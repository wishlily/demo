package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"reflect"
	"time"
)

var gUrl = "user=test dbname=test password=111111 host=192.168.61.128 port=5432 sslmode=disable"

type store struct {
	OutradeNo string      `db:"out_trade_no"`
	Time      int64       `db:"time"`
	Base      interface{} `db:"base"`
	Extend    interface{} `db:"extend"`
}

type Base struct {
	Valid      bool
	Time       string `key:"notify_time"`
	OutTradeNo string `key:"out_trade_no"`
	TradeNo    string `key:"trade_no"`
}

func (m Base) Value() (driver.Value, error) {
	if !m.Valid {
		return nil, nil
	}
	return json.Marshal(m)
}

func (m *Base) Scan(src interface{}) error {
	v := reflect.ValueOf(src)
	if !v.IsValid() || v.IsNil() {
		return nil
	}
	if data, ok := src.([]byte); ok {
		return json.Unmarshal(data, m)
	}
	return fmt.Errorf("Could not not decode type %T -> %T", src, m)
}

func main() {
	url := gUrl
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS test (out_trade_no text, time bigint, base json, extend json)"); err != nil {
		log.Panic(err)
	}

	v := store{
		"Hello",
		time.Now().Unix(),
		Base{true, "111", "22", "33"},
		sql.NullString{Valid: false},
	}
	q1 := `INSERT INTO test (out_trade_no, time, base, extend) VALUES (:out_trade_no, :time, :base, :extend)`
	_, err = db.NamedExec(q1, v)
	if err != nil {
		log.Panic(err)
	}
	var m store
	err = db.Get(&m, "SELECT * FROM test LIMIT 1")
	if err != nil {
		log.Panic(err)
	}
	log.Println(m)

	log.Println("list:")
	rows, err := db.Queryx("SELECT * FROM test")
	if err != nil {
		log.Panic(err)
	}
	for rows.Next() {
		var m store
		err = rows.StructScan(&m)
		if err != nil {
			log.Panic(err)
		}
		log.Println(m)
	}

	if _, err := db.Exec("DROP TABLE test"); err != nil {
		log.Panic(err)
	}
}
