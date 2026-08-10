package drivermongo

import (
	c_ "ecs_govel/configs"
	"ecs_govel/database/migration"
	"ecs_govel/pkg/pkglog"
	"fmt"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	optionsv1 "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"

	// optionsv1 "go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/ssh"
)

func index_test_mongo() []mongo.IndexModel {
	indexUsr := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "id", Value: 1}, // _id no se crea index why mongo create automaticaly
			},
			Options: optionsv1.Index().SetUnique(true).SetName("idx_test_mongo_id"),
		},
		{
			Keys: bson.D{
				{Key: "created_at", Value: 1},
			},
			Options: optionsv1.Index().SetName("idx_test_mongo_created_at"),
		},
	}

	return indexUsr
}

func create_index_mongo() {
	mgm.Coll(&migration.TestMongoMigration{}).Indexes().CreateMany(mgm.Ctx(), index_test_mongo())
}

func migrationMongoDB() {
	create_index_mongo()
}

func MongoDB() {
	clientOpts := &optionsv1.ClientOptions{}
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
			fmt.Println("Error al conectar al servidor SSH: " + err.Error())
		}
		sshDialer := &sshDialer{client: sshClient}
		clientOpts.SetDialer(sshDialer)
	}

	MONGO_DB_CONNSTR := fmt.Sprintf("mongodb://%s:%s@%s:%s/?compressors=snappy,zlib,MONGO_zstd", c_.MONGO_DB_USER, c_.MONGO_DB_PASSWORD, c_.MONGO_DB_HOST, c_.MONGO_DB_PORT)
	clientOpts = optionsv1.Client().ApplyURI(MONGO_DB_CONNSTR)

	err = mgm.SetDefaultConfig(nil, c_.MONGO_DB_NAME, clientOpts)
	if err != nil {
		fmt.Println("Error creating mgm default config: " + err.Error())
	}

	// VERIFICACIÓN REAL: Ping al servidor
	_, client, _, _ := mgm.DefaultConfigs()
	err = client.Ping(mgm.Ctx(), nil)
	if err != nil {
		pkglog.Log.Sub("configs").Errorf(`prepare_db {type: "%s", date_time: "%s", error: "%s%s%s"}%s`, c_.DB_TYPE, time.Now().Format("2006-01-02 15:04:05"), "\033[31m", err.Error(), "\033[0m", "\n")
	} else {
		pkglog.Log.Sub("configs").Infof(`prepare_db {type: "%s", date_time: "%s", connection_str: "%s"}%s`, c_.DB_TYPE, time.Now().Format("2006-01-02 15:04:05"), MONGO_DB_CONNSTR, "\n")

		migrationMongoDB()
	}
}
