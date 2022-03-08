package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuites(t *testing.T) {
	suite.Run(t, new(AuthSuite))
	suite.Run(t, new(UnauthenticatedSuite))
	suite.Run(t, new(StudentSuite))
	suite.Run(t, new(TutorSuite))
	suite.Run(t, new(CalendarSuite))
	suite.Run(t, new(AdminSuite))
}
