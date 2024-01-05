# Introducing the drone-gcp-oidc Plugin

At Harness, we are committed to enhancing Continuous Integration (CI) and Continuous Deployment (CD) processes by providing tools that simplify complex workflows. We understand the importance of seamlessly integrating Google Cloud Platform (GCP) OpenID Connect (OIDC) authentication with your CI/CD pipelines. That's why we are excited to introduce the **drone-gcp-oidc** plugin. This plugin facilitates the retrieval of Google Cloud access tokens, streamlining the authentication process in your CI/CD workflows.

### What is the drone-gcp-oidc plugin?

The **drone-gcp-oidc plugin** is a versatile tool designed to simplify the integration of GCP OIDC authentication with your CI/CD pipelines. This plugin automates the process of obtaining Google Cloud access tokens, allowing you to seamlessly authenticate with GCP services during your pipeline executions.

### Build the Docker Image

Using the plugin is straightforward. You can run the plugin directly using the following command:

    PLUGIN_OIDC_ID_TOKEN=OIDC_ID_TOKEN \
    PLUGIN_PROJECT_NUMBER=PROJECT_NUMBER \
    PLUGIN_POOL_ID=POOL_ID \
    PLUGIN_PROVIDER_ID=PROVIDER_ID \
    PLUGIN_SERVICE_ACCOUNT_EMAIL=SERVICE_ACCOUNT_EMAIL \
    DRONE_OUTPUT=DRONE_OUTPUT \
    go run main.go

Additionally, you can build the Docker image with these commands:

    docker buildx build -t DOCKER_ORG/drone-maven-version-docker-build --platform linux/amd64 .

### Usage in Harness CI

Integrating the drone-gcp-oidc Plugin into your Harness CI pipeline is seamless. You can use Docker to run the plugin with environment variables. Here's how:

    docker run --rm \
    -e PLUGIN_OIDC_ID_TOKEN=${OIDC_ID_TOKEN} \
    -e PLUGIN_PROJECT_NUMBER=${PROJECT_NUMBER} \
    -e PLUGIN_POOL_ID=${POOL_ID} \
    -e PLUGIN_PROVIDER_ID=${PROVIDER_ID} \
    -e PLUGIN_SERVICE_ACCOUNT_EMAIL=${SERVICE_ACCOUNT_EMAIL} \
    -e DRONE_OUTPUT=${DRONE_OUTPUT} \
    -v $(pwd):$(pwd) \
    -w $(pwd) \
    harnesscommunity/drone-gcp-oidc

In your Harness CI pipeline, you can define the plugin as a step, like this:

    - step:
        type: Plugin
        name: drone-gcp-oidc-plugin
        identifier: gcp_oidc_plugin
        spec:
            connectorRef: docker-connector
            image: harnesscommunity/drone-gcp-oidc
            settings:
                oidc_id_token: your-oidc-id-token
                project_number: your-project-number
                pool_id: your-pool-id
                provider_id: your-provider-id
                service_account_email: your-service-account-email

### Plugin Options

The drone-gcp-oidc Plugin offers the following customization options:

- **oidc_id_token**: The OIDC ID token used for authentication.

- **project_number**: The project number associated with your GCP project.

- **pool_id**: The pool ID for OIDC authentication.

- **provider_id**: The provider ID for OIDC authentication.

- **service_account_email**: The email address of the service account.

These environment variables are crucial for configuring and customizing the behavior of the drone-gcp-oidc plugin when executed as a Docker container. They allow you to provide specific values required for obtaining the Google Cloud access token.

### Get Started with the GCP OIDC Plugin

Whether you are an experienced DevOps professional or new to CI/CD, the drone-gcp-oidc plugin can simplify your GCP authentication process. Give it a try and witness how it streamlines your CI/CD pipelines!

For more information, documentation, and updates, please visit our GitHub repository: [drone-gcp-oidc](https://github.com/harness-community/drone-gcp-oidc).
