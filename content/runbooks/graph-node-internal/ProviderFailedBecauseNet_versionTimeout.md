---
title: ProviderFailedBecauseNet_versionTimeout
description: Troubleshooting for alert ProviderFailedBecauseNet_versionTimeout
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProviderFailedBecauseNet_versionTimeout

net_version timeout for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`

<details>
  <summary>Alert Rule</summary>

{{% rule "graph-node/graph-node-internal.yml" "ProviderFailedBecauseNet_versionTimeout" %}}

{{% comment %}}

```yaml
alert: ProviderFailedBecauseNet_versionTimeout
expr: eth_rpc_status == 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Provider failed because net_version timeout (instance {{ $labels.instance }})
    description: |-
        net_version timeout for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/graph-node-internal/ProviderFailedBecauseNet_versionTimeout.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The `ProviderFailedBecauseNet_versionTimeout` alert is triggered when the `eth_rpc_status` metric returns a value of 3, indicating that a provider has failed due to a net_version timeout. This alert is critical and requires immediate attention.

## Impact

The impact of this alert is that the provider is unable to function correctly, which can lead to:

* Disruption of critical services dependent on the provider
* Loss of data or transactions
* Inconsistent network state
* Inability to retrieve or update information from the provider

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `eth_rpc_status` metric to confirm that it is still returning a value of 3.
2. Investigate the provider logs for errors related to net_version timeouts.
3. Verify that the provider is configured correctly and that there are no issues with the underlying network or infrastructure.
4. Check the Grafana dashboard for any other related alerts or issues.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the provider service to attempt to recover from the timeout.
2. Check the provider configuration and adjust any settings that may be contributing to the timeout.
3. Investigate and resolve any underlying network or infrastructure issues.
4. If the issue persists, consider increasing the timeout value or implementing retry logic to improve the provider's resilience.
5. Notify the relevant teams and stakeholders of the issue and ensure that it is being actively worked on.

Remember to check the alert's annotations for specific details about the affected provider and instance, and refer to the linked runbook for further guidance.