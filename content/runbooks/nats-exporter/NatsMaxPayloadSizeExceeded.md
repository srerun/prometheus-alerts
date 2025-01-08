---
title: NatsMaxPayloadSizeExceeded
description: Troubleshooting for alert NatsMaxPayloadSizeExceeded
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsMaxPayloadSizeExceeded

The max payload size allowed by NATS has been exceeded (1MB)

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsMaxPayloadSizeExceeded" %}}

{{% comment %}}

```yaml
alert: NatsMaxPayloadSizeExceeded
expr: max(gnatsd_varz_max_payload) > 1024 * 1024
for: 5m
labels:
    severity: critical
annotations:
    summary: Nats max payload size exceeded (instance {{ $labels.instance }})
    description: |-
        The max payload size allowed by NATS has been exceeded (1MB)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsMaxPayloadSizeExceeded.md

```

{{% /comment %}}

</details>


## Meaning

The `NatsMaxPayloadSizeExceeded` alert is triggered when the maximum payload size allowed by NATS exceeds 1MB. This alert indicates that the NATS messaging system is receiving messages with payloads larger than the allowed limit, which can lead to performance issues, message loss, and even system instability.

## Impact

* Performance degradation: Excessive payload sizes can cause NATS to slow down, leading to slower message processing and potentially causing backlogs.
* Message loss: If the payload size exceeds the allowed limit, messages may be dropped, resulting in data loss and potential system instability.
* System instability: Prolonged exposure to oversized payloads can cause the NATS system to become unstable, leading to crashes or restarts.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NATS server logs for error messages related to oversized payloads.
2. Verify that the `gnatsd_varz_max_payload` metric is accurately reporting the maximum payload size.
3. Investigate the source of the oversized payloads, which may be due to:
	* Misconfigured clients or publishers
	* Incorrectly formatted messages
	* Unintended data growth
4. Review the NATS configuration to ensure that the maximum payload size is set correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and fix the source of the oversized payloads:
	* Adjust client or publisher settings to adhere to the maximum payload size limit.
	* Implement payload compression or splitting to reduce message size.
2. Adjust the NATS configuration to increase the maximum payload size limit, if necessary.
3. Monitor the `gnatsd_varz_max_payload` metric to ensure the issue is resolved and prevent future occurrences.
4. Consider implementing payload size monitoring and alerting to detect potential issues proactively.

For more detailed information and guidance, refer to the runbook linked in the alert annotation: <https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsMaxPayloadSizeExceeded.md>