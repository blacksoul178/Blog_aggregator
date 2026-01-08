package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	Commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	value, exists := c.Commands[cmd.Name]
	if exists != true {
		return fmt.Errorf("command does not exist")
	}
	// Could have simplified here with :  return value(s, cmd)
	// but it seems cleaner in case of errors this way

	err := value(s, cmd)
	if err != nil {
		return fmt.Errorf("Error running command: %w", err)
	}

	return nil
}
