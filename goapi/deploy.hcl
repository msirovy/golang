job "goapi" {
	datacenters = ["dc1", "dev", "infra"]
	type = "service"

	group "webapi" {
		network {
			port "rest" {
				to = 8080
			}
		}

		service {
			name = "${JOB}-http"
			port = "rest"

			check {
        name     = "${NOMAD_JOB_NAME} - alive"
        type     = "http"
        path     = "/api"
				interval = "10s"
				timeout  = "5s"
			}
		}

		task "api" {
			driver = "docker"
			
			config {
				image = "webapi:0.1.0"
			}
		}

	} # Group
	
}
