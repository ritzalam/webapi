package main

type Meeting struct {
	Name string
	ExtMeetingId string
	IntMeetingId string
	Duration int
	CreatedTime int64
	StartTime int64
	EndTime int64
	ForciblyEnded bool
	TelVoice string
	WebVoice string
	ModeratorPass string
	ViewerPass string
	WelcomeMsg string
	ModOnlyMessage string
	LogoutUrl string
	MaxUsers string
	Record bool
	AutoStartRecording bool
	AllowStartStopRecording bool
	DialNumber string
	DefaultAvatarURL string
	DefaultConfigToken string
	UserHasJoined bool
}