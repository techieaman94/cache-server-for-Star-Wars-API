package cacheServer

type People struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Homeworld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
	Created   string   `json:"created"`
	Edited    string   `json:"edited"`
	Url       string   `json:"url"`
}

type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	Url            string   `json:"url"`
}

type Film struct {
	Title        string   `json:"title"`
	EpisodeId    string   `json:"episode_id"`
	OpeningCrawl string   `json:"opening_crawl"`
	Director     string   `json:"director"`
	Producer     string   `json:"producer"`
	ReleaseDate  string   `json:"release_date"`
	Characters   []string `json:"characters"`
	Planets      []string `json:"planets"`
	Starships    []string `json:"starships"`
	Vehicles     []string `json:"vehicles"`
	Species      []string `json:"species"`
	Created      string   `json:"created"`
	Edited       string   `json:"edited"`
	Url          string   `json:"url"`
}

type Species struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	SkinColors      string   `json:"skin_colors"`
	HairHolors      string   `json:"hair_colors"`
	EyeColors       string   `json:"eye_colors"`
	AverageLifespan string   `json:"average_lifespan"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	People          []string `json:"people"`
	Films           []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	Url             string   `json:"url"`
}

type Vehicle struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `json:"vehicle_class"`
	Pilots               []string `json:"pilots"`
	Films                []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	Url                  string   `json:"url"`
}

type Starship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"MGLT"`
	StarshipClass        string   `json:"starship_class"`
	Pilots               []string `json:"pilots"`
	Films                []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	Url                  string   `json:"url"`
}

type PeopleAll struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []People `json:"results"`
}

type FilmAll struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Film `json:"results"`
}
type SpeciesAll struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Species `json:"results"`
}
type VehicleAll struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Vehicle `json:"results"`
}
type StarshipAll struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Starship `json:"results"`
}
type PlanetAll struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Planet `json:"results"`
}
