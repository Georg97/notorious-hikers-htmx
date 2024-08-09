package env

import (
    "os"
    "log"
    "strconv"

    "github.com/joho/godotenv"
)

var (
    Port int = 3000
    Environment string = "dev"
    initialized bool = false
)

func InitiateEnv() error {
    if initialized {
        log.Println("Tried to Load Env, when it was already initialized")
        return nil
    }

    env := os.Getenv("TARGET_ENV")
    if env != "" {
        Environment = env
    }

    err := godotenv.Load(".env." + Environment)
    if err != nil {
        log.Println("Error loading .env file. Using default values")
        return err
    }

    port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
    if err != nil {
        log.Printf("Given port '%s' was not a number", port)
    } else {
        Port = port
    }

    return nil
}
