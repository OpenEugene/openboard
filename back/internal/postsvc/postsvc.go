package postsvc

import (
	"github.com/OpenEugene/openboard/back/internal/pb"
)

var _ pb.PostServer = &PostSvc{}

// PostSvc encapsulates dependencies and data required to implement the
// pb.PostServer interface.
type PostSvc struct{}

// New returns a pointer to a PostSvc instance or an error.
func New() (*PostSvc, error) {
	return &PostSvc{}, nil
}

/*type PostServer interface {
	AddType(context.Context, *AddTypeReq) (*TypeResp, error)
	AddPost(context.Context, *AddPostReq) (*PostResp, error)
	FndPosts(context.Context, *FndPostsReq) (*PostsResp, error)
	OvrPost(context.Context, *OvrPostReq) (*PostResp, error)
	RmvPost(context.Context, *RmvPostReq) (*RmvPostResp, error)
}*/
