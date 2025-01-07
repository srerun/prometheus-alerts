---
title: JenkinsOutdatedPlugins
description: Troubleshooting for alert JenkinsOutdatedPlugins
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsOutdatedPlugins

## Meaning
[//]: # "Short paragraph that explains what the alert means"
{{ $value }} plugins need update

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: JenkinsOutdatedPlugins
expr: sum(jenkins_plugins_withUpdate) by (instance) > 3
for: 1d
labels:
    severity: warning
annotations:
    summary: Jenkins outdated plugins (instance {{ $labels.instance }})
    description: |-
        {{ $value }} plugins need update
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/JenkinsOutdatedPlugins

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
