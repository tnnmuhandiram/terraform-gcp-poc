
variable "gcp_project_id" {
  description = "The ID of the GCP project in which these resources will be created."
}

variable "instance_name" {
  description = "The Name to use for the Cloud Instance."
  default     = "terratest-example"
}

variable "machine_type" {
  description = "The Machine Type to use for the Cloud Instance."
  default     = "f1-micro"
}

variable "zone" {
  description = "The Zone to launch the Cloud Instance into."
  default     = "us-central1-a"
}

variable "bucket_name" {
  description = "The Name of the example Bucket to create."
  default     = "gruntwork-terratest-bucket"
}

variable "bucket_location" {
  description = "The location to store the Bucket. This value can be regional or multi-regional."
  default     = "US"
}

variable "credentails_json" {
  description = "The creadentails to authntication Google Cloud Platform"
}
