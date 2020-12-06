# [CryptoSim](https://github.com/jzlotek/bc/README.md)

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
- Kubernetes

# Instructions
You can specify the ports, flags, and tags. These are some example commands
 - `docker pull jzlotek/bcfrontend:1.4.1`
 - `docker pull jzlotek/bcbackend:1.4.1`
 - `docker run -d -p 8080:80 jzlotek/bcfrontend:1.4.1`
 - `docker run -d -p 8081:8000 jzlotek/bcbackend:1.4.1`