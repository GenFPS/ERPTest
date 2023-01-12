package main

import (
	"database/sql" // позволяет создавать базу данных через интерфейсы Go

	_ "github.com/mattn/go-sqlite3" // драйвер sqllite3 для подключения Go к SQlite
)

func creatorDb() {
    // создадим базу данных product.db
	database, err := sql.Open("sqlite3", "databases/product.db")
    if err != nil {
        panic(err)
    }
    defer database.Close()

	// создаем таблицы для выдачи необходимой информации
    table1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS productOrder (IdOrder INTEGER PRIMARY KEY, ProductType VARCHAR, Quantity INTEGER, RegDate VARCHAR, ExecDate VARCHAR)")
    table1.Exec()
    table1, _ = database.Prepare("INSERT INTO productOrder (ProductType, Quantity, RegDate, ExecDate) VALUES ('ГП5', 2, '09-01-2023', '12-03-2023')")
    table1.Exec()

	table2, _ := database.Prepare("CREATE TABLE IF NOT EXISTS techCardStep1 (CommodityValue VARCHAR, Quantity INTEGER)")
	table2.Exec()
	table2, _ = database.Prepare("INSERT INTO techCardStep1 (CommodityValue, Quantity) VALUES ('Материал 2', 15), ('Материал 3', 20), ('Материал 4', 40), ('Материал 9', 1000), ('Материал 10', 25), ('Материал 11', 3)")
	table2.Exec()

    table3, _ := database.Prepare("CREATE TABLE IF NOT EXISTS techCardStep2 (CommodityValue VARCHAR, Quantity INTEGER)")
	table3.Exec()
	table3, _ = database.Prepare("INSERT INTO techCardStep2 (CommodityValue, Quantity) VALUES ('Материал 2', 8), ('Материал 3', 10), ('Материал 4', 50), ('Материал 9', 30), ('Материал 10', 25), ('Материал 11', 3)")
	table3.Exec()

    table4, _ := database.Prepare("CREATE TABLE IF NOT EXISTS StockBalance (CommodityValue VARCHAR, Quantity INTEGER, Value INTEGER, SUM INTEGER)")
    table4.Exec()
    table4, _ = database.Prepare("INSERT INTO StockBalance (CommodityValue, Quantity, Value, SUM) VALUES ('Материал 1', 450, 120, 54000), ('Материал 2', 890, 115, 102350), ('Материал 3', 800, 52, 41600), ('Материал 4', 630, 32, 20160), ('Материал 5', 40000, 100, 4000000), ('Материал 6', 999, 200, 199800), ('Материал 7', 8700, 50, 435000), ('Материал 8', 10000, 60, 600000), ('Материал 9', 65000, 45, 2925000), ('Материал 10', 4500, 89, 400500), ('Материал 11', 890, 98, 87220)")
    table4.Exec()

    table5, _ := database.Prepare("CREATE TABLE IF NOT EXISTS ChartOfAcc1 (Chart VARCHAR, TypeOfChart VARCHAR, InitialRem INTEGER, WriteOffStepOne INTEGER, WriteOffSteTwo INTEGER, Result INTEGER)")
    table5.Exec()
    table5, _ = database.Prepare("INSERT INTO ChartOfAcc1 (Chart, TypeOfChart, InitialRem, WriteOffStepOne, WriteOffSteTwo, Result) VALUES ('Материалы', 'A', 8865630, -103128, -13818, 8748684), ('ОсновноеПроизводство', 'A', 3430000, 103128, 13818, 3546946)")
    table5.Exec()

    table6, _ := database.Prepare("CREATE TABLE IF NOT EXISTS ChartOfAcc2 (Chart VARCHAR, TypeOfChart VARCHAR, ResultAfterPostings INTEGER, DepartmentQuality INTEGER, Result INTEGER)")
    table6.Exec()
    table6, _ = database.Prepare("INSERT INTO ChartOfAcc2 (Chart, TypeOfChart, ResultAfterPostings, DepartmentQuality, Result) VALUES ('ОсновноеПроизводство', 'A', 3546946, -116946, 3430000), ('ГотоваяПрдукция', 'A', 22000000, 116946, 22116946)")
    table6.Exec()

}