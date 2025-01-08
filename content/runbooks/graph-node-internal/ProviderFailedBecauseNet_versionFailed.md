---
title: ProviderFailedBecauseNet_versionFailed
description: Troubleshooting for alert ProviderFailedBecauseNet_versionFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProviderFailedBecauseNet_versionFailed

Failed net_version for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`

<details>
  <summary>Alert Rule</summary>

{{% rule "graph-node/graph-node-internal.yml" "ProviderFailedBecauseNet_versionFailed" %}}

{{% comment %}}

```yaml
alert: ProviderFailedBecauseNet_versionFailed
expr: eth_rpc_status == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Provider failed because net_version failed (instance {{ $labels.instance }})
    description: |-
        Failed net_version for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/graph-node-internal/ProviderFailedBecauseNet_versionFailed.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

This alert is triggered when the `eth_rpc_status` metric returns a value of 1, indicating that the Provider has failed because of a failed `net_version` check. This is a critical alert, as it affects the functionality of the Graph node.

## Impact

The impact of this alert is that the Provider is unable to function correctly, which can lead to issues with data retrieval and processing. This can have a cascading effect on dependent services and applications, causing data inconsistencies and errors.

## Diagnosis

To diagnose the root cause of this issue, follow these steps:

1. Check the Graph node logs for errors related to `net_version` checks.
2. Verify that the Provider is configured correctly and that the `net_version` API is accessible.
3. Check the `eth_rpc_status` metric to see if there are any trends or patterns that can indicate the cause of the failure.
4. Investigate any recent changes to the Provider configuration or the Graph node environment that may have caused the issue.

## Mitigation

To mitigate this issue, follow these steps:

1. Check the Provider configuration and ensure that it is correct and up-to-date.
2. Restart the Provider service to try to recover from the failed `net_version` check.
3. If the issue persists, try to debug the `net_version` API to identify the root cause of the failure.
4. If necessary, roll back any recent changes to the Provider configuration or Graph node environment to a known good state.
5. If the issue is still not resolved, escalate to the development team for further investigation and resolution.