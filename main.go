package main

import (
    "fmt"
    "AnimalsBD/models"
    "AnimalsBD/database"
    "AnimalsBD/errors"
)

func inputAnimal() (string, animal.Animal) {
    var animalType string
    var age int

    fmt.Print("Введите тип животного с большой буквы (Лев, Жираф, Змея): ")
    fmt.Scan(&animalType)
    fmt.Print("Введите возраст животного: ")
    fmt.Scan(&age)

    switch animalType {
    case "Лев":
        return animalType, &animal.Lion{Age: age}
    case "Жираф":
        return animalType, &animal.Giraffe{Age: age}
    case "Змея":
        return animalType, &animal.Snake{Age: age}
    default:
        fmt.Println("Неизвестный тип животного!")
        return "", nil
    }
}

func main() {
    dbUser := "your_db_user"
    dbPassword := "your_db_password"
    dbName := "your_db_name"

    db, err := database.ConnectToDB(dbUser, dbPassword, dbName)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    defer db.Close()

    err = database.CreateTable(db)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }

    var count int
    fmt.Print("Сколько животных вы хотите добавить? ")
    fmt.Scan(&count)

    for i := 0; i < count; i++ {
        fmt.Printf("Введите данные для животного %d:\n", i+1)
        animalType, animal := inputAnimal()
        if animal != nil {
            insertErr := database.InsertAnimal(db, animalType, animal)
            if insertErr != nil {
                errHandlerErr := errorhandler.HandleInsertError(db, animalType, animal, insertErr)
                if errHandlerErr != nil {
                    fmt.Println("Ошибка при обработке данных:", errHandlerErr)
                    return
                }
            }
        }
    }

    fmt.Println("\nИнформация о животных в БД:")
    rows, err := db.Query("SELECT * FROM animals")
    if err != nil {
        fmt.Println("Ошибка при запросе:", err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var animalType, sound, move string
        var age int
        err = rows.Scan(&id, &animalType, &sound, &move, &age)
        if err != nil {
            fmt.Println("Ошибка при чтении строки:", err)
            return
        }
        fmt.Printf("ID: %d, Тип: %s, Звук: %s, Движение: %s, Возраст: %d\n", id, animalType, sound, move, age)
    }
}
