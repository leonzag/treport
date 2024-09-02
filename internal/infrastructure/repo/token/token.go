package token

// TODO: убать
//
//	func NewSQLiteRepo(ctx context.Context) (repo.TokenRepo, error) {
//		dbPath, err := config.PrepareSQLiteDBPathDefault()
//		if err != nil {
//			return nil, err
//		}
//		repo, err := sqlite.NewTokenRepo(dbPath)
//		if err != nil {
//			return nil, err
//		}
//		if !repo.IsInited() {
//			err = repo.Init(ctx)
//		}
//		return repo, err
//	}
