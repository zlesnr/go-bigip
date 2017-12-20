package bigip

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

func (s *SysTestSuite) TestModifyGlobalSettings() {

	myModifedGlobalSettings := &GlobalSettings{Hostname: "asd"}

	s.Client.ModifyGlobalSettings(myModifedGlobalSettings)

	assert.Equal(s.T(), "PUT", s.LastRequest.Method)
	assert.Equal(s.T(), fmt.Sprintf("/mgmt/tm/%s/%s", uriSys, uriGlobalSettings), s.LastRequest.URL.Path)
	assert.Equal(s.T(), `{"hostname":"asd"}`, s.LastRequestBody)
}
