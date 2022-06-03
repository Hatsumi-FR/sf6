package main

import (
	"github.com/sirupsen/logrus"
	"sf6/internal/store"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		ForceColors:   true,
	})
	_, err := store.NewStore()
	if err != nil {
		logrus.Fatalf("failed to load database %v", err)
	}
}
