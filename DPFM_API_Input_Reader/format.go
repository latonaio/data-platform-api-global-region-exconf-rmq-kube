package dpfm_api_input_reader

import (
	"data-platform-api-global-region-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGlobalRegion() *requests.GlobalRegion {
	data := sdc.GlobalRegion
	return &requests.GlobalRegion{
		GlobalRegion: data.GlobalRegion,
	}
}
