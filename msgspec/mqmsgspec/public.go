package mqmsgspec

import "github.com/byrnedo/apibase/natsio/protobuf"

// Wrapper for protobuf generated structs to
// add setters
type GetPostsResponse struct {
	*InnerGetPostsResponse
}

func (w *GetPostsResponse) SetContext(ctx *protobuf.NatsContext) {
	w.Context = ctx
}

func NewGetPostsResponse(r *InnerGetPostsResponse) *GetPostsResponse {
	return &GetPostsResponse{r}
}

type GetPostsRequest struct {
	*InnerGetPostsRequest
}

func (w *GetPostsRequest) SetContext(ctx *protobuf.NatsContext) {
	w.Context = ctx
}

func NewGetPostsRequest(r *InnerGetPostsRequest) *GetPostsRequest {
	return &GetPostsRequest{r}
}
