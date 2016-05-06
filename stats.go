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
  BoardsId    []string
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

  for boardid := range config.BoardsId {
    board, err := trello.Board(config.BoardsId[boardid])
    if err != nil {
      fmt.Println("error:", err)
      return
    }
    fmt.Printf("%s\n", board.Name)
    lists, err := board.Lists()
    if err != nil {
      fmt.Println("error:", err)
      return
    }
    for listid := range lists {
      
    }
  }

}
