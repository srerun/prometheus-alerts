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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
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


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
