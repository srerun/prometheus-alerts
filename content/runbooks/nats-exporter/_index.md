---
title: nats-exporter

bookCollapseSection: true
bookFlatSection: true
weight: 1
---


Below is a structured runbook for the alerts defined in the provided Prometheus alert rules file. Each alert includes sections for meaning, impact, diagnosis, and mitigation.

---

## Runbook for NATS Alerts

Below is the updated runbook with links to the relevant runbook pages for each alert:

---

### **1. NatsHighConnectionCount**
- **Meaning**: The number of NATS connections exceeds 100 for more than 3 minutes.
- **Impact**: Could indicate resource exhaustion or potential misuse of the system.
- **Diagnosis**:
  - Check the number of connections using `gnatsd_varz_connections`.
  - Identify which clients are connected and their connection rates.
- **Mitigation**:
  - Ensure legitimate usage patterns.
  - Increase server capacity or adjust connection limits if required.
  - Investigate and terminate any unauthorized connections.
- **Runbook Link**: [NatsHighConnectionCount](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighConnectionCount.md)

---

### **2. NatsHighPendingBytes**
- **Meaning**: Pending bytes in NATS connections exceed 100,000 for more than 3 minutes.
- **Impact**: Potential delays in message delivery, leading to application performance issues.
- **Diagnosis**:
  - Inspect `gnatsd_connz_pending_bytes` for the affected instance.
  - Identify publishers with high message rates or slow consumers.
- **Mitigation**:
  - Tune client configurations to optimize publishing and consumption rates.
  - Scale consumers or increase processing capacity.
- **Runbook Link**: [NatsHighPendingBytes](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighPendingBytes.md)

---

### **3. NatsHighSubscriptionsCount**
- **Meaning**: The number of subscriptions exceeds 50 for more than 3 minutes.
- **Impact**: May indicate suboptimal client behavior or over-subscription to channels.
- **Diagnosis**:
  - Examine `gnatsd_connz_subscriptions` for the instance.
  - Check for duplicate or unnecessary subscriptions.
- **Mitigation**:
  - Optimize application subscription logic.
  - Reduce redundant or overlapping subscriptions.
- **Runbook Link**: [NatsHighSubscriptionsCount](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighSubscriptionsCount.md)

---

### **4. NatsHighRoutesCount**
- **Meaning**: The number of routes in the cluster exceeds 10 for more than 3 minutes.
- **Impact**: Could lead to increased resource usage and complexity in routing messages.
- **Diagnosis**:
  - Check `gnatsd_varz_routes` for active routes.
  - Review the cluster topology for misconfigurations or unexpected peers.
- **Mitigation**:
  - Simplify the cluster design by reducing unnecessary routes.
  - Review and correct any misconfigured peer connections.
- **Runbook Link**: [NatsHighRoutesCount](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighRoutesCount.md)

---

### **5. NatsHighMemoryUsage**
- **Meaning**: Memory usage of the NATS server exceeds 200 MB for 5 minutes.
- **Impact**: May result in degraded server performance or crashes.
- **Diagnosis**:
  - Analyze memory usage with `gnatsd_varz_mem`.
  - Check for memory-intensive workloads or leaks.
- **Mitigation**:
  - Optimize message size and retention policies.
  - Restart the server if necessary to clear memory leaks.
- **Runbook Link**: [NatsHighMemoryUsage](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighMemoryUsage.md)

---

### **6. NatsSlowConsumers**
- **Meaning**: Slow consumers are detected for more than 3 minutes.
- **Impact**: May delay or drop messages, affecting downstream applications.
- **Diagnosis**:
  - Identify slow consumers using `gnatsd_varz_slow_consumers`.
  - Investigate network or processing bottlenecks.
- **Mitigation**:
  - Increase consumer processing capacity.
  - Optimize message handling to reduce delays.
- **Runbook Link**: [NatsSlowConsumers](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsSlowConsumers.md)

---

### **7. NatsHighCpuUsage**
- **Meaning**: CPU usage exceeds 80% for 5 minutes.
- **Impact**: Could lead to degraded performance or timeouts.
- **Diagnosis**:
  - Monitor `rate(gnatsd_varz_cpu[5m])` for sustained high usage.
  - Identify and optimize CPU-intensive workloads.
- **Mitigation**:
  - Distribute workloads across multiple nodes.
  - Upgrade server hardware if necessary.

---

### **8. NatsHighJetstreamStoreUsage**
- **Meaning**: JetStream store usage exceeds 80% capacity for 5 minutes.
- **Impact**: Risk of message loss due to storage exhaustion.
- **Diagnosis**:
  - Review `gnatsd_varz_jetstream_stats_storage` and `gnatsd_varz_jetstream_config_max_storage`.
  - Check for high storage usage patterns.
- **Mitigation**:
  - Increase storage capacity.
  - Implement retention policies to manage usage.

---

### **9. NatsHighJetstreamMemoryUsage**
- **Meaning**: JetStream memory usage exceeds 80% of the configured limit for 5 minutes.
- **Impact**: May result in message processing slowdowns or failures.
- **Diagnosis**:
  - Check `gnatsd_varz_jetstream_stats_memory` and `gnatsd_varz_jetstream_config_max_memory`.
  - Look for memory spikes due to high message throughput or retention policies.
- **Mitigation**:
  - Increase JetStream memory allocation.
  - Optimize message sizes and retention policies.

---

### **10. NatsHighNumberOfSubscriptions**
- **Meaning**: The number of subscriptions exceeds 1,000 for more than 5 minutes.
- **Impact**: May lead to increased resource usage and delays in message delivery.
- **Diagnosis**:
  - Monitor `gnatsd_connz_subscriptions` for the instance.
  - Check if any clients are creating excessive subscriptions.
- **Mitigation**:
  - Limit the number of subscriptions per client.
  - Optimize subscription patterns to prevent duplication.

---

### **11. NatsTooManyErrors**
- **Meaning**: API errors in JetStream increase for more than 5 minutes.
- **Impact**: Indicates potential instability or misconfiguration in JetStream.
- **Diagnosis**:
  - Inspect `increase(gnatsd_varz_jetstream_stats_api_errors[5m])`.
  - Review logs for error details and patterns.
- **Mitigation**:
  - Address configuration issues.
  - Resolve application-level errors causing API failures.

---

### **12. NatsJetstreamConsumersExceeded**
- **Meaning**: The number of JetStream consumers exceeds 100 for more than 5 minutes.
- **Impact**: Could lead to excessive resource usage or message processing bottlenecks.
- **Diagnosis**:
  - Examine `sum(gnatsd_varz_jetstream_stats_accounts)` for consumer counts.
  - Identify which consumers are contributing to the high count.
- **Mitigation**:
  - Optimize consumer creation logic in applications.
  - Distribute workload among fewer, more efficient consumers.

---

### **13. NatsFrequentAuthenticationTimeouts**
- **Meaning**: More than 5 authentication timeouts occur within 5 minutes.
- **Impact**: Indicates issues with authentication, potentially blocking connections.
- **Diagnosis**:
  - Review `increase(gnatsd_varz_auth_timeout[5m])`.
  - Check authentication server or configuration logs for anomalies.
- **Mitigation**:
  - Address authentication server issues.
  - Adjust authentication timeout settings.

---

### **14. NatsMaxPayloadSizeExceeded**
- **Meaning**: Payload size exceeds the configured maximum of 1MB for 5 minutes.
- **Impact**: May lead to message rejection or delivery failures.
- **Diagnosis**:
  - Monitor `max(gnatsd_varz_max_payload)`.
  - Identify clients or applications sending oversized messages.
- **Mitigation**:
  - Reconfigure clients to respect payload size limits.
  - Increase maximum payload size if feasible.

---

### **15. NatsLeafNodeConnectionIssue**
- **Meaning**: No leaf node connections have been established within 5 minutes.
- **Impact**: Indicates potential issues with leaf node connectivity, leading to message routing problems.
- **Diagnosis**:
  - Inspect `increase(gnatsd_varz_leafnodes[5m])`.
  - Verify network connectivity and leaf node configurations.
- **Mitigation**:
  - Resolve network or configuration issues preventing leaf node connections.
  - Restart affected nodes if necessary.

---

### **16. NatsMaxPingOperationsExceeded**
- **Meaning**: Ping operations exceed 50 for more than 5 minutes.
- **Impact**: Indicates potential instability in connection health checks.
- **Diagnosis**:
  - Review `gnatsd_varz_ping_max` for the instance.
  - Check logs for ping operation patterns.
- **Mitigation**:
  - Optimize client configurations to reduce ping frequency.
  - Ensure stable network conditions.

---

### **17. NatsWriteDeadlineExceeded**
- **Meaning**: Write deadlines are exceeded, indicating potential message delivery issues.
- **Impact**: Could lead to dropped or delayed messages.
- **Diagnosis**:
  - Monitor `gnatsd_varz_write_deadline`.
  - Investigate client write speeds and network conditions.
- **Mitigation**:
  - Increase write deadline limits.
  - Optimize client-side configurations and network conditions.

---

