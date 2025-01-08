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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Vault cluster is not healthy {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "hashicorp-vault/hashicorp-vault-internal.yml" "VaultClusterHealth" %}}

<!-- Rule when generated

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

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
