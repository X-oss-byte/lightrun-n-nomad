job "lightrun-example" {

  datacenters = ["dc1"]
  type = "batch"

  group "lightrun-group" {
    task "web" {
      driver = "lightrun-java"

      artifact {
        source      = "<path_to_your_jar>"
        destination = "local/example.jar"
        mode        = "file"
        options {}
      }

      config {
        class       = "<class_name>"
        class_path  = "local/example.jar"
        jvm_options = ["-Xmx2048m", "-Xms256m"]
        lightrun_server = "https://app.lightrun.com/company/<company_name>"
        lightrun_secret = "<agent_secret>"
        lightrun_certificate = "ee80811b38e7e6c2dc4cc372cbea86bd86b446b012e427f2e19bf094afba5d12"
      }
    }
  }
}

