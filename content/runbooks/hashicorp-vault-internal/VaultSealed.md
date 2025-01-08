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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
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


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
