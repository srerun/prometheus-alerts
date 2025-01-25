---
title: FreeRADIUSStatsErrorDetected
description: Troubleshooting for alert FreeRADIUSStatsErrorDetected
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSStatsErrorDetected

A stats error has been reported in the FreeRADIUS server. Investigate the source of the error.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSStatsErrorDetected" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSStatsErrorDetected
expr: freeradius_stats_error == 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Stats Error Detected
    description: A stats error has been reported in the FreeRADIUS server. Investigate the source of the error.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiusstatserrordetected/

```

{{% /comment %}}

</details>


## Meaning

The `FreeRADIUSStatsErrorDetected` alert is triggered when the FreeRADIUS server reports a stats error. This error can indicate a problem with the FreeRADIUS server's ability to collect or process statistical data, which can impact the accuracy of billing, authentication, and authorization processes.

## Impact

The impact of this alert is potentially high, as it can lead to:

* Inaccurate billing and revenue loss
* Authentication and authorization issues
* Difficulty in troubleshooting and debugging issues due to incomplete or inaccurate statistical data

## Diagnosis

To diagnose the root cause of the `FreeRADIUSStatsErrorDetected` alert, follow these steps:

1. Check the FreeRADIUS server logs for error messages related to stats collection and processing.
2. Verify that the FreeRADIUS server is properly configured to collect and process statistical data.
3. Check the FreeRADIUS exporter configuration to ensure it is correctly configured to scrape stats data from the FreeRADIUS server.
4. Review the network and system logs to identify any underlying issues that may be contributing to the stats error.

## Mitigation

To mitigate the impact of the `FreeRADIUSStatsErrorDetected` alert, follow these steps:

1. Immediately investigate and resolve the underlying cause of the stats error.
2. Restart the FreeRADIUS server and exporter to ensure that the stats collection and processing are restarted.
3. Verify that the stats data is being collected and processed correctly after the restart.
4. Implement additional monitoring and logging to detect and alert on similar issues in the future.
5. Consider implementing redundant stats collection and processing mechanisms to mitigate the impact of future errors.