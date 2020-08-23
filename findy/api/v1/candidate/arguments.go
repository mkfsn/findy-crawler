package candidate

type Argument func(*arguments)

type arguments struct {
	// Skill ID
	skillId *int
	// Sort method, options are: newest, like
	sort *string
	// Value of cookie `_excalibur_session`
	session string
}

func WithSkillId(skillId int) Argument {
	return func(a *arguments) {
		a.skillId = &skillId
	}
}

func WithSort(sort string) Argument {
	return func(a *arguments) {
		a.sort = &sort
	}
}

func WithSession(session string) Argument {
	return func(a *arguments) {
		a.session = session
	}
}
