package main

import (
	"fmt"
	"supermoney/models"
	"supermoney/repositories"
	"supermoney/service"
	"time"
)

func main() {
	topicRepo := repositories.NewTopicRepository()
	publisherRepo := repositories.NewPublisherRepository()
	consumerRepo := repositories.NewConsumerRepository()

	topicService := service.NewTopicService(topicRepo)
	publisherService := service.NewPublisherService(topicRepo, publisherRepo)
	consumerService := service.NewConsumerService(topicRepo, consumerRepo)

	topicName := "TechTalks"
	topic, err := topicService.CreateTopic(topicName, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Topic created:", topic.Name)

	publisher1 := models.NewPublisher("Publisher1", topic)
	err = publisherRepo.AddPublisher(publisher1)
	if err != nil {
		fmt.Println(err)
		return
	}

	publisher2 := models.NewPublisher("Publisher2", topic)
	err = publisherRepo.AddPublisher(publisher2)
	if err != nil {
		fmt.Println(err)
		return
	}

	consumer1 := models.NewConsumer("Consumer1", topic)
	err = consumerRepo.AddConsumer(consumer1)
	if err != nil {
		fmt.Println(err)
		return
	}

	consumer2 := models.NewConsumer("Consumer2", topic)
	err = consumerRepo.AddConsumer(consumer2)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = topicService.AddConsumerToTopic(topicName, consumer1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = topicService.AddConsumerToTopic(topicName, consumer2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: to Fix Publish messages in parallel
	go func() {
		err := publisherService.PublishMessage("Publisher1", "Hello from Publisher1!")
		if err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		err := publisherService.PublishMessage("Publisher2", "Tech news update from Publisher2!")
		if err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		err := publisherService.PublishMessage("Publisher1", "New Golang release announced!")
		if err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(1 * time.Second)

	err = consumerService.ConsumeMessages("Consumer1")
	if err != nil {
		fmt.Println(err)
	}
	err = consumerService.ConsumeMessages("Consumer2")
	if err != nil {
		fmt.Println(err)
	}

	// Reset Consumer1's offset and replay messages
	fmt.Println("Resetting Consumer1's offset...")
	err = consumerService.ResetConsumerOffset("Consumer1", 0)
	if err != nil {
		return
	}
	err = consumerService.ConsumeMessages("Consumer1") // Consumer1 replays messages from offset 0
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Waiting for message retention cleanup...")
	time.Sleep(12 * time.Second)

	// Cleanup expired messages
	topicService.CleanupExpiredMessages()

	// Try consuming again (should not get expired messages)
	err = consumerService.ConsumeMessages("Consumer1")
	if err != nil {
		fmt.Println(err)
	}
	err = consumerService.ConsumeMessages("Consumer2")
	if err != nil {
		fmt.Println(err)
	}
}
