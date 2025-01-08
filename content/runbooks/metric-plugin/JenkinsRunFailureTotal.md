---
title: JenkinsRunFailureTotal
description: Troubleshooting for alert JenkinsRunFailureTotal
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsRunFailureTotal

Job run failures: ({{$value}}) {{$labels.jenkins_job}}. Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

{{% rule "jenkins/metric-plugin.yml" "JenkinsRunFailureTotal" %}}

{{% comment %}}

```yaml
alert: JenkinsRunFailureTotal
expr: delta(jenkins_runs_failure_total[1h]) > 100
for: 0m
labels:
    severity: warning
annotations:
    summary: Jenkins run failure total (instance {{ $labels.instance }})
    description: |-
        Job run failures: ({{$value}}) {{$labels.jenkins_job}}. Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/metric-plugin/JenkinsRunFailureTotal.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `JenkinsRunFailureTotal`:

## Meaning

The `JenkinsRunFailureTotal` alert is triggered when the total number of Jenkins job run failures exceeds 100 in a 1-hour period. This indicates a significant problem with the Jenkins instance, potentially caused by misconfiguration, resource issues, or application errors.

## Impact

This alert has a significant impact on the development and deployment pipeline, as failed Jenkins jobs can prevent code changes from being built, tested, and deployed. This can lead to delayed releases, reduced productivity, and increased errors. Furthermore, failed jobs can also indicate underlying issues with the application or infrastructure, which if left unchecked, can lead to more severe consequences.

## Diagnosis

To diagnose the root cause of the Jenkins run failures, follow these steps:

1. Check the Jenkins job console output for error messages.
2. Review the Jenkins instance logs for any errors or exceptions.
3. Verify that the Jenkins instance has sufficient resources (e.g., CPU, memory, disk space).
4. Check for any recent changes to the Jenkins configuration, plugins, or job definitions.
5. Investigate any underlying infrastructure issues (e.g., networking, database connectivity).

## Mitigation

To mitigate the impact of Jenkins run failures, follow these steps:

1. Investigate and resolve the root cause of the failures (as identified in the diagnosis step).
2. Restart the failed Jenkins jobs to resume the build and deployment pipeline.
3. Consider increasing the resources allocated to the Jenkins instance to prevent similar failures in the future.
4. Implement additional monitoring and logging to detect and alert on potential issues before they cause significant failures.
5. Verify that the Jenkins instance is properly configured and up-to-date, including all plugins and dependencies.