package config

var ConnDB = connDB{}

type connDB struct {
	Driver   string
	User     string
	Password string
	Database string
	Host     string
	Port     string
	SslMode  bool
}

func (c *connDB) ConnPath() string {
	return c.Driver + "://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.Database + "?sslmode=disable"
}

func init() {
	ConnDB = getDBConn()
}

func getDBConn() connDB {
	return connDB{
		Driver:   "postgresql",
		Port:     "5432",
		User:     "olegsika",
		Database: "movie_service",
		Password: "very-secret-password",
		Host:     "db-service",
	}
}
