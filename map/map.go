package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	//added new item in map
	websites["Linedin"] = "https://linkedin.com"

	fmt.Println(websites)

	//to delete item in map
	delete(websites, "Amazon")
	fmt.Println(websites)

	//make function

	userNames := make([]string, 2, 5) //capacity is 5 , maximum number of elements it can hold for array

	userNames[0] = "dinesh"
	userNames[1] = "ramesh"
	fmt.Println(userNames)

	//typealias for map
	courses := make(floatMap, 3) //allocate memory to map
	//code make more efficient using make function
	courses["go"] = 4.7
	courses["react"] = 4.8
	courses["angular"] = 4.9 //map is reference type

	courses.output()

	//for loops for map, slics and arrays

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courses {
		fmt.Println(key, value)
	}

}

//map and structs are used to store data
//map is used to store data in key value pairs
//map is unordered list
//map is reference type
//struct is used to store data in key value pairs when you know the data structure already
//struct is ordered list
//struct is value type
//map is used to store data when you dont know the data structure already
