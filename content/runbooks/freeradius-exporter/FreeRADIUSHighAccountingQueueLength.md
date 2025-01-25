---
title: FreeRADIUSHighAccountingQueueLength
description: Troubleshooting for alert FreeRADIUSHighAccountingQueueLength
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSHighAccountingQueueLength

The accounting queue length ({{ $value }}) is above the threshold of 100.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSHighAccountingQueueLength" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSHighAccountingQueueLength
expr: freeradius_queue_len_acct > 100
for: 5m
labels:
    severity: warning
annotations:
    summary: High Accounting Queue Length
    description: The accounting queue length ({{ $value }}) is above the threshold of 100.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiushighaccountingqueuelength/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning
The FreeRADIUSHighAccountingQueueLength alert is triggered when the length of the Freeradius accounting queue exceeds 100 for a period of 5 minutes. This alert indicates that the accounting queue is growing and may lead to issues with authentication and authorization.

## Impact
A high accounting queue length can cause:

* Delays in authentication and authorization
* Increased latency for users
* Potential loss of accounting data
* Decreased overall system performance

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the Freeradius server logs for any errors or issues related to accounting.
2. Verify that the accounting database is functioning correctly and not experiencing any issues.
3. Check the system resources (CPU, memory, disk space) to ensure they are not overwhelmed.
4. Review the Freeradius configuration to ensure it is correctly set up and not causing any bottlenecks.
5. Check the accounting queue length over time to determine if it is consistently high or if it's a one-time spike.

## Mitigation
To mitigate the issue, follow these steps:

1. Check the accounting database and ensure it is running optimally. Consider optimizing database queries or indexing to improve performance.
2. Verify that the Freeradius server is configured correctly and not causing any bottlenecks. Consider tuning server settings or adjusting the configuration to improve performance.
3. Investigate and resolve any system resource issues (CPU, memory, disk space) that may be contributing to the high accounting queue length.
4. Consider increasing the accounting queue size or adjusting the queue processing timeout to prevent further issues.
5. If the issue persists, consider reaching out to the Freeradius maintainers or a system administrator for further assistance.

Remember to update the alert threshold and adjust the runbook according to your specific system requirements and needs.