package drivermongo

import (
	"context"
	c_ "ecs_govel/configs"
	"ecs_govel/pkg/pkglog"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"

	logecs "github.com/ecsavigne/logecs/log"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/ssh"
)

func index_test_mongo() []mongo.IndexModel {
	indexUsr := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "id", Value: 1}, // _id no se crea index why mongo create automaticaly
			},
			Options: options.Index().SetUnique(true).SetName("idx_test_mongo_id"),
		},
		{
			Keys: bson.D{
				{Key: "created_at", Value: 1},
			},
			Options: options.Index().SetName("idx_test_mongo_created_at"),
		},
	}

	return indexUsr
}

func create_index_mongo() {
	ctx := context.Background()
	_, e := mongoDB.Collection("test_mongo_migrations").
		Indexes().
		CreateMany(ctx, index_test_mongo())

	if e != nil {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivermongo.create_index_mongo", Name: "create_index_mongo", Content: map[string]any{"error": e.Error()}})
	}
}

func migrationMongoDB() {
	create_index_mongo()
}

func MongoDB() {
	clientOpts := &options.ClientOptions{}
	var (
		err       error
		sshClient *ssh.Client
	)

	// 2. Configura el túnel SSH
	if c_.SSH_ENABLE {
		sshConfig := &ssh.ClientConfig{
			User: c_.SSH_USER,
			Auth: []ssh.AuthMethod{
				ssh.Password(c_.SSH_PASS), // O usa llaves privadas
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		sshClient, err = ssh.Dial("tcp", fmt.Sprintf("%s:%s", c_.SSH_HOST, c_.SSH_PORT), sshConfig)
		if err != nil {
			pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivermongo.MongoDB", Name: "mongo_db_connect_ssh", Content: map[string]any{"error": err.Error()}})
		}
		sshDialer := &sshDialer{client: sshClient}
		clientOpts.SetDialer(sshDialer)
	}

	MONGO_DB_CONNSTR := fmt.Sprintf("mongodb://%s:%s@%s:%s/?compressors=snappy,zlib,MONGO_zstd", c_.MONGO_DB_USER, c_.MONGO_DB_PASSWORD, c_.MONGO_DB_HOST, c_.MONGO_DB_PORT)
	clientOpts = options.Client().ApplyURI(MONGO_DB_CONNSTR)

	clientMongo, err := mongo.Connect(clientOpts)
	if err != nil {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivermongo.MongoDB", Name: "mongo_db_connect_client_mongo", Content: map[string]any{"error": err.Error()}})
	}

	// err = mgm.SetDefaultConfig(nil, c_.MONGO_DB_NAME, clientOpts)
	mongoDB = clientMongo.Database(c_.MONGO_DB_NAME, options.Database())
	if err != nil {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivermongo.MongoDB", Name: "mongo_db_get_db_handle", Content: map[string]any{"error": err.Error()}})
	}

	// VERIFICACIÓN REAL: Ping al servidor
	// _, client, _, _ := mgm.DefaultConfigs()
	// err = client.Ping(mgm.Ctx(), nil)
	ctx := context.Background()
	err = clientMongo.Ping(ctx, nil)
	if err != nil {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "drivermongo.MongoDB", Name: "mongo_db_send_ping", Content: map[string]any{"error": fmt.Sprintf("%s%s%s", "\033[31m", err.Error(), "\033[0m"), "date_time": time.Now().Format("2006-01-02 15:04:05"), "type": c_.DB_TYPE}})
	} else {
		pkglog.Log.Create(logecs.InfoLog{Type: logecs.Info, Sub: "drivermongo.MongoDB", Name: "mongo_db_send_ping", Content: map[string]any{"date_time": time.Now().Format("2006-01-02 15:04:05"), "type": c_.DB_TYPE, "connection_str": MONGO_DB_CONNSTR}})

		migrationMongoDB()
	}
}

func GetDB() *mongo.Database {
	return mongoDB
}
