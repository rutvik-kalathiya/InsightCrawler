# 🕵️‍♂️ InsightCrawler

InsightCrawler is a full-stack web application that analyzes any website URL and extracts structured metadata such as HTML version, headings, links, broken links, and presence of login forms. It provides a real-time dashboard to view analysis results and manage URLs.

---

## ✨ Features

- ✅ Submit any website URL for analysis
- 📊 View metadata: HTML version, title, headings, internal/external links
- 🚫 Detect broken links (4xx/5xx)
- 🔐 Detect login forms
- 📈 Dashboard with paginated, sortable, filterable table
- 📤 Bulk reanalyze or delete URLs
- ⚡ Real-time status updates (queued → running → done / error)

---

## 🧰 Tech Stack

### Frontend
- [React](https://reactjs.org/) + [TypeScript](https://www.typescriptlang.org/)
- [Vite](https://vitejs.dev/) for fast builds
- [Tailwind CSS](https://tailwindcss.com/) for styling
- [React Router](https://reactrouter.com/) for routing
- [Chart.js](https://www.chartjs.org/) via `react-chartjs-2` for data visualization

### Backend
- [Go](https://golang.org/) + [Gin](https://gin-gonic.com/)
- [GORM](https://gorm.io/) for ORM
- [MySQL](https://www.mysql.com/) for persistent storage
- [GoQuery](https://github.com/PuerkitoBio/goquery) for HTML crawling

---

## 📦 Project Structure

```
InsightCrawler/
├── backend/                # Go server, crawler logic, API
├── frontend/               # React app
├── docker-compose.yml
├── .env                    # Env vars for local dev
├── README.md
```

---

## 🚀 Getting Started

### Prerequisites
- Node.js (v18+)
- Go (v1.20+)
- Docker & Docker Compose

---

### 1️⃣ Clone the Repo

```bash
git clone https://github.com/yourusername/InsightCrawler.git
cd InsightCrawler
```

### 2️⃣ Run Frontend
```bash
cd frontend
npm install
npm run dev
```
Runs on: http://localhost:5173

### 3️⃣ Run Backend
```bash
cd backend
go run main.go
```
Runs on: http://localhost:8080

### 4️⃣ (Optional) Run with Docker
```bash
docker-compose up --build
```

---

## 🔐 API Authentication
All backend requests require an Authorization token via headers:

```http
Authorization: Bearer <your-token>
```

---

## 🧪 Testing
```bash
cd frontend
npm run test
```
(Backend testing via `go test ./...` coming soon)


## ✍️ Author
Built with ❤️ by Rutvik Kalathiya