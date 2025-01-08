---
title: HostClockNotSynchronising
description: Troubleshooting for alert HostClockNotSynchronising
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostClockNotSynchronising

Clock not synchronising. Ensure NTP is configured on this host.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostClockNotSynchronising" %}}

{{% comment %}}

```yaml
alert: HostClockNotSynchronising
expr: (min_over_time(node_timex_sync_status[1m]) == 0 and node_timex_maxerror_seconds >= 16) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host clock not synchronising (instance {{ $labels.instance }})
    description: |-
        Clock not synchronising. Ensure NTP is configured on this host.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostClockNotSynchronising.md

```

{{% /comment %}}

</details>


## Meaning

The HostClockNotSynchronising alert rule is triggered when a host's clock is not synchronizing with an NTP (Network Time Protocol) server. This alert is raised when the `node_timex_sync_status` metric reports a value of 0, indicating that the clock is not synchronized, and the `node_timex_maxerror_seconds` metric exceeds 16 seconds, indicating a significant deviation from the correct time.

## Impact

If left unchecked, a desynchronized clock can cause a range of issues, including:

* Inaccurate timestamps on log entries and other system events
* Potential security vulnerabilities due to outdated system time
* Inconsistent behavior or failures in applications that rely on accurate timekeeping
* Incorrect scheduling of tasks and jobs

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `node_timex_sync_status` and `node_timex_maxerror_seconds` metrics to confirm the clock is not synchronized and the error is significant.
2. Verify that NTP is configured and enabled on the affected host.
3. Check the system logs for errors related to NTP or time synchronization.
4. Verify that the host can reach the NTP server(s) configured in the system.

## Mitigation

To mitigate the issue, follow these steps:

1. Ensure NTP is correctly configured and enabled on the affected host.
2. Restart the NTP service to attempt to resynchronize the clock.
3. Verify that the host can reach the NTP server(s) configured in the system.
4. Check the system logs to ensure that NTP is functioning correctly and the clock is synchronized.
5. If the issue persists, consider adjusting the NTP configuration or seeking assistance from a system administrator.

For more detailed instructions and troubleshooting steps, refer to the runbook at [https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostClockNotSynchronising.md](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostClockNotSynchronising.md).