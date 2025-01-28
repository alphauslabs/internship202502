package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type contents struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

var ps = NewPubSub()

func publishHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var c contents
	if err := json.Unmarshal(body, &c); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	if c.Topic == "" || c.Message == "" {
		http.Error(w, "Topic and message are required", http.StatusBadRequest)
		return
	}
	ps.Publish(c.Topic, c.Message)
	w.WriteHeader(http.StatusOK)
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	if topic == "" {
		http.Error(w, "Topic is required", http.StatusBadRequest)
		return
	}
	ch := make(chan string)
	ps.Subscribe(topic, ch)

	for msg := range ch {
		_, err := w.Write([]byte(msg + "\n"))
		if err != nil {
			return
		}
		w.(http.Flusher).Flush()
	}
}

type PubSub struct {
	topics map[string][]chan string // key: topic, val: list of subscribers
	mu     sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		topics: make(map[string][]chan string),
	}
}

func (ps *PubSub) Publish(topic string, message string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	if subscribers, ok := ps.topics[topic]; ok {
		for _, ch := range subscribers {
			ch <- message
		}
	}
}

func (ps *PubSub) Subscribe(topic string, ch chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.topics[topic] = append(ps.topics[topic], ch)
}

func main() {
	http.HandleFunc("/publish", publishHandler)
	http.HandleFunc("/subscribe", subscribeHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
