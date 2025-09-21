# Go Elasticsearch Fun 🚀

A collection of Go scripts and tools for learning and experimenting with Elasticsearch, Terraform, and Go development for myself. This repository serves as a hands-on learning playground for building practical Elasticsearch integrations and infrastructure automation. 

## 🎯 Purpose

This repository is designed to:
- Explore Go programming through practical Elasticsearch integrations
- Learn Terraform for infrastructure as code with Elasticsearch deployments
- Build reusable tools and utilities for Elasticsearch operations
- Practice modern DevOps workflows and best practices

## 🛠️ Tech Stack

- **Go**: Primary programming language for tools and scripts
- **Elasticsearch**: Version 9
- **Terraform**: Infrastructure as code for deployment automation
- **GitHub Actions**: CI/CD pipelines (planned)

## 📁 Projects

### 001_go_es_health
**Elasticsearch Cluster Health Monitor**

A dynamic cluster health checker that connects to Elastic Cloud deployments using API key authentication.

**Features:**
- 🔐 Secure API key authentication for Elastic Cloud 【3】
- 🎨 Color-coded health status indicators (green/yellow/red)
- 📊 Dynamic field display - automatically shows all cluster health metrics 【1】
- 🔄 Future-proof against Elasticsearch API changes
- ⚡ Real-time cluster monitoring capabilities

**Usage:**
```bash
cd 001_go_es_health
export ELASTIC_CLOUD_ID=&quot;your-cloud-id&quot;
export ELASTIC_API_KEY=&quot;your-api-key&quot;
go run main.go

```

## 🗺️ Learning Roadmap (to keep me honest)

### Phase 1: Go Fundamentals (Week 1-2)
**Awareness & Desire: Connect Go to your existing ES/Terraform work**

- [ ] Go syntax and basic patterns
- [ ] Error handling (critical for SRE tools)
- [ ] JSON marshaling/unmarshaling (for ES APIs)
- [ ] HTTP clients and REST API interaction

**Projects:**
- `001_go_es_health` - Cluster health monitoring with dynamic JSON handling

### Phase 2: Go + Elasticsearch Integration (Week 3)
**Knowledge & Ability: Build practical tools**

- Elasticsearch Go client library mastery [^4][^2]
- Index management automation
- Query builders and search operations
- Monitoring and alerting tools

**Planned Projects:**
- `002_go_es_indexer` - Bulk document indexing with performance metrics
- `003_go_es_search` - Advanced search query builder and executor
- `004_go_es_monitor` - Custom metrics collection and alerting

### Phase 3: Infrastructure as Code with Go (Week 4)
**Reinforcement: Advanced SRE automation**

- Terraform provider development basics
- Go-based infrastructure validation tools
- Configuration management utilities

**Planned Projects:**
- `005_terraform_es_cluster` - Terraform modules for Elasticsearch deployment
- `006_go_terraform_validator` - Infrastructure validation tools
- `007_go_config_manager` - Configuration management utilities

### Phase 4: SRE Project Synthesis (Week 5-6)
**Destiny: Real-world application**

- Build a complete monitoring/alerting system
- Automate ES cluster health checks
- Create Terraform modules with Go validation

**Capstone Projects:**
- `008_sre_monitoring_suite` - Complete monitoring/alerting system
- `009_automated_es_ops` - Full cluster lifecycle management
- `010_terraform_go_integration` - Terraform + Go workflow automation

## 🎯 Development Philosophy

**Start Small → Scale Up → Integrate**

1. **Start small**: Build a simple Go program that does ONE of my existing curl/terraform commands
2. **Add complexity**: Handle version detection and headers
3. **Scale up**: Make it work with all my sample data types
4. **Integrate**: Connect it to your Terraform workflow

Each project follows this progression, building from simple scripts to complex integrated systems.

## 🚀 Getting Started

### Prerequisites
- Go 1.19+ installed
- Elasticsearch Cloud deployment (or local cluster)
- API key with cluster monitoring permissions

### Quick Start
1. Clone the repository:
   ```bash
   git clone https://github.com/ZBrisson/go-es-fun.git
   cd go-es-fun
   ```

2. Navigate to a project directory:
   ```bash
   cd 001_go_es_health
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Set up environment variables and run:
   ```bash
   export ELASTIC_CLOUD_ID="your-cloud-id"
   export ELASTIC_API_KEY="your-api-key"
   go run main.go
   ```

## 🤝 Contributing

This is a learning repository, but contributions and suggestions are welcome! Feel free to:
- Open issues for bugs or feature requests
- Submit pull requests with improvements
- Share ideas for new projects or tools

## 📚 Learning Resources

- [Elasticsearch Go Client Documentation](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/) [^4][^2]
- [Go Programming Language](https://golang.org/doc/)
- [Terraform Elasticsearch EC Provider](https://registry.terraform.io/providers/elastic/ec/latest)
- [Terraform Elasticsearch Stack Provider](https://registry.terraform.io/providers/elastic/elasticstack/latest)
- [Elasticsearch API Reference](https://www.elastic.co/guide/en/elasticsearch/reference/current/) [^1]
