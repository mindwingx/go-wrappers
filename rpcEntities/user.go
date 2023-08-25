package rpcentities

import "fmt"

type (
	UserRpc struct{}

	Args struct {
		Name string
	}

	Reply struct {
		Message string
	}
)

func (u *UserRpc) GetUser(args *Args, reply *Reply) error {
	reply.Message = fmt.Sprintf("JOOOOOOON %s", args.Name)
	return nil
}
