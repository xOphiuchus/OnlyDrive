### MOVED TO AN ORGANIZATION NAME IS AxolotlDrive (same thing with different name Minecraft themed!!)

# OnlyDrive

An open-source, self-hosted Google Drive alternative built with Go.

> **NOTE:** OnlyDrive is a complete rewrite of RustyDrive in Go. We realized that while Rust provides excellent performance, Go offers a more accessible entry point for contributors and is perfectly suited for self-hosted applications.

## Description

OnlyDrive is a lightweight, privacy-focused file storage and synchronization solution designed to give you complete control over your data. Host it yourself, own your files, and enjoy a seamless cloud storage experience without relying on third-party services.

## Features

- 📁 File upload and download
- 🔄 File synchronization across devices
- 🔒 Self-hosted and privacy-first
- ⚡ Built with Go for simplicity and performance
- 🌐 RESTful API with clean architecture
- 👥 Multi-user support
- 💚 Well-loved by the community
- 🪶 Lightweight, fast, and easy to deploy
- 🔐 Rate limiting and CORS support
- 📊 Comprehensive logging with Zerolog

## Tech Stack

- **Backend:** Go, Fiber (web framework)
- **Database:** PostgreSQL
- **Logger:** Zerolog
- **ORM:** GORM
- **API:** RESTful with v1 versioning
- **Client** Sveltekit, Native Android

## Getting Started

### Prerequisites

- Go 1.25+
- PostgreSQL 17+
- Docker & Docker Compose (optional)

### Installation

```bash
git clone https://github.com/xOphiuchus/OnlyDrive.git
cd OnlyDrive
go mod download
```

### Environment Setup

Create a `.env` file in the root directory:

```env
APP_ENV=development
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=onlydrive
DB_PASS=your_secure_password
DB_NAME=onlydrive_dev
LOG_LEVEL=debug
```

### Running

**Development:**

```bash
go run ./cmd/main.go
```

**Docker:**

```bash
docker-compose up -d
```

Visit `http://localhost:8080` in your browser.

### Health Check

```bash
curl http://localhost:8080/api/v1/healthz
```

## Configuration

All configuration is managed through environment variables. See `.env.development` for available options.

- `APP_ENV`: Application environment (development/production)
- `APP_PORT`: Server port (default: 8080)
- `DB_HOST`: PostgreSQL host
- `DB_PORT`: PostgreSQL port (default: 5432)
- `DB_USER`: Database user
- `DB_PASS`: Database password
- `DB_NAME`: Database name
- `LOG_LEVEL`: Logging level (debug/info/warn/error)

## API Endpoints

All endpoints are prefixed with `/api/v1`

- `GET /healthz` - Health check endpoint

More endpoints coming soon!

## Testing

```bash
go test ./...
```

## Contributing

Contributions are welcome! We love community involvement. Feel free to:

- Submit pull requests
- Open issues for bugs and feature requests
- Suggest improvements
- Help with documentation

Please ensure your code follows Go conventions and includes tests.

## License

This project is open source and licensed under the MIT License. See the LICENSE file for details.

## Author

**xOphiuchus**

Made with ❤️ for the open source community.

---

**Why Go instead of Rust?**

While Rust provides excellent performance and safety guarantees, Go offers a more accessible learning curve and is perfectly suited for self-hosted applications. Go's simplicity and fast compilation make it ideal for contributors and maintainers, allowing the community to grow and collaborate more effectively.
