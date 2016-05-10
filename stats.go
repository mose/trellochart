package main

import (
  "os"
  "fmt"
  "time"
  "regexp"
  "strings"
  "text/template"
  "encoding/json"
  "github.com/VojtechVitek/go-trello"
)

type Config struct {
  ApiKey      string
  Token       string
  User        string
}

type ListsStats struct {
  Name          string   `json:"name"`
  Count         int      `json:"count"`
}

type BoardStats struct {
  Name          string       `json:"name"`
  Urlname       string       `json:"urlname"`
  Lists        []ListsStats `json:"lists"`
}

type Stats struct {
  DateTaken string        `json:"date"`
  Boards    []BoardStats  `json:"boards"`
}

func buildIndex(s Stats) {
  const index = `
<html><head><title>Trello Stats</title></head><body><ul>
{{range .Boards}}<li><a href="{{.Urlname}}.html">{{.Name}}</a></li>
{{end}}
<div>Last update: {{.DateTaken}}</div>
</ul></body></html>
`
  t := template.New("Index")
  t, _ = t.Parse(index)
  t.Execute(os.Stdout, s)
}

func check(e error) {
  if e != nil {
    fmt.Println("error: ", e)
    panic(e)
  }
}

func urlize(s string) string {
  reg, _ := regexp.Compile("[^A-Za-z0-9]+")
  s = reg.ReplaceAllString(s, "-")
  s = strings.ToLower(strings.Trim(s, "-"))
  return s
}

func main() {
  file, _ := os.Open("config.json")
  decoder := json.NewDecoder(file)
  config := Config{}
  stats := Stats{}
  stats.DateTaken = time.Now().Format("Sat Mar  7 11:06:39 PST 2015")
  err := decoder.Decode(&config)
  check(err)

  trello, err := trello.NewAuthClient(config.ApiKey, &config.Token)
  check(err)

  user, err := trello.Member(config.User)
  check(err)

  boards, err := user.Boards()
  check(err)

  if len(boards) > 0 {
    for _, board := range boards {
      boardStat := BoardStats{}
      boardStat.Name = board.Name
      boardStat.Urlname = urlize(board.Name)
      // fmt.Printf("Board: %s\n", board.Name)
      lists, err := board.Lists()
      check(err)
      for _, list := range lists {
        // fmt.Printf("|  list: %s\n", list.Name)
        cards, _ := list.Cards()
        boardStat.Lists = append(boardStat.Lists, ListsStats{list.Name, len(cards)} )
      }
      stats.Boards = append(stats.Boards, boardStat)
    }
  }
  // fmt.Println("-------------------\n")
  // fmt.Printf("%+v\n", stats)
  fmt.Println("-------------------\n")
  buildIndex(stats)


}
