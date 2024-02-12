package models

type Artist struct {
	ID           int          `json:"id"`
	Image        string       `json:"image"`
	Name         string       `json:"name"`
	Members      []string     `json:"members"`
	CreationDate int          `json:"creationDate"`
	FirstAlbum   string       `json:"firstAlbum"`
	Locations    Locations    `json:"-"`
	ConcertDates ConcertDates `json:"-"`
	Relations    Relations    `json:"-"`
}

type (
	IndexLocations struct {
		Locations []Locations `json:"index"`
	}
	IndexConcerts struct {
		ConcertDates []ConcertDates `json:"index"`
	}
	IndexRelations struct {
		Relations []Relations `json:"index"`
	}
)

type (
	Locations struct {
		Id           int      `json:"id"`
		Location     []string `json:"locations"`
		ConcertDates ConcertDates
	}
	ConcertDates struct {
		Id           int      `json:"id"`
		ConcertDates []string `json:"dates"`
	}
	Relations struct {
		Id        int                 `json:"id"`
		Relations map[string][]string `json:"datesLocations"`
	}
)

type AllArtists []Artist
