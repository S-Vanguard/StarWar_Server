package swagger

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/boltdb/bolt"
)

func traversal() {
	db, err := bolt.Open("db/my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("planets")) //这个 桶 必须存在！！！
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
	defer db.Close()
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
