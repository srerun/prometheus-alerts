---
title: VaultSealed
description: Troubleshooting for alert VaultSealed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# VaultSealed

Vault instance is sealed on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "hashicorp-vault/hashicorp-vault-internal.yml" "VaultSealed" %}}

{{% comment %}}

```yaml
alert: VaultSealed
expr: vault_core_unsealed == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Vault sealed (instance {{ $labels.instance }})
    description: |-
        Vault instance is sealed on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/hashicorp-vault-internal/VaultSealed.md

```

{{% /comment %}}

</details>


Here is a runbook for the VaultSealed alert:

## Meaning

The VaultSealed alert indicates that a HashiCorp Vault instance has become sealed, meaning it is no longer accessible and all secrets are encrypted. This is a critical issue that requires immediate attention, as it can impact the availability of dependent services and applications.

## Impact

The impact of a sealed Vault instance can be significant, leading to:

* Inaccessible secrets and sensitive data
* Disruption of dependent services and applications
* Downtime and potential revenue loss
* Increased risk of security breaches due to delayed secret rotation and access

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Vault instance logs for errors or warnings that may indicate the cause of the seal.
2. Verify the Vault instance configuration and ensure it is correct and up-to-date.
3. Check the Vault instance's storage backend for any issues or errors.
4. Review the system's overall health and resource utilization to identify any potential bottlenecks or issues.

## Mitigation

To mitigate the VaultSealed alert, follow these steps:

1. Immediately investigate the cause of the seal and take corrective action to unseal the Vault instance.
2. Check the Vault instance's configuration and storage backend to ensure they are correct and functional.
3. Perform a rolling restart of the Vault instance to ensure all nodes are properly unsealed.
4. Verify the dependent services and applications are functioning correctly after the Vault instance is unsealed.
5. Implement additional monitoring and logging to detect potential issues before they cause a seal.
6. Schedule regular maintenance and backups to ensure business continuity in case of a seal.

Remember to update the runbook with specific steps and procedures relevant to your environment and requirements.