package db 

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5430/simple_bank?sslmode=disable"
)

func testMain(m *testing.M) {
	conn, err := db.Conn(dbDriver, dbSource) 
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())

}