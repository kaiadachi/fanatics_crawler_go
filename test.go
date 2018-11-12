package main

import (
    "github.com/PuerkitoBio/goquery"
    "fmt"
    "reflect"
)

func main() {
    doc, err := goquery.NewDocument("https://www.fanatics.com/nfl/arizona-cardinals/accessories/o-4605+t-58482364+d-42226612+z-9-3255659524")
    if err != nil {
        fmt.Print("url scarapping failed")
    }

    var a string = doc.Url.Html()
    fmt.Println(reflect.TypeOf(a))
}
