package handler

import (
	"log"
	"os"
	"testing"

	"github.com/iris-contrib/httpexpect"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"github.com/stretchr/testify/suite"
	"github.com/tpphu/visual-job/mock"
)

type JobHandlerTestSuite struct {
	suite.Suite
	jobRepo *mock.JobRepoImpl
	Expect  *httpexpect.Expect
}

func TestJobHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(JobHandlerTestSuite))
}

func (s *JobHandlerTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)

	app := iris.Default()
	s.jobRepo = new(mock.JobRepoImpl)
	jobHanler := &jobHandlerImpl{
		jobRepo: s.jobRepo,
		log:     logger,
	}
	jobHanler.inject(app)

	s.Expect = httptest.New(s.T(), app)
}

func (s *JobHandlerTestSuite) TearDownTest() {
}
