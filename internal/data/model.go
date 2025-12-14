package data

// API Response

type BulkGameResponse struct {
	Data       []BulkGameDto `json:"data"`
	Pagination PaginationDto `json:"pagination"`
}

type LeaderboardResponse struct {
	Data LeaderboardDto `json:"data"`
}

type UserRunsResponse struct {
	Data       []RunDto      `json:"data"`
	Pagination PaginationDto `json:"pagination"`
}

type GameResponse struct {
	Data GameDto `json:"data"`
}

// DTO

type RunDto struct {
	Id        string            `json:"id"`
	Weblink   string            `json:"weblink"`
	Game      string            `json:"game"`
	Level     *string           `json:"level"`
	Category  *string           `json:"category"`
	Videos    RunVideoDto       `json:"videos"`
	Comment   string            `json:"comment"`
	Status    RunStatusDto      `json:"status"`
	Players   []PlayerDto       `json:"players"`
	Date      string            `json:"date"`
	Submitted string            `json:"submitted"`
	Times     RunTimesDto       `json:"times"`
	System    RunSystemDto      `json:"system"`
	Splits    *struct{}         `json:"splits"`
	Values    map[string]string `json:"values"`
	Links     []LinkDto         `json:"links"`
}

type ValuesDto struct{}

type LeaderboardDto struct {
	Weblink   string    `json:"weblink"`
	Game      string    `json:"game"`
	Category  *string   `json:"category"`
	Level     *string   `json:"level"`
	Platform  *string   `json:"platform"`
	Region    *string   `json:"region"`
	Emulators *string   `json:"emulators"`
	VideoOnly bool      `json:"video-only"`
	Timing    string    `json:"timing"`
	Values    ValuesDto `json:"values"`
	Runs      []struct {
		Place int                 `json:"place"`
		Run   LeaderboardEntryDto `json:"run"`
	} `json:"runs"`
	Links []LinkDto `json:"links"`
}

type LeaderboardEntryDto struct {
	Id        string            `json:"id"`
	Weblink   string            `json:"weblink"`
	Game      string            `json:"game"`
	Level     *string           `json:"level"`
	Category  *string           `json:"category"`
	Videos    RunVideoDto       `json:"videos"`
	Comment   string            `json:"comment"`
	Status    RunStatusDto      `json:"status"`
	Players   []PlayerDto       `json:"players"`
	Date      string            `json:"date"`
	Submitted string            `json:"submitted"`
	Times     RunTimesDto       `json:"times"`
	System    RunSystemDto      `json:"system"`
	Splits    *struct{}         `json:"splits"`
	Values    map[string]string `json:"values"`
}

type RunVideoDto struct {
	Links []struct {
		uri string `json:"uri"`
	} `json:"links"`
}

type LinkDto struct {
	Rel string `json:"rel"`
	Uri string `json:"uri"`
}

type RunSystemDto struct {
	Platform string `json:"platform"`
	Emulated bool   `json:"emulated"`
	Region   string `json:"region"`
}

type RunStatusDto struct {
	Status     string `json:"status"`
	Examiner   string `json:"examiner"`
	VerifyDate string `json:"verify-date"`
}

type RunTimesDto struct {
	Primary          string  `json:"primary"`
	PrimaryT         float64 `json:"primary_t"`
	Realtime         string  `json:"realtime"`
	RealtimeT        float64 `json:"realtime_t"`
	RealtimeNoLoads  string  `json:"realtime_noloads"`
	RealtimeNoLoadsT float64 `json:"realtime_noloads_t"`
	InGame           string  `json:"ingame"`
	InGameT          float64 `json:"ingame_t"`
}

type PlayerDto struct {
	Rel string `json:"rel"`
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

type GameDto struct {
	Id    string `json:"id"`
	Names struct {
		International string  `json:"international"`
		Japanese      *string `json:"japanese"`
		Twitch        string  `json:"twitch"`
	} `json:"names"`
	BoostReceived       int    `json:"boostReceived"`
	BoostDistinctDonors int    `json:"boostDistinctDonors"`
	Abbreviation        string `json:"abbreviation"`
	Weblink             string `json:"weblink"`
	Discord             string `json:"discord"`
	Released            int    `json:"released"`
	ReleaseDate         string `json:"release-date"`
	Ruleset             struct {
		ShowMilliseconds    bool     `json:"show-milliseconds"`
		RequireVerification bool     `json:"require-verification"`
		RequireVideo        bool     `json:"require-video"`
		RunTimes            []string `json:"run-times"`
		DefaultTime         string   `json:"default-time"`
		EmulatorsAlowed     bool     `json:"emulators-allowed"`
	} `json:"ruleset"`
	Romhack    bool              `json:"romhack"`
	Gametypes  []string          `json:"gametypes"`
	Platforms  []string          `json:"platforms"`
	Regions    []string          `json:"regions"`
	Genres     []string          `json:"genres"`
	Engines    []string          `json:"engines"`
	Developers []string          `json:"developers"`
	Publishers []string          `json:"publishers"`
	Moderators map[string]string `json:"moderators"`
	Created    *string           `json:"created"`
	Assets     map[string]struct {
		Uri *string `json:"uri"`
	} `json:"assets"`
	Links  []LinkDto `json:"links"`
	Levels struct {
		Data []LevelDto `json:"data"`
	} `json:"levels"`
	Categories struct {
		Data []CategoryDto `json:"data"`
	} `json:"categories"`
	Variables struct {
		Data []VariableDto `json:"data"`
	} `json:"variables"`
}

type LevelDto struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Weblink string    `json:"weblink"`
	Rules   *string   `json:"rules"`
	Links   []LinkDto `json:"links"`
}

type CategoryDto struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Weblink string  `json:"weblink"`
	Type    string  `json:"type"`
	Rules   *string `json:"rules"`
	Players struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"players"`
	Miscellaneous bool      `json:"miscellaneous"`
	Links         []LinkDto `json:"links"`
}

type VariableDto struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Category *string `json:"category"`
	Scope    struct {
		Type string `json:"type"`
	} `json:"scope"`
	Mandatory   bool `json:"mandatory"`
	UserDefined bool `json:"user-defined"`
	Obsoletes   bool `json:"obsoletes"`
	Values      struct {
		Note    string            `json:"_note"`
		Choices map[string]string `json:"choices"`
		Values  map[string]struct {
			Label string `json:"label"`
			Rules string `json:"rules"`
			Flags struct {
				Miscellaneous bool `json:"miscellaneous"`
			} `json:"flags"`
		} `json:"values"`
		Default string `json:"default"`
	} `json:"values"`
	IsSubcategory bool      `json:"is-subcategory"`
	Links         []LinkDto `json:"links"`
}

type BulkGameDto struct {
	Id    string `json:"id"`
	Names struct {
		International string  `json:"international"`
		Japanese      *string `json:"japanese"`
	} `json:"names"`
	Abbreviation string `json:"abbreviation"`
	Weblink      string `json:"weblink"`
}

type PaginationDto struct {
	Offset int `json:"offset"`
	Max    int `json:"max"`
	Size   int `json:"size"`
	Links  []struct {
		Rel string `json:"rel"`
		Uri string `json:"uri"`
	} `json:"links"`
}
