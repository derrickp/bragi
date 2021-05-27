package spotify

type HistoricalListen struct {
	EndTime    string `json:"endTime" binding:"required"`
	ArtistName string `json:"artistName" binding:"required"`
	TrackName  string `json:"trackName" binding:"required"`
	MsPlayed   int    `json:"msPlayed" binding:"required"`
}
