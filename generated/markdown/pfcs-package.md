# Palantir Federal Cloud Service

**Palantir Technologies Inc.**

The Palantir Federal Cloud Service (PFCS) is a dedicated environment for the purpose of delivering Palantir software to federal government customers as a cloud service. Palantir software, including Foundry, AIP, Gotham, Apollo, and supporting products utilizing the same infrastructure, enables a multitude of collaborative and operational workflows for end government users. Palantir enables organizations to take advantage of best-in-class Artificial Intelligence, Machine Learning, Data Integration, Data Storage, Data Processing, Analytics, Visualization, Operations, Cybersecurity, and Software Deployment capabilities. The PFCS allows customers to acquire Palantir to quickly deliver value against their hardest problems. PFCS includes deployments on AWS GovCloud, Azure Government, Azure Commercial, Google Cloud Platform (GCP), and AWS US East/West. PFCS combines the formerly separated PFCS High (FR2434554673) and PFCS Moderate (FR1912671248) packages.

## Offering

| Field | Value |
|---|---|
| Certification | Rev5 |
| FedRAMP ID | FR2434554673 |
| UEI | FSY4LVSBGWB7 |
| Service acronym | PFCS |
| Service model | PaaS, SaaS |
| Deployment model | Public Cloud |
| Product website | [https://www.palantir.com/](https://www.palantir.com/) |

## Business categories

Analytics · Collaboration · Cybersecurity & Risk Management · Data Management · Development Tools · Finance · Fleet Management · Law Enforcement · Operations Management · Research

## Contacts and assessor

| Role | Contact |
|---|---|
| Security | FedRAMP ISSO · [fedramp-isso@palantir.com](mailto:fedramp-isso@palantir.com) |
| Sales | Palantir Federal Sales · [fedramp@palantir.com](mailto:fedramp@palantir.com) |
| Independent assessor | Schellman Compliance, LLC · FedRAMP ID 136571 |

## Certified services

7 services are included in this certification. Services not listed here are outside the FedRAMP Minimum Assessment Scope.

| Service | Available | Category |
|---|---|---|
| AIP | 2023-04-07 | High |
| Apollo | 2019-12-18 | High |
| Apollo Agent | 2019-12-18 | High |
| Foundry | 2019-12-18 | High |
| Gotham | 2019-12-18 | High |
| Gotham Thick Client | 2019-12-18 | High |
| Palantir Data Connector | 2019-12-18 | High |

**About service dates**

> dateAvailable identifies when each service was available in PFCS or entered a predecessor PFCS authorization package; it is not the date of the current FedRAMP High certification. PFCS first received FedRAMP Moderate authorization on 2019-12-18 under FR1912671248 and was later certified at the High baseline on 2024-11-19 under FR2434554673. PFCS now combines the formerly separate Moderate and High packages.

## Certified service details

### AIP

**FIPS 199 High · Available since 2023-04-07**

Palantir's Artificial Intelligence Platform (AIP) connects large language models (LLMs) and other AI technologies with customer data and operations within PFCS. AIP provides unified access to a range of open-source, self-hosted, and commercial LLMs through the AIP Model Catalog, and supports customer-connected models through its Bring Your Own Model capability. AIP includes tools for building, deploying, and managing models throughout their lifecycle, with LLM capacity management controls that allow administrators to govern model availability and resource consumption. The platform facilitates the conversion of LLM logic flows into secure, governed automations, with support for staging ontology edits for human review, while end-to-end traceability ensures rigorous auditing of all automation execution and downstream effects.

### Apollo

**FIPS 199 High · Available since 2019-12-18**

Apollo is Palantir's continuous delivery and operations platform, providing a single control layer to deploy, upgrade, monitor, and manage software both for Palantir for PFCS and customers using Apollo to deploy their own applications. Apollo uses a Hub-and-Spoke architecture in which a central Hub environment orchestrates deployments to Spoke environments through an Orchestration Engine. Apollo manages the full deployment lifecycle, including release channel promotion pipelines, ramped rollouts, automated rollbacks, maintenance windows, and bulk release recalls, across cloud, on-premises, and disconnected or air-gapped environments. Apollo also provides real-time vulnerability scanning with SBOM and CVE visibility into deployed services.

### Apollo Agent

**FIPS 199 High · Available since 2019-12-18**

Apollo upgrades components via a local agent which regularly queries the Apollo server to determine if any updates or configuration changes are available. Apollo maintains a catalog of versions of services from the artifact repository and is informed of the current state of service installations by agents running alongside those installations. Apollo applies upgrades subject to defined rules, automatically adjudicates releases by observing performance metrics and error states, and gradually rolls out releases that pass adjudication.

### Foundry

**FIPS 199 High · Available since 2019-12-18**

Foundry is Palantir's data operations platform, providing a unified environment for data connectivity and integration, pipeline development, ontology building, analytics, model development, and application delivery. Foundry ingests and integrates data from diverse sources and formats, transforming it through automated pipelines into a semantic data layer known as the Ontology. Foundry enforces security and governance as an integrated layer across all platform capabilities, with access controls managed through organizations, spaces, projects, roles, users, groups, and markings, and maintains comprehensive audit logging and data lineage tracking across all operations.

### Gotham

**FIPS 199 High · Available since 2019-12-18**

Gotham is Palantir's platform for intelligence analysis, investigative operations, and mission planning. It integrates structured and unstructured data from disparate sources into a unified data asset built on a dynamic ontology of entities, properties, and relationships. Gotham provides an integrated workspace for geospatial mapping, network analysis, and temporal pattern detection, supporting the complete operational lifecycle from intelligence collection and production through mission planning, execution, and after-action review. Every piece of data ingested is tethered to its original source with traceable lineage, and all user and administrator interactions are recorded in audit logs.

### Gotham Thick Client

**FIPS 199 High · Available since 2019-12-18**

Palantir Gotham's Titanium is the workspace for Intel-Ops fusion. Palantir Gotham integrates, enhances, and fuses data from many systems to enable decision making at every echelon and in every domain. Thick clients can be rolled out either through central infrastructure of the customer, or via a download from the Gotham landing page's Palantir Global Launcher link.

### Palantir Data Connector

**FIPS 199 High · Available since 2019-12-18**

The Foundry Data Connector supports connection with all types of source systems, structured or unstructured, and supports batch, micro-batch, or streaming. It supports on-premises and cloud-based sources, including edge devices. Users can customize syncs with querying parameters, retries, and time- or trigger-based schedules, and can monitor sync performance for tuning and latency of data ingestion.

## Trust Center

| Field | Value |
|---|---|
| Repository | [https://pfcsdocs.palantirgov.com/](https://pfcsdocs.palantirgov.com/) |
| Type | Trust Center |
| Description | PFCS Documentation Repository. Access is restricted to verified federal government accounts and to FedRAMP Recognized independent assessment services supporting an assessment of this offering. |
| Authentication required | Yes |
| Access instructions | Federal agencies and FedRAMP Recognized independent assessment services request access by contacting FedRAMP-ISSO@palantir.com with the requesting organization, the point of contact to be provisioned, and the authorization or assessment activity the access supports. Access persists for the duration of that activity. |

## Secure Configuration Guidance

| Field | Value |
|---|---|
| Repository | [https://pfcsdocs.palantirgov.com/workspace/compass/view/ri.compass.main.folder.306e6509-fc47-4ca4-b1b4-bd702dbea683](https://pfcsdocs.palantirgov.com/workspace/compass/view/ri.compass.main.folder.306e6509-fc47-4ca4-b1b4-bd702dbea683) |
| Type | Secure Configuration Guidance |
| Description | PFCS Secure Configuration Guidance, providing an offering-specific index of applicable guidance. The underlying public Palantir product documentation is available at https://www.palantir.com/docs. |
| Authentication required | Yes |
| Access instructions | Request access by contacting FedRAMP-ISSO@palantir.com. |

## Continuous monitoring

| Field | Value |
|---|---|
| Next Ongoing Certification Report | 2026-10-20 |
| Next Quarterly Review | 2026-10-27, 13:30:00-04:00 |
| Quarterly Review registration | [Request an invitation](mailto:FedRAMP-ISSO@palantir.com?subject=PFCS%20Quarterly%20Review%20registration). Request an invitation by email to FedRAMP-ISSO@palantir.com. |
| OCR feedback and questions | [Email the PFCS FedRAMP ISSO](mailto:FedRAMP-ISSO@palantir.com?subject=PFCS%20Ongoing%20Certification%20Report%20feedback). Send feedback or questions about any Ongoing Certification Report by email to FedRAMP-ISSO@palantir.com. Questions and answers are published, anonymized and desensitized, in the Feedback Summary of a subsequent report. |

## Certification Data documentation

Availability describes access to artifacts, not certification status.

| Document | Formats | Availability |
|---|---|---|
| [Public offering information](https://github.com/palantir/palantir-fedramp-marketplace/blob/develop/generated/markdown/pfcs-package.md) (CDS-CSO-PUB) | Markdown, JSON | Public |
| [Secure Configuration Guidance](https://pfcsdocs.palantirgov.com/workspace/compass/view/ri.compass.main.folder.306e6509-fc47-4ca4-b1b4-bd702dbea683) (SCG-CSO-RSC) | Human-readable | Offering-specific index available to authorized parties; underlying product documentation is public. |
| [PFCS Certification Data](https://pfcsdocs.palantirgov.com/) | Human-readable, Machine-readable where required | Available to authorized parties through the PFCS Documentation Repository; additional CR26 materials are added as completed. |

---

Version 1.0.0 · last updated 2026-09-23T01:29:04Z · source https://github.com/palantir/palantir-fedramp-marketplace · responsible official PFCS ISSO (Information System Security Officer, [FedRAMP-ISSO@palantir.com](mailto:FedRAMP-ISSO@palantir.com)).
