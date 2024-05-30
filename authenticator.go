package authn

type Authenticator interface {
	Authenticate(key *Key) (data any, err error)
}
