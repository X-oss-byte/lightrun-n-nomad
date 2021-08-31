variable "lightrun-server" {
  type = string
}
variable "lightrun-secret" {
  type = string
}
variable "lightrun-certificate" {
  type = string
}


job "lightrun-standalone-example" {
  datacenters = ["dc1"]
  type = "batch"

  group "lightrun-group" {
    task "web" {
      driver = "java"

      artifact {
        source      = "<path_to_your_jar_here>"
        destination = "local/<example>.jar"
        mode        = "file"
        options {}
      }


      artifact {
        source      = "https://lightrun-server-files.s3.amazonaws.com/Nomad/agent.zip"
        destination = "local/"
        mode        = "any"
        options {}
      }

      config {
        class       = "<className_here>"
        class_path  = "local/<example>.jar"
        jvm_options = ["-Xmx2048m", "-Xms256m", "-agentpath:local/lightrun_agent.so", "-Dcom.lightrun.server=${var.lightrun-server}", "-Dcom.lightrun.secret=${var.lightrun-secret}", "-Dpinned_certs=${var.lightrun-certificate}", "-Dtransmission_bulk_max_size=10"]
      }
    }
  }
}
