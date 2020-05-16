resource google_container_cluster gke01 {
  provider                 = google-beta
  name                     = "gke01"
  location                 = "asia-northeast1-a"
  network                  = google_compute_subnetwork.ne1_subnet.network
  subnetwork               = google_compute_subnetwork.ne1_subnet.self_link
  initial_node_count       = "1"
  min_master_version       = "1.16.8-gke.15"
  node_version             = "1.16.8-gke.15"
  remove_default_node_pool = true

  lifecycle {
    ignore_changes = [
      master_auth,
      node_version,
    ]
  }

  ip_allocation_policy {
    cluster_secondary_range_name  = "pod-shared"
    services_secondary_range_name = "svc01"
  }

  # disable basic auth and client certificate publish.
  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }

  logging_service    = "logging.googleapis.com/kubernetes"
  monitoring_service = "monitoring.googleapis.com/kubernetes"
}

resource google_container_node_pool gke01_node_pool01 {
  name               = "${google_container_cluster.gke01.name}-node-pool01"
  location           = "asia-northeast1-a"
  cluster            = google_container_cluster.gke01.name
  initial_node_count = 1

  management {
    auto_repair  = true
    auto_upgrade = true
  }

  node_config {
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform",
    ]

    preemptible  = "true"
    machine_type = "n1-standard-1"
    tags         = ["server", "gke"]
  }
}