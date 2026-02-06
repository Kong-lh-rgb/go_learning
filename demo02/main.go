package main

import "fmt"

type Exercise struct {
	Name   string
	Weight float32
	Sets   int
	Reps   int
}

type Workout struct {
	group []Exercise
}

func (w *Workout) AddExercise(e Exercise) {
	w.group = append(w.group, e)
}

func (w *Workout) PrintSummary() {
	fmt.Println("今日训练总结")
	var totalVolume float32 = 0
	for _, e := range w.group {
		fmt.Printf("动作: %s | 重量: %.1fkg | 组数: %d | 次数: %d\n", e.Name, e.Weight, e.Sets, e.Reps)
		totalVolume += e.Weight * float32(e.Sets) * float32(e.Reps)
	}
	fmt.Printf("--------------------\n")
	fmt.Printf("🔥 今日总容量: %.1f kg\n", totalVolume)
}

func main() {
	myWorkout := &Workout{}
	benchPress := Exercise{
		Name:   "平板卧推",
		Weight: 60.0,
		Sets:   4,
		Reps:   10,
	}
	fly := Exercise{
		Name:   "飞鸟",
		Weight: 10.0,
		Sets:   4,
		Reps:   12,
	}

	myWorkout.AddExercise(benchPress)
	myWorkout.AddExercise(fly)

	myWorkout.PrintSummary()
}
