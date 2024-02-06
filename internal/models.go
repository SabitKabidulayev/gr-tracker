package internal

type (
	Artists struct {
		Id           int                `json:"id"`
		Image        string             `json:"image"`
		Name         string             `json:"name"`
		Members      []string           `json:"members"`
		CreationDate int                `json:"creationDate"`
		FirstAlbum   string             `json:"firstalbum"`
		Locations    StructLocations    `json:"-"`
		ConcertDates StructConcertDates `json:"-"`
		Relations    StructRelations    `json:"-"`
	}

	StructLocations struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	}

	StructConcertDates struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	}

	StructRelations struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
)
