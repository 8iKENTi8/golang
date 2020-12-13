package main

import(
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "fmt")

type User struct{
  Name string `json:"name"`
  Age uint16  `json:"age"`
}

func main(){
  fmt.Println("Работа с MySQL")


  db,err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
  if err != nil{
    panic(err)
    }

    defer db.Close()

    // Установка данных
    // insert, err := db.Query("INSERT INTO `users` (`id`, `name`, `age`) VALUES (NULL, 'Bob', '35')")
    // if err != nil {
    //   panic(err)
    // }
    //
    // defer insert.Close()

    res, err := db.Query("SELECT `name`, `age` FROM `users`")
    if err != nil{
      panic(err)
    }

    for res.Next(){
      var user User
      err = res.Scan(&user.Name, &user.Age)

      if err != nil{
        panic(err)
      }

      fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
    }

    fmt.Println("Подключено к MySQL")
}
