package dpfm_api_processing_formatter

type HeaderUpdates struct {
	Post							int		`json:"Post"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
}

type FriendUpdates struct {
	Post			int     `json:"Post"`
	Friend			int		`json:"Friend"`
	LastChangeDate	string	`json:"LastChangeDate"`
	LastChangeTime	string	`json:"LastChangeTime"`
}
