---
title: FreeswitchSessionsCritical
description: Troubleshooting for alert FreeswitchSessionsCritical
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeswitchSessionsCritical

High sessions usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "freeswitch/znerol-freeswitch-exporter.yml" "FreeswitchSessionsCritical" %}}

{{% comment %}}

```yaml
alert: FreeswitchSessionsCritical
expr: (freeswitch_session_active * 100 / freeswitch_session_limit) > 90
for: 5m
labels:
    severity: critical
annotations:
    summary: Freeswitch Sessions Critical (instance {{ $labels.instance }})
    description: |-
        High sessions usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/znerol-freeswitch-exporter/FreeswitchSessionsCritical.md

```

{{% /comment %}}

</details>


## Meaning

The FreeswitchSessionsCritical alert is triggered when the percentage of active FreeSWITCH sessions exceeds 90% of the total session limit. This indicates that the system is experiencing high usage and may be at risk of exhausting available resources.

## Impact

If left unaddressed, high session usage can lead to:

* Degraded system performance
* Increased latency and response times
* Potential crashes or failures of the FreeSWITCH server
* Impact on overall system stability and reliability

## Diagnosis

To diagnose the root cause of the high session usage, follow these steps:

1. Check the FreeSWITCH server logs for any errors or anomalies
2. Verify that the session limit is set correctly and adjust if necessary
3. Investigate any recent changes to the system or configuration that may be contributing to the high usage
4. Review the FreeSWITCH debug logs to identify any specific issues or bottlenecks
5. Check the system resource utilization (CPU, memory, etc.) to ensure it is within acceptable limits

## Mitigation

To mitigate the high session usage, follow these steps:

1. Increase the session limit if possible
2. Identify and terminate any idle or unnecessary sessions
3. Optimize system configuration and tuning to improve performance and reduce resource utilization
4. Implement rate limiting or other traffic control measures to prevent sudden spikes in session usage
5. Consider load balancing or distributing traffic across multiple FreeSWITCH servers to reduce the load on individual instances