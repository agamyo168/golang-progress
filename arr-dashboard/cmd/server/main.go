package main

import (
	"arr-dashboard/internal/config"
	"arr-dashboard/internal/handler"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := config.EnsureConfig(); err != nil {
		log.Printf("Warning: Failed to create config: %v", err)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/services", handler.Services)
	http.HandleFunc("/api/refresh", handler.Refresh)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("arr-dashboard running on http://0.0.0.0:%s\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>arr-dashboard</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: #1a1a2e;
            color: #eee;
            min-height: 100vh;
        }
        .container {
            max-width: 900px;
            margin: 0 auto;
            padding: 2rem;
        }
        header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 2rem;
            padding-bottom: 1rem;
            border-bottom: 1px solid #333;
        }
        h1 {
            font-size: 1.5rem;
            font-weight: 600;
        }
        .refresh-btn {
            background: #4361ee;
            color: white;
            border: none;
            padding: 0.5rem 1rem;
            border-radius: 6px;
            cursor: pointer;
            font-size: 0.9rem;
            transition: background 0.2s;
        }
        .refresh-btn:hover {
            background: #3a56d4;
        }
        .refresh-btn:disabled {
            background: #555;
            cursor: not-allowed;
        }
        .services-grid {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
            gap: 1rem;
        }
        .service-card {
            background: #16213e;
            border-radius: 8px;
            padding: 1.25rem;
            display: flex;
            align-items: center;
            gap: 1rem;
            transition: transform 0.2s, box-shadow 0.2s;
            text-decoration: none;
            color: inherit;
        }
        .service-card:hover {
            transform: translateY(-2px);
            box-shadow: 0 4px 12px rgba(0,0,0,0.3);
        }
        .status-dot {
            width: 12px;
            height: 12px;
            border-radius: 50%;
            flex-shrink: 0;
        }
        .status-dot.active {
            background: #10b981;
            box-shadow: 0 0 8px #10b981;
        }
        .status-dot.inactive {
            background: #ef4444;
            box-shadow: 0 0 8px #ef4444;
        }
        .service-info {
            flex: 1;
        }
        .service-name {
            font-weight: 500;
            margin-bottom: 0.25rem;
        }
        .service-url {
            font-size: 0.8rem;
            color: #888;
        }
        .last-refresh {
            text-align: center;
            margin-top: 2rem;
            color: #666;
            font-size: 0.85rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>arr-dashboard</h1>
            <button class="refresh-btn" onclick="refresh()">Refresh</button>
        </header>
        <div class="services-grid" id="services"></div>
        <div class="last-refresh" id="lastRefresh"></div>
    </div>

    <script>
        let services = [];

        async function fetchServices() {
            const res = await fetch('/api/services');
            const data = await res.json();
            services = data.services;
            render();
        }

        function render() {
            const container = document.getElementById('services');
            container.innerHTML = services.map(svc => \`
                <a href="\${svc.url}" target="_blank" class="service-card">
                    <div class="status-dot \${svc.status}"></div>
                    <div class="service-info">
                        <div class="service-name">\${svc.name}</div>
                        <div class="service-url">\${svc.url}</div>
                    </div>
                </a>
            \`).join('');

            const refreshEl = document.getElementById('lastRefresh');
            refreshEl.textContent = services.length > 0 ? 'Last updated: ' + new Date().toLocaleTimeString() : '';
        }

        async function refresh() {
            const btn = document.querySelector('.refresh-btn');
            btn.disabled = true;
            btn.textContent = 'Refreshing...';
            
            try {
                await fetch('/api/refresh', { method: 'POST' });
                await fetchServices();
            } finally {
                btn.disabled = false;
                btn.textContent = 'Refresh';
            }
        }

        fetchServices();
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
