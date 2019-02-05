# Model
model:
  rest_name: x509certificate
  resource_name: x509certificates
  entity_name: X509Certificate
  package: barret
  group: internal/x509
  description: This API allows to retrieve an client certifcate for api authentication.
  private: true

# Attributes
attributes:
  v1:
  - name: CSR
    description: CSR contains the Certificate Signing Request as a PEM encoded string.
    type: string
    exposed: true
    required: true
    creation_only: true
    example_value: |-
      -----BEGIN CERTIFICATE REQUEST-----
      MIICvDCCAaQCAQAwdzELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxDzANBgNV
      BAcMBkxpbmRvbjEWMBQGA1UECgwNRGlnaUNlcnQgSW5jLjERMA8GA1UECwwIRGln
      aUNlcnQxHTAbBgNVBAMMFGV4YW1wbGUuZGlnaWNlcnQuY29tMIIBIjANBgkqhkiG
      9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8+To7d+2kPWeBv/orU3LVbJwDrSQbeKamCmo
      wp5bqDxIwV20zqRb7APUOKYoVEFFOEQs6T6gImnIolhbiH6m4zgZ/CPvWBOkZc+c
      1Po2EmvBz+AD5sBdT5kzGQA6NbWyZGldxRthNLOs1efOhdnWFuhI162qmcflgpiI
      WDuwq4C9f+YkeJhNn9dF5+owm8cOQmDrV8NNdiTqin8q3qYAHHJRW28glJUCZkTZ
      wIaSR6crBQ8TbYNE0dc+Caa3DOIkz1EOsHWzTx+n0zKfqcbgXi4DJx+C1bjptYPR
      BPZL8DAeWuA8ebudVT44yEp82G96/Ggcf7F33xMxe0yc+Xa6owIDAQABoAAwDQYJ
      KoZIhvcNAQEFBQADggEBAB0kcrFccSmFDmxox0Ne01UIqSsDqHgL+XmHTXJwre6D
      hJSZwbvEtOK0G3+dr4Fs11WuUNt5qcLsx5a8uk4G6AKHMzuhLsJ7XZjgmQXGECpY
      Q4mC3yT3ZoCGpIXbw+iP3lmEEXgaQL0Tx5LFl/okKbKYwIqNiyKWOMj7ZR/wxWg/
      ZDGRs55xuoeLDJ/ZRFf9bI+IaCUd1YrfYcHIl3G87Av+r49YVwqRDT0VDV7uLgqn
      29XI1PpVUNCPQGn9p/eX6Qo7vpDaPybRtA2R7XLKjQaF9oXWeCUqy1hvJac9QFO2
      97Ob1alpHPoZ7mWiEuJwjBPii6a9M9G30nUo39lBi1w=
      -----END CERTIFICATE REQUEST-----

  - name: ID
    description: ID contains the identifier of the certificate.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    filterable: true
    identifier: true

  - name: certificate
    description: Certificate contains the certificate data in PEM format.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: expirationDate
    description: ExpirationDate contains the requested expiration date.
    type: time
    exposed: true
    subtype: string
    creation_only: true

  - name: extensions
    description: |-
      Extensions is a list of extensions that can be added as SAN extensions to the
      certificate.
    type: list
    exposed: true
    subtype: string
    creation_only: true

  - name: signer
    description: Selects what CA should sign the certificate.
    type: enum
    exposed: true
    allowed_choices:
    - Public
    - System
    default_value: Public

  - name: unrevocable
    description: |-
      If set to true, the certificate is considered short lived and it will not be
      possible to revoke it.
    type: boolean
    exposed: true

  - name: usage
    description: Usage defines the requested key usage.
    type: enum
    exposed: true
    allowed_choices:
    - Client
    - Server
    - ServerClient
    default_value: Client
