# sqs
VERY simple amazon sqs send recv example

For this example to work you need two aws config files

> ~/.aws/credentials
```
[default]
aws_access_key_id = your_keyid
aws_secret_access_key = your_secret, both are coming from IAM console / Access Keys
```
> ~/.aws/config
```
[default]
region=eu-west-2 <--- your region
output=json
```
