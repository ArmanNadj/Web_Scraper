package main

import (
  //"io"
  "io/ioutil"
  "log"
  "net/http"
  "fmt"
  "bufio"
  "strings"
  "os"
  )

func KeywordSearch(html_file string) {
  fmt.Println("\nIn the Go language, strings are implicitly immutable. Thus, they cannot be changed...")
  fmt.Print("Because of this, please enter a single, case-sensitive keyword to be searched for: ")
  var keyword string
  fmt.Scanln(&keyword)

  //opening file to traverse
  file, error := os.Open(html_file)

  if error != nil {
    log.Fatal(error)
  }
  scanner := bufio.NewScanner(file)

  var scan_iterator, scan_count int = 0, 0
  fmt.Println("Beginning a search on the HTML file...\n")
  for scanner.Scan() {
    scan_iterator++
    line := scanner.Text()
    if strings.Contains(line, keyword) {
      fmt.Printf("%s exists on line %d of %s\n", keyword, scan_iterator, html_file)
      scan_count++
    }
  }

  fmt.Printf("\n%s exists on a total of %d lines within the %s file\n\n", keyword, scan_count, html_file);
}

func main() {
  html_file := "website.html"
  fmt.Print("\nEnter a URL to be scraped: ")
  var url string
  fmt.Scanln(&url)
  website_html, error := http.Get(url)

  if error != nil {
    log.Fatal(error)
  }

  defer website_html.Body.Close()
  website_as_string, _ := ioutil.ReadAll(website_html.Body)
  ioutil.WriteFile(html_file, website_as_string, 0600)

  fmt.Printf("\nAll of the HTML code for the given webpage has been stored into the provided %s file\n\n", html_file)

  //KeywordSearch(html_file)

 }
