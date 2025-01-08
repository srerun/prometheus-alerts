---
title: VaultTooManyInfinityTokens
description: Troubleshooting for alert VaultTooManyInfinityTokens
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# VaultTooManyInfinityTokens

Too many infinity tokens {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "hashicorp-vault/hashicorp-vault-internal.yml" "VaultTooManyInfinityTokens" %}}

{{% comment %}}

```yaml
alert: VaultTooManyInfinityTokens
expr: vault_token_count_by_ttl{creation_ttl="+Inf"} > 3
for: 5m
labels:
    severity: warning
annotations:
    summary: Vault too many infinity tokens (instance {{ $labels.instance }})
    description: |-
        Too many infinity tokens {{ $labels.instance }}: {{ $value | printf "%.2f"}}%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/hashicorp-vault-internal/VaultTooManyInfinityTokens.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "VaultTooManyInfinityTokens":

## Meaning

The "VaultTooManyInfinityTokens" alert is triggered when more than 3 infinity tokens are present in Vault. Infinity tokens are tokens with a creation TTL (time to live) set to infinity, meaning they never expire. This alert indicates a potential security risk, as an excessive number of infinity tokens can lead to unauthorized access to sensitive resources.

## Impact

The presence of too many infinity tokens in Vault can have the following impacts:

* Increased risk of unauthorized access to sensitive resources
* Complexity in token management and revocation
* Potential security breaches due to uncontrolled token usage

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Vault token count: Run the command `vault token count` to get the total number of tokens in Vault.
2. Identify infinity tokens: Run the command `vault token list -format=json | jq '.[] | select(.creation_ttl == "+Inf")'` to get a list of infinity tokens.
3. Check token usage: Analyze the token usage patterns to identify potential misuse or abuse.

## Mitigation

To mitigate the issue, follow these steps:

1. Revoke excess infinity tokens: Use the `vault token revoke` command to revoke the excess infinity tokens.
2. Implement token TTL: Configure a sensible TTL for tokens to ensure they expire after a reasonable period.
3. Implement token counting and alerting: Set up monitoring and alerting to detect and notify about excessive token creation.
4. Review and adjust token creation policies: Review and adjust token creation policies to ensure they align with security best practices.

Remember to investigate the root cause of the issue and take necessary steps to prevent similar incidents in the future.