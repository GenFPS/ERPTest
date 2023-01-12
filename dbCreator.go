package main

import (
	"database/sql" // позволяет создавать базу данных через интерфейсы Go

	_ "github.com/mattn/go-sqlite3" // драйвер sqllite3 для подключения Go к SQlite
)

func creator() {
    // создадим базу данных product.db
    database, err := sql.Open("sqlite3", "databases/product.db")
    if err != nil {
        panic(err)
    }
    defer database.Close()

	// создаем таблицы для выдачи необходимой информации
    table1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS productOrder (IdOrder INTEGER PRIMARY KEY, ProductType VARCHAR, Кол-во INTEGER, RegDate VARCHAR, ExecDate VARCHAR)")
    table1.Exec()
    table1, _ = database.Prepare("INSERT INTO productOrder (ProductType, Quantity, RegDate, ExecDate) VALUES ('ГП5', 2, '09-01-2023', '12-03-2023')")
    table1.Exec()

	table2, _ := database.Prepare("CREATE TABLE IF NOT EXISTS techCardStep1 (ТМЦ VARCHAR, Quantity INTEGER)")
	table2.Exec()
	table2, _ = database.Prepare("INSERT INTO techCardStep1 (ТМЦ VARCHAR, Quantity INTEGER) VALUES ('Материал 2', 15), ('Материал 3', 20), (Материал 4', 40), (Материал 9', 1000), (Материал 10', 25), (Материал 11', 3)")
	table2.Exec()
}