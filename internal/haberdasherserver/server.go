package haberdasherserver

import (
	"context"
	"math/rand"

	pb "github.com/rynop/twirpl/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

// Server implements the Haberdasher service
type Server struct{}

//MakeHat makes a cool hat!
func (s *Server) MakeHat(ctx context.Context, size *pb.Size) (*pb.Hat, error) {
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}
	return &pb.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}
