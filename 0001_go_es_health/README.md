### 001_go_es_health
**Elasticsearch Cluster Health Monitor**

A dynamic cluster health checker that connects to Elastic Cloud deployments using API key authentication utilizing the _cluster/health API [^1]

**Features:**
- 🔐 Secure API key authentication for Elastic Cloud
- 🎨 Color-coded health status indicators (green/yellow/red)
- 📊 Dynamic field display - automatically shows all cluster health metrics 
- 🔄 Future-proof against Elasticsearch API changes
- ⚡ Real-time cluster monitoring capabilities

**Usage:**
```bash
cd 001_go_es_health
export ELASTIC_CLOUD_ID="your-cloud-id"
export ELASTIC_API_KEY="your-api-key"
go run main.go
```

**Sample Output:**
```
=== Elasticsearch Cluster Health ===
Status: 🟡 yellow
Active Primary Shards: 56
Active Shards: 56
Active Shards Percent As Number: 88.89
Cluster Name: 7a960ec2650a4452be33dc1ed29e05f5
Delayed Unassigned Shards: 0
Initializing Shards: 0
Number Of Data Nodes: 1
Number Of In Flight Fetch: 0
Number Of Nodes: 1
Number Of Pending Tasks: 0
Relocating Shards: 0
Task Max Waiting In Queue Millis: 0
Timed Out: false
Unassigned Primary Shards: 0
Unassigned Shards: 7

=== Health Status Interpretation ===
⚠️  All primary shards are allocated, but some replica shards are not
⚠️  Warning: 7 unassigned shards detected
```
## References

- [Cluster Health API](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-cluster-health) [^1]


