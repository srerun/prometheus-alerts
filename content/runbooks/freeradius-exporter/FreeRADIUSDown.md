---
title: FreeRADIUSDown
description: Troubleshooting for alert FreeRADIUSDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSDown

The FreeRADIUS server is not reachable. Please check the server immediately.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSDown" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSDown
expr: freeradius_up == 0
for: 1m
labels:
    severity: critical
annotations:
    summary: FreeRADIUS is down
    description: The FreeRADIUS server is not reachable. Please check the server immediately.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiusdown/

```

{{% /comment %}}

</details>


Here is a runbook for the FreeRADIUSDown alert:

## Meaning
The FreeRADIUSDown alert indicates that the FreeRADIUS server is not reachable and has been down for at least 1 minute. This alert is critical because it affects the ability to authenticate and authorize users, which can impact the availability of critical services.

## Impact
The impact of this alert is high, as it can cause:

* Disruption to user authentication and authorization
* Inability to access critical services and applications
* Potential security risks due to unauthorized access
* Increased support tickets and user frustration

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the FreeRADIUS server logs for errors or anomalies
2. Verify that the server is running and that there are no network connectivity issues
3. Check the system resources (CPU, memory, disk space) to ensure they are within normal limits
4. Verify that the FreeRADIUS exporter is properly configured and running

## Mitigation
To mitigate the issue, follow these steps:

1. Restart the FreeRADIUS server and verify that it is running properly
2. Check and resolve any network connectivity issues
3. Check and resolve any system resource issues (e.g., clear disk space, restart services)
4. Verify that the FreeRADIUS exporter is properly configured and running
5. Perform a rolling restart of the FreeRADIUS server to ensure that it is running with the correct configuration
6. Monitor the server for any further issues and take corrective action as needed

Note: It is essential to resolve this issue as quickly as possible to minimize the impact on users and services.