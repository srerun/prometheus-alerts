---
title: FreeRADIUSDroppedRequests
description: Troubleshooting for alert FreeRADIUSDroppedRequests
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSDroppedRequests

The FreeRADIUS server has dropped requests. Investigate potential performance issues or client misconfigurations.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSDroppedRequests" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSDroppedRequests
expr: freeradius_total_acct_dropped_requests > 0 or freeradius_total_auth_dropped_requests > 0
for: 5m
labels:
    severity: warning
annotations:
    summary: Dropped Requests Detected
    description: The FreeRADIUS server has dropped requests. Investigate potential performance issues or client misconfigurations.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiusdroppedrequests/

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule:

## Meaning

The FreeRADIUSDroppedRequests alert is triggered when the FreeRADIUS server has dropped one or more requests in the last 5 minutes. This can be either accounting requests (freeradius_total_acct_dropped_requests) or authentication requests (freeradius_total_auth_dropped_requests). Dropped requests can indicate performance issues or misconfigurations on the client-side or on the FreeRADIUS server itself.

## Impact

The impact of dropped requests can vary depending on the specific scenario, but it may lead to:

* Authentication failures for users
* Inaccurate accounting data
* Performance degradation of the FreeRADIUS server
* Frustration for users and administrators due to unpredictable behavior

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the FreeRADIUS server logs for any errors or warnings related to dropped requests.
2. Verify that the FreeRADIUS server is properly configured and that there are no misconfigurations on the client-side.
3. Check the system resources (CPU, memory, disk space) of the FreeRADIUS server to ensure it is not experiencing performance issues.
4. Investigate any recent changes to the FreeRADIUS server configuration or software updates that may be causing the issue.
5. Review the Prometheus metrics to identify patterns or trends in dropped requests (e.g., are they occurring during peak usage hours?).

## Mitigation

To mitigate the issue, follow these steps:

1. Resolve any errors or warnings found in the FreeRADIUS server logs.
2. Adjust the FreeRADIUS server configuration to optimize performance and reduce the likelihood of dropped requests.
3. Verify that clients are properly configured to connect to the FreeRADIUS server.
4. Consider increasing system resources (CPU, memory, disk space) if the FreeRADIUS server is experiencing performance issues.
5. Implement measures to prevent dropped requests in the future, such as load balancing or redundancy configurations.

Remember to monitor the situation closely after implementing these mitigation steps to ensure the issue is resolved and does not reoccur.