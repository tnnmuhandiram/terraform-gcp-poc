
provider "google" {
  credentials = "${var.credentails_json}"
  project     = "${var.gcp_project_id}"
  region      = "us-central1"
  zone        = "us-central1-c"
}
resource "google_compute_instance" "example" {
  project = "${var.gcp_project_id}"
  name = "${var.instance_name}"
  machine_type = "${var.machine_type}"
  zone = "${var.zone}"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-1604-lts"
    }
  }

  network_interface {
    network = "default"
    access_config {}
  }
}

resource "google_storage_bucket" "example_bucket" {
  project = "${var.gcp_project_id}"
  name = "${var.bucket_name}"
  location = "${var.bucket_location}"
}
