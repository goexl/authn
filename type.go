package authn

const (
	TypeUsername Type = "username"
	TypeMobile   Type = "mobile"
	TypeEmail    Type = "email"
	TypeQQ       Type = "qq"
	TypeWechat   Type = "wechat"
)

type Type string
