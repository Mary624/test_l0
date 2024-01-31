package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Env         string
	Port        int
	ClusterId   string
	ClientId    string
	SubjectNats string
	DurableName string
	CacheLimit  int
	DBConfig
}

type DBConfig struct {
	MigrationsPath  string
	MigrationsTable string
	UserDB          string
	PassDB          string
	HostDB          string
	PortDB          int
	DBName          string
}

const (
	envName         = "ENV"
	portName        = "PORT"
	subjectNatsName = "NATS_SUBJECT"

	clusterNatsName = "CLUSTER_ID"
	clientNatsName  = "CLIENT_ID"
	durableNatsName = "DURABLE_NAME"
	cacheLimitName  = "CACHE_LIMIT"

	migrationsPathName  = "MIGRATIONS_PATH"
	migrationsTableName = "MIGRATIONS_TABLE"
	userDBName          = "USER_DB"
	passDBName          = "PASS_DB"
	hostDBName          = "HOST_DB"
	portDBName          = "PORT_DB"
	dbName              = "DB_NAME"
)

func MustLoad() Config {
	emptyName := ""
	defer emptyNameErr(emptyName)
	env := os.Getenv(envName)
	if env == "" {
		emptyName = envName
		return Config{}
	}
	clusterID := os.Getenv(clusterNatsName)
	if clusterID == "" {
		emptyName = clusterNatsName
		return Config{}
	}
	clientID := os.Getenv(clientNatsName)
	if clientID == "" {
		emptyName = clientNatsName
		return Config{}
	}
	subjectNats := os.Getenv(subjectNatsName)
	if subjectNats == "" {
		emptyName = subjectNatsName
		return Config{}
	}
	durableName := os.Getenv(durableNatsName)
	if durableName == "" {
		emptyName = durableNatsName
		return Config{}
	}
	portStr := os.Getenv(portName)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		emptyName = portName
		return Config{}
	}
	limitCacheStr := os.Getenv(cacheLimitName)
	cacheLimit, err := strconv.Atoi(limitCacheStr)
	if err != nil {
		emptyName = cacheLimitName
		return Config{}
	}
	return Config{
		Env:         env,
		Port:        port,
		ClusterId:   clusterID,
		ClientId:    clientID,
		SubjectNats: subjectNats,
		DurableName: durableName,
		CacheLimit:  cacheLimit,
		DBConfig:    MustLoadDB(),
	}
}

func MustLoadDB() DBConfig {
	emptyName := ""
	defer emptyNameErr(emptyName)
	migrationsPath := os.Getenv(migrationsPathName)
	if migrationsPath == "" {
		emptyName = migrationsPathName
		return DBConfig{}
	}
	migrationsTable := os.Getenv(migrationsTableName)
	if migrationsTable == "" {
		emptyName = migrationsTableName
		return DBConfig{}
	}
	userDB := os.Getenv(userDBName)
	if userDB == "" {
		emptyName = userDBName
		return DBConfig{}
	}
	passDB := os.Getenv(passDBName)
	if passDB == "" {
		emptyName = passDBName
		return DBConfig{}
	}
	hostDB := os.Getenv(hostDBName)
	if hostDB == "" {
		emptyName = hostDBName
		return DBConfig{}
	}
	db := os.Getenv(dbName)
	if db == "" {
		emptyName = dbName
		return DBConfig{}
	}
	portStr := os.Getenv(portDBName)
	portDB, err := strconv.Atoi(portStr)
	if err != nil {
		emptyName = portName
		return DBConfig{}
	}

	return DBConfig{
		MigrationsPath:  migrationsPath,
		MigrationsTable: migrationsTable,
		UserDB:          userDB,
		PassDB:          passDB,
		HostDB:          hostDB,
		PortDB:          portDB,
		DBName:          db,
	}
}

func emptyNameErr(emptyName string) {
	if emptyName != "" {
		log.Fatal(fmt.Sprintf("%s is not set", emptyName))
	}
}
