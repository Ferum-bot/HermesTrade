package profitability_5_spreads

import "context"

type Worker struct {
}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) Start(ctx context.Context) error {
	return nil
}
