package events

import "github.com/infinitybotlist/grevolt/types"

type ReportCreate struct {
	Event
	*types.Report
}
