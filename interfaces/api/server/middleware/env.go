package middleware

import "github.com/joho/godotenv"

func ExportDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
