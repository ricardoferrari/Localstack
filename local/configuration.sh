export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="us-east-1"

aws configure set region us-east-1 --profile localstack
aws configure set output json --profile localstack
aws configure set aws_access_key_id test --profile localstack
aws configure set aws_secret_access_key test --profile localstack
aws configure set aws_session_token ls-PIcibAWa-wIlU-DUce-roRO-8559qEGU159b --profile localstack

aws sns create-topic --name failed-resize-topic --profile localstack
aws sns list-topics --profile localstack

aws --endpoint-url=http://localhost:4566 sqs list-queues --profile localstack

