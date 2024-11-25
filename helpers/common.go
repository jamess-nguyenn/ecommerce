package helpers

import (
	"ecommerce/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/mod/modfile"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func init() {
	// load environment variables
	if err := LoadEnv(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	config.LoadConfigs()
}

func LoadEnv() error {
	return godotenv.Load()
}

func GetEnv() string {
	return os.Getenv("APP_ENV")
}

func GetServerHost() string {
	return os.Getenv("SERVER_HOST")
}

func GetServerPort() string {
	return os.Getenv("SERVER_PORT")
}

func GetServerAddress() string {
	address := fmt.Sprintf("%s:%s",
		GetServerHost(),
		GetServerPort(),
	)
	
	return address
}

func GetParams(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func GetTrace(skip int) (string, string) {
	pc, filePath, _, _ := runtime.Caller(skip)
	functionInfo := runtime.FuncForPC(pc).Name()

	return functionInfo, filePath
}

func GetFunctionAfterDot(functionInfo string) string {
	dotIndex := strings.LastIndex(functionInfo, ".")

	return functionInfo[dotIndex+1:]
}

func GetModuleName() (string, error) {
	file := "go.mod"

	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	modFile, err := modfile.Parse(file, data, nil)
	if err != nil {
		return "", err
	}

	return modFile.Module.Mod.Path, nil
}
