# BotNana

A VDA5050-compliant AMR fleet engine built in Go. It uses MQTT for robot traffic and REST for human-facing workflows and integrations.

**Status**: Work in progress. Not production-ready.

## Details
For a more detailed specification of how and why this came to be, have a look at the readme-original.md.

## Overview
BotNana is a backend-first fleet platform that speaks VDA5050 over MQTT and exposes a REST API for UI, reporting, and external integrations. The design separates robot messaging from user workflows so the fleet can function even without a UI layer.

## Key ideas
- VDA5050 over MQTT for robot communication
- REST API for operator workflows and external systems
- Adapter approach for non-VDA5050 robots (vendor-to-VDA5050 mappers)
- Strong observability and structured event storage

## Architecture (high level)
- **MQTT layer**: VDA5050 message handling and robot state
- **REST layer**: Human-facing API and orchestration
- **Datastores**:
  - Relational DB for users, orders, and templates (default: MariaDB/MySQL)
  - Document/observability store for events and telemetry (default: Elasticsearch)
- **UI**: Companion frontend (separate repo)

## Features (current)
- MQTT message ingestion and inspection
- Active robot discovery and in-memory state
- REST endpoints for basic fleet visibility
- Map storage in the relational DB

## Features (planned)
- Full VDA5050 coverage
- Adapter SDK for non-compliant robots
- Fleet-to-fleet interoperability
- CI integration tests with MQTT recordings
- Simulator-driven regression tests

## Requirements
- Go (1.21+ recommended)
- Docker + Docker Compose
- An MQTT broker (included in docker-compose)
- MariaDB/MySQL (included in docker-compose)
- Elasticsearch + Kibana (included in docker-compose)

## Getting started
1) Copy config files:
   - `./config/.env.example` → `./.env`
   - `./config/example.yaml` → `./config/development.yaml`
2) Start the dev stack:
   - `docker compose up -d`
3) Run the backend:
   - `go run main.go -e development`

### Full stack profile
If you want the heavier dev stack (set the correct `.env` values first):
- `docker compose --profile full-masters-cluster up -d`

## Configuration
BotNana reads configuration from two places:

- `.env` (used by Docker Compose only)
  - Database credentials
  - Elastic stack version/ports
  - Kibana/Fleet settings
- `config/development.yaml` (used by the Go app)
  - `mqttBroker`, `apiPort`, `appUrl`, `keycloakUrl`, `OAuthUrl`
  - `mysql.*` credentials
  - `logging.*` toggles
  - `addTestData` to seed a test map from `assets/maps/*` when the DB is empty

If you prefer a different environment name, run `go run main.go -e <name>` and create `config/<name>.yaml`.

## REST API
- Base path: `/api/v1`
- Auth: Keycloak (see setup below)
- Swagger UI: `/swagger/index.html`
- Specs: `docs/swagger.yaml` and `docs/swagger.json`

### Resource groups
- `amrs`: `/amrs/all`, `/amrs/positiondata`, `/amrs/info`
- `maps`: `/maps/all`, `/maps/:mapID`, `POST /maps`
- `nodes`: `/nodes/all`, `/nodes/all/:mapid`, `POST /nodes`
- `edges`: `/edges/all`, `POST /edges`
- `actions`: `/actions/all`, `POST /actions`, `GET /actions/allactionparameters`, `POST /actions/actionparameter`, `GET /actions/allinstantactions`, `POST /actions/instantaction`
- `orders`: `/orders/all`, `POST /orders`, `POST /orders/assign`

## MQTT topics
Subscribed topics:
- `state` (robot state updates)
- `connection` (robot connection status)

Published topics:
- `order` (assign a VDA5050 order)
- `instantAction` (publish instant actions)

## Authentication (Keycloak)
Keycloak is the default OAuth provider. You can use another provider if you prefer.

1) Admin UI: `http://localhost:7080`
2) Create a realm named `botnana`
3) Import `docs/realm-export.json` or configure manually based on `docs/realm-settings.png`
4) Set `clientID` and `clientSecret` in `config/development.yaml`
5) Create a test user

## Observability
Kibana and APM are included for observability.
- Kibana: `https://localhost:5601`
- Create an APM integration with default values and `https://localhost`
- Enroll the APM agent policy in the Fleet server agent

## Services and ports (default)
- MQTT broker: `localhost:1883` (websocket `:9001`)
- MariaDB: `localhost:23312`
- phpMyAdmin: `http://localhost:8183`
- Keycloak: `http://localhost:7080`
- REST API: `http://localhost:8002`
- Kibana: `https://localhost:5601`

## Project layout
- `main.go`: App entrypoint
- `config/`: App config and examples
- `configs/`: Elastic and HAProxy configs
- `controllers/`: MQTT and REST controllers
- `routes/`: REST and MQTT routing
- `models/`: GORM models and DB setup
- `conn/`: Database and Elastic clients
- `docs/`: Swagger specs and Keycloak realm export
- `mqtt-broker/`: Mosquitto configs
- `docker/`: Elastic/Kibana/Fleet helper configs
- `assets/`: Maps and other static data

## Frontend
The reference UI lives in a separate repo:
- `https://github.com/TobiasFP/BananaUI`

## Business model
- Core backend: Open source (MIT)
- UI and simulator: Source-available, commercial licensing by agreement
- Paid offering: Integration, support, and custom adapter development

## Roadmap (near term)
- Complete basic VDA5050 message flow
- Reference adapter for one non-compliant robot
- Integration tests with MQTT recordings
- Minimal operator workflow in the UI

## Contributing
BotNana is early-stage. Feedback, issues, and PRs are welcome. If you want to integrate your own robots or build adapters, open an issue to align on the interface.

## License
MIT for the backend core. Separate terms apply to the UI and simulator.

## Links
- VDA5050 spec: `https://github.com/VDA5050/VDA5050/blob/main/VDA5050_EN.md`
- Maps (PGM): `https://netpbm.sourceforge.net/doc/pgm.html`
- SIEM overview: `https://en.wikipedia.org/wiki/Security_information_and_event_management`
