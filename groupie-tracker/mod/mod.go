package mod

// Data is a struct representing the artist data.
type Data struct {
	ID          int      `json:"id"`
	Image       string   `json:"image"`
	Name        string   `json:"name"`
	Members     []string `json:"members"`
	Creation    int      `json:"creationDate"`
	FirstAlbum  string   `json:"firstAlbum"`
	Locations   string   `json:"locations"`
	ConcertDates string  `json:"concertDates"`
	Relations   string   `json:"relations"`
}