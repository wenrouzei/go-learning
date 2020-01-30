package main

import (
	"log"
	"time"

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

func main() {
	client := asynq.NewClient(redis)

	// Create a task with typename and payload.
	t1 := asynq.NewTask(
		"send_welcome_email",
		map[string]interface{}{"user_id": 42})

	t2 := asynq.NewTask(
		"send_reminder_email",
		map[string]interface{}{"user_id": 42})

	// Process the task immediately.
	err := client.Schedule(t1, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	// Process the task 24 hours later.
	err = client.Schedule(t2, time.Now().Add(24*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
}
