---
title: VaultTooManyPendingTokens
description: Troubleshooting for alert VaultTooManyPendingTokens
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# VaultTooManyPendingTokens

Too many pending tokens {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "hashicorp-vault/hashicorp-vault-internal.yml" "VaultTooManyPendingTokens" %}}

{{% comment %}}

```yaml
alert: VaultTooManyPendingTokens
expr: avg(vault_token_create_count - vault_token_store_count) > 0
for: 5m
labels:
    severity: warning
annotations:
    summary: Vault too many pending tokens (instance {{ $labels.instance }})
    description: |-
        Too many pending tokens {{ $labels.instance }}: {{ $value | printf "%.2f"}}%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/hashicorp-vault-internal/VaultTooManyPendingTokens.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "VaultTooManyPendingTokens":

## Meaning

The "VaultTooManyPendingTokens" alert is triggered when the average difference between the number of token creations and token stores in Vault exceeds 0 over a 5-minute period. This indicates that there are too many pending tokens in Vault, which can lead to performance issues and potential security risks.

## Impact

If left unaddressed, this issue can cause:

* Performance degradation in Vault and dependent systems
* Increased latency for token requests
* Potential security risks due to unauthorized access or token leakage
* In extreme cases, Vault may become unavailable or crash due to excessive memory usage

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Vault logs for any errors or warnings related to token creation or storage.
2. Verify that the `vault_token_create_count` and `vault_token_store_count` metrics are correct and up-to-date.
3. Investigate any recent changes to Vault configurations, plugins, or dependent systems that may be causing the issue.
4. Check the Vault instance's resource utilization (CPU, memory, disk space) to ensure it has sufficient capacity.

## Mitigation

To mitigate the issue, follow these steps:

1. Reduce the load on Vault by identifying and addressing any excessive token requests or misconfigured clients.
2. Verify that Vault is properly configured for token storage and rotation.
3. Consider increasing the resources (CPU, memory, disk space) allocated to the Vault instance.
4. Implement additional monitoring and logging to detect and prevent similar issues in the future.
5. If necessary, consult the Vault documentation and HashiCorp support resources for further guidance.

Remember to update the `vault_token_create_count` and `vault_token_store_count` metrics to reflect any changes made to Vault configurations or dependent systems.