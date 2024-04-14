package main

import "os"

func commandExit(cfg *config, name *string) error {
	os.Exit(0)
	return nil
}
