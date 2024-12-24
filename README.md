# Golang Project 🚀🚀

Welcome to my Go (Golang) project! In this project, I developed a HTTP server, which processes incoming requests and includes a graceful shutdown mechanism. The project is hosted on AWS using Elastic Beanstalk for easy scaling and management. This was a fantastic learning experience where I got hands-on with Go, AWS services, and web application deployment.

## Features

- **HTTP Server**: Handles GET requests and responds with a message.
- **Graceful Shutdown**: Proper handling of shutdown signals for a smooth server stop.
- **AWS Deployment**: Deployed on AWS Elastic Beanstalk for scalability and management.

## Technologies Used

- **Go (Golang)**: Backend programming language.
- **AWS Elastic Beanstalk**: For deployment, load balancing, and auto-scaling.
- **AWS EC2**: Virtual servers to run the application.
- **SQLite**: Managed relational database service for data persistence.
- **AWS CloudWatch**: For monitoring application logs and health checks.

## Getting Started

### Prerequisites

- **Go**: You need Go installed on your local machine. You can download it from [here](https://golang.org/dl/).
- **AWS CLI**: To interact with AWS from the command line, you'll need the AWS CLI installed. Follow the [AWS CLI installation guide](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html).
- **Elastic Beanstalk CLI**: Install the AWS Elastic Beanstalk CLI. You can follow the installation steps [here](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/eb-cli3-install.html).

### How to Clone the Project

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/your-username/golang-first-project.git
    cd golang-first-project
    ```

2. Make sure you have Go installed:

    ```bash
    go version
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Run the application locally:

    ```bash
    go run main.go -config config/local.yaml
    ```

   This should start the server locally at `localhost:8080`.

## Features

- **HTTP GET handler**: Responds to requests at `/` with a simple message.
- **Graceful Shutdown**: Handles system shutdown signals (`SIGINT`, `SIGTERM`) gracefully.
- **AWS Deployment**: Automatically scales the application based on demand.

## Troubleshooting

- **Port already in use**: If you encounter the "address already in use" error, ensure that no other process is using port 8080 by running:

    ```bash
    lsof -i :8080
    ```

    You can kill the process using:

    ```bash
    kill -9 <PID>
    ```

- **Deployment issues on AWS**: Make sure the Elastic Beanstalk CLI is correctly configured and the AWS environment is properly set up.

## Contributing

If you'd like to contribute to this project, feel free to fork the repository, submit a pull request, or open an issue for feedback. All contributions are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
