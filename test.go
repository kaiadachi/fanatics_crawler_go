package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "strconv"
  "sync"
)

func importDoc(url string) *goquery.Document{
  doc, err := goquery.NewDocument(url)
  if err != nil { fmt.Print("url scarapping failed") }
  return doc
}

func getUrlList(doc *goquery.Document, target_css string, start_url string) []string{
  var urls []string
  doc.Find(target_css).Each(func(_ int, s *goquery.Selection) {
    href, _ := s.Attr("href")
    urls = append(urls, start_url + href)
  })
  return urls
}

func getItems(start_url string, target_url string) []string{
  doc := importDoc(target_url)
  target_css := "h4.product-card-title > a"
  item_urls := getUrlList(doc, target_css, start_url)
  return item_urls
}

func getTotalItems(doc *goquery.Document, total_pages int, total_item_urls[] string, start_url string, category_url string) []string{
  for i := 1; i <= total_pages; i++ {
      url_with_param := category_url + "?pageSize=72&pageNumber=" + strconv.Itoa(i)
      fmt.Println(url_with_param)
      item_urls := getItems(start_url, url_with_param)
      total_item_urls = append(total_item_urls, item_urls...)
  }
  return total_item_urls
}

type TargetItem struct{
  Name string
}

func (t *TargetItem) getTargetItems(doc *goquery.Document){
  t.Name = doc.Find("div.product-title-container > h1").Text()
}

func main() {
  // init_setting
  const start_url = "https://www.fanatics.com"
  target_url := "https://www.fanatics.com/nfl/arizona-cardinals/accessories/o-4605+t-58482364+d-42226612+z-9-3255659524"
  urls := getItems(start_url, target_url)

    var wg sync.WaitGroup
    for _, url := range urls{
      wg.Add(1)
      go func(){
        defer wg.Done()
        doc := importDoc(url)
        items := &TargetItem{}
        items.getTargetItems(doc)
        fmt.Println(items.Name)
      }()
    }
    wg.Wait()
  }
