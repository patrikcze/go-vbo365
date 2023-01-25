[![Build Status](https://travis-ci.org/patrikcze/go-vbo365.svg?branch=master)](https://travis-ci.org/patrikcze/go-vbo365)
[![GitHub version](https://badge.fury.io/gh/patrikcze%2Fgo-vbo365.svg)](https://badge.fury.io/gh/patrikcze%2Fgo-vbo365)
[![Lint](https://github.com/patrikcze/go-vbo365/actions/workflows/lint.yaml/badge.svg)](https://github.com/patrikcze/go-vbo365/actions/workflows/lint.yaml)
# go-vbo365
This project is about to rewrite Bash script which has been developed by **jorgedlcruz**  into Go Lang. Possibly to be able to deploy this Go Binary into Docker Container and deploy complete solution to Docker or Kubernetes.  Let's see how well it goes.

# Source Jorge's Bash script description

The Bash script [**jorgedlcruz**](https://github.com/jorgedlcruz/veeam-backup-for-microsoft365-grafana/blob/e393ddb3c55c7d3568cc5d01ac9d02712a6024f1/veeam_microsoft365.sh) provided collects various information from the `Veeam Backup for Microsoft 365 REST API` and sends it to `InfluxDB`. The script is divided into several sections, each of which makes a different type of API call to the `Veeam server` to retrieve information.

- **`The first section`** retrieves an access token from the Veeam server using a `curl` command that makes a `POST` request to the token endpoint with the specified `username` and `password`. The `token` is then used to authenticate all subsequent API requests.

- **`The second section`** retrieves the `version` of `Veeam Backup for Microsoft 365` using a GET request to the service instance endpoint, and sends the version to `InfluxDB`.

- **`The third section`** retrieves information about the `organization` and `licensing` information, it makes a GET request to the organizations endpoint, and iterates over the results to retrieve the `id`, `name`, `licensedUsers`, and `newUsers` for `each organization`, and sends this information to `InfluxDB`.

- **`The fourth section`** retrieves information about the `Backup Proxies`, it makes a GET request to the proxies endpoint, and iterates over the results to retrieve the `hostName`, `threadsNumber`, and `status` of each proxy and sends this information to `InfluxDB`.

- **`The fifth section`** of the script retrieves detailed information about the `Backup Proxies`. It sets the URL for the REST API call, retrieves the JSON response, and parses it to extract the `hostName`, `threadsNumber`, and status for each proxy. It then sends this information to `InfluxDB` using the same method as the previous sections.

- **`The sixth section`** of the script retrieves information about the `Backup Jobs`. It sets the URL for the REST API call, retrieves the JSON response, and parses it to extract the name and id of each job. It then makes another REST API call to retrieve information about the `Job Sessions` for each `job`, and parses the JSON response to extract the `status`, `creationTime`, and `endTime` of each session. It then sends this information to `InfluxDB` using the same method as the previous sections.

- **`The seventh section`** of the script retrieves information about the `RBAC Roles`. It sets the URL for the REST API call, retrieves the JSON response, and parses it to extract the `id`, `name`, `description`, `roleType`, and `organizationId` of each `role`. It then makes another REST API call to retrieve the organization name associated with the role and sends this information to `InfluxDB` using the same method as the previous sections. It then makes another REST API call to retrieve information about the RBAC `Role Operators`, and parses the JSON response to extract the `type` and `name` of each operator. It then sends this information to `InfluxDB` using the same method as the previous sections.

**`Overall`**, the Bash script is designed to collect various types of data related to Veeam Backup for Microsoft 365 service. The data is collected via a series of REST API calls to the Veeam Backup for Microsoft 365 server using the `curl` command. The collected data is then parsed using `jq` and the `awk` command, and sent to an `InfluxDB database` for storage and `visualization`. The script is divided into several sections, each of which is responsible for collecting a specific type of data, such as version information, `licensing` information, `backup proxy information`, `backup job information`, and `RBAC role` information.

One of the key things to note is that the script **`relies heavily`** on `environmental variables` to configure the various parameters needed to connect to the `Veeam Backup for Microsoft 365 server` and the `InfluxDB` database. This includes the server URL and port, username and password, InfluxDB URL and port, and InfluxDB token and organization.

The script also makes use of a number of command line tools such as `jq`, `awk`, `curl`, and `echo` to parse and manipulate the data, and `for` and `if` statements to iterate over the collected data and perform different actions based on the data.

In terms of optimization, the script could be improved by reducing the number of API calls made to the Veeam Backup for Microsoft 365 server and by making use of more efficient data manipulation techniques. Additionally, error handling and logging could be added to improve the script's robustness and troubleshooting capabilities.

# Project Goal

1. Goal here is to rewrite script into `Go` Language (i'm total beginner ;-)).
2. Avoid using `jq`, `awk`, `curl` or other cmd line tools, to have it run purely as a `binary`.
3. Avoid using `Crontab` or K8S `crontabjobs` (eg. implement `time.Tick` function).
4. Stick with current setup using `Environment variables` (eg. can be helpful in case of K8S deployment)
    - Create configuration file with parameters for `"Schedule"` every 30minutes, 1h, or so.
    - Also make some of the `Environment variables` being able to be configured within configuration file (eg. Binary would or could run Stand-alone)
5. `Build` an image with this `Go Binary` (using `multi-stage` docker image build process)
    - save an amount of space (possibly use some basic Alpine Linux images)
    - consider using it as `commandline tool` too (eg. Using Go `Cobra/Viper`)
6. Use type `Sructures` where possible instead of parsing `JSON`.
7. Make it `updatable` in terms of REST API changes on both sides `Veeam M365 API` and `InfluxDB API`. - *Nice-to-have*
8. Implement `Go` Code `lint` and some basic workflows for `GitHub` and `Gitlab`. - *Nice-to-have*
9. Maybe later on, implement individual `Visualization` purely with `Go` over some REST API Call. (eg. `Go-Echarts`) - *Nice-to-have*
10. Include Deployment of `InfluxDB 2.x` and `Grafana` too with complete configuration (Out-of-the-box solution).  - *Good-to-have*


# Updates :

- 25.1.2023 - started to work.
