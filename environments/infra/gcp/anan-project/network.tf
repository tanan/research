resource "google_compute_network" "anan_project" {
  name                    = "anan-project"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "ne1_subnet" {
  name          = "ne1-subnet"
  ip_cidr_range = "10.146.0.0/20"
  region        = "asia-northeast1"
  network       = google_compute_network.anan_project.self_link

  secondary_ip_range {
    range_name    = "svc01"
    ip_cidr_range = "10.146.128.0/20"
  }

  secondary_ip_range {
    range_name    = "svc02"
    ip_cidr_range = "10.146.144.0/20"
  }

  secondary_ip_range {
    range_name    = "svc03"
    ip_cidr_range = "10.146.160.0/20"
  }

  secondary_ip_range {
    range_name    = "svc04"
    ip_cidr_range = "10.146.176.0/20"
  }

  secondary_ip_range {
    range_name    = "pod-shared"
    ip_cidr_range = "10.146.192.0/18"
  }
}

###############################################################################
# private access for managed services
# https://cloud.google.com/vpc/docs/configure-private-services-access
resource "google_compute_global_address" "anan_project_servicenetworking" {
  name          = "anan-project-servicenetworking"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  address       = "10.147.0.0"
  prefix_length = 16
  network       = google_compute_network.anan_project.self_link
}

resource "google_service_networking_connection" "servicenetworking_googleapis_com" {
  network                 = google_compute_network.anan_project.self_link
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.anan_project_servicenetworking.name]
}