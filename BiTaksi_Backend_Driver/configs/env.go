package configs

import (
	"BiTaksi_Backend_Driver/tools/errors"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)

	errors.ServerErrorWithErrorLog(err)

	errLoad := godotenv.Load(dir + "\\.env")
	errors.ServerErrorWithErrorLog(errLoad)

	return os.Getenv("MONGOURI")
}
