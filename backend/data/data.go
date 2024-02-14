package data

// в данном пакете берутся данные из предоставленных JSON и далее преобразуются в формат с которым может работать GO

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/backend/models"
	"io"
	"net/http"
	"strconv"
	"unicode"
)

// объявляем переменную Artist - массив в котором хранятся структуры типа models.Artist (данную структуру мы обозначили в пакете models) для хранения информации о всех артистах
var Artists []models.Artist

// объявляем функцию FetchDataFromJSON - в качестве параметров она принимает пустой интерфейс data (в который будут записываться данные) и ссылку url - ссылка в которой содержаться записываемые данные. Функция возвращает ошибку если что то пошло не так

// В функции FetchDataFromJSON, интерфейс используется потому что пустой интерфейс не имеет определенного типа данных и, следовательно, может хранить значения любого типа. Это позволяет функции FetchData быть более универсальной и использоваться для получения данных различных типов. В момент вызова функции FetchData, передается конкретный тип данных (например, &location, &date, &relation), и функция использует пустой интерфейс для работы с этими данными без необходимости знания конкретного типа заранее

func FetchDataFromJSON(data interface{}, url string) error {
	// Отправляем GET-запрос по указанному URL
	response, err := http.Get(url)
	// Если произошла ошибка при выполнении GET-запроса, выводим сообщение и возвращаем ошибку
	if err != nil {
		fmt.Println("Error making GET request: ", err)
		return err
	}
	defer response.Body.Close() // Закрываем тело ответа после завершения функции. Это важно для предотвращения утечек ресурсов. (memory leak)

	// Читаем тело ответа
	body, err := io.ReadAll(response.Body)
	// Если произошла ошибка при чтении тела ответа, возвращаем ошибку
	if err != nil {
		return err
	}

	// Распаковываем JSON-данные в структуру, указанную в параметре data
	// &data - это взятие адреса переменной, или указатель (в нашем случае указатель на пустой интерфейс). Использование указателя в Go позволяет изменять существующую переменную вместо создания новой (так как меняется содержимое внутри переменно на адрес которой ссылается указатель)
	err = json.Unmarshal(body, &data)
	// Если произошла ошибка при распаковке JSON, возвращаем ошибку
	if err != nil {
		return err
	}
	// Возвращаем nil, если все операции завершились успешно
	return nil
}

// func BasicData() функция, которая перебирает срез Artist и вызывает AdditionalData для каждого артиста.

func BasicData() error {
	for i := 0; i < len(Artists); i++ {
		err := AdditionalData(i + 1)
		if err != nil {
			return err
		}
	}
	return nil
}

func AdditionalData(id int) error {
	location := models.StructLocations{}
	FetchDataFromJSON(&location, "https://groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(id))

	Artists[id-1].Locations = location

	date := models.StructConcertDates{}
	FetchDataFromJSON(&date, "https://groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(id))

	Artists[id-1].ConcertDates = date

	relation := models.StructRelations{}
	FetchDataFromJSON(&relation, "https://groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(id))
	Artists[id-1].Relations = relation

	return nil
}

func IsValid(id string) bool {
	if id == "" {
		return false
	}
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}

func IsRange(id int) bool {
	if id < 1 || id > 52 {
		return false
	}
	return true
}

func ContainsZero(id string) bool {
	if id[0] == '0' {
		return true
	}
	return false
}
