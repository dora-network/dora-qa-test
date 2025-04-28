# DORA Developer Test

## Usage

### Requirements
To use this project, you will need to have the following installed:

Docker 27.5.0+
Go 1.22+
Buf v1.32.1+

### Docker network

Create a docker network to allow the containers to communicate with each other.

```bash
docker network create dev-test-network
```

If you are want to call the network something else, you will need to update the .env file
to reflect the new network name.

### Starting the services required for the project

The test will test your knowledge of the following technologies:

- Kafka
- Redis

To start the services, run the following command:

```bash
make compose-up
```

To stop the services, run the following command:

```bash
make compose-down
```

### To launch the mock service run:

```bash
go run main.go
```


