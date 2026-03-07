# ⚽ Soccer Club Comparator

> AI-powered historical soccer club comparison. Pick two clubs, specify the year, and get a detailed head-to-head analysis powered by LLaMA 3.3 via Groq.

**Example:** São Paulo FC 2005 vs Internacional 2007

---

## Tech Stack

| Layer | Technology |
|---|---|
| Backend | Go + Gin |
| AI | Groq API — LLaMA 3.3 70B |
| Styling | Tailwind CSS |

---

## Features

- Compare any two soccer clubs from any year
- AI-generated analysis covering squad, coach, tactics, achievements, and a head-to-head verdict
- Clean REST API built in Go
- Fast inference via Groq (free tier supported)

---

## Project Structure

```
soccer-compare/
├── backend/           # Go + Gin REST API
│   ├── main.go
│   ├── handlers/
│   └── services/
```

---

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js 18+
- A free [Groq API key](https://console.groq.com)

### Backend

```bash
cd backend
cp .env.example .env
# Add your GROQ_API_KEY to .env

go mod tidy
make start
```

The API will be available at `http://localhost:8080`.


## API

### `GET /api/compare`

| Param | Type | Example |
|---|---|---|
| `team1` | string | `Sao Paulo FC 2005` |
| `team2` | string | `Internacional 2007` |

**Example request:**

```bash
curl "http://localhost:8080/api/compare?team1=Sao+Paulo+FC+2005&team2=Internacional+2007"
```

**Example response:**

```json
{
  "team1": "Sao Paulo FC 2005",
  "team2": "Internacional 2007",
  "result": "## 1. Squad & Key Players\n..."
}
```

---

## Environment Variables

### `backend/.env`

```
GROQ_API_KEY=your_key_here
PORT=8080
```

---