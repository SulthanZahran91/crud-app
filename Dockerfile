# syntax=docker/dockerfile:1

FROM node:20-alpine AS frontend-build
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ ./
RUN npm run build

FROM golang:1.24-alpine AS backend-build
RUN apk add --no-cache build-base
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
COPY --from=frontend-build /app/frontend/dist ./frontend/dist
RUN CGO_ENABLED=1 go build -o /app/bin/app .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates sqlite
WORKDIR /app
COPY --from=backend-build /app/bin/app ./app
COPY --from=frontend-build /app/frontend/dist ./frontend/dist
RUN mkdir -p /app/data
ENV PORT=8080
ENV DB_PATH=/app/data/app.db
EXPOSE 8080
CMD ["./app"]
