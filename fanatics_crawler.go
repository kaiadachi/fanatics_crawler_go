package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "flag"
  "strconv"
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

func getClubs(start_url string, target_url string, target_css string) []string{
  doc := importDoc(target_url)
  club_urls := getUrlList(doc, target_css, start_url)
  return club_urls
}

func getCategories(start_url string, target_url string) []string{
  doc := importDoc(target_url)
  target_css := "div.allDepartmentsBoxes > ul > li > a"
  category_urls := getUrlList(doc, target_css, start_url)
  return category_urls
}

func getItems(start_url string, target_url string) []string{
  doc := importDoc(target_url)
  target_css := "h4.product-card-title > a"
  item_urls := getUrlList(doc, target_css, start_url)
  return item_urls
}

func countPageNum(doc *goquery.Document) int {
  total_items_s := doc.Find("span.page-count-quantity").Text()
  total_items, _ := strconv.Atoi(total_items_s)
  shou := total_items/72;
  var total_pages int
  if amari := total_items % 72; amari != 0 {
    total_pages = shou + 1
  } else {
    total_pages = shou
  }

  return total_pages
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

func (t TargetItem) getTargetItems(doc *goquery.Document) string{
  t.Name = doc.Find("div.product-title-container > h1").Text()
  return t.Name
}

func main() {
  flag.Parse()
  args := flag.Args()
  fmt.Println(args)

  // init_setting
  const start_url = "https://www.fanatics.com"
  target_url := start_url + "/" + args[0]
  target_css := ".team-list-link"

  // club
  club_urls := getClubs(start_url, target_url, target_css)
  //fmt.Println(club_urls)

  // category
  for _, club_url := range club_urls{
    category_urls := getCategories(start_url, club_url)

    // item
    for _, category_url := range category_urls{
      doc := importDoc(category_url)
      total_pages := countPageNum(doc)
      var total_item_urls[] string
      total_item_urls = getTotalItems(doc, total_pages, total_item_urls, start_url, category_url)

      // target_item
      for _, fainal_target_url := range total_item_urls{
        doc := importDoc(fainal_target_url)
        items := TargetItem{}
        test := items.getTargetItems(doc)
        fmt.Println(test)
      }
    }
  }
}
