package downstream_test_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/jamsman94/integration-test/ginkgo/pkg"
)

var host string

var _ = BeforeSuite(func() {
	// base test host here
	host = "http://localhost:8081"
	// get host from env if possible
	env := os.Getenv("HOST")
	if env != "" {
		host = env
	}

	fmt.Println("Downstream Test Start")
	fmt.Println("The host to be tested is:", host)
})

var _ = Describe("Downstream", func(){
	Context("Downstream Service API Test", func() {
		It("Should Success when no request body is provided", func() {
			request := &TestRequest{}
			statusCode, response, err := CallAPI("GET", host, "withoutUpstreamSvc", request)
			Expect(err).To(BeNil(), "call downstream API should success")
			Expect(statusCode).To(Equal(200), "status code should be 200.")
			Expect(response.Code).To(Equal(0), "service code should be 0")
		})

		It("Should Success when random request is given", func() {
			request := &TestRequest{
				Expectation: "random_string",
			}
			statusCode, response, err := CallAPI("GET", host, "withoutUpstreamSvc", request)
			Expect(err).To(BeNil(), "call downstream API should success")
			Expect(statusCode).To(Equal(200), "status code should be 200.")
			Expect(response.Code).To(Equal(0), "service code should be 0")
		})

		It("Should Fail when a wrong request is given", func() {
			request := &BadRequest{
				Expectation: 1,
			}
			statusCode, response, err := CallAPI("GET", host, "withoutUpstreamSvc", request)
			Expect(err).To(BeNil(), "call downstream API should success")
			Expect(statusCode).To(Equal(400), "status code should be 200.")
			Expect(response.Code).To(Equal(1), "service code should be 0")
		})

		It("Should Fail when given trigger is passed", func() {
			request := &TestRequest{
				Expectation: "downstream_error",
			}
			statusCode, response, err := CallAPI("GET", host, "withoutUpstreamSvc", request)
			Expect(err).To(BeNil(), "call downstream API should success")
			Expect(statusCode).To(Equal(400), "status code should be 200.")
			Expect(response.Code).To(Equal(2), "service code should be 0")
		})


	})
})
