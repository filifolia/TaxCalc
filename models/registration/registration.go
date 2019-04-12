package registration

func init(){
	orm.RegisterDriver("postgres", orm.DRPostgres)

	cfg := config.Config.PgCfg
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?connect_timeout=%s&sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Database, strconv.Itoa(cfg.Timeout))
	log.Info("Connecting to ", connString)

	orm.DefaultTimeLoc = time.UTC

	if config.Environment == "dev" || config.Environment == "test" {
		orm.Debug = true
	}

	orm.RegisterDataBase("default", "postgres", connString, cfg.MaxIdleConn, cfg.MaxConn)

	orm.RegisterModel(new(models.SalesOrder))
}
