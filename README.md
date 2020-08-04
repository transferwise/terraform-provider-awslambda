# terraform-provider-awslambda

## Build

```console
$ go build -o terraform-provider-awslambda 
```

## Installation

```console
$ cp terraform-provider-awslambda ~/.terraform.d/plugins
```

## Usage

Defining the provider. This provider only supports three variables. They are `region`, `profile`, and `role_arn`.

```terraform
provider "awslambda" {
  region   = "us-west-1"
  profile  = "default"
  role_arn = "arn:aws:iam::account-id:role/role-name-with-path"
}
```

Example:

```terraform
resource "aws_lambda_invocation" "example" {
  function_name = "${aws_lambda_function.lambda_function_test.function_name}"

  input = <<JSON
 {
   "key1": "value1",
   "key2": "value2"
 }
 JSON
}

output "result_entry" {
  value = jsondecode(aws_lambda_invocation.example.result)["key1"]
}
```
