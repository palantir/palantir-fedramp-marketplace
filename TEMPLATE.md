# {{.serviceIdentification.serviceName}}

**{{.serviceIdentification.providerName}}**

{{.serviceIdentification.serviceDescription}}

## Offering

| Field | Value |
|---|---|
| Certification | {{.serviceIdentification.certificationType}} |
| FedRAMP ID | {{.serviceIdentification.fedRampPackageId}} |
{{if (index .serviceIdentification "ueiNumber")}}| UEI | {{.serviceIdentification.ueiNumber}} |
{{end}}| Service acronym | {{.serviceIdentification.serviceAcronym}} |
| Service model | {{join .serviceProperties.serviceType ", "}} |
| Deployment model | {{.serviceProperties.deploymentModel}} |
| Product website | [{{.serviceIdentification.website}}]({{.serviceIdentification.website}}) |

## Business categories

{{join .serviceProperties.businessCategory " · "}}

## Contacts and assessor

| Role | Contact |
|---|---|
{{range $contact := .contactInformation}}| {{$contact.contactType}} | {{$contact.contactName}} · [{{$contact.contactEmail}}](mailto:{{$contact.contactEmail}}) |
{{end}}| Independent assessor | {{.assessor.name}} · FedRAMP ID {{.assessor.assessorID}} |

## Certified services

{{len .certifiedServices}} services are included in this certification. Services not listed here are outside the FedRAMP Minimum Assessment Scope.

| Service | Available | Category |
|---|---|---|
{{range $service := sortServices .certifiedServices}}| {{$service.serviceName}} | {{with (index $service "dateAvailable")}}{{.}}{{else}}Not specified{{end}} | {{$service.securityCategory}} |
{{end}}
**About service dates**

> {{.certifiedServicesNote.dateAvailableBasis}}

## Certified service details

{{range $service := sortServices .certifiedServices}}### {{$service.serviceName}}

**FIPS 199 {{$service.securityCategory}}{{with (index $service "dateAvailable")}} · Available since {{.}}{{end}}**

{{$service.serviceDescription}}
{{if or (index $service "providerWebsite") (index $service "securityAdminGuideUrl") (index $service "supplementalDocuments")}}
{{if (index $service "providerWebsite")}}[Provider website]({{$service.providerWebsite}}){{end}}{{if (index $service "securityAdminGuideUrl")}} · [Security administration guide]({{$service.securityAdminGuideUrl}}){{end}}{{if (index $service "supplementalDocuments")}} · [Supplemental documents]({{$service.supplementalDocuments}}){{end}}{{end}}
{{end}}{{if index . "servicesOutsideMinimumAssessmentScope"}}## Services outside the minimum assessment scope

{{range $service := .servicesOutsideMinimumAssessmentScope}}### {{$service.serviceName}}

{{$service.reason}}

{{if (index $service "supplementalDocuments")}}[Supplemental documents]({{$service.supplementalDocuments}}){{end}}
{{end}}{{end}}## Trust Center

| Field | Value |
|---|---|
| Repository | [{{.serviceProperties.trustCenter.url}}]({{.serviceProperties.trustCenter.url}}) |
| Type | {{join .serviceProperties.trustCenter.repositoryType ", "}} |
| Description | {{.serviceProperties.trustCenter.repositoryDescription}} |
| Authentication required | {{if .serviceProperties.trustCenter.authenticationRequired}}Yes{{else}}No{{end}} |
{{if (index .serviceProperties.trustCenter "accessRequestInstructions")}}| Access instructions | {{.serviceProperties.trustCenter.accessRequestInstructions}} |
{{end}}
## Secure Configuration Guidance

| Field | Value |
|---|---|
| Repository | [{{.serviceProperties.secureConfigurationGuidance.url}}]({{.serviceProperties.secureConfigurationGuidance.url}}) |
| Type | {{join .serviceProperties.secureConfigurationGuidance.repositoryType ", "}} |
| Description | {{.serviceProperties.secureConfigurationGuidance.repositoryDescription}} |
| Authentication required | {{if .serviceProperties.secureConfigurationGuidance.authenticationRequired}}Yes{{else}}No{{end}} |
{{if (index .serviceProperties.secureConfigurationGuidance "accessRequestInstructions")}}| Access instructions | {{.serviceProperties.secureConfigurationGuidance.accessRequestInstructions}} |
{{end}}
## Continuous monitoring

| Field | Value |
|---|---|
| Next Ongoing Certification Report | {{.serviceProperties.nextOngoingCertificationReportDate}} |
| Next Quarterly Review | {{.serviceProperties.nextQuarterlyReview.date}}, {{.serviceProperties.nextQuarterlyReview.startTime}} |
| Quarterly Review registration | [Request an invitation]({{.serviceProperties.nextQuarterlyReview.registrationUrl}}). {{.serviceProperties.nextQuarterlyReview.registrationInstructions}} |
| OCR feedback and questions | [Email the {{.serviceIdentification.serviceAcronym}} FedRAMP ISSO]({{.serviceProperties.ongoingCertificationReportFeedbackUrl}}). {{.serviceProperties.ongoingCertificationReportFeedbackInstructions}} |

## Certification Data documentation

Availability describes access to artifacts, not certification status.

| Document | Formats | Availability |
|---|---|---|
{{range $item := .documentationOverview}}| {{if (index $item "url")}}[{{$item.name}}]({{$item.url}}){{else}}{{$item.name}}{{end}}{{if (index $item "rule")}} ({{$item.rule}}){{end}} | {{join $item.formats ", "}} | {{$item.availability}} |
{{end}}
---

Version {{.metadata.version}} · last updated {{.metadata.lastUpdated}} · source {{.metadata.sourceOfUpdate}} · responsible official {{.metadata.responsibleAccountableOfficial.name}} ({{.metadata.responsibleAccountableOfficial.title}}, [{{.metadata.responsibleAccountableOfficial.email}}](mailto:{{.metadata.responsibleAccountableOfficial.email}})).
