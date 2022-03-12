package eventtype

import (
	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
)

func Init(api *operations.TutorAPI, db *ent.Client) {

	typeService := &Service{Type: db.EventType}

	api.EventTypeListTypeHandler = &listType{eventType: typeService}
	api.EventTypeCreateTypeHandler = &createType{eventType: typeService}

	api.EventTypeReadTypeHandler = &readType{eventType: typeService}
	api.EventTypeUpdateTypeHandler = &updateType{eventType: typeService}
	api.EventTypeDeleteTypeHandler = &deleteType{eventType: typeService}
	api.EventTypeListTypeEventsHandler = &listTypeEvents{eventType: typeService}
}
