package swagger

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/boltdb/bolt"
)

func GetPlanetsByPage(page int) [10]string {
	return traversalAndFetch("planets", (page-1)*10)
}

func GetStarshipsByPage(page int) [10]string {
	return traversalAndFetch("starships", (page-1)*10)
}

func GetFilmsByPage(page int) [10]string {
	return traversalAndFetch("films", (page-1)*10)
}

func GetSpeciesByPage(page int) [10]string {
	return traversalAndFetch("species", (page-1)*10)
}

func GetVehiclesByPage(page int) [10]string {
	return traversalAndFetch("vehicles", (page-1)*10)
}

func GetPeopleByPage(page int) [10]string {
	return traversalAndFetch("people", (page-1)*10)
}

func traversalAndFetch(bucketName string, firstItem int) [10]string {
	var result [10]string
	db, err := bolt.Open("db/my.db", 0600, nil)
	//fmt.Println("5")
	if err != nil {
		log.Fatal(err)
	}
	var count int
	var number int
	count = 0
	number = 0
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName)) //这个桶必须存在！！！
		b.ForEach(func(k, v []byte) error {
			if count < 10 && number >= firstItem {
				result[count] = string(v[:])
				count++
			}
			number++
			return nil
		})
		return nil
	})
	defer db.Close()
	return result
}

func GetPlanetByID(id string) string {
	return queryInformation(id, "planets")
}

func GetStarshipByID(id string) string {
	return queryInformation(id, "starships")
}

func GetFilmByID(id string) string {
	return queryInformation(id, "films")
}

func GetSpeciesByID(id string) string {
	return queryInformation(id, "species")
}

func GetVehicleByID(id string) string {
	return queryInformation(id, "vehicles")
}

func GetPeopleByID(id string) string {
	return queryInformation(id, "people")
}

func queryInformation(index string, bucketName string) string {
	var result string = ""
	//fmt.Println("4")
	db, err := bolt.Open("db/my.db", 0600, nil)
	// fmt.Println("5")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(index))
		result = string(v[:])
		// fmt.Println(result)
		return nil
	}); err != nil {
		log.Fatal("view error :", err.Error())
	}
	defer db.Close()
	// fmt.Println(result)
	return result
}

func initDB() {
	xlsx, err := excelize.OpenFile("db/StarWars.xlsx")
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
			//fmt.Print(colCell, "\t")
		}
		insertFunction(str1, str2, "planets")
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
		insertFunction(str1, str2, "starships")
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
		insertFunction(str1, str2, "vehicles")
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
		insertFunction(str1, str2, "people")
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
		insertFunction(str1, str2, "films")
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
		insertFunction(str1, str2, "species")
	}

}

func insertFunction(str1 string, str2 string, bucketName string) {
	db, err := bolt.Open("db/my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketName)); err != nil { //判断是否存在
			log.Fatal("create failed", err.Error())
			return err
		}
		b := tx.Bucket([]byte(bucketName))
		err = b.Put([]byte(str1), []byte(str2))
		return err
	}); err != nil {
		log.Fatal("update error is:", err.Error())
	}
	defer db.Close()
}

// 修改密码
func ChangePassword(username string, newPassword string) {
	insertFunction(username, newPassword, "password")
}

// 修改邮箱
func ChangeEmail(username string, email string) {
	insertFunction(username, email, "email")
}

// 查询用户
func QueryUser(username string) [2]string {
	var result [2]string
	result[0] = queryInformation(username, "password")
	result[1] = queryInformation(username, "email")
	return result
}

// 插入用户
func InsertUser(username string, password string, email string) {
	insertFunction(username, password, "password")
	insertFunction(username, email, "email")
}

// 判断一个用户是否存在
func IsExist(username string) bool {
	var isexist bool
	isexist = false
	db, err := bolt.Open("db/my.db", 0600, nil)
	//fmt.Println("5")
	if err != nil {
		log.Fatal(err)
	}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("password")) //这个桶必须存在！！！
		b.ForEach(func(k, v []byte) error {
			if string(k[:]) == username {
				isexist = true
			}
			return nil
		})
		return nil
	})
	defer db.Close()
	return isexist
}
