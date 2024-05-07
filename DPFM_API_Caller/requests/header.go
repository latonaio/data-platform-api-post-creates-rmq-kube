package requests

type Header struct {
	Post							*int	`json:"Post"`
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
