package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-post-creates-rmq-kube/DPFM_API_Input_Reader"
	//dpfm_api_processing_formatter "data-platform-api-point-transaction-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-post-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToFriendCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Friend, error) {
	friends := make([]Friend, 0)

	for _, data := range *subfuncSDC.Message.Friend {
		friend, err := TypeConverter[*Friend](data)
		if err != nil {
			return nil, err
		}

		friends = append(friends, *friend)
	}

	return &friends, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToFriendCreates(subfuncSDC *sub_func_complementer.SDC) (*Friend, error) {
	data := subfuncSDC.Message.Friend

	friend, err := TypeConverter[*Friend](data)
	if err != nil {
		return nil, err
	}

	return friend, nil
}

func ConvertToFriendUpdates(friendUpdates *[]dpfm_api_processing_formatter.FriendUpdates) (*[]Friend, error) {
	friends := make([]Friend, 0)

	for _, data := range *friendUpdates {
		friend, err := TypeConverter[*Friend](data)
		if err != nil {
			return nil, err
		}

		friends = append(friends, *friend)
	}

	return &friends, nil
}

func ConvertToHeader(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.Header = &sub_func_complementer.Header{
		Post:							*input.Header.Post,
		PostType:						input.Header.PostType,
		PostOwner:						input.Header.PostDate,
		PostOwnerBusinessPartnerRole:	input.Header.PostTime,
		Description:					input.Header.Description,
		LongText						input.Header.LongText,
		Tag1:                           input.Header.Tag1,
		Tag2:                           input.Header.Tag2,
		Tag3:                           input.Header.Tag3,
		Tag4:                           input.Header.Tag4,
		CreationDate:					input.Header.CreationDate,
		CreationTime:					input.Header.CreationTime,
		LastChangeDate:					input.Header.LastChangeDate,
		LastChangeTime:					input.Header.LastChangeTime,
		IsMarkedForDeletion:			input.Header.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToFriend(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var friends []sub_func_complementer.Friend

	friends = append(
		friends,
		sub_func_complementer.Friend{
			Post:           		*input.Header.Post,
			Friend:      			input.Header.Friend[0].Friend,
			CreationDate:			input.Header.Friend[0].CreationDate,
			CreationTime:			input.Header.Friend[0].CreationTime,
			LastChangeDate:			input.Header.Friend[0].LastChangeDate,
			LastChangeTime:			input.Header.Friend[0].LastChangeTime,
			IsMarkedForDeletion:	input.Header.Friend[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.Friend = &friends

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
