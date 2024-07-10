package enums

type Channel string

var (
	EMAIL    Channel = "EMAIL"
	PUSH     Channel = "PUSH"
	APP      Channel = "APP"
	SMS      Channel = "SMS"
	CALL     Channel = "CALL"
	TELECALL Channel = "TELECALL"
)
