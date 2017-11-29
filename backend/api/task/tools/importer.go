package main

import "github.com/bragfoo/saman/backend/api/task"

func main() {
	for i := 1; i <= 7; i++ {
		task.GetDailyByOffset(i)
	}
}
