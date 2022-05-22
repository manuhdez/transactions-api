package bootstrap

type Dependencies struct {
	Repositories Repositories
	Services     Services
}

func Deps() Dependencies {
	dbs := bootstrapDatabases()
	repos := bootstrapRepositories(dbs)
	services := bootstrapServices(repos)

	return Dependencies{
		Repositories: repos,
		Services:     services,
	}
}
