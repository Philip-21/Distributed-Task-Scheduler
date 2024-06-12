package common

import (
	"time"

	"google.golang.org/protobuf/types/known/anypb"
)

//Define the Task struct in the common package, which can be used by both scheduler and worker. To avoid import cycle errors
type Task struct {
	ID        string
	Payload   *anypb.Any
	ExecuteAt time.Time
}