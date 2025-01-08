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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many pending tokens {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "hashicorp-vault/hashicorp-vault-internal.yml" "VaultTooManyPendingTokens" %}}

<!-- Rule when generated

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

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
