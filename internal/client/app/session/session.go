package session

type Session struct {
	password string
	token    string
}

func (s *Session) SetPassword(p string) {
	s.password = p
}

func (s *Session) GetPassword() string {
	return s.password
}

func (s *Session) SetToken(token string) {
	s.token = token
}

func (s *Session) GetToken() string {
	return s.token
}

func NewSession() *Session {
	return &Session{}
}
