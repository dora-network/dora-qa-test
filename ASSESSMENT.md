# QA Challenge: Prices Tick Service

## Overview

This repository contains a simple implementation of a service whose specifications are laid out below. 
Your task is to design and implement an end-to-end test for the service,and find out what issues there
are with the service. The service is a simple gRPC server that consumes price ticks from a Kafka topic
and persists them to a data store. The service also exposes a gRPC API to fetch the persisted tick data.

## Business requirements

The trading platform matches orders from buyers and sellers. When a trade is executed, the price at which
the trade was executed is called a "tick". The trading platform needs to persist these ticks for later
analysis. The data should include the asset ID, the price, volume and the timestamp of the tick. 
Additional the best bid and best ask prices at the time of the tick should be included.

When a trade is executed, the tick data should be written to a message queue and persisted to a data store.
The data store should be able to handle high throughput and low latency. The service should be able to 
scale horizontally to handle increased load.

An API should be provided to fetch the persisted tick data. The API should support filtering by asset ID and time range.
The API should also support pagination to limit the number of results returned.

## Task

The mock service
simulates the trading platform and generates ticks that are persisted to a data store, and provides an API to fetch the persisted tick data.

Discuss the requirements with the development team and design a test to ensure the service meets the business. 

### Considerations

This is a paired programming exercise, so you must share your screen.
You can use whatever language and framework you are most comfortable with.
You must disable any AI code generation tools in your IDE. Code completion for variable names, and function signatures etc., is allowed.
You may use Google to search for information, about any libraries you are not familiar with, e.g. Kafka, Redis, but nothing else.
You may ask the team any questions you may have about the requirements, and the implementation of the service.


## Evaluation Criteria

We will be primarily analyzing you approach to the problem, how you communicate your thought process and the overall design of your test. For example:

- What types of problems do you think you might encounter?
- How comprehensive is your test?
- What is your thought process when you're asked to design a test?
- Are you asking relevant questions to complete the task?
- Do you identify edge cases?
- Can you bring additional value to the task, e.g. identify areas for improvement in the code, suggest additional tests, etc.?

We will also be looking at your coding style, and how you structure your code. For example:

- Is your code easy to read and understand?
- Do you document your code where appropriate?
- How reusable is your code?
- How familiar are you with your chosen tools and libraries, do you use them effectively to complete the task at hand?


## Submission
- Fork this repository and work on your fork.
- Once completed, create a pull request with your changes.

## Notes
- The codebase is generic and uses public libraries, so it should be straightforward to work with.
- If you have any questions or need clarification, feel free to reach out.

Good luck, and we look forward to seeing your solution!

