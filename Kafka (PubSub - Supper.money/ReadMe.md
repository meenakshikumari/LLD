In-memory message broker library

The library should have the notion of a Topic, Publisher, Consumer. The publisher publishes to a topic, multiple consumers can consume from a topic.
One topic can have multiple consumers and publishers.
The library should be able to manage multiple topics.
Create topics
Delete topics
Consumers should be able to consume when a message is received for the topic.
Ability to publish messages in parallel.
Graceful handling of exceptions.
Offset Management for Consumers.
Topics should have a max retention period beyond which the messages should be deleted. This is a topic level property.
[Bonus-1] Ability to reset offset and replay messages from a particular offset.
[Bonus-2] Ability to provide visibility at a consumer and topic based on last offset read or lag.


Assumptions/Clarifications -
1. Messages can be in simple string format.
2. One publisher will be publishing to one topic and one Consumer will be consuming from only one topic. However, multiple publishers can publish to a topic and multiple consumers can consume from a topic.
3. Every consumer will manage its own offset and should consume accordingly.
4. Once the retention period has passed for a message, the message should be deleted permanently and should not be consumed by any consumer after that time.
5. For point 10, we need a simple function that will get the current lag and offset for a consumer.
6. The library should allow publishing in parallel, meaning multiple publishers can send messages simultaneously to a topic without blocking each other.




