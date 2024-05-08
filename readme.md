# Ormuco Technical Test Solution

This monorepo contains the solution for the technical test. The solution is divided into a frontend and backend. For the frontend, I used Angular, for the backend I decided to use Golang and to persist data I used Redis.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

To execute the solution, follow these steps:

1. Create a new folder on your computer and download the ZIP file from this repository into that folder. Extract the contents of the ZIP file.

2. Open your terminal and navigate to the project's root directory, where the `docker-compose.yaml` file is located.

3. Run the following command to start the Docker containers:

```bash
docker compose up -d
```

4. Once the containers are running, you can access the application:

   - Frontend: [http://localhost:4200](http://localhost)
   - Backend: [http://localhost:8080](http://localhost:8080) rest api
   - Backend websocket: [ws://localhost:8081/ws](ws://localhost:8081/ws) web socket

The Angular frontend will be available at `http://localhost`, and the Golang backend will be accessible at `http://localhost:8080`.

If you want to access the backend API directly, you can use tools like cURL or Postman to send HTTP requests to `http://localhost:8080/api/v1`.

if you want to see all the endpoint and documenation of the backend you can access to `http://localhost:8080/api/v1/swagger`


## Project Structure

- **frontend/**: Contains the Angular frontend application.
  - **src/app/**: The main source code for the Angular application.
  - **src/assets/**: Static assets like images, fonts, etc.
  - **src/environments/**: Environment configuration files.
  - **angular.json**: Angular CLI configuration file.
  - **package.json**: npm package manager file.

- **backend/**: Contains the Golang backend application.
  - **cmd**: The entry point of the Golang application
  - **build/**: Contains the binary to execute the application.
  - **config/**: Contains the configuration files that use our web api.
  - **docs/**: Contains the file for the docs our api.
  - **internal/**: Private application and library code.
  - **pkg/**: Library code that is ok to use by external application in our case we use to struct our model.
  - **test/**: Additional external test apps and test data
  - **utils/**: Utility functions and helpers.
  - **go.mod**: Go module definition file.
  - **go.sum**: Go module checksums file.
  - **images/**.png: Contains the images attached to this readme.md
  - **ngninx/**: Contains the file to set up a Dockerizer nginx web server
  - **redis/**: Contains the files to set up a Dockerizer redis database
  - docker-compose.yaml: file to set up our application


The infrastructure design 
![infract design](images/aws-infrastructure.png)



## System Design

The system architecture is designed to facilitate real-time communication and efficient handling of Least Recently Used (LRU) data. Below is an overview of the key components and their interactions:

### Client Interaction

- Users interact with the system to create LRU data entries via the frontend interface.
- These entries are stored in the backend for future reference.

### WebSocket Implementation

- WebSocket protocol is employed to enable real-time communication between the backend server and connected clients.
- Clients establish WebSocket connections with the server to receive updates and notifications regarding LRU data changes.

### LRU Data Expiry Notification

- Redis serves as the data store for LRU entries and monitors the expiration of data entries.
- When a data entry expires, Redis triggers a notification event.

### Server Handling of Redis Events

- The Golang backend server is responsible for handling Redis notification events.
- Upon receiving a notification about an expired LRU data entry, the server broadcasts this information to all connected WebSocket clients.

### Client Response to Expiry Events

- Clients subscribed to WebSocket channels relevant to the expired data entry receive notifications in real-time.
- Upon receiving a notification, clients update their local state or user interface to reflect the expired data.

### Benefits of the Architecture

- Efficient handling of data expiry events.
- Seamless updates across all connected clients.
- Real-time synchronization between backend and frontend components.

This architecture ensures that clients stay synchronized with changes to LRU data, providing a responsive and interactive user experience.



## screenshot exercise 1
![image 1](images/1.png)


## screenshot exercise 2
![image 2](images/2.png)


## screenshot exercise 3
![image 3](images/3.png)









