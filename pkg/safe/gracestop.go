package safe

import "context"

func GraceResidentProc(ctx context.Context, f func() error) {
	go func() {
		var isDone bool
		go func() {
			select {
			case <-ctx.Done():
				isDone = true
			}
		}

	}
}
