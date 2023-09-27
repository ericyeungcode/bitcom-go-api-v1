package apivo

type MmpConfigVo struct {
	WindowMs       int64  `json:"window_ms"`
	FrozenPeriodMs int64  `json:"frozen_period_ms"`
	QtyLimit       string `json:"qty_limit"`
	DeltaLimit     string `json:"delta_limit"`
}

type MmpStateVo struct {
	MmpFrozenUntilMs int64 `json:"mmp_frozen_until_ms"`
	MmpFrozen        bool  `json:"mmp_frozen"`
}

type MmpItemVo struct {
	Pair      string      `json:"pair"`
	MmpConfig MmpConfigVo `json:"mmp_config"`
	MmpState  MmpStateVo  `json:"mmp_state"`
}

type MmpVo struct {
	MmpEnabled          bool        `json:"mmp_enabled"`
	MmpUserConfigurable bool        `json:"mmp_user_configurable"`
	MmpData             []MmpItemVo `json:"mmp_data"`
}
