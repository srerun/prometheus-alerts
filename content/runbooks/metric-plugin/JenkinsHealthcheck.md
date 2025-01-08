---
title: JenkinsHealthcheck
description: Troubleshooting for alert JenkinsHealthcheck
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsHealthcheck

Jenkins healthcheck score: {{$value}}. Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

{{% rule "jenkins/metric-plugin.yml" "JenkinsHealthcheck" %}}

{{% comment %}}

```yaml
alert: JenkinsHealthcheck
expr: jenkins_health_check_score < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Jenkins healthcheck (instance {{ $labels.instance }})
    description: |-
        Jenkins healthcheck score: {{$value}}. Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/metric-plugin/JenkinsHealthcheck.md

```

{{% /comment %}}

</details>


## Meaning

The JenkinsHealthcheck alert is triggered when the Jenkins health check score falls below 1, indicating that the Jenkins instance is experiencing health issues. This alert is critical and should be addressed promptly to prevent disruption to the development pipeline.

## Impact

The impact of this alert is high, as a unhealthy Jenkins instance can cause:

* Delays in code builds and deployments
* Failure of automated tests and validation
* Inability to monitor and track changes to the codebase
* Potential data loss or corruption

If left unaddressed, this issue can lead to significant disruptions to the development team and potentially impact the overall quality of the product.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Jenkins UI to identify the specific health check that is failing.
2. Review the Jenkins logs to gather more information about the error.
3. Verify that the Jenkins instance has sufficient resources (CPU, memory, and disk space) to operate properly.
4. Check for any recent changes to the Jenkins configuration or plugins that may be causing the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Jenkins instance to reset the health check score.
2. Investigate and resolve the underlying issue causing the health check failure.
3. Verify that the Jenkins instance is properly configured and that all required plugins are up-to-date.
4. Consider increasing the resources allocated to the Jenkins instance if it is experiencing high load or usage.
5. Monitor the Jenkins instance closely to ensure that the health check score returns to a normal level.

Remember to refer to the JenkinsHealthcheck runbook for more detailed instructions and troubleshooting steps.