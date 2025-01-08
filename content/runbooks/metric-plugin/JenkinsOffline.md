---
title: JenkinsOffline
description: Troubleshooting for alert JenkinsOffline
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsOffline

Jenkins offline: `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

{{% rule "jenkins/metric-plugin.yml" "JenkinsOffline" %}}

{{% comment %}}

```yaml
alert: JenkinsOffline
expr: jenkins_node_offline_value > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Jenkins offline (instance {{ $labels.instance }})
    description: |-
        Jenkins offline: `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/metric-plugin/JenkinsOffline.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "JenkinsOffline":

## Meaning

The "JenkinsOffline" alert is triggered when the `jenkins_node_offline_value` metric exceeds a value of 1, indicating that one or more Jenkins instances are offline. This alert is critical, as it may impact the availability of Jenkins services and the ability to execute build jobs.

## Impact

* Jenkins services may be unavailable, preventing users from executing build jobs and accessing Jenkins functionality.
* Builds may fail or stall, leading to delays in software development and deployment.
* The unavailability of Jenkins may also impact dependent services and systems that rely on Jenkins for automation and orchestration.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Jenkins instance specified in the alert label `instance` and verify its status. Is the instance running, and are all required services operational?
2. Review Jenkins logs to identify any errors or issues that may be contributing to the offline state.
3. Verify that the Jenkins instance has connectivity to required dependencies, such as databases, file systems, and networking resources.
4. Check for any recent changes or updates to Jenkins plugins, configuration, or environment that may be causing the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Attempt to restart the Jenkins instance or service to see if it can be recovered.
2. Check Jenkins logs for any error messages or exceptions that may indicate the root cause of the issue.
3. Verify that all required Jenkins plugins and dependencies are up-to-date and functional.
4. If the issue persists, consider escalating to the Jenkins administrator or a senior team member for further assistance and troubleshooting.

Remember to refer to the original alert definition and labels for specific details about the affected Jenkins instance, realm, environment, and region.