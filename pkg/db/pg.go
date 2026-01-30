package db

import (
	"database/sql"
	"fmt"

	"github.com/intentional_mitsake/db_shit/pkg/config"
	// side-effect import for postgres driver
	_ "github.com/lib/pq"
)

type PGClient struct {
	db *sql.DB
	//every time user enters a command, the db info is bound to viper in initConfig()
	//when we call LoadDBConfig(), that info is passed
	config config.DatabaseConfig
}

// a func that returns an instance of PGCLient struct
// the struct returned is a PGClient instance with the given config
// the config is the latest info user has entered
// this function is needed as we first must create an instance of PGCLient struct
// then use that instance to call the methods like Connect that act on it
func NewPGClient(config config.DatabaseConfig) *PGClient {
	return &PGClient{
		config: config,
	}
}

// this is a method that connects to the postgres db
// NOT a function, diff is that method has a receiver (p *PGClient) before the name and not an argument like func Connect(connStr string) error
// what this means is that this method is associated with the PGClient struct
// that is, it can be called on instances of PGClient--> p is the instance of PGClient on which this method is called
// so we can do something like:
// pgClient := &PGClient{connectionString: "your_connection_string"}
// err := pgClient.Connect()
// here pgClient is the receiver and Connect is called on it
// meaning Connect acts on obj of type PGClient and can access its fields and other methods and can only be called on PGClient instances
func (p *PGClient) Connect() error {
	//construct connection string using the config fields
	//format: postgres://username:password@host:port/dbname
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		p.config.Username,
		p.config.Password,
		p.config.Host,
		p.config.Port,
		p.config.Database,
	)
	//from go doc, the func takes driver(db type) and data source name(connStr)
	//and returns a pointer to sql.DB and error
	//the p.db is a var of type sql.DB
	// append sslmode setting for local/dev
	connStr = connStr + "?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	p.db = db

	// verify connection
	if err := p.db.Ping(); err != nil {
		return err
	}

	return nil
}

// Close closes the underlying DB connection.
func (p *PGClient) Close() error {
	if p.db == nil {
		return nil
	}
	return p.db.Close()
}
