package data

// в данном пакете берутся данные из предоставленных JSON и далее преобразуются в формат с которым может работать GO

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/backend/models"
	"io"
	"net/http"
	"strconv"
)

// объявляем переменную Artist - массив в котором хранятся структуры типа models.Artist (данную структуру мы обозначили в пакете models) для хранения информации о всех артистах
var Artists []models.Artist

// объявляем функцию FetchDataFromJSON - в качестве параметров она принимает пустой интерфейс data (в который будут записываться данные) и ссылку url - ссылка в которой содержаться записываемые данные. Функция возвращает ошибку если что то пошло не так. Эта функция берет данные JSON по ссылке url и заполняют ими data interface{}

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

// объявляем функцию GetDataForArtist в качестве параметров она принимает id (по этому id находится нужный артист из переменной Artist, которую мы создали в начале пакета). Функция возвращает ошибку если что то пошло не так (ошибки могут возникнуть во время использования FetchDataFromJSON). С помощью этой функции получаем данные о конкретной группе (где и когда будут выступать)

func GetDataForArtist(id int) error {
	// объявляем переменную location - структуру типа models.StructLocations{} (данную структуру мы обозначили в пакете models) в ней хранятся данные о местах выступлений
	location := models.StructLocations{}
	// заполняем структуру location с помощью функции FetchDataFromJSON которую мы объявили ранее
	FetchDataFromJSON(&location, "https://groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(id))
	// сохраняем структуру location в структуре Artist из массива Artists []models.Artist по индексу id - 1. (поскольку массивы индексируются с нуля, нам нужно вычесть 1 из id, чтобы получить правильный индекс в массиве.)
	Artists[id-1].Locations = location

	// аналогично location получаем данные о датах выступления
	date := models.StructConcertDates{}
	FetchDataFromJSON(&date, "https://groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(id))
	Artists[id-1].ConcertDates = date

	// аналогично location получаем данные какие даты соответствуют каким местам выступления
	relation := models.StructRelations{}
	FetchDataFromJSON(&relation, "https://groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(id))
	Artists[id-1].Relations = relation

	// если во время использования функции FetchDataFromJSON не произошло никаких ошибок возвращаем nil
	return nil
}
