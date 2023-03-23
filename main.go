package main

import (
	"fmt"
	"os"
)

func main() {
	// var person = []map[int]string{
	// 	{
	// 		1 : "Didik",
	// 		2 : "Yanto",
	// 	},
	// 	{
	// 		1 : "dimas",
	// 	},
	// }
	// for index, value :=range(person) {
	// 	fmt.Println("Index %v: %s", index, value)
	// }
	//fmt.Println(person)

	person := map[string]map[string]string{
		"1": {
			"nama":      "Dian",
			"alamat":    "Jl.GajahMada No.1",
			"pekerjaan": "Data Science",
			"alasan":    "Loremipsum",
		},
		"2": {
			"nama":      "Joko",
			"alamat":    "Jl.Kintelan No.4B",
			"pekerjaan": "Programmer",
			"alasan":    "Loremipsum dolorsitamet",
		},
		"3": {
			"nama":      "Fadil",
			"alamat":    "Jl.Pantura",
			"pekerjaan": "Karyawan",
			"alasan":    "Menambah Wawasan",
		},
		"4": {
			"nama":      "Eka",
			"alamat":    "Jl.Pantura",
			"pekerjaan": "Tukang AC",
			"alasan":    "Mencari Ilmu",
		},
		"5": {
			"nama":      "Eka",
			"alamat":    "Jl.Pantura 2",
			"pekerjaan": "Tukang AC",
			"alasan":    "Mencari Ilmu",
		},
	}
		temp := os.Args[1]
		fmt.Println(person[temp])
}


//Mapping
