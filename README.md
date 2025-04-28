> If you see this section, you've just created a repository using [PoC Innovation's Open-Source project template](https://github.com/PoCInnovation/open-source-project-template). Check the [getting started guide](./.github/getting-started.md).

# PoC Defender

## Overview
PoC Defender is a project that sets up a Go application using the Gin framework to proxy requests to a backend service. It also integrates Prometheus for monitoring and Grafana for visualization of metrics.

## How does it work?

It works by creating a proxy server that listens for incoming requests. When a request is received, the server checks if the request is malicious or not. If it is, the server blocks the request and returns an error message. If it is not, the server forwards the request to the intended destination.

## Project Structure
```
PoC-Defender
├── src
│   └── main.go          # Main application code
├── docker-compose.yml    # Docker Compose configuration
├── Dockerfile            # Dockerfile for building the Go application
├── prometheus
│   └── prometheus.yml   # Prometheus configuration file
└── README.md            # Project documentation
```

## Getting Started

### Prerequisites
- Docker
- Docker Compose

### Setup
1. Clone the repository:
   ```
   git clone <repository-url>
   cd PoC-Defender
   ```

2. Build and run the application with Docker Compose:
   ```
   docker compose up --build
   ```

### Accessing the Application
- The Go application will be available at `http://127.0.0.1:9000`.
- Prometheus will be available at `http://127.0.0.1:9090`.
- Grafana will be available at `http://127.0.0.1:3000`.

### Monitoring
- Configure Prometheus to scrape metrics from the Go application by modifying the `prometheus/prometheus.yml` file as needed.

### Grafana
- Use Grafana to visualize the metrics collected by Prometheus. You can set up dashboards to monitor the performance and health of the Go application.

## Get involved

You're invited to join this project ! Check out the [contributing guide](./CONTRIBUTING.md).

If you're interested in how the project is organized at a higher level, please contact the current project manager.

## Our PoC team ❤️

Developers
| [<img src="https://github.com/flav-code.png?size=85" width=85><br><sub>flav-code</sub>](https://github.com/flav-code) | [<img src="https://github.com/toutcourtlll.png?size=85" width=85><br><sub>toutcourtlll</sub>](https://github.com/toutcourtlll) | [<img src="https://github.com/t1m0t-p.png?size=85" width=85><br><sub>t1m0t-p</sub>](https://github.com/t1m0t-p)
| :---: | :---: | :---: |

Manager
| [<img src="https://github.com/maitrecraft1234.png?size=85" width=85><br><sub>maitrecraft1234</sub>](https://github.com/maitrecraft1234)
| :---: |

<h2 align=center>
Organization
</h2>

<p align='center'>
    <a href="https://www.linkedin.com/company/pocinnovation/mycompany/">
        <img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" alt="LinkedIn logo">
    </a>
    <a href="https://www.instagram.com/pocinnovation/">
        <img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white" alt="Instagram logo"
>
    </a>
    <a href="https://twitter.com/PoCInnovation">
        <img src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white" alt="Twitter logo">
    </a>
    <a href="https://discord.com/invite/Yqq2ADGDS7">
        <img src="https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white" alt="Discord logo">
    </a>
</p>
<p align=center>
    <a href="https://www.poc-innovation.fr/">
        <img src="https://img.shields.io/badge/WebSite-1a2b6d?style=for-the-badge&logo=GitHub Sponsors&logoColor=white" alt="Website logo">
    </a>
</p>

> 🚀 Don't hesitate to follow us on our different networks, and put a star 🌟 on `PoC's` repositories

> Made with ❤️ by PoC
