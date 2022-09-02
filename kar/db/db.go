package db

import (
	"helo/models"
	"log"

	// "github.com/tutorials/go/crud/pkg/models"
	//"github.com/ThotPrime/Project/tree/master/Project/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitP() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}

/*
func intiMongo() *mongo.collection {

	const connectionString = "mongodb+srv://Umerchu:123@cluster0.ubdz78d.mongodb.net/?retryWrites=true&w=majority"
	const dbName = "netflix"
	const colName = "watchlist"

	// MOST IMPORTANT
	var collection *mongo.Collection

	// connect with monogoDB

	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
	return collection
}
*/
