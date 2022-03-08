package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type TutorSuite struct {
	suite.Suite
	client      *http.Client
	token       *models.Token
	userID      uuid.UUID
	eventTypeID uuid.UUID
	eventID     uuid.UUID
}

func (s *TutorSuite) SetupSuite() {
	s.client = &http.Client{}
	s.token = &models.Token{}

	resp, err := Get(s.client, "/auth/login")
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(s.token))
	defer resp.Body.Close()
}

func (s *TutorSuite) TestTutor() {

	s.Run("ListUsers", s.ListUsers)
	s.Run("ListEvents", s.ListEvents)
	s.Run("ListEventTypes", s.ListEventTypes)

	s.Run("ReadAuthUser", s.ReadAuthUser)
	s.Run("ReadUser", s.ReadUser)
	s.Run("ListAuthUserRoles", s.ListAuthUserRoles)
	s.Run("ListUserRoles", s.ListUserRoles)
	s.Run("ReadEvent", s.ReadEvent)
	s.Run("GetEventType", s.GetEventType)
	s.Run("ReadEventType", s.ReadEventType)
	s.Run("ListEventTypeEvents", s.ListEventTypeEvents)

	s.Run("SubscribeAuthUser", s.SubscribeAuthUser)
	s.Run("ListAuthUserEvents", s.ListAuthUserEvents)
	s.Run("ListUserEvents", s.ListUserEvents)
	s.Run("ListEventUsers", s.ListEventUsers)
	s.Run("UnsubscribeAuthUser", s.UnsubscribeAuthUser)
	s.Run("UpdateAuthUser", s.UpdateAuthUser)
	s.Run("DeleteAuthUser", s.DeleteAuthUser)

	s.Run("CreateEvent", s.CreateEvent)
	s.Run("UpdateEvent", s.UpdateEvent)
	s.Run("DeleteEvent", s.DeleteEvent)
	s.Run("SetEventType", s.SetEventType)

	s.Run("CreateEventType", s.CreateEventType)
	s.Run("UpdateEventType", s.UpdateEventType)
	s.Run("DeleteEventType", s.DeleteEventType)

	s.Run("CreateUser", s.CreateUser)
	s.Run("UpdateUser", s.UpdateUser)
	s.Run("DeleteUser", s.DeleteUser)
	s.Run("SubscribeUser", s.SubscribeUser)
	s.Run("UnsubscribeUser", s.UnsubscribeUser)

	s.Run("AddTutor", s.AddTutor)
	s.Run("RemoveTutor", s.RemoveTutor)
	s.Run("AddCalendar", s.AddCalendar)
	s.Run("RemoveCalendar", s.RemoveCalendar)
	s.Run("AddAdmin", s.AddAdmin)
	s.Run("RemoveAdmin", s.RemoveAdmin)

}

func (s *TutorSuite) ListUsers() {
	resp, err := GetAuth(s.client, "/users", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(4, len(body))
}

func (s *TutorSuite) ListEvents() {
	resp, err := GetAuth(s.client, "/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("test", body[0].Name)
	s.eventID = body[0].ID
}

func (s *TutorSuite) ListEventTypes() {
	resp, err := GetAuth(s.client, "/events/types", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.EventType{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("test", body[0].Name)
	s.Require().Equal("test", body[0].Color)
	s.eventTypeID = body[0].ID
}

func (s *TutorSuite) ReadAuthUser() {
	resp, err := GetAuth(s.client, "/users/me", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("tutor", body.Login)
	s.userID = body.ID
}

func (s *TutorSuite) ReadUser() {
	resp, err := GetAuth(s.client, "/users/"+s.userID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("tutor", body.Login)
}

func (s *TutorSuite) ListAuthUserRoles() {
	resp, err := GetAuth(s.client, "/users/me/roles", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Role{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("tutor", body[0].Name)
}

func (s *TutorSuite) ListUserRoles() {
	resp, err := GetAuth(s.client, "/users/"+s.userID.String()+"/roles", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Role{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("tutor", body[0].Name)
}

func (s *TutorSuite) ReadEvent() {
	resp, err := GetAuth(s.client, "/events/"+s.eventID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("test", body.Name)
}

func (s *TutorSuite) GetEventType() {
	resp, err := GetAuth(s.client, "/events/"+s.eventID.String()+"/types", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.EventType{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("test", body.Name)
	s.Require().Equal("test", body.Color)
}

func (s *TutorSuite) ReadEventType() {
	resp, err := GetAuth(s.client, "/events/types/"+s.eventTypeID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.EventType{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("test", body.Name)
	s.Require().Equal("test", body.Color)
}

func (s *TutorSuite) ListEventTypeEvents() {
	resp, err := GetAuth(s.client, "/events/types/"+s.eventTypeID.String()+"/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("test", body[0].Name)
}

func (s *TutorSuite) SubscribeAuthUser() {
	resp, err := PostAuth(s.client, "/users/me/events/"+s.eventID.String(), s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(201, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(0, len(body))
}

func (s *TutorSuite) ListAuthUserEvents() {
	resp, err := GetAuth(s.client, "/users/me/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("test", body[0].Name)
}

func (s *TutorSuite) ListUserEvents() {
	resp, err := GetAuth(s.client, "/users/"+s.userID.String()+"/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("test", body[0].Name)
}

func (s *TutorSuite) ListEventUsers() {
	resp, err := GetAuth(s.client, "/events/"+s.eventID.String()+"/users", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("tutor", body[0].Login)
}

func (s *TutorSuite) UnsubscribeAuthUser() {
	resp, err := DeleteAuth(s.client, "/users/me/events/"+s.eventID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *TutorSuite) UpdateAuthUser() {
	resp, err := PatchAuth(s.client, "/users/me", s.token.AccessToken, `{"displayName": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("tutor", body.Login)
	s.Require().Equal("melissa", *body.DisplayName)
}

func (s *TutorSuite) DeleteAuthUser() {
	resp, err := DeleteAuth(s.client, "/users/me", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *TutorSuite) CreateEvent() {
	resp, err := PostAuth(s.client, "/events", s.token.AccessToken, fmt.Sprintf(
		`{"name": "exam", "startAt": "%v", "endAt": "%v"}`,
		time.Now().Add(time.Minute),
		time.Now().Add(time.Minute+time.Hour)))
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) UpdateEvent() {
	resp, err := PatchAuth(s.client, "/events/42", s.token.AccessToken, fmt.Sprintf(
		`{"name": "exam", "startAt": "%v", "endAt": "%v"}`,
		time.Now().Add(time.Minute),
		time.Now().Add(time.Minute+time.Hour)))
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) DeleteEvent() {
	resp, err := DeleteAuth(s.client, "/events/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) SetEventType() {
	resp, err := PatchAuth(s.client, "/events/42/types/21", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) CreateEventType() {
	resp, err := PostAuth(s.client, "/events/types", s.token.AccessToken, `{"name": "exam", "color": "red"}`)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) UpdateEventType() {
	resp, err := PatchAuth(s.client, "/events/types/42", s.token.AccessToken, `{"name": "Exam", "color": "red"}`)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) DeleteEventType() {
	resp, err := DeleteAuth(s.client, "/events/types/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) CreateUser() {
	resp, err := PostAuth(s.client, "/users", s.token.AccessToken, `{"login": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) UpdateUser() {
	resp, err := PatchAuth(s.client, "/users/42", s.token.AccessToken, `{"displayName": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) DeleteUser() {
	resp, err := DeleteAuth(s.client, "/users/42", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) SubscribeUser() {
	resp, err := PostAuth(s.client, "/users/42/events/21", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) UnsubscribeUser() {
	resp, err := DeleteAuth(s.client, "/users/42/events/21", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) AddTutor() {
	resp, err := PostAuth(s.client, "/users/42/role/tutor", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) RemoveTutor() {
	resp, err := DeleteAuth(s.client, "/users/42/role/tutor", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) AddCalendar() {
	resp, err := PostAuth(s.client, "/users/42/role/calendar", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) RemoveCalendar() {
	resp, err := DeleteAuth(s.client, "/users/42/role/calendar", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) AddAdmin() {
	resp, err := PostAuth(s.client, "/users/42/role/admin", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}

func (s *TutorSuite) RemoveAdmin() {
	resp, err := DeleteAuth(s.client, "/users/42/role/admin", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(403, resp.StatusCode)
}
