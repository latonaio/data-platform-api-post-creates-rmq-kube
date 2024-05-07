package sub_func_complementer

type SDC struct {
	ConnectionKey       string   `json:"connection_key"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	APIType             string   `json:"api_type"`
	Message             Message  `json:"message"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Message struct {
	Header	*Header		`json:"PointTransaction"`
	Friend	*[]Friend	`json:"Friend"`
}

type Header struct {
	Post							int		`json:"Post"`
	PostType						string	`json:"PostType"`
	PostOwner						int		`json:"PostOwner"`
	PostOwnerBusinessPartnerRole	string	`json:"PostOwnerBusinessPartnerRole"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}

type Friend struct {
	Post                 	int     `json:"Post"`
	Friend		            int		`json:"Friend"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}

type NumberRange struct {
	NumberRangeID            string `json:"NumberRangeID"`
	ServiceLabel             string `json:"ServiceLabel"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	LatestNumber             int    `json:"LatestNumber"`
}
type CalculatePostQueryGets struct {
	NumberRangeID                string `json:"NumberRangeID"`
	ServiceLabel                 string `json:"service_label"`
	FieldNameWithNumberRange     string `json:"FieldNameWithNumberRange"`
	PostLatestNumber			 int    `json:"PostLatestNumber"`
}
