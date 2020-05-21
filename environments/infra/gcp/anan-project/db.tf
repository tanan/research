resource "google_sql_database_instance" "research03" {
  provider         = google-beta
  name             = "research03"
  region           = "asia-northeast1"
  database_version = "POSTGRES_11"

  settings {
    tier = "db-f1-micro"

    disk_autoresize = true

    # disk_size = 20
    disk_type = "PD_SSD"

    ip_configuration {
      ipv4_enabled    = "false"
      private_network = google_compute_network.anan_project.self_link
    }

    maintenance_window {
      day  = 7
      hour = 22
    }

    backup_configuration {
      binary_log_enabled = false
      enabled            = true
      start_time         = "16:00"
    }

    database_flags {
      name  = "temp_file_limit"
      value = "104857600"
    }

    database_flags {
      name  = "max_connections"
      value = "5000"
    }
  }
}

resource "google_sql_user" "research03_iskandar" {
  name     = "dolphin"
  instance = google_sql_database_instance.research03.name
  password = "dolphin"
}