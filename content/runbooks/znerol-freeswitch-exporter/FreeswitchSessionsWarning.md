---
title: FreeswitchSessionsWarning
description: Troubleshooting for alert FreeswitchSessionsWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeswitchSessionsWarning

High sessions usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "freeswitch/znerol-freeswitch-exporter.yml" "FreeswitchSessionsWarning" %}}

{{% comment %}}

```yaml
alert: FreeswitchSessionsWarning
expr: (freeswitch_session_active * 100 / freeswitch_session_limit) > 80
for: 10m
labels:
    severity: warning
annotations:
    summary: Freeswitch Sessions Warning (instance {{ $labels.instance }})
    description: |-
        High sessions usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/znerol-freeswitch-exporter/FreeswitchSessionsWarning.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `FreeswitchSessionsWarning`:

## Meaning

The `FreeswitchSessionsWarning` alert is triggered when the percentage of active FreeSWITCH sessions exceeds 80% of the maximum allowed sessions for more than 10 minutes. This indicates that the system is approaching its capacity limit and may start to experience performance issues or dropped calls if the trend continues.

## Impact

If left unaddressed, high session usage can lead to:

* Decreased call quality and dropped calls
* Increased latency and response times
* Potential system crashes or instability
* Negative impact on business operations and revenue

## Diagnosis

To diagnose the root cause of the issue, follow these steps:

1. Check the FreeSWITCH logs for any errors or warnings related to session handling.
2. Verify that the system has sufficient resources (CPU, memory, and disk space) to handle the current load.
3. Investigate recent changes or updates to the system or configuration that may have caused the increased session usage.
4. Review call volume and traffic patterns to identify any unusual or unexpected activity.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the `freeswitch_session_limit` configuration value to allow for more sessions, if possible.
2. Optimize system resources (e.g., add more CPU or memory) to handle the increased load.
3. Investigate and address any underlying issues causing excessive session usage (e.g., fix errors in the FreeSWITCH configuration or dialplan).
4. Consider implementing rate limiting or traffic shaping to prevent sudden spikes in call volume.
5. If the issue persists, consider scaling the system horizontally (adding more nodes) to distribute the load.