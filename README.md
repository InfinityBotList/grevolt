# Grevolt

Package ``grevolt`` provides golang typings for the revolt api.

## Type generation datasheet

To generate/update types for Revolt API, use SwaggerHub (or ``swagger-codegen``) to import https://api.revolt.chat/openapi.json, export the SDK, then concatenate all the ``model_*.go`` files into one file, and remove the ``package`` declarations and comments:

See ``src-openapi/parse.sh`` for an script that will automate all of this for you with patching as well

**Heavy work in progress**