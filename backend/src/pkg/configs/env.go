package configs

import "github.com/joho/godotenv"

func SetupEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}
