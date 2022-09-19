package main

import "fmt"

func main() {
	bob := "Bob"
	subscribe(bob)
	newJob("Senior")
	sendAll()
	jim := "Jim"
	subscribe(jim)
	sendAll()
	unsubscribe(jim)
	sendAll()
}

var s Site

type Observer interface {
	react()
}

type Observable interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	sendAll()
}

type Site struct {
	subs []string
	jobs []string
}

func newJob(v string) {
	s.jobs = append(s.jobs, v)
	sendAll()
}

func deleteJob(index string) []string {
	for i, v := range s.jobs {
		if v == index {
			s.jobs = append(s.jobs[:i], s.jobs[i+1:]...)
		}
	}
	return s.jobs
}

func subscribe(v string) {
	s.subs = append(s.subs, v)
}

func unsubscribe(index string) []string {
	for i, v := range s.subs {
		if v == index {
			s.subs = append(s.subs[:i], s.subs[i+1:]...)
		}
	}
	return s.subs
}

func sendAll() {
	for _, value := range s.subs {
		fmt.Println("Hi dear", value)
		fmt.Println("Vacancies updated: ")
		react()
	}
}

func react() {
	for _, value := range s.jobs {
		fmt.Println(value)
	}
}
