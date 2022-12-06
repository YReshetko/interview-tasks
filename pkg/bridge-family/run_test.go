package bridge_family_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	bridge_family "github.com/YReshetko/interview-tasks/pkg/bridge-family"
)

func TestRun(t *testing.T) {
	res := bridge_family.Run()
	assert.Equal(t, "[ПМБР] | [ПМБР]-> [ПМ](2) ->[] | [БР]<- [М](2) <-[ПМ] | [БРМ]-> [БР](10) ->[П] | [М]<- [П](1) <-[ПБР] | [МП]-> [МП](2) ->[БР] | [БРМП] || = 17", res)
}
