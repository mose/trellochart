package main

import (
  "os"
  "fmt"
  "time"
  "encoding/json"
  "github.com/VojtechVitek/go-trello"
)

type Config struct {
  ApiKey      string
  Token       string
  User        string
}

type ListsStats struct {
  Listname      string   `json:"name"`
  Count         int      `json:"count"`
}

type BoardStats struct {
  Boardname    string       `json:"name"`
  Lists        []ListsStats `json:"lists"`
}

type Stats struct {
  DateTaken string        `json:"date"`
  Boards    []BoardStats  `json:"boards"`
}

func main() {

  file, _ := os.Open("config.json")
  decoder := json.NewDecoder(file)
  config := Config{}
  stats := Stats{}
  stats.DateTaken = time.Now().Format("Sat Mar  7 11:06:39 PST 2015")
  err := decoder.Decode(&config)
  if err != nil {
    fmt.Println("error:", err)
    return
  }

  trello, err := trello.NewAuthClient(config.ApiKey, &config.Token)
  if err != nil {
    fmt.Println("error:", err)
    return
  }

  user, err := trello.Member(config.User)
  if err != nil {
    fmt.Println("error:", err)
    return
  }

  boards, err := user.Boards()
  if err != nil {
    fmt.Println("error:", err)
    return
  }

  if len(boards) > 0 {
    for _, board := range boards {
      boardStat := BoardStats{}
      boardStat.Boardname = board.Name
      fmt.Printf("Board: %s\n", board.Name)
      lists, err := board.Lists()
      if err != nil {
        fmt.Println("error:", err)
        return
      }
      for _, list := range lists {
        fmt.Printf("|  list: %s\n", list.Name)
        cards, _ := list.Cards()
        boardStat.Lists = append(boardStat.Lists, ListsStats{list.Name, len(cards)} )
      }
      stats.Boards = append(stats.Boards, boardStat)
    }
  }
  fmt.Println("-------------------\n")
  fmt.Printf("%+v\n", stats)
}
