# go-shellcron

This program looks through a chosen folder for shell scripts (files ending in *.sh).
It finds scheduling details inside each script and runs them at the set times.
It helps to arrange tasks that need to run regularly, set up their times directly in the
scripts, and run them automatically using Go.

## Installation

Before running the program, follow these steps to set it up:

1. **Clone the repository**:
   Open your command line or terminal and run the following command to clone the Git repository:

```bash
git clone https://github.com/JexSrs/go-shellcron.git
```

2. **Navigate into the project directory**:
   After cloning, change into the project directory:

```bash
cd go-shellcron
```

3. **Install necessary Go package**:
   Install the `cron` package required for scheduling tasks:

```bash
go mod tidy
```

After these setup steps, the program will be ready for use.

## Docker

### Using Docker command

Before using Docker, ensure Docker is installed on your system. Visit [Docker's official installation guide](https://docs.docker.com/get-docker/) for help with the installation.

1. **Build the Docker Image**:
   From within the project directory (where the `Dockerfile` is located), build your Docker image:
   ```bash
   docker build -t go-shellcron .
   ```

2. **Run your Docker Container**:
   You can run your Docker container with the following command:
   ```bash
   docker run --name go-shellcron -v $(pwd)/scripts:/scripts -d go-shellcron
   ```

### Using Docker Compose

For an even simpler setup or deployment, you can utilize Docker Compose.

1. **Create a `docker-compose.yml` file**:
   Make sure you have a `docker-compose.yml` file in the root of your project with the following content:
   ```yaml
   services:
     go-shellcron:
       container_name: go-shellcron
       restart: always
       image: go-shellcron
       build:
         context: ./
         dockerfile: Dockerfile
       env_file: ".env"
       volumes:
         - ./scripts:/scripts
   ```

2. **Run using Docker Compose**:
   If you wish to run the services in the background (detached mode), use:
   ```bash
   docker-compose up -d --build
   ```
   
## Script structure
Here is a sample script to be processed by this scheduler:
```shell
#!/bin/bash
# CRON: */10 * * * *

# Echo current date and time
echo "Current date and time:"
date
```
The script above is scheduled to run every 10 minutes.

## Contributing
Contributions to this project are welcome.
Please fork the repository, make your changes, and submit a pull request for review.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
