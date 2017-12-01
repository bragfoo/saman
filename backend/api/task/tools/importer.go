package main

import "github.com/bragfoo/saman/backend/api/task"

func main() {
	for i := 1; i <= 1; i++ {
		o := i - 1
		task.GetTotalByOffset(o)
		task.GetDailyByOffset(i)
	}
}
