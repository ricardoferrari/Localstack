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

Now it is possible to call: aws s3 ls --profile localstack