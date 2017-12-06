package main

import "github.com/bragfoo/saman/backend/api/task"

func main() {
	for i := 1; i <= 1; i++ {
		task.GetTotalByOffset(i)
		task.GetDailyByOffset(i)
	}
}
