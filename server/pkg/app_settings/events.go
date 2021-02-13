package app_settings

import (
	"encoding/json"

	"github.com/freshteapot/learnalist-api/server/pkg/event"
	"github.com/freshteapot/learnalist-api/server/pkg/openapi"
	"github.com/freshteapot/learnalist-api/server/pkg/spaced_repetition/dripfeed"
)

func (s AppSettingsService) OnEvent(entry event.Eventlog) {
	switch entry.Kind {
	case dripfeed.EventDripfeedAdded:
		b, _ := json.Marshal(entry.Data)
		var moment openapi.SpacedRepetitionOvertimeInfo
		json.Unmarshal(b, &moment)
		AppendAndSaveSpacedRepetition(s.userRepo, moment.UserUuid, moment.AlistUuid)
	case dripfeed.EventDripfeedRemoved:
		fallthrough
	case dripfeed.EventDripfeedFinished:
		b, _ := json.Marshal(entry.Data)
		var moment openapi.SpacedRepetitionOvertimeInfo
		json.Unmarshal(b, &moment)
		RemoveAndSaveSpacedRepetition(s.userRepo, moment.UserUuid, moment.AlistUuid)
	}
}
