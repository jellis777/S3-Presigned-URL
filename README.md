# S3 Presigned URL Generator (Go + AWS S3)

This project is a small Go command-line tool that uploads a file to Amazon S3 and generates a presigned URL for sharing that file.

It demonstrates how to use the AWS SDK for Go to interact with S3, upload objects, and securely share files without making the bucket public.

---

## What This Tool Does

1. Accepts a file name and S3 bucket name as command-line arguments
2. Uploads the file to the specified S3 bucket
3. Generates a presigned URL for the uploaded object
4. Prints the URL, which can be used to download the file

The presigned URL expires after a fixed amount of time.

---

## Why This Exists

This project was built to practice:

* Using Amazon S3 with the AWS SDK for Go
* Uploading objects to S3 programmatically
* Generating presigned URLs for secure file sharing
* Working with AWS credentials and configuration
* Structuring a small Go CLI application

---

## Project Structure

```
s3-presigned-url
├── main/
│   └── main.go        # CLI entry point
├── s3share.go         # S3 client setup
├── upload.go          # File upload logic
├── share.go           # Presigned URL generation
├── testdata/
│   └── text.txt       # Sample file
├── go.mod
└── go.sum
```

---

## How It Works (High Level)

* AWS credentials are loaded using the default AWS configuration
* An S3 client is created using the AWS SDK for Go
* The file is uploaded to S3 using `PutObject`
* A presigned URL is generated using `PresignGetObject`
* The URL is valid for one hour

---

## Requirements

* Go installed
* AWS credentials configured locally
* An existing S3 bucket

---

## How to Run

```bash
go run main/main.go -file <filename> -bucket <bucket-name>
```

Example:

```bash
go run main/main.go -file testdata/text.txt -bucket my-s3-bucket
```

---

## Notes

* The S3 bucket must already exist
* The AWS credentials must allow `s3:PutObject` and `s3:GetObject`
* Presigned URLs allow temporary access without making the object public
* Error handling is kept simple for clarity

---

## Purpose of This Project

This project is intentionally small.

It is meant to demonstrate:

* Secure file sharing using S3 presigned URLs
* Basic AWS SDK usage in Go
* Command-line input handling
* Cloud-native file access patterns

It is not intended to be a full production file-sharing system.
