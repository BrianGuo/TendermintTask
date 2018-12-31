package main

import "fmt"
import "os"
import "strings"

type City struct {
    name string
    adjacent []*City
}

func containsCity(city_arr []*City, str string) (bool) {
    for _, entry := range(city_arr) {
        if (*entry).name == str {
            return true
        }
    }
    return false
}

func main() {
    file, _ := os.Open("test.map")
    name_to_city := make(map[string]*City)
    b := make([]byte, 100)
    file.Read(b)
    fmt.Println(cap(b))
    fmt.Println(len(b))
    c := string(b)
    fmt.Println(c)
    d := strings.Split(c, "\n")
    for _, element := range d {
        data := strings.Split(element, " ")
        name := data[0]
        var city *City
        if name_to_city[name] != nil {
            city = name_to_city[name]
        } else {
            city = &City{name, make([]*City, 0)}
        }
        fmt.Println(name)
        for _, entry := range data[1:] {
            cityName := entry[strings.Index(entry, "=") + 1 : ]
            if name_to_city[cityName] != nil {
                // City Exists
                adjacent_city := name_to_city[cityName]
                if !containsCity((*city).adjacent, cityName){
                    (*city).adjacent = append((*city).adjacent, adjacent_city)
                    adjacent_city.adjacent = append(adjacent_city.adjacent, city)
                }
            } else {
                // City doesn't exist
                fmt.Printf("%s doesn't exist\n", cityName)
            }
        }
    }
}
