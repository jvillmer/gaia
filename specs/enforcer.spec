# Model
model:
  rest_name: enforcer
  resource_name: enforcers
  entity_name: Enforcer
  package: squall
  group: core/enforcer
  description: |-
    An Enforcer contains all parameters associated with a registered enforcer. The
    object is mainly by maintained by the enforcers themselves. Users can read the
    object in order to understand the current status of the enforcers.
  indexes:
  - - :shard
    - zone
    - zHash
  - - namespace
  - - namespace
    - name
  - - namespace
    - normalizedTags
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: q
        description: Backward compat for enforcer <=v2.2.1. does not have any effect.
        type: string
        example_value: no effect

      - name: tag
        description: Backward compat for enforcer <=v2.2.1. does not have any effect.
        type: string
        multiple: true
        example_value: no effect
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: FQDN
    description: FQDN contains the fqdn of the server where the enforcer is running.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: server1.domain.com
    orderable: true

  - name: certificate
    description: Certificate is the certificate of the enforcer.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    orderable: true

  - name: certificateExpirationDate
    description: CertificateExpirationDate is the expiration date of the certificate.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: certificateKey
    description: |-
      CertificateKey is the secret key of the enforcer. Returned only when a enforcer
      is created or the certificate is updated.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    orderable: true

  - name: certificateRequest
    description: |-
      CertificateRequest can be used to generate a certificate from that CSR instead
      of letting the server generate your private key for you.
    type: string
    exposed: true
    transient: true

  - name: certificateRequestEnabled
    description: |-
      If set during creation,the server will not initially generate a certificate. In
      that case you have to provide a valid CSR through certificateRequest during an
      update.
    type: boolean
    exposed: true
    stored: true
    creation_only: true

  - name: collectInfo
    description: CollectInfo indicates to the enforcer it needs to collect information.
    type: boolean
    exposed: true
    stored: true

  - name: collectedInfo
    description: CollectedInfo represents the latest info collected by the enforcer.
    type: external
    exposed: true
    subtype: map_of_string_of_strings
    stored: true

  - name: currentVersion
    description: |-
      CurrentVersion holds the enforcerd binary version that is currently associated
      to this object.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: enforcerProfileID
    description: Contains the ID of the profile used by the instance of enforcerd.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: lastCollectionTime
    description: |-
      LastCollectionTime represents the date and time when the info have been
      collected.
    type: time
    exposed: true
    stored: true

  - name: lastSyncTime
    description: LastSyncTime holds the last heart beat time.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: lastValidHostServices
    description: |-
      LastValidHostServices is a read only attribute that stores the list valid host
      services that have been applied to this enforcer. This list might be different
      from the list retrieved through policy, if the dynamically calculated list leads
      into conflicts.
    type: refList
    subtype: hostservice
    stored: true
    autogenerated: true

  - name: localCA
    description: |-
      LocalCA contains the initial chain of trust for the enforcer. This valud is only
      given when you retrieve a single enforcer.
    type: string
    exposed: true
    autogenerated: true
    transient: true

  - name: machineID
    description: |-
      MachineID holds a unique identifier for every machine as detected by the
      enforcer. It is based on hardware information such as the SMBIOS UUID, MAC
      addresses of interfaces or cloud provider IDs.
    type: string
    exposed: true
    stored: true
    example_value: 3F23E8DF-C56D-45CF-89B8-A867F3956409
    filterable: true

  - name: operationalStatus
    description: OperationalStatus tells the status of the enforcer.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Registered
    - Connected
    - Disconnected
    - Initialized
    - Unknown
    default_value: Registered
    filterable: true

  - name: publicToken
    description: |-
      PublicToken is the public token of the server that will be included in the
      datapath and its signed by the private CA.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    transient: true

  - name: startTime
    description: |-
      startTime holds the time this enforcerd was started. This time-stamp is reported
      by the enforcer and is is preserved across disconnects.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: updateAvailable
    description: Tells if the the version of enforcerd is outdated and should be updated.
    type: boolean
    exposed: true
    stored: true
    orderable: true

# Relations
relations:
- rest_name: enforcerprofile
  get:
    description: Returns the enforcer profile that must be used by an enforcer.

- rest_name: auditprofile
  get:
    description: Returns a list of the audit profiles that must be applied to this
      enforcer.

- rest_name: hostservice
  get:
    description: Returns a list of the host services policies that apply to this enforcer.
    parameters:
      entries:
      - name: appliedServices
        description: Valid when retrieved for a given enforcer and returns the applied
          services.
        type: boolean

      - name: setServices
        description: Instructs the backend to cache the services that were resolved.
          services.
        type: boolean

- rest_name: poke
  get:
    description: Sends a poke empty object. This is used to ensure an enforcer is
      up and running.
    parameters:
      entries:
      - name: cpuload
        description: If set, provides the total cpu usage in percentage of vCPUs.
        type: float
        example_value: 1000

      - name: memory
        description: If set, provides the total resident memory used in bytes.
        type: integer
        example_value: 1000

      - name: processes
        description: If set, defines the number of current processes.
        type: integer
        example_value: 10

      - name: ts
        description: time of report. If not set, local server time will be used.
        type: time
