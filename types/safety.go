package types

import (
	"encoding/json"

	"github.com/infinitybotlist/grevolt/types/timestamp"
)

// UserReportReason : Reason for reporting a user
type UserReportReason string

// List of UserReportReason
const (
	NONE_SPECIFIED_UserReportReason        UserReportReason = "NoneSpecified"
	UNSOLICITED_SPAM_UserReportReason      UserReportReason = "UnsolicitedSpam"
	SPAM_ABUSE_UserReportReason            UserReportReason = "SpamAbuse"
	INAPPROPRIATE_PROFILE_UserReportReason UserReportReason = "InappropriateProfile"
	IMPERSONATION_UserReportReason         UserReportReason = "Impersonation"
	BAN_EVASION_UserReportReason           UserReportReason = "BanEvasion"
	UNDERAGE_UserReportReason              UserReportReason = "Underage"
)

// ReportStatusString : Just the status of the report
type ReportStatusString string

// List of ReportStatusString
const (
	CREATED_ReportStatusString  ReportStatusString = "Created"
	REJECTED_ReportStatusString ReportStatusString = "Rejected"
	RESOLVED_ReportStatusString ReportStatusString = "Resolved"
)

// ContentReportReason : Reason for reporting content (message or server)
type ContentReportReason string

// List of ContentReportReason
const (
	NONE_SPECIFIED_ContentReportReason      ContentReportReason = "NoneSpecified"
	ILLEGAL_ContentReportReason             ContentReportReason = "Illegal"
	ILLEGAL_GOODS_ContentReportReason       ContentReportReason = "IllegalGoods"
	ILLEGAL_EXTORTION_ContentReportReason   ContentReportReason = "IllegalExtortion"
	ILLEGAL_PORNOGRAPHY_ContentReportReason ContentReportReason = "IllegalPornography"
	ILLEGAL_HACKING_ContentReportReason     ContentReportReason = "IllegalHacking"
	EXTREME_VIOLENCE_ContentReportReason    ContentReportReason = "ExtremeViolence"
	PROMOTES_HARM_ContentReportReason       ContentReportReason = "PromotesHarm"
	UNSOLICITED_SPAM_ContentReportReason    ContentReportReason = "UnsolicitedSpam"
	RAID_ContentReportReason                ContentReportReason = "Raid"
	SPAM_ABUSE_ContentReportReason          ContentReportReason = "SpamAbuse"
	SCAMS_FRAUD_ContentReportReason         ContentReportReason = "ScamsFraud"
	MALWARE_ContentReportReason             ContentReportReason = "Malware"
	HARASSMENT_ContentReportReason          ContentReportReason = "Harassment"
)

// Snapshot of some content with required data to render
type SnapshotWithContext struct {
	// Users involved in snapshot
	Users []User `json:"_users"`
	// Channels involved in snapshot
	Channels []Channel `json:"_channels"`
	// Server involved in snapshot
	Server *Server `json:"_server,omitempty"`
	// Unique Id
	Id string `json:"_id"`
	// Report parent Id
	ReportId string `json:"report_id"`
	// Snapshot of content
	Content *SnapshotContent `json:"content"`
}

type SnapshotContent struct {
	Type string `json:"_type"`

	// If type is Message
	Message *SnapshotMessage `json:"message,omitempty"`

	// If type is Server
	Server *Server `json:"server,omitempty"`

	// If type is User
	User *User `json:"user,omitempty"`
}

// Snapshot message data
type SnapshotMessage struct {
	// Underlying message
	*Message
	// Context before the message
	PriorContext []*Message `json:"_prior_context,omitempty"`
	// Context after the message
	LeadingContext []*Message `json:"_leading_context,omitempty"`
}

// Special decoder for snapshot content
func (s *SnapshotContent) UnmarshalJSON(b []byte) error {
	// First get type
	var typ struct {
		Type string `json:"_type"`
	}

	err := json.Unmarshal(b, &typ)

	if err != nil {
		return err
	}

	s = &SnapshotContent{
		Type: typ.Type,
	}

	// Now decode based on type
	switch typ.Type {
	case "Message":
		var msg *SnapshotMessage

		err = json.Unmarshal(b, &msg)

		if err != nil {
			return err
		}

		s.Message = msg
	case "Server":
		var srv *Server

		err = json.Unmarshal(b, &srv)

		if err != nil {
			return err
		}

		s.Server = srv
	case "User":
		var usr *User

		err = json.Unmarshal(b, &usr)

		if err != nil {
			return err
		}

		s.User = usr
	}

	return nil
}

// Special function for encoding snapshot content
func (s *SnapshotContent) MarshalJSON() ([]byte, error) {
	switch s.Type {
	case "Message":
		var marshalData = map[string]any{
			"_type":   s.Type,
			"message": s.Message,
		}

		return json.Marshal(marshalData)
	case "Server":
		var marshalData = map[string]any{
			"_type":  s.Type,
			"server": s.Server,
		}

		return json.Marshal(marshalData)
	case "User":
		var marshalData = map[string]any{
			"_type": s.Type,
			"user":  s.User,
		}
		return json.Marshal(marshalData)
	}

	return nil, nil
}

// Status of the report
type ReportStatus struct {
	Status string `json:"status"`

	// If status is Rejected, this is the reason
	RejectionReason string `json:"rejection_reason,omitempty"`

	// If status is Resolved/Rejected, this is when it was closed
	ClosedAt timestamp.Timestamp `json:"closed_at,omitempty"`
}

// Additional report description
type DataReportContent struct {
	// Content being reported
	Content *DataReportContentContent `json:"content"`
	// Additional report description
	AdditionalContext string `json:"additional_context,omitempty"`
}

// The content being reported
type DataReportContentContent struct {
	// Type of content being reported, either Message, Server or User
	Type string `json:"type"`
	// ID of the message/server/user being reported
	//
	// <this is mandatory for all types>
	Id string `json:"id"`
	// Reason for reporting this message/server/user
	//
	// <this is mandatory for all types>
	ReportReason string `json:"report_reason"`
	// Message context (only is Type is Message)
	MessageId string `json:"message_id,omitempty"`
}

// User-generated platform moderation report.
type Report struct {
	// Unique Id
	Id string `json:"_id"`
	// Id of the user creating this report
	AuthorId string `json:"author_id"`
	// Reported content
	Content string `json:"content"`
	// Additional report context
	AdditionalContext string `json:"additional_context"`
	// Additional notes included on the report
	Notes string `json:"notes,omitempty"`
	// Status of the report
	Status *ReportStatus `json:"status"`
}

type DataEditReport struct {
	// New report status
	Status *ReportStatus `json:"status,omitempty"`
	// Report notes
	Notes string `json:"notes,omitempty"`
}

// -- Strikes --

// New strike information
type DataCreateStrike struct {
	// Id of reported user
	UserId string `json:"user_id"`
	// Attached reason
	Reason string `json:"reason"`
}

// New strike information
type DataEditAccountStrike struct {
	// New attached reason
	Reason string `json:"reason"`
}

// Account Strike on a user
type AccountStrike struct {
	// Strike Id
	Id string `json:"_id"`
	// Id of reported user
	UserId string `json:"user_id"`
	// Attached reason
	Reason string `json:"reason"`
}
