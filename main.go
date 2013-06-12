package main
 
import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
)
 
func main() {
    //file, err := os.Open("names.txt")
    file, err := os.Open("move.backup.10Jun2013.EventEntry.csv")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()
    reader := csv.NewReader(file)
    first := true
    var headers [] string
    reverse_map := map[string]int{}
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
            return
        }
        if first { 
            first = false
            headers = record
            fmt.Println(record) // record has the type []string
            //fmt.Println(headers) // record has the type []string
            for key,header := range headers {
                //fmt.Println(key, header)
                reverse_map[header] = key
            }
            //fmt.Println(reverse_map)
            //fmt.Println(reverse_map["created"])
        } else {
            if record[reverse_map["checkout_id"]] != "" {
                fmt.Println(record[reverse_map["created"]], record[reverse_map["eventid"]] , record[reverse_map["cart2"]], record[reverse_map["first"]],record[reverse_map["last"]],record[reverse_map["age"]],record[reverse_map["gender"]], record[reverse_map["checkout_id"]])
            }
        }

    }
}

