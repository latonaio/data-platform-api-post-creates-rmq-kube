package requests

type Friend struct {
	Post                 	int     `json:"Post"`
	Friend		            int		`json:"Friend"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}
