package enums

type Channel string

const (
	EMAIL    Channel = "EMAIL"
	PUSH     Channel = "PUSH"
	APP      Channel = "APP"
	SMS      Channel = "SMS"
	CALL     Channel = "CALL"
	TELECALL Channel = "TELECALL"
)

func (c Channel) IsValid() bool {
	switch c {
	case EMAIL, PUSH, APP, SMS, CALL, TELECALL:
		return true
	}

	return false
}
