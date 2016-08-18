package alexa

type Request struct {
	Version string
	Session struct {
		User struct {
			AccessToken string
		}
	}
	Request struct {
		Type, Timestamp string
		Intent          Intent
	}
}

type Intent struct {
	Name  string
	Slots map[string]struct{ Name, Value string }
}
