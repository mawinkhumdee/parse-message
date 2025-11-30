# parse-message

Messaging pipeline that ingests raw user text, parses fields with Gemini, stores results in MongoDB, and updates message status through Kafka.

## Concept
- gRPC API (`InsertMessage`) accepts a message (user, content, source).
- Service stores the message in MongoDB, then publishes a JSON payload to the `parse-message` Kafka topic with `id`, `content`, `user_id`, and `source`.
- Parse consumer calls Gemini to extract fields and intent, saves a `ParseResult`, and publishes an update payload to the `update-message` topic.
- Update consumer reads the update payload and marks the original message status in MongoDB.

## Prerequisites
- Docker and Docker Compose (recommended), or Go 1.25+ with access to Kafka and MongoDB.
- Gemini API key if you want real parsing (otherwise stub Gemini as needed).

## Run with Docker Compose
```bash
docker compose up -d --build
```
- Services: Kafka (`kafka:9092`), Zookeeper, MongoDB (`mongo:27017`), and the app (`:50051`).
- Configs are loaded from `resources/config.master.yml` and `resources/config.secret.yml`. Update `gemini.api-key` in `resources/config.secret.yml` as needed.

## Run locally (without Docker for the app)
1. Start Kafka, Zookeeper, and MongoDB (e.g., `docker compose up kafka zookeeper mongo`).
2. Ensure `resources/config.master.yml` points to those hosts (defaults already match the compose network).
3. Run the service:
   ```bash
   go run ./main
   ```

## gRPC usage
The gRPC server listens on `:50051`.

Example `grpcurl` call:
```bash
grpcurl -plaintext -d '{
  "user_id": "u123",
  "content": "Bought coffee 120 baht at Amazon Cafe",
  "source": "text"
}' localhost:50051 proto.MessageService/InsertMessage
```
Response: `{"success":true,"id":"pending"}`

## Kafka payloads
- Parse topic (`parse-message`): JSON sent by the app
  ```json
  {"id":"<message_id>","content":"Bought coffee 120 baht","user_id":"u123","source":"text"}
  ```
- Update topic (`update-message`): produced by the parser
  ```json
  {"id":"<message_id>","status":"success"}
  ```

## Notes
- Mongo connection: `mongodb://vector:secretvector@mongo:27017/?authSource=admin` (see `resources/config.secret.yml`).
- Topics auto-create is enabled in the provided Kafka config.
