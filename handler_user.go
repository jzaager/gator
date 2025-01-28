package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jzaager/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}
	userName := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	// nil err means user was found
	if err != nil {
		return fmt.Errorf("Could not find user %w", err)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User set to %q\n", userName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}

	// context.Background() creates an empty context
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("coulnd't create user: %w", err)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully!")
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:		%v", user.ID)
	fmt.Printf(" * Name:	%v", user.Name)
}
