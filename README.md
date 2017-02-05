# Go Forum

This repository contains the code of our backend forum api.

## How to use

Run the following command:

`http_port=":port" cb_url="couchbase://url" cb_bucket="bucketname" cb_pass="bucketpassword" go run main.go`

## Error codes

1. Not all environment variables are set
2. Recheck the couchbase url
3. Authentication error (the bucket does not exist or the password is wrong)