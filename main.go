package main

type Database struct {
	user string
}

type Server struct {
	db *Database // uinttr => 8 bytes long
}

func (s Server) GetUserFromDB() string {
	// golang is going to "dereference" the db pointer
	// it's going to lookup the memory address of that pointer
	if s.db == nil {
		panic("database is not initialized")
	}
	return s.db.user
}

func main() {
	s := &Server{}
	s.GetUserFromDB()
}
