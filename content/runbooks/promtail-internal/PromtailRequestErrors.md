---
title: PromtailRequestErrors
description: Troubleshooting for alert PromtailRequestErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PromtailRequestErrors

The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}% errors.

<details>
  <summary>Alert Rule</summary>

{{% rule "promtail/promtail-internal.yml" "PromtailRequestErrors" %}}

{{% comment %}}

```yaml
alert: PromtailRequestErrors
expr: 100 * sum(rate(promtail_request_duration_seconds_count{status_code=~"5..|failed"}[1m])) by (namespace, job, route, instance) / sum(rate(promtail_request_duration_seconds_count[1m])) by (namespace, job, route, instance) > 10
for: 5m
labels:
    severity: critical
annotations:
    summary: Promtail request errors (instance {{ $labels.instance }})
    description: |-
        The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}% errors.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/promtail-internal/PromtailRequestErrors.md

```

{{% /comment %}}

</details>

## Meaning
This alert triggers when more than 10% of Promtail requests result in 5xx errors or failed responses over a 1-minute window, sustained for at least 5 minutes. It indicates that a significant portion of requests are failing, potentially impacting log ingestion and monitoring pipelines.

## Impact
- **Log Ingestion Failure:** Logs may not be ingested properly, leading to gaps in observability.
- **Service Degradation:** Downstream services relying on logs for monitoring, debugging, or auditing could be affected.
- **Increased Latency:** The issue may lead to bottlenecks or retries, further stressing the system.

---

## Diagnosis
### Key Metrics to Investigate
1. **Error Rate:**
   - Inspect the percentage of errors in Promtail requests:
     ```
     sum(rate(promtail_request_duration_seconds_count{status_code=~"5..|failed"}[1m])) by (namespace, job, route, instance)
     ```
   - Compare it with the total request count:
     ```
     sum(rate(promtail_request_duration_seconds_count[1m])) by (namespace, job, route, instance)
     ```
2. **Affected Instances:**
   - Identify the specific instance(s), namespace(s), job(s), or route(s) showing high error rates from alert labels.

### Logs and Debugging Steps
1. **Promtail Logs:**
   - Look for errors or warnings in Promtail logs related to the affected instance:
     ```
     kubectl logs -n <namespace> <promtail-pod> | grep -i error
     ```
2. **Status Codes:**
   - Analyze the nature of 5xx errors:
     - Are they consistent or intermittent?
     - Do they originate from a specific route or endpoint?
3. **Downstream Dependencies:**
   - Check if external services or storage backends (e.g., Loki) are causing these failures.
4. **Network Issues:**
   - Verify network connectivity between Promtail and dependent services.
5. **Resource Constraints:**
   - Check if the Promtail pods are resource-starved:
     ```
     kubectl top pods -n <namespace>
     ```

---

## Mitigation
### Immediate Actions
1. **Restart Promtail Pods:**
   - Restart the affected Promtail pods to resolve transient issues:
     ```
     kubectl rollout restart deployment <promtail-deployment> -n <namespace>
     ```
2. **Increase Resources:**
   - If Promtail is resource-starved, scale up resources or increase limits in the deployment:
     ```yaml
     resources:
       requests:
         cpu: 500m
         memory: 512Mi
       limits:
         cpu: 1
         memory: 1Gi
     ```
3. **Route Isolation:**
   - Temporarily disable or throttle problematic routes, if identifiable.

### Long-term Fixes
1. **Error Budget Monitoring:**
   - Set up monitoring for error budgets and thresholds to catch issues earlier.
2. **Improve Retries and Backoff:**
   - Optimize retry and backoff logic for Promtail requests.
3. **Dependency Health:**
   - Monitor and improve the stability of downstream systems (e.g., Loki, storage backends).
4. **Scaling:**
   - Ensure Promtail deployments are scaled appropriately for the workload.
5. **Alert Refinement:**
   - Review alert thresholds and adjust if necessary to reduce false positives.

---

## Additional Resources
- [Promtail Documentation](https://grafana.com/docs/loki/latest/clients/promtail/)
- [Prometheus Querying Basics](https://prometheus.io/docs/prometheus/latest/querying/basics/)
