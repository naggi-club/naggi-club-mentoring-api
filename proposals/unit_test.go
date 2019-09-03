package proposals

import (
	"testing"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naggi-club/naggi-club-mentoring-api/common"
)

func newProposal() Proposal {
	return Proposal{
		Name:                "Test",
		Email:               "test@naggi.club",
		PrograminExperience: false,
		RequestMotivation:   1,
		Content:             "よろしくおねがいします。",
	}
}

func resetDBWithMock() {
	common.TestDBClean()
	test_db := common.TestDBInit()
	newProposal := newProposal()
	AutoMigrate()
	test_db.Save(&newProposal)
}

var requestTests = []common.RequestTest{
	{
		Init: func(req *http.Request) {
			resetDBWithMock()
		},
		Url: "/api/proposals/",
		Method: "GET",
		BodyData: ``,
		ExpectedCode: http.StatusOK,
		ResponseRegexg: `{"proposals":\[.*\]`,
		Msg: "should get list of proposals",
	},
	{
		Init: func(req *http.Request) {
			resetDBWithMock()
		},
		Url: "/api/proposals/",
		Method: "POST",
		BodyData: `{"name":"テスト太郎","email":"test2@naggi.club","programming_experience":true,"request_motivation":1,"content":"宜しくおねがいします"}`,
		ExpectedCode: http.StatusCreated,
		ResponseRegexg: `{"message":"Proposal created successfully","resourceId":\d,"status":201}`,
		Msg: "should create proposal",
	},
	{
		Init: func(req *http.Request) {
			resetDBWithMock()
		},
		Url: "/api/proposals/",
		Method: "POST",
		BodyData: `{"name":"テスト太郎","email":"test@naggi.club","programming_experience":true,"request_motivation":1,"content":"宜しくおねがいします"}`,
		ExpectedCode: http.StatusUnprocessableEntity,
		ResponseRegexg: `{"errors":\["Error 1062: Duplicate entry 'test@naggi.club' for key 'email'"\],"status":422}`,
		Msg: "should not create proposal with existed email",
	},
}

func TestProposal(t *testing.T) {
	r := gin.New()
	ProposalsRegister(r.Group("/api"))

	common.ExecuteTests(t, r, requestTests)
}
