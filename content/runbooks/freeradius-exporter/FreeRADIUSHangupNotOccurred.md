---
title: FreeRADIUSHangupNotOccurred
description: Troubleshooting for alert FreeRADIUSHangupNotOccurred
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSHangupNotOccurred

The FreeRADIUS server has not been hup'd since it started. Consider restarting it if necessary.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSHangupNotOccurred" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSHangupNotOccurred
expr: freeradius_start_time == freeradius_hup_time
for: 5m
labels:
    severity: warning
annotations:
    summary: FreeRADIUS server has not been restarted
    description: The FreeRADIUS server has not been hup'd since it started. Consider restarting it if necessary.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiushangupnotoccurred/

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule:

## Meaning

The FreeRADIUSHangupNotOccurred alert indicates that the FreeRADIUS server has not been restarted (hup'd) since it started. This means that the server has not received a hangup signal (SIGHUP) to reload its configuration and restart. This can lead to issues with the server's functionality and performance.

## Impact

The impact of this alert can be significant, as a failure to restart the FreeRADIUS server can result in:

* Inconsistent or outdated configuration
* Memory leaks or resource exhaustion
* Inability to apply security patches or updates
* Disruption to authentication and authorization services

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the FreeRADIUS server logs for any errors or issues related to the last restart or configuration reload.
2. Verify that the FreeRADIUS server is configured to receive SIGHUP signals and restart correctly.
3. Check the system logs for any errors or issues related to the system's ability to send SIGHUP signals to the FreeRADIUS server.
4. Verify that the FreeRADIUS server is running with the correct configuration and that there are no issues with the configuration file.

## Mitigation

To mitigate the issue, perform the following steps:

1. Restart the FreeRADIUS server to reload its configuration and apply any pending changes.
2. Verify that the FreeRADIUS server is configured to receive SIGHUP signals and restart correctly.
3. Check the system logs to ensure that the system is able to send SIGHUP signals to the FreeRADIUS server.
4. Consider implementing a scheduled restart or reload of the FreeRADIUS server to prevent this issue from occurring in the future.

Note: This runbook provides general guidance and may need to be tailored to your specific environment and requirements.