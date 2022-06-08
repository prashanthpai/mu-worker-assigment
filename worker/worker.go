package worker

import (
	"context"

	"github.com/Unacademy/mu-worker/pkg/db"
)

type Store interface {
	UpdateFeedback(ctx context.Context, feedback *db.Feedback) error
}

type Queue interface {
	ReceiveMessage(ctx context.Context) ([]byte, error)
}

func Run(ctx context.Context, s Store, q Queue) error {
	for {
		// run forever
		//   receive message
		//   process message i.e insert to DB

		// ensure this forever running loop can be cleanly
		// exited by sending a signal

	}
}
