package main

import (
	"fmt"
)

func main() {
	sub1 := Subscriber{"Bob"}
	s := Site{[]Subscriber{sub1}, []string{"Flutter Developer", "Java Developer"}}
	s.newJob("Golang Developer")
}

type Observer interface {
	react(jobs []string)
}

type Observable interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	sendAll()
}

type Site struct {
	subs []Subscriber
	jobs []string
}

type Subscriber struct {
	name string
}

func (s *Subscriber) react(jobs []string) {
	fmt.Println("Hi dear", s.name)
	fmt.Println("Vacancies updated: ")
	for _, value := range jobs {
		fmt.Println(value)
	}
}

func (s *Site) newJob(v string) {
	s.jobs = append(s.jobs, v)
	s.sendAll()
}

func (s *Site) deleteJob(v string) {
	for i, value := range s.jobs {
		if value == v {
			s.jobs = append(s.jobs[:i], s.jobs[i+1:]...)
		}
	}
	s.sendAll()
}

func (s *Site) subscribe(v Subscriber) {
	s.subs = append(s.subs, v)
}

func (s *Site) unsubscribe(v Subscriber) {
	for i, value := range s.subs {
		if value == v {
			s.subs = append(s.subs[:i], s.subs[i+1:]...)
		}
	}
}

func (s *Site) sendAll() {
	for index := range s.subs {
		s.subs[index].react(s.jobs)
	}
}
