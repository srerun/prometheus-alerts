---
title: JenkinsBuildTestsFailing
description: Troubleshooting for alert JenkinsBuildTestsFailing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsBuildTestsFailing

Last build tests failed: {{$labels.jenkins_job}}. Failed build Tests for job `{{$labels.jenkins_job}}` on {{$labels.instance}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

{{% rule "jenkins/metric-plugin.yml" "JenkinsBuildTestsFailing" %}}

{{% comment %}}

```yaml
alert: JenkinsBuildTestsFailing
expr: default_jenkins_builds_last_build_tests_failing > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Jenkins build tests failing (instance {{ $labels.instance }})
    description: |-
        Last build tests failed: {{$labels.jenkins_job}}. Failed build Tests for job `{{$labels.jenkins_job}}` on {{$labels.instance}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/metric-plugin/JenkinsBuildTestsFailing.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule "JenkinsBuildTestsFailing":

## Meaning

The "JenkinsBuildTestsFailing" alert is triggered when the number of failing tests in the last build of a Jenkins job exceeds 0. This indicates that there is an issue with the build or the tests themselves, which can impact the quality and reliability of the software being built.

## Impact

The impact of this alert is that the software build is not reliable, and the failing tests may indicate underlying issues that need to be addressed. This can lead to:

* Delays in software releases
* Lower quality software
* Increased risk of defects and bugs
* Decreased confidence in the software development process

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Jenkins job that triggered the alert and review the build logs to identify the failing tests.
2. Analyze the test failures to determine the root cause of the issue.
3. Check the Jenkins job configuration and the test scripts to ensure they are correct and up-to-date.
4. Verify that the build environment and dependencies are correct and up-to-date.
5. Check the Jenkins instance and environment to ensure they are functioning correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Fix the failing tests by updating the test scripts or the Jenkins job configuration.
2. Rerun the build to ensure the tests pass.
3. Verify that the build is successful and the software is reliable.
4. Implement monitoring and testing strategies to catch similar issues earlier in the development process.
5. Consider implementing automated testing and CI/CD pipelines to improve the software development process.

Remember to update the annotations and labels in the Prometheus alert rule to reflect the root cause of the issue and the steps taken to mitigate it.