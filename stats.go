package main

import (
  "fmt"
  "os"
  "encoding/json"
  "github.com/VojtechVitek/go-trello"
)

type Config struct {
  ApiKey      string
  Token       string
  User        string
}

type ListsStats struct {
  Listname      string
  Count         int
}

type BoardStats struct {
  Boardname    string
  Lists        []ListsStats
}

type Stats struct {
  Boards    []BoardStats
}

func main() {

  file, _ := os.Open("config.json")
  decoder := json.NewDecoder(file)
  config := Config{}
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
      fmt.Printf("Board: %s\n", board.Name)
      lists, err := board.Lists()
      if err != nil {
        fmt.Println("error:", err)
        return
      }
      for _, list := range lists {
        fmt.Printf("|  list: %s\n", list.Name)
        cards, _ := list.Cards()

      }
    }
  }

}
