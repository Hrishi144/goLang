package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/PuerkitoBio/goquery"
)

func main() {
     var typ string
     var url string
    fmt.Println("Enter the section name in html:")
    fmt.Scanln(&typ)
    fmt.Println("Enter the url of the site:")
    fmt.Scanln(&url)

    if url == "" || typ =="" {
      log.Fatal("Both the url and selector must be provided")
    }
  res ,err :=http.Get(url)
  if err != nil {
     log.Fatal(err)
   }
  defer res.Body.Close()

  if res.StatusCode != 200 {
     log.Fatalf("Status code Error : %d %s",res.StatusCode,res.Status)
  }

   doc,err :=goquery.NewDocumentFromReader(res.Body)
   if err != nil {
     log.Fatal(err)
   }

   doc.Find(typ).Each(func(i int,s *goquery.Selection) {
       fmt.Println("Heading:",s.Text())
        })
}
