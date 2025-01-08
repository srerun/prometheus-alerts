---
title: JenkinsLastBuildFailed
description: Troubleshooting for alert JenkinsLastBuildFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsLastBuildFailed

Last build failed: {{$labels.jenkins_job}}. Failed build for job `{{$labels.jenkins_job}}` on {{$labels.instance}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

{{% rule "jenkins/metric-plugin.yml" "JenkinsLastBuildFailed" %}}

{{% comment %}}

```yaml
alert: JenkinsLastBuildFailed
expr: default_jenkins_builds_last_build_result_ordinal == 2
for: 0m
labels:
    severity: warning
annotations:
    summary: Jenkins last build failed (instance {{ $labels.instance }})
    description: |-
        Last build failed: {{$labels.jenkins_job}}. Failed build for job `{{$labels.jenkins_job}}` on {{$labels.instance}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/metric-plugin/JenkinsLastBuildFailed.md

```

{{% /comment %}}

</details>


Here is a runbook for the JenkinsLastBuildFailed alert:

## Meaning

The JenkinsLastBuildFailed alert is triggered when the last build of a Jenkins job fails. This alert indicates that there is an issue with the build process that needs to be addressed to ensure the integrity of the software development lifecycle.

## Impact

The impact of a failed Jenkins build can be significant, leading to:

* Delays in software delivery
* Incomplete or faulty software releases
* Increased risk of errors and bugs in production
* Decreased confidence in the development team's ability to deliver high-quality software

## Diagnosis

To diagnose the root cause of the failed build, follow these steps:

1. Check the Jenkins job logs for errors or exceptions
2. Verify that all dependencies and required plugins are installed and up-to-date
3. Review the build configuration and script for any changes or modifications
4. Check the system resources and infrastructure for any issues or bottlenecks
5. Consult with the development team and subject matter experts to gather more information about the build failure

## Mitigation

To mitigate the failed build, follow these steps:

1. Identify and fix the root cause of the build failure
2. Rerun the failed build to verify that the issue is resolved
3. Update the build configuration and script as necessary to prevent similar failures in the future
4. Consider implementing additional monitoring and logging to detect build failures earlier
5. Communicate with stakeholders about the build failure and the steps being taken to resolve the issue