package tv

type TvShowDetails struct {
	BackdropPath        string                `json:"backdrop_path"`
	CreatedBy           []CreatedBy           `json:"created_by"`
	EpisodeRunTime      []int                 `json:"episode_run_time"`
	FirstAirDate        string                `json:"first_air_date"`
	Genres              []Genres              `json:"genres"`
	Homepage            string                `json:"homepage"`
	ID                  int                   `json:"id"`
	InProduction        bool                  `json:"in_production"`
	Languages           []string              `json:"languages"`
	LastAirDate         string                `json:"last_air_date"`
	LastEpisodeToAir    LastEpisodeToAir      `json:"last_episode_to_air"`
	Name                string                `json:"name"`
	NextEpisodeToAir    interface{}           `json:"next_episode_to_air"`
	Networks            []Networks            `json:"networks"`
	NumberOfEpisodes    int                   `json:"number_of_episodes"`
	NumberOfSeasons     int                   `json:"number_of_seasons"`
	OriginCountry       []string              `json:"origin_country"`
	OriginalLanguage    string                `json:"original_language"`
	OriginalName        string                `json:"original_name"`
	Overview            string                `json:"overview"`
	Popularity          float64               `json:"popularity"`
	PosterPath          string                `json:"poster_path"`
	ProductionCompanies []ProductionCompanies `json:"production_companies"`
	ProductionCountries []ProductionCountries `json:"production_countries"`
	Seasons             []Seasons             `json:"seasons"`
	SpokenLanguages     []SpokenLanguages     `json:"spoken_languages"`
	Status              string                `json:"status"`
	Tagline             string                `json:"tagline"`
	Type                string                `json:"type"`
	VoteAverage         float64               `json:"vote_average"`
	VoteCount           int                   `json:"vote_count"`
}
type CreatedBy struct {
	ID          int    `json:"id"`
	CreditID    string `json:"credit_id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}
type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type LastEpisodeToAir struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}
type Networks struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}
type ProductionCompanies struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}
type ProductionCountries struct {
	Iso31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}
type Seasons struct {
	AirDate      string `json:"air_date"`
	EpisodeCount int    `json:"episode_count"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
}
type SpokenLanguages struct {
	EnglishName string `json:"english_name"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}
