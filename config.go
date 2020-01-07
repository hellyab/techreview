package techreview

type config struct {
	DatabaseName string
	ConnString   string
}

// DBConnfigurations sturct containing db configuration
var DBConnfigurations = config{
	DatabaseName: "postgres",
	ConnString:   "postgres://postgres:Binaman1!@localhost/testdb?sslmode=disable",
}
				