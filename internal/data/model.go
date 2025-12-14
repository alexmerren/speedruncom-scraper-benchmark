package data

// API Response

type BulkGameResponse struct {
	Data       BulkGameDto
	Pagination PaginationDto
}

type LeaderboardResponse struct {
	Data LeaderboardDto
}

type UserRunsResponse struct {
	Data       []RunDto
	Pagination PaginationDto
}

type GameResponse struct {
	Data GameDto
}

// DTO

type RunDto struct {
	Id        string
	Weblink   string
	Game      string
	Level     *string
	Category  *string
	Videos    RunVideoDto
	Comment   string
	Status    RunStatusDto
	Players   []PlayerDto
	Date      string
	Submitted string
	Times     RunTimesDto
	System    RunSystemDto
	Splits    *struct{}
	Values    map[string]string
	Links     []LinkDto
}

type ValuesDto struct{}

type LeaderboardDto struct {
	Weblink   string
	Game      string
	Category  *string
	Level     *string
	Platform  *string
	Region    *string
	Emulators *string
	VideoOnly bool // video-only
	Timing    string
	Values    ValuesDto
	Runs      []struct {
		Place int
		Run   LeaderboardEntryDto
	}
	Links []LinkDto
}

type LeaderboardEntryDto struct {
	Id        string
	Weblink   string
	Game      string
	Level     *string
	Category  *string
	Videos    RunVideoDto
	Comment   string
	Status    RunStatusDto
	Players   []PlayerDto
	Date      string
	Submitted string
	Times     RunTimesDto
	System    RunSystemDto
	Splits    *struct{}
	Values    map[string]string
}

type RunVideoDto struct {
	Links []struct {
		uri string
	}
}

type LinkDto struct {
	Rel string
	Uri string
}

type RunSystemDto struct {
	Platform string
	Emulated bool
	Region   string
}

type RunStatusDto struct {
	Status     string
	Examiner   string
	VerifyDate string // verify-date
}

type RunTimesDto struct {
	Primary          string  // primary
	PrimaryT         float64 // primary_t
	Realtime         string  // realtime
	RealtimeT        float64 // realtime_t
	RealtimeNoLoads  string  // realtime_noloads
	RealtimeNoLoadsT float64 // realtime_noloads_t
	InGame           string  // ingame
	InGameT          float64 // ingame_t
}

type PlayerDto struct {
	Rel string
	Id  string
	Uri string
}

type GameDto struct {
	Id    string
	Names struct {
		International string
		Japanese      *string
		Twitch        string
	}
	BoostReceived       int
	BoostDistinctDonors int
	Abbreviation        string
	Weblink             string
	Discord             string
	Released            int
	ReleaseDate         string // release-date
	Ruleset             struct {
		ShowMilliseconds    bool     // show-milliseconds
		RequireVerification bool     // require-verification
		RequireVideo        bool     // require-video
		RunTimes            []string // run-times
		DefaultTime         string   // default-time
		EmulatorsAlowed     bool     // emulators-allowed
	}
	Romhack    bool
	Gametypes  []string
	Platforms  []string
	Regions    []string
	Genres     []string
	Engines    []string
	Developers []string
	Publishers []string
	Moderators map[string]string
	Created    *string
	Assets     map[string]struct {
		Uri *string
	}
	Links  []LinkDto
	Levels struct {
		Data []LevelDto
	}
	Categories struct {
		Data []CategoryDto
	}
	Variables struct {
		Data []VariableDto
	}
}

type LevelDto struct {
	Id      string
	Name    string
	Weblink string
	Rules   *string
	Links   []LinkDto
}

type CategoryDto struct {
	Id      string
	Name    string
	Weblink string
	Type    string
	Rules   *string
	Players struct {
		Type  string
		Value int
	}
	Miscellaneous bool
	Links         []LinkDto
}

type VariableDto struct {
	Id       string
	Name     string
	Category *string
	Scope    struct {
		Type string
	}
	Mandatory   bool
	UserDefined bool // user-defined
	Obsoletes   bool
	Values      struct {
		Note    string // _note
		Choices map[string]string
		Values  map[string]struct {
			Label string
			Rules string
			Flags struct {
				Miscellaneous bool
			}
		}
		Default string
	}
	IsSubcategory bool // is-subcategory
	Links         []LinkDto
}

type BulkGameDto struct {
	Id    string
	Names struct {
		International string
		Japanese      *string
	}
	Abbreviation string
	Weblink      string
}

type PaginationDto struct {
	Offset int
	Max    int
	Size   int
	Links  struct {
		Rel string
		Uri string
	}
}
