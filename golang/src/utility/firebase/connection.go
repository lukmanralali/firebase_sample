package config
 
import (
	"context"
	"log"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)
 
func Connection() (context.Context, *db.Client) {
	ctx := context.Background()
	conf := &firebase.Config{
        DatabaseURL: "https://sample-project-267411.firebaseio.com",
	}
	opt := option.WithCredentialsFile("utility/firebase/credential.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
        log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
        log.Fatalln("Error initializing database client:", err)
	}
	return ctx, client
}