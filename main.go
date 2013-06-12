package main
 
import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "sort"
)
 

type Entry struct {
    Values [] string
}

type ByCreateDate [] Entry

func (this ByCreateDate) Len() int {
    //return len(this.Entry)
    //return len(this[0])
    fmt.Println("Len: %v\n", this)
    return len(this)
}


func (this ByCreateDate) Less(i, j int) bool {
    fmt.Println("Less: %v i: %v j: %v\n", this,i,j)
    return this[i].Values[0] < this[j].Values[0]
}
func (this ByCreateDate) Swap(i, j int) {
    fmt.Println("Swap: %v i: %v j: %v\n", this,i,j)
    this[i], this[j] = this[j], this[i]
}

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
    //var data [][] string
    data2 := make([] Entry, 10000)
    data := make([][]string, 10000)
    //data := make([][]string, 10)
    counter := 0
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
            //fmt.Println(record) // record has the type []string
            //fmt.Println(headers) // record has the type []string
            for key,header := range headers {
                //fmt.Println(key, header)
                reverse_map[header] = key
            }
            //fmt.Println(reverse_map)
            //fmt.Println(reverse_map["created"])
        } else {
            if record[reverse_map["checkout_id"]] != "" {
                //fmt.Println(record[reverse_map["created"]], record[reverse_map["eventid"]] , record[reverse_map["cart2"]], record[reverse_map["first"]],record[reverse_map["last"]],record[reverse_map["age"]],record[reverse_map["gender"]], record[reverse_map["checkout_id"]])
                //data[key] := [record[reverse_map["created"]], record[reverse_map["eventid"]] , record[reverse_map["cart2"]], record[reverse_map["first"]],record[reverse_map["last"]],record[reverse_map["age"]],record[reverse_map["gender"]], record[reverse_map["checkout_id"]]]
                //data[reverse_map["created"]][0] := record[reverse_map["created"]]

                //data[counter] = make([] string, 100)
                //data[counter] = make([] string, len(headers))
                data[counter] = make([] string, len(headers))
                var entry Entry
                entry.Values = make([] string, len(headers))
                data2[counter] = entry
                //counterinside := 0
                for key,value := range record {
                    data[counter][key] = value
                    data2[counter].Values[key] = value
                    //data[counter][counterinside] = value
                    //counterinside++
                }

                /*
                for key,value := range record {
                    data[key][0] := value
                }
                */
                //fmt.Println(data2[counter])
                //fmt.Println(ByCreateDate(data2).Len())

                counter++

            }
        }
        //fmt.Println(len(headers))
    }
    //fmt.Println(sort.Strings(data))
    //fmt.Println(sort.Sort(data))

    sort.Sort(ByCreateDate(data2))
    //fmt.Println(data2)

    //fmt.Println(data)
}

