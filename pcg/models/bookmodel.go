package models

import (
	"BookApiGin/pcg/config"
	"gorm.io/gorm"
)

//db - это глобальная переменная, хранящая указатель на экземпляр базы данных GORM. Она используется для выполнения запросов к базе данных.

var db *gorm.DB

//Book - это структура, представляющая модель книги. Она включает в себя поля:
//gorm.Model: Это встроенная структура GORM, содержащая поля ID, CreatedAt, UpdatedAt, и DeletedAt для поддержки стандартных конвенций GORM.
//Name, Autor, Publication: Эти поля представляют собой данные книги и имеют теги для JSON сериализации.

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Autor       string `json:"autor"`
	Publication string `json:"publication"`
}

// инициализация базы данных
//init() - это специальная функция в Go, которая вызывается автоматически при инициализации пакета.
//config.Connect(): Это функция из пакета config, которая устанавливает соединение с базой данных.
//db = config.GetDb(): Здесь получаем уже установленное соединение с базой данных и присваиваем его глобальной переменной db.
//db.AutoMigrate(&Book{}): Этот метод GORM автоматически мигрирует схему базы данных, чтобы она соответствовала структуре модели Book.
//	Он создаст таблицу, добавит недостающие столбцы, индексы и ограничения, но не удалит неиспользуемые столбцы, чтобы защитить данные

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

//var Books []Book: Здесь объявляется слайс структуры Book для хранения результатов запроса.
//db.Find(&Books): Этот метод GORM выполняет SELECT запрос, который загружает все записи из таблицы books и заполняет слайс Books.
//return Books: Функция возвращает слайс с полученными книгами.

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func (b *Book) CreateBook() *Book {
	config.GetDb().Create(&b)
	return b
}
