package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-post-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		Post:				*data.Post,
		Description:		data.Description,
		LongText			data.LongText,
		Tag1:				data.Tag1,
		Tag2:				data.Tag2,
		Tag3:				data.Tag3,
		Tag4:				data.Tag4,
		LastChangeDate:		data.LastChangeDate,
		LastChangeTime:		data.LastChangeTime,
	}
}
