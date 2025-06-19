# Creating Infra on AWS

Before creating infra make sure you set up AWS creadentials on your system, If not then do so by following instructions

```bash
export AWS_ACCESS_KEY_ID="YOUR_AWS_KEY_ID"
export AWS_SECRET_ACCESS_KEY="YOUR_AWS_SECRET_KEY"
```


Once you have setup credentials, Make sure you have S3 bucket `all-terraform-state-files` & DynamoDB with table `assessments-particle41-terraformtfstate` required for statefile and Locking management respectively otherwise make changes in main.tf backend configuration block or comment that out to manage statefile locally.

then execute following command in order to have setup infra on AWS:

```bash
terraform init
terraform plan # optional to have resource modification report
terraform apply
```
