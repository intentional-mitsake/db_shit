package db

//so interface in go are contracts that define a set of methods
// here these methods are general operations any db type needs to implement
// instead of coding separate funcs for each type of db, we instead create these general funcs for all
//basically what operations the db client can perform
type DBClient interface {
	Backup(destination string) error //takes file name returns err/nill
	Connect() error                  // takes nothing returns errr/nil
	Restore(source string) error
	Create(database string) error
	Ping() error
	List() ([]string, error) //takes nothing returns an array of strings or err
	Close() error
}
