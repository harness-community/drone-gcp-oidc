# drone-gcp-oidc

- [Synopsis](#Synopsis)
- [Parameters](#Parameters)
- [Notes](#Notes)
- [Plugin Image](#Plugin-Image)
- [Examples](#Examples)

## Synopsis

This plugin generates an access token through the OIDC token and outputs it as an environment variable. This variable can be utilized in subsequent pipeline steps to control Google Cloud Services through the gcloud CLI or API using curl.

To learn how to utilize Drone plugins in Harness CI, please consult the provided [documentation](https://developer.harness.io/docs/continuous-integration/use-ci/use-drone-plugins/run-a-drone-plugin-in-ci).

## Parameters

| Parameter                                                                                                                              | Choices/<span style="color:blue;">Defaults</span> | Comments                                                        |
| :------------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------ | --------------------------------------------------------------- |
| project_id <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>               |                                                   | The project id associated with your GCP project.                |
| pool_id <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>                  |                                                   | The pool ID for OIDC authentication.                            |
| provider_id <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>              |                                                   | The provider ID for OIDC authentication.                        |
| service_account_email_id <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span> |                                                   | The email address of the service account.                       |
| duration <span style="font-size: 10px"><br/>`string`</span>                                                                            | Default: `3600`                                   | The lifecycle duration of the access token generated in seconds |
| create_application_credentials_file <span style="font-size: 10px"><br/>`boolean`</span>                                                | Default: `false`                                  | Create application_default_credentials.json                     |

## Notes

This plugin also requires an OIDC token `PLUGIN_OIDC_TOKEN_ID`, provided as a stage variable.

Please provide the `duration` in seconds, for example, the default value is 1 hour, i.e, 3600 seconds. The service account must have the `iam.allowServiceAccountCredentialLifetimeExtension` permission to set a custom duration.

The plugin created `application_default_credentials.json` if the `create_application_credentials_file` flag is set to `true`, only to be used with `gradlew`, for use with `gcloud cli`, use `GCLOUD_ACCESS_TOKEN`. The `credentials.json` path is exported to `$HOME/.config/gcloud/application_default_credentials.json` variable, which can be used as `<+steps.STEP_ID.output.outputVariables.GOOGLE_APPLICATION_CREDENTIALS>`

The plugin outputs the access token in the form of an environment variable that can be accessed in the subsequent pipeline steps like this: `<+steps.STEP_ID.output.outputVariables.GCLOUD_ACCESS_TOKEN>`

## Plugin Image

The plugin `plugins/gcp-oidc` is available for the following architectures:

| OS            | Tag                                |
| ------------- | ---------------------------------- |
| latest        | `linux-amd64/arm64, windows-amd64` |
| linux/amd64   | `linux-amd64`                      |
| linux/arm64   | `linux-arm64`                      |
| windows/amd64 | `windows-amd64`                    |

## Examples

```
# Plugin YAML
- step:
    type: Plugin
    name: drone-gcp-oidc-plugin
    identifier: drone_gcp_oidc_plugin
    spec:
        connectorRef: harness-docker-connector
        image: plugins/gcp-oidc
        settings:
                project_id: 22819301
                pool_id: d8291ka22
                service_account_email_id: test-gcp@harness.io
                provider_id: svr-account1

- step:
    type: Plugin
    name: drone-gcp-oidc-plugin
    identifier: drone_gcp_oidc_plugin
    spec:
        connectorRef: harness-docker-connector
        image: plugins/gcp-oidc
        settings:
                project_id: 22819301
                pool_id: d8291ka22
                service_account_email_id: test-gcp@harness.io
                provider_id: svr-account1
                duration: 7200

- step:
    type: Plugin
    name: drone-gcp-oidc-plugin
    identifier: drone_gcp_oidc_plugin
    spec:
        connectorRef: harness-docker-connector
        image: plugins/gcp-oidc
        settings:
                project_id: 22819301
                pool_id: d8291ka22
                service_account_email_id: test-gcp@harness.io
                provider_id: svr-account1
                duration: 7200
                create_application_credentials_file: true

# Run step to use the access token to list the compute zones
- step:
    type: Run
    name: List Compute Engine Zone
    identifier: list_zones
    spec:
        shell: Sh
        command: |-
            curl -H "Authorization: Bearer <+steps.STEP_ID.output.outputVariables.GCLOUD_ACCESS_TOKEN>" \
            "https://compute.googleapis.com/compute/v1/projects/[PROJECT_ID]/zones/[ZONE]/instances"
```

> <span style="font-size: 14px; margin-left:5px; background-color: #d3d3d3; padding: 4px; border-radius: 4px;">ℹ️ If you notice any issues in this documentation, you can [edit this document](https://github.com/harness-community/drone-gcp-oidc/blob/main/README.md) to improve it.</span>
