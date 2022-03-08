package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type StudentSuite struct {
	suite.Suite
	client *http.Client
	token  *models.Token
	// roleID      uuid.UUID
	// userID      uuid.UUID
	// eventTypeID uuid.UUID
	// eventID     uuid.UUID
}

func (s *StudentSuite) SetupSuite() {
	s.client = &http.Client{}
	s.token = &models.Token{}

	resp, err := Get(s.client, "/auth/login")
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(s.token))
	defer resp.Body.Close()
}

func (s *StudentSuite) TestStudent() {

	s.Run("ListEvents", s.ListEvents)
	s.Run("ListUsers", s.ListUsers)

	s.Run("ReadAuthUser", s.ReadAuthUser)
	s.Run("UpdateAuthUser", s.UpdateAuthUser)
	s.Run("DeleteAuthUser", s.DeleteAuthUser)

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

	s.Run("CreateUser", s.CreateUser)
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

func (s *StudentSuite) ListEvents() {
	resp, err := GetAuth(s.client, "/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("test", body[0].Name)
}

func (s *StudentSuite) ListUsers() {
	resp, err := GetAuth(s.client, "/users", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(5, len(body))
}

func (s *StudentSuite) ReadAuthUser() {
	resp, err := GetAuth(s.client, "/users/me", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("student", body.Login)
}

func (s *StudentSuite) UpdateAuthUser() {
	resp, err := PatchAuth(s.client, "/users/me", s.token.AccessToken, `{"displayName": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("student", body.Login)
	s.Require().Equal("melissa", *body.DisplayName)
}

func (s *StudentSuite) DeleteAuthUser() {
	resp, err := DeleteAuth(s.client, "/users/me", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *StudentSuite) CreateEvent() {
	resp, err := PostAuth(s.client, "/events", s.token.AccessToken, fmt.Sprintf(
		`{"name": "exam", "startAt": "%v", "endAt": "%v"}`,
		time.Now().Add(time.Minute),
		time.Now().Add(time.Minute+time.Hour)))
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ReadEvent() {
	resp, err := GetAuth(s.client, "/events/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) UpdateEvent() {
	resp, err := PatchAuth(s.client, "/events/42", s.token.AccessToken, fmt.Sprintf(
		`{"name": "exam", "startAt": "%v", "endAt": "%v"}`,
		time.Now().Add(time.Minute),
		time.Now().Add(time.Minute+time.Hour)))
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) DeleteEvent() {
	resp, err := DeleteAuth(s.client, "/events/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ListEventUsers() {
	resp, err := GetAuth(s.client, "/events/42/users", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) GetEventType() {
	resp, err := GetAuth(s.client, "/events/42/types", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) SetEventType() {
	resp, err := PatchAuth(s.client, "/events/42/types/21", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ListEventTypes() {
	resp, err := GetAuth(s.client, "/events/types", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) CreateEventType() {
	resp, err := PostAuth(s.client, "/events/types", s.token.AccessToken, `{"name": "exam", "color": "red"}`)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ReadEventType() {
	resp, err := GetAuth(s.client, "/events/types/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) UpdateEventType() {
	resp, err := PatchAuth(s.client, "/events/types/42", s.token.AccessToken, `{"name": "Exam", "color": "red"}`)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) DeleteEventType() {
	resp, err := DeleteAuth(s.client, "/events/types/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ListEventTypeEvents() {
	resp, err := GetAuth(s.client, "/events/types/42/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) CreateUser() {
	resp, err := PostAuth(s.client, "/users", s.token.AccessToken, `{"login": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ListAuthUserEvents() {
	resp, err := GetAuth(s.client, "/users/me/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ListAuthUserRoles() {
	resp, err := GetAuth(s.client, "/users/me/roles", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) SubscribeAuthUser() {
	resp, err := PostAuth(s.client, "/users/me/events/42", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) UnsubscribeAuthUser() {
	resp, err := DeleteAuth(s.client, "/users/me/events/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ReadUser() {
	resp, err := GetAuth(s.client, "/users/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) UpdateUser() {
	resp, err := PatchAuth(s.client, "/users/42", s.token.AccessToken, `{"displayName": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) DeleteUser() {
	resp, err := DeleteAuth(s.client, "/users/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ListUserEvents() {
	resp, err := GetAuth(s.client, "/users/42/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) ListUserRoles() {
	resp, err := GetAuth(s.client, "/users/42/roles", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) SubscribeUser() {
	resp, err := PostAuth(s.client, "/users/42/events/21", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) UnsubscribeUser() {
	resp, err := DeleteAuth(s.client, "/users/42/events/21", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) AddTutor() {
	resp, err := PostAuth(s.client, "/users/42/role/tutor", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) RemoveTutor() {
	resp, err := DeleteAuth(s.client, "/users/42/role/tutor", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) AddCalendar() {
	resp, err := PostAuth(s.client, "/users/42/role/calendar", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) RemoveCalendar() {
	resp, err := DeleteAuth(s.client, "/users/42/role/calendar", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) AddAdmin() {
	resp, err := PostAuth(s.client, "/users/42/role/admin", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *StudentSuite) RemoveAdmin() {
	resp, err := DeleteAuth(s.client, "/users/42/role/admin", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}
