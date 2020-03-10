package req


// Request is struct of request.
type Request struct {
	Num int
	Resp chan Response
}

// Response is struct of response.
type Response struct {
	Num int
	WorkerID int
}

// PlusOneService starts go routine with parameter reqs(channel of Request) and workerID
func PlusOneService(reqs <-chan Request, workerID int) {
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerID}
		}(req)
	}
}

