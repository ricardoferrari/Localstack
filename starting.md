# How to initiate **Localstack** environment

Initially you would have to download Docker, then run the docker compose command:

```bash
docker-compose up -d
```

This will create a container with the Localstack image, and it will run in the background.

To check if Localstack is running correctly, you can run:

```bash
docker ps
```

You should see a container with the name `localstack_main` running. If you don't see it, make sure that Docker is running and that you have the correct permissions to run Docker commands.

To stop Localstack, you can run:

```bash
docker-compose down
```

# Environment variables

You'll need to set localstack Auth Token variabless:

LOCALSTACK_AUTH_TOKEN

# Setting AWS CLI configuration

~/.aws/config

[profile localstack]
region=us-east-1
output=json
endpoint_url = http://localhost:4566

~/.aws/credentials
[localstack]
aws_access_key_id=test
aws_secret_access_key=test

Check with:
awslocal sts get-caller-identity --profile localstack

Now it is possible to call: aws s3 ls --profile localstack

Attention: AWS CLI V2 is incompatible

Another option is to instal: pip install awscli-local[ver1]

# Go configuration

Add AWS sdk
go get -u github.com/aws/aws-sdk-go/...

## SDK Examples
https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/sqs/ReceiveMessage/ReceiveMessage.go


# Usage

To send a message call:
go run SendMessage.go -q game-update-queue -s "Teste do sender"

Te receive and delete the message use:
go run ReceiveMessage.go -q game-update-queue