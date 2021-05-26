package spotify

type Listen struct {
	EndTime    string `json:"endTime"`
	ArtistName string `json:"artistName"`
	TrackName  string `json:"trackName"`
	MsPlayed   int    `json:"msPlayed"`
}
