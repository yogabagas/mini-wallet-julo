package constant

type WalletStatus int

const (
	Disabled WalletStatus = 0
	Enabled  WalletStatus = 1
)

func (ws WalletStatus) Int() int {
	return int(ws)
}

func ToInt(s string) int {
	switch s {
	case "Enabled":
		return 1
	case "Disabled":
		return 0
	default:
		return 400
	}
}
