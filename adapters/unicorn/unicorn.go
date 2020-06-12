package unicorn

import (
	"github.com/mxmCherry/openrtb"
	"github.com/prebid/prebid-server/adapters"
)

type UnicornAdapter struct {
	endpoint string
}

func NewUnicornBidder(endpoint string) *UnicornAdapter {
	return &UnicornAdapter{
		endpoint,
	}
}

func (a *UnicornAdapter) MakeBids(internalRequest *openrtb.BidRequest, externalRequest *adapters.RequestData, response *adapters.ResponseData) (*adapters.BidderResponse, []error) {
	return nil, nil
}

func (a *UnicornAdapter) MakeRequests(request *openrtb.BidRequest, reqInfo *adapters.ExtraRequestInfo) ([]*adapters.RequestData, []error) {
	return nil, nil
}
