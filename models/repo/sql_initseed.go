package repo

func (sr *SqlRepository) Init() {
	if _, err := sr.Commands.Init.ExecContext(sr.Context); err != nil {
		sr.Logger.Panic("Cannot exec init command")
	}
}

func (sr *SqlRepository) Seed() {
	if _, err := sr.Commands.Seed.ExecContext(sr.Context); err != nil {
		sr.Logger.Panic("Cannot exec seed command")
	}
}
