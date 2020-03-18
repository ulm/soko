package config

import "os"

func PortDir() string {
	return getEnv("SOKO_PORT_DIR", "/mnt/packages-tree/gentoo")
}

func PostgresUser() string {
	return getEnv("SOKO_POSTGRES_USER", "root")
}

func PostgresPass() string {
	return getEnv("SOKO_POSTGRES_PASS", "root")
}

func PostgresDb() string {
	return getEnv("SOKO_POSTGRES_DB", "soko")
}

func PostgresHost() string {
	return getEnv("SOKO_POSTGRES_HOST", "db")
}

func PostgresPort() string {
	return getEnv("SOKO_POSTGRES_PORT", "5432")
}

func Debug() string {
	return getEnv("SOKO_DEBUG", "false")
}

func Quiet() string {
	return getEnv("SOKO_QUIET", "false")
}

func Version() string {
	return getEnv("SOKO_VERSION", "v0.1.0")
}

func Port() string {
	return getEnv("SOKO_PORT", "8080")
}


func getEnv(key string, fallback string) string{
	if os.Getenv("") != ""{
		return os.Getenv(key)
	}else{
		return fallback
	}
}
