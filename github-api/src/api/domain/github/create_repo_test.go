package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoTestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang intro",
		Description: "a golang introduction repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t,
		`{"name":"golang intro","description":"a golang introduction repo","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":false,"has_wiki":true}`,
		string(bytes))

}
