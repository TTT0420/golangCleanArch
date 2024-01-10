package infrastructure

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			PassWord string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			PassWord string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.DB.Production.Host = "localhost"
	c.DB.Production.Username = "root"
	c.DB.Production.PassWord = "root"
	c.DB.Production.DBName = "mate"

	c.DB.Production.Host = "localhost"
	c.DB.Production.Username = "root"
	c.DB.Production.PassWord = "root"
	c.DB.Production.DBName = "mate"

	c.Routing.Port = ":8080"

	return c
}
