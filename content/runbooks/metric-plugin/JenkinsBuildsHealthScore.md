---
title: JenkinsBuildsHealthScore
description: Troubleshooting for alert JenkinsBuildsHealthScore
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsBuildsHealthScore

Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

{{% rule "jenkins/metric-plugin.yml" "JenkinsBuildsHealthScore" %}}

{{% comment %}}

```yaml
alert: JenkinsBuildsHealthScore
expr: default_jenkins_builds_health_score < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Jenkins builds health score (instance {{ $labels.instance }})
    description: |-
        Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/metric-plugin/JenkinsBuildsHealthScore.md

```

{{% /comment %}}

</details>


## Meaning

The JenkinsBuildsHealthScore alert is triggered when the default Jenkins builds health score falls below 1. This score is a measure of the overall health of Jenkins builds across all instances, realms, environments, and regions. A score below 1 indicates that one or more Jenkins builds are experiencing issues, which can impact the reliability and efficiency of the development pipeline.

## Impact

The impact of this alert is high, as it can lead to:

* Delays in code deployment and release
* Increased risk of errors and bugs in production
* Decreased developer productivity
* Inefficiencies in the development pipeline

If not addressed promptly, this issue can cause significant disruptions to the development team and the overall business.

## Diagnosis

To diagnose the root cause of this issue, follow these steps:

1. Check the Jenkins dashboard for any failed or stuck builds
2. Review the Jenkins build logs for errors or warnings
3. Verify that all Jenkins instances are running and healthy
4. Check the network connectivity and communication between Jenkins instances
5. Investigate any recent changes to the Jenkins configuration or plugins

## Mitigation

To mitigate this issue, follow these steps:

1. Identify and cancel any stuck or failed builds
2. Restart the affected Jenkins instance(s)
3. Verify that all Jenkins instances are running and healthy
4. Review and update the Jenkins configuration and plugins as needed
5. Implement additional monitoring and logging to prevent similar issues in the future

For more detailed instructions and troubleshooting steps, refer to the [JenkinsBuildsHealthScore runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/metric-plugin/JenkinsBuildsHealthScore.md).