package dbconnection
import(
	"log"
	_"github.com/go-sql-driver/mysql"
	
	"database/sql"
)
var db *sql.DB
func Connection()*sql.DB{

	Db,err := sql.Open("mysql","uscolleges:Ancons927@tcp(uscolleges.c96w2gvk26br.us-west-2.rds.amazonaws.com:3306)/uscolleges")
	
	if err!=nil{
		log.Fatalln(err)
	}
	return Db
}
