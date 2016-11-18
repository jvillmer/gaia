{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Configuration stores information needed to authenticate an user using any servers like LDAP/Google/Certificate ",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": "free",
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "configuration",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "setter": false,
            "stored": true,
            "subtype": "auth_config",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Certificate",
                "Key",
                "LDAP",
                "OAUTH"
            ],
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Method for authenticator (Certificate, Key, LDAP, Google, etc)",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "method",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Name of the authenticator",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "free",
            "getter": true,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "name",
            "orderable": true,
            "primary_key": true,
            "read_only": false,
            "required": true,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": true,
            "delete": false,
            "deprecated": null,
            "get": true,
            "relationship": "child",
            "rest_name": "user",
            "update": false
        }
    ],
    "model": {
        "create": false,
        "delete": true,
        "description": "Authenticator defines all the configuration needed to authenticate an user.",
        "entity_name": "Authenticator",
        "extends": [
            "@base",
            "@described",
            "@identifiable-nopk-stored",
            "@named"
        ],
        "get": true,
        "package": "Infrastructure",
        "resource_name": "authenticators",
        "rest_name": "authenticator",
        "root": null,
        "update": true
    }
}
