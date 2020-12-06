# [CryptoSim](https://github.com/jzlotek/bc/blob/main/README.md)

![Action Status](https://github.com/jzlotek/bc/workflows/Deploy%20to%20docker%20hub/badge.svg)

## Purpose

Project for SE 575 at Drexel University

Learn software architecting principles through creating a modular interface for a blockchain system.

## Technologies Used

### Frontend

- Vue

### Backend

- Go

### Infrastructure

- Docker
- Docker Compose
- Kubernetes
- Traefik

# Instructions
You can specify the ports, flags, and tags. These are some example commands

- From source:
  - Run: `docker-compose up --build`
  - Go to `http://localhost:8000`
- From pre-built images:
  - `docker pull jzlotek/bcfrontend:1.5.0`
  - `docker pull jzlotek/bcbackend:1.5.0`
  - `docker run -d -p 8080:80 jzlotek/bcfrontend:1.5.0`
  - `docker run -d -p 8081:8000 jzlotek/bcbackend:1.5.0`
