# Container Security Checklist
- Ignite talk slides: [https://bit.ly/container-security-slides](https://bit.ly/container-security-slides)

### My Other Related Repos 🔥
- [Shift Left (Preventive) Security for Containers/Kubernetes](https://github.com/gunjan5/shift-left-security)
- [Container Security Checklist with Risk Scores](https://github.com/gunjan5/cloud-native-security)
- [Fun: Container from Scratch](https://github.com/gunjan5/container-from-scratch)

![](security-integration-points.png)

## Intro

Container technologies are enabling new application architectures for legacy apps, refactored apps, and microservices, among others. However, these new application architectures also introduce new attack vectors including control plane attacks against the orchestrator, network-based attacks across the infrastructure, container repository attacks, and host operating system attacks. According to a December 2018 Datadog survey, containers are now churning 12 times faster than VMs and have an average lifetime of just 12 hours. As a result, container security is a top challenge. In a recent Tripwire survey, 98 percent of respondents felt they needed additional container security capabilities, 75 percent of respondents with more than 100 containers in production had reported an incident, and 42 percent had either delayed or limited container adoption due to security concerns. Although containers are inherently more secure than many other technologies by design, you need to take additional steps to fully secure your container deployments throughout their lifecycle. 

## Secure the Build Pipeline

- [x] Verify the image source (registry)
- [x] Use official base images 
- [x] Lock down access to the image registry (who can push/pull) 
- [x] Scan container image layers for Common Vulnerabilities and Exposures (CVEs) 
- [x] Scan configuration files for security and compliance checks in continuous integration (CI)
- [x] Do a static analysis of the code and libraries used by the code to surface any vulnerabilities in the code and its dependencies
- [x] Tag and automatically prevent vulnerable images from running in certain clusters or prevent them from talking to other containers in the cluster 


## Secure the Host
- [x] Lock down the operating system (like Google’s Container Optimized OS (COS)) 
- [x] Use secure computing (seccomp) to restrict host system call (syscall) access from containers
- [x] Use Security-Enhanced Linux (SELinux) to further isolate containers 
- [x] Utilize container sandboxing projects like gVisor, Kata Containers, etc. to reduce the attack surface


## Secure the Container Runtimes and Configure Admission Control

- [x] Ensure security configurations span across container runtimes, especially if the environment has multiple container runtimes in the cluster (for example, different runtimes for the orchestrator control plane and workloads)
- [x] Use policies (for example, pod security policy in Kubernetes) to restrict which containers can run in the cluster including policies to restrict privileged containers, containers that don’t need write access to a specific volume, and containers that need certain syscalls
- [x] Restrict access to container runtime daemon/APIs


## Secure the Network 

- [x] Secure services that are exposed to the Internet using a firewall 
- [x] Lock down Layer 3 and 4 access for the services using network policy
- [x] Create granular Layer 7 policies using service mesh (such as Istio)
- [x] Use Mutual Transport Layer Security (mTLS) to mutually authenticate containerized workloads (for example, using Istio)
- [x] Segregate containerized workloads with a mix of host segregation and network isolation (for example, separate group of hosts for workload segregation and/or network policies to isolate different group of containerized workloads) 
- [x] Log unsuccessful connection attempts

## Secure the Orchestrator Configuration

- [x] Implement version control for orchestrator service definitions and configurations (using git) for auditing
- [x] Ensure cluster-level policies (such as security policies, network policies, and so on) go through your change request, review, and approval process
- [x] Implement orchestrator API access security using role-based access control (RBAC) and network policies
- [x] Be aware of which third-party plugins (such as Container Network Interfaces [CNIs], Container Storage Interface [CSIs], and Container Runtime Interfaces [CRIs]) are running (binary/DaemonSet/controller), what access they have, and whether they are running as privileged containers
- [x] Control access to the orchestrator control plane APIs from third-party plugins using RBAC and service accounts
- [x] Enable access logs for all API requests to the orchestrator control plane (for example, audit logs in Kubernetes)  
- [x] Scan orchestrator manifests for containerized apps (such as Kubernetes deployment manifests) for security best practices and applicable compliance standards in the CI phase 
- [x] If you have any sensitive configuration information in your cluster that needs to be accessed by containers at runtime make sure the configuration is encrypted (such as encrypted secrets in kubernetes)
- [x] Rorate encryption keys that are used for communication between orchestrator components (for example kubernetes API server and etcd)


## Secure the Cloud Environment

- [x] If you’re running your containers in a cloud, remember the default security configuration (for orchestrator, container runtime, and operating system) can be different for different cloud providers 
- [x] Understand the version of orchestrator and container runtime components your cloud provider is running by default, and whether those components are modified from their open source version
- [x] Scan environment deployment configurations (such as Terraform, Cloud Formation templates, and Azure ARM templates) for security best practices and compliance misalignments 


## Secure the Data

- [x] Use a proper filesystem encryption technology for container storage
- [x] Provide write/execute access only to the containers that need to modify the data in a specific host filesystem path
- [x] Reduce write/execute filesystem access for the host filesystem to a minimum using constructs like Pod Security Policy (for example, only allowing Read-only Root Filesystem access, listing allowed host filesystem paths to mount, and listing allowed Flex volume drivers) 
- [x] Automatically scan container images for sensitive data such as tokens, private keys, and so on, before pushing them to a container registry (can be done locally and in CI) 
- [x] Limit storage related syscalls and capabilities to prevent runtime privilege escalation 
- [x] Log all successful and unsuccessful attempts to access sensitive data  


# Other links

- Prisma Free API Tutorial: [Prisma_Scanning_Tutorial.pdf](https://github.com/gunjan5/container-security/blob/master/Prisma%20Scanning%20Tutorial.pdf)

- CircleCI CI build example: [https://circleci.com/gh/gunjan5/container-security](https://circleci.com/gh/gunjan5/container-security)

- CircleCI config: [config.yml](https://github.com/gunjan5/container-security/blob/master/.circleci/config.yml)

- [VS Code](https://code.visualstudio.com/download) IDE Plugin config:

    - Install [Save and Run](https://marketplace.visualstudio.com/items?itemName=wk-j.save-and-run)
    - Edit config and add this config:[vscode-settings.json](https://github.com/gunjan5/container-security/blob/master/vscode-settings.json)
