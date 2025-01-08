---
title: IstioHighTotalRequestRate
description: Troubleshooting for alert IstioHighTotalRequestRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioHighTotalRequestRate

Global request rate in the service mesh is unusually high.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioHighTotalRequestRate" %}}

{{% comment %}}

```yaml
alert: IstioHighTotalRequestRate
expr: sum(rate(istio_requests_total{reporter="destination"}[5m])) > 1000
for: 2m
labels:
    severity: warning
annotations:
    summary: Istio high total request rate (instance {{ $labels.instance }})
    description: |-
        Global request rate in the service mesh is unusually high.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioHighTotalRequestRate.md

```

{{% /comment %}}

</details>


**Alert Name**: `IstioHighTotalRequestRate`  
**Severity**: Warning  
**Alert Expression**:  
```prometheus
sum(rate(istio_requests_total{reporter="destination"}[5m])) > 1000
```

---

## Meaning

This alert fires when the total rate of requests within the Istio service mesh exceeds 1000 requests per second over the last 5 minutes. The `istio_requests_total` metric is typically used to track the number of requests handled by Istio proxies. This specific rule monitors the rate at which requests are being sent to service instances.

- **VALUE**: The number of requests per second calculated from the rate expression (`sum(rate(istio_requests_total{reporter="destination"}[5m]))`).
- **LABELS**: Additional contextual information such as the instance of the service.

## Impact

- **Performance Impact**: A high request rate may indicate abnormal traffic patterns, potentially leading to service degradation or an overload of application resources (CPU, memory, network).
- **Potential Causes**:
  - A sudden surge in traffic from clients or misconfigured services.
  - DDoS-like attack or bot traffic flooding the service mesh.
  - Unanticipated behavior in the application, such as excessive retries or loops.
  
Failure to address this issue could lead to service instability, high latencies, or outages if the infrastructure cannot handle the load.

## Diagnosis

To investigate the cause of this alert:
1. **Check the Request Rate**:
   - Examine the raw request rate data over time using Prometheus queries or Grafana dashboards. Confirm if the spike is sustained or just a transient anomaly.
   - Investigate which specific services or instances are driving the high traffic. Use the `instance` label to pinpoint affected nodes.
   
2. **Review Application Logs**:
   - Check application logs for unusual request patterns or potential errors, retries, or loops within the services.
   
3. **Monitor Proxy Configuration**:
   - Inspect Istio proxy configuration to verify that routing, retries, and circuit breakers are set appropriately.
   
4. **Check for External Factors**:
   - Investigate if there are any external factors such as a surge in client traffic, bot activity, or potential DDoS attacks.

5. **Traffic Source**:
   - Determine the source of the high traffic (e.g., specific client, internal system, or external network).

## Mitigation

To mitigate the impact of this alert:
1. **Adjust Rate Limits**:
   - If traffic is legitimate but overwhelming, consider implementing rate limiting or throttling to control the incoming request rate.
   
2. **Scale Application Resources**:
   - If the high request rate is expected and normal, scale the application horizontally by adding more replicas to handle the increased load.
   
3. **Traffic Shaping or Control**:
   - Use Istio's traffic control policies (e.g., request mirroring, retries, circuit breaking) to prevent service overload and gracefully handle high traffic.
   
4. **Investigate and Block Malicious Traffic**:
   - If malicious traffic is suspected (e.g., DDoS), investigate the source and block or throttle incoming requests accordingly using Istio's ingress/egress controls or external WAF tools.
   
5. **Optimize Application Behavior**:
   - Review and optimize application logic to avoid excessive request generation, such as unnecessary retries or resource-hungry processes.

6. **Alert Tuning**:
   - If the threshold of 1000 requests per second is too sensitive for your environment, consider adjusting the alert threshold or duration.


