package main

import (
	"context"
	"fmt"
	"time"

	"github.com/blacksoul178/Blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Username error: login requires exactly one username argument \n usage: %s <name>", cmd.Name) // added the usage 2nd line, clarifies the intended use on top of the error
	}
	name := cmd.Args[0]
	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("Couldn't set current user : %w", err)
	}
	fmt.Printf("Username set to %s\n", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Username error: register requires exactly one username argument \n usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("Error creating user: %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Error setting user: %w", err)
	} else {
		fmt.Printf("User %s created successfuly", user.Name)
	}
	return nil

}
