package main

import (
	"fmt"

	"github.com/hibiken/asynq"
)

var redis = &asynq.RedisClientOpt{
	Addr: "localhost:6379",
	// Omit if no password is required
	Password: "",
	// Use a dedicated db number for asynq.
	// By default, Redis offers 16 databases (0..15)
	DB: 0,
}

// Dispatcher is used to dispatch tasks to registered handlers.
type Dispatcher struct {
	mapping map[string]asynq.HandlerFunc
}

// HandleFunc registers a task handler
func (d *Dispatcher) HandleFunc(taskType string, fn asynq.HandlerFunc) {
	d.mapping[taskType] = fn
}

// ProcessTask processes a task.
//
// NOTE: Dispatcher satisfies asynq.Handler interface.
func (d *Dispatcher) ProcessTask(task *asynq.Task) error {
	fn, ok := d.mapping[task.Type]
	if !ok {
		return fmt.Errorf("no handler registered for %q", task.Type)
	}
	return fn(task)
}

func main() {
	d := &Dispatcher{mapping: make(map[string]asynq.HandlerFunc)}
	d.HandleFunc("send_welcome_email", sendWelcomeEmail)
	d.HandleFunc("send_reminder_email", sendReminderEmail)

	bg := asynq.NewBackground(redis, &asynq.Config{
		Concurrency: 10,
	})
	bg.Run(d)
}

func sendWelcomeEmail(t *asynq.Task) error {
	id, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	fmt.Printf("Send Welcome Email to User %d\n", id)
	return nil
}

func sendReminderEmail(t *asynq.Task) error {
	id, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	fmt.Printf("Send Welcome Email to User %d\n", id)
	return nil
}
