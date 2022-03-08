package integration

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stretchr/testify/suite"
)

type UnauthenticatedSuite struct {
	suite.Suite
	client *http.Client
}

func (s *UnauthenticatedSuite) SetupSuite() {
	s.client = &http.Client{}
}

func (s *UnauthenticatedSuite) TestUnauthenticated() {

	s.Run("ListEvents", s.ListEvents)
	s.Run("CreateEvent", s.CreateEvent)
	s.Run("ReadEvent", s.ReadEvent)
	s.Run("UpdateEvent", s.UpdateEvent)
	s.Run("DeleteEvent", s.DeleteEvent)
	s.Run("ListEventUsers", s.ListEventUsers)
	s.Run("GetEventType", s.GetEventType)
	s.Run("SetEventType", s.SetEventType)

	s.Run("ListEventTypes", s.ListEventTypes)
	s.Run("CreateEventType", s.CreateEventType)
	s.Run("ReadEventType", s.ReadEventType)
	s.Run("UpdateEventType", s.UpdateEventType)
	s.Run("DeleteEventType", s.DeleteEventType)
	s.Run("ListEventTypeEvents", s.ListEventTypeEvents)

	s.Run("ListUsers", s.ListUsers)
	s.Run("CreateUser", s.CreateUser)
	s.Run("ReadAuthUser", s.ReadAuthUser)
	s.Run("UpdateAuthUser", s.UpdateAuthUser)
	s.Run("DeleteAuthUser", s.DeleteAuthUser)
	s.Run("ListAuthUserEvents", s.ListAuthUserEvents)
	s.Run("ListAuthUserRoles", s.ListAuthUserRoles)
	s.Run("SubscribeAuthUser", s.SubscribeAuthUser)
	s.Run("UnsubscribeAuthUser", s.UnsubscribeAuthUser)
	s.Run("ReadUser", s.ReadUser)
	s.Run("UpdateUser", s.UpdateUser)
	s.Run("DeleteUser", s.DeleteUser)
	s.Run("ListUserEvents", s.ListUserEvents)
	s.Run("ListUserRoles", s.ListUserRoles)
	s.Run("SubscribeUser", s.SubscribeUser)
	s.Run("UnsubscribeUser", s.UnsubscribeUser)

	s.Run("AddTutor", s.AddTutor)
	s.Run("RemoveTutor", s.RemoveTutor)
	s.Run("AddCalendar", s.AddCalendar)
	s.Run("RemoveCalendar", s.RemoveCalendar)
	s.Run("AddAdmin", s.AddAdmin)
	s.Run("RemoveAdmin", s.RemoveAdmin)
}

func (s *UnauthenticatedSuite) ListEvents() {
	resp, err := GetAuth(s.client, "/events", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) CreateEvent() {
	resp, err := PostAuth(s.client, "/events", "", fmt.Sprintf(
		`{"name": "exam", "startAt": "%v", "endAt": "%v"}`,
		time.Now().Add(time.Minute),
		time.Now().Add(time.Minute+time.Hour)))
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ReadEvent() {
	resp, err := GetAuth(s.client, "/events/42", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) UpdateEvent() {
	resp, err := PatchAuth(s.client, "/events/42", "", fmt.Sprintf(
		`{"name": "exam", "startAt": "%v", "endAt": "%v"}`,
		time.Now().Add(time.Minute),
		time.Now().Add(time.Minute+time.Hour)))
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) DeleteEvent() {
	resp, err := DeleteAuth(s.client, "/events/42", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ListEventUsers() {
	resp, err := GetAuth(s.client, "/events/42/users", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) GetEventType() {
	resp, err := GetAuth(s.client, "/events/42/types", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) SetEventType() {
	resp, err := PatchAuth(s.client, "/events/42/types/21", "", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ListEventTypes() {
	resp, err := GetAuth(s.client, "/events/types", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) CreateEventType() {
	resp, err := PostAuth(s.client, "/events/types", "", `{"name": "exam", "color": "red"}`)
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ReadEventType() {
	resp, err := GetAuth(s.client, "/events/types/42", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) UpdateEventType() {
	resp, err := PatchAuth(s.client, "/events/types/42", "", `{"name": "Exam", "color": "red"}`)
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) DeleteEventType() {
	resp, err := DeleteAuth(s.client, "/events/types/42", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ListEventTypeEvents() {
	resp, err := GetAuth(s.client, "/events/types/42/events", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ListUsers() {
	resp, err := GetAuth(s.client, "/users", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) CreateUser() {
	resp, err := PostAuth(s.client, "/users", "", `{"login": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ReadAuthUser() {
	resp, err := GetAuth(s.client, "/users/me", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) UpdateAuthUser() {
	resp, err := PatchAuth(s.client, "/users/me", "", `{"displayName": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) DeleteAuthUser() {
	resp, err := DeleteAuth(s.client, "/users/me", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ListAuthUserEvents() {
	resp, err := GetAuth(s.client, "/users/me/events", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ListAuthUserRoles() {
	resp, err := GetAuth(s.client, "/users/me/roles", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) SubscribeAuthUser() {
	resp, err := PostAuth(s.client, "/users/me/events/42", "", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) UnsubscribeAuthUser() {
	resp, err := DeleteAuth(s.client, "/users/me/events/42", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ReadUser() {
	resp, err := GetAuth(s.client, "/users/42", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) UpdateUser() {
	resp, err := PatchAuth(s.client, "/users/42", "", `{"displayName": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) DeleteUser() {
	resp, err := DeleteAuth(s.client, "/users/42", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ListUserEvents() {
	resp, err := GetAuth(s.client, "/users/42/events", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) ListUserRoles() {
	resp, err := GetAuth(s.client, "/users/42/roles", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) SubscribeUser() {
	resp, err := PostAuth(s.client, "/users/42/events/21", "", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) UnsubscribeUser() {
	resp, err := DeleteAuth(s.client, "/users/42/events/21", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) AddTutor() {
	resp, err := PostAuth(s.client, "/users/42/role/tutor", "", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) RemoveTutor() {
	resp, err := DeleteAuth(s.client, "/users/42/role/tutor", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) AddCalendar() {
	resp, err := PostAuth(s.client, "/users/42/role/calendar", "", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) RemoveCalendar() {
	resp, err := DeleteAuth(s.client, "/users/42/role/calendar", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) AddAdmin() {
	resp, err := PostAuth(s.client, "/users/42/role/admin", "", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}

func (s *UnauthenticatedSuite) RemoveAdmin() {
	resp, err := DeleteAuth(s.client, "/users/42/role/admin", "")
	s.Require().NoError(err)
	s.Require().Equal(401, resp.StatusCode)
}
