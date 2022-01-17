package pkg

type TestRequest struct {
	Expectation string `json:"expectation"`
}

type BadRequest struct {
	Expectation int `json:"expectation"`
}

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data string `json:"data"`
}
