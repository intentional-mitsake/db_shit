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
	//this attr is of type DatabaseConfig from the config.go file
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
func (p *PGClient) Connect(existing bool) error {
	//the db name is in p.config.Database, if we are connecting to this existing db, that var is fine
	//if we are creaeting it tho, we have to make sure p.config.Databse is the new db
	//postgres comes wih a default db callde 'postgres', we can use that to open a connection instead of p.config.Database
	//using the default 'postgres' db if we are creating a new db to open connection
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres?sslmode=disable",
		p.config.Username,
		p.config.Password,
		p.config.Host,
		p.config.Port,
	)
	if existing {
		//construct connection string using the config fields
		//format: postgres://username:password@host:port/dbname
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			p.config.Username,
			p.config.Password,
			p.config.Host,
			p.config.Port,
			p.config.Database,
		)

	}
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

func (p *PGClient) Create() error {
	//first we need to connect to the pg server
	c_err := p.Connect(false) //false cuz we are creating a new db
	//Connect returns an error if connection fails and nil if executed
	if c_err != nil {
		return c_err
	}
	//name of new db is the name user gave with create command
	//quote identifier to handle mixed-case and special chars safely
	_, err := p.db.Exec(fmt.Sprintf(`CREATE DATABASE "%s";`, p.config.Database))
	if err != nil {
		return err
	}
	//closes the connection once function is done
	//defer executes after func returns or exc is complete i think
	defer p.Close()
	return nil
}

func (p *PGClient) Close() error {
	if p.db == nil {
		return nil
	}
	return p.db.Close()
}

// returns a list of existing dbs and error if any
func (p *PGClient) List() ([]string, error) {
	c_err := p.Connect(false) //false cuz we r listing existing dbs and using the default 'postgres' db to connect
	if c_err != nil {
		return nil, c_err
	}
	//dataistemplate false excludes template dbs; in thss case we only want user-created dbs and aere ingoring 'postgres'
	//datname is a column in pg_database that holds the db names
	//pg_database is a system catalog table that holds info about all dbs in the pg server
	rows, err := p.db.Query(`SELECT datname FROM pg_database WHERE datistemplate = false;`)
	//above query returns *sql.Rows and error if any
	//we stored the *sql.Rows in rows var now we can do two things:
	//1. iterate over rows to get each db name
	//2. close rows once done
	if err != nil {
		return nil, err
	}
	var listDBs []string
	for rows.Next() {
		var dbName string
		err := rows.Scan(&dbName)
		if err != nil {
			return nil, err
		}
		//appending each db name to the list
		listDBs = append(listDBs, dbName)
	}
	//close to release resources
	rows.Close()
	if err != nil {
		return nil, err
	}
	defer p.Close()
	//if no errors, return the list of db names and nil error
	return listDBs, nil
}
