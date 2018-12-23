package swagger

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_Driver = "root:123456@tcp(127.0.0.1:3306)/starWar?charset=utf8"
)

func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	return isOpen, db
}

func QueryFromDB(tableName string, index int) string {
	opend, db := OpenDB()
	var id int
	var information string
	if opend {
		rows, err := db.Query("SELECT * FROM " + tableName + " where id=" + strconv.Itoa(index))
		if err != nil {
			fmt.Println(err)
			return ""
		}
		for rows.Next() {
			err = rows.Scan(&id, &information)
		}
	} else {
		fmt.Println("open database failed")
	}
	//fmt.Println("test is:" + test)
	defer db.Close()
	return information
}

func insertToDB(tableName string, str1 string, str2 string) {
	opend, db := OpenDB()
	if opend {
		stmt, _ := db.Prepare("insert " + tableName + " set id=?,information=?")
		var id int
		id, _ = strconv.Atoi(str1)
		res, err := stmt.Exec(id, str2)
		index, _ := res.LastInsertId()
		var temp int64
		temp = index
		temp++
		if err != nil {
			fmt.Println("插入数据失败")
			fmt.Println(err)
		} else {
			fmt.Println("插入数据成功")
		}
	} else {
		fmt.Println("open database failed")
	}
	defer db.Close()
}

func initDB() {
	xlsx, err := excelize.OpenFile("StarWars.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the planets.
	rows := xlsx.GetRows("planets")
	var str1 string = ""
	var str2 string = ""
	for _, row := range rows {
		for i, colCell := range row {
			if i == 0 {
				str1 = colCell
			}
			if i == 1 {
				str2 = colCell
			}
		}
		//生成表并插入信息
		insertToDB("planets", str1, str2)
	}
	// Get all the rows in the starships
	rows = xlsx.GetRows("starships")
	for _, row := range rows {
		for i, colCell := range row {
			if i == 0 {
				str1 = colCell
			}
			if i == 1 {
				str2 = colCell
			}
			//fmt.Print(colCell, "\t")
		}
		//生成表并插入信息
		insertToDB("starships", str1, str2)
	}

	// Get all the rows in the vehicles
	rows = xlsx.GetRows("vehicles")
	for _, row := range rows {
		for i, colCell := range row {
			if i == 0 {
				str1 = colCell
			}
			if i == 1 {
				str2 = colCell
			}
			//fmt.Print(colCell, "\t")
		}
		//生成表并插入信息
		insertToDB("vehicles", str1, str2)
	}

	// Get all the rows in the people
	rows = xlsx.GetRows("people")
	for _, row := range rows {
		for i, colCell := range row {
			if i == 0 {
				str1 = colCell
			}
			if i == 1 {
				str2 = colCell
			}
			//fmt.Print(colCell, "\t")
		}
		//生成表并插入信息
		insertToDB("people", str1, str2)
	}

	// Get all the rows in the films
	rows = xlsx.GetRows("films")
	for _, row := range rows {
		for i, colCell := range row {
			if i == 0 {
				str1 = colCell
			}
			if i == 1 {
				str2 = colCell
			}
			//fmt.Print(colCell, "\t")
		}
		//生成表并插入信息
		insertToDB("films", str1, str2)
	}

	// Get all the rows in the species
	rows = xlsx.GetRows("species")
	for _, row := range rows {
		for i, colCell := range row {
			if i == 0 {
				str1 = colCell
			}
			if i == 1 {
				str2 = colCell
			}
			//fmt.Print(colCell, "\t")
		}
		//生成表并插入信息
		insertToDB("species", str1, str2)
	}
}

func GetPage(tableName string, firstItem int) [10]string {
	opend, db := OpenDB()
	var count int
	var number int
	var id int
	var information string
	var result [10]string
	count = 0
	number = 0
	if opend {
		rows, err := db.Query("SELECT * FROM " + tableName)
		if err != nil {
			fmt.Println("查询失败")
		}
		for rows.Next() {
			err = rows.Scan(&id, &information)
			if count < 10 && number >= firstItem {
				result[count] = information
				count++
			}
			number++
		}
	} else {
		fmt.Println("open database failed")
	}
	defer db.Close()
	return result
}

func QueryUserFromDB(user string) [2]string {
	opend, db := OpenDB()
	var result [2]string
	var username string
	var password string
	var email string
	if opend {
		rows, err := db.Query("SELECT * FROM USER")
		if err != nil {
			fmt.Println("查询失败")
		}
		for rows.Next() {
			err = rows.Scan(&username, &password, &email)
			if username == user {
				result[0] = password
				result[1] = email
				break
			}
		}
	} else {
		fmt.Println("open database failed")
	}
	defer db.Close()
	return result
}

func UpdateDB(username string, changeAttr string, attrName string) {
	opend, db := OpenDB()
	if opend {
		stmt, _ := db.Prepare("update USER set " + attrName + "=? where username=?")
		res, _ := stmt.Exec(changeAttr, username)
		affect, _ := res.RowsAffected()
		fmt.Println("更新数据：", affect)
	} else {
		fmt.Println("open database failed")
	}
	defer db.Close()
}

//------------------------------------------------------------------------------------------------------------------
func GetPlanetByID(id string) string {
	temp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(" 操作失败")
	}
	//fmt.Println(temp)
	return QueryFromDB("planets", temp)
}

func GetStarshipByID(id string) string {
	temp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(" 操作失败")
	}
	return QueryFromDB("starships", temp)
}

func GetFilmByID(id string) string {
	temp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(" 操作失败")
	}
	return QueryFromDB("films", temp)
}

func GetSpeciesByID(id string) string {
	temp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(" 操作失败")
	}
	return QueryFromDB("species", temp)
}

func GetVehicleByID(id string) string {
	temp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(" 操作失败")
	}
	return QueryFromDB("vehicles", temp)
}

func GetPeopleByID(id string) string {
	temp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(" 操作失败")
	}
	return QueryFromDB("people", temp)
}

func GetPlanetsByPage(page int) [10]string {
	return GetPage("planets", (page-1)*10)
}

func GetStarshipsByPage(page int) [10]string {
	return GetPage("starships", (page-1)*10)
}

func GetFilmsByPage(page int) [10]string {
	return GetPage("films", (page-1)*10)
}

func GetSpeciesByPage(page int) [10]string {
	return GetPage("species", (page-1)*10)
}

func GetVehiclesByPage(page int) [10]string {
	return GetPage("vehicles", (page-1)*10)
}

func GetPeopleByPage(page int) [10]string {
	return GetPage("people", (page-1)*10)
}

// 插入用户
func insertUser(username string, password string, email string) {
	opend, db := OpenDB()
	if opend {
		stmt, _ := db.Prepare("insert USER set username=?, password=?, email=?")
		res, err := stmt.Exec(username, password, email)
		index, _ := res.LastInsertId()
		var temp int64
		temp = index
		temp++
		if err != nil {
			fmt.Println("数据插入失败")
		} else {
			fmt.Println("数据插入成功")
		}
	} else {
		fmt.Println("open database failed")
	}
	defer db.Close()
}

// 查询用户
func QueryUser(username string) [2]string {
	return QueryUserFromDB(username)
}

// 修改密码
func ChangePassword(username string, newPassword string) {
	UpdateDB(username, newPassword, "password")
}

// 修改邮箱
func ChangeEmail(username string, email string) {
	UpdateDB(username, email, "email")
}

// 判断一个用户是否存在
func IsExist(user string) bool {
	opend, db := OpenDB()
	var isexist bool
	isexist = false
	var username string
	var password string
	var email string
	if opend {
		rows, err := db.Query("SELECT * FROM USER")
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			err = rows.Scan(&username, &password, &email)
			if user == username {
				isexist = true
				break
			}
		}
	} else {
		fmt.Println("open database failed")
	}
	defer db.Close()
	return isexist
}
