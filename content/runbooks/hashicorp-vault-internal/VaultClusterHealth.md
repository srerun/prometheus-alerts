---
title: VaultClusterHealth
description: Troubleshooting for alert VaultClusterHealth
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# VaultClusterHealth

Vault cluster is not healthy {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "hashicorp-vault/hashicorp-vault-internal.yml" "VaultClusterHealth" %}}

{{% comment %}}

```yaml
alert: VaultClusterHealth
expr: sum(vault_core_active) / count(vault_core_active) <= 0.5
for: 0m
labels:
    severity: critical
annotations:
    summary: Vault cluster health (instance {{ $labels.instance }})
    description: |-
        Vault cluster is not healthy {{ $labels.instance }}: {{ $value | printf "%.2f"}}%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/hashicorp-vault-internal/VaultClusterHealth.md

```

{{% /comment %}}

</details>


Here is a runbook for the VaultClusterHealth alert rule:

## Meaning

The Vault cluster health alert indicates that the Vault cluster is not in a healthy state. This is determined by checking the ratio of active Vault cores to the total number of Vault cores. If the ratio falls below 50%, the alert is triggered.

## Impact

A unhealthy Vault cluster can have significant implications on the overall security and reliability of the system. It can lead to:

* Unavailability of secrets and sensitive data
* Increased risk of data breaches
* Disruption to dependent applications and services
* Potential loss of business critical data

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Vault cluster status using the Vault CLI or Web UI
2. Verify that the Vault cores are properly configured and running
3. Check the system logs for any errors or warnings related to Vault
4. Investigate any recent changes or deployments that may have caused the issue
5. Review the Vault cluster topology and ensure it is correctly configured

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the root cause of the unhealthy cluster state
2. Take corrective action to restore the cluster to a healthy state
	* This may involve restarting Vault cores, fixing configuration issues, or rolling back recent changes
3. Verify that the cluster is healthy by checking the status and monitoring metrics
4. Perform a thorough review of the Vault cluster configuration and topology to prevent similar issues in the future
5. Consider implementing additional monitoring and alerting to detect potential issues earlier