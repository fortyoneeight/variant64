package classic

type RequestNewBoard struct{}

func (r *RequestNewBoard) PerformAction() (*ClassicBoard, error) {
	return New(), nil
}
