# Distributed Train Management System

## Overview

This project implements a distributed microservices-based system for managing trains, schedules, tickets, users, and notifications.  
Each service is independent and communicates through REST APIs. Some components also use MQTT for asynchronous notifications.

A web frontend is included to allow both administrators and customers to interact with the system.

---

## Architecture

The system is composed of the following services:

### users-service
Handles user registration and stores user information.

### auth-service
Manages user authentication and JWT token generation.

### trains-service
Manages trains, schedules, and related data.

### tickets-service
Handles ticket purchases and ticket management.

### notifications-service
Manages user subscriptions to trains and sends notifications through MQTT.

### database
MySQL database used by all services.

### server_mosquitto
MQTT broker used for notifications.

### frontend
Web interface for interacting with the system.

---

## Requirements

- Docker
- Docker Compose

No additional dependencies are required.

---

## How to Run

To start the entire system, move to the folder /project-dpwims and run:
```bash
docker compose up -–build
```

Docker will:

- build all service images
- start the containers
- initialize the databases
- automatically create two demo users (admin and customer)
- start the frontend

The system is ready when all containers are running.

---

## Demo Users

Two users are automatically created at startup:

### Admin
- Email: admin@example.com
- Password: Password123!
- Role: admin

### Customer
- Email: customer@example.com
- Password: Password123!
- Role: customer

These accounts can be used to test the system through the frontend.

---

## Frontend Access

Once the system is running, the frontend is available at:


Docker will:

- build all service images
- start the containers
- initialize the databases
- automatically create two demo users (admin and customer)
- start the frontend

The system is ready when all containers are running.

---

## Demo Users

Two users are automatically created at startup:

### Admin
- Email: admin@example.com
- Password: Password123!
- Role: admin

### Customer
- Email: customer@example.com
- Password: Password123!
- Role: customer

These accounts can be used to test the system through the frontend.

---

## Login Access

Once the system is running, the frontend is available at:

http://localhost:5173


From the frontend, it is possible to:

- log in as admin or customer
- manage users, trains, stations, tickets, payments and schedules as admin
- purchase tickets, manage subscriptions and receive notifications as customer 

---

## Stopping the System

To stop all services:
```bash
docker compose down -v
```

---

## Notes

- The system is fully containerized and requires no manual setup.
- Demo users are created automatically at startup through a seed script.
- The project is intended for demonstration and testing purposes.

