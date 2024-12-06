package errorhandler

import (
    "database/sql"
    "fmt"
    "AnimalsBD/models"
    "AnimalsBD/database"
)

// HandleInsertError обрабатывает ошибку вставки в таблицу
func HandleInsertError(db *sql.DB, animalType string, animal animal.Animal, err error) error {
    if err != nil {
        fmt.Printf("Ошибка вставки данных для типа %s: %v\n", animalType, err)

        // Попытка повторной записи
        fmt.Println("Повторная попытка записи данных...")
        retryErr := database.InsertAnimal(db, animalType, animal)
        if retryErr != nil {
            return fmt.Errorf("Не удалось записать данные для типа %s после повторной попытки: %w", animalType, retryErr)
        }
    }
    return nil
}
