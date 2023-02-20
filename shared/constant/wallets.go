package constant

type (
	WalletStatus int

	WalletType int
)

const (
	Disabled WalletStatus = 0
	Enabled  WalletStatus = 1

	Deposit    WalletType = 1
	Withdrawal WalletType = 2
)

func (ws WalletStatus) Int() int {
	return int(ws)
}

func ToInt(s string) int {
	switch s {
	case "disabled":
		return 0
	case "enabled":
		return 1
	default:
		return 400
	}
}

func (ws WalletStatus) String() string {
	switch ws.Int() {
	case Disabled.Int():
		return "disabled"
	case Enabled.Int():
		return "enabled"
	default:
		return ""
	}
}

func (wt WalletType) Int() int {
	return int(wt)
}

func (wt WalletType) String() string {
	switch wt.Int() {
	case Deposit.Int():
		return "deposit"
	case Withdrawal.Int():
		return "withdrawal"
	default:
		return ""
	}
}
