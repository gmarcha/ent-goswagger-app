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

type AdminSuite struct {
	suite.Suite
	client      *http.Client
	token       *models.Token
	userID      uuid.UUID
	eventTypeID uuid.UUID
	eventID     uuid.UUID
}

func (s *AdminSuite) SetupSuite() {
	s.client = &http.Client{}
	s.token = &models.Token{}

	resp, err := Get(s.client, "/auth/login")
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(s.token))
	defer resp.Body.Close()
}

func (s *AdminSuite) TestAdmin() {

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
	s.Run("SubscribeUser", s.SubscribeUser)
	s.Run("UnsubscribeUser", s.UnsubscribeUser)

	s.Run("CreateEvent", s.CreateEvent)
	s.Run("UpdateEvent", s.UpdateEvent)
	s.Run("SetEventType", s.SetEventType)
	s.Run("DeleteEvent", s.DeleteEvent)

	s.Run("CreateEventType", s.CreateEventType)
	s.Run("UpdateEventType", s.UpdateEventType)
	s.Run("DeleteEventType", s.DeleteEventType)

	s.Run("CreateUser", s.CreateUser)
	s.Run("UpdateUser", s.UpdateUser)
	s.Run("AddTutor", s.AddTutor)
	s.Run("RemoveTutor", s.RemoveTutor)
	s.Run("AddCalendar", s.AddCalendar)
	s.Run("RemoveCalendar", s.RemoveCalendar)
	s.Run("AddAdmin", s.AddAdmin)
	s.Run("RemoveAdmin", s.RemoveAdmin)
	s.Run("DeleteUser", s.DeleteUser)

	s.Run("UpdateAuthUser", s.UpdateAuthUser)
	s.Run("DeleteAuthUser", s.DeleteAuthUser)
}

func (s *AdminSuite) ListUsers() {
	resp, err := GetAuth(s.client, "/users", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(2, len(body))
}

func (s *AdminSuite) ListEvents() {
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

func (s *AdminSuite) ListEventTypes() {
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

func (s *AdminSuite) ReadAuthUser() {
	resp, err := GetAuth(s.client, "/users/me", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("admin", body.Login)
	s.userID = body.ID
}

func (s *AdminSuite) ReadUser() {
	resp, err := GetAuth(s.client, "/users/"+s.userID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("admin", body.Login)
}

func (s *AdminSuite) ListAuthUserRoles() {
	resp, err := GetAuth(s.client, "/users/me/roles", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Role{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("admin", body[0].Name)
}

func (s *AdminSuite) ListUserRoles() {
	resp, err := GetAuth(s.client, "/users/"+s.userID.String()+"/roles", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Role{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("admin", body[0].Name)
}

func (s *AdminSuite) ReadEvent() {
	resp, err := GetAuth(s.client, "/events/"+s.eventID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("test", body.Name)
}

func (s *AdminSuite) GetEventType() {
	resp, err := GetAuth(s.client, "/events/"+s.eventID.String()+"/types", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.EventType{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("test", body.Name)
	s.Require().Equal("test", body.Color)
}

func (s *AdminSuite) ReadEventType() {
	resp, err := GetAuth(s.client, "/events/types/"+s.eventTypeID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.EventType{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("test", body.Name)
	s.Require().Equal("test", body.Color)
}

func (s *AdminSuite) ListEventTypeEvents() {
	resp, err := GetAuth(s.client, "/events/types/"+s.eventTypeID.String()+"/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("test", body[0].Name)
}

func (s *AdminSuite) SubscribeAuthUser() {
	resp, err := PostAuth(s.client, "/users/me/events/"+s.eventID.String(), s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(201, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(0, len(body))
}

func (s *AdminSuite) ListAuthUserEvents() {
	resp, err := GetAuth(s.client, "/users/me/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("test", body[0].Name)
}

func (s *AdminSuite) ListUserEvents() {
	resp, err := GetAuth(s.client, "/users/"+s.userID.String()+"/events", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("test", body[0].Name)
}

func (s *AdminSuite) ListEventUsers() {
	resp, err := GetAuth(s.client, "/events/"+s.eventID.String()+"/users", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := []*ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(1, len(body))
	s.Require().Equal("admin", body[0].Login)
}

func (s *AdminSuite) UnsubscribeAuthUser() {
	resp, err := DeleteAuth(s.client, "/users/me/events/"+s.eventID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *AdminSuite) SubscribeUser() {
	resp, err := PostAuth(s.client, "/users/"+s.userID.String()+"/events/"+s.eventID.String(), s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(201, resp.StatusCode)
	body := []*ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(0, len(body))
}

func (s *AdminSuite) UnsubscribeUser() {
	resp, err := DeleteAuth(s.client, "/users/"+s.userID.String()+"/events/"+s.eventID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *AdminSuite) CreateEvent() {
	resp, err := PostAuth(s.client, "/events", s.token.AccessToken, fmt.Sprintf(
		`{"name": "exam", "startAt": "%v", "endAt": "%v"}`, time.Now().Add(time.Hour).Format(time.RFC3339), time.Now().Add(time.Hour*2).Format(time.RFC3339),
	))
	s.Require().NoError(err)
	s.Require().Equal(201, resp.StatusCode)
	body := ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("exam", body.Name)
	s.eventID = body.ID
}

func (s *AdminSuite) UpdateEvent() {
	resp, err := PatchAuth(s.client, "/events/"+s.eventID.String(), s.token.AccessToken, fmt.Sprintf(
		`{"name": "exam stud", "startAt": "%v", "endAt": "%v"}`, time.Now().Add(time.Hour).Format(time.RFC3339), time.Now().Add(time.Hour*2).Format(time.RFC3339),
	))
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("exam stud", body.Name)
}

func (s *AdminSuite) SetEventType() {
	resp, err := PatchAuth(s.client, "/events/"+s.eventID.String()+"/types/"+s.eventTypeID.String(), s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.Event{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("exam stud", body.Name)
}

func (s *AdminSuite) DeleteEvent() {
	resp, err := DeleteAuth(s.client, "/events/"+s.eventID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *AdminSuite) CreateEventType() {
	resp, err := PostAuth(s.client, "/events/types", s.token.AccessToken, `{"name": "exam", "color": "red"}`)
	s.Require().NoError(err)
	s.Require().Equal(201, resp.StatusCode)
	body := ent.EventType{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("exam", body.Name)
	s.Require().Equal("red", body.Color)
	s.eventTypeID = body.ID
}

func (s *AdminSuite) UpdateEventType() {
	resp, err := PatchAuth(s.client, "/events/types/"+s.eventTypeID.String(), s.token.AccessToken, `{"name": "exam stud", "color": "blue"}`)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.EventType{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("exam stud", body.Name)
	s.Require().Equal("blue", body.Color)
}

func (s *AdminSuite) DeleteEventType() {
	resp, err := DeleteAuth(s.client, "/events/types/"+s.eventTypeID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *AdminSuite) CreateUser() {
	resp, err := PostAuth(s.client, "/users", s.token.AccessToken, `{"login": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(201, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("melissa", body.Login)
	s.userID = body.ID
}

func (s *AdminSuite) UpdateUser() {
	resp, err := PatchAuth(s.client, "/users/"+s.userID.String(), s.token.AccessToken, `{"displayName": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("melissa", body.Login)
	s.Require().Equal("melissa", *body.DisplayName)
}

func (s *AdminSuite) AddTutor() {
	resp, err := PostAuth(s.client, "/users/"+s.userID.String()+"/role/tutor", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(201, resp.StatusCode)
	body := []*ent.Role{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(0, len(body))
}

func (s *AdminSuite) RemoveTutor() {
	resp, err := DeleteAuth(s.client, "/users/"+s.userID.String()+"/role/tutor", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *AdminSuite) AddCalendar() {
	resp, err := PostAuth(s.client, "/users/"+s.userID.String()+"/role/calendar", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(201, resp.StatusCode)
	body := []*ent.Role{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(0, len(body))
}

func (s *AdminSuite) RemoveCalendar() {
	resp, err := DeleteAuth(s.client, "/users/"+s.userID.String()+"/role/calendar", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *AdminSuite) AddAdmin() {
	resp, err := PostAuth(s.client, "/users/"+s.userID.String()+"/role/admin", s.token.AccessToken, "")
	s.Require().NoError(err)
	s.Require().Equal(201, resp.StatusCode)
	body := []*ent.Role{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal(0, len(body))
}

func (s *AdminSuite) RemoveAdmin() {
	resp, err := DeleteAuth(s.client, "/users/"+s.userID.String()+"/role/admin", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *AdminSuite) DeleteUser() {
	resp, err := DeleteAuth(s.client, "/users/"+s.userID.String(), s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}

func (s *AdminSuite) UpdateAuthUser() {
	resp, err := PatchAuth(s.client, "/users/me", s.token.AccessToken, `{"displayName": "melissa"}`)
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	body := ent.User{}
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(&body))
	defer resp.Body.Close()
	s.Require().Equal("admin", body.Login)
	s.Require().Equal("melissa", *body.DisplayName)
}

func (s *AdminSuite) DeleteAuthUser() {
	resp, err := DeleteAuth(s.client, "/users/me", s.token.AccessToken)
	s.Require().NoError(err)
	s.Require().Equal(204, resp.StatusCode)
}
