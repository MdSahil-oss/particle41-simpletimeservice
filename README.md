# Repository struction

```bash
.
├── app
│   ├── Dockerfile
│   ├── go.mod
│   ├── main.go     # application Golang file
│   └── README.md   # application/docker build Readme
├── README.md       # root Readme
└── terraform
    ├── main.tf
    ├── outputs.tf
    ├── README.md   # terraform configuration Readme
    └── variables.tf
```

# Have code locally

To get this repo on your system execute following commands

```bash
git clone https://github.com/MdSahil-oss/particle41-simpletimeservice.git
cd particle41-simpletimeservice
```

# Build & Run container

If you want to build the container image of this microservice execute following commands

```bash
cd app/
docker build -t mdsahiloss/simple-time-service:1.0 .
```

Once the image has built then run following command to run an instance of container image

```bash
docker run -p 8080:8080 mdsahiloss/simple-time-service:1.0
```


# Creating Infra on AWS

Before creating infra make sure you set up AWS creadentials on your system, If not then do so by following instructions

```bash
export AWS_ACCESS_KEY_ID="YOUR_AWS_KEY_ID"
export AWS_SECRET_ACCESS_KEY="YOUR_AWS_SECRET_KEY"
```


Once you have setup credentials, Make sure you have S3 bucket `all-terraform-state-files` & DynamoDB with table `assessments-particle41-terraformtfstate` required for statefile and Locking management respectively otherwise make changes in main.tf backend configuration block or comment that out to manage statefile locally.

then execute following command in order to have setup infra on AWS:

```bash
cd terraform/
terraform init
terraform plan # optional to have resource modification report
terraform apply
```
