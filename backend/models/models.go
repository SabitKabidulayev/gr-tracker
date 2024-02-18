// привет, начинай отсюда
package models

// в данном пакете задаются структуры в которых будут храниться данные, полученные из предоставленных данных в формете JSON

// Go структура (struct) представляет собой тип данных, который позволяет объединить различные поля с разными типами данных в одной записи. Структуры обеспечивают удобный способ организации данных в программе.

// JSON (JavaScript Object Notation) - это формат обмена данными, основанный на способе записи объектов в JavaScript
// пример:
//  {
// 	"имя": "Иван",
// 	"возраст": 30,
// 	"город": "Москва",
// 	"хобби": ["плавание", "фотография"]
//   }

// структура Artist содержит информацию о группе
// данные для структуры Artist берутся из предоставленной ссылки на JSON файл https://groupietrackers.herokuapp.com/api/artists", эти данные будут сохранены в структуре с помощью демаршалинга
// демаршалинг (unmarshaling) означает процесс преобразования данных в формате JSON в формат Go-структуры . Это нужно для того чтобы мы могли работать с полученными данными нашем коде (который написан на GO :)
type (
	Artist struct {
		Id           int                `json:"id"`           // при демаршалинге (преобразовании из JSON в Go-структуру), значение, найденное в JSON-ключе "id", будет присвоено полю Id (аналогично для остальных)
		Image        string             `json:"image"`        // ссылка на изображение группы в формате строки
		Name         string             `json:"name"`         // имя группы в формате строки
		Members      []string           `json:"members"`      // массив строк с членами группы
		CreationDate int                `json:"creationDate"` // дата создания группы в формате числа
		FirstAlbum   string             `json:"firstalbum"`   // название первого альбома в формате строки
		Locations    StructLocations    `json:"-"`            // вложенная структура.  `json:"-"` означает что при демаршалинге если в JSON-данных присутствует поле, соответствующее Locations, оно будет проигнорировано при преобразовании в Go-структуру
		ConcertDates StructConcertDates `json:"-"`            // так же как и с Locations
		Relations    StructRelations    `json:"-"`            // так же как и с Locations
	}

	// структура Locations содержит информацию о местах выступления группы
	// данные для структуры Artist берутся из предоставленной ссылки на JSON файл https://groupietrackers.herokuapp.com/api/locations"
	StructLocations struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"` // массив строк с местами выступления
		Dates     string   `json:"dates"`     // ссылка на JSON с датами выступления в формате строки
	}

	// структура ConcertDates содержит информацию о датах выступления группы
	// данные для структуры Artist берутся из предоставленной ссылки на JSON файл https://groupietrackers.herokuapp.com/api/dates"
	StructConcertDates struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"` // массив строк с датами выступления
	}

	// структура Relations связывает даты выступления с местами выступления
	// данные для структуры Artist берутся из предоставленной ссылки на JSON файл https://groupietrackers.herokuapp.com/api/relations"
	StructRelations struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"` // это map в которой ключи являются местом выступления (в формате строки) а значения ключей это массив строк с датами выступлений
	}
)

// структуры в этом пакете будут заполняться с помощью функции FetchData() которая находится в пакете data (backend/data), переходи туда
// backend/data/data.go
