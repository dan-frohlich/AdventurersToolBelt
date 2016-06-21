package skills

type Skill interface {
	DisplayName() string
	BasicDescription() string
	AdvancedDescription() string
	HasAdvenced() bool
	HasFocus() bool
	FocusList() []string
}

type skill struct {
	displayName         string
	basicDescription    string
	advancedDescription string
	hasAdvenced         bool
	hasFocus            bool
	focusList           []string
}

func (skl *skill) DisplayName() string { return skl.displayName }

func (skl *skill) BasicDescription() string { return skl.basicDescription }

func (skl *skill) AdvancedDescription() string { return skl.advancedDescription }

func (skl *skill) HasAdvenced() bool { return skl.hasAdvenced }

func (skl *skill) HasFocus() bool { return skl.hasFocus }

func (skl *skill) FocusList() []string { return skl.focusList }
