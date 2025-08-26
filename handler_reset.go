package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't clear table: %w", err)
	}

	fmt.Printf("Table reset!")
	return nil
}
