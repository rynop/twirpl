package blogserver

import (
	"context"

	pb "github.com/rynop/twirpl/rpc/publicservices"
	"github.com/twitchtv/twirp"
)

// Server implements the Blog service
type Server struct{}

//Subscribe creates a blog subscription for user
func (s *Server) Subscribe(ctx context.Context, user *pb.User) (*pb.User, error) {
	if user.Email == "" {
		return nil, twirp.InvalidArgumentError("email", "Cant be blank")
	}
	if user.Name == "" {
		return nil, twirp.InvalidArgumentError("name", "Cant be blank")
	}

	return &pb.User{
		Id: "aaa-bbb-ccc-dd",
	}, nil
}
