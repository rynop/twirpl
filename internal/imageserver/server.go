package imageserver

import (
	"context"

	pb "github.com/rynop/twirpl/rpc/image"
	"github.com/twitchtv/twirp"
)

// Server implements the Image service
type Server struct{}

//CreateGiphy produces a Giphy from a search term
func (s *Server) CreateGiphy(ctx context.Context, search *pb.Search) (*pb.Giphy, error) {
	if search.Term == "" {
		return nil, twirp.InvalidArgumentError("term", "Cant be blank")
	}

	return &pb.Giphy{
		Url: "https://giphy.com/gifs/nhl-hockey-3o7bugwwXAUrBLflkY",
	}, nil
}
